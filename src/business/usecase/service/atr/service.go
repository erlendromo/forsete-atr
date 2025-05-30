package atr

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	"github.com/erlendromo/forsete-atr/src/business/domain/model"
	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
	imagerepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/image"
	modelrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/model"
	outputrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/output"
	pipelinerepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/pipeline"
	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

type ATRService struct {
	ModelRepo    *modelrepository.ModelRepository
	PipelineRepo *pipelinerepository.PipelineRepository
	ImageRepo    *imagerepository.ImageRepository
	OutputRepo   *outputrepository.OutputRepository
}

func NewATRService(m *modelrepository.ModelRepository, p *pipelinerepository.PipelineRepository, i *imagerepository.ImageRepository, o *outputrepository.OutputRepository) *ATRService {
	return &ATRService{
		ModelRepo:    m,
		PipelineRepo: p,
		ImageRepo:    i,
		OutputRepo:   o,
	}
}

func (a *ATRService) CreatePipeline(ctx context.Context, models []*model.Model) (*pipeline.Pipeline, error) {
	var names []string
	for _, model := range models {
		names = append(names, model.Name)
	}

	pipelineName := strings.Join(names, "_")
	pipeline, err := a.PipelineRepo.RegisterPipeline(ctx, pipelineName, util.PIPELINES_PATH)
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		switch model.ModelTypeID {
		// ModelTypeID = 1 -> ModelType = regionsegmentation (yolo)
		// ModelTypeID = 2 -> ModelType = linesegmentation (yolo)
		case 1, 2:
			pipeline = pipeline.AppendYoloStep(model.Path)
		// ModelTypeID = 3 -> ModelType = textrecognition (trocr)
		case 3:
			pipeline = pipeline.AppendTrOCRStep(model.Path)
		default:
			return nil, err
		}

		if err := a.PipelineRepo.RegisterPipelineModel(ctx, pipeline.ID, model.ID); err != nil {
			return nil, err
		}
	}

	pipeline = pipeline.AppendOrderStep("OrderLines").AppendExportStep("json", util.TEMP_OUTPUTS_PATH)

	if err := pipeline.CreateLocal(); err != nil {
		return nil, err
	}

	return pipeline, nil
}

func (a *ATRService) RunATROnImages(ctx context.Context, pipelineID int, userID uuid.UUID, imageIDs []uuid.UUID) ([]*output.Output, error) {
	outputs := make([]*output.Output, 0)
	for _, imageID := range imageIDs {
		output, err := a.runATROnImage(ctx, pipelineID, imageID, userID)
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, output)
	}

	return outputs, nil
}

// Helper-function to run atr on image
func (a *ATRService) runATROnImage(ctx context.Context, pipelineID int, imageID, userID uuid.UUID) (*output.Output, error) {
	// Get pipeline
	p, err := a.PipelineRepo.PipelineByID(ctx, pipelineID)
	if err != nil {
		return nil, err
	}

	// Get image
	i, err := a.ImageRepo.ImageByID(ctx, imageID, userID)
	if err != nil {
		return nil, err
	}

	// Execute htrflow
	cmd := exec.Command(util.BIN_BASH, util.HTRFLOW_SH_PATH, p.PathToFile(), i.PathToFile())
	if htrflowResponse, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(htrflowResponse))
		return nil, err
	}

	// Register output
	outputPath := path.Join(util.USERS_PATH, userID.String(), util.OUTPUTS)
	o, err := a.OutputRepo.RegisterOutput(ctx, i.Name, "json", outputPath, i.ID, userID)
	if err != nil {
		return nil, err
	}

	// Defer removal of temp-file
	pathToTempFile := fmt.Sprintf("%s/%s/%s.%s", util.TEMP_OUTPUTS_PATH, util.IMAGES, o.ImageID.String(), o.Format)
	defer os.Remove(pathToTempFile)

	// Read data from temp-file
	file, err := os.Open(pathToTempFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	data, err := util.DecodeJSON[output.ATRResponse](file)
	if err != nil {
		return nil, err
	}

	// Recreate file for the user
	if err := o.CreateLocal(data); err != nil {
		return nil, err
	}

	return o, nil
}

func (a *ATRService) UploadImages(ctx context.Context, userID uuid.UUID, fileHeaders []*multipart.FileHeader) ([]*image.Image, error) {
	images := make([]*image.Image, 0)
	errs := make([]error, 0)

	for _, fileHeader := range fileHeaders {
		img, err := a.uploadImage(ctx, userID, fileHeader)
		if err != nil {
			errs = append(errs, err)
		}

		images = append(images, img)
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("unable to upload images: %+v", errs)
	}

	return images, nil
}

// Helper-function to upload image
func (a *ATRService) uploadImage(ctx context.Context, userID uuid.UUID, fileHeader *multipart.FileHeader) (*image.Image, error) {
	originalName := strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename))
	name := strings.ToLower(originalName)
	path := path.Join(util.USERS_PATH, userID.String(), util.IMAGES)

	img, err := a.ImageRepo.RegisterImage(ctx, name, "png", path, userID)
	if err != nil {
		return nil, err
	}

	if err := img.CreateLocal(fileHeader); err != nil {
		if err := a.ImageRepo.DeleteImageByID(ctx, img.ID, userID); err != nil {
			return nil, err
		}

		return nil, err
	}

	return img, nil
}

func (a *ATRService) DeleteImageAndOutputs(ctx context.Context, imageID, userID uuid.UUID) error {
	outputs, err := a.OutputRepo.OutputsByImageID(ctx, imageID, userID)
	for _, output := range outputs {
		if err := output.DeleteLocal(); err != nil {
			return err
		}
	}

	if err := a.OutputRepo.DeleteOutputsByImageID(ctx, imageID, userID); err != nil {
		return err
	}

	image, err := a.ImageRepo.ImageByID(ctx, imageID, userID)
	if err != nil {
		return err
	}

	if err := a.ImageRepo.DeleteImageByID(ctx, imageID, userID); err != nil {
		return err
	}

	if err := image.DeleteLocal(); err != nil {
		return err
	}

	return nil
}

func (a *ATRService) DeleteUserOutputsAndImages(ctx context.Context, userID uuid.UUID) error {
	if err := a.OutputRepo.DeleteUserOutputs(ctx, userID); err != nil {
		return err
	}

	if err := a.ImageRepo.DeleteUserImages(ctx, userID); err != nil {
		return err
	}

	return nil
}

// N.B! This function is not fully implemented.
func (a *ATRService) UploadModel(ctx context.Context, name, path string, model_type_id int, fileHeaders []*multipart.FileHeader) (*model.Model, error) {
	model, err := a.ModelRepo.RegisterModel(ctx, name, path, model_type_id)
	if err != nil {
		return nil, err
	}

	// TODO
	// check if fileHeaders match required files (here???)
	// revert on error?

	if err := model.CreateLocal(fileHeaders); err != nil {
		return nil, err
	}

	return model, nil
}

// Setup function to create pipelines on launch
func (a *ATRService) CreatePipelines(ctx context.Context) ([]*pipeline.Pipeline, error) {
	/*
		regionModels, err := a.ModelRepo.ModelsByType(ctx, util.REGION_SEGMENTATION)
		if err != nil {
			return nil, err
		}
	*/

	lineModels, err := a.ModelRepo.ModelsByType(ctx, util.LINE_SEGMENTATION)
	if err != nil {
		return nil, err
	}

	textModels, err := a.ModelRepo.ModelsByType(ctx, util.TEXT_RECOGNITION)
	if err != nil {
		return nil, err
	}

	var createdPipelines []*pipeline.Pipeline

	// Creates pipelines with all possible model-combinations
	for _, textModel := range textModels {

		// TextModels
		pipeline, err := a.CreatePipeline(ctx, []*model.Model{textModel})
		if err != nil {
			return nil, err
		}

		createdPipelines = append(createdPipelines, pipeline)

		for _, lineModel := range lineModels {

			// LineModels + TextModels
			pipeline, err := a.CreatePipeline(ctx, []*model.Model{lineModel, textModel})
			if err != nil {
				return nil, err
			}

			createdPipelines = append(createdPipelines, pipeline)

			/*
				for _, regionModel := range regionModels {

					// RegionModels + LineModels + TextModels
					pipeline, err := a.CreatePipeline(ctx, []*model.Model{regionModel, lineModel, textModel})
					if err != nil {
						return nil, err
					}

					createdPipelines = append(createdPipelines, pipeline)
				}
			*/
		}
	}

	return createdPipelines, nil
}

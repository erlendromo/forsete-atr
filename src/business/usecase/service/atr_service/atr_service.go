package atrservice

import (
	"context"
	"fmt"
	"mime/multipart"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/erlendromo/forsete-atr/src/business/domain/image"
	"github.com/erlendromo/forsete-atr/src/business/domain/model"
	"github.com/erlendromo/forsete-atr/src/business/domain/output"
	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline"
	modelrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/model_repository"
	outputrepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/output_repository"
	pipelinerepository "github.com/erlendromo/forsete-atr/src/business/usecase/repository/pipeline_repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ATRService struct {
	ModelRepo    *modelrepository.ModelRepository
	PipelineRepo *pipelinerepository.PipelineRepository
	OutputRepo   *outputrepository.OutputRepository
	db           *sqlx.DB
}

func NewATRService(db *sqlx.DB) *ATRService {
	return &ATRService{
		ModelRepo:    modelrepository.NewModelRepository(db),
		PipelineRepo: pipelinerepository.NewPipelineRepository(db),
		OutputRepo:   outputrepository.NewOutputRepository(db),
		db:           db,
	}
}

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

func (a *ATRService) CreatePipelines(ctx context.Context) ([]*pipeline.Pipeline, error) {
	/*
		regionModels, err := a.ModelRepo.ModelsByType(ctx, "regionsegmentation")
		if err != nil {
			return nil, err
		}
	*/

	lineModels, err := a.ModelRepo.ModelsByType(ctx, "linesegmentation")
	if err != nil {
		return nil, err
	}

	textModels, err := a.ModelRepo.ModelsByType(ctx, "textrecognition")
	if err != nil {
		return nil, err
	}

	var createdPipelines []*pipeline.Pipeline

	// Creates pipelines with all possible model-combinations
	for _, textModel := range textModels {
		pipeline, err := a.CreatePipeline(ctx, []*model.Model{textModel})
		if err != nil {
			return nil, err
		}

		createdPipelines = append(createdPipelines, pipeline)

		for _, lineModel := range lineModels {
			pipeline, err := a.CreatePipeline(ctx, []*model.Model{lineModel, textModel})
			if err != nil {
				return nil, err
			}

			createdPipelines = append(createdPipelines, pipeline)

			/*
				for _, regionModel := range regionModels {
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

func (a *ATRService) CreatePipeline(ctx context.Context, models []*model.Model) (*pipeline.Pipeline, error) {
	var names []string
	for _, model := range models {
		names = append(names, model.Name)
	}

	pipelineName := strings.Join(names, "_")
	pipelinePath := path.Join("assets", "pipelines")

	pipeline, err := a.PipelineRepo.RegisterPipeline(ctx, pipelineName, pipelinePath)
	if err != nil {
		return nil, err
	}

	for _, model := range models {
		switch model.ModelTypeID {
		case 1, 2:
			pipeline = pipeline.AppendYoloStep(model.Path)
		case 3:
			pipeline = pipeline.AppendTrOCRStep(model.Path)
		default:
			return nil, err
		}

		_, err := a.PipelineRepo.RegisterPipelineModel(ctx, pipeline.ID, model.ID)
		if err != nil {
			return nil, err
		}
	}

	pipeline = pipeline.AppendOrderStep("OrderLines").AppendExportStep("json", "assets/outputs")

	if err := pipeline.CreateLocal(); err != nil {
		return nil, err
	}

	return pipeline, nil
}

func (a *ATRService) RunATROnImage(ctx context.Context, image *image.Image, pipeline *pipeline.Pipeline, userID uuid.UUID) (*output.Output, error) {
	pathToPipelineFile := fmt.Sprintf("%s/%s.yaml", pipeline.Path, pipeline.Name)
	pathToImageFile := fmt.Sprintf("%s/%s.%s", image.Path, image.ID.String(), image.Format)

	cmd := exec.Command("/bin/bash", "assets/scripts/htrflow.sh", pathToPipelineFile, pathToImageFile)
	if output, err := cmd.CombinedOutput(); err != nil {
		fmt.Println(string(output))
		return nil, err
	}

	outputPath := path.Join("assets", "users", userID.String(), "outputs")

	output, err := a.OutputRepo.RegisterOutput(ctx, image.Name, "json", outputPath, image.ID)
	if err != nil {
		return nil, err
	}

	// Read from temporary file...
	fullPathToFile := fmt.Sprintf("assets/outputs/images/%s.%s", output.ImageID.String(), output.Format)

	resp, err := output.ReadJson(fullPathToFile)
	if err != nil {
		return nil, err
	}

	// Remove temporary created file...
	defer os.Remove(fullPathToFile)

	// Recreate file for the user
	if err := output.CreateLocal(resp); err != nil {
		return nil, err
	}

	return output, nil
}

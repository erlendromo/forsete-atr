package pipeline

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type TipNotePipeline struct {
	Steps []Step
}

func NewTipNotePipeline(regionSegmentationModel, lineSegmentationModel, textRecognitionModel, device string) *TipNotePipeline {
	return &TipNotePipeline{
		Steps: []Step{
			ModelStep{
				StepName: "Segmentation",
				Settings: ModelStepSettings{
					ModelType: "yolo",
					ModelSettings: ModelSettings{
						Model:  regionSegmentationModel,
						Device: device,
					},
				},
			},
			ModelStep{
				StepName: "Segmentation",
				Settings: ModelStepSettings{
					ModelType: "yolo",
					ModelSettings: ModelSettings{
						Model:  lineSegmentationModel,
						Device: device,
					},
				},
			},
			ModelStep{
				StepName: "TextRecognition",
				Settings: ModelStepSettings{
					ModelType: "TrOCR",
					ModelSettings: ModelSettings{
						Model:  textRecognitionModel,
						Device: device,
					},
				},
			},
			OrderStep{
				StepName: "OrderLines",
			},
			ExportStep{
				StepName: "Export",
				Settings: ExportStepSettings{
					Format:      "json",
					Destination: "tmp/outputs",
				},
			},
		},
	}
}

func (t *TipNotePipeline) Encode(destination, filename string) (string, error) {
	filepath := fmt.Sprintf("%s/%s", destination, filename)

	if file, err := os.Open(filepath); err == nil {
		return file.Name(), nil
	}

	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	payload, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}

	if _, err := file.Write(payload); err != nil {
		return "", err
	}

	return file.Name(), nil
}

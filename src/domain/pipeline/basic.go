package pipeline

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type BasicPipeline struct {
	Steps []Step
}

func (b *BasicPipeline) Encode(destination, filename string) (string, error) {
	file, err := os.Create(fmt.Sprintf("%s/%s", destination, filename))
	if err != nil {
		return "", err
	}

	payload, err := yaml.Marshal(b)

	_, err = file.Write(payload)
	if err != nil {
		return "", err
	}

	return file.Name(), nil
}

func NewBasicPipeline(lineSegmentationModel, textRecognitionModel string) Pipeline {
	return &BasicPipeline{
		Steps: []Step{
			ModelStep{
				StepName: "Segmentation",
				Settings: ModelStepSettings{
					ModelType: "yolo",
					ModelSettings: ModelSettings{
						Model:  lineSegmentationModel,
						Device: "cpu",
					},
				},
			},
			ModelStep{
				StepName: "TextRecognition",
				Settings: ModelStepSettings{
					ModelType: "TrOCR",
					ModelSettings: ModelSettings{
						Model:  textRecognitionModel,
						Device: "cpu",
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

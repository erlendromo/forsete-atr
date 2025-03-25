package pipeline

type TipNotePipeline struct {
	Steps []Step
}

func (t *TipNotePipeline) Encode(destination, filename string) (string, error) {
	return "", nil
}

func NewTipNotePipeline(regionSegmentationModel, lineSegmentationModel, textRecognitionModel, device string) Pipeline {
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

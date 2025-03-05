package pipeline

type TipNotePipeline struct {
	Steps []Step
}

func (t *TipNotePipeline) Encode(destination, filename string) (string, error) {
	return "", nil
}

func NewTipNotePipeline(regionSegmentationModel, lineSegmentationModel, textRecognitionModel string) Pipeline {
	return &TipNotePipeline{
		Steps: []Step{
			ModelStep{
				StepName: "RegionSegmentation",
				Settings: ModelStepSettings{
					ModelType: "yolo",
					ModelSettings: ModelSettings{
						Model:  regionSegmentationModel,
						Device: "cpu",
					},
				},
			},
			ModelStep{
				StepName: "LineSegmentation",
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
					Destination: "/tmp/outputs",
				},
			},
		},
	}
}

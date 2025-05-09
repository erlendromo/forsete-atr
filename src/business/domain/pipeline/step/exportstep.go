package step

type ExportStep struct {
	StepName string             `yaml:"step"`
	Settings ExportStepSettings `yaml:"settings"`
}

type ExportStepSettings struct {
	Format string `yaml:"format"`
	Dst    string `yaml:"dest"`
}

func NewExportStep(stepName, format, dst string) *ExportStep {
	return &ExportStep{
		StepName: stepName,
		Settings: ExportStepSettings{
			Format: format,
			Dst:    dst,
		},
	}
}

package pipeline

type Pipeline interface {
	Encode(destination, filename string) (string, error)
}

type Step interface{}

type ModelStep struct {
	StepName string            `yaml:"step"`
	Settings ModelStepSettings `yaml:"settings"`
}

type ModelStepSettings struct {
	ModelType     string        `yaml:"model"`
	ModelSettings ModelSettings `yaml:"model_settings"`
}

type ModelSettings struct {
	Model  string `yaml:"model"`
	Device string `yaml:"device"`
}

type OrderStep struct {
	StepName string `yaml:"step"`
}

type ExportStep struct {
	StepName string             `yaml:"step"`
	Settings ExportStepSettings `yaml:"settings"`
}

type ExportStepSettings struct {
	Format      string `yaml:"format"`
	Destination string `yaml:"dest"`
}

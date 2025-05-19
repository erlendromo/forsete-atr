package step

import "github.com/erlendromo/forsete-atr/src/config"

type ModelStep struct {
	StepName string            `yaml:"step"`
	Settings ModelStepSettings `yaml:"settings"`
}

type ModelStepSettings struct {
	ModelType     string        `yaml:"model"`
	ModelSettings ModelSettings `yaml:"model_settings"`
}

type ModelSettings struct {
	PathToModel string `yaml:"model"`
	Device      string `yaml:"device"`
}

func NewModelStep(stepName, modelType, pathToModel string) *ModelStep {
	return &ModelStep{
		StepName: stepName,
		Settings: ModelStepSettings{
			ModelType: modelType,
			ModelSettings: ModelSettings{
				PathToModel: pathToModel,
				Device:      config.GetConfig().APIConfig().DEVICE,
			},
		},
	}
}

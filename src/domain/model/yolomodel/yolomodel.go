package yolomodel

type YoloModel struct {
	ModelName string `json:"name"`
	modelPath string
	modelType string
}

func NewYoloModel(modelName, modelPath, modelType string) *YoloModel {
	return &YoloModel{
		ModelName: modelName,
		modelPath: modelPath,
		modelType: modelType,
	}
}

func (ym *YoloModel) Name() string {
	return ym.ModelName
}

func (ym *YoloModel) Path() string {
	return ym.modelPath
}

func (ym *YoloModel) Type() string {
	return ym.modelType
}

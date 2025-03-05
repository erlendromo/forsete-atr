package yolomodel

type YoloModel struct {
	modelName string
	modelPath string
	modelType string
}

func NewYoloModel(modelName, modelPath, modelType string) *YoloModel {
	return &YoloModel{
		modelName: modelName,
		modelPath: modelPath,
		modelType: modelType,
	}
}

func (ym *YoloModel) Name() string {
	return ym.modelName
}

func (ym *YoloModel) Path() string {
	return ym.modelPath
}

func (ym *YoloModel) Type() string {
	return ym.modelType
}

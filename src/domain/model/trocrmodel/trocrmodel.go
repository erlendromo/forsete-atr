package trocrmodel

type TrOCRModel struct {
	ModelName string `json:"name"`
	modelPath string
	modelType string
}

func NewTrOCRModel(modelName, modelPath, modelType string) *TrOCRModel {
	return &TrOCRModel{
		ModelName: modelName,
		modelPath: modelPath,
		modelType: modelType,
	}
}

func (tom *TrOCRModel) Name() string {
	return tom.ModelName
}

func (tom *TrOCRModel) Path() string {
	return tom.modelPath
}

func (tom *TrOCRModel) Type() string {
	return tom.modelType
}

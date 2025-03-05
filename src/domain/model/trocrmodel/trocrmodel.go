package trocrmodel

type TrOCRModel struct {
	modelName string
	modelPath string
	modelType string
}

func NewTrOCRModel(modelName, modelPath, modelType string) *TrOCRModel {
	return &TrOCRModel{
		modelName: modelName,
		modelPath: modelPath,
		modelType: modelType,
	}
}

func (tom *TrOCRModel) Name() string {
	return tom.modelName
}

func (tom *TrOCRModel) Path() string {
	return tom.modelPath
}

func (tom *TrOCRModel) Type() string {
	return tom.modelType
}

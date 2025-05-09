package step

type OrderStep struct {
	StepName string `yaml:"step"`
}

func NewOrderStep(stepName string) *OrderStep {
	return &OrderStep{
		StepName: stepName,
	}
}

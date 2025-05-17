package pipeline

import (
	"fmt"
	"os"

	"github.com/erlendromo/forsete-atr/src/business/domain/pipeline/step"
	"gopkg.in/yaml.v3"
)

// Pipeline
//
//	@Summary		Pipeline
//	@Description	Pipeline containing id, name etc
type Pipeline struct {
	ID   int    `db:"id" json:"id" yaml:"-"`
	Name string `db:"name" json:"name" yaml:"-"`
	Path string `db:"path" json:"-" yaml:"-"`

	Steps []step.Step `db:"-" json:"-" yaml:"steps"`
}

func (p *Pipeline) AppendYoloStep(pathToYoloModel string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewModelStep(
		"Segmentation", "yolo", pathToYoloModel),
	)

	return p
}

func (p *Pipeline) AppendTrOCRStep(pathToTrOCRModel string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewModelStep(
		"TextRecognition", "TrOCR", pathToTrOCRModel),
	)

	return p
}

func (p *Pipeline) AppendOrderStep(order string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewOrderStep(order))

	return p
}

func (p *Pipeline) AppendExportStep(format, dst string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewExportStep("Export", format, dst))

	return p
}

func (p *Pipeline) CreateLocal() error {
	if len(p.Steps) < 2 {
		return fmt.Errorf("invalid pipeline configuration, add more steps")
	}

	filepath := fmt.Sprintf("%s/%s.yaml", p.Path, p.Name)
	if _, err := os.Open(filepath); err == nil {
		return nil
	}

	file, err := os.Create(filepath)
	if err != nil {
		return err
	}

	payload, err := yaml.Marshal(p)
	if err != nil {
		return err
	}

	if _, err := file.Write(payload); err != nil {
		return err
	}

	return nil
}

func (p *Pipeline) PathToFile() string {
	return fmt.Sprintf("%s/%s.yaml", p.Path, p.Name)
}

package pipeline

import (
	"fmt"
	"os"
	"strings"

	"github.com/erlendromo/forsete-atr/src/domain/pipeline/step"
	"gopkg.in/yaml.v3"
)

type Pipeline struct {
	Steps     []step.Step `yaml:"steps"`
	device    string
	filename  string
	fileDst   string
	outputDst string
}

func NewPipeline(device, filename string) (*Pipeline, error) {
	if device != "cpu" && device != "cuda" {
		return nil, fmt.Errorf("invalid device '%s', supported devices are cpu and cuda", device)
	}

	if len(filename) < 5 || strings.ContainsAny(filename, "*)(/&%$#!:;><@`´+±€£™©∞§|[]≈…‚") {
		return nil, fmt.Errorf("invalid filename, cannot contain special characters except dash and underscore")
	}

	return &Pipeline{
		Steps:     make([]step.Step, 0),
		device:    device,
		filename:  filename,
		fileDst:   "assets/pipelines",
		outputDst: "assets/outputs",
	}, nil
}

func (p *Pipeline) AppendYoloStep(pathToYoloModel string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewModelStep(
		"Segmentation", "yolo", pathToYoloModel, p.device),
	)

	return p
}

func (p *Pipeline) AppendTrOCRStep(pathToTrOCRModel string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewModelStep(
		"TextRecognition", "TrOCR", pathToTrOCRModel, p.device),
	)

	return p
}

func (p *Pipeline) AppendOrderStep(order string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewOrderStep(order))

	return p
}

func (p *Pipeline) AppendExportStep(format string) *Pipeline {
	p.Steps = append(p.Steps, *step.NewExportStep("Export", format, p.outputDst))

	return p
}

func (p *Pipeline) CreateLocalYaml() (string, error) {
	if len(p.Steps) < 2 {
		return "", fmt.Errorf("missing pipeline steps")
	}

	filepath := fmt.Sprintf("%s/%s.yaml", p.fileDst, p.filename)
	if file, err := os.Open(filepath); err == nil {
		return file.Name(), nil
	}

	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	payload, err := yaml.Marshal(p)
	if err != nil {
		return "", err
	}

	if _, err := file.Write(payload); err != nil {
		return "", err
	}

	return file.Name(), nil
}

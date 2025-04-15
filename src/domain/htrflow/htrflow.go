package htrflow

import (
	"fmt"
	"os"
	"os/exec"
)

type HTRflow struct {
	yamlPath  string
	imagePath string
	resultDst string
}

func NewHTRflow(yamlPath, imagePath, resultDst string) *HTRflow {
	return &HTRflow{
		yamlPath:  yamlPath,
		imagePath: imagePath,
		resultDst: resultDst,
	}
}

func (h *HTRflow) Run() (*os.File, error) {
	cmd := exec.Command("/bin/bash", "assets/scripts/htrflow.sh", h.yamlPath, h.imagePath)
	if output, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("%s", string(output))
	}

	return os.Open(h.resultDst)
}

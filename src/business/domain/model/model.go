package model

import (
	"fmt"
	"mime/multipart"
	"sync"

	"github.com/erlendromo/forsete-atr/src/util"
)

// Model
//
//	@Summary		Model
//	@Description	Model containing id, name etc.
type Model struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Path        string `db:"path" json:"-"`
	ModelTypeID int    `db:"model_type_id" json:"-"`
	ModelType   string `db:"model_type" json:"model_type"`
}

func (m *Model) CreateLocal(fileHeaders []*multipart.FileHeader) error {
	if len(fileHeaders) < 1 {
		return fmt.Errorf("missing model-files")
	}

	var wg sync.WaitGroup
	errs := make([]error, 0)

	for _, fileHeader := range fileHeaders {
		wg.Add(1)

		go func(fileHeader *multipart.FileHeader, path string) {
			name, ext := util.NameAndExtFromFileHeader(fileHeader)

			if err := util.CreateLocal(fileHeader, path, name, ext); err != nil {
				errs = append(errs, err)
				util.NewInternalErrorLog("CREATE FILE ERROR", err).PrintLog("SERVER ERROR")
			}

			wg.Done()
		}(fileHeader, m.Path)
	}

	wg.Wait()

	if len(errs) > 0 {
		return fmt.Errorf("error creating local files: %+v", errs)
	}

	return nil
}

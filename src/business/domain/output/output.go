package output

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/erlendromo/forsete-atr/src/util"
	"github.com/google/uuid"
)

// Output
//
//	@Summary		Output
//	@Description	Output containing id, name, format, confirmed, image_id etc.
type Output struct {
	ID        uuid.UUID  `db:"id" json:"id"`
	Name      string     `db:"name" json:"name"`
	Format    string     `db:"format" json:"format"`
	Path      string     `db:"path" json:"-"`
	CreatedAt time.Time  `db:"created_at" json:"-"`
	UpdatedAt time.Time  `db:"updated_at" json:"-"`
	DeletedAt *time.Time `db:"deleted_at" json:"-"`
	Confirmed bool       `db:"confirmed" json:"confirmed"`
	ImageID   uuid.UUID  `db:"image_id" json:"image_id"`
}

type ATRResponse struct {
	Filename  string `json:"file_name"`
	Imagename string `json:"image_name"`
	Label     string `json:"label"`
	Contains  []struct {
		Segment struct {
			BBox struct {
				XMin int `json:"xmin"`
				YMin int `json:"ymin"`
				XMax int `json:"xmax"`
				YMax int `json:"ymax"`
			} `json:"bbox"`
			Polygon struct {
				Points []struct {
					X int `json:"x"`
					Y int `json:"y"`
				} `json:"points"`
			} `json:"polygon"`
			Score      float64  `json:"score"`
			ClassLabel string   `json:"class_label"`
			OrigShape  []int    `json:"orig_shape"`
			Data       struct{} `json:"data,omitempty"`
		} `json:"segment"`
		TextResult struct {
			Texts  []string  `json:"texts"`
			Scores []float64 `json:"scores"`
			Label  string    `json:"label"`
		} `json:"text_result"`
	} `json:"contains"`
}

func (o *Output) ReadJson(fullPathToFile string) (*ATRResponse, error) {
	file, err := os.Open(fullPathToFile)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return util.DecodeJSON[ATRResponse](file)
}

// Only supports json
func (o *Output) CreateLocal(data *ATRResponse) error {
	fullPathToFile := fmt.Sprintf("%s/%s.%s", o.Path, o.ID, o.Format)

	file, err := os.Create(fullPathToFile)
	if err != nil {
		return err
	}

	defer file.Close()

	if err := json.NewEncoder(file).Encode(data); err != nil {
		return err
	}

	return nil
}

func (o *Output) DeleteLocal() error {
	return os.Remove(fmt.Sprintf("%s/%s.%s", o.Path, o.ID.String(), o.Format))
}

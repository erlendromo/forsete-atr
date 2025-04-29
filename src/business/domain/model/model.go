package model

type Model struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Path        string `db:"path" json:"-"`
	ModelTypeID int    `db:"model_type_id" json:"-"`
	ModelType   string `db:"model_type" json:"model_type"`
}

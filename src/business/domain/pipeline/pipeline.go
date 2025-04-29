package pipeline

type Pipeline struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Path string `db:"path" json:"-"`
}

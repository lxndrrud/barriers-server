package classes

type Building struct {
	Id   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

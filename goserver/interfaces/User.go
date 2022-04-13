package interfaces

type User struct {
	Id         int64 `db:"id"`
	Firstname  string
	Middlename string
	Lastname   string
	SkudCard   string
}

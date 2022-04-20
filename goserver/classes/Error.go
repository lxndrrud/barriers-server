package classes

type CustomError struct {
	Text string
	Code int
}

func (e CustomError) Error() string {
	return e.Text
}

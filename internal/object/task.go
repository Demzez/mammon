package object

type Status int

const (
	TODO Status = iota
	INPROGRESS
	DONE
)

type Task struct {
	Name        string
	Description string
	Status      Status
}

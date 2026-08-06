package object

type Status int

const (
	TODO Status = iota
	INPROGRESS
	DONE
)

type Task struct {
	name        string
	description string
	status      Status
}

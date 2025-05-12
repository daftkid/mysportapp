package target

type Target struct {
	id        string
	numOfDays int
	numOfReps int
	comment   string
	userId    string
}

type TargetRepository interface {
	FindAll() ([]Target, error)
	CreateTarget(target Target) (int, error)
	ReadTarget(int) (Target, error)
	UpdateTarget(target Target) (int, error)
	DeleteTarget(int) error
}

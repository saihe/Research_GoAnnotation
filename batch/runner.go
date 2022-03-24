package batch

type runner struct {
	task Task
}

func NewRunner(t Task) runner {
	return runner{t}
}

func (runner) Run() int {

	return 0
}

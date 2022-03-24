package batch

import "goannotation/service"

type Task struct {
	serivce service.SampleService // `Autowired``
}

func (t *Task) Execute(msg string) {
	t.serivce.Hello(msg)
}

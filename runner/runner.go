package runner

import (
	"log"
	"sync"

	"github.com/ueqri/actions/loadbalancer"
	"github.com/ueqri/actions/task"
)

type Runner struct {
	Tasks    []task.Task
	Balancer loadbalancer.Balancer
}

func (r *Runner) Run() {
	localTasks := r.Balancer.LocalTasks(r.Tasks)

	log.Println("Running starts!")

	var wg sync.WaitGroup
	for _, v := range localTasks {
		wg.Add(1)
		go func(task task.Task) {
			defer wg.Done()
			task.StartTask()
			log.Println(task.TaskName(), "finished!")
		}(v)
	}
	wg.Wait()

	log.Println("Running completes!")
}

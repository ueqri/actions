package runner

import (
	"log"
	"sync"
	"sync/atomic"

	"github.com/ueqri/actions/collector"
	"github.com/ueqri/actions/loadbalancer"
	"github.com/ueqri/actions/task"
)

type Runner struct {
	PrefixName string
	Tasks      []task.Task
	Balancer   loadbalancer.Balancer
	Collector  *collector.Collector
}

func (r *Runner) Run() {
	localTasks := r.Balancer.LocalTasks(r.Tasks)
	numTasks := len(localTasks)

	log.Println("Running starts!")

	var status uint64 = 0
	var wg sync.WaitGroup
	for _, v := range localTasks {
		wg.Add(1)

		go func(task task.Task, wg *sync.WaitGroup, ops *uint64, num int) {
			defer wg.Done()

			task.StartTask()
			task.FinishTask()

			done := atomic.AddUint64(ops, 1)
			log.Printf("[%d/%d] %s finished!", done, num, task.TaskName())
		}(v, &wg, &status, numTasks)

	}
	wg.Wait()

	log.Println("Running completes!")
	if r.Collector != nil {
		log.Println("Report to collector...")
		r.Collector.Run(localTasks)
	}
}

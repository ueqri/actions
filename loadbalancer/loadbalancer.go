package loadbalancer

import "github.com/ueqri/actions/task"

type Balancer interface {
	LocalTasks(alltasks []task.Task) []task.Task
}

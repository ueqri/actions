package collector

import "github.com/ueqri/actions/task"

type Send interface {
	Messages(name string, msgs []string)
	Artifacts(name string, arts []string)
}
type Receive interface {
	OrganizedMessages() [][]string
	OrganizedArtifaces() [][]string
}
type Collector struct {
	Sender   Send
	Receiver Receive
}

func (c *Collector) Run(localTasks []task.Task) {
	for _, t := range localTasks {
		c.Sender.Messages(t.TaskName(), t.OutputMessage())
		c.Sender.Artifacts(t.TaskName(), t.ExportArtifacts())
	}
}

func (c *Collector) Collect() {
	// Plugin
}

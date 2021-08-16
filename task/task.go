package task

type Task interface {
	TaskName() string
	StartTask()
	FinishTask()
	OutputMessage() []string
	ExportArtifacts() []string
}

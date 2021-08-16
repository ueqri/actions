package main

import (
	"github.com/ueqri/actions/task"
)

type BasicTask struct {
	Benchmark
	outputMessage string
}

func (b *BasicTask) TaskName() string {
	return b.Name
}

func (b *BasicTask) StartTask() {
	b.outputMessage = task.ExecuteCommand(b.Path, "go build && "+b.Cmd)
	// fmt.Println(b.outputMessage)
}

func (b *BasicTask) FinishTask() {
	if !task.CheckArtifactsExist(b.Export) {
		panic("The artifact to export is not existed")
	}
}

func (b *BasicTask) OutputMessage() []string {
	return []string{b.outputMessage}
}

func (b *BasicTask) ExportArtifacts() []string {
	return []string{b.Export}
}

package main

import (
	"github.com/ueqri/actions/util"
)

type BasicTask struct {
	Benchmark
	outputMessage string
}

func (b *BasicTask) TaskName() string {
	return b.Name
}

func (b *BasicTask) StartTask() {
	b.outputMessage = util.ExecuteCommand(b.Path, "go build && "+b.Cmd)
	// fmt.Println(b.outputMessage)
}

func (b *BasicTask) FinishTask() {
	if !util.CheckArtifactsExist(b.Export) {
		panic("The artifact to export is not existed")
	}
}

func (b *BasicTask) OutputMessage() []string {
	return []string{b.outputMessage}
}

func (b *BasicTask) ExportArtifacts() []string {
	return []string{b.Export}
}

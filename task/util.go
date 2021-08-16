package task

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func ExecuteCommand(dir string, command string) string {
	var cmd *exec.Cmd

	var stdout bytes.Buffer
	var stderr bytes.Buffer

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/k", command)
	} else if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		cmd = exec.Command("/bin/sh", "-c", command)
	}

	cmd.Dir = dir
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()

	// fmt.Println("stdout:", stdout.String())
	if err != nil {
		fmt.Println("stderr:", stderr.String())
		panic(err)
	}

	return strings.TrimSpace(stdout.String())
}

func CheckArtifactsExist(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

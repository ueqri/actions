package task

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func ExecuteCommand(dir string, command string) (string, error) {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		fmt.Println("=== Command running in Windows ===")
		cmd = exec.Command("cmd", "/k")
	} else if runtime.GOOS == "darwin" {
		fmt.Println("=== Command running in macos ===")
		cmd = exec.Command("/bin/bash")
	} else {
		fmt.Println("=== Command running in Linux ===")
		cmd = exec.Command("/bin/sh")
	}

	in := bytes.NewBuffer(nil)

	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdin = in       // Binding Input
	cmd.Stdout = &out    // Binding Output
	cmd.Stderr = &stderr // Binding Error Output
	cmd.Dir = dir

	go func() {
		// Write your multiline commands, use `\n` represent a new line
		in.WriteString(command)
	}()
	log.Println(cmd.String())
	err := cmd.Start()

	if err != nil {
		return "Command start with error:" + err.Error() + ": " +
			stderr.String(), err
	}

	err = cmd.Wait()
	if err != nil {
		return "Command finished with error: " + err.Error() + ": " +
			stderr.String(), err
	}

	return out.String(), nil
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

package util

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path"
	"path/filepath"
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

func CheckCommandAvailable(command string) bool {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/k", command, "/?")
	} else if runtime.GOOS == "darwin" || runtime.GOOS == "linux" {
		cmd = exec.Command("/bin/sh", "-c", command, "--help")
	}

	if err := cmd.Run(); err != nil {
		return false
	}
	return true
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

func CopyFileToRemote(srcLocal, dstRemote, loginSSH string) {
	if !CheckCommandAvailable("scp") {
		panic("scp not exists")
	}
	cmd := fmt.Sprintf("scp %s %s%s", srcLocal, dstRemote, loginSSH)
	ExecuteCommand(path.Dir(srcLocal), cmd)
}

func WriteStringToFile(data, file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}

	writer := bufio.NewWriter(f)
	defer f.Close()

	fmt.Fprintln(writer, data)
	writer.Flush()
}

func RemoveFile(file string) {
	err := os.Remove(file)
	if err != nil {
		log.Fatal(err)
	}
}

func RenameFile(from, to string) {
	err := os.Rename(from, to)
	if err != nil {
		log.Fatal(err)
	}
}

func FindFiles(root, extension string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		if filepath.Ext(d.Name()) == extension {
			a = append(a, s)
		}
		return nil
	})
	return a
}

func ExpandTilde(path string) string {
	if len(path) == 0 || path[0] != '~' {
		return path
	}

	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return filepath.Join(usr.HomeDir, path[1:])
}

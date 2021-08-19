package collector

import (
	"fmt"
	"path/filepath"

	"github.com/ueqri/actions/util"
)

type SingleSender struct {
	Dst string
}

func (s *SingleSender) Messages(name string, msgs []string) {
	var data string
	for _, msg := range msgs {
		data = data + msg
	}
	file := name + ".log"
	util.WriteStringToFile(data, file)
	util.CopyFile(file, filepath.Join(s.Dst, file))
	util.RemoveFile(file)
}

func (s *SingleSender) Artifacts(name string, arts []string) {
	for _, art := range arts {
		dir, originFile := filepath.Split(art)
		renameFile := fmt.Sprintf("%s_%s", name, originFile)
		renamePath := filepath.Join(dir, renameFile)
		util.RenameFile(art, renamePath)
		util.CopyFile(renamePath, filepath.Join(s.Dst, renameFile))
		util.RemoveFile(renamePath)
	}
}

type SingleReceiver struct {
	Dir string
}

func (r *SingleReceiver) OrganizedMessages() [][]string {
	return [][]string{util.FindFiles(r.Dir, ".log")}
}

func (r *SingleReceiver) OrganizedArtifaces() [][]string {
	return [][]string{util.FindFiles(r.Dir, ".csv")}
}

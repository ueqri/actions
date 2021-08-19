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
	util.CopyFile(file, s.Dst)
}

func (s *SingleSender) Artifacts(name string, arts []string) {
	for _, art := range arts {
		dir, origin := filepath.Split(art)
		rename := filepath.Join(dir, fmt.Sprintf("%s_%s", name, origin))
		util.RenameFile(origin, rename)
		util.CopyFile(rename, s.Dst)
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

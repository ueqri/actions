package collector

import (
	"fmt"
	"path/filepath"

	"github.com/ueqri/actions/util"
)

type LocalSender struct {
	Dst      string
	LoginSSH string
}

func (s *LocalSender) Messages(name string, msgs []string) {
	var data string
	for _, msg := range msgs {
		data = data + msg
	}
	file := name + ".log"
	// util.RemoveFile(file)
	util.WriteStringToFile(data, file)
	util.CopyFileToRemote(file, s.Dst, s.LoginSSH)
}

func (s *LocalSender) Artifacts(name string, arts []string) {
	for _, art := range arts {
		dir, origin := filepath.Split(art)
		rename := filepath.Join(dir, fmt.Sprintf("%s_%s", name, origin))
		util.RenameFile(origin, rename)
		util.CopyFileToRemote(rename, s.Dst, s.LoginSSH)
	}
}

type LocalReceiver struct {
	Dir string
}

func (r *LocalReceiver) OrganizedMessages() [][]string {
	return [][]string{util.FindFiles(r.Dir, ".log")}
}

func (r *LocalReceiver) OrganizedArtifaces() [][]string {
	return [][]string{util.FindFiles(r.Dir, ".csv")}
}

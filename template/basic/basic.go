package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/ueqri/actions/collector"
	"github.com/ueqri/actions/loadbalancer"
	"github.com/ueqri/actions/runner"
)

func main() {
	rankPtr := flag.Int("rank", -1, "Local rank of this runner")
	configPrt := flag.String("config", "", "Path of the configuration")
	templatePrt := flag.String("template", "", "Path of the actions template")
	flag.Parse()

	alltasks, numRanks := ParserConfig(*configPrt, *templatePrt)

	// fmt.Printf("%v\n%#v", numRanks, alltasks)

	balancer := loadbalancer.MakeBasicBalancer().
		WithNumberOfRanks(numRanks).
		WithRankID(*rankPtr)

	dir, _ := os.Getwd()
	c := new(collector.Collector)
	c.Sender = &collector.SingleSender{
		Dst: filepath.Join(dir, ".collect"),
	}
	c.Receiver = &collector.SingleReceiver{
		Dir: filepath.Join(dir, ".collect"),
	}

	runner := new(runner.Runner)
	runner.Balancer = &balancer
	runner.Collector = c
	for _, t := range alltasks {
		runner.Tasks = append(runner.Tasks, t)
	}

	runner.Run()

	if *rankPtr == 0 {
		c.Collect()
	}
}

package main

import (
	"flag"

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

	c := new(collector.Collector)
	c.Sender = &collector.LocalSender{
		Dst:      "",
		LoginSSH: "",
	}
	c.Receiver = &collector.LocalReceiver{
		Dir: "",
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

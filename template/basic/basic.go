package main

import (
	"flag"

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

	runner := new(runner.Runner)
	runner.Balancer = &balancer
	for _, t := range alltasks {
		runner.Tasks = append(runner.Tasks, t)
	}
	runner.Run()
}

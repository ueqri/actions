package loadbalancer

import "github.com/ueqri/actions/task"

type BasicBalancer struct {
	RankID   int
	NumRanks int
}

func MakeBasicBalancer() BasicBalancer {
	b := BasicBalancer{-1, -1}
	return b
}

func (b BasicBalancer) WithRankID(id int) BasicBalancer {
	b.RankID = id
	return b
}

func (b BasicBalancer) WithNumberOfRanks(num int) BasicBalancer {
	b.NumRanks = num
	return b
}

func (b *BasicBalancer) LocalTasks(alltasks []task.Task) []task.Task {
	if b.RankID < 0 {
		panic("Rank ID must be set to a positive number")
	}
	if b.NumRanks < 0 {
		panic("Number of ranks must be set to a positive number")
	}
	if b.RankID >= b.NumRanks {
		panic("Rank ID should not be equal to or greater than number of ranks")
	}

	var ret []task.Task
	batch := len(alltasks)
	load := batch / b.NumRanks
	rest := batch % b.NumRanks

	if b.RankID < rest {
		startTask := b.RankID * (load + 1)
		endTask := startTask + load + 1
		for j := startTask; j < endTask; j++ {
			ret = append(ret, alltasks[j])
		}
	} else {
		startTask := b.RankID*load + rest
		endTask := startTask + load
		for j := startTask; j < endTask; j++ {
			ret = append(ret, alltasks[j])
		}
	}
	return ret
}

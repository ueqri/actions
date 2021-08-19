package runner

import "time"

type Timer struct {
	start  time.Time
	finish time.Time
}

func (t *Timer) Register(name string) {
	t.start = time.Now()

}

func (t *Timer) Record(name string) {
	t.finish = time.Now()
}

func (t *Timer) Elapsed(name string) time.Duration {
	return t.finish.Sub(t.start)
}

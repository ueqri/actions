package runner

import "time"

type timer struct {
	start  map[string]time.Time
	finish map[string]time.Time
}

func (t *timer) Register(name string) {
	_, ok := t.start[name]
	if ok {
		panic("duplicated registration for " + name + " timer")
	} else {
		t.start[name] = time.Now()
	}
}

func (t *timer) Record(name string) {
	_, ok := t.finish[name]
	if ok {
		panic("duplicated record for " + name + " timer")
	} else {
		t.finish[name] = time.Now()
	}
}

func (t *timer) Elapsed(name string) time.Duration {
	if _, ok := t.start[name]; ok {
		panic("no registration for " + name + " timer")
	}
	if _, ok := t.finish[name]; ok {
		panic("no record for " + name + " timer")
	}
	return t.finish[name].Sub(t.start[name])
}

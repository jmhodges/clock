package clock

import "time"

type Timer struct {
	C <-chan time.Time

	timer     *time.Timer
	fakeTimer *fakeTimer
}

func (t *Timer) Reset(d time.Duration) bool {
	if t.timer != nil {
		return t.timer.Reset(d)
	}
	return t.fakeTimer.Reset(d)
}

func (t *Timer) Stop() bool {
	if t.timer != nil {
		return t.timer.Stop()
	}
	return t.fakeTimer.Stop()
}

type fakeTimer struct {
	c      chan<- time.Time
	target time.Time
	clk    *fake
	active bool
}

func (ft *fakeTimer) Reset(d time.Duration) bool {
	ft.clk.Lock()
	defer ft.clk.Unlock()
	ft.target = ft.clk.t.Add(d)
	active := ft.active
	ft.active = true
	if !active { // FIXME
		ft.clk.timers = append(ft.clk.timers, ft)
	}
	ft.clk.sendTimes()
	return active
}

func (ft *fakeTimer) Stop() bool {
	ft.clk.Lock()
	defer ft.clk.Unlock()
	active := ft.active
	ft.active = false
	ft.clk.sendTimes()
	return active
}

type sortedFakeTimers []*fakeTimer

func (s sortedFakeTimers) Len() int {
	return len(s)
}

func (s sortedFakeTimers) Less(i, j int) bool {
	return s[i].target.Before(s[j].target)
}

func (s sortedFakeTimers) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

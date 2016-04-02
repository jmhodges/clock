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
	clk    *fake
	active bool
	// sends is where we store all the sends made by this timer so we can
	// deactivate the old ones when Reset or Stop is called.
	sends []*send
}

func (ft *fakeTimer) Reset(d time.Duration) bool {
	ft.clk.Lock()
	defer ft.clk.Unlock()
	target := ft.clk.t.Add(d)
	active := ft.active
	ft.active = true
	for _, s := range ft.sends {
		s.active = false
	}
	s := ft.clk.addSend(target, ft)
	ft.sends = []*send{s}
	ft.clk.sendTimes()
	return active
}

func (ft *fakeTimer) Stop() bool {
	ft.clk.Lock()
	defer ft.clk.Unlock()
	active := ft.active
	ft.active = false
	for _, s := range ft.sends {
		s.active = false
	}
	ft.sends = nil
	ft.clk.sendTimes()
	return active
}

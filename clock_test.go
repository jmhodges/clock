package clock

import (
	"testing"
	"time"
)

func TestFakeClockGoldenPath(t *testing.T) {
	clk := NewFake()
	second := NewFake()
	if !clk.Now().Equal(second.Now()) {
		t.Errorf("clocks must start out at the same time but didn't: %#v vs %#v", clk.Now(), second.Now())
	}
	clk.Add(3 * time.Second)
	if clk.Now().Equal(second.Now()) {
		t.Errorf("clocks different must differ: %#v vs %#v", clk.Now(), second.Now())
	}
}

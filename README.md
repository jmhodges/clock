clock
====

[![Build Status](https://travis-ci.org/jmhodges/clock.png?branch=master)](https://travis-ci.org/jmhodges/clock)

Package clock provides an abstraction for system time that enables
testing of time-sensitive code.

Where you'd use time.Now, instead use clk.Now where clk is an instance
of Clock.

Pass in a Clock given by Default() when running your code in
production and pass it an instance of Clock from NewFake() when
running it in your tests.

When you do that, you can use FakeClock's the Add and Set methods to
control how time behaves in your code making them more reliable while
also expanding the space of problems you can test.

Be sure to test Time equality with time.Time#Equal, not ==.Where you
would use `time.Now()`, instead use `clk.Now()` where `clk` is an
instance of `clock.Clock`. You can make clocks with `clock.Default()`
to get the "actual time" or `clock.NewFake()` to get a time that you
control with `FakeClock.Add` and `FakeClock.Set`.

For documentation, see the
[godoc](http://godoc.org/github.com/jmhodges/clock).

// Package
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
	}

	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, "Go!")
}

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

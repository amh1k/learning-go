package main

import (
	"fmt"
	"io"
	"os"
	"time"
)
const finalWord = "Go!"
const countdownStart = 3
const sleep = "sleep"
const write = "write"

// interace used to mock sleep. It lets us then use a real Sleeper in main 
// and a spy sleeper in our tests. By using an interface our Countdown function is 
// oblivious to this and adds some flexibility for the caller.
type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep() {
	s.Calls++;
}
//Spies are a kind of mock which can record how a dependency is used. 
// They can record the arguments sent in, how many times it has been called, etc. 
// In our case, we're keeping track of 
// how many times Sleep() is called so we can check it in our test.


type DefaultSleeper struct {
}
func (d *DefaultSleeper)Sleep() {
	time.Sleep(1 * time.Second)
}

//The issue with the above spy functions is that only takes into account
// no of sleep calls made and not their order
type SpyCountDownOperations struct {
	Calls []string
}
func (s * SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}
// this implements io writer
func (s *SpyCountDownOperations) Write(p []byte)(n int, err error) {
	s.Calls = append(s.Calls, write)
	return

}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}
type SpyTime struct {
	durationSlept time.Duration

}
func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)

}


func main() {
	// sleeper := &DefaultSleeper{}
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}


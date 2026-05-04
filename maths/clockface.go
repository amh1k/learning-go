package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	thirtySeconds := time.Date(312, time.October, 28, 0, 0, 30, 0, time.UTC)
	want := math.Pi
	got := secondsInRadians(thirtySeconds)

	if want != got {
		t.Fatalf("Wanted %v radians, but got %v", want, got)
	}
}

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}
func secondsInRadians(t time.Time) float64 {
	return math.Pi
}

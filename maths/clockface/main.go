package main

import (
	clockface "learninggolang/maths"
	"os"
	"time"
)
func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
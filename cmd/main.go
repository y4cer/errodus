package main

import (
	"github.com/y4cer/errodus/internal/math"
)

func main() {
	var m math.Matrix

	m.ParseFromFile("examples/2x2")
	m.Print("2x2")
}

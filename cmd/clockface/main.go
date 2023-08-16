package main

import (
	"os"
	"time"

	"github.com/brunoquindeler/go-with-tests/math/clockface/pkg/svg"
)

func main() {
	t := time.Now()
	svg.SVGWriter(os.Stdout, t)
}

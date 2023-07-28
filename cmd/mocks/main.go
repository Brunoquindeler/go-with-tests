package main

import (
	"os"
	"time"

	"github.com/brunoquindeler/go-with-tests/mocks"
)

func main() {
	sleeper := mocks.NewConfigurableSleeper(time.Second, time.Sleep)
	mocks.Countdown(os.Stdout, sleeper)
}

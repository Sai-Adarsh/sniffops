package main

import (
	"time"

	memObj "./lib"
)

func main() {
	memObj.CpuHeartBeat()
	time.Sleep(time.Millisecond * 1000)
}

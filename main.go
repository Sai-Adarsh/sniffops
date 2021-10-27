package main

import (
	"time"

	memObj "./lib"
)

func main() {
	memObj.MemoryHeartBeat()
	time.Sleep(time.Millisecond * 1000)
}

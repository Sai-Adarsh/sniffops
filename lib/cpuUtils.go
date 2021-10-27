package lib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mackerelio/go-osstat/cpu"
)

// This func must be Exported, Capitalized, and comment added.
func CpuFunction() {
	before, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	time.Sleep(time.Duration(1) * time.Second)

	after, err := cpu.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	total := float64(after.Total - before.Total)
	fmt.Printf("cpu user: %f %%\n", float64(after.User-before.User)/total*100)
	fmt.Printf("cpu system: %f %%\n", float64(after.System-before.System)/total*100)
	fmt.Printf("cpu idle: %f %%\n\n", float64(after.Idle-before.Idle)/total*100)

	file, err := os.OpenFile("./src/media/cpu.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("failed to open file: %d", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	// Using Write
	row := []string{
		strconv.Itoa(int(float64(after.User-before.User) / total * 100)),
		strconv.Itoa(int(float64(after.System-before.System) / total * 100)),
		strconv.Itoa(int(float64(after.Idle-before.Idle) / total * 100)),
	}

	if err := w.Write(row); err != nil {
		fmt.Printf("error writing record to file: %d", err)
	}
}

// This func must be Exported, Capitalized, and comment added.
func CpuHeartBeat() {
	for range time.Tick(time.Millisecond * 1000) {
		CpuFunction()
	}
}

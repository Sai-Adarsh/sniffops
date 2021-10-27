package lib

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mackerelio/go-osstat/memory"
)

// This func must be Exported, Capitalized, and comment added.
func CreateLogFile() {
	csvFile, err := os.Create("memory.csv")
	if err != nil {
		fmt.Printf("failed creating file: %s", err)
	}
	csvFile.Close()
}

// This func must be Exported, Capitalized, and comment added.
func MemoryFunction() {
	memory, err := memory.Get()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	fmt.Printf("memory total: %d bytes\n", memory.Total)
	fmt.Printf("memory used: %d bytes\n", memory.Used)
	fmt.Printf("memory cached: %d bytes\n", memory.Cached)
	fmt.Printf("memory free: %d bytes\n\n", memory.Free)

	file, err := os.OpenFile("memory.csv", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("failed to open file: %d", err)
	}

	w := csv.NewWriter(file)
	defer w.Flush()

	// Using Write
	row := []string{
		strconv.Itoa(int(memory.Total)),
		strconv.Itoa(int(memory.Used)),
		strconv.Itoa(int(memory.Cached)),
		strconv.Itoa(int(memory.Free)),
	}

	if err := w.Write(row); err != nil {
		fmt.Printf("error writing record to file: %d", err)
	}
}

// This func must be Exported, Capitalized, and comment added.
func MemoryHeartBeat() {
	for range time.Tick(time.Millisecond * 500) {
		MemoryFunction()
	}
}

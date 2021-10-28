package main

import (
	"fmt"
	"log"
	"time"

	memObj "./lib"
	"github.com/cshenton/seer-golang/client"
)

func main() {
	memObj.CpuHeartBeat()
	time.Sleep(time.Millisecond * 1000)

	c, err := client.New("localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	// Create a stream
	_, err = c.CreateStream("myStream", 3600)
	if err != nil {
		log.Fatal(err)
	}

	// Add in data
	_, err = c.UpdateStream(
		"myStream",
		[]time.Time{
			time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2016, 1, 2, 0, 0, 0, 0, time.UTC),
			time.Date(2016, 1, 3, 0, 0, 0, 0, time.UTC),
		},
		[]float64{10, 9, 6},
	)
	if err != nil {
		log.Fatal(err)
	}

	// Generate and display forecast
	f, err := c.GetForecast("myStream", 10)
	fmt.Println(f)
}

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	_, err := exec.Command("bash", "-c", "pkill -SIGINT client").Output()

	if err == nil {
		fmt.Println("Success")
	} else {
		fmt.Println("Failure")
	}
}

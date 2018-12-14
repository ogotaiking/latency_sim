package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd, err := exec.Command("ifconfig", "enp1s0f0", "1.1.1.1/24", "up").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmd))
}

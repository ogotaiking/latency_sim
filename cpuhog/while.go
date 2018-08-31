package main

import (
	"sync"
)

func loopwhile() {
	for {
		//time.Sleep(time.Nanosecond * time.Duration(1))
	}
}
func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go loopwhile()
	go loopwhile()
	go loopwhile()
	go loopwhile()
	go loopwhile()
	go loopwhile()
	go loopwhile()
	go loopwhile()
	wg.Wait()

}

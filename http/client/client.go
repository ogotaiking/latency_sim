package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func HTTPClient(host string, latency int, size int, tick int) {
	rand.Seed(time.Now().UnixNano())

	for {
		tc := time.After(time.Microsecond * time.Duration(tick*1000))
		ratio := (float32(rand.Int31n(100))) / float32(100)
		sizeR := int(float32(size) * (0.1 + ratio))
		latency = int(float32(latency) * (0.5 + ratio/2))
		url := fmt.Sprintf("http://%s/test/%d/%d", host, sizeR, latency)
		response, err := http.Get(url)
		if err != nil {
			logrus.Warn(err)
		} else {
			_, err = ioutil.ReadAll(response.Body)
                if err != nil {
                        logrus.Warn(err)
                }
			response.Body.Close()
		}

		<-tc
	}

}

var commandOptions = struct {
	Worker  int
	URL     string
	Tick    int
	Latency int
	Size    int
}{
	100,
	"127.0.0.1:8080",
	1000,
	100,
	1000,
}

func init() {
	flag.IntVar(&commandOptions.Worker, "n", commandOptions.Worker, "Number of Client Worker")
	flag.IntVar(&commandOptions.Tick, "t", commandOptions.Tick, "tick time(ms)")
	flag.StringVar(&commandOptions.URL, "u", commandOptions.URL, "ServerURL")
	flag.IntVar(&commandOptions.Latency, "l", commandOptions.Latency, "Server Response Latency(ms)")
	flag.IntVar(&commandOptions.Size, "s", commandOptions.Size, "Server Response size(kb)")

	flag.Parse()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < commandOptions.Worker; i++ {
		go HTTPClient(commandOptions.URL, commandOptions.Latency, commandOptions.Size, commandOptions.Tick)
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(100)))
	}
	wg.Wait()

}

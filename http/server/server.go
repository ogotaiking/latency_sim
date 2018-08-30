package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var a []byte

func init() {
	const maxsize int = 1024 * 1024 * 1024
	a = make([]byte, maxsize, maxsize)
}

func greet(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1000 * time.Millisecond)
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func TestTraffic(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	if _l, ok := vars["latency"]; ok {

		latency, _ := strconv.ParseInt(_l, 10, 64)
		//fmt.Fprintf(w, "sleeping.........%d\n\n\n", latency)
		time.Sleep(time.Duration(latency) * time.Millisecond)
	}

	if s, ok := vars["size"]; ok {
		size, _ := strconv.ParseInt(s, 10, 64)
		size = size * 512
		fmt.Fprintf(w, "%b", a[:size])
	}

	//fmt.Fprintf(w, "size: %v\n", vars["size"])
	//fmt.Fprintf(w, "latency: %v\n", vars["latency"])
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/test/{size}/{latency}", TestTraffic)
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:8080", nil)
}

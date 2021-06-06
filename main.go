package main

import (
	"fmt"
	"net/http"
	"stuff/skittlespin"
	"sync"
)

type pinRequestHandler struct {
	pin skittlespin.Skittlespin
}

func (pr pinRequestHandler) pinHandler(w http.ResponseWriter, r *http.Request) {
	globalLock.Lock()
	pr.pin.ActuatePin()
	globalLock.Unlock()
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Up and Atom!")

}

var (
	globalLock sync.Mutex
)

func main() {
	pr := pinRequestHandler{pin: *skittlespin.NewSkittlesPin(21, "relay", "output")}
	http.HandleFunc("/garage", pr.pinHandler)
	http.HandleFunc("/health", healthCheck)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Print(err)
	}

}

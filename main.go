package main

import (
	"fmt"
	"net/http"
	"sync"
	"zerrofour4/go-garage/skittlespin"

	"github.com/d2r2/go-dht"
)

type pinRequestHandler struct {
	pin skittlespin.Skittlespin
}

func (pr pinRequestHandler) pinHandler(w http.ResponseWriter, r *http.Request) {
	globalLock.Lock()
	pr.pin.ActuatePin()
	globalLock.Unlock()
}

func (pr pinRequestHandler) sensorHandler(w http.ResponseWriter, r *http.Request) {
	temperature, humidity, _, temperr :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 26, false, 5)
	if temperr != nil {
		fmt.Fprintf(w, "error getting temp %s", temperr)
		return
	}
	tempF := temperature*1.8 + 32
	fmt.Fprintf(w, "tempC: %v\ntempF: %v\nhumidity: %v%%", temperature, tempF, humidity)

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
	http.HandleFunc("/temp", pr.sensorHandler)

	httperr := http.ListenAndServe(":8000", nil)
	if httperr != nil {
		fmt.Print(httperr)
	}

}

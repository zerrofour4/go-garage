package main

import (
	"fmt"
	"net/http"

	"stuff/skittlespin"
)

type pinRequestHandler struct {
	pin skittlespin.Skittlespin
}

func (pr pinRequestHandler) pinHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
	pr.pin.ActuatePin()
}

func main() {
	pr := pinRequestHandler{pin: *skittlespin.NewSkittlesPin(21, "relay", "output")}
	http.HandleFunc("/", pr.pinHandler)
	http.ListenAndServe(":8000", nil)

}

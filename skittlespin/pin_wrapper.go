package skittlespin

import (
	"time"

	"github.com/stianeikeland/go-rpio"
)

type Skittlespin struct {
	label  string
	number int
	mode   string
	pin    rpio.Pin
}

func NewSkittlesPin(number int, label string, mode string) *Skittlespin {

	rpio.Open()
	rpiopin := rpio.Pin(number)

	if mode == "output" {
		rpiopin.Output()
	} else {
		rpiopin.Input()
	}

	s := Skittlespin{
		label:  label,
		number: number,
		mode:   mode,
		pin:    rpiopin}

	return &s
}

func (s Skittlespin) ActuatePin() {
	s.pin.High() // Set pin High
	time.Sleep(1 * time.Second)
	s.pin.Low() // Set pin Low
}

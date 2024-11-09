package pins

import (
	"fmt"

	"github.com/stianeikeland/go-rpio/v4"
)

type PinsService struct {
	pins map[int]rpio.Pin
}

func NewPinsService() *PinsService {
	return &PinsService{
		pins: make(map[int]rpio.Pin),
	}
}

func (s *PinsService) InitOutPin(pin int) {
	s.pins[pin] = rpio.Pin(pin)
	s.pins[pin].Output()
}

func (s *PinsService) SetHigh(number int) error {
	pin, exists := s.pins[number]
	if !exists {
		return fmt.Errorf("pin is not initialized: %d", number)
	}
	pin.High()
	return nil
}

func (s *PinsService) SetLow(number int) error {
	pin, exists := s.pins[number]
	if !exists {
		return fmt.Errorf("pin is not initialized: %d", number)
	}
	pin.High()
	return nil
}

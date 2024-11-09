package relays

import (
	"beaver/thing/pins"
	"fmt"
)

type Relay struct {
	pin   int
	state bool
}

type RelaysService struct {
	pinsService *pins.PinsService
	relays      map[string]*Relay
}

func NewRelaysService() *RelaysService {
	relaysService := &RelaysService{
		pinsService: pins.NewPinsService(),
	}
	return relaysService
}

func (s *RelaysService) InitRelays(name string, pin int) error {
	_, exists := s.relays[name]
	if exists {
		fmt.Errorf("relays already exists")
	}
	s.pinsService.InitOutPin(pin)
	err := s.pinsService.SetHigh(pin)
	if err != nil {
		fmt.Errorf("unable to initRelays: %w", err)
	}
	s.relays[name] = &Relay{
		pin:   pin,
		state: false,
	}
}

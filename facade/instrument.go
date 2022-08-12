package facade

import (
	"fmt"

	"github.com/go-facade-example/instruments"
)

type InstrumentsFacade struct {
	Bass *instruments.Bass
}

func NewInstrumentsFacade() *InstrumentsFacade {
	return &InstrumentsFacade{
		&instruments.Bass{},
	}
}

func (i *InstrumentsFacade) PlayMusic() {
	fmt.Println(i.Bass.PlayBass())
}

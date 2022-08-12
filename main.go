package main

import (
	"github.com/go-facade-example/facade"
)

func main() {
	instrument := facade.NewInstrumentsFacade()
	instrument.PlayMusic()
}

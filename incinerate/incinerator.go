package main

import (
	"fmt"
	"time"
)

type Incinerator struct {
	name    string
	gophers chan *Gopher
	errors  chan error
}

func NewIncinerator(name string) (i *Incinerator) {
	i = &Incinerator{
		name:    name,
		gophers: make(chan *Gopher),
		errors:  make(chan error),
	}

	go i.run()

	return
}

func (i *Incinerator) run() {
	for {
		select {
		case g := <-i.gophers:
			status := fmt.Sprintf("Burning book for Gopher %s", g.name)
			i.ReportStatus(status)
			time.Sleep(1 * time.Second)
			i.errors <- nil
		}
	}
}

func (i *Incinerator) ReportStatus(status string) {
	fmt.Printf("Incinerator %s: %s\n", i.name, status)
}

func (i *Incinerator) RequestBurn(g *Gopher) error {
	i.gophers <- g
	return <-i.errors
}

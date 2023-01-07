package main

import (
	"log"

	v1 "github.com/punchguard/v1"
)

type MainSignal struct {
	v1.BaseSignal
}

func (d *MainSignal) Run() *v1.Control { return nil }

type MainOtr struct {
	v1.BaseOtr
}

type MainPunch struct {
	v1.BasePunch
}

func (d *MainPunch) RunOnce() error { return nil }

type MainGuard struct {
	v1.BaseGuard
}

func (d *MainGuard) Run() *v1.Control { return nil }

func main() {
	log.SetPrefix("punchguard: ")
	log.SetFlags(0)

	log.Print("version 0.0.1")

	var id v1.Id = v1.Id{
		Signal: "s",
		Otr:    "o",
		Guard:  "g",
	}

	signal, otr, punch, guard := v1.NewFlowParams[MainSignal, MainOtr, MainPunch, MainGuard]()
	if err := v1.InitFlowParams(signal, otr, punch, guard); err != nil {
		log.Fatal(err)
	}

	flow, err := v1.NewOnDemandFlow(id, signal, otr, punch, guard)
	if err != nil {
		log.Fatal(err)
	}

	flow.Run()
}

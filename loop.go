package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/punchguard/v0"
	"github.com/punchguard/v0/unstable/flow"
	"github.com/punchguard/v0/unstable/guard"
	"github.com/punchguard/v0/unstable/otr"
	"github.com/punchguard/v0/unstable/punch"
	"github.com/punchguard/v0/unstable/signaling"
)

type Main struct{}

func main() {
	log.SetPrefix("punchguard: ")
	log.SetFlags(0)

	log.Print("version 0.0.1")

	var id1 punchguard.Id = punchguard.NewId("s1", "o1", "g1")
	var id2 punchguard.Id = punchguard.NewId("s2", "o2", "g2")

	flow1, err := punchguard.Init[signaling.Tox, otr.Otr3, punch.IceLite, guard.WireGuard, flow.PunchOnDemand](id1)
	if err != nil {
		log.Fatal(err)
	}

	f1, _ := flow1.(*flow.PunchOnDemand)
	p1, _ := f1.GetPunch().(*punch.IceLite)
	p1.Id = "p1"

	flow2, err := punchguard.Init[signaling.Tox, otr.Otr3, punch.IceLite, guard.WireGuard, flow.PunchOnDemand](id2)
	if err != nil {
		log.Fatal(err)
	}

	f2, _ := flow2.(*flow.PunchOnDemand)
	p2, _ := f2.GetPunch().(*punch.IceLite)
	p2.Id = "p2"

	stopped1 := flow1.Start()
	stopped2 := flow2.Start()

	go func() {
		term := make(chan os.Signal, 1)

		signal.Notify(term, syscall.SIGTERM)
		signal.Notify(term, os.Interrupt)

		<-term
		log.Print("main: flow.Stop")
		flow1.Stop()
		flow2.Stop()
	}()

	<-stopped1
	<-stopped2
	log.Print("main: flow stopped")

	log.Print("main: exiting")
}

package v1_test

import (
	"log"
	"testing"

	v1 "github.com/punchguard/v1"
)

var (
	DummyIdTest     v1.Id
	DummySignalTest *DummySignal
	DummyOtrTest    *DummyOtr
	DummyPunchTest  *DummyPunch
	DummyGuardTest  *DummyGuard
)

func init() {
	DummyIdTest = v1.NewId("s", "o", "g")

	if d, err := v1.NewSignal[DummySignal](DummyIdTest.Signal); err != nil {
		log.Fatal(err)
	} else {
		DummySignalTest = d.(*DummySignal)
	}

	if d, err := v1.NewOtr[DummyOtr](DummyIdTest.Otr, DummySignalTest.GetSend(), DummySignalTest.GetOnReceive()); err != nil {
		log.Fatal(err)
	} else {
		DummyOtrTest = d.(*DummyOtr)
	}

	if d, err := v1.NewPunch[DummyPunch](DummyOtrTest.GetPunchSend(), DummyOtrTest.GetPunchOnReceive()); err != nil {
		log.Fatal(err)
	} else {
		DummyPunchTest = d.(*DummyPunch)
	}

	if d, err := v1.NewGuard[DummyGuard](DummyIdTest.Guard, DummyPunchTest.GetOnPeers()); err != nil {
		log.Fatal(err)
	} else {
		DummyGuardTest = d.(*DummyGuard)
	}
}

type DummySignal struct {
	v1.BaseSignal
}

func (d *DummySignal) Start() v1.Controlling { return nil }

type DummySignalOnReadyNil struct {
	DummySignal
}

func (b *DummySignalOnReadyNil) GetOnReady() <-chan struct{} { return nil }

type DummyOtr struct {
	v1.BaseOtr
}

type DummyPunch struct {
	v1.BasePunch
}

func (d *DummyPunch) RunOnce() error { return nil }

type DummyGuard struct {
	v1.BaseGuard
}

func (d *DummyGuard) Start() v1.Controlling { return nil }

type DummyGuardOnConnectedNil struct {
	DummyGuard
}

func (b *DummyGuardOnConnectedNil) GetOnConnected() <-chan struct{} { return nil }

type DummyGuardOnDisconnectedNil struct {
	DummyGuard
}

func (b *DummyGuardOnDisconnectedNil) GetOnDisconnected() <-chan struct{} { return nil }

type DummyFlow struct {
	v1.BaseFlow
}

func (d *DummyFlow) Start() v1.Controlling { return nil }

func TestDummySignalRunIsNil(t *testing.T) {
	d := new(DummySignal)

	run := d.Start()
	if run != nil {
		t.Fatalf("DummySignal.Start() = '%v', want '%v'", run, nil)
	}
}

func TestDummyPunchRunOnceIsNil(t *testing.T) {
	d := new(DummyPunch)

	err := d.RunOnce()
	if err != nil {
		t.Fatalf("DummyPunch.RunOnce() = '%v', want '%v'", err, nil)
	}
}

func TestDummyGuardRunIsNil(t *testing.T) {
	d := new(DummyGuard)

	run := d.Start()
	if run != nil {
		t.Fatalf("DummyGuard.Start() = '%v', want '%v'", run, nil)
	}
}

func TestDummyFlowRunIsNil(t *testing.T) {
	d := new(DummyFlow)

	run := d.Start()
	if run != nil {
		t.Fatalf("DummyFlow.Start() = '%v', want '%v'", run, nil)
	}
}

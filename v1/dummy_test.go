package v1

import (
	"errors"
	"testing"
)

var (
	dummyId     Id
	dummySignal *DummySignal
	dummyOtr    *DummyOtr
	dummyPunch  *DummyPunch
	dummyGuard  *DummyGuard
)

func init() {
	dummyId = Id{
		Signal: "s",
		Otr:    "o",
		Guard:  "g",
	}

	dummySignal, dummyOtr, dummyPunch, dummyGuard = NewFlowParams[DummySignal, DummyOtr, DummyPunch, DummyGuard]()

	err := InitFlowParams(dummySignal, dummyOtr, dummyPunch, dummyGuard)
	if err != nil {
		panic(err)
	}
}

type DummySignal struct {
	BaseSignal
}

func (d *DummySignal) Run() *Control { return nil }

type DummySignalInitFail struct {
	DummySignal
}

func (d *DummySignalInitFail) Init() error {
	return errors.New("err")
}

type DummyOtr struct {
	BaseOtr
}

type DummyPunch struct {
	BasePunch
}

func (d *DummyPunch) RunOnce() error { return nil }

type DummyGuard struct {
	BaseGuard
}

func (d *DummyGuard) Run() *Control { return nil }

func TestDummySignalRunIsNil(t *testing.T) {
	d := new(DummySignal)

	run := d.Run()
	if run != nil {
		t.Fatalf("DummySignal.Run() = '%v', want '%v'", run, nil)
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

	run := d.Run()
	if run != nil {
		t.Fatalf("DummyGuard.Run() = '%v', want '%v'", run, nil)
	}
}

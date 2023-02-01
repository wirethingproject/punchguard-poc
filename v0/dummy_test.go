package punchguard_test

// import (
// 	"log"
// 	"testing"

// 	v1 "github.com/punchguard/v1"
// )

// var (
// 	DummyIdTest     v1.Id
// 	DummySignalingTest *DummySignaling
// 	DummyOtrTest    *DummyOtr
// 	DummyPunchTest  *DummyPunch
// 	DummyGuardTest  *DummyGuard
// )

// func init() {
// 	DummyIdTest = v1.NewId("s", "o", "g")

// 	if d, err := v1.NewSignaling[DummySignaling](DummyIdTest.Signaling); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		DummySignalingTest = d.(*DummySignaling)
// 	}

// 	if d, err := v1.NewOtr[DummyOtr](DummyIdTest.Otr, DummySignalingTest.GetSend(), DummySignalingTest.GetOnReceive()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		DummyOtrTest = d.(*DummyOtr)
// 	}

// 	if d, err := v1.NewPunch[DummyPunch](DummyOtrTest.GetPunchSend(), DummyOtrTest.GetPunchOnReceive()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		DummyPunchTest = d.(*DummyPunch)
// 	}

// 	if d, err := v1.NewGuard[DummyGuard](DummyIdTest.Guard, DummyPunchTest.GetOnPeers()); err != nil {
// 		log.Fatal(err)
// 	} else {
// 		DummyGuardTest = d.(*DummyGuard)
// 	}
// }

// type DummySignaling struct {
// 	v1.BaseSignaling
// }

// func (d *DummySignaling) Main(stopping v1.Stopping) error { return nil }
// func (d *DummySignaling) Start() v1.Stopped               { return nil }
// func (d *DummySignaling) Connect()                        {}
// func (d *DummySignaling) Disconnect()                     {}

// type DummySignalingOnReadyNil struct {
// 	DummySignaling
// }

// func (b *DummySignalingOnReadyNil) GetOnReady() <-chan struct{} { return nil }

// type DummyOtr struct {
// 	v1.BaseOtr
// }

// func (d *DummyOtr) Main() error { return nil }

// type DummyPunch struct {
// 	v1.BasePunch
// }

// func (d *DummyPunch) Main() error    { return nil }
// func (d *DummyPunch) RunOnce() error { return nil }

// type DummyGuard struct {
// 	v1.BaseGuard
// }

// func (d *DummyGuard) Main(stopping v1.Stopping) error { return nil }
// func (d *DummyGuard) Start() v1.Controlling           { return nil }
// func (m *DummyGuard) SetPeers(peers v1.Peers)         {}

// type DummyGuardOnConnectedNil struct {
// 	DummyGuard
// }

// func (b *DummyGuardOnConnectedNil) GetOnConnected() <-chan struct{} { return nil }

// type DummyGuardOnDisconnectedNil struct {
// 	DummyGuard
// }

// func (b *DummyGuardOnDisconnectedNil) GetOnDisconnected() <-chan struct{} { return nil }

// type DummyFlow struct {
// 	v1.BaseFlow
// }

// func (d *DummyFlow) Main() error           { return nil }
// func (d *DummyFlow) Start() v1.Controlling { return nil }

// func TestDummySignalingRunIsNil(t *testing.T) {
// 	d := new(DummySignaling)

// 	run := d.Start()
// 	if run != nil {
// 		t.Fatalf("DummySignaling.Start() = '%v', want '%v'", run, nil)
// 	}
// }

// func TestDummyPunchRunOnceIsNil(t *testing.T) {
// 	d := new(DummyPunch)

// 	err := d.RunOnce()
// 	if err != nil {
// 		t.Fatalf("DummyPunch.RunOnce() = '%v', want '%v'", err, nil)
// 	}
// }

// func TestDummyGuardRunIsNil(t *testing.T) {
// 	d := new(DummyGuard)

// 	run := d.Start()
// 	if run != nil {
// 		t.Fatalf("DummyGuard.Start() = '%v', want '%v'", run, nil)
// 	}
// }

// func TestDummyFlowRunIsNil(t *testing.T) {
// 	d := new(DummyFlow)

// 	run := d.Start()
// 	if run != nil {
// 		t.Fatalf("DummyFlow.Start() = '%v', want '%v'", run, nil)
// 	}
// }

package v1

import (
	"testing"
)

func TestBaseFlowInitSuccess(t *testing.T) {
	b := new(BaseFlow)
	err := b.Init(dummyId, dummySignal, dummyOtr, dummyPunch, dummyGuard)

	if err != nil {
		t.Fatalf("BaseFlow.init() = '%s' | want '%v'", err, nil)
	}

	if b.id != dummyId {
		t.Fatalf("BaseFlow.id = '%s' | want '%s'", b.id, dummyId)
	}

	if b.signal != dummySignal {
		t.Fatalf("BaseFlow.signal = '%v' | want '%v'", b.signal, dummySignal)
	}

	if b.otr != dummyOtr {
		t.Fatalf("BaseFlow.otr = '%v' | want '%v'", b.otr, dummyOtr)
	}

	if b.punch != dummyPunch {
		t.Fatalf("BaseFlow.punch = '%v' | want '%v'", b.punch, dummyPunch)
	}

	if b.guard != dummyGuard {
		t.Fatalf("BaseFlow.guard = '%v' | want '%v'", b.guard, dummyGuard)
	}

	if dummySignal.id != dummyId.Signal {
		t.Fatalf("BaseFlow.signal.id = '%v' | want '%v'", dummySignal.id, dummyId.Signal)
	}

	if dummyOtr.id != dummyId.Otr {
		t.Fatalf("BaseFlow.otr.id = '%v' | want '%v'", dummyOtr.id, dummyId.Otr)
	}

	if dummyGuard.id != dummyId.Guard {
		t.Fatalf("BaseFlow.guard.id = '%v' | want '%v'", dummyGuard.id, dummyId.Guard)
	}

	if dummyOtr.signalSend != dummySignal.GetSend() {
		t.Fatalf("BaseFlow.otr.signalSend = '%v' | want '%v'", dummyOtr.signalSend, dummySignal.GetSend())
	}

	if dummyOtr.signalOnReceive != dummySignal.GetOnReceive() {
		t.Fatalf("BaseFlow.otr.signalOnReceive = '%v' | want '%v'", dummyOtr.signalOnReceive, dummySignal.GetOnReceive())
	}

	if dummyPunch.send != dummyOtr.GetSend() {
		t.Fatalf("BaseFlow.punch.send = '%v' | want '%v'", dummyPunch.send, dummyOtr.GetSend())
	}

	if dummyPunch.onReceive != dummyOtr.GetOnReceive() {
		t.Fatalf("BaseFlow.punch.onReceive = '%v' | want '%v'", dummyPunch.onReceive, dummyOtr.GetOnReceive())
	}

	if dummyGuard.onPeers != dummyPunch.GetOnPeers() {
		t.Fatalf("BaseFlow.guard.onPeers ='%v' | want '%v'", dummyGuard.onPeers, dummyPunch.GetOnPeers())
	}

	if b.signalOnReady != dummySignal.GetOnReady() {
		t.Fatalf("BaseFlow.signalOnReady ='%v' | want '%v'", b.signalOnReady, dummySignal.GetOnReady())
	}

	if b.guardOnConnected != dummyGuard.GetOnConnected() {
		t.Fatalf("BaseFlow.guardOnConnected ='%v' | want '%v'", b.guardOnConnected, dummyGuard.GetOnConnected())
	}

	if b.guardOnDisconnected != dummyGuard.GetOnDisconnected() {
		t.Fatalf("BaseFlow.guardOnDisconnected ='%v' | want '%v'", b.guardOnDisconnected, dummyGuard.GetOnDisconnected())
	}
}

func TestBaseFlowIdSignalIsEmpty(t *testing.T) {
	id := Id{
		Signal: "",
		Otr:    "o",
		Guard:  "g",
	}

	b := new(BaseFlow)
	errStr := errorToString(b.Init(id, dummySignal, dummyOtr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [id.Signal]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowIdOtrIsEmpty(t *testing.T) {
	id := Id{
		Signal: "s",
		Otr:    "",
		Guard:  "g",
	}

	b := new(BaseFlow)
	errStr := errorToString(b.Init(id, dummySignal, dummyOtr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [id.Otr]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowIdGuardIsEmpty(t *testing.T) {
	id := Id{
		Signal: "s",
		Otr:    "o",
		Guard:  "",
	}

	b := new(BaseFlow)
	errStr := errorToString(b.Init(id, dummySignal, dummyOtr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [id.Guard]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowSignalIsNil(t *testing.T) {
	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, nil, dummyOtr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [signal]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowOtrIsNil(t *testing.T) {
	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, nil, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [otr]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowPunchIsNil(t *testing.T) {
	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, dummyOtr, nil, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [punch]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowGuardIsNil(t *testing.T) {
	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, dummyOtr, dummyPunch, nil))

	want := "*v1.BaseFlow.init: nil or empty parameters [guard]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowSignalGetSendIsNil(t *testing.T) {
	signal := new(DummySignal)
	signal.Init()
	signal.send = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, signal, dummyOtr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [signal.GetSend]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowSignalGetOnReceiveIsNil(t *testing.T) {
	signal := new(DummySignal)
	signal.Init()
	signal.onReceive = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, signal, dummyOtr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [signal.GetOnReceive]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowOtrGetSendIsNil(t *testing.T) {
	otr := new(DummyOtr)
	otr.Init()
	otr.punchSend = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, otr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [otr.GetSend]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowOtrGetOnReceiveIsNil(t *testing.T) {
	otr := new(DummyOtr)
	otr.Init()
	otr.punchOnReceive = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, otr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [otr.GetOnReceive]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowPunchGetOnPeersIsNil(t *testing.T) {
	punch := new(DummyPunch)
	punch.Init()
	punch.onPeers = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, dummyOtr, punch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [punch.GetOnPeers]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowSignalGetOnReadyIsNil(t *testing.T) {
	signal := new(DummySignal)
	signal.Init()
	signal.onReady = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, signal, dummyOtr, dummyPunch, dummyGuard))

	want := "*v1.BaseFlow.init: nil or empty parameters [signal.GetOnReady]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowGuardGetOnConnectedIsNil(t *testing.T) {
	guard := new(DummyGuard)
	guard.Init()
	guard.onConnected = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, dummyOtr, dummyPunch, guard))

	want := "*v1.BaseFlow.init: nil or empty parameters [guard.GetOnConnected]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

func TestBaseFlowGuardGetOnDisconnectedIsNil(t *testing.T) {
	guard := new(DummyGuard)
	guard.Init()
	guard.onDisconnected = nil

	b := new(BaseFlow)
	errStr := errorToString(b.Init(dummyId, dummySignal, dummyOtr, dummyPunch, guard))

	want := "*v1.BaseFlow.init: nil or empty parameters [guard.GetOnDisconnected]"
	if errStr != want {
		t.Fatalf("BaseFlow.init() ='%s' | want '%s'", errStr, want)
	}
}

package v1

import (
	"strings"
	"testing"
)

func TestBaseSignalInit(t *testing.T) {
	b := new(BaseSignal)
	init := b.Init()

	if init != nil {
		t.Fatalf("BaseSignal.Init() = '%s', want '%v'", init, nil)
	}

	if strings.TrimSpace(b.id) != "" {
		t.Fatalf("BaseSignal.id = '%s', want ''", b.id)
	}

	if ready, open := b.GetOnReady(), receiveIsOpen(b.GetOnReady()); ready == nil || !open {
		t.Fatalf("BaseSignal.GetOnReady() = '%v', open = '%v' | want '%s', open = '%v'", ready, open, "not <nil>", true)
	}

	if send, open := b.GetSend(), sendIsOpen(b.GetSend()); send == nil || !open {
		t.Fatalf("BaseSignal.GetSend() = '%v', open = '%v' | want '%s', open = '%v'", send, open, "not <nil>", true)
	}

	if receive, open := b.GetOnReceive(), receiveIsOpen(b.GetOnReceive()); receive == nil || !open {
		t.Fatalf("BaseSignal.GetOnReceive() = '%v', open = '%v' | want '%s', open = '%v'", receive, open, "not <nil>", true)
	}
}

func TestBaseSignalSet(t *testing.T) {
	b := new(BaseSignal)

	wantId := "id"
	b.SetId(wantId)
	if b.id != wantId {
		t.Fatalf("BaseSignal.SetId() = '%s', want '%v'", b.id, wantId)
	}
}

func TestBaseOtrInit(t *testing.T) {
	b := new(BaseOtr)
	init := b.Init()

	if init != nil {
		t.Fatalf("BaseOtr.Init() = '%s', want '%v'", init, nil)
	}

	if strings.TrimSpace(b.id) != "" {
		t.Fatalf("BaseOtr.id = '%s', want ''", b.id)
	}

	if b.signalSend != nil {
		t.Fatalf("BaseOtr.signalSend = '%v', want '%v'", b.signalSend, nil)
	}

	if b.signalOnReceive != nil {
		t.Fatalf("BaseOtr.signalOnReceive = '%v', want '%v'", b.signalOnReceive, nil)
	}

	if send, open := b.GetSend(), sendIsOpen(b.GetSend()); send == nil || !open {
		t.Fatalf("BaseOtr.GetSend() = '%v', open = '%v' | want '%s', open = '%v'", send, open, "not <nil>", true)
	}

	if receive, open := b.GetOnReceive(), receiveIsOpen(b.GetOnReceive()); receive == nil || !open {
		t.Fatalf("BaseOtr.GetOnReceive() = '%v', open = '%v' | want '%s', open = '%v'", receive, open, "not <nil>", true)
	}
}

func TestBaseOtrSet(t *testing.T) {
	b := new(BaseOtr)

	wantId := "id"
	b.SetId(wantId)
	if b.id != wantId {
		t.Fatalf("BaseOtr.SetId() = '%s', want '%v'", b.id, wantId)
	}

	wantSend := make(chan<- string)
	b.SetSend(wantSend)
	if b.signalSend != wantSend {
		t.Fatalf("BaseOtr.SetSend() = '%v', want '%v'", b.signalSend, wantSend)
	}

	wantOnReceive := make(<-chan string)
	b.SetOnReceive(wantOnReceive)
	if b.signalOnReceive != wantOnReceive {
		t.Fatalf("BaseOtr.SetOnReceive() = '%v', want '%v'", b.signalOnReceive, wantOnReceive)
	}
}

func TestBasePunchInit(t *testing.T) {
	b := new(BasePunch)
	init := b.Init()

	if init != nil {
		t.Fatalf("BasePunch.Init() = '%s', want '%v'", init, nil)
	}

	if b.send != nil {
		t.Fatalf("BasePunch.send = '%v', want '%v'", b.send, nil)
	}

	if b.onReceive != nil {
		t.Fatalf("BasePunch.signalOnReceive = '%v', want '%v'", b.onReceive, nil)
	}

	if receive, open := b.GetOnPeers(), receiveIsOpen(b.GetOnPeers()); receive == nil || !open {
		t.Fatalf("BasePunch.GetOnPeers() = '%v', open = '%v' | want '%s', open = '%v'", receive, open, "not <nil>", true)
	}
}

func TestBasePunchSet(t *testing.T) {
	b := new(BasePunch)

	wantSend := make(chan<- string)
	b.SetSend(wantSend)
	if b.send != wantSend {
		t.Fatalf("BasePunch.SetSend() = '%v', want '%v'", b.send, wantSend)
	}

	wantOnReceive := make(<-chan string)
	b.SetOnReceive(wantOnReceive)
	if b.onReceive != wantOnReceive {
		t.Fatalf("BasePunch.SetOnReceive() = '%v', want '%v'", b.onReceive, wantOnReceive)
	}
}

func TestBaseGuardInit(t *testing.T) {
	b := new(BaseGuard)
	init := b.Init()

	if init != nil {
		t.Fatalf("BaseGuard.Init() = '%s', want '%v'", init, nil)
	}

	if strings.TrimSpace(b.id) != "" {
		t.Fatalf("BaseGuard.id = '%s', want ''", b.id)
	}

	if b.onPeers != nil {
		t.Fatalf("BaseGuard.onPeers = '%v', want '%v'", b.onPeers, nil)
	}

	if connected, open := b.GetOnConnected(), receiveIsOpen(b.GetOnConnected()); connected == nil || !open {
		t.Fatalf("BaseGuard.GetOnConnected() = '%v', open = '%v' | want '%s', open = '%v'", connected, open, "not <nil>", true)
	}

	if disconnected, open := b.GetOnDisconnected(), receiveIsOpen(b.GetOnDisconnected()); disconnected == nil || !open {
		t.Fatalf("BaseGuard.GetOnDisconnected() = '%v', open = '%v' | want '%s', open = '%v'", disconnected, open, "not <nil>", true)
	}
}

func TestBaseGuardSet(t *testing.T) {
	b := new(BaseGuard)

	wantId := "id"
	b.SetId(wantId)
	if b.id != wantId {
		t.Fatalf("BaseGuard.SetId() = '%s', want '%v'", b.id, wantId)
	}

	wantOnPeers := make(<-chan Peers)
	b.SetOnPeers(wantOnPeers)
	if b.onPeers != wantOnPeers {
		t.Fatalf("BaseGuard.SetPeers() = '%v', want '%v'", b.onPeers, wantOnPeers)
	}
}

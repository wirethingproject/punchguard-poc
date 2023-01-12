package v1_test

import (
	"testing"

	v1 "github.com/punchguard/v1"
)

func TestBaseSignalInitSuccess(t *testing.T) {
	wantId := "s"

	b := new(v1.BaseSignal)
	err := b.Init(wantId)

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	if id := b.GetId(); id != wantId {
		t.Fatalf("%T.GetId() = '%v', want '%v'", b, id, wantId)
	}

	if ready, open := ReceiveIsOpen(b.GetOnReady()); ready == nil || !open {
		t.Fatalf("%T.GetOnReady() = '%v', open = '%v' | want '%s', open = '%v'", b, ready, open, "not <nil>", true)
	}

	if send, open := SendIsOpen(b.GetSend()); send == nil || !open {
		t.Fatalf("%T.GetSend() = '%v', open = '%v' | want '%s', open = '%v'", b, send, open, "not <nil>", true)
	}

	if receive, open := ReceiveIsOpen(b.GetOnReceive()); receive == nil || !open {
		t.Fatalf("%T.GetOnReceive() = '%v', open = '%v' | want '%s', open = '%v'", b, receive, open, "not <nil>", true)
	}
}

func TestBaseGuardSetIsReady(t *testing.T) {
	b := new(v1.BaseSignal)
	err := b.Init("s")

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	go b.SetIsReady()

	wantReady := *new(struct{})
	if ready := <-b.GetOnReady(); ready != wantReady {
		t.Fatalf("%T.GetOnReady() = '%v', want '%v'", b, ready, wantReady)
	}
}

func TestBaseSignalInitIdError(t *testing.T) {
	wantErr := "*v1.BaseSignal.Init id is empty ''"

	b := new(v1.BaseSignal)
	err := b.Init("")

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseOtrInitSuccess(t *testing.T) {
	wantId := "o"
	wantSignalSend := make(chan<- string)
	wantSignalOnReceive := make(<-chan string)

	b := new(v1.BaseOtr)
	err := b.Init(wantId, wantSignalSend, wantSignalOnReceive)

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	if id := b.GetId(); id != wantId {
		t.Fatalf("%T.GetId() = '%v', want '%v'", b, id, wantId)
	}

	if send := b.GetSignalSend(); send != wantSignalSend {
		t.Fatalf("%T.GetSignalSend() = '%v', want '%v'", b, send, wantSignalSend)
	}

	if receive := b.GetSignalOnReceive(); receive != wantSignalOnReceive {
		t.Fatalf("%T.GetSignalOnReceive() = '%v', want '%v'", b, receive, wantSignalOnReceive)
	}

	if send, open := SendIsOpen(b.GetPunchSend()); send == nil || !open {
		t.Fatalf("%T.GetPunchSend() = '%v', open = '%v' | want '%s', open = '%v'", b, send, open, "not <nil>", true)
	}

	if receive, open := ReceiveIsOpen(b.GetPunchOnReceive()); receive == nil || !open {
		t.Fatalf("%T.GetPunchOnReceive() = '%v', open = '%v' | want '%s', open = '%v'", b, receive, open, "not <nil>", true)
	}
}

func TestBaseOtrInitIdError(t *testing.T) {
	wantErr := "*v1.BaseOtr.Init id is empty ''"

	b := new(v1.BaseOtr)
	err := b.Init("", nil, nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseOtrInitSignalSendError(t *testing.T) {
	wantErr := "*v1.BaseOtr.Init signalSend is '<nil>'"

	b := new(v1.BaseOtr)
	err := b.Init("o", nil, nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseOtrInitSignalOnReceiveError(t *testing.T) {
	wantErr := "*v1.BaseOtr.Init signalOnReceive is '<nil>'"

	b := new(v1.BaseOtr)
	err := b.Init("o", make(chan<- string), nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBasePunchInitSuccess(t *testing.T) {
	wantSend := make(chan<- string)
	wantOnReceive := make(<-chan string)

	b := new(v1.BasePunch)
	err := b.Init(wantSend, wantOnReceive)

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	if send := b.GetSend(); send != wantSend {
		t.Fatalf("%T.GetSend() = '%v', want '%v'", b, send, wantSend)
	}

	if receive := b.GetOnReceive(); receive != wantOnReceive {
		t.Fatalf("%T.GetOnReceive() = '%v', want '%v'", b, receive, wantOnReceive)
	}

	if peers, open := ReceiveIsOpen(b.GetOnPeers()); peers == nil || !open {
		t.Fatalf("%T.GetOnPeers() = '%v', open = '%v' | want '%s', open = '%v'", b, peers, open, "not <nil>", true)
	}
}

func TestBasePunchInitSendError(t *testing.T) {
	wantErr := "*v1.BasePunch.Init send is '<nil>'"

	b := new(v1.BasePunch)
	err := b.Init(nil, nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBasePunchInitOnReceiveError(t *testing.T) {
	wantErr := "*v1.BasePunch.Init onReceive is '<nil>'"

	b := new(v1.BasePunch)
	err := b.Init(make(chan<- string), nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBasePunchSetPeers(t *testing.T) {
	b := new(v1.BasePunch)
	err := b.Init(make(chan<- string), make(<-chan string))

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	wantPeers := v1.NewPeers("local", 1, "remote", 2)
	go b.SetPeers(wantPeers)

	if peers := <-b.GetOnPeers(); peers != wantPeers {
		t.Fatalf("%T.SetPeers() = '%v', want '%v'", b, peers, wantPeers)
	}
}

func TestBaseGuardInitSuccess(t *testing.T) {
	wantId := "g"
	wantOnPeers := make(<-chan v1.Peers)

	b := new(v1.BaseGuard)
	err := b.Init(wantId, wantOnPeers)

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	if id := b.GetId(); id != wantId {
		t.Fatalf("%T.GetId() = '%v', want '%v'", b, id, wantId)
	}

	if peers := b.GetOnPeers(); peers != wantOnPeers {
		t.Fatalf("%T.GetOnPeers() = '%v', want '%v'", b, peers, wantOnPeers)
	}

	if connect, open := ReceiveIsOpen(b.GetOnConnected()); connect == nil || !open {
		t.Fatalf("%T.GetOnConnected() = '%v', open = '%v' | want '%s', open = '%v'", b, connect, open, "not <nil>", true)
	}

	if disconnect, open := ReceiveIsOpen(b.GetOnDisconnected()); disconnect == nil || !open {
		t.Fatalf("%T.GetOnDisconnected() = '%v', open = '%v' | want '%s', open = '%v'", b, disconnect, open, "not <nil>", true)
	}
}
func TestBaseGuardInitIdError(t *testing.T) {
	wantErr := "*v1.BaseGuard.Init id is empty ''"

	b := new(v1.BaseGuard)
	err := b.Init("", nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseGuardInitOnPeersError(t *testing.T) {
	wantErr := "*v1.BaseGuard.Init onPeers is '<nil>'"

	b := new(v1.BaseGuard)
	err := b.Init("o", nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseGuardSetConnected(t *testing.T) {
	b := new(v1.BaseGuard)
	err := b.Init("o", make(<-chan v1.Peers))

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	go b.SetConnected()

	wantConnected := *new(struct{})
	if connected := <-b.GetOnConnected(); connected != wantConnected {
		t.Fatalf("%T.SetConnected() = '%v', want '%v'", b, connected, wantConnected)
	}
}

func TestBaseGuardSetDisconnected(t *testing.T) {
	b := new(v1.BaseGuard)
	err := b.Init("o", make(<-chan v1.Peers))

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	go b.SetDisconnected()

	wantDisconnected := *new(struct{})
	if Disconnected := <-b.GetOnDisconnected(); Disconnected != wantDisconnected {
		t.Fatalf("%T.SetDisconnected() = '%v', want '%v'", b, Disconnected, wantDisconnected)
	}
}

func TestBaseFlowInitSuccess(t *testing.T) {

	b := new(v1.BaseFlow)
	err := b.Init(DummySignalTest, DummyOtrTest, DummyPunchTest, DummyGuardTest)

	if err != nil {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
	}

	if signal := b.GetSignal(); signal != DummySignalTest {
		t.Fatalf("%T.GetSignal() = '%v', want '%v'", b, signal, DummySignalTest)
	}

	if otr := b.GetOtr(); otr != DummyOtrTest {
		t.Fatalf("%T.GetOtr() = '%v', want '%v'", b, otr, DummyOtrTest)
	}

	if punch := b.GetPunch(); punch != DummyPunchTest {
		t.Fatalf("%T.GetPunch() = '%v', want '%v'", b, punch, DummyPunchTest)
	}

	if guard := b.GetGuard(); guard != DummyGuardTest {
		t.Fatalf("%T.GetGuard() = '%v', want '%v'", b, guard, DummyGuardTest)
	}

	if ready := b.GetSignalOnReady(); ready != DummySignalTest.GetOnReady() {
		t.Fatalf("%T.GetSignalOnReady() = '%v', want '%v'", b, ready, DummySignalTest.GetOnReady())
	}

	if connected := b.GetGuardOnConnected(); connected != DummyGuardTest.GetOnConnected() {
		t.Fatalf("%T.GetGuardOnConnected() = '%v', want '%v'", b, connected, DummyGuardTest.GetOnConnected())
	}

	if Disconnected := b.GetGuardOnDisconnected(); Disconnected != DummyGuardTest.GetOnDisconnected() {
		t.Fatalf("%T.GetGuardOnDisconnected() = '%v', want '%v'", b, Disconnected, DummyGuardTest.GetOnDisconnected())
	}
}

func TestBaseFlowInitSignalError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init signal is '<nil>'"

	b := new(v1.BaseFlow)
	err := b.Init(nil, DummyOtrTest, DummyPunchTest, DummyGuardTest)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseFlowInitOtrError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init otr is '<nil>'"

	b := new(v1.BaseFlow)
	err := b.Init(DummySignalTest, nil, DummyPunchTest, DummyGuardTest)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseFlowInitPunchError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init punch is '<nil>'"

	b := new(v1.BaseFlow)
	err := b.Init(DummySignalTest, DummyOtrTest, nil, DummyGuardTest)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseFlowInitGuardError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init guard is '<nil>'"

	b := new(v1.BaseFlow)
	err := b.Init(DummySignalTest, DummyOtrTest, DummyPunchTest, nil)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseFlowInitSignalOnReadyError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init signal.GetOnReady() is '<nil>'"

	signal := new(DummySignalOnReadyNil)
	signal.Init("s")

	b := new(v1.BaseFlow)
	err := b.Init(signal, DummyOtrTest, DummyPunchTest, DummyGuardTest)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseFlowInitGuardOnConnectedError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init guard.GetOnConnected() is '<nil>'"

	guard := new(DummyGuardOnConnectedNil)
	guard.Init("g", make(<-chan v1.Peers))

	b := new(v1.BaseFlow)
	err := b.Init(DummySignalTest, DummyOtrTest, DummyPunchTest, guard)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

func TestBaseFlowInitGuardOnDisconnectedError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init guard.GetOnDisconnected() is '<nil>'"

	guard := new(DummyGuardOnDisconnectedNil)
	guard.Init("g", make(<-chan v1.Peers))

	b := new(v1.BaseFlow)
	err := b.Init(DummySignalTest, DummyOtrTest, DummyPunchTest, guard)

	if err == nil {
		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
	}
}

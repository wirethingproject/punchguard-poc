package punchguard_test

// import (
// 	"testing"

// 	v1 "github.com/punchguard/v1"
// )

// func TestBaseSignalingInitSuccess(t *testing.T) {
// 	wantId := "s"

// 	b := new(v1.BaseSignaling)
// 	err := b.Init(wantId)

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	if id := b.GetId(); id != wantId {
// 		t.Fatalf("%T.GetId() = '%v', want '%v'", b, id, wantId)
// 	}

// 	if ready, open := ReceiveIsOpen(b.GetOnReady()); ready == nil || !open {
// 		t.Fatalf("%T.GetOnReady() = '%v', open = '%v' | want '%s', open = '%v'", b, ready, open, "not <nil>", true)
// 	}

// 	if send, open := SendIsOpen(b.GetSend()); send == nil || !open {
// 		t.Fatalf("%T.GetSend() = '%v', open = '%v' | want '%s', open = '%v'", b, send, open, "not <nil>", true)
// 	}

// 	if receive, open := ReceiveIsOpen(b.GetOnReceive()); receive == nil || !open {
// 		t.Fatalf("%T.GetOnReceive() = '%v', open = '%v' | want '%s', open = '%v'", b, receive, open, "not <nil>", true)
// 	}
// }

// func TestBaseGuardSetIsReady(t *testing.T) {
// 	b := new(v1.BaseSignaling)
// 	err := b.Init("s")

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	go b.SetOnReady()

// 	wantReady := *new(struct{})
// 	if ready := <-b.GetOnReady(); ready != wantReady {
// 		t.Fatalf("%T.GetOnReady() = '%v', want '%v'", b, ready, wantReady)
// 	}
// }

// func TestBaseSignalingInitIdError(t *testing.T) {
// 	wantErr := "*v1.BaseSignaling.Init id is empty ''"

// 	b := new(v1.BaseSignaling)
// 	err := b.Init("")

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseOtrInitSuccess(t *testing.T) {
// 	wantId := "o"
// 	wantSignalingSend := make(chan<- string)
// 	wantSignalingOnReceive := make(<-chan string)

// 	b := new(v1.BaseOtr)
// 	err := b.Init(wantId, wantSignalingSend, wantSignalingOnReceive)

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	if id := b.GetId(); id != wantId {
// 		t.Fatalf("%T.GetId() = '%v', want '%v'", b, id, wantId)
// 	}

// 	if send := b.GetSignalingSend(); send != wantSignalingSend {
// 		t.Fatalf("%T.GetSignalingSend() = '%v', want '%v'", b, send, wantSignalingSend)
// 	}

// 	if receive := b.GetSignalingOnReceive(); receive != wantSignalingOnReceive {
// 		t.Fatalf("%T.GetSignalingOnReceive() = '%v', want '%v'", b, receive, wantSignalingOnReceive)
// 	}

// 	if send, open := SendIsOpen(b.GetPunchSend()); send == nil || !open {
// 		t.Fatalf("%T.GetPunchSend() = '%v', open = '%v' | want '%s', open = '%v'", b, send, open, "not <nil>", true)
// 	}

// 	if receive, open := ReceiveIsOpen(b.GetPunchOnReceive()); receive == nil || !open {
// 		t.Fatalf("%T.GetPunchOnReceive() = '%v', open = '%v' | want '%s', open = '%v'", b, receive, open, "not <nil>", true)
// 	}
// }

// func TestBaseOtrInitIdError(t *testing.T) {
// 	wantErr := "*v1.BaseOtr.Init id is empty ''"

// 	b := new(v1.BaseOtr)
// 	err := b.Init("", nil, nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseOtrInitSignalingSendError(t *testing.T) {
// 	wantErr := "*v1.BaseOtr.Init signalingSend is '<nil>'"

// 	b := new(v1.BaseOtr)
// 	err := b.Init("o", nil, nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseOtrInitSignalingOnReceiveError(t *testing.T) {
// 	wantErr := "*v1.BaseOtr.Init signalingOnReceive is '<nil>'"

// 	b := new(v1.BaseOtr)
// 	err := b.Init("o", make(chan<- string), nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBasePunchInitSuccess(t *testing.T) {
// 	wantSend := make(chan<- string)
// 	wantOnReceive := make(<-chan string)

// 	b := new(v1.BasePunch)
// 	err := b.Init(wantSend, wantOnReceive)

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	if send := b.GetSend(); send != wantSend {
// 		t.Fatalf("%T.GetSend() = '%v', want '%v'", b, send, wantSend)
// 	}

// 	if receive := b.GetOnReceive(); receive != wantOnReceive {
// 		t.Fatalf("%T.GetOnReceive() = '%v', want '%v'", b, receive, wantOnReceive)
// 	}

// 	if peers, open := ReceiveIsOpen(b.GetOnPeers()); peers == nil || !open {
// 		t.Fatalf("%T.GetOnPeers() = '%v', open = '%v' | want '%s', open = '%v'", b, peers, open, "not <nil>", true)
// 	}
// }

// func TestBasePunchInitSendError(t *testing.T) {
// 	wantErr := "*v1.BasePunch.Init send is '<nil>'"

// 	b := new(v1.BasePunch)
// 	err := b.Init(nil, nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBasePunchInitOnReceiveError(t *testing.T) {
// 	wantErr := "*v1.BasePunch.Init onReceive is '<nil>'"

// 	b := new(v1.BasePunch)
// 	err := b.Init(make(chan<- string), nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBasePunchSetPeers(t *testing.T) {
// 	b := new(v1.BasePunch)
// 	err := b.Init(make(chan<- string), make(<-chan string))

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	wantPeers := v1.NewPeers("local", 1, "remote", 2)
// 	go b.SetOnPeers(wantPeers)

// 	if peers := <-b.GetOnPeers(); peers != wantPeers {
// 		t.Fatalf("%T.SetPeers() = '%v', want '%v'", b, peers, wantPeers)
// 	}
// }

// func TestBaseGuardInitSuccess(t *testing.T) {
// 	wantId := "g"
// 	wantOnPeers := make(<-chan v1.Peers)

// 	b := new(v1.BaseGuard)
// 	err := b.Init(wantId, wantOnPeers)

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	if id := b.GetId(); id != wantId {
// 		t.Fatalf("%T.GetId() = '%v', want '%v'", b, id, wantId)
// 	}

// 	if peers := b.GetOnPeers(); peers != wantOnPeers {
// 		t.Fatalf("%T.GetOnPeers() = '%v', want '%v'", b, peers, wantOnPeers)
// 	}

// 	if connect, open := ReceiveIsOpen(b.GetOnConnected()); connect == nil || !open {
// 		t.Fatalf("%T.GetOnConnected() = '%v', open = '%v' | want '%s', open = '%v'", b, connect, open, "not <nil>", true)
// 	}

// 	if disconnect, open := ReceiveIsOpen(b.GetOnDisconnected()); disconnect == nil || !open {
// 		t.Fatalf("%T.GetOnDisconnected() = '%v', open = '%v' | want '%s', open = '%v'", b, disconnect, open, "not <nil>", true)
// 	}
// }
// func TestBaseGuardInitIdError(t *testing.T) {
// 	wantErr := "*v1.BaseGuard.Init id is empty ''"

// 	b := new(v1.BaseGuard)
// 	err := b.Init("", nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseGuardInitOnPeersError(t *testing.T) {
// 	wantErr := "*v1.BaseGuard.Init onPeers is '<nil>'"

// 	b := new(v1.BaseGuard)
// 	err := b.Init("o", nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseGuardSetConnected(t *testing.T) {
// 	b := new(v1.BaseGuard)
// 	err := b.Init("o", make(<-chan v1.Peers))

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	go b.SetOnConnected()

// 	wantConnected := *new(struct{})
// 	if connected := <-b.GetOnConnected(); connected != wantConnected {
// 		t.Fatalf("%T.SetConnected() = '%v', want '%v'", b, connected, wantConnected)
// 	}
// }

// func TestBaseGuardSetDisconnected(t *testing.T) {
// 	b := new(v1.BaseGuard)
// 	err := b.Init("o", make(<-chan v1.Peers))

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	go b.SetOnDisconnected()

// 	wantDisconnected := *new(struct{})
// 	if Disconnected := <-b.GetOnDisconnected(); Disconnected != wantDisconnected {
// 		t.Fatalf("%T.SetDisconnected() = '%v', want '%v'", b, Disconnected, wantDisconnected)
// 	}
// }

// func TestBaseFlowInitSuccess(t *testing.T) {

// 	b := new(v1.BaseFlow)
// 	err := b.Init(DummySignalingTest, DummyOtrTest, DummyPunchTest, DummyGuardTest)

// 	if err != nil {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, nil)
// 	}

// 	if signaling := b.GetSignaling(); signaling != DummySignalingTest {
// 		t.Fatalf("%T.GetSignaling() = '%v', want '%v'", b, signaling, DummySignalingTest)
// 	}

// 	if otr := b.GetOtr(); otr != DummyOtrTest {
// 		t.Fatalf("%T.GetOtr() = '%v', want '%v'", b, otr, DummyOtrTest)
// 	}

// 	if punch := b.GetPunch(); punch != DummyPunchTest {
// 		t.Fatalf("%T.GetPunch() = '%v', want '%v'", b, punch, DummyPunchTest)
// 	}

// 	if guard := b.GetGuard(); guard != DummyGuardTest {
// 		t.Fatalf("%T.GetGuard() = '%v', want '%v'", b, guard, DummyGuardTest)
// 	}
// }

// func TestBaseFlowInitSignalingError(t *testing.T) {
// 	wantErr := "*v1.BaseFlow.Init signaling is '<nil>'"

// 	b := new(v1.BaseFlow)
// 	err := b.Init(nil, DummyOtrTest, DummyPunchTest, DummyGuardTest)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseFlowInitOtrError(t *testing.T) {
// 	wantErr := "*v1.BaseFlow.Init otr is '<nil>'"

// 	b := new(v1.BaseFlow)
// 	err := b.Init(DummySignalingTest, nil, DummyPunchTest, DummyGuardTest)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseFlowInitPunchError(t *testing.T) {
// 	wantErr := "*v1.BaseFlow.Init punch is '<nil>'"

// 	b := new(v1.BaseFlow)
// 	err := b.Init(DummySignalingTest, DummyOtrTest, nil, DummyGuardTest)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseFlowInitGuardError(t *testing.T) {
// 	wantErr := "*v1.BaseFlow.Init guard is '<nil>'"

// 	b := new(v1.BaseFlow)
// 	err := b.Init(DummySignalingTest, DummyOtrTest, DummyPunchTest, nil)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseFlowInitSignalingOnReadyError(t *testing.T) {
// 	wantErr := "*v1.BaseFlow.Init signaling.GetOnReady() is '<nil>'"

// 	signaling := new(DummySignalingOnReadyNil)
// 	signaling.Init("s")

// 	b := new(v1.BaseFlow)
// 	err := b.Init(signaling, DummyOtrTest, DummyPunchTest, DummyGuardTest)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseFlowInitGuardOnConnectedError(t *testing.T) {
// 	wantErr := "*v1.BaseFlow.Init guard.GetOnConnected() is '<nil>'"

// 	guard := new(DummyGuardOnConnectedNil)
// 	guard.Init("g", make(<-chan v1.Peers))

// 	b := new(v1.BaseFlow)
// 	err := b.Init(DummySignalingTest, DummyOtrTest, DummyPunchTest, guard)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

// func TestBaseFlowInitGuardOnDisconnectedError(t *testing.T) {
// 	wantErr := "*v1.BaseFlow.Init guard.GetOnDisconnected() is '<nil>'"

// 	guard := new(DummyGuardOnDisconnectedNil)
// 	guard.Init("g", make(<-chan v1.Peers))

// 	b := new(v1.BaseFlow)
// 	err := b.Init(DummySignalingTest, DummyOtrTest, DummyPunchTest, guard)

// 	if err == nil {
// 		t.Fatalf("%T.Init() = '%v', want not '%v'", b, err, nil)
// 	}

// 	if err.Error() != wantErr {
// 		t.Fatalf("%T.Init() = '%v', want '%v'", b, err, wantErr)
// 	}
// }

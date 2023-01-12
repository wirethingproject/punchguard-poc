package v1_test

import (
	"testing"

	v1 "github.com/punchguard/v1"
)

func TestNewSignalSuccess(t *testing.T) {
	wantId := "s"
	value, err := v1.NewSignal[DummySignal](wantId)

	if err != nil {
		t.Fatalf("NewSignal.err = '%v'", err)
	}

	if value == nil {
		t.Fatalf("NewSignal = '%v' | want not '%v'", value, nil)
	}

	if id := value.GetId(); id != wantId {
		t.Fatalf("NewSignal.GetId() = '%v' | want '%v'", id, wantId)
	}
}

func TestNewSignalInitError(t *testing.T) {
	wantErr := "*v1.BaseSignal.Init id is empty ''"

	value, err := v1.NewSignal[DummySignal]("")

	if err == nil {
		t.Fatalf("NewSignal.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewSignal.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewSignal = '%v' | want '%v'", value, nil)
	}
}

func TestNewSignalTypeError(t *testing.T) {
	wantErr := "*int is not *v1.Signal"

	value, err := v1.NewSignal[int]("")

	if err == nil {
		t.Fatalf("NewSignal.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewSignal.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewSignal = '%v' | want '%v'", value, nil)
	}
}

func TestNewOtrSuccess(t *testing.T) {
	wantId := "o"
	value, err := v1.NewOtr[DummyOtr](wantId, make(chan<- string), make(<-chan string))

	if err != nil {
		t.Fatalf("NewOtr.err = '%v'", err)
	}

	if value == nil {
		t.Fatalf("NewOtr = '%v' | want not '%v'", value, nil)
	}

	if id := value.GetId(); id != wantId {
		t.Fatalf("NewOtr.GetId() = '%v' | want '%v'", id, wantId)
	}
}

func TestNewOtrInitError(t *testing.T) {
	wantErr := "*v1.BaseOtr.Init id is empty ''"

	value, err := v1.NewOtr[DummyOtr]("", nil, nil)

	if err == nil {
		t.Fatalf("NewOtr.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewOtr.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewOtr = '%v' | want '%v'", value, nil)
	}
}

func TestNewOtrTypeError(t *testing.T) {
	wantErr := "*int is not *v1.Otr"

	value, err := v1.NewOtr[int]("", nil, nil)

	if err == nil {
		t.Fatalf("NewOtr.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewOtr.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewOtr = '%v' | want '%v'", value, nil)
	}
}

func TestNewPunchSuccess(t *testing.T) {
	wantSend := make(chan<- string)
	value, err := v1.NewPunch[DummyPunch](wantSend, make(<-chan string))

	if err != nil {
		t.Fatalf("NewPunch.err = '%v'", err)
	}

	if value == nil {
		t.Fatalf("NewPunch = '%v' | want not '%v'", value, nil)
	}

	if send := value.GetSend(); send != wantSend {
		t.Fatalf("NewPunch.GetSend() = '%v' | want '%v'", send, wantSend)
	}
}

func TestNewPunchInitError(t *testing.T) {
	wantErr := "*v1.BasePunch.Init send is '<nil>'"

	value, err := v1.NewPunch[DummyPunch](nil, nil)

	if err == nil {
		t.Fatalf("NewPunch.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewPunch.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewPunch = '%v' | want '%v'", value, nil)
	}
}

func TestNewPunchTypeError(t *testing.T) {
	wantErr := "*int is not *v1.Punch"

	value, err := v1.NewPunch[int](nil, nil)

	if err == nil {
		t.Fatalf("NewPunch.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewPunch.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewPunch = '%v' | want '%v'", value, nil)
	}
}

func TestNewGuardSuccess(t *testing.T) {
	wantId := "g"
	value, err := v1.NewGuard[DummyGuard](wantId, make(<-chan v1.Peers))

	if err != nil {
		t.Fatalf("NewGuard.err = '%v'", err)
	}

	if value == nil {
		t.Fatalf("NewGuard = '%v' | want not '%v'", value, nil)
	}

	if id := value.GetId(); id != wantId {
		t.Fatalf("NewGuard.GetId() = '%v' | want '%v'", id, wantId)
	}
}

func TestNewGuardInitError(t *testing.T) {
	wantErr := "*v1.BaseGuard.Init id is empty ''"

	value, err := v1.NewGuard[DummyGuard]("", nil)

	if err == nil {
		t.Fatalf("NewGuard.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewGuard.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewGuard = '%v' | want '%v'", value, nil)
	}
}

func TestNewGuardTypeError(t *testing.T) {
	wantErr := "*int is not *v1.Guard"

	value, err := v1.NewGuard[int]("", nil)

	if err == nil {
		t.Fatalf("NewGuard.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewGuard.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewGuard = '%v' | want '%v'", value, nil)
	}
}

func TestNewFlowSuccess(t *testing.T) {
	wantSignal := DummySignalTest
	value, err := v1.NewFlow[DummyFlow](wantSignal, DummyOtrTest, DummyPunchTest, DummyGuardTest)

	if err != nil {
		t.Fatalf("NewFlow.err = '%v'", err)
	}

	if value == nil {
		t.Fatalf("NewFlow = '%v' | want not '%v'", value, nil)
	}

	if signal := value.GetSignal(); signal != wantSignal {
		t.Fatalf("NewFlow.GetSignal() = '%v' | want '%v'", signal, wantSignal)
	}
}

func TestNewFlowInitError(t *testing.T) {
	wantErr := "*v1.BaseFlow.Init signal is '<nil>'"

	value, err := v1.NewFlow[DummyFlow](nil, nil, nil, nil)

	if err == nil {
		t.Fatalf("NewFlow.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewFlow.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewFlow = '%v' | want '%v'", value, nil)
	}
}

func TestNewFlowTypeError(t *testing.T) {
	wantErr := "*int is not *v1.Flow"

	value, err := v1.NewFlow[int](nil, nil, nil, nil)

	if err == nil {
		t.Fatalf("NewFlow.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("NewFlow.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("NewFlow = '%v' | want '%v'", value, nil)
	}
}

func TestInitSuccess(t *testing.T) {
	wantId := v1.NewId("s", "o", "g")
	value, err := v1.Init[DummySignal, DummyOtr, DummyPunch, DummyGuard, DummyFlow](wantId)

	if err != nil {
		t.Fatalf("Init.err = '%v'", err)
	}

	if value == nil {
		t.Fatalf("Init = '%v' | want not '%v'", value, nil)
	}

	if signal := value.GetSignal(); signal == nil {
		t.Fatalf("Init.GetSignal() = '%v' | want not '%v'", signal, nil)
	}

	if otr := value.GetOtr(); otr == nil {
		t.Fatalf("Init.GetOtr() = '%v' | want not '%v'", otr, nil)
	}

	if punch := value.GetPunch(); punch == nil {
		t.Fatalf("Init.GetPunch() = '%v' | want not '%v'", punch, nil)
	}

	if guard := value.GetGuard(); guard == nil {
		t.Fatalf("Init.GetGuard() = '%v' | want not '%v'", guard, nil)
	}

	if id := value.GetSignal().GetId(); id != wantId.Signal {
		t.Fatalf("Init.GetSignal().GetId() = '%v' | want '%v'", id, wantId.Signal)
	}

	if id := value.GetOtr().GetId(); id != wantId.Otr {
		t.Fatalf("Init.GetOtr().GetId() = '%v' | want '%v'", id, wantId.Otr)
	}

	if send := value.GetOtr().GetSignalSend(); send != value.GetSignal().GetSend() {
		t.Fatalf("Init.GetOtr().GetSignalSend() = '%v' | want '%v'", send, value.GetSignal().GetSend())
	}

	if receive := value.GetOtr().GetSignalOnReceive(); receive != value.GetSignal().GetOnReceive() {
		t.Fatalf("Init.GetOtr().GetSignalOnReceive() = '%v' | want '%v'", receive, value.GetSignal().GetOnReceive())
	}

	if send := value.GetPunch().GetSend(); send != value.GetOtr().GetPunchSend() {
		t.Fatalf("Init.GetPunch().GetSend() = '%v' | want '%v'", send, value.GetOtr().GetPunchSend())
	}

	if receive := value.GetPunch().GetOnReceive(); receive != value.GetOtr().GetPunchOnReceive() {
		t.Fatalf("Init.GetPunch().GetOnReceive() = '%v' | want '%v'", receive, value.GetOtr().GetPunchOnReceive())
	}

	if id := value.GetGuard().GetId(); id != wantId.Guard {
		t.Fatalf("Init.GetGuard().GetId() = '%v' | want '%v'", id, wantId.Guard)
	}

	if peers := value.GetGuard().GetOnPeers(); peers != value.GetPunch().GetOnPeers() {
		t.Fatalf("Init.GetGuard().GetOnPeers() = '%v' | want '%v'", peers, value.GetPunch().GetOnPeers())
	}
}

func TestInitSignalError(t *testing.T) {
	wantErr := "*int is not *v1.Signal"
	value, err := v1.Init[int, DummyOtr, DummyPunch, DummyGuard, DummyFlow](v1.NewId("s", "o", "g"))

	if err == nil {
		t.Fatalf("Init.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("Init.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("Init = '%v' | want '%v'", value, nil)
	}
}

func TestInitOtrError(t *testing.T) {
	wantErr := "*int is not *v1.Otr"
	value, err := v1.Init[DummySignal, int, DummyPunch, DummyGuard, DummyFlow](v1.NewId("s", "o", "g"))

	if err == nil {
		t.Fatalf("Init.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("Init.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("Init = '%v' | want '%v'", value, nil)
	}
}

func TestInitPunchError(t *testing.T) {
	wantErr := "*int is not *v1.Punch"
	value, err := v1.Init[DummySignal, DummyOtr, int, DummyGuard, DummyFlow](v1.NewId("s", "o", "g"))

	if err == nil {
		t.Fatalf("Init.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("Init.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("Init = '%v' | want '%v'", value, nil)
	}
}

func TestInitGuardError(t *testing.T) {
	wantErr := "*int is not *v1.Guard"
	value, err := v1.Init[DummySignal, DummyOtr, DummyPunch, int, DummyFlow](v1.NewId("s", "o", "g"))

	if err == nil {
		t.Fatalf("Init.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("Init.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("Init = '%v' | want '%v'", value, nil)
	}
}

func TestInitFlowError(t *testing.T) {
	wantErr := "*int is not *v1.Flow"
	value, err := v1.Init[DummySignal, DummyOtr, DummyPunch, DummyGuard, int](v1.NewId("s", "o", "g"))

	if err == nil {
		t.Fatalf("Init.err = '%v' | want not '%v'", err, nil)
	}

	if err.Error() != wantErr {
		t.Fatalf("Init.err.Error() = '%v' | want '%v'", err.Error(), wantErr)
	}

	if value != nil {
		t.Fatalf("Init = '%v' | want '%v'", value, nil)
	}
}

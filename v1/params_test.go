package v1

import "testing"

func TestNewFlowParamsSuccess(t *testing.T) {
	var (
		signal Signal
		otr    Otr
		punch  Punch
		guard  Guard
	)

	signal, otr, punch, guard = NewFlowParams[DummySignal, DummyOtr, DummyPunch, DummyGuard]()

	if signal.(*DummySignal) == nil {
		t.Fatalf("NewFlowParams.signal = '%v' | want not '%v'", signal, nil)
	}

	if otr.(*DummyOtr) == nil {
		t.Fatalf("NewFlowParams.otr = '%v' | want not '%v'", otr, nil)
	}

	if punch.(*DummyPunch) == nil {
		t.Fatalf("NewFlowParams.punch = '%v' | want not '%v'", punch, nil)
	}

	if guard.(*DummyGuard) == nil {
		t.Fatalf("NewFlowParams.guard = '%v' | want not '%v'", guard, nil)
	}
}

func TestInitFlowParamsSuccess(t *testing.T) {
	signal, otr, punch, guard := NewFlowParams[DummySignal, DummyOtr, DummyPunch, DummyGuard]()

	err := InitFlowParams(signal, otr, punch, guard)

	if err != nil {
		t.Fatalf("InitFlowParams.err = '%v' | want '%v'", err, nil)
	}

	if signal.onReceive == nil {
		t.Fatal("signal.Init() was not called")
	}

	if otr.punchOnReceive == nil {
		t.Fatal("otr.Init() was not called")
	}

	if punch.onPeers == nil {
		t.Fatal("punch.Init() was not called")
	}

	if guard.onDisconnected == nil {
		t.Fatal("guard.Init() was not called")
	}
}

func TestInitFlowParamsError(t *testing.T) {
	signal, otr, punch, guard := NewFlowParams[DummySignalInitFail, DummyOtr, DummyPunch, DummyGuard]()

	err := InitFlowParams(signal, otr, punch, guard)

	if err == nil {
		t.Fatalf("InitFlowParams.err = '%v' | want not '%v'", err, nil)
	}
}

func TestInvalidParamAppend(t *testing.T) {
	invalid := new(invalidParams)

	wantName := "param"
	invalid.append("param")

	wantLen := 1
	if len(invalid.get()) != wantLen {
		t.Fatalf("invalid.len = '%v', want '%v'", len(invalid.get()), wantLen)
	}

	if invalid.get()[0] != wantName {
		t.Fatalf("invalid.get() = '%v', want '%v'", invalid.get()[0], wantName)
	}
}

func TestGetAndSetOrInvalidSuccess(t *testing.T) {
	invalid := new(invalidParams)

	var setted <-chan string
	wantSetted := make(<-chan string)

	getAndSetOrInvalid(invalid, "get", func() <-chan string {
		return wantSetted
	}, func(c <-chan string) {
		setted = c
	})

	wantLen := 0
	if len(invalid.get()) != wantLen {
		t.Fatalf("invalid.len = '%v', want '%v'", len(invalid.get()), wantLen)
	}

	if setted != wantSetted {
		t.Fatalf("setted = '%v', want '%v'", setted, wantSetted)
	}
}

func TestGetAndSetOrInvalidNil(t *testing.T) {
	invalid := new(invalidParams)

	var setted <-chan string
	wantName := "get"

	getAndSetOrInvalid(invalid, wantName, func() <-chan string {
		return nil
	}, func(c <-chan string) {
		setted = c
	})

	wantLen := 1
	if len(invalid.get()) != wantLen {
		t.Fatalf("invalid.len = '%v', want '%v'", len(invalid.get()), wantLen)
	}

	if invalid.get()[0] != wantName {
		t.Fatalf("invalid.get() = '%v', want '%v'", invalid.get()[0], wantName)
	}

	var wantSetted <-chan string
	if setted != wantSetted {
		t.Fatalf("setted = '%v', want '%v'", setted, wantSetted)
	}
}

func TestSetStringOrInvalidSuccess(t *testing.T) {
	invalid := new(invalidParams)

	var setted string
	want := "s"

	setStringOrInvalid(invalid, "value", want, func(s string) {
		setted = s
	})

	wantLen := 0
	if len(invalid.get()) != wantLen {
		t.Fatalf("invalid.len = '%v', want '%v'", len(invalid.get()), wantLen)
	}

	if setted != want {
		t.Fatalf("setted = '%v', want '%v'", setted, want)
	}
}

func TestSetStringOrInvalidEmpty(t *testing.T) {
	invalid := new(invalidParams)

	var setted string
	wantName := "value"
	wantSetted := ""

	setStringOrInvalid(invalid, wantName, wantSetted, func(s string) {
		setted = s
	})

	wantLen := 1
	if len(invalid.get()) != wantLen {
		t.Fatalf("invalid.len = '%v', want '%v'", len(invalid.get()), wantLen)
	}

	if invalid.get()[0] != wantName {
		t.Fatalf("invalid.get() = '%v', want '%v'", invalid.get()[0], wantName)
	}

	if setted != wantSetted {
		t.Fatalf("setted = '%v', want '%v'", setted, wantSetted)
	}
}

func TestSetValueOrInvalidSuccess(t *testing.T) {
	invalid := new(invalidParams)

	var setted string
	want := "s"

	setValueOrInvalid(invalid, "value", false, want, func(s string) {
		setted = s
	})

	wantLen := 0
	if len(invalid.get()) != wantLen {
		t.Fatalf("invalid.len = '%v', want '%v'", len(invalid.get()), wantLen)
	}

	if setted != want {
		t.Fatalf("setted = '%v', want '%v'", setted, want)
	}
}

func TestSetValueOrInvalidNil(t *testing.T) {
	invalid := new(invalidParams)

	var setted string
	wantName := "value"
	wantSetted := ""

	setValueOrInvalid(invalid, wantName, true, wantSetted, func(s string) {
		setted = s
	})

	wantLen := 1
	if len(invalid.get()) != wantLen {
		t.Fatalf("invalid.len = '%v', want '%v'", len(invalid.get()), wantLen)
	}

	if invalid.get()[0] != wantName {
		t.Fatalf("invalid.get() = '%v', want '%v'", invalid.get()[0], wantName)
	}

	if setted != wantSetted {
		t.Fatalf("setted = '%v', want '%v'", setted, wantSetted)
	}
}

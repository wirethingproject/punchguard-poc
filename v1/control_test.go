package v1_test

import (
	"testing"
	"time"

	v1 "github.com/punchguard/v1"
)

type testControl struct{}

func TestNewControlSuccess(t *testing.T) {
	wantControlled := new(testControl)

	controlling, controlled := v1.NewControl(wantControlled)

	if c, is := controlling.(*v1.Control); is == false {
		t.Fatalf("NewControl.controlling = '%T' | want '%T'", controlling, c)
	}

	if c, is := controlled.(*v1.Control); is == false {
		t.Fatalf("NewControl.controlled = '%T' | want '%T'", controlled, c)
	}

	if controlling.(*v1.Control) != controlled.(*v1.Control) {
		t.Fatalf("NewControl.controlling = '%v' and NewControl.controlled '%v' differ | want equal", controlling, controlled)
	}

	if c := controlled.(*v1.Control).GetControlled(); c != wantControlled {
		t.Fatalf("%T.GetControlled() = '%v' | want '%v'", controlled, c, wantControlled)
	}

	if stopping, open := ReceiveIsOpen(controlled.GetStopping()); stopping == nil || !open {
		t.Fatalf("%T.GetStopping() = '%v', open = '%v' | want '%s', open = '%v'", controlled, stopping, open, "not <nil>", true)
	}
}

func TestControlStopAndClose(t *testing.T) {
	want := *new(struct{})
	wantCloseWait := make(chan struct{})
	controlling, controlled := v1.NewControl(new(testControl))

	go func() {
		time.Sleep(1 * time.Second)
		controlling.Stop(new(testControl))
		wantCloseWait <- want
	}()

	if stopping := <-controlled.GetStopping(); stopping != want {
		t.Fatalf("%T.GetStopping() = '%v', want '%v'", controlled, stopping, want)
	}

	time.Sleep(1 * time.Second)
	controlled.Close()

	if stopping, open := ReceiveIsOpen(controlled.GetStopping()); stopping == nil || open {
		t.Fatalf("%T.GetStopping() = '%v', open = '%v' | want '%s', open = '%v'", controlled, stopping, open, "not <nil>", true)
	}

	if closed := <-wantCloseWait; closed != want {
		t.Fatalf("%T.wantCloseWait = '%v' | want '%v'", controlled, closed, want)
	}
}

package v1

import "testing"

func TestNewOnDemandSuccess(t *testing.T) {
	o, err := NewOnDemand(dummyId, dummySignal, dummyOtr, dummyPunch, dummyGuard)

	if err != nil {
		t.Fatalf("NewOnDemand.err = '%v' | want '%v'", err, nil)
	}

	if o == nil {
		t.Fatalf("NewOnDemand.ondemandflow = '%v' | want not '%v'", o, nil)
	}
}

func TestNewOnDemandFailure(t *testing.T) {
	o, err := NewOnDemand(dummyId, nil, dummyOtr, dummyPunch, dummyGuard)

	if err == nil {
		t.Fatalf("NewOnDemand.err = '%v' | want not '%v'", err, nil)
	}

	if o != nil {
		t.Fatalf("NewOnDemand.ondemandflow = '%v' | want '%v'", o, nil)
	}
}

func TestOnDemandFlowRunIsNil(t *testing.T) {
	o, _ := NewOnDemand(dummyId, dummySignal, dummyOtr, dummyPunch, dummyGuard)

	run := o.Run()
	if run != nil {
		t.Fatalf("OnDemandFlow.Run() = '%v', want '%v'", run, nil)
	}
}

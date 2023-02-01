package flow_test

// import (
// 	"testing"
// 	"time"

// 	v1 "github.com/punchguard/v1"
// )

// type testPunchOnDemand struct{}

// func TestPunchOnDemandInitSuccess(t *testing.T) {
// 	p, err := v1.NewFlow[v1.PunchOnDemand](DummySignalingTest, DummyOtrTest, DummyPunchTest, DummyGuardTest)

// 	if err != nil {
// 		t.Fatalf("%T.NewFlow.err = '%v' | want '%v'", p, err, nil)
// 	}

// 	if p == nil {
// 		t.Fatalf("%T.NewFlow = '%v' | want not '%v'", p, err, nil)
// 	}
// }

// func TestPunchOnDemandStartSuccess(t *testing.T) {
// 	o, _ := v1.NewFlow[v1.PunchOnDemand](DummySignalingTest, DummyOtrTest, DummyPunchTest, DummyGuardTest)

// 	control := o.Start()

// 	if control != nil {
// 		t.Fatalf("%T.Start() = '%v' | want not '%v'", o, control, nil)
// 	}
// 	<-time.After(3 * time.Second)
// }

package v1_test

import (
	"testing"

	v1 "github.com/punchguard/v1"
)

type channelTypes interface {
	string | struct{} | v1.Peers
}

func ReceiveIsOpen[T channelTypes](ch <-chan T) (<-chan T, bool) {
	open := ch != nil
	select {
	case _, open = <-ch:
	default:
	}
	return ch, open
}

func SendIsOpen[T channelTypes](ch chan<- T) (chan<- T, bool) {
	defer func() {
		recover()
	}()

	open := ch != nil

	select {
	case ch <- *new(T):
	default:
	}

	return ch, open
}

func TestReceiveIsOpenTrue(t *testing.T) {
	ch := make(<-chan string)

	if receive, open := ReceiveIsOpen(ch); receive == nil || !open {
		t.Fatalf("ReceiveIsOpen() = '%v', open = '%v' | want '%s', open = '%v'", receive, open, "not <nil>", true)
	}
}

func TestReceiveIsOpenFalse(t *testing.T) {
	ch := make(chan string)
	close(ch)

	if receive, open := ReceiveIsOpen(ch); receive == nil || open {
		t.Fatalf("ReceiveIsOpen() = '%v', open = '%v' | want '%s', open = '%v'", receive, open, "not <nil>", false)
	}
}

func TestReceiveIsOpenNil(t *testing.T) {
	var ch chan string

	if receive, open := ReceiveIsOpen(ch); receive != nil || open {
		t.Fatalf("ReceiveIsOpen() = '%v', open = '%v' | want '%s', open = '%v'", receive, open, "<nil>", false)
	}
}

func TestSendIsOpenTrue(t *testing.T) {
	ch := make(chan string)

	if send, open := SendIsOpen(ch); send == nil || !open {
		t.Fatalf("SendIsOpen() = '%v', open = '%v' | want '%s', open = '%v'", send, open, "not <nil>", true)
	}
}

func TestSendIsOpenFalse(t *testing.T) {
	ch := make(chan string)
	close(ch)

	if send, open := SendIsOpen(ch); send != nil || open {
		t.Fatalf("SendIsOpen() = '%v', open = '%v' | want '%s', open = '%v'", send, open, "not <nil>", false)
	}
}

func TestSendIsOpenNil(t *testing.T) {
	var ch chan string

	if send, open := SendIsOpen(ch); send != nil || open {
		t.Fatalf("ReceiveIsOpen() = '%v', open = '%v' | want '%s', open = '%v'", send, open, "<nil>", false)
	}
}

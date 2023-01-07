package v1

import (
	"errors"
	"testing"
)

type channelTypes interface {
	string | struct{} | Peers
}

func receiveIsOpen[T channelTypes](ch <-chan T) bool {
	open := ch != nil
	select {
	case _, open = <-ch:
	default:
	}
	return open
}

func sendIsOpen[T channelTypes](ch chan<- T) bool {
	defer func() {
		recover()
	}()

	open := ch != nil

	select {
	case ch <- *new(T):
	default:
	}

	return open
}

func errorToString(err error) string {
	if err == nil {
		return "nil"
	}
	return err.Error()
}

func TestReceiveIsOpenTrue(t *testing.T) {
	c := make(<-chan string)

	open := receiveIsOpen(c)

	want := true
	if open != want {
		t.Fatalf("receiveIsOpen() = '%v', want '%v'", open, want)
	}
}

func TestReceiveIsOpenFalse(t *testing.T) {
	c := make(chan string)
	close(c)

	open := receiveIsOpen(c)

	want := false
	if open != want {
		t.Fatalf("receiveIsOpen() = '%v', want '%v'", open, want)
	}
}

func TestReceiveIsOpenNil(t *testing.T) {
	var c chan string

	open := receiveIsOpen(c)

	want := false
	if open != want {
		t.Fatalf("receiveIsOpen() = '%v', want '%v'", open, want)
	}
}

func TestSendIsOpenTrue(t *testing.T) {
	c := make(chan string)

	open := sendIsOpen(c)

	want := true
	if open != want {
		t.Fatalf("sendIsOpen() = '%v', want '%v'", open, want)
	}
}

func TestSendIsOpenFalse(t *testing.T) {
	c := make(chan string)
	close(c)

	open := sendIsOpen(c)

	want := false
	if open != want {
		t.Fatalf("sendIsOpen() = '%v', want '%v'", open, want)
	}
}

func TestSendIsOpenNil(t *testing.T) {
	var c chan string

	open := sendIsOpen(c)

	want := false
	if open != want {
		t.Fatalf("sendIsOpen() = '%v', want '%v'", open, want)
	}
}

func TestErrorToStringIsNil(t *testing.T) {
	s := errorToString(nil)

	want := "nil"
	if s != "nil" {
		t.Fatalf("errorToString() = '%v', want '%v'", s, want)
	}
}

func TestErrorToStringNotNil(t *testing.T) {
	s := errorToString(errors.New("err"))

	want := "err"
	if s != "err" {
		t.Fatalf("errorToString() = '%v', want '%v'", s, want)
	}
}

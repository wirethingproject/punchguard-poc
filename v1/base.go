package v1

import (
	"fmt"
	"strings"
)

type BaseSignal struct {
	id        string
	onReady   chan struct{}
	send      chan string
	onReceive chan string
}

func (b *BaseSignal) isValid(id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("%T.Init id is empty '%v'", b, id)
	}

	return nil
}

func (b *BaseSignal) Init(id string) error {
	if err := b.isValid(id); err != nil {
		return err
	}

	b.id = id
	b.onReady = make(chan struct{})
	b.send = make(chan string)
	b.onReceive = make(chan string)

	return nil
}

func (b *BaseSignal) GetId() string               { return b.id }
func (b *BaseSignal) GetOnReady() <-chan struct{} { return b.onReady }
func (b *BaseSignal) GetSend() chan<- string      { return b.send }
func (b *BaseSignal) GetOnReceive() <-chan string { return b.onReceive }

func (b *BaseSignal) SetIsReady() { b.onReady <- *new(struct{}) }

type BaseOtr struct {
	id              string
	signalSend      chan<- string
	signalOnReceive <-chan string
	punchSend       chan string
	punchOnReceive  chan string
}

func (b *BaseOtr) isValid(id string, signalSend chan<- string, signalOnReceive <-chan string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("%T.Init id is empty '%v'", b, id)
	}

	if signalSend == nil {
		return fmt.Errorf("%T.Init signalSend is '%v'", b, signalSend)
	}

	if signalOnReceive == nil {
		return fmt.Errorf("%T.Init signalOnReceive is '%v'", b, signalOnReceive)
	}

	return nil
}

func (b *BaseOtr) Init(id string, signalSend chan<- string, signalOnReceive <-chan string) error {
	if err := b.isValid(id, signalSend, signalOnReceive); err != nil {
		return err
	}

	b.id = id
	b.signalSend = signalSend
	b.signalOnReceive = signalOnReceive
	b.punchSend = make(chan string)
	b.punchOnReceive = make(chan string)

	return nil
}

func (b *BaseOtr) GetId() string                     { return b.id }
func (b *BaseOtr) GetSignalSend() chan<- string      { return b.signalSend }
func (b *BaseOtr) GetSignalOnReceive() <-chan string { return b.signalOnReceive }
func (b *BaseOtr) GetPunchSend() chan<- string       { return b.punchSend }
func (b *BaseOtr) GetPunchOnReceive() <-chan string  { return b.punchOnReceive }

type BasePunch struct {
	send      chan<- string
	onReceive <-chan string
	onPeers   chan Peers
}

func (b *BasePunch) isValid(send chan<- string, onReceive <-chan string) error {
	if send == nil {
		return fmt.Errorf("%T.Init send is '%v'", b, send)
	}

	if onReceive == nil {
		return fmt.Errorf("%T.Init onReceive is '%v'", b, onReceive)
	}

	return nil
}

func (b *BasePunch) Init(send chan<- string, onReceive <-chan string) error {
	if err := b.isValid(send, onReceive); err != nil {
		return err
	}

	b.send = send
	b.onReceive = onReceive
	b.onPeers = make(chan Peers)

	return nil
}

func (b *BasePunch) SetPeers(peers Peers) { b.onPeers <- peers }

func (b *BasePunch) GetSend() chan<- string      { return b.send }
func (b *BasePunch) GetOnReceive() <-chan string { return b.onReceive }
func (b *BasePunch) GetOnPeers() <-chan Peers    { return b.onPeers }

type BaseGuard struct {
	id             string
	onPeers        <-chan Peers
	onConnected    chan struct{}
	onDisconnected chan struct{}
}

func (b *BaseGuard) isValid(id string, onPeers <-chan Peers) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("%T.Init id is empty '%v'", b, id)
	}

	if onPeers == nil {
		return fmt.Errorf("%T.Init onPeers is '%v'", b, onPeers)
	}

	return nil
}

func (b *BaseGuard) Init(id string, onPeers <-chan Peers) error {
	if err := b.isValid(id, onPeers); err != nil {
		return err
	}

	b.id = id
	b.onPeers = onPeers
	b.onConnected = make(chan struct{})
	b.onDisconnected = make(chan struct{})

	return nil
}

func (b *BaseGuard) SetConnected() {
	b.onConnected <- *new(struct{})
}
func (b *BaseGuard) SetDisconnected() {
	b.onDisconnected <- *new(struct{})
}

func (b *BaseGuard) GetId() string                      { return b.id }
func (b *BaseGuard) GetOnPeers() <-chan Peers           { return b.onPeers }
func (b *BaseGuard) GetOnConnected() <-chan struct{}    { return b.onConnected }
func (b *BaseGuard) GetOnDisconnected() <-chan struct{} { return b.onDisconnected }

type BaseFlow struct {
	signal              Signal
	otr                 Otr
	punch               Punch
	guard               Guard
	signalOnReady       <-chan struct{}
	guardOnConnected    <-chan struct{}
	guardOnDisconnected <-chan struct{}
}

func (b *BaseFlow) GetSignal() Signal                    { return b.signal }
func (b *BaseFlow) GetOtr() Otr                          { return b.otr }
func (b *BaseFlow) GetPunch() Punch                      { return b.punch }
func (b *BaseFlow) GetGuard() Guard                      { return b.guard }
func (b *BaseFlow) GetSignalOnReady() <-chan struct{}    { return b.signalOnReady }
func (b *BaseFlow) GetGuardOnConnected() <-chan struct{} { return b.guardOnConnected }
func (b *BaseFlow) GetGuardOnDisconnected() <-chan struct{} {
	return b.guardOnDisconnected
}

func (b *BaseFlow) isValid(signal Signal, otr Otr, punch Punch,
	guard Guard) error {
	if signal == nil {
		return fmt.Errorf("%T.Init signal is '%v'", b, signal)
	}

	if ready := signal.GetOnReady(); ready == nil {
		return fmt.Errorf("%T.Init signal.GetOnReady() is '%v'", b, ready)
	}

	if otr == nil {
		return fmt.Errorf("%T.Init otr is '%v'", b, otr)
	}

	if punch == nil {
		return fmt.Errorf("%T.Init punch is '%v'", b, punch)
	}

	if guard == nil {
		return fmt.Errorf("%T.Init guard is '%v'", b, guard)
	}

	if connected := guard.GetOnConnected(); connected == nil {
		return fmt.Errorf("%T.Init guard.GetOnConnected() is '%v'", b, connected)
	}

	if disconnected := guard.GetOnDisconnected(); disconnected == nil {
		return fmt.Errorf("%T.Init guard.GetOnDisconnected() is '%v'", b, disconnected)
	}

	return nil
}

func (b *BaseFlow) Init(signal Signal, otr Otr, punch Punch,
	guard Guard) error {
	if err := b.isValid(signal, otr, punch, guard); err != nil {
		return err
	}

	b.signal = signal
	b.otr = otr
	b.punch = punch
	b.guard = guard
	b.signalOnReady = signal.GetOnReady()
	b.guardOnConnected = guard.GetOnConnected()
	b.guardOnDisconnected = guard.GetOnDisconnected()

	return nil
}

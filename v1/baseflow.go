package v1

import (
	"fmt"
)

type BaseFlow struct {
	id                  Id
	signal              Signal
	otr                 Otr
	punch               Punch
	guard               Guard
	signalOnReady       <-chan struct{}
	guardOnConnected    <-chan struct{}
	guardOnDisconnected <-chan struct{}
}

func (b *BaseFlow) setId(id Id)                                     { b.id = id }
func (b *BaseFlow) setSignal(signal Signal)                         { b.signal = signal }
func (b *BaseFlow) setOtr(otr Otr)                                  { b.otr = otr }
func (b *BaseFlow) setPunch(punch Punch)                            { b.punch = punch }
func (b *BaseFlow) setGuard(guard Guard)                            { b.guard = guard }
func (b *BaseFlow) setSignalOnReady(onReady <-chan struct{})        { b.signalOnReady = onReady }
func (b *BaseFlow) setGuardOnConnected(onConnected <-chan struct{}) { b.guardOnConnected = onConnected }
func (b *BaseFlow) setGuardOnDisconnected(onDisconnected <-chan struct{}) {
	b.guardOnDisconnected = onDisconnected
}

func (b *BaseFlow) Init(id Id, signal Signal, otr Otr, punch Punch,
	guard Guard) error {

	invalid := new(invalidParams)

	b.setId(id)

	setValueOrInvalid(invalid, "signal", signal == nil, signal, b.setSignal)
	setValueOrInvalid(invalid, "otr", otr == nil, otr, b.setOtr)
	setValueOrInvalid(invalid, "punch", punch == nil, punch, b.setPunch)
	setValueOrInvalid(invalid, "guard", guard == nil, guard, b.setGuard)

	if len(invalid.get()) != 0 {
		return fmt.Errorf("%T.init: nil or empty parameters %v", b, invalid.get())
	}

	setStringOrInvalid(invalid, "id.Signal", b.id.Signal, b.signal.SetId)
	setStringOrInvalid(invalid, "id.Otr", b.id.Otr, b.otr.SetId)
	setStringOrInvalid(invalid, "id.Guard", b.id.Guard, b.guard.SetId)

	getAndSetOrInvalid(invalid, "signal.GetSend", b.signal.GetSend, b.otr.SetSend)
	getAndSetOrInvalid(invalid, "signal.GetOnReceive", b.signal.GetOnReceive, b.otr.SetOnReceive)

	getAndSetOrInvalid(invalid, "otr.GetSend", b.otr.GetSend, b.punch.SetSend)
	getAndSetOrInvalid(invalid, "otr.GetOnReceive", b.otr.GetOnReceive, b.punch.SetOnReceive)

	getAndSetOrInvalid(invalid, "punch.GetOnPeers", b.punch.GetOnPeers, b.guard.SetOnPeers)

	getAndSetOrInvalid(invalid, "signal.GetOnReady", b.signal.GetOnReady, b.setSignalOnReady)
	getAndSetOrInvalid(invalid, "guard.GetOnConnected", b.guard.GetOnConnected, b.setGuardOnConnected)
	getAndSetOrInvalid(invalid, "guard.GetOnDisconnected", b.guard.GetOnDisconnected, b.setGuardOnDisconnected)

	if len(invalid.get()) != 0 {
		return fmt.Errorf("%T.init: nil or empty parameters %v", b, invalid.get())
	}

	return nil
}

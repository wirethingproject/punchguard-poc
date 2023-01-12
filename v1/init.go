package v1

import "fmt"

func asInterface[I any](value any) (I, error) {
	as, isType := value.(I)
	if !isType {
		return as, fmt.Errorf("%T is not %T", value, new(I))
	}
	return as, nil
}

func newType[T any, I any](init func(I) error) (I, error) {
	var nilValue I
	as, err := asInterface[I](new(T))
	if err != nil {
		return nilValue, err
	}
	if err := init(as); err != nil {
		return nilValue, err
	}
	return as, nil
}

func NewSignal[S any](id string) (Signal, error) {
	return newType[S](func(signal Signal) error {
		return signal.Init(id)
	})
}

func NewOtr[S any](id string, signalSend chan<- string, signalOnReceive <-chan string) (Otr, error) {
	return newType[S](func(otr Otr) error {
		return otr.Init(id, signalSend, signalOnReceive)
	})
}

func NewPunch[S any](send chan<- string, onReceive <-chan string) (Punch, error) {
	return newType[S](func(punch Punch) error {
		return punch.Init(send, onReceive)
	})
}

func NewGuard[S any](id string, onPeers <-chan Peers) (Guard, error) {
	return newType[S](func(guard Guard) error {
		return guard.Init(id, onPeers)
	})
}

func NewFlow[F any](signal Signal, otr Otr, punch Punch,
	guard Guard) (Flow, error) {
	return newType[F](func(f Flow) error {
		return f.Init(signal, otr, punch, guard)
	})
}

func Init[S any, O any, P any, G any, F any](id Id) (Flow, error) {
	signal, err := NewSignal[S](id.Signal)
	if err != nil {
		return nil, err
	}

	otr, err := NewOtr[O](id.Otr, signal.GetSend(), signal.GetOnReceive())
	if err != nil {
		return nil, err
	}

	punch, err := NewPunch[P](otr.GetPunchSend(), otr.GetPunchOnReceive())
	if err != nil {
		return nil, err
	}

	guard, err := NewGuard[G](id.Guard, punch.GetOnPeers())
	if err != nil {
		return nil, err
	}

	f, err := NewFlow[F](signal, otr, punch, guard)
	if err != nil {
		return nil, err
	}
	return f, nil
}

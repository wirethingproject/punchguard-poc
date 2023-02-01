package punchguard

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

func NewSignaling[S any](id string) (Signaling, error) {
	return newType[S](func(signaling Signaling) error {
		return signaling.Init(id)
	})
}

func NewOtr[S any](id string) (Otr, error) {
	return newType[S](func(otr Otr) error {
		return otr.Init(id)
	})
}

func NewPunch[S any]() (Punch, error) {
	return newType[S](func(punch Punch) error {
		return punch.Init()
	})
}

func NewGuard[S any](id string) (Guard, error) {
	return newType[S](func(guard Guard) error {
		return guard.Init(id)
	})
}

func NewFlow[F any](signaling Signaling, otr Otr, punch Punch,
	guard Guard) (Flow, error) {
	return newType[F](func(f Flow) error {
		return f.Init(signaling, otr, punch, guard)
	})
}

func Init[S any, O any, P any, G any, F any](id Id) (Flow, error) {
	signaling, err := NewSignaling[S](id.Signaling)
	if err != nil {
		return nil, err
	}

	otr, err := NewOtr[O](id.Otr)
	if err != nil {
		return nil, err
	}

	punch, err := NewPunch[P]()
	if err != nil {
		return nil, err
	}

	guard, err := NewGuard[G](id.Guard)
	if err != nil {
		return nil, err
	}

	f, err := NewFlow[F](signaling, otr, punch, guard)
	if err != nil {
		return nil, err
	}
	return f, nil
}

package v1

import "strings"

func NewFlowParams[S any, O any, P any, G any]() (*S, *O, *P, *G) {
	return new(S), new(O), new(P), new(G)
}

func InitFlowParams(signal Signal, otr Otr, punch Punch,
	guard Guard) error {
	deps := []withInit{signal, otr, punch, guard}
	for _, dep := range deps {
		if err := dep.Init(); err != nil {
			return err
		}
	}
	return nil
}

type channelParams interface {
	<-chan string | chan<- string | <-chan struct{} | <-chan Peers
}

type invalidParams struct {
	params []string
}

func (n *invalidParams) append(name string) {
	n.params = append(n.params, name)
}

func (n *invalidParams) get() []string {
	return n.params
}

func getAndSetOrInvalid[T channelParams](invalid *invalidParams, getterName string, getter func() T, setter func(T)) {
	value := getter()
	if value != nil {
		setter(value)
	} else {
		invalid.append(getterName)
	}
}

func setStringOrInvalid(invalid *invalidParams, valueName string, value string, setter func(string)) {
	if strings.TrimSpace(value) != "" {
		setter(value)
	} else {
		invalid.append(valueName)
	}
}

func setValueOrInvalid[T interface{}](invalid *invalidParams, valueName string, isNil bool, value T, setter func(T)) {
	if !isNil {
		setter(value)
	} else {
		invalid.append(valueName)
	}
}

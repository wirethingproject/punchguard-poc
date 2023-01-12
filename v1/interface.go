package v1

type Controlling interface {
	Stop(controlling any)
}

type Controlled interface {
	GetStopping() <-chan struct{}
	Close()
}

type Signal interface {
	Init(id string) error
	GetId() string
	GetOnReady() <-chan struct{}
	GetSend() chan<- string
	GetOnReceive() <-chan string
	Start() Controlling
}

type Otr interface {
	Init(id string, signalSend chan<- string, signalOnReceive <-chan string) error
	GetId() string
	GetSignalSend() chan<- string
	GetSignalOnReceive() <-chan string
	GetPunchSend() chan<- string
	GetPunchOnReceive() <-chan string
}

type Punch interface {
	Init(send chan<- string, onReceive <-chan string) error
	GetSend() chan<- string
	GetOnReceive() <-chan string
	GetOnPeers() <-chan Peers
	RunOnce() error
}

type Guard interface {
	Init(id string, onPeers <-chan Peers) error
	GetId() string
	GetOnPeers() <-chan Peers
	GetOnConnected() <-chan struct{}
	GetOnDisconnected() <-chan struct{}
	Start() Controlling
}

type Flow interface {
	Init(signal Signal, otr Otr, punch Punch,
		guard Guard) error
	GetSignal() Signal
	GetOtr() Otr
	GetPunch() Punch
	GetGuard() Guard
	Start() Controlling
}

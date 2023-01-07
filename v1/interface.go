package v1

type withId interface {
	SetId(string)
}

type withInit interface {
	Init() error
}

type withRunOnce interface {
	RunOnce() error
}

type withRun interface {
	Run() *Control
}

type Signal interface {
	withInit
	withId
	withRun
	GetOnReady() <-chan struct{}
	GetSend() chan<- string
	GetOnReceive() <-chan string
}

type Otr interface {
	withInit
	withId
	SetSend(chan<- string)
	SetOnReceive(<-chan string)
	GetSend() chan<- string
	GetOnReceive() <-chan string
}

type Punch interface {
	withInit
	withRunOnce
	SetSend(chan<- string)
	SetOnReceive(<-chan string)
	GetOnPeers() <-chan Peers
}

type Guard interface {
	withInit
	withId
	withRun
	SetOnPeers(<-chan Peers)
	GetOnConnected() <-chan struct{}
	GetOnDisconnected() <-chan struct{}
}

type Flow interface {
	Init(id Id, signal Signal, otr Otr, punch Punch,
		guard Guard) error
	withRun
}

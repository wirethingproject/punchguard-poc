package v1

type BaseSignal struct {
	id        string
	onReady   chan struct{}
	send      chan string
	onReceive chan string
}

func (b *BaseSignal) Init() error {
	b.id = ""
	b.onReady = make(chan struct{})
	b.send = make(chan string)
	b.onReceive = make(chan string)
	return nil
}

func (b *BaseSignal) SetId(id string)             { b.id = id }
func (b *BaseSignal) GetOnReady() <-chan struct{} { return b.onReady }
func (b *BaseSignal) GetSend() chan<- string      { return b.send }
func (b *BaseSignal) GetOnReceive() <-chan string { return b.onReceive }

type BaseOtr struct {
	id              string
	signalSend      chan<- string
	signalOnReceive <-chan string
	punchSend       chan string
	punchOnReceive  chan string
}

func (b *BaseOtr) Init() error {
	b.id = ""
	b.signalSend = nil
	b.signalOnReceive = nil
	b.punchSend = make(chan string)
	b.punchOnReceive = make(chan string)
	return nil
}

func (b *BaseOtr) SetId(id string)                      { b.id = id }
func (b *BaseOtr) SetSend(send chan<- string)           { b.signalSend = send }
func (b *BaseOtr) SetOnReceive(onReceive <-chan string) { b.signalOnReceive = onReceive }
func (b *BaseOtr) GetSend() chan<- string               { return b.punchSend }
func (b *BaseOtr) GetOnReceive() <-chan string          { return b.punchOnReceive }

type BasePunch struct {
	send      chan<- string
	onReceive <-chan string
	onPeers   chan Peers
}

func (b *BasePunch) Init() error {
	b.send = nil
	b.onReceive = nil
	b.onPeers = make(chan Peers)
	return nil
}

func (b *BasePunch) SetSend(send chan<- string)           { b.send = send }
func (b *BasePunch) SetOnReceive(onReceive <-chan string) { b.onReceive = onReceive }
func (b *BasePunch) GetOnPeers() <-chan Peers             { return b.onPeers }

type BaseGuard struct {
	id             string
	onPeers        <-chan Peers
	onConnected    chan struct{}
	onDisconnected chan struct{}
}

func (b *BaseGuard) Init() error {
	b.id = ""
	b.onPeers = nil
	b.onConnected = make(chan struct{})
	b.onDisconnected = make(chan struct{})
	return nil
}

func (b *BaseGuard) SetId(id string)                    { b.id = id }
func (b *BaseGuard) SetOnPeers(onPeers <-chan Peers)    { b.onPeers = onPeers }
func (b *BaseGuard) GetOnConnected() <-chan struct{}    { return b.onConnected }
func (b *BaseGuard) GetOnDisconnected() <-chan struct{} { return b.onDisconnected }

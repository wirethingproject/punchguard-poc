package punchguard

type Service interface {
	Start() StoppedEvent
	Stop()
}

type Signaling interface {
	Service
	Init(id string) error
	Connect()
	Disconnect()
	OnReadyEvent() ReceiveEvent
	Send(msg string)
	OnReceiveEvent() ReceiveString
}

type Otr interface {
	// Service
	Init(id string) error
	Close()
	Query() ([]string, error)
	End() ([]string, error)
	OnReadyEvent() ReceiveEvent
	Encode(msg string) ([]string, error)
	Decode(msg string) (string, []string, error)
}

type Punch interface {
	Service
	Init() error
	Punch()
	OnPeersEvent() ReceivePeers
	OnSendEvent() ReceiveString
	Receive(msg string)
}

type Guard interface {
	Service
	Init(id string) error
	Peers(peers Peers)
	OnConnectedEvent() ReceiveEvent
	OnDisconnectedEvent() ReceiveEvent
}

type Flow interface {
	Service
	Init(signaling Signaling, otr Otr, punch Punch,
		guard Guard) error
}

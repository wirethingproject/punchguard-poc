package punchguard

type Id struct {
	Signaling string
	Otr       string
	Guard     string
}

type Peer struct {
	Address string
	Port    int
}

type Peers struct {
	Local  Peer
	Remote Peer
}

type Event struct{}
type ReceiveEvent <-chan Event
type StoppedEvent ReceiveEvent
type ReceiveString <-chan string
type ReceivePeers <-chan Peers

func NewId(signaling, otr, guard string) Id {
	return Id{
		Signaling: signaling,
		Otr:       otr,
		Guard:     guard,
	}
}

func NewPeers(localAddress string, localPort int, remoteAddress string, remotePort int) Peers {
	return Peers{
		Local: Peer{
			Address: localAddress,
			Port:    localPort,
		},
		Remote: Peer{
			Address: remoteAddress,
			Port:    remotePort,
		},
	}
}

func NewEvent() Event {
	return Event{}
}

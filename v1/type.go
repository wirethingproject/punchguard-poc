package v1

type Id struct {
	Signal string
	Otr    string
	Guard  string
}

type Peer struct {
	Address string
	Port    int
}

type Peers struct {
	Local  Peer
	Remote Peer
}

func NewId(signal, otr, guard string) Id {
	return Id{
		Signal: signal,
		Otr:    otr,
		Guard:  guard,
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

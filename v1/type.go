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

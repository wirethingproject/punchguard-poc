package punch

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pionice "github.com/pion/ice/v2"
	"github.com/pion/randutil"

	"github.com/punchguard/v0"
)

type IceLite struct {
	punchguard.BasePunch
	remoteAuthChannel      chan string
	remoteCandidateChannel chan pionice.Candidate
	config                 IceConfig
	Id                     string
}

func (i *IceLite) Init() error {
	// log.Printf("%T.Init", i)

	if err := i.InitBase(); err != nil {
		return err
	}

	i.remoteAuthChannel = make(chan string, 100)
	i.remoteCandidateChannel = make(chan pionice.Candidate, 100)
	i.config = i.IceInit()

	return nil
}

type Event struct {
	Action string
	Data   string
}

type AuthEvent struct {
	Ufrag string
	Pwd   string
}

type CandidateEvent struct {
	Candidate string
}

// type IceOnRemoteCandidate func(string)
// type IceOnRemoteAuth func(string, string)
// type IceOnError func(error)

// type IceTransport interface {
// 	Open() error
// 	Close() error
// 	SendRemoteAuth(string, string) error
// 	SendRemoteCandidate(string) error
// 	OnRemoteCandidate(IceOnRemoteCandidate)
// 	OnRemoteAuth(IceOnRemoteAuth)
// 	OnError(IceOnError)
// }

type IceConfig struct {
	StunHost string
	StunPort int
}

// type IceCredential struct {
// 	ufrag string
// 	pwd   string
// }

func (i *IceLite) IceInit() IceConfig {
	return IceConfig{
		StunHost: "stun.stunprotocol.org",
		StunPort: 3478,
	}
}

func (i *IceLite) Ice() (*punchguard.Peers, error) { //nolint

	var (
		err  error
		conn *pionice.Conn
	)

	running := true

	i.remoteAuthChannel = make(chan string, 100)
	i.remoteCandidateChannel = make(chan pionice.Candidate, 100)

	stunServerURL := &pionice.URL{
		Scheme: pionice.SchemeTypeSTUN,
		Host:   i.config.StunHost,
		Port:   i.config.StunPort,
		Proto:  pionice.ProtoTypeUDP,
	}

	iceAgent, err := pionice.NewAgent(&pionice.AgentConfig{
		Urls:         []*pionice.URL{stunServerURL},
		NetworkTypes: []pionice.NetworkType{pionice.NetworkTypeUDP4},
		// Lite: true,
	})
	if err != nil {
		return nil, err
	}

	i.Async.WhenRunningAsync(func() {
		log.Printf("%T.Ice: AddRemoteCandidate started", i)
		for running {
			select {
			case c, open := <-i.remoteCandidateChannel:
				if open {
					log.Printf("%T.Ice: AddRemoteCandidate '%v'", i, c)
					if err := iceAgent.AddRemoteCandidate(c); err != nil {
						panic(err)
					}
				}
			default:
			}
			<-time.After(100 * time.Millisecond)
		}
		log.Printf("%T.Ice: AddRemoteCandidate stopped", i)
	})

	// When we have gathered a new ICE Candidate send it to the remote peer
	if err = iceAgent.OnCandidate(func(c pionice.Candidate) {
		log.Printf("%T.Ice: OnCandidate '%v'", i, c)
		if c == nil {
			return
		}

		err := i.SendRemoteCandidate(c.Marshal())
		if err != nil {
			panic(err)
		}
	}); err != nil {
		return nil, err
	}

	// When ICE Connection state has change print to stdout
	if err = iceAgent.OnConnectionStateChange(func(c pionice.ConnectionState) {
		log.Printf("%T.OnConnectionStateChange '%v'", i, c.String())
		// if c.String() == "Failed" {
		// 	running = false
		// 	iceAgent.Close()
		// }
	}); err != nil {
		return nil, err
	}

	// Get the local auth details and send to remote peer
	localUfrag, localPwd, err := iceAgent.GetLocalUserCredentials()
	if err != nil {
		return nil, err
	}

	log.Printf("%T.SendRemoteAuth '", i)
	err = i.SendRemoteAuth(localUfrag, localPwd)
	if err != nil {
		return nil, err
	}

	remoteUfrag := <-i.remoteAuthChannel
	remotePwd := <-i.remoteAuthChannel

	log.Printf("%T.ReceiveRemoteAuth '", i)

	if err = iceAgent.GatherCandidates(); err != nil {
		return nil, err
	}

	// Start the ICE Agent. One side must be controlled, and the other must be controlling
	if i.Id == "p1" {
		log.Printf("%T.Dial '", i)
		conn, err = iceAgent.Dial(context.TODO(), remoteUfrag, remotePwd)
		log.Printf("%T.Dialed '", i)
	} else {
		log.Printf("%T.Accept '", i)
		conn, err = iceAgent.Accept(context.TODO(), remoteUfrag, remotePwd)
		log.Printf("%T.Accepted '", i)
	}
	if err != nil {
		return nil, err
	}

	log.Printf("%T.Ice.conn: local '%v' remote '%v'", i, conn.LocalAddr(), conn.RemoteAddr())

	// fmt.Printf("conn: %s\n", conn)

	p, err := iceAgent.GetSelectedCandidatePair()
	if err != nil {
		return nil, err
	}
	// fmt.Printf("pair: %s\n", p)

	var pair punchguard.Peers

	related := p.Local.RelatedAddress()
	if related != nil {
		pair = punchguard.NewPeers(p.Local.RelatedAddress().Address, p.Local.RelatedAddress().Port, p.Remote.Address(), p.Remote.Port())
	} else {
		pair = punchguard.NewPeers(p.Local.Address(), p.Local.Port(), p.Remote.Address(), p.Remote.Port())
	}

	val, err := randutil.GenerateCryptoRandomString(15, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	if err != nil {
		return nil, err
	}
	if _, err = conn.Write([]byte(val)); err != nil {
		return nil, err
	}

	buf := make([]byte, 1500)
	_, err = conn.Read(buf)
	if err != nil {
		return nil, err
	}

	err = conn.Close()
	if err != nil {
		return nil, err
	}

	running = false
	return &pair, nil
}

func (i *IceLite) SendRemoteAuth(ufrag, pwd string) error {
	// fmt.Printf("tox.SendRemoteAuth\n")

	data, err := json.Marshal(AuthEvent{ufrag, pwd})
	if err != nil {
		return err
	}

	message, err := json.Marshal(Event{"auth", string(data)})
	if err != nil {
		return err
	}

	i.Send(string(message))

	return nil
}

func (i *IceLite) SendRemoteCandidate(candidate string) error {
	// fmt.Printf("tox.SendRemoteCandidate %s\n", candidate)

	data, err := json.Marshal(CandidateEvent{candidate})
	if err != nil {
		return err
	}

	message, err := json.Marshal(Event{"candidate", string(data)})
	if err != nil {
		return err
	}

	i.Send(string(message))

	return nil
}

func (i *IceLite) OnRemoteCandidate(e CandidateEvent) {
	// fmt.Printf("HTTP OnRemoteCandidate: %s\n", e.Candidate)

	c, err := pionice.UnmarshalCandidate(e.Candidate)
	if err != nil {
		panic(err)
	}

	if c.Address() == "192.168.0.11" {
		return
	}

	// fmt.Printf("ICE remote Candidate: %s\n", c.String())

	i.remoteCandidateChannel <- c
}

func (i *IceLite) OnRemoteAuth(e AuthEvent) {
	// fmt.Printf("HTTP OnRemoteAuth: %s %s\n", e.Ufrag, e.Pwd)

	i.remoteAuthChannel <- e.Ufrag
	i.remoteAuthChannel <- e.Pwd
}

func (i *IceLite) onCallbackFriendMessage(message string) {
	var e Event
	err := json.Unmarshal([]byte(message), &e)
	// fmt.Printf("tox friend message event: %s %s\n", message, e)
	switch e.Action {
	case "auth":
		var ae AuthEvent
		_ = json.Unmarshal([]byte(e.Data), &ae)
		// fmt.Printf("tox friend message auth event: %s %s\n", ae, err)
		i.OnRemoteAuth(ae)
	case "candidate":
		var ce CandidateEvent
		_ = json.Unmarshal([]byte(e.Data), &ce)
		// fmt.Printf("tox friend message candidate event: %s %s\n", ce, err)
		i.OnRemoteCandidate(ce)
	default:
		log.Printf("%T.onCallbackFriendMessage error: '%v' '%v' '%v'", i, message, err, e)
	}
}

func (i *IceLite) Start() punchguard.StoppedEvent {
	return i.MainLoop(func() {
		log.Printf("%T.MainLoop: started", i)
	}, func() {
		select {
		case <-i.OnPunchEvent():
			log.Printf("%T.OnPunchEvent", i)
			i.WhenRunningAsync(func() {
				if peers, err := i.Ice(); err != nil {
					log.Printf("%T.Ice: error '%v'", i, err)
				} else {
					i.Peers(*peers)
				}
			})
		case msg := <-i.OnReceiveEvent():
			i.WhenRunningAsync(func() {
				i.onCallbackFriendMessage(msg)
			})
		default:
		}
	}, func() {
		log.Printf("%T.MainLoop: stopped", i)
		i.Close()
	})
}

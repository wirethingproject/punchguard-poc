package signaling

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/coyim/otr3"
	pionice "github.com/pion/ice/v2"
	"github.com/pion/randutil"

	"github.com/punchguard/v0"
	"github.com/punchguard/v0/unstable/otr"
	"github.com/punchguard/v0/unstable/punch"
)

type LoopOtr3Ice struct {
	punchguard.BaseSignaling
	Conversation           *otr3.Conversation
	remoteAuthChannel      chan string
	remoteCandidateChannel chan pionice.Candidate
	config                 punch.IceConfig
}

func (t *LoopOtr3Ice) Init(id string) error {
	// log.Printf("%T.Init", t)

	if err := t.InitBase(id); err != nil {
		return err
	}

	t.Conversation = otr.NewConversation()
	t.remoteAuthChannel = make(chan string, 3)
	t.remoteCandidateChannel = make(chan pionice.Candidate)
	t.config = t.IceInit()

	return nil
}

func (i *LoopOtr3Ice) IceInit() punch.IceConfig {
	return punch.IceConfig{
		StunHost: "stun.stunprotocol.org",
		StunPort: 3478,
	}
}

func (i *LoopOtr3Ice) Ice() (*punchguard.Peers, error) { //nolint

	var (
		err  error
		conn *pionice.Conn
	)

	running := true

	i.remoteCandidateChannel = make(chan pionice.Candidate)

	// if i.config.IsControlling {
	// 	fmt.Println("Local Agent is controlling")
	// } else {
	// 	fmt.Println("Local Agent is controlled")
	// }

	stunServerURL := &pionice.URL{
		Scheme: pionice.SchemeTypeSTUN,
		Host:   i.config.StunHost,
		Port:   i.config.StunPort,
		Proto:  pionice.ProtoTypeUDP,
	}

	iceAgent, err := pionice.NewAgent(&pionice.AgentConfig{
		Urls:         []*pionice.URL{stunServerURL},
		NetworkTypes: []pionice.NetworkType{pionice.NetworkTypeUDP4},
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
				<-time.After(100 * time.Millisecond)
			}
		}
		log.Printf("%T.Ice: AddRemoteCandidate stopped", i)
	})

	// When we have gathered a new ICE Candidate send it to the remote peer
	if err = iceAgent.OnCandidate(func(c pionice.Candidate) {
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
		if c.String() == "Failed" {
			iceAgent.Close()
		}
	}); err != nil {
		return nil, err
	}

	// Get the local auth details and send to remote peer
	localUfrag, localPwd, err := iceAgent.GetLocalUserCredentials()
	if err != nil {
		return nil, err
	}

	err = i.SendRemoteAuth(localUfrag, localPwd)
	if err != nil {
		return nil, err
	}

	remoteUfrag := <-i.remoteAuthChannel
	remotePwd := <-i.remoteAuthChannel

	if err = iceAgent.GatherCandidates(); err != nil {
		return nil, err
	}

	// Start the ICE Agent. One side must be controlled, and the other must be controlling
	// if i.config.IsControlling {
	// 	conn, err = iceAgent.Dial(context.TODO(), remoteUfrag, remotePwd)
	// } else {
	conn, err = iceAgent.Accept(context.TODO(), remoteUfrag, remotePwd)
	// }
	if err != nil {
		return nil, err
	}

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

func (i *LoopOtr3Ice) SendRemoteAuth(ufrag, pwd string) error {
	// fmt.Printf("tox.SendRemoteAuth\n")

	data, err := json.Marshal(punch.AuthEvent{Ufrag: ufrag, Pwd: pwd})
	if err != nil {
		return err
	}

	message, err := json.Marshal(punch.Event{Action: "auth", Data: string(data)})
	if err != nil {
		return err
	}

	i.Return(string(message))

	return nil
}

func (i *LoopOtr3Ice) SendRemoteCandidate(candidate string) error {
	// fmt.Printf("tox.SendRemoteCandidate %s\n", candidate)

	data, err := json.Marshal(punch.CandidateEvent{Candidate: candidate})
	if err != nil {
		return err
	}

	message, err := json.Marshal(punch.Event{Action: "candidate", Data: string(data)})
	if err != nil {
		return err
	}

	i.Return(string(message))

	return nil
}

func (i *LoopOtr3Ice) OnRemoteCandidate(e punch.CandidateEvent) {
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

func (i *LoopOtr3Ice) OnRemoteAuth(e punch.AuthEvent) {
	// fmt.Printf("HTTP OnRemoteAuth: %s %s\n", e.Ufrag, e.Pwd)

	i.remoteAuthChannel <- e.Ufrag
	i.remoteAuthChannel <- e.Pwd
}

func (i *LoopOtr3Ice) onCallbackFriendMessage(message string) {
	var e punch.Event
	err := json.Unmarshal([]byte(message), &e)
	// fmt.Printf("tox friend message event: %s %s\n", message, e)
	switch e.Action {
	case "auth":
		var ae punch.AuthEvent
		_ = json.Unmarshal([]byte(e.Data), &ae)
		// fmt.Printf("tox friend message auth event: %s %s\n", ae, err)
		i.OnRemoteAuth(ae)
	case "candidate":
		var ce punch.CandidateEvent
		_ = json.Unmarshal([]byte(e.Data), &ce)
		// fmt.Printf("tox friend message candidate event: %s %s\n", ce, err)
		i.OnRemoteCandidate(ce)
	default:
		log.Printf("%T.onCallbackFriendMessage error: '%v' '%v' '%v'", i, message, err, e)
	}
}

func (m *LoopOtr3Ice) Return(msg string) {
	toSend, err := m.Conversation.Send([]byte(msg))

	if err != nil {
		log.Printf("%T.Conversation.Send: err '%v' ", m, err)
	}

	for _, s := range toSend {
		m.Receive(string(s))
	}
}

func (m *LoopOtr3Ice) Start() punchguard.StoppedEvent {
	return m.MainLoop(func() {
		log.Printf("%T.MainLoop: started", m)
	}, func() {
		select {
		case <-m.OnConnectEvent():
			log.Printf("%T.OnConnectEvent", m)
			// TODO sync.RunOnce
			m.WhenRunningAsync(func() {
				m.Ready()
				if peers, err := m.Ice(); err != nil {
					log.Printf("%T.Ice: error '%v'", m, err)
				} else {
					log.Printf("%T.Ice: peers '%v'", m, *peers)
				}
			})
		case <-m.OnDisconnectEvent():
			log.Printf("%T.OnDisconnectEvent", m)
			m.IsReady = false
		case msg := <-m.OnSendEvent():
			m.WhenRunningAsync(func() {
				if !m.IsReady {
					log.Printf("%T.OnSendEvent: err '%v' ", m, "IsReady = false")
					return
				}

				plain, toSend, err := m.Conversation.Receive([]byte(msg))

				if err != nil {
					log.Printf("%T.Conversation.Receive: err '%v' ", m, err)
				}

				for _, s := range toSend {
					m.Receive(string(s))
				}

				if string(plain) != "" {
					m.onCallbackFriendMessage(string(plain))
				}
			})
		default:
		}
	}, func() {
		log.Printf("%T.MainLoop: stopped", m)
		m.Close()
	})
}

package signaling

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/coyim/otr3"
	"github.com/punchguard/v0"
	"github.com/punchguard/v0/unstable/otr"
)

type LoopOtr3 struct {
	punchguard.BaseSignaling
	Conversation *otr3.Conversation
}

func (t *LoopOtr3) Init(id string) error {
	// log.Printf("%T.Init", t)

	if err := t.InitBase(id); err != nil {
		return err
	}

	t.Conversation = otr.NewConversation()

	return nil
}

func (m *LoopOtr3) Start() punchguard.StoppedEvent {
	return m.MainLoop(func() {
		log.Printf("%T.MainLoop: started", m)
	}, func() {
		select {
		case <-m.OnConnectEvent():
			log.Printf("%T.OnConnectEvent", m)
			// TODO sync.RunOnce
			m.WhenRunningAsync(func() {
				<-time.After(2 * time.Second)
				m.Ready()
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
					i, _ := strconv.Atoi(string(plain))
					next := fmt.Sprintf("%v", i+1)

					// log.Printf("%T: received '%v' sended '%v'", m, string(plain), string(next))

					toSend, err := m.Conversation.Send([]byte(next))

					if err != nil {
						log.Printf("%T.Conversation.Send: err '%v' ", m, err)
					}

					for _, s := range toSend {
						m.Receive(string(s))
					}
				}
			})
		default:
		}
	}, func() {
		log.Printf("%T.MainLoop: stopped", m)
		m.Close()
	})
}

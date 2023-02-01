package flow

import (
	"log"
	"time"

	"github.com/punchguard/v0"
)

type PunchOnDemand struct {
	punchguard.BaseFlow
}

func (p *PunchOnDemand) Init(signaling punchguard.Signaling, otr punchguard.Otr, punch punchguard.Punch,
	guard punchguard.Guard) error {
	// log.Printf("%T.Init", p)

	if err := p.InitBase(signaling, otr, punch, guard); err != nil {
		return err
	}

	return nil
}

func (p *PunchOnDemand) Stop() {
	p.StopService(p.GetSignaling())
	// p.StopService(p.otr)
	p.StopService(p.GetPunch())
	p.StopService(p.GetGuard())
	p.Async.StopAsync()
}

func WhenOpen(open bool, sync func()) {
	if open {
		sync()
	}
}

func (p *PunchOnDemand) SendAll(toSend []string) {
	for _, s := range toSend {
		p.GetSignaling().Send(s)
	}
}

func (p *PunchOnDemand) Otr() {
	select {
	case _, open := <-p.GetSignaling().OnReadyEvent():
		WhenOpen(open, func() {
			toSend, err := p.GetOtr().Query()
			if err != nil {
				log.Printf("%T.OtrQuery: error '%v'", p, err)
			} else {
				p.SendAll(toSend)
			}
		})
	case msg, open := <-p.GetPunch().OnSendEvent():
		// log.Printf("%T.OnSendEvent: '%v' open '%v'", p, msg, open)
		WhenOpen(open, func() {
			toSend, err := p.GetOtr().Encode(msg)
			if err != nil {
				log.Printf("%T.Otr: encode error '%v'", p, err)
			} else {
				p.SendAll(toSend)
			}
		})
	case msg, open := <-p.GetSignaling().OnReceiveEvent():
		// log.Printf("%T.OnReceiveEvent: '%v' open '%v'", p, msg, open)
		WhenOpen(open, func() {
			plain, toSend, err := p.GetOtr().Decode(msg)
			if err != nil {
				log.Printf("%T.Otr: decode error '%v'", p, err)
			} else {
				p.SendAll(toSend)
				if plain != "" {
					p.GetPunch().Receive(plain)
				}
			}
		})
	default:
	}
}

func (p *PunchOnDemand) Disconnect() {
	toSend, err := p.GetOtr().End()
	if err != nil {
		log.Printf("%T.Disconnect: otr.End error '%v'", p, err)
	} else {
		p.SendAll(toSend)
	}
	p.GetSignaling().Disconnect()
}

func (p *PunchOnDemand) Main() {
	select {
	case _, open := <-p.GetOtr().OnReadyEvent():
		WhenOpen(open, func() {
			p.WhenRunningAsync(func() {
				<-time.After(5 * time.Second)
				p.GetPunch().Punch()
			})
		})
	case peers, open := <-p.GetPunch().OnPeersEvent():
		WhenOpen(open, func() {
			p.GetGuard().Peers(peers)
		})
	case _, open := <-p.GetGuard().OnConnectedEvent():
		WhenOpen(open, func() {
			p.Disconnect()
		})
	case _, open := <-p.GetGuard().OnDisconnectedEvent():
		WhenOpen(open, func() {
			p.GetSignaling().Connect()
		})
	default:
	}
}

func (p *PunchOnDemand) Start() punchguard.StoppedEvent {
	p.StartService(p.GetGuard())
	p.StartService(p.GetPunch())
	// p.StartService(p.GetOtr())
	p.StartService(p.GetSignaling())

	return p.MainLoop(func() {
		p.GetSignaling().Connect()
		log.Printf("%T.MainLoop: started", p)
	}, func() {
		p.Otr()
		p.Main()
	}, func() {
		log.Printf("%T.MainLoop: stopped", p)
		p.GetOtr().Close()
		p.Close()
	})
}

package punchguard

import (
	"fmt"
	"log"
	"strings"
)

type BaseSignaling struct {
	Async
	id         string
	IsReady    bool
	connect    chan Event
	disconnect chan Event
	ready      chan Event
	send       chan string
	receive    chan string
}

func isValidId(b any, id string) error {
	if strings.TrimSpace(id) == "" {
		return fmt.Errorf("%T.Init id is empty '%v'", b, id)
	}

	return nil
}

func (b *BaseSignaling) InitBase(id string) error {
	// log.Printf("%T.Init", b)

	if err := isValidId(b, id); err != nil {
		return err
	}

	b.id = id
	b.IsReady = false
	b.connect = make(chan Event, 1)
	b.disconnect = make(chan Event, 1)
	b.ready = make(chan Event, 1)
	b.send = make(chan string, 100)
	b.receive = make(chan string, 100)

	b.InitAsync()

	return nil
}

func (b *BaseSignaling) Close() {
	b.IsReady = false
	close(b.connect)
	close(b.disconnect)
	close(b.ready)
	close(b.send)
	close(b.receive)
	log.Printf("%T.Close", b)
}

func (b *BaseSignaling) Stop() {
	log.Printf("%T.Stop", b)
	b.StopAsync()
}

func (b *BaseSignaling) GetId() string { return b.id }

func (b *BaseSignaling) OnConnectEvent() ReceiveEvent { return b.connect }

func (b *BaseSignaling) Connect() {
	b.WhenRunningSync(func() {
		log.Printf("%T.Connect", b)
		select {
		case b.connect <- NewEvent():
		default:
			log.Printf("%T.Connect: channel error", b)
		}
	})
}

func (b *BaseSignaling) OnDisconnectEvent() ReceiveEvent { return b.disconnect }

func (b *BaseSignaling) Disconnect() {
	b.WhenRunningSync(func() {
		log.Printf("%T.Disconnect", b)
		select {
		case b.disconnect <- NewEvent():
		default:
			log.Printf("%T.Disconnect: channel error", b)
		}
	})
}

func (b *BaseSignaling) OnReadyEvent() ReceiveEvent { return b.ready }

func (b *BaseSignaling) Ready() {
	b.WhenRunningSync(func() {
		log.Printf("%T.Ready", b)
		b.IsReady = true
		select {
		case b.ready <- NewEvent():
		default:
			log.Printf("%T.Ready: channel error", b)
		}
	})
}

func (b *BaseSignaling) OnSendEvent() ReceiveString { return b.send }

func (b *BaseSignaling) Send(msg string) {
	b.WhenRunningSync(func() {
		select {
		case b.send <- msg:
		default:
			log.Printf("%T.Send: channel error", b)
		}
	})
}

func (b *BaseSignaling) OnReceiveEvent() ReceiveString { return b.receive }

func (b *BaseSignaling) Receive(msg string) {
	b.WhenRunningSync(func() {
		select {
		case b.receive <- msg:
		default:
			log.Printf("%T.Receive: channel error", b)
		}
	})
}

type BaseOtr struct {
	// Async
	id      string
	IsReady bool
	ready   chan Event
	// signalingSend    chan string
	// signalingReceive chan string
	// punchSend        chan string
	// punchReceive     chan string
}

func (b *BaseOtr) InitBase(id string) error {
	// log.Printf("%T.Init", b)

	if err := isValidId(b, id); err != nil {
		return err
	}

	b.id = id
	b.IsReady = false
	b.ready = make(chan Event, 1)
	// b.signalingSend = make(chan string, 1)
	// b.signalingReceive = make(chan string, 1)
	// b.punchSend = make(chan string, 1)
	// b.punchReceive = make(chan string, 1)

	// b.InitAsync()

	return nil
}

func (b *BaseOtr) Close() {
	b.IsReady = false
	close(b.ready)
	log.Printf("%T.Close", b)
}

// func (b *BaseOtr) Close() {
// 	// close(b.punchSend)
// 	// close(b.punchReceive)
// 	log.Printf("%T.Close", b)
// }

// func (b *BaseOtr) Stop() {
// 	log.Printf("%T.Stop", b)
// 	// b.StopAsync()
// }

func (b *BaseOtr) GetId() string { return b.id }

func (b *BaseOtr) OnReadyEvent() ReceiveEvent { return b.ready }

func (b *BaseOtr) Ready() {
	// b.WhenRunningSync(func() {
	log.Printf("%T.Ready", b)
	b.IsReady = true
	select {
	case b.ready <- NewEvent():
	default:
		log.Printf("%T.Ready: channel error", b)
	}
	// })
}

// func (b *BaseOtr) OnSignalingSendEvent() ReceiveString { return b.signalingSend }

// func (b *BaseOtr) SignalingSend(msg string) {
// 	b.WhenRunningSync(func() {
// 		select {
// 		case b.signalingSend <- msg:
// 		default:
// 			log.Printf("%T.SignalingSend: channel error", b)
// 		}
// 	})
// }

// func (b *BaseOtr) OnSignalingReceiveEvent() ReceiveString { return b.signalingReceive }

// func (b *BaseOtr) SignalingReceive(msg string) {
// 	b.WhenRunningSync(func() {
// 		select {
// 		case b.signalingReceive <- msg:
// 		default:
// 			log.Printf("%T.SignalingReceive: channel error", b)
// 		}
// 	})
// }

// func (b *BaseOtr) OnPunchSendEvent() ReceiveString { return b.punchSend }

// func (b *BaseOtr) PunchSend(msg string) {
// 	b.WhenRunningSync(func() {
// 		select {
// 		case b.punchSend <- msg:
// 		default:
// 			log.Printf("%T.PunchSend: channel error", b)
// 		}
// 	})
// }

// func (b *BaseOtr) OnPunchReceiveEvent() ReceiveString { return b.punchReceive }

// func (b *BaseOtr) PunchReceive(msg string) {
// 	b.WhenRunningSync(func() {
// 		select {
// 		case b.punchReceive <- msg:
// 		default:
// 			log.Printf("%T.PunchReceive: channel error", b)
// 		}
// 	})
// }

type BasePunch struct {
	Async
	punch   chan Event
	peers   chan Peers
	send    chan string
	receive chan string
}

func (b *BasePunch) InitBase() error {
	// log.Printf("%T.Init", b)

	b.punch = make(chan Event, 1)
	b.peers = make(chan Peers, 1)
	b.send = make(chan string, 100)
	b.receive = make(chan string, 100)

	b.InitAsync()

	return nil
}

func (b *BasePunch) Close() {
	close(b.punch)
	close(b.peers)
	close(b.send)
	close(b.receive)
	log.Printf("%T.Close", b)
}

func (b *BasePunch) Stop() {
	log.Printf("%T.Stop", b)
	b.StopAsync()
}

func (b *BasePunch) OnPunchEvent() ReceiveEvent { return b.punch }

func (b *BasePunch) Punch() {
	b.WhenRunningSync(func() {
		log.Printf("%T.Punch", b)
		select {
		case b.punch <- NewEvent():
		default:
			log.Printf("%T.Punch: channel error", b)
		}
	})
}

func (b *BasePunch) OnPeersEvent() ReceivePeers { return b.peers }

func (b *BasePunch) Peers(peers Peers) {
	b.WhenRunningSync(func() {
		log.Printf("%T.Peers: '%v'", b, peers)
		select {
		case b.peers <- peers:
		default:
			log.Printf("%T.Peers: channel error", b)
		}
	})
}

func (b *BasePunch) OnSendEvent() ReceiveString { return b.send }

func (b *BasePunch) Send(msg string) {
	b.WhenRunningSync(func() {
		select {
		case b.send <- msg:
		default:
			log.Printf("%T.Send: channel error", b)
		}
	})
}

func (b *BasePunch) OnReceiveEvent() ReceiveString { return b.receive }

func (b *BasePunch) Receive(msg string) {
	b.WhenRunningSync(func() {
		select {
		case b.receive <- msg:
		default:
			log.Printf("%T.Receive: channel error", b)
		}
	})
}

type BaseGuard struct {
	Async
	id           string
	peers        chan Peers
	connected    chan Event
	disconnected chan Event
}

func (b *BaseGuard) InitBase(id string) error {
	// log.Printf("%T.Init", b)

	if err := isValidId(b, id); err != nil {
		return err
	}

	b.id = id
	b.peers = make(chan Peers, 1)
	b.connected = make(chan Event, 1)
	b.disconnected = make(chan Event, 1)

	b.InitAsync()

	return nil
}

func (b *BaseGuard) Close() {
	close(b.peers)
	close(b.connected)
	close(b.disconnected)
	log.Printf("%T.Close", b)
}

func (b *BaseGuard) Stop() {
	log.Printf("%T.Stop", b)
	b.StopAsync()
}

func (b *BaseGuard) GetId() string { return b.id }

func (b *BaseGuard) OnPeersEvent() ReceivePeers { return b.peers }

func (b *BaseGuard) Peers(peers Peers) {
	b.WhenRunningSync(func() {
		log.Printf("%T.Peers: '%v'", b, peers)
		select {
		case b.peers <- peers:
		default:
			log.Printf("%T.Peers: channel error", b)
		}
	})
}

func (b *BaseGuard) OnConnectedEvent() ReceiveEvent { return b.connected }

func (b *BaseGuard) Connected() {
	b.WhenRunningSync(func() {
		log.Printf("%T.Connected", b)
		select {
		case b.connected <- NewEvent():
		default:
			log.Printf("%T.Connected: channel error", b)
		}
	})
}

func (b *BaseGuard) OnDisconnectedEvent() ReceiveEvent { return b.disconnected }

func (b *BaseGuard) Disconnected() {
	b.WhenRunningSync(func() {
		log.Printf("%T.Disconnected", b)
		select {
		case b.disconnected <- NewEvent():
		default:
			log.Printf("%T.Disconnected: channel error", b)
		}
	})
}

type BaseFlow struct {
	Async
	signaling Signaling
	otr       Otr
	punch     Punch
	guard     Guard
}

func (b *BaseFlow) GetSignaling() Signaling { return b.signaling }
func (b *BaseFlow) GetOtr() Otr             { return b.otr }
func (b *BaseFlow) GetPunch() Punch         { return b.punch }
func (b *BaseFlow) GetGuard() Guard         { return b.guard }

func (b *BaseFlow) isValid(signaling Signaling, otr Otr, punch Punch,
	guard Guard) error {
	if signaling == nil {
		return fmt.Errorf("%T.Init signaling is '%v'", b, signaling)
	}

	if otr == nil {
		return fmt.Errorf("%T.Init otr is '%v'", b, otr)
	}

	if punch == nil {
		return fmt.Errorf("%T.Init punch is '%v'", b, punch)
	}

	if guard == nil {
		return fmt.Errorf("%T.Init guard is '%v'", b, guard)
	}

	return nil
}

func (b *BaseFlow) InitBase(signaling Signaling, otr Otr, punch Punch,
	guard Guard) error {
	// log.Printf("%T.Init", b)
	if err := b.isValid(signaling, otr, punch, guard); err != nil {
		return err
	}

	b.signaling = signaling
	b.otr = otr
	b.punch = punch
	b.guard = guard

	b.InitAsync()

	return nil
}

func (b *BaseFlow) Close() {
	b.signaling = nil
	b.otr = nil
	b.punch = nil
	b.guard = nil
	log.Printf("%T.Close", b)
}

func (b *BaseFlow) Stop() {
	log.Printf("%T.Stop", b)
	b.StopAsync()
}

package punchguard_test

// import (
// 	"testing"

// 	v1 "github.com/punchguard/v1"
// )

// func TestNewId(t *testing.T) {
// 	wantSignaling := "s"
// 	wantOtr := "o"
// 	wantGuard := "g"

// 	id := v1.NewId(wantSignaling, wantOtr, wantGuard)

// 	if id.Signaling != wantSignaling {
// 		t.Fatalf("NewId().Signaling = '%v', want '%v'", id.Signaling, wantSignaling)
// 	}

// 	if id.Otr != wantOtr {
// 		t.Fatalf("NewId().Otr = '%v', want '%v'", id.Otr, wantOtr)
// 	}

// 	if id.Guard != wantGuard {
// 		t.Fatalf("NewId().Guard = '%v', want '%v'", id.Guard, wantGuard)
// 	}
// }
// func TestNewPeers(t *testing.T) {
// 	wantLocalAddress := "local"
// 	wantLocalPort := 1
// 	wantRemoteAddress := "remote"
// 	wantRemotePort := 2

// 	peers := v1.NewPeers(wantLocalAddress, wantLocalPort, wantRemoteAddress, wantRemotePort)

// 	if peers.Local.Address != wantLocalAddress {
// 		t.Fatalf("NewPeers().Local.Address = '%v', want '%v'", peers.Local.Address, wantLocalAddress)
// 	}

// 	if peers.Local.Port != wantLocalPort {
// 		t.Fatalf("NewPeers().Local.Port = '%v', want '%v'", peers.Local.Port, wantLocalPort)
// 	}

// 	if peers.Remote.Address != wantRemoteAddress {
// 		t.Fatalf("NewPeers().Remote.Address = '%v', want '%v'", peers.Remote.Address, wantRemoteAddress)
// 	}

// 	if peers.Remote.Port != wantRemotePort {
// 		t.Fatalf("NewPeers().Remote.Port = '%v', want '%v'", peers.Remote.Port, wantRemotePort)
// 	}
// }

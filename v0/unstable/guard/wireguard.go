package guard

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"github.com/punchguard/v0"
	"golang.zx2c4.com/wireguard/conn"
	"golang.zx2c4.com/wireguard/device"
	"golang.zx2c4.com/wireguard/ipc"
	"golang.zx2c4.com/wireguard/tun"
)

type WireGuard struct {
	punchguard.BaseGuard
	state *State
}

func (w *WireGuard) Init(id string) error {
	// log.Printf("%T.Init", w)

	if err := w.InitBase(id); err != nil {
		return err
	}

	return nil
}

const Version = "0.0.0"

const (
	ExitSetupSuccess = 0
	ExitSetupFailed  = 1
)

const (
	ENV_WG_TUN_FD             = "WG_TUN_FD"
	ENV_WG_UAPI_FD            = "WG_UAPI_FD"
	ENV_WG_PROCESS_FOREGROUND = "WG_PROCESS_FOREGROUND"
)

// func printUsage() {
// 	log.Printf("Usage: %s [-f/--foreground] INTERFACE-NAME\n", os.Args[0])
// }

func warning() {
	switch runtime.GOOS {
	case "linux", "freebsd", "openbsd":
		if os.Getenv(ENV_WG_PROCESS_FOREGROUND) == "1" {
			return
		}
	default:
		return
	}

	log.Print("┌──────────────────────────────────────────────────────┐")
	log.Print("│                                                      │")
	log.Print("│   Running wireguard-go is not required because this  │")
	log.Print("│   kernel has first class support for WireGuard. For  │")
	log.Print("│   information on installing the kernel module,       │")
	log.Print("│   please visit:                                      │")
	log.Print("│         https://www.wireguard.com/install/           │")
	log.Print("│                                                      │")
	log.Print("└──────────────────────────────────────────────────────┘")
}

type InterfaceConfig struct {
	Port       int
	PrivateKey string
}

type PeerConfig struct {
	Endpoint                    string
	PublicKey                   string
	AllowedIp                   string
	PersistentKeepaliveInterval uint
}

func base64ToHex(value string) string {
	bytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		log.Printf("base64ToHex %s\n", err)
	}
	return hex.EncodeToString(bytes)
}

func formatInterfaceConfig(interfaceConfig *InterfaceConfig) string {
	return fmt.Sprintf(`private_key=%s
listen_port=%d`, base64ToHex(interfaceConfig.PrivateKey), interfaceConfig.Port)
}

func formatPeerConfig(peerConfig *PeerConfig) string {
	return fmt.Sprintf(`public_key=%s
endpoint=%s
allowed_ip=%s
persistent_keepalive_interval=%d
`, base64ToHex(peerConfig.PublicKey), peerConfig.Endpoint, peerConfig.AllowedIp, peerConfig.PersistentKeepaliveInterval)
}

type State struct {
	term          chan struct{}
	errs          chan error
	device        *device.Device
	uapi          net.Listener
	logger        *device.Logger
	interfaceName string
}

func wireguardInit(id string, peers *punchguard.Peers, network string) (InterfaceConfig, PeerConfig) {
	var privateKey, publicKey string

	if id == "g1" {
		privateKey = "mBXRGMWfyFSASXlSCNXptikFufPZKCrWu+mQRT5AAFg="
		publicKey = "xs45uRfwxXv208Uab267cyV3P4+CA6OX3TYO9jFQ9Ws="
	} else {
		privateKey = "MKO4pZryVtCP5WC7A0TQuEtidDWBfkAqY52dSPt9JW8="
		publicKey = "OxKOJeHypzZv4xLbv4zVA0lQbRPo2UfAx4iqndGLDw0="
	}

	interfaceConfig := InterfaceConfig{
		Port:       peers.Local.Port,
		PrivateKey: privateKey,
	}

	peerConfig := PeerConfig{
		Endpoint:                    fmt.Sprintf("%s:%d", peers.Remote.Address, peers.Remote.Port),
		PublicKey:                   publicKey,
		AllowedIp:                   network,
		PersistentKeepaliveInterval: 10,
		// PersistentKeepaliveInterval: 25,
	}

	return interfaceConfig, peerConfig
}

func Wireguard(interfaceConfig *InterfaceConfig, peerConfig *PeerConfig) *State {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		log.Printf("wireguard-go v%s\n\nUserspace WireGuard daemon for %s-%s.\nInformation available at https://www.wireguard.com.\nCopyright (C) Jason A. Donenfeld <Jason@zx2c4.com>.\n", Version, runtime.GOOS, runtime.GOARCH)
		return nil
	}

	warning()

	var foreground bool
	var interfaceName string
	/* if len(os.Args) < 2 || len(os.Args) > 3 {
		printUsage()
		return
	}

	switch os.Args[1] {

	case "-f", "--foreground":
		foreground = true
		if len(os.Args) != 3 {
			printUsage()
			return
		}
		interfaceName = os.Args[2]

	default: */
	foreground = true
	/* if len(os.Args) != 2 {
		printUsage()
		return
	} */
	switch runtime.GOOS {
	case "linux":
		interfaceName = "pg10"
	case "darwin":
		interfaceName = "utun"
	default:
		log.Printf("OS [%s] not supported", runtime.GOOS)
	}
	//}

	if !foreground {
		foreground = os.Getenv(ENV_WG_PROCESS_FOREGROUND) == "1"
	}

	// get log level (default: info)

	logLevel := func() int {
		switch os.Getenv("LOG_LEVEL") {
		case "verbose", "debug":
			return device.LogLevelVerbose
		case "error":
			return device.LogLevelError
		case "silent":
			return device.LogLevelSilent
		}
		return device.LogLevelError
	}()

	// open TUN device (or use supplied fd)

	tun, err := func() (tun.Device, error) {
		tunFdStr := os.Getenv(ENV_WG_TUN_FD)
		if tunFdStr == "" {
			return tun.CreateTUN(interfaceName, device.DefaultMTU)
		}

		// construct tun device from supplied fd

		fd, err := strconv.ParseUint(tunFdStr, 10, 32)
		if err != nil {
			return nil, err
		}

		err = syscall.SetNonblock(int(fd), true)
		if err != nil {
			return nil, err
		}

		file := os.NewFile(uintptr(fd), "")
		return tun.CreateTUNFromFile(file, device.DefaultMTU)
	}()

	if err == nil {
		realInterfaceName, err2 := tun.Name()
		if err2 == nil {
			interfaceName = realInterfaceName
		}
	}

	logger := device.NewLogger(
		logLevel,
		fmt.Sprintf("(%s) ", interfaceName),
	)

	log.Printf("Starting wireguard-go version %s", Version)

	if err != nil {
		log.Printf("Failed to create TUN device: %v", err)
		os.Exit(ExitSetupFailed)
	}

	// open UAPI file (or use supplied fd)

	fileUAPI, err := func() (*os.File, error) {
		uapiFdStr := os.Getenv(ENV_WG_UAPI_FD)
		if uapiFdStr == "" {
			return ipc.UAPIOpen(interfaceName)
		}

		// use supplied fd

		fd, err := strconv.ParseUint(uapiFdStr, 10, 32)
		if err != nil {
			return nil, err
		}

		return os.NewFile(uintptr(fd), ""), nil
	}()
	if err != nil {
		log.Printf("UAPI listen error: %v", err)
		os.Exit(ExitSetupFailed)
		return nil
	}
	// daemonize the process

	if !foreground {
		env := os.Environ()
		env = append(env, fmt.Sprintf("%s=3", ENV_WG_TUN_FD))
		env = append(env, fmt.Sprintf("%s=4", ENV_WG_UAPI_FD))
		env = append(env, fmt.Sprintf("%s=1", ENV_WG_PROCESS_FOREGROUND))
		files := [3]*os.File{}
		if os.Getenv("LOG_LEVEL") != "" && logLevel != device.LogLevelSilent {
			files[0], _ = os.Open(os.DevNull)
			files[1] = os.Stdout
			files[2] = os.Stderr
		} else {
			files[0], _ = os.Open(os.DevNull)
			files[1], _ = os.Open(os.DevNull)
			files[2], _ = os.Open(os.DevNull)
		}
		attr := &os.ProcAttr{
			Files: []*os.File{
				files[0], // stdin
				files[1], // stdout
				files[2], // stderr
				tun.File(),
				fileUAPI,
			},
			Dir: ".",
			Env: env,
		}

		path, err := os.Executable()
		if err != nil {
			log.Printf("Failed to determine executable: %v", err)
			os.Exit(ExitSetupFailed)
		}

		process, err := os.StartProcess(
			path,
			os.Args,
			attr,
		)
		if err != nil {
			log.Printf("Failed to daemonize: %v", err)
			os.Exit(ExitSetupFailed)
		}
		process.Release()
		return nil
	}

	device := device.NewDevice(tun, conn.NewDefaultBind(), logger)

	err = device.IpcSet(formatInterfaceConfig(interfaceConfig))
	if err != nil {
		log.Printf("Failed to set interface: %v", err)
		os.Exit(ExitSetupFailed)
	}

	err = device.IpcSet(formatPeerConfig(peerConfig))
	if err != nil {
		log.Printf("Failed to set peer: %v", err)
		os.Exit(ExitSetupFailed)
	}

	log.Printf("Device started")

	state := State{
		errs:          make(chan error, 1),
		term:          make(chan struct{}, 1),
		logger:        logger,
		device:        device,
		interfaceName: interfaceName,
	}

	state.uapi, err = ipc.UAPIListen(interfaceName, fileUAPI)
	if err != nil {
		log.Printf("Failed to listen on uapi socket: %v", err)
		os.Exit(ExitSetupFailed)
	}

	go func() {
		for {
			conn, err := state.uapi.Accept()
			if err != nil {
				state.errs <- err
				return
			}
			go device.IpcHandle(conn)
		}
	}()

	log.Printf("UAPI listener started")

	// wait for program to terminate
	return &state
}

func (state *State) GetInterfaceName() string {
	return state.interfaceName
}

func (state *State) Wait() {

	select {
	case <-state.term:
	case <-state.errs:
	case <-state.device.Wait():
	}

	// clean up

}

func (state *State) Close() {
	state.uapi.Close()
	state.device.Close()

	log.Printf("Shutting down")
}

func (m *WireGuard) Start() punchguard.StoppedEvent {
	return m.MainLoop(func() {
		log.Printf("%T.MainLoop: started", m)
	}, func() {
		select {
		case peers := <-m.OnPeersEvent():
			log.Printf("%T.OnPeersEvent", m)
			m.WhenRunningAsync(func() {
				netConfig := netInit()

				interfaceConfig, peerConfig := wireguardInit(m.GetId(), &peers, netConfig.Network)

				m.state = Wireguard(&interfaceConfig, &peerConfig)

				err := netConfig.SetIp(m.state.GetInterfaceName())
				if err != nil {
					log.Printf("init.setip: %s\n", err)
				}

				defer func() {
					m.state.Close()
					m.Disconnected()
				}()

				m.WhenRunningAsync(func() {
					log.Print("")
					log.Printf("%T.OnPeersEvent sleep 10s", m)
					log.Print("")
					<-time.After(10 * time.Second)
					log.Print("")
					log.Printf("%T.OnPeersEvent: 10s timeout", m)
					log.Print("")
					m.state.term <- *new(struct{})
				})

				m.Connected()
				m.state.Wait()

			})
		default:
		}
	}, func() {
		log.Printf("%T.MainLoop: stopped", m)
		m.Close()
	})
}

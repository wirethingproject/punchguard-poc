package guard

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

func netInit() NetConfig {
	var ip string

	// if isControlling {
	ip = "100.64.0.1"
	// } else {
	// 	ip = "100.64.0.2"
	// }

	return NetConfig{
		Network: "100.64.0.0/10",
		Ip:      ip,
	}
}

type NetConfig struct {
	Network string
	Ip      string
}

func command(name string, arg ...string) error {
	fmt.Printf("command run: %s %s\n", name, arg)
	cmd := exec.Command(name, arg...)
	combinedOut, runErr := cmd.CombinedOutput()
	combinedOutStr := string(combinedOut)
	if combinedOutStr != "" {
		fmt.Printf("command out: %s\n", combinedOutStr)
	}
	if runErr != nil {
		fmt.Printf("command err: %s\n", runErr)
		return runErr
	}
	return nil
}

func (netConfig *NetConfig) SetIpDarwin(interfaceName string) error {
	masked := netConfig.Ip + "/" + strings.Split(netConfig.Network, "/")[1]
	err := command("ifconfig", interfaceName, "inet", masked, netConfig.Ip, "alias")
	if err == nil {
		err = command("route", "-q", "-n", "add", "-inet", netConfig.Network, "-interface", interfaceName)
	}
	return err
}

func (netConfig *NetConfig) SetIpLinux(interfaceName string) error {
	masked := netConfig.Ip + "/" + strings.Split(netConfig.Network, "/")[1]
	/* err := command("ip", "link", "add", interfaceName, "type", "wireguard")
	if err != nil {
			return err
	} */
	_ = command("ip", "-4", "address", "add", masked, "dev", interfaceName)
	_ = command("ip", "link", "set", "mtu", "1420", "up", "dev", interfaceName)
	_ = command("ip", "-4", "route", "add", netConfig.Network, "dev", interfaceName)
	return nil
}

func (netConfig *NetConfig) SetIp(interfaceName string) error {
	switch runtime.GOOS {
	case "linux":
		return netConfig.SetIpLinux(interfaceName)
	case "darwin":
		return netConfig.SetIpDarwin(interfaceName)
	default:
		return errors.New(fmt.Sprintf("OS [%s] not supported", runtime.GOOS))
	}
}

**INFORMATION**: _This repository served to prove the initial idea, but was archived because it ended up being too complex to be useful._

# PunchGuard - behind nat p2p vpn

This project aims to establish peer-to-peer vpn connections between devices that are behind NAT in the simplest and secure way possible without the need to use a service from a company that requires a login or manage a public facing server yourself.

Project status: proof-of-concept worked, now the focus is on making it generally usable.

## Security Warning

This project's code and some of the cryptographic libraries it uses have not been formally audited and therefore there is no way to guarantee that it's free of security bugs. Before use, consider this information in your risk assessment.

The biggest problem in using a software like that is the protection of your private keys, if one of the devices is compromised these keys can be used to establish a connection with the other pair. However, if your device has been compromised, chances are these private keys are the least of your problems.

# History

This project was born from the impossibility of making a simple VPN connection between two devices in separate locations without depending on an external service or opening ports on a router. Although [WireGuard](https://www.wireguard.com/) is simple, fast and secure, having an open door to the public Internet will always be uncomfortable. Also, due [Carrier-grade NAT](https://en.wikipedia.org/wiki/Carrier-grade_NAT) is often used for mitigating [IPv4 exhaustion](https://en.wikipedia.org/wiki/IPv4_address_exhaustion), opening a port on a router no longer works for many providers.

It's possible to solve this problem using the [ICE](https://en.wikipedia.org/wiki/Interactive_Connectivity_Establishment) protocol that do a [hole punching](https://en.wikipedia.org/wiki/Hole_punching_(networking)) for NAT traversal, the same used for VoIP communications. The problem with this protocol is that it relies on an already established connection between the two devices to exchange information on how to establish this new connection. It's the chicken-and-egg problem and this is where other solutions need an open Internet service and require a login from you.

There aren't many alternatives, but luckily it's possible to make two devices communicate without login using p2p instant messaging. Initially [Tox](https://github.com/TokTok/c-toxcore) will be used, because it's simple to integrate and meets the requirement of not having a login and central servers (the network is made up of its users), in the future other alternatives will be considered.

Finally, as it's very difficult to guarantee the confidentiality of the signaling, [Off-the-Record](https://otr.cypherpunks.ca/Protocol-v3-4.0.0.html) protocol will be used for the communication to occur in a secure way.

## Why Go?

These things run concurrently, Go makes solving concurrent problems easier, although solving concurrent problems will never be easy.

Also for portability, safe memory, speed, and simplicity, but mostly for the single binary thing.

# Roadmap

The plan is to have a stable version 1.0.0 as soon as possible. To achieve this, only the minimum necessary features to work will be included.

## 1.0.0

- Signaling using [go-toxcore](https://github.com/TokTok/go-toxcore-c)
- Off-the-Record using [OTR3](https://github.com/coyim/otr3)
- Punch using [Pion ICE](https://github.com/pion/ice)
- Guard using [wireguard-go](https://git.zx2c4.com/wireguard-go/about/)
- A desktop-to-desktop flow that punches and reconnects when guard disconnects
- A stable implementation-agnostic API
- Command line and daemon in a single binary
- One-to-one/single peer only configuration
- Manual setup and pairing, using file transfer or copy-and-paste
- UDP only ICE
- Comprehensive test coverage
- Setup and configuration of network interfaces
- Security awareness and considerations
- OS/Arch:
  - Linux: amd64, arm
  - Darwin: amd64

## Future Work

- Replace Tox C implementation by a native Go code or use other library
- Mobile library for embedding
- A mobile-to-desktop flow that constantly punches for the best connection
- If possible, a device-to-server flow that supports many-to-one/multiple peers configuration
- If possible, a network routing flow, I don't think it's a good idea for this to be used by someone with no experience in network administration
- Secure firewall rules for this VPN connection
- A SOCKS Proxy flow that I don't think it's a good idea since using WireGuard is more secure, but maybe someone can't use UDP
- TCP ICE
- Protobuf protocol for no manual pairing using the signaling implementation
- Signaling using TOR
- Signaling using a central service or server
- Guard using OpenVPN
- OS/Arch:
  - Darwin: arm64
  - Windows

# License

PunchGuard source code is licensed under [MIT License](https://opensource.org/licenses/MIT). You can find the complete text in [LICENSE](LICENSE).

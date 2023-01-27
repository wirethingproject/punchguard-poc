# PunchGuard - behind nat p2p vpn

This project aims to establish peer-to-peer vpn connections between devices that are behind NAT in the simplest way possible without the need to use a service from a company that has a login or manage a public facing server yourself.

## Security Warning

This project's code and some of the cryptographic libraries it uses have not been formally audited and therefore there is no way to guarantee that it's free of security issues. Before use, consider this information in your risk analysis.

## History

This project was born from the impossibility of making a simple VPN connection between two devices in separate locations without depending on an external service or opening ports on a router. Although [WireGuard](https://www.wireguard.com/) is simple, fast and secure, having an open door to the public Internet will always be uncomfortable. Also, due [Carrier-grade NAT](https://en.wikipedia.org/wiki/Carrier-grade_NAT) is often used for mitigating [IPv4 exhaustion](https://en.wikipedia.org/wiki/IPv4_address_exhaustion), opening a port on a router no longer works for many providers.

It is possible to solve this problem using the [ICE](https://en.wikipedia.org/wiki/Interactive_Connectivity_Establishment) protocol that do a [hole punching](https://en.wikipedia.org/wiki/Hole_punching_(networking)) for NAT traversal, the same used for VoIP communications. The problem with this protocol is that it relies on an already established connection between the two devices to exchange information on how to establish this new connection. It's the chicken-and-egg problem and this is where other solutions need an open Internet service and require a login from you.

There aren't many alternatives, but luckily it's possible to make two devices communicate without login using p2p instant messaging. Initially [Tox](https://github.com/TokTok/c-toxcore) will be used, because it is simple to integrate and meets the requirement of not having a login and central servers (the network is made up of its users), in the future other alternatives will be considered.

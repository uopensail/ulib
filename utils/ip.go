package utils

import (
	"errors"
	"net"
)

// GetLocalIP returns the first non-loopback IPv4 address of the local machine.
func GetLocalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ip := getIPFromAddr(addr); ip != nil {
				return ip.String(), nil
			}
		}
	}
	return "", errors.New("no suitable IPv4 address found")
}

// GetAllLocalIPs returns all non-loopback IPv4 addresses of the local machine.
func GetAllLocalIPs() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 || iface.Flags&net.FlagLoopback != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, addr := range addrs {
			if ip := getIPFromAddr(addr); ip != nil {
				ips = append(ips, ip.String())
			}
		}
	}

	if len(ips) == 0 {
		return nil, errors.New("no suitable IPv4 addresses found")
	}
	return ips, nil
}

// GetLocalIPByInterface returns the IPv4 address of the named network interface.
func GetLocalIPByInterface(interfaceName string) (string, error) {
	iface, err := net.InterfaceByName(interfaceName)
	if err != nil {
		return "", err
	}
	if iface.Flags&net.FlagUp == 0 {
		return "", errors.New("interface is down")
	}

	addrs, err := iface.Addrs()
	if err != nil {
		return "", err
	}
	for _, addr := range addrs {
		if ip := getIPFromAddr(addr); ip != nil {
			return ip.String(), nil
		}
	}
	return "", errors.New("no suitable IPv4 address found on interface " + interfaceName)
}

// IsValidIPv4 reports whether s is a valid IPv4 address.
func IsValidIPv4(s string) bool {
	ip := net.ParseIP(s)
	return ip != nil && ip.To4() != nil
}

// GetOutboundIP returns the preferred outbound IP of this machine by opening a
// UDP socket toward 8.8.8.8 (no packet is actually sent).
func GetOutboundIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.String(), nil
}

// getIPFromAddr extracts an IPv4 address from addr, returning nil for loopback
// or non-IPv4 addresses.
func getIPFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	default:
		return nil
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	return ip.To4()
}

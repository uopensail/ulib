package utils

import (
	"errors"
	"net"
)

// GetLocalIp retrieves the first non-loopback IPv4 address of the local machine
//
// @return: Local IPv4 address as string and error if no suitable address found
// @note: Returns the first valid non-loopback IPv4 address found
func GetLocalIp() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		// Skip interfaces that are down or loopback
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue // Skip this interface if address retrieval fails
		}

		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("no suitable IPv4 address found - check network connectivity")
}

// getIpFromAddr extracts the IP address from a net.Addr interface
// Filters out loopback and non-IPv4 addresses
//
// @param addr: Network address interface
// @return: IPv4 address or nil if not valid
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	default:
		return nil // Unsupported address type
	}

	if ip == nil || ip.IsLoopback() {
		return nil
	}

	ip = ip.To4()
	if ip == nil {
		return nil // not an IPv4 address
	}

	return ip
}

// GetAllLocalIps retrieves all non-loopback IPv4 addresses of the local machine
//
// @return: Slice of local IPv4 addresses as strings and error if retrieval fails
func GetAllLocalIps() ([]string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	var ips []string
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}

		addrs, err := iface.Addrs()
		if err != nil {
			continue // Skip this interface
		}

		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip != nil {
				ips = append(ips, ip.String())
			}
		}
	}

	if len(ips) == 0 {
		return nil, errors.New("no suitable IPv4 addresses found")
	}
	return ips, nil
}

// GetLocalIpByInterface retrieves the IPv4 address of a specific network interface
//
// @param interfaceName: Name of the network interface (e.g., "eth0", "en0")
// @return: IPv4 address as string and error if interface not found or no suitable address
func GetLocalIpByInterface(interfaceName string) (string, error) {
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
		ip := getIpFromAddr(addr)
		if ip != nil {
			return ip.String(), nil
		}
	}

	return "", errors.New("no suitable IPv4 address found on interface " + interfaceName)
}

// IsValidIPv4 checks if a string is a valid IPv4 address
//
// @param ipStr: String to validate as IPv4
// @return: true if valid IPv4, false otherwise
func IsValidIPv4(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}
	return ip.To4() != nil
}

// GetOutboundIP gets the preferred outbound IP address of this machine
// by establishing a connection to a public DNS server
//
// @return: Outbound IPv4 address as string and error if connection fails
func GetOutboundIP() (string, error) {
	// Using Google's public DNS server to determine outbound IP
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}

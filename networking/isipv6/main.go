package main

import (
	"fmt"
	"net/netip"
)

func main() {
	fmt.Println("Hello, 世界")
	// http.HTTP2Config{}
	ipStrs := []string{
		"2001:db8::68",     // Valid IPv6
		"192.0.2.1",        // Valid IPv4
		"::ffff:192.0.2.1", // Valid IPv4-mapped IPv6
		"invalid-ip",       // Invalid IP
		"2001:db8::68::1",  // Invalid IPv6
		"256.256.256.256",  // Invalid IPv4
	}
	for _, ipStr := range ipStrs {
		addr, err := netip.ParseAddr(ipStr)
		if err != nil {
			fmt.Printf("Failed to parse IP address: %v\n", err)
		} else {
			fmt.Printf("Parsed Address: %v %v\n", addr, addr.Is6())
		}
	}
}

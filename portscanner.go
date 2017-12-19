
package portscanner

import (
		"fmt"
		"net"
		"time"
		"regexp"
		"strings"
		"strconv"
		)

// IPv4 is the type used for IP addresses.
type IPv4 [4]int

// Check checks if an error is nil or a significant error. 
func Check(err error) {

	if err != nil {
		panic(err)
	}
}

// ToInt converts a string to integer, as strconv.Atoi does, but without 
// returning errors.
func ToInt(s string) int {

	i, _ := strconv.Atoi(s)
	return i
}

// PrintBanner presents a simple banner just to identify program name.
func PrintBanner() {

	fmt.Println("\n +----------------+")
	fmt.Println(" | Go PortScanner |")
	fmt.Println(" +----------------+")
}

// ToString converts an IP from IPv4 type to string.
func (ip *IPv4)ToString() string {
	
	ip_stringed := strconv.Itoa(ip[0])
	for i := 1; i < 4; i++ {
		str_i := strconv.Itoa(ip[i])
		ip_stringed += "." + str_i
	}
	return ip_stringed
}

// IsValid checks an IP address as valid or not.
func (ip *IPv4)IsValid() bool { 

	for i, oct := range ip {
		if i == 0 || i == 3 {
			if oct < 1 || oct > 255 {
				return false
			}
		} else {
			if oct < 1 || oct > 255 {
				return false
			}
		}
	}
	return true
}

// PlusPlus increments an IPv4 value.
func (ip *IPv4)PlusPlus() *IPv4 {

	if ip[3] <= 254 {
		ip[3] = ip[3] + 1
	} else {
		if ip[2] < 255 {
			ip[2] = ip[2] + 1
		} else {
			if ip[1] < 255 {
				ip[1] = ip[1] + 1
			} else {
				if ip[0] < 255 {
					ip[0] = ip[0] + 1
				}
			}
		}
	}
	return ip
}

// ToIPv4 converts an string to a IPv4.
func ToIPv4(ip string) IPv4 {

	var new_ip IPv4
	
	ip_s := strings.Split(ip, ".")
	
	for i, v := range (ip_s) {
		new_ip[i], _ = strconv.Atoi(v)
	}

	return new_ip
}

// ParseIPSequence gets a sequence of IP addresses correspondent from an 
// "init-end" entry.
func ParseIPSequence(ip_sequence string) []IPv4 {

	var array_ips []IPv4

	series, _ := regexp.Compile("([0-9]+)")

	// For sequence ips, using '-'
	l_series := series.FindAllStringSubmatch(ip_sequence,-1)

	for i := ToInt(l_series[3][0]); i <= ToInt(l_series[4][0]); i++ {
		array_ips = append(array_ips, IPv4{
								ToInt(l_series[0][0]), 
								ToInt(l_series[1][0]), 
								ToInt(l_series[2][0]), 
								i})
	}
	return array_ips
}

// ParsePortList gets a port list from its port entry in arguments.
func ParsePortList(raw_ports string) []string {

	var ports []string

	individuals, _ := regexp.Compile("([0-9]+)[,]*")
	series, _ := regexp.Compile("([0-9]+)[-]([0-9]+)")

	// For individual ports, separated by ','
	l_individuais := individuals.FindAllStringSubmatch(raw_ports,-1)

	// For sequence ports, using '-'
	l_series := series.FindAllStringSubmatch(raw_ports,-1)

	if len(l_series)>0 {
		for _, s := range l_series {
			init, _ := strconv.Atoi(s[1])
			end, _  := strconv.Atoi(s[2])
			for i := init; i < end; i++ {
				ports = append(ports, strconv.Itoa(i))
			}
		}
	}
	for _, port := range l_individuais {
		ports = append(ports, port[1])
	}
	return ports
}

// PresentResults presents all results in console.
func PresentResults(open_ports []string) {

	fmt.Println("Port:	Description:")
	for _, port := range open_ports {
		fmt.Println(port + "\t" + port_short_list[port])
	}
}

// PortScanner scans IP:port pairs looking for open ports on IP addresses.
func PortScanner(ip IPv4, port_list []string) []string {

	var open 	[]string

	for _, port := range port_list {

		conn, err := net.DialTimeout("tcp", 
								ip.ToString()+ ":"+ port, 
								100 * time.Millisecond)

		if err == nil {
			conn.Close()
			open = append(open, port)
		}
	}
	return open
}

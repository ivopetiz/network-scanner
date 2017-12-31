package portscanner

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
)

// IPv4 is the type used for IP addresses.
type IPv4 [4]int

/*struct IPPortList{
	ip		IPv4
	ports	[]string
}*/

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

// ToString converts an IP from IPv4 type to string.
func (ip *IPv4) ToString() string {

	ip_stringed := strconv.Itoa(ip[0])
	for i := 1; i < 4; i++ {
		str_i := strconv.Itoa(ip[i])
		ip_stringed += "." + str_i
	}
	return ip_stringed
}

// IsValid checks an IP address as valid or not.
func (ip *IPv4) IsValid() bool {

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
func (ip *IPv4) PlusPlus() *IPv4 {

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

	for i, v := range ip_s {
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
	l_series := series.FindAllStringSubmatch(ip_sequence, -1)

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
	l_individuals := individuals.FindAllStringSubmatch(raw_ports, -1)

	// For sequence ports, using '-'
	l_series := series.FindAllStringSubmatch(raw_ports, -1)

	if len(l_series) > 0 {
		for _, s := range l_series {
			init, _ := strconv.Atoi(s[1])
			end, _ := strconv.Atoi(s[2])
			for i := init; i < end; i++ {
				ports = append(ports, strconv.Itoa(i))
			}
		}
	}
	for _, port := range l_individuals {
		ports = append(ports, port[1])
	}
	return ports
}

// GetAllIPsClassC returns a slice of IPv4 with all IP addresses 
// from a Class C.
//func GetAllIPsClassC(ip IPv4) []IPv4 {}

// PresentResults presents all results in console.
func PresentResults(ip IPv4, ports []string) {

	fmt.Println(" \n>" + ip.ToString())
	fmt.Println(" Port:	Description:")
	for _, port := range ports {
		fmt.Println(" " + port + "\t" + port_short_list[port])
	}
}

// PortScanner scans IP:port pairs looking for open ports on IP addresses.
func PortScanner(ip IPv4, port_list []string) []string {

	var open []string

/*	var wg sync.WaitGroup

	for _, port := range port_list {
		wg.Add(1)
		go func(port string) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp",
										ip.ToString()+":"+port,
										100*time.Millisecond)
			
			if err == nil {
				conn.Close()
				open = append(open, port)
			}
		}(port)
	}

	wg.Wait()
*/
	for _, port := range port_list {

		conn, err := net.DialTimeout("tcp",
									ip.ToString()+":"+port,
									100*time.Millisecond)
			
		if err == nil {
			conn.Close()
			open = append(open, port)
		}
	}
	
	return open
}


// IPScanner scans all IP addresses in ip_list for every port in port_list.
func IPScanner(ip_str []string, port_str []string, print_results bool) map[IPv4][]string {
	
	m := make(map[IPv4][]string)

	var ip_list []IPv4
	var port_list []string
	
	var wg sync.WaitGroup

	if len(port_str) == 1 {
		port_list = ParsePortList(port_str[0])
	} else {
		port_list = port_str
	}

	if len(ip_str) == 0 {
		ip_list = append(ip_list, IPv4{127, 0, 0, 1})
	} else {
		for _, i := range ip_str {
			if strings.Contains(i, "-") {
				ip_list = append(ip_list, ParseIPSequence(i)...)
			} else {
				ip := ToIPv4(i)
				if ip.IsValid() {
					ip_list = append(ip_list, ip)
				}
			}
		}
	}
/*
	for _, ip := range ip_list {
		result := PortScanner(ip, port_list)
		if len(result) > 0 {
			m[ip] = result
			if print_results {
				PresentResults(ip, result)
			}
		}
	}
*/


	for _, ip := range ip_list {
		wg.Add(1)
		go func(ip IPv4) {
			defer wg.Done()
			result := PortScanner(ip, port_list)
			if len(result) > 0 {
				m[ip] = result
				if print_results {
					PresentResults(ip, result)
				}
			}
		}(ip)
	}

	wg.Wait()
	
	return m
}

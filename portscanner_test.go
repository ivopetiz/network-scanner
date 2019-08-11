package portscanner

import (
	"reflect"
	"strconv"
	"testing"
)

func TestToInt(t *testing.T) {

	var v int
	v = ToInt("123")

	i, _ := strconv.Atoi("123")

	if v != 123 {
		t.Error("Expected 123, got ", i)
	}
}

func TestToString(t *testing.T) {

	var ip IPv4
	var v string

	ip = IPv4{127, 0, 0, 1}

	v = ip.ToString()

	if v != "127.0.0.1" {
		t.Error("Expected 127.0.0.1, got ", v)
	}
}

func TestIsValid(t *testing.T) {

	var ip IPv4

	ip = IPv4{127, 0, 0, 1}
	if ip.IsValid() != true {
		t.Error("Expected true, got", ip.IsValid())
	}

	ip = IPv4{192, 168, 1, 100}
	if ip.IsValid() != true {
		t.Error("Expected true, got", ip.IsValid())
	}

	ip = IPv4{254, 254, 254, 254}
	if ip.IsValid() != true {
		t.Error("Expected true, got", ip.IsValid())
	}

	ip = IPv4{254, 254, 254, 0}
	if ip.IsValid() != false {
		t.Error("Expected false, got", ip.IsValid())
	}

	ip = IPv4{123, 234, 456, 2}
	if ip.IsValid() != false {
		t.Error("Expected false, got", ip.IsValid())
	}

	ip = IPv4{254, 255, 254, 250}
	if ip.IsValid() != true {
		t.Error("Expected true, got", ip.IsValid())
	}
}

func TestPlusPlus(t *testing.T) {

	var ip, nextIP IPv4

	ip = IPv4{183, 145, 10, 10}
	nextIP = IPv4{183, 145, 10, 11}

	if ip.PlusPlus().ToString() != nextIP.ToString() {
		t.Error("Expected 183.145.10.11, got", ip.PlusPlus().ToString())
	}

	ip = IPv4{183, 145, 10, 254}
	nextIP = IPv4{183, 145, 11, 1}

	if ip.PlusPlus().ToString() != nextIP.ToString() {
		t.Error("Expected 183.145.11.1, got", ip.PlusPlus().ToString())
	}

	ip = IPv4{183, 145, 255, 254}
	nextIP = IPv4{183, 146, 1, 1}

	if ip.PlusPlus().ToString() != nextIP.ToString() {
		t.Error("Expected 183.146.1.1, got", ip.PlusPlus().ToString())
	}

	ip = IPv4{183, 255, 255, 254}
	nextIP = IPv4{184, 1, 1, 1}

	if ip.PlusPlus().ToString() != nextIP.ToString() {
		t.Error("Expected 183.146.1.1, got", ip.PlusPlus().ToString())
	}
}

func TestToIPv4(t *testing.T) {

	var ipString string
	var ip IPv4

	ipString = "183.145.10.11"
	ip = IPv4{183, 145, 10, 11}

	if ToIPv4(ipString) != ip {
		t.Error("Expected 183.145.10.11, got", ip.ToString())
	}
}

func TestParseIPSequence(t *testing.T) {

	var IPsec string
	var listIPv4 []IPv4

	IPsec = "192.168.1.1-10"
	listIPv4 = append(listIPv4,
		IPv4{192, 168, 1, 1},
		IPv4{192, 168, 1, 2},
		IPv4{192, 168, 1, 3},
		IPv4{192, 168, 1, 4},
		IPv4{192, 168, 1, 5},
		IPv4{192, 168, 1, 6},
		IPv4{192, 168, 1, 7},
		IPv4{192, 168, 1, 8},
		IPv4{192, 168, 1, 9},
		IPv4{192, 168, 1, 10},
	)

	if !reflect.DeepEqual(ParseIPSequence(IPsec), listIPv4) {
		t.Error("Expected IPs from 1 to 10, got", ParseIPSequence(IPsec))
	}
}

func TestParsePortList(t *testing.T) {

	var portSeq string
	var ports []string

	portSeq = "1234,1235,1236,1237,1238"
	ports = append(ports, "1234",
		"1235",
		"1236",
		"1237",
		"1238")

	if !reflect.DeepEqual(ParsePortList(portSeq), ports) {
		t.Error("Expected ports 1234,1235,1236,1237,1238, got", ParsePortList(portSeq))
	}

	portSeq = "1234-1238"

	if !reflect.DeepEqual(ParsePortList(portSeq), ports) {
		t.Error("Expected ports 1234,1235,1236,1237,1238, got", ParsePortList(portSeq))
	}
}

//func TestPresentResults(t *testing.T) {}

func TestPortScanner(t *testing.T) {

	var ip IPv4
	var open, portList []string

	ip = IPv4{127, 0, 0, 1}
	portList = append(portList, "123456")

	if !reflect.DeepEqual(PortScanner(ip, portList), open) {
		t.Error("Expected empty list, got", PortScanner(ip, portList))
	}
}

func TestIPScanner(t *testing.T) {

	var ipStr []string
	//var ipList []IPv4
	var portList []string

	openMap := make(map[IPv4][]string)

	ipStr = append(ipStr, "127.0.0.1")

	portList = append(portList, "123456")

	IPScanner(ipStr, portList, true)

	if !reflect.DeepEqual(IPScanner(ipStr, portList, false), openMap) {
		t.Error("Expected empty list, got",
			IPScanner(ipStr, portList, false), openMap)
	}
}

//ver: cold little heart

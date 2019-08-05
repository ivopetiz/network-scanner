package portscanner

import (
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

	/*	if ip.PlusPlus().ToString() != nextIP.ToString() {
		t.Error("Expected 183.146.1.1, got", ip.PlusPlus().ToString())
	}*/
}

//func TestIsValid(t *testing.T) {}

//func TestIsValid(t *testing.T) {}

//func TestIsValid(t *testing.T) {}

package main

import (
	"net"
	"testing"
)

type parseIPTest struct {
	input    string
	isV4     bool
	isV6     bool
	ip       net.IP
	erroring bool
}

var parseIPTests = []parseIPTest{
	{"10.0.0.1", true, false, net.IP{10, 0, 0, 1}, false},
	{"fe80::1", false, true, net.IP{254, 128, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}, false},
	{"10.0.0.1", true, false, net.IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 255, 10, 0, 0, 1}, false},
	{"NOT A VALID IP", false, false, nil, true},
	{"256.256.256.256", false, false, nil, true},
}

func TestParseIP(t *testing.T) {
	for _, test := range parseIPTests {
		isV4, isV6, ip, err := ParseIP(test.input)

		if isV4 != test.isV4 {
			t.Errorf("isV4 %t not equal to expected %t", isV4, test.isV4)
		}

		if isV6 != test.isV6 {
			t.Errorf("isV6 %t not equal to expected %t", isV6, test.isV6)
		}

		if !ip.Equal(test.ip) {
			t.Errorf("IP %s not equal to expected %s (%+v %+v)", ip, test.ip, []byte(ip), []byte(test.ip))
		}

		if (err != nil && !test.erroring) || (err == nil && test.erroring) {
			t.Errorf("Error %+v received while function should error: %t", err, test.erroring)
		}
	}
}

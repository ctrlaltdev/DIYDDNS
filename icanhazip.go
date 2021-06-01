package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var (
	client = http.Client{}
)

func IsIPv4(ip net.IP) bool {
	IPB := []byte(ip)

	if len(IPB) == 4 {
		return true
	}

	if len(IPB) == 16 {
		sum10B := byte(0x00)
		for i := 0; i < 10; i++ {
			sum10B = sum10B + IPB[i]
		}

		if sum10B == 0x00 && IPB[10] == 0xFF && IPB[11] == 0xFF {
			return true
		}
	}

	return false
}

func IsIPv6(ip net.IP) bool {
	IPB := []byte(ip)

	if len(IPB) == 16 {
		sum10B := byte(0x00)
		for i := 0; i < 10; i++ {
			sum10B = sum10B + IPB[i]
		}

		if sum10B != 0x00 && IPB[10] != 0xFF && IPB[11] != 0xFF {
			return true
		}
	}

	return false
}

func ParseIP(input string) (isV4 bool, isV6 bool, ip net.IP, err error) {
	ip = net.ParseIP(strings.TrimSpace(input))

	if ip == nil {
		return isV4, isV6, ip, errors.New(fmt.Sprintf("%s is not a valid IP", strings.TrimSpace(input)))
	}

	return IsIPv4(ip), IsIPv6(ip), ip, err
}

func GetIP() (isV4 bool, isV6 bool, ip net.IP, err error) {
	req, err := http.NewRequest("GET", "https://icanhazip.com/", nil)
	if err != nil {
		return isV4, isV6, ip, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("DIYDDNS %s - https://ctrlalt.dev/DIYDDNS", VERSION))

	res, err := client.Do(req)
	if err != nil {
		return isV4, isV6, ip, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return isV4, isV6, ip, err
	}
	if len(body) <= 0 {
		return isV4, isV6, ip, errors.New("no IP received")
	}

	isV4, isV6, ip, err = ParseIP(string(body))
	return isV4, isV6, ip, err
}

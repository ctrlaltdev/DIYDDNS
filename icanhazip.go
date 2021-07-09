package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
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

func GetIPv(v6 bool) (ip net.IP, err error) {
	var url string

	if v6 {
		url = "https://ipv6.icanhazip.com"
	} else {
		url = "https://ipv4.icanhazip.com"
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ip, err
	}

	req.Header.Set("User-Agent", fmt.Sprintf("DIYDDNS %s - https://ctrlalt.dev/DIYDDNS", VERSION))

	res, err := client.Do(req)
	if err != nil {
		return ip, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ip, err
	}

	_, _, ip, err = ParseIP(string(body))
	return ip, err
}

type IPs struct {
	v4 net.IP
	v6 net.IP
}

func GetIPs() (ips IPs, err error) {
	ipv4, err := GetIPv(false)
	if err != nil {
		log.Println(err)
	}

	ipv6, err := GetIPv(true)
	if err != nil {
		log.Println(err)
	}

	return IPs{ipv4, ipv6}, err
}

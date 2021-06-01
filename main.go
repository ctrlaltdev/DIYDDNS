package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/cloudflare/cloudflare-go"
	"gopkg.in/yaml.v3"
)

var (
	VERSION = "v2.0.0"

	shouldInit = flag.Bool("init", false, "Run the initialization process")
	FQDN       = flag.String("fqdn", "", "The full domain that should be used")

	DOMAIN string

	CONF CFConf
)

func runInit() {
	home, err := os.UserHomeDir()
	CheckErr(err)

	CreateFolderIfNotExists(filepath.Join(home, ".DIYDDNS"), 0700)

	var conf CFConf

	fmt.Print("CloudFlare API Email: ")
	fmt.Scanln(&conf.API_EMAIL)

	fmt.Print("CloudFlare API Key: ")
	fmt.Scanln(&conf.API_KEY)

	serial, err := yaml.Marshal(conf)
	CheckErr(err)

	WriteFile(filepath.Join(home, ".DIYDDNS", "conf.yaml"), string(serial))

	os.Exit(0)
}

func loadConf() {
	home, err := os.UserHomeDir()
	CheckErr(err)

	data := ReadFile(filepath.Join(home, ".DIYDDNS", "conf.yaml"))

	if data == "" {
		log.Fatalln("No configuration found. Please run DIYDDNS -init first.")
	}

	err = yaml.Unmarshal([]byte(data), &CONF)
	CheckErr(err)
}

func runCheck() {
	isV4, isV6, ip, err := GetIP()
	CheckErr(err)

	var records []cloudflare.DNSRecord

	if isV4 {
		records, err = GetRecords(DOMAIN, *FQDN, "A")
	}
	if isV6 {
		records, err = GetRecords(DOMAIN, *FQDN, "AAAA")
	}
	CheckErr(err)

	if len(records) > 1 {
		log.Fatalf("%d records were found in cloudflare, we only support single record at the moment", len(records))
	}

	if len(records) == 0 {

		var recordType string

		if isV4 {
			recordType = "A"
		}
		if isV6 {
			recordType = "AAAA"
		}

		err = CreateRecord(DOMAIN, *FQDN, recordType, ip)
		CheckErr(err)

	} else {

		if ip.String() == records[0].Content {
			os.Exit(0)
		}

		err := UpdateRecord(records[0], ip)
		CheckErr(err)

	}
}

func main() {
	flag.Parse()

	if *shouldInit {
		runInit()
	}

	if *FQDN != "" {
		loadConf()
		DOMAIN = GetRootDomain(*FQDN)
		runCheck()
	} else {
		fmt.Println("No FQDN passed, you need to define it with the fqdn flag. Run DIYDDN -h to get more info")
	}
}

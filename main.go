package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudflare/cloudflare-go"
	_ "github.com/joho/godotenv/autoload"
)

var (
	VERSION = "v2.0.0"

	DOMAIN = os.Getenv("DOMAIN")
	FQDN   = fmt.Sprintf("devlocal.%s", DOMAIN)
)

func main() {
	isV4, isV6, ip, err := GetIP()
	CheckErr(err)

	var records []cloudflare.DNSRecord

	if isV4 {
		records, err = GetRecords(DOMAIN, FQDN, "A")
	}
	if isV6 {
		records, err = GetRecords(DOMAIN, FQDN, "AAAA")
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

		err = CreateRecord(DOMAIN, FQDN, recordType, ip)
		CheckErr(err)

	} else {

		if ip.String() == records[0].Content {
			os.Exit(0)
		}

		err := UpdateRecord(records[0], ip)
		CheckErr(err)

	}

}

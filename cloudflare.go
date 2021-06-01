package main

import (
	"context"
	"net"
	"os"

	"github.com/cloudflare/cloudflare-go"
	_ "github.com/joho/godotenv/autoload"
)

var (
	CF_API_KEY   = os.Getenv("CF_API_KEY")
	CF_API_EMAIL = os.Getenv("CF_API_EMAIL")
)

func GetRecords(domain string, name string, recordType string) ([]cloudflare.DNSRecord, error) {
	api, err := cloudflare.New(CF_API_KEY, CF_API_EMAIL)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()

	_, err = api.UserDetails(ctx)
	if err != nil {
		return nil, err
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		return nil, err
	}

	host := cloudflare.DNSRecord{
		Name: name,
		Type: recordType,
	}

	records, err := api.DNSRecords(ctx, zoneID, host)

	return records, err
}

func CreateRecord(domain string, name string, recordType string, dst net.IP) error {
	api, err := cloudflare.New(CF_API_KEY, CF_API_EMAIL)
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = api.UserDetails(ctx)
	if err != nil {
		return err
	}

	zoneID, err := api.ZoneIDByName(domain)
	if err != nil {
		return err
	}

	rec := cloudflare.DNSRecord{
		Name:    name,
		Type:    recordType,
		Content: dst.String(),
	}

	_, err = api.CreateDNSRecord(ctx, zoneID, rec)
	if err != nil {
		return err
	}

	return nil
}

func UpdateRecord(record cloudflare.DNSRecord, dst net.IP) error {
	api, err := cloudflare.New(CF_API_KEY, CF_API_EMAIL)
	if err != nil {
		return err
	}

	ctx := context.Background()

	_, err = api.UserDetails(ctx)
	if err != nil {
		return err
	}

	rec := record
	rec.Content = dst.String()

	err = api.UpdateDNSRecord(ctx, record.ZoneID, record.ID, rec)
	if err != nil {
		return err
	}

	return nil
}

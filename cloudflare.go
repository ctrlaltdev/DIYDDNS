package main

import (
	"context"
	"net"

	"github.com/cloudflare/cloudflare-go"
)

type CFConf struct {
	API_KEY   string `yaml:"CF_API_KEY"`
	API_EMAIL string `yaml:"CF_API_EMAIL"`
}

func GetRecords(domain string, name string, recordType string) ([]cloudflare.DNSRecord, error) {
	api, err := cloudflare.New(CONF.API_KEY, CONF.API_EMAIL)
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
	api, err := cloudflare.New(CONF.API_KEY, CONF.API_EMAIL)
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
	api, err := cloudflare.New(CONF.API_KEY, CONF.API_EMAIL)
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

package main

import (
	"encoding/json"
	"os"

	"github.com/alyx/go-daddy/daddy"
)

type Godaddy struct {
	client  *daddy.Client
	domains []daddy.DomainSummary
}

var _ Registrar = (*Godaddy)(nil)

func (g *Godaddy) Load() error {
	client, err := daddy.NewClient(os.Getenv("GODADDY_KEY"), os.Getenv("GODADDY_SECRET"), false)
	if err != nil {
		return err
	}
	g.client = client

	myDomains, err := client.Domains.List([]string{"ACTIVE"}, nil, 0, "", nil, "")
	if err != nil {
		return err
	}
	g.domains = myDomains

	return nil
}
func (g *Godaddy) UpdateDomain(ip, domain string, names []string) error {
	records := make([]GoDaddyDNSRecord, len(names))
	for i, name := range names {
		records[i] = GoDaddyDNSRecord{
			Data: ip,
			Name: name,
			TTL:  3600,
		}
	}

	return ReplaceRecordsByType(g.client, domain, "A", records)
}

func (g *Godaddy) HasDomain(domain string) (bool, error) {
	for _, d := range g.domains {
		if d.Domain == domain {
			return true, nil
		}
	}
	return false, nil
}

// GoDaddyDNSRecord represents an individual DNS record
type GoDaddyDNSRecord struct {
	Data string `json:"data"`
	Name string `json:"name"`
	TTL  int    `json:"ttl"`
}

// ReplaceRecordsByType replaces all DNS Records for the specified Domain with
// the specified Type
func ReplaceRecordsByType(client *daddy.Client, domain string, dnstype string, body []GoDaddyDNSRecord) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = client.Put("/v1/domains/"+domain+"/records/"+dnstype, enc)

	return err
}

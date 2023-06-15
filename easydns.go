package main

import (
	"os"

	"github.com/abibby/autodns/easydns"
)

type EasyDNS struct {
	client  *easydns.Client
	domains []string
}

var _ Registrar = (*EasyDNS)(nil)

func (e *EasyDNS) Load() error {
	e.client = easydns.NewClient(os.Getenv("EASYDNS_TOKEN"), os.Getenv("EASYDNS_KEY"), os.Getenv("EASYDNS_ENVIRONMENT"))

	u, err := e.client.User()
	if err != nil {
		return err
	}

	domains, err := e.client.DomainsList(u.Data.User)
	if err != nil {
		return err
	}
	e.domains = make([]string, len(domains.Domains))
	for i, d := range domains.Domains {
		e.domains[i] = d.Name
	}
	return nil
}

func (e *EasyDNS) HasDomain(domain string) (bool, error) {
	for _, d := range e.domains {
		if d == domain {
			return true, nil
		}
	}
	return false, nil
}

func (e *EasyDNS) UpdateDomain(ip, domain string, names []string) error {
	records, err := e.client.ZoneRecordsAll(domain)
	if err != nil {
		return err
	}
	// spew.Dump(records)
	for _, name := range names {
		found := false
		for _, record := range records.Data {
			if record.Host != name || record.Type != "A" {
				continue
			}
			_, err := e.client.ZoneRecords(record.ID, &easydns.ZoneRecordsRequest{
				Domain: domain,
				Host:   name,
				TTL:    3600,
				Type:   "A",
				RData:  ip,
			})
			if err != nil {
				return err
			}
			found = true
		}
		if !found {
			_, err := e.client.ZoneRecordsAdd(domain, "A", &easydns.ZoneRecordsAddRequest{
				Domain: domain,
				Host:   name,
				TTL:    3600,
				Type:   "A",
				RData:  ip,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

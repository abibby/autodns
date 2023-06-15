package main

type Registrar interface {
	Load() error
	HasDomain(domain string) (bool, error)
	UpdateDomain(ip, domain string, names []string) error
}

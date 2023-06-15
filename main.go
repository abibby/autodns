package main

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type IP struct {
	Query string `json:"query"`
}

func main() {
	godotenv.Load("./.env")

	domains := map[string][]string{}

	for _, d := range strings.Split(os.Getenv("DOMAINS"), ",") {
		parts := strings.SplitN(d, ".", 2)
		if len(parts) < 2 {
			log.Printf("invalid domain %s, must contain subdomain", d)
		}
		host, ok := domains[parts[1]]
		if !ok {
			host = make([]string, 0, 1)
		}
		domains[parts[1]] = append(host, parts[0])
	}

	registrars := []Registrar{
		&Godaddy{},
		&EasyDNS{},
	}

	ticker := time.NewTicker(time.Hour)
	ip, err := publicIP()
	if err != nil {
		log.Print(err)
	} else {
		updateDNSRecords(registrars, domains, ip)
	}
	for range ticker.C {
		lastIP := ip
		ip, err = publicIP()
		if err != nil {
			log.Print(err)
			continue
		}

		if ip == lastIP {
			log.Print("no changes")
			continue
		}

		updateDNSRecords(registrars, domains, ip)
	}

}

func updateDNSRecords(registrars []Registrar, domains map[string][]string, ip string) {

	for _, r := range registrars {
		err := r.Load()
		if err != nil {
			log.Print(err)
			continue
		}

		for domain, names := range domains {
			has, err := r.HasDomain(domain)
			if err != nil {
				log.Print(err)
				continue
			}
			if !has {
				continue
			}
			log.Printf("update %s", domain)
			err = r.UpdateDomain(ip, domain, names)
			if err != nil {
				log.Print(err)
			}
		}
	}
}

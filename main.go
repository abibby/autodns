package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/abibby/bob/set"
	"github.com/alyx/go-daddy/daddy"
	"github.com/joho/godotenv"
)

type IP struct {
	Query string `json:"query"`
}

func main() {
	godotenv.Load("./.env")
	domains := set.New[string]()

	for _, d := range strings.Split(os.Getenv("DOMAINS"), ",") {
		domains.Add(strings.TrimSpace(d))
	}
	ip := publicIP()
	err := updateGodaddy(ip, domains)
	if err != nil {
		log.Fatal(err)
	}
}

func publicIP() string {
	req, err := http.Get("http://ip-api.com/json/")
	if err != nil {
		return err.Error()
	}
	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err.Error()
	}

	var ip IP
	json.Unmarshal(body, &ip)

	return ip.Query
}

func updateGodaddy(ip string, domains set.Set[string]) error {
	client, err := daddy.NewClient(os.Getenv("GODADDY_KEY"), os.Getenv("GODADDY_SECRET"), false)
	if err != nil {
		return err
	}

	myDomains, err := client.Domains.List([]string{"ACTIVE"}, nil, 0, "", nil, "")
	if err != nil {
		return err
	}

	for _, value := range myDomains {
		if !domains.Has(value.Domain) {
			continue
		}

		err := ReplaceRecordsByType(client, value.Domain, "A", []DNSRecord{
			{
				Data: ip,
				Name: "@",
				TTL:  3600,
			},
			{
				Data: ip,
				Name: "*",
				TTL:  3600,
			},
		})
		if err != nil {
			return err
		}
		log.Printf("Updated %s", value.Domain)
	}
	return nil
}

// DNSRecord represents an individual DNS record
type DNSRecord struct {
	Data string `json:"data"`
	Name string `json:"name"`
	TTL  int    `json:"ttl"`
}

// ReplaceRecordsByType replaces all DNS Records for the specified Domain with
// the specified Type
func ReplaceRecordsByType(client *daddy.Client, domain string, dnstype string, body []DNSRecord) error {
	enc, err := json.Marshal(body)
	if err != nil {
		return err
	}

	_, err = client.Put("/v1/domains/"+domain+"/records/"+dnstype, enc)

	return err
}

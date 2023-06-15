package easydns

type Record struct {
	ID        string `json:"id"`
	Domain    string `json:"domain"`
	Host      string `json:"host"`
	TTL       string `json:"ttl"`
	Prio      string `json:"prio"`
	Type      string `json:"type"`
	RData     string `json:"rdata"`
	GeoZoneID string `json:"geozone_id"`
	LastMod   string `json:"last_mod"`
}

type ZoneRecordsAllResponse struct {
	Request
	PaginatedRequest
	Data []*Record
}

func (c *Client) ZoneRecordsAll(domain string) (*ZoneRecordsAllResponse, error) {
	resp := &ZoneRecordsAllResponse{}
	err := c.get("/zones/records/all/"+domain, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ZoneRecordsRequest struct {
	Domain    string `json:"domain"`
	Host      string `json:"host"`
	TTL       int    `json:"ttl,omitempty"`
	Prio      int    `json:"prio,omitempty"`
	Type      string `json:"type"`
	RData     string `json:"rdata"`
	GeoZoneID int    `json:"geozone_id,omitempty"`
}
type ZoneRecordsResponse struct{}

func (c *Client) ZoneRecords(id string, req *ZoneRecordsRequest) (*ZoneRecordsResponse, error) {
	resp := &ZoneRecordsResponse{}
	err := c.post("/zones/records/"+id, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

type ZoneRecordsAddRequest ZoneRecordsRequest
type ZoneRecordsAddResponse struct{}

func (c *Client) ZoneRecordsAdd(domain, recordType string, req *ZoneRecordsAddRequest) (*ZoneRecordsAddResponse, error) {
	resp := &ZoneRecordsAddResponse{}
	err := c.put("/zones/records/add/"+domain+"/"+recordType, req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

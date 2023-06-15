package easydns

import (
	"encoding/json"
	"fmt"
)

type Domain struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

type DomainsListResponse struct {
	Request
	PaginatedRequest

	Data map[string]json.RawMessage

	Domains []*Domain
}

func (c *Client) DomainsList(user string) (*DomainsListResponse, error) {
	resp := &DomainsListResponse{}
	err := c.get("/domains/list/"+user, resp)
	if err != nil {
		return nil, err
	}

	resp.Domains = make([]*Domain, 0, resp.Count)
	for i := 0; i < resp.Count; i++ {
		d := &Domain{}
		b, ok := resp.Data[fmt.Sprint(i)]
		if !ok {
			continue
		}
		err = json.Unmarshal(b, d)
		if err != nil {
			return nil, fmt.Errorf("unmarshal domain %d: %w", i, err)
		}
		resp.Domains = append(resp.Domains, d)
	}

	return resp, nil
}

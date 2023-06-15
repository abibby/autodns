package easydns

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	token string
	key   string
	host  string
}

func NewClient(token, key, environment string) *Client {
	host := "https://sandbox.rest.easydns.net"
	if environment == "prod" {
		host = "https://rest.easydns.net"
	}
	return &Client{
		token: token,
		key:   key,
		host:  host,
	}
}

type Request struct {
	Message string `json:"msg"`
	Time    int    `json:"tm"`
	Status  int    `json:"status"`
}
type PaginatedRequest struct {
	Total int `json:"total"`
	Start int `json:"start"`
	Max   int `json:"max"`
	Count int `json:"count"`
}

func (c *Client) request(method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return req, err
	}

	req.SetBasicAuth(c.token, c.key)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	return req, nil
}

func (c *Client) doRequest(method, path string, body, v any) error {
	// fmt.Println(c.host + path)

	var bodyReader io.Reader = http.NoBody
	if body != nil {
		buff := &bytes.Buffer{}
		err := json.NewEncoder(buff).Encode(body)
		if err != nil {
			return err
		}
		bodyReader = buff
	}

	r, err := c.request(method, c.host+path, bodyReader)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("request failed with status %s, %s", resp.Status, b)
	}

	// fmt.Printf("%s\n", b)

	return json.Unmarshal(b, v)
}

func (c *Client) get(path string, v any) error {
	return c.doRequest(http.MethodGet, path, nil, v)
}
func (c *Client) put(path string, body, v any) error {
	return c.doRequest(http.MethodPut, path, body, v)
}
func (c *Client) post(path string, body, v any) error {
	return c.doRequest(http.MethodPost, path, body, v)
}

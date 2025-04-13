package amocrm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	baseURL           = "https://integraciacrm.amocrm.ru"
	endpointCompanies = "/api/v4/companies"
	endpointContacts  = "/api/v4/contacts"
	endpointLink      = "/api/v4/%s/%d/link"
	endpointLeads     = "/api/v4/leads"
	endpointNotes     = "/api/v4/%s/notes"
)

type Client struct {
	client *http.Client
	token  string
}

func NewClient(token string) *Client {
	return &Client{
		client: retryablehttp.NewClient().StandardClient(),
		token:  token,
	}
}

func (c *Client) do(ctx context.Context, endpoint string, request interface{}, response interface{}) error {
	buf := new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(request)
	if err != nil {
		return fmt.Errorf("encode request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+endpoint, buf)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+c.token)

	reqDump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("REQUEST:\n%s", string(reqDump))

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("RESPONSE:\n%s", string(respDump))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status !ok (%d)", resp.StatusCode)
	}

	if response != nil {
		err = json.NewDecoder(resp.Body).Decode(&response)
		if err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

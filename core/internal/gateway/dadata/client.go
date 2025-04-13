package dadata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	client *http.Client
	apiKey string
}

func NewClient(apiKey string) *Client {
	return &Client{
		client: &http.Client{
			Timeout: 1 * time.Minute,
		},
		apiKey: apiKey,
	}
}

func (c *Client) FindByINN(ctx context.Context, inn string) (FindByInnResponse, error) {
	request := request{
		Query: inn,
	}

	buffer := new(bytes.Buffer)
	err := json.NewEncoder(buffer).Encode(request)
	if err != nil {
		return FindByInnResponse{}, fmt.Errorf("encode request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, baseURL+findByIdPartyEndpoint, buffer)
	if err != nil {
		return FindByInnResponse{}, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.apiKey))

	resp, err := c.client.Do(req)
	if err != nil {
		return FindByInnResponse{}, fmt.Errorf("do request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return FindByInnResponse{}, fmt.Errorf("status code not OK")
	}

	var response FindByInnResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return FindByInnResponse{}, fmt.Errorf("decide response: %w", err)
	}

	if len(response.Suggestions) == 0 {
		return FindByInnResponse{}, fmt.Errorf("empty suggestions")
	}

	return response, nil
}

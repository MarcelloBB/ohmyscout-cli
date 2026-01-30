package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client interface {
	Get(ctx context.Context, url string, out any) error
}

type httpClient struct {
	http *http.Client
}

func NewClient() Client {
	return &httpClient{
		http: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *httpClient) Get(ctx context.Context, url string, out any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "application/vnd.github+json")

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("github api error: %s", resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(out)
}

package httpclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type HTTPClient struct {
	client     *http.Client
	retries    int
	retryDelay time.Duration
}

func NewHTTPClient(timeout time.Duration, retries int, retryDelay time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
		retries:    retries,
		retryDelay: retryDelay,
	}
}

func (hc *HTTPClient) Get(url string, target interface{}) error {
	var lastErr error

	for i := 0; i < hc.retries; i++ {
		resp, err := hc.client.Get(url)
		if err != nil {
			lastErr = err
		} else {
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				lastErr = fmt.Errorf("error HTTP: %s", resp.Status)
			} else {
				body, _ := io.ReadAll(resp.Body)
				if err = json.Unmarshal(body, target); err != nil {
					return fmt.Errorf("error unmarshalling response: %w", err)
				}

				return nil
			}
		}

		if i < hc.retries {
			time.Sleep(hc.retryDelay)
		}
	}

	return fmt.Errorf("Fail to GET %s despues de %d intentos: %w", url, hc.retries+1, lastErr)
}

package climacell

import (
	"net/http"
	"time"
)

// BaseURL is the base url of the climacell API
var BaseURL string

func init() {
	BaseURL = "https://api.climacell.co/v3"
}

// Client represents the climacell API client
type Client struct {
	httpClient *http.Client
	baseURL    string
	apiKey     string
}

// NewClient returns a new climacell Client and checks for the
// validity of the provided baseURL
func NewClient(apiKey string, httpClient *http.Client) (*Client, error) {
	client := &Client{}

	if BaseURL == "" {
		return nil, ErrInvalidBaseURL
	}

	if apiKey == "" {
		return nil, ErrInvalidAPIKey
	}

	client.baseURL = BaseURL
	client.apiKey = apiKey

	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 15 * time.Second,
		}
	}

	client.httpClient = httpClient

	return client, nil
}

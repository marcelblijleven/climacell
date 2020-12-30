package climacell

import (
	"fmt"
	"io/ioutil"
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

func (c *Client) makeRequest(endpoint string, parameters map[string]interface{}) (*http.Request, error) {
	u, err := getURL(c.baseURL, endpoint)

	if err != nil {
		return nil, err
	}

	q := u.Query()

	for key, value := range parameters {
		q.Set(key, fmt.Sprintf("%v", value))
	}
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("apikey", c.apiKey)
	return req, nil
}

func (c *Client) doRequest(req *http.Request, endpoint string) ([]byte, error) {
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	err = checkHTTPError(resp, endpoint)

	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return data, nil
}

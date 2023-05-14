package librenms

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// BaseURL - Default librenms base API endpoint
const BaseURL string = "http://localhost:8000/api/v0"

// Client -
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
}

// NewClient -
func NewClient(api_token *string, url *string, insecure *bool) (*Client, error) {

	tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: *insecure},
	}
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: tr},
		// Default librenms URL
		BaseURL: BaseURL,
	}

	if url != nil {
		c.BaseURL = *url
	}

	// If api_token not provided, return empty client
	if api_token == nil {
		return &c, nil
	}

	c.Token = *api_token
	return &c, nil
}

func (c *Client) doRequest(req *http.Request, apiToken *string) ([]byte, error) {
	token := c.Token

	if apiToken != nil {
		token = *apiToken
	}

	req.Header.Set("X-Auth-Token", token)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

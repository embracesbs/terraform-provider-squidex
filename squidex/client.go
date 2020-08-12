package squidex

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
	"time"
)

// Client -
type Client struct {
	HTTPClient  *http.Client
	HostURL     string
	AppName     string
	TokenType   string
	AccessToken string
}

// AuthResponse -
type AuthResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

// NewClient -
func NewClient(host, appName, clientID, clientSecret *string) (*Client, error) {
	client := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
		AppName:    *appName,
	}

	if host != nil {
		client.HostURL = *host
	}

	if (clientID != nil) && (clientSecret != nil) {
		// form request body
		data := url.Values{}
		data.Set("grant_type", "client_credentials")
		data.Set("client_id", *clientID)
		data.Set("client_secret", *clientSecret)
		data.Set("scope", "squidex-api")

		u, err := url.Parse(client.HostURL)
		if err != nil {
			return nil, err
		}
		// boilerplate to create absolute URL
		u.Path = path.Join(u.Path, "/identity-server/connect/token")

		// authenticate
		req, err := http.NewRequest("POST", u.String(), strings.NewReader(data.Encode()))
		if err != nil {
			return nil, err
		}

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		body, err := client.doRequest(req)

		// parse response body
		ar := AuthResponse{}
		err = json.Unmarshal(body, &ar)
		if err != nil {
			return nil, err
		}

		client.TokenType = ar.TokenType
		client.AccessToken = ar.AccessToken
	}

	return &client, nil
}

func (client *Client) doRequest(req *http.Request) ([]byte, error) {
	if !strings.HasSuffix(req.URL.Path, "/identity-server/connect/token") {
		req.Header.Set("Authorization", fmt.Sprintf("%s %s", client.TokenType, client.AccessToken))
	}

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}

package squidex

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Clients -
type Clients struct {
	Items []Client    `json:"items"`
	Links interface{} `json:"_links"`
}

// Client -
type Client struct {
	ID     string      `json:"id"`
	Name   string      `json:"name"`
	Role   string      `json:"role"`
	Secret string      `json:"secret"`
	Links  interface{} `json:"_links"`
}

func getClient(name string, clients []Client) *Client {
	for i := range clients {
		if clients[i].Name == name {
			return &clients[i]
		}
	}

	return nil
}

// GetClientByName -
func (apiClient *APIClient) GetClientByName(name string) (*Client, error) {
	clients, err := apiClient.GetClients()
	if err != nil {
		return nil, err
	}

	return getClient(name, clients), nil
}

// CreateClient -
func (apiClient *APIClient) CreateClient(newClient *Client) (*Client, error) {
	client := map[string]string{"id": newClient.Name}
	rb, err := json.Marshal(client)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(apiClient.HostURL)
	if err != nil {
		return nil, err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", apiClient.AppName, "clients")

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := apiClient.doRequest(req)
	if err != nil {
		return nil, err
	}

	clients := Clients{}
	err = json.Unmarshal(body, &clients)

	if err != nil {
		return nil, err
	}

	return getClient(newClient.Name, clients.Items), nil
}

// UpdateClient -
func (apiClient *APIClient) UpdateClient(updateClient *Client) (*Client, error) {
	client := map[string]string{"name": updateClient.Name, "role": updateClient.Role}
	rb, err := json.Marshal(client)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(apiClient.HostURL)
	if err != nil {
		return nil, err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", apiClient.AppName, "clients", updateClient.Name)

	req, err := http.NewRequest("PUT", u.String(), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := apiClient.doRequest(req)
	if err != nil {
		return nil, err
	}

	clients := Clients{}
	err = json.Unmarshal(body, &clients)

	if err != nil {
		return nil, err
	}

	return getClient(updateClient.Name, clients.Items), nil
}

// DeleteClient -
func (apiClient *APIClient) DeleteClient(id string) error {
	u, err := url.Parse(apiClient.HostURL)
	if err != nil {
		return err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", apiClient.AppName, "clients", id)

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}

	body, err := apiClient.doRequest(req)
	if err != nil {
		return err
	}

	clients := Clients{}
	err = json.Unmarshal(body, &clients)
	if getClient(id, clients.Items) != nil {
		return errors.New(string(body))
	}

	return nil
}

// GetClients -
func (apiClient *APIClient) GetClients() ([]Client, error) {
	u, err := url.Parse(apiClient.HostURL)
	if err != nil {
		return nil, err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", apiClient.AppName, "clients")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	body, err := apiClient.doRequest(req)
	if err != nil {
		return nil, err
	}

	clients := Clients{}
	err = json.Unmarshal(body, &clients)
	if err != nil {
		return nil, err
	}

	return clients.Items, nil
}

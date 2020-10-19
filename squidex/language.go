package squidex

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"path"
	"strings"
)

// Languages -
type Languages struct {
	Items []Language  `json:"items"`
	Links interface{} `json:"_links"`
}

// Language -
type Language struct {
	Iso2Code    string      `json:"iso2Code"`
	EnglishName string      `json:"englishName"`
	Fallback    []string    `json:"fallback"`
	IsMaster    bool        `json:"isMaster"`
	IsOptional  bool        `json:"isOptional"`
	Links       interface{} `json:"_links"`
}

func getLanguage(iso2Code string, languages []Language) *Language {
	for i := range languages {
		if languages[i].Iso2Code == iso2Code {
			return &languages[i]
		}
	}

	return nil
}

// GetLanguage -
func (client *APIClient) GetLanguage(iso2Code string) (*Language, error) {
	languages, err := client.GetLanguages()
	if err != nil {
		return nil, err
	}

	return getLanguage(iso2Code, languages), nil
}

// CreateLanguage -
func (client *APIClient) CreateLanguage(newLanguage *Language) (*Language, error) {
	language := map[string]string{"language": newLanguage.Iso2Code}
	rb, err := json.Marshal(language)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(client.HostURL)
	if err != nil {
		return nil, err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", client.AppName, "languages")

	req, err := http.NewRequest("POST", u.String(), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	languages := Languages{}
	err = json.Unmarshal(body, &languages)

	if err != nil {
		return nil, err
	}

	return getLanguage(newLanguage.Iso2Code, languages.Items), nil
}

// UpdateLanguage -
func (client *APIClient) UpdateLanguage(updateLanguage *Language) (*Language, error) {
	language := map[string]bool{"isMaster": updateLanguage.IsMaster, "isOptional": false}
	rb, err := json.Marshal(language)
	if err != nil {
		return nil, err
	}

	u, err := url.Parse(client.HostURL)
	if err != nil {
		return nil, err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", client.AppName, "languages", updateLanguage.Iso2Code)

	req, err := http.NewRequest("PUT", u.String(), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	body, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	languages := Languages{}
	err = json.Unmarshal(body, &languages)

	if err != nil {
		return nil, err
	}

	return getLanguage(updateLanguage.Iso2Code, languages.Items), nil
}

// DeleteLanguage -
func (client *APIClient) DeleteLanguage(iso2Code string) error {
	u, err := url.Parse(client.HostURL)
	if err != nil {
		return err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", client.AppName, "languages", iso2Code)

	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return err
	}

	body, err := client.doRequest(req)
	if err != nil {
		return err
	}

	languages := Languages{}
	err = json.Unmarshal(body, &languages)
	if getLanguage(iso2Code, languages.Items) != nil {
		return errors.New(string(body))
	}

	return nil
}

// GetLanguages -
func (client *APIClient) GetLanguages() ([]Language, error) {
	u, err := url.Parse(client.HostURL)
	if err != nil {
		return nil, err
	}
	// boilerplate to create absolute URL
	u.Path = path.Join(u.Path, "/api/apps", client.AppName, "languages")

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	body, err := client.doRequest(req)
	if err != nil {
		return nil, err
	}

	languages := Languages{}
	err = json.Unmarshal(body, &languages)
	if err != nil {
		return nil, err
	}

	return languages.Items, nil
}

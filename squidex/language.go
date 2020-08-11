package squidex

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func getLanguage(iso2Code string, languages []Language) *Language {
	for i := range languages {
		if languages[i].Iso2Code == iso2Code {
			log.Println("Kreirani ", languages[i])
			return &languages[i]
		}
	}

	return nil
}

// GetLanguage -
func (client *Client) GetLanguage(iso2Code string) (*Language, error) {
	languages, err := client.GetLanguages()
	if err != nil {
		return nil, err
	}

	return getLanguage(iso2Code, languages), nil
}

// CreateLanguage -
func (client *Client) CreateLanguage(newLanguage *Language) (*Language, error) {
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
		log.Println("Greska 1", err)
		return nil, err
	}

	languages := Languages{}
	err = json.Unmarshal(body, &languages)

	if err != nil {
		log.Println("Greska 2", err)
		return nil, err
	}

	return getLanguage(newLanguage.Iso2Code, languages.Items), nil
}

// DeleteLanguage -
func (client *Client) DeleteLanguage(iso2Code string) error {
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
func (client *Client) GetLanguages() ([]Language, error) {
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

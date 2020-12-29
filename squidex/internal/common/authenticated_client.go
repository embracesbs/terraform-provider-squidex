package common

import (
	"net/http"

	"golang.org/x/net/context"
	cc "golang.org/x/oauth2/clientcredentials"
)

//This client will refresh token in background
func NewClient(clientId, clientSecret, tokenUrl, scope string) *http.Client {

	// this should match whatever service has given you
	// client credential access
	config := &cc.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     tokenUrl,
		Scopes:       []string{scope},
	}

	// you can modify the client (for example ignoring bad certs or otherwise)
	// by modifying the context
	ctx := context.Background()
	client := config.Client(ctx)

	return client
}

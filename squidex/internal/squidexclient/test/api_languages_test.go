/*
Squidex API

Testing LanguagesApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package squidexclient

import (
	"context"
	"testing"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_openapi_LanguagesApiService(t *testing.T) {

	configuration := squidexclient.NewConfiguration()
	apiClient := squidexclient.NewAPIClient(configuration)

	t.Run("Test LanguagesApiService LanguagesGetLanguages", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.LanguagesApi.LanguagesGetLanguages(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

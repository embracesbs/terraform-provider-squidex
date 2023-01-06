/*
Squidex API

Testing TemplatesApiService

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

func Test_openapi_TemplatesApiService(t *testing.T) {

	configuration := squidexclient.NewConfiguration()
	apiClient := squidexclient.NewAPIClient(configuration)

	t.Run("Test TemplatesApiService TemplatesGetTemplate", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var name string

		resp, httpRes, err := apiClient.TemplatesApi.TemplatesGetTemplate(context.Background(), name).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TemplatesApiService TemplatesGetTemplates", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.TemplatesApi.TemplatesGetTemplates(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

/*
Squidex API

Testing DiagnosticsApiService

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

func Test_openapi_DiagnosticsApiService(t *testing.T) {

	configuration := squidexclient.NewConfiguration()
	apiClient := squidexclient.NewAPIClient(configuration)

	t.Run("Test DiagnosticsApiService DiagnosticsGetDump", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DiagnosticsApi.DiagnosticsGetDump(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DiagnosticsApiService DiagnosticsGetGCDump", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		httpRes, err := apiClient.DiagnosticsApi.DiagnosticsGetGCDump(context.Background()).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}

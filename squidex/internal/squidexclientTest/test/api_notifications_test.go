/*
Squidex API

Testing NotificationsApiService

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

func Test_openapi_NotificationsApiService(t *testing.T) {

	configuration := squidexclient.NewConfiguration()
	apiClient := squidexclient.NewAPIClient(configuration)

	t.Run("Test NotificationsApiService UserNotificationsDeleteComment", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string
		var commentId string

		httpRes, err := apiClient.NotificationsApi.UserNotificationsDeleteComment(context.Background(), userId, commentId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationsApiService UserNotificationsGetNotifications", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId string

		resp, httpRes, err := apiClient.NotificationsApi.UserNotificationsGetNotifications(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
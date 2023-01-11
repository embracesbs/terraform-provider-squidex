package common

import (
	// "github.com/hashicorp/errwrap"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/embracesbs/terraform-provider-squidex/squidex/internal/squidexclient"
)

// APIError is the return in http response body of Squidex Api.
type APIError struct {
	Message   string   `json:"message"`
	TraceID   *string  `json:"traceId,omitempty"`
	ErrorType *string  `json:"type,omitempty"`
	Details   []string `json:"details,omitempty"`
	Code      int      `json:"statusCode,omitempty"`
}

func (e APIError) Error() string {
	val := strconv.Itoa(e.Code) + "\n" + e.Message + "\n" + strings.Join(e.Details, "\n")
	return val
}

func HandleAPIError(response *http.Response, err interface{}, deleteFunction bool) error {

	if response.StatusCode == http.StatusOK ||
		response.StatusCode == http.StatusCreated ||
		response.StatusCode == http.StatusNoContent {
		// do nothing
		return nil
	}

	if deleteFunction && response.StatusCode == http.StatusNotFound {
		return nil
	}

	if response.StatusCode == http.StatusUnauthorized || response.StatusCode == http.StatusForbidden {
		// do something with status
		return &APIError{
			Message: response.Status + " - " + response.Request.URL.String(),
			Code:    response.StatusCode,
			Details: make([]string, 0),
		}
	}

	genericMessage := fmt.Sprintf(
		"error sending %s request to %s: %s.",
		response.Request.Method,
		response.Request.URL.Path,
		response.Status)

	if err != nil {
		apiError := err.(*squidexclient.GenericOpenAPIError)

		responseBody := apiError.Body()

		if len(responseBody) != 0 && response.StatusCode >= 400 {
			var apierror APIError
			err := json.Unmarshal(responseBody, &apierror)
			if err != nil {
				// custom create error
				log.Printf("[DEBUG] Response body: %s", responseBody)
				val := string(responseBody)
				return &APIError{
					Message: genericMessage,
					Code:    response.StatusCode,
					Details: []string{val},
				}
			}
			// mapped response to error
			apierror.Message = genericMessage + "\n" + apierror.Message
			return apierror
		}
	}

	// something else went wrong, generic http error
	return &APIError{
		Message: genericMessage,
		Code:    response.StatusCode,
		Details: make([]string, 0),
	}

}

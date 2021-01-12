package common

import (
	// "github.com/hashicorp/errwrap"
	"net/http"
	"io/ioutil"
	"log"
)

// APIError is the return in http response body of Squidex Api.
type APIError struct {
	Message string `json:"message"`
	TraceID *string `json:"traceId,omitempty"`
	ErrorType *string `json:"type,omitempty"`
	Details *[]string `json:"details,omitempty"`
	StatusCode *int `json:"statusCode,omitempty"`
}

func (e *APIError) Error() string {
	return e.Message
}

// HandleAPIError handles http responses from Squidex Api, 
// ignores status 200 OK all other responses will panic() and output the response body
func HandleAPIError(resp *http.Response) {
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		// do nothing
		return
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	log.Panicf("Bad Api Response:\n%s", bodyString)
}

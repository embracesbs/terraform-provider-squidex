/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

import (
	"time"
)

// HistoryEventDto struct for HistoryEventDto
type HistoryEventDto struct {
	// The message for the event.
	Message string `json:"message"`
	// The type of the original event.
	EventType string `json:"eventType"`
	// The user who called the action.
	Actor string `json:"actor"`
	// Gets a unique id for the event.
	EventId string `json:"eventId,omitempty"`
	// The time when the event happened.
	Created time.Time `json:"created,omitempty"`
	// The version identifier.
	Version int64 `json:"version,omitempty"`
}

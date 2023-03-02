/*
 * Squidex API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 7.0.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package squidexclient

// SkipReason the model 'SkipReason'
type SkipReason string

// List of SkipReason
const (
	SkipReasonNONE                              SkipReason = "None"
	SkipReasonCONDITION_DOES_NOT_MATCH          SkipReason = "ConditionDoesNotMatch"
	SkipReasonCONDITION_PRECHECK_DOES_NOT_MATCH SkipReason = "ConditionPrecheckDoesNotMatch"
	SkipReasonDISABLED                          SkipReason = "Disabled"
	SkipReasonFAILED                            SkipReason = "Failed"
	SkipReasonFROM_RULE                         SkipReason = "FromRule"
	SkipReasonNO_ACTION                         SkipReason = "NoAction"
	SkipReasonNO_TRIGGER                        SkipReason = "NoTrigger"
	SkipReasonTOO_OLD                           SkipReason = "TooOld"
	SkipReasonWRONG_EVENT                       SkipReason = "WrongEvent"
	SkipReasonWRONG_EVENT_FOR_TRIGGER           SkipReason = "WrongEventForTrigger"
)
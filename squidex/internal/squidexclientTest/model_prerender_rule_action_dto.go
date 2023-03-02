/*
Squidex API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 7.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package squidexclient

import (
	"encoding/json"
)

// PrerenderRuleActionDto struct for PrerenderRuleActionDto
type PrerenderRuleActionDto struct {
	RuleAction
	// The prerender token from your account.
	Token string `json:"token"`
	// The url to recache.
	Url string `json:"url"`
}

// NewPrerenderRuleActionDto instantiates a new PrerenderRuleActionDto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrerenderRuleActionDto(token string, url string, actionType NullableString) *PrerenderRuleActionDto {
	this := PrerenderRuleActionDto{}
	this.ActionType = actionType
	this.Token = token
	this.Url = url
	return &this
}

// NewPrerenderRuleActionDtoWithDefaults instantiates a new PrerenderRuleActionDto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrerenderRuleActionDtoWithDefaults() *PrerenderRuleActionDto {
	this := PrerenderRuleActionDto{}
	return &this
}

// GetToken returns the Token field value
func (o *PrerenderRuleActionDto) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *PrerenderRuleActionDto) GetTokenOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *PrerenderRuleActionDto) SetToken(v string) {
	o.Token = v
}

// GetUrl returns the Url field value
func (o *PrerenderRuleActionDto) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *PrerenderRuleActionDto) GetUrlOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *PrerenderRuleActionDto) SetUrl(v string) {
	o.Url = v
}

func (o PrerenderRuleActionDto) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedRuleAction, errRuleAction := json.Marshal(o.RuleAction)
	if errRuleAction != nil {
		return []byte{}, errRuleAction
	}
	errRuleAction = json.Unmarshal([]byte(serializedRuleAction), &toSerialize)
	if errRuleAction != nil {
		return []byte{}, errRuleAction
	}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullablePrerenderRuleActionDto struct {
	value *PrerenderRuleActionDto
	isSet bool
}

func (v NullablePrerenderRuleActionDto) Get() *PrerenderRuleActionDto {
	return v.value
}

func (v *NullablePrerenderRuleActionDto) Set(val *PrerenderRuleActionDto) {
	v.value = val
	v.isSet = true
}

func (v NullablePrerenderRuleActionDto) IsSet() bool {
	return v.isSet
}

func (v *NullablePrerenderRuleActionDto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrerenderRuleActionDto(val *PrerenderRuleActionDto) *NullablePrerenderRuleActionDto {
	return &NullablePrerenderRuleActionDto{value: val, isSet: true}
}

func (v NullablePrerenderRuleActionDto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrerenderRuleActionDto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


# WebhookRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | The url to the webhook. | 
**Method** | [**WebhookMethod**](WebhookMethod.md) |  | 
**Payload** | Pointer to **NullableString** | Leave it empty to use the full event as body. | [optional] 
**PayloadType** | Pointer to **NullableString** | The mime type of the payload. | [optional] 
**Headers** | Pointer to **NullableString** | The message headers in the format &#39;[Key]&#x3D;[Value]&#39;, one entry per line. | [optional] 
**SharedSecret** | Pointer to **NullableString** | The shared secret that is used to calculate the payload signature. | [optional] 

## Methods

### NewWebhookRuleActionDto

`func NewWebhookRuleActionDto(url string, method WebhookMethod, ) *WebhookRuleActionDto`

NewWebhookRuleActionDto instantiates a new WebhookRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookRuleActionDtoWithDefaults

`func NewWebhookRuleActionDtoWithDefaults() *WebhookRuleActionDto`

NewWebhookRuleActionDtoWithDefaults instantiates a new WebhookRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *WebhookRuleActionDto) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebhookRuleActionDto) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebhookRuleActionDto) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *WebhookRuleActionDto) GetMethod() WebhookMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *WebhookRuleActionDto) GetMethodOk() (*WebhookMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *WebhookRuleActionDto) SetMethod(v WebhookMethod)`

SetMethod sets Method field to given value.


### GetPayload

`func (o *WebhookRuleActionDto) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *WebhookRuleActionDto) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *WebhookRuleActionDto) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *WebhookRuleActionDto) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### SetPayloadNil

`func (o *WebhookRuleActionDto) SetPayloadNil(b bool)`

 SetPayloadNil sets the value for Payload to be an explicit nil

### UnsetPayload
`func (o *WebhookRuleActionDto) UnsetPayload()`

UnsetPayload ensures that no value is present for Payload, not even an explicit nil
### GetPayloadType

`func (o *WebhookRuleActionDto) GetPayloadType() string`

GetPayloadType returns the PayloadType field if non-nil, zero value otherwise.

### GetPayloadTypeOk

`func (o *WebhookRuleActionDto) GetPayloadTypeOk() (*string, bool)`

GetPayloadTypeOk returns a tuple with the PayloadType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadType

`func (o *WebhookRuleActionDto) SetPayloadType(v string)`

SetPayloadType sets PayloadType field to given value.

### HasPayloadType

`func (o *WebhookRuleActionDto) HasPayloadType() bool`

HasPayloadType returns a boolean if a field has been set.

### SetPayloadTypeNil

`func (o *WebhookRuleActionDto) SetPayloadTypeNil(b bool)`

 SetPayloadTypeNil sets the value for PayloadType to be an explicit nil

### UnsetPayloadType
`func (o *WebhookRuleActionDto) UnsetPayloadType()`

UnsetPayloadType ensures that no value is present for PayloadType, not even an explicit nil
### GetHeaders

`func (o *WebhookRuleActionDto) GetHeaders() string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *WebhookRuleActionDto) GetHeadersOk() (*string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *WebhookRuleActionDto) SetHeaders(v string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *WebhookRuleActionDto) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### SetHeadersNil

`func (o *WebhookRuleActionDto) SetHeadersNil(b bool)`

 SetHeadersNil sets the value for Headers to be an explicit nil

### UnsetHeaders
`func (o *WebhookRuleActionDto) UnsetHeaders()`

UnsetHeaders ensures that no value is present for Headers, not even an explicit nil
### GetSharedSecret

`func (o *WebhookRuleActionDto) GetSharedSecret() string`

GetSharedSecret returns the SharedSecret field if non-nil, zero value otherwise.

### GetSharedSecretOk

`func (o *WebhookRuleActionDto) GetSharedSecretOk() (*string, bool)`

GetSharedSecretOk returns a tuple with the SharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedSecret

`func (o *WebhookRuleActionDto) SetSharedSecret(v string)`

SetSharedSecret sets SharedSecret field to given value.

### HasSharedSecret

`func (o *WebhookRuleActionDto) HasSharedSecret() bool`

HasSharedSecret returns a boolean if a field has been set.

### SetSharedSecretNil

`func (o *WebhookRuleActionDto) SetSharedSecretNil(b bool)`

 SetSharedSecretNil sets the value for SharedSecret to be an explicit nil

### UnsetSharedSecret
`func (o *WebhookRuleActionDto) UnsetSharedSecret()`

UnsetSharedSecret ensures that no value is present for SharedSecret, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



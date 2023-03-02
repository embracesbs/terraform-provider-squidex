# ElasticSearchRuleActionDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The url to the elastic search instance or cluster. | 
**IndexName** | **string** | The name of the index. | 
**Username** | Pointer to **NullableString** | The optional username. | [optional] 
**Password** | Pointer to **NullableString** | The optional password. | [optional] 
**Document** | Pointer to **NullableString** | The optional custom document. | [optional] 
**Delete** | Pointer to **NullableString** | The condition when to delete the document. | [optional] 

## Methods

### NewElasticSearchRuleActionDtoAllOf

`func NewElasticSearchRuleActionDtoAllOf(host string, indexName string, ) *ElasticSearchRuleActionDtoAllOf`

NewElasticSearchRuleActionDtoAllOf instantiates a new ElasticSearchRuleActionDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticSearchRuleActionDtoAllOfWithDefaults

`func NewElasticSearchRuleActionDtoAllOfWithDefaults() *ElasticSearchRuleActionDtoAllOf`

NewElasticSearchRuleActionDtoAllOfWithDefaults instantiates a new ElasticSearchRuleActionDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ElasticSearchRuleActionDtoAllOf) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ElasticSearchRuleActionDtoAllOf) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ElasticSearchRuleActionDtoAllOf) SetHost(v string)`

SetHost sets Host field to given value.


### GetIndexName

`func (o *ElasticSearchRuleActionDtoAllOf) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *ElasticSearchRuleActionDtoAllOf) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *ElasticSearchRuleActionDtoAllOf) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetUsername

`func (o *ElasticSearchRuleActionDtoAllOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ElasticSearchRuleActionDtoAllOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ElasticSearchRuleActionDtoAllOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ElasticSearchRuleActionDtoAllOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ElasticSearchRuleActionDtoAllOf) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ElasticSearchRuleActionDtoAllOf) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ElasticSearchRuleActionDtoAllOf) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ElasticSearchRuleActionDtoAllOf) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ElasticSearchRuleActionDtoAllOf) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ElasticSearchRuleActionDtoAllOf) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ElasticSearchRuleActionDtoAllOf) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ElasticSearchRuleActionDtoAllOf) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDocument

`func (o *ElasticSearchRuleActionDtoAllOf) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ElasticSearchRuleActionDtoAllOf) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ElasticSearchRuleActionDtoAllOf) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ElasticSearchRuleActionDtoAllOf) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *ElasticSearchRuleActionDtoAllOf) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *ElasticSearchRuleActionDtoAllOf) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetDelete

`func (o *ElasticSearchRuleActionDtoAllOf) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ElasticSearchRuleActionDtoAllOf) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ElasticSearchRuleActionDtoAllOf) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ElasticSearchRuleActionDtoAllOf) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *ElasticSearchRuleActionDtoAllOf) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *ElasticSearchRuleActionDtoAllOf) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



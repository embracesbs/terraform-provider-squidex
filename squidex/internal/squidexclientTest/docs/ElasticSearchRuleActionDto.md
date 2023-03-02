# ElasticSearchRuleActionDto

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

### NewElasticSearchRuleActionDto

`func NewElasticSearchRuleActionDto(host string, indexName string, ) *ElasticSearchRuleActionDto`

NewElasticSearchRuleActionDto instantiates a new ElasticSearchRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticSearchRuleActionDtoWithDefaults

`func NewElasticSearchRuleActionDtoWithDefaults() *ElasticSearchRuleActionDto`

NewElasticSearchRuleActionDtoWithDefaults instantiates a new ElasticSearchRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *ElasticSearchRuleActionDto) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ElasticSearchRuleActionDto) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ElasticSearchRuleActionDto) SetHost(v string)`

SetHost sets Host field to given value.


### GetIndexName

`func (o *ElasticSearchRuleActionDto) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *ElasticSearchRuleActionDto) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *ElasticSearchRuleActionDto) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetUsername

`func (o *ElasticSearchRuleActionDto) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ElasticSearchRuleActionDto) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ElasticSearchRuleActionDto) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ElasticSearchRuleActionDto) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ElasticSearchRuleActionDto) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ElasticSearchRuleActionDto) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *ElasticSearchRuleActionDto) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ElasticSearchRuleActionDto) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ElasticSearchRuleActionDto) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ElasticSearchRuleActionDto) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *ElasticSearchRuleActionDto) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ElasticSearchRuleActionDto) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetDocument

`func (o *ElasticSearchRuleActionDto) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *ElasticSearchRuleActionDto) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *ElasticSearchRuleActionDto) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *ElasticSearchRuleActionDto) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *ElasticSearchRuleActionDto) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *ElasticSearchRuleActionDto) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetDelete

`func (o *ElasticSearchRuleActionDto) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *ElasticSearchRuleActionDto) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *ElasticSearchRuleActionDto) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *ElasticSearchRuleActionDto) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *ElasticSearchRuleActionDto) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *ElasticSearchRuleActionDto) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



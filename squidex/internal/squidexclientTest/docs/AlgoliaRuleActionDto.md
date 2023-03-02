# AlgoliaRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppId** | **string** | The application ID. | 
**ApiKey** | **string** | The API key to grant access to Squidex. | 
**IndexName** | **string** | The name of the index. | 
**Document** | Pointer to **NullableString** | The optional custom document. | [optional] 
**Delete** | Pointer to **NullableString** | The condition when to delete the entry. | [optional] 

## Methods

### NewAlgoliaRuleActionDto

`func NewAlgoliaRuleActionDto(appId string, apiKey string, indexName string, ) *AlgoliaRuleActionDto`

NewAlgoliaRuleActionDto instantiates a new AlgoliaRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlgoliaRuleActionDtoWithDefaults

`func NewAlgoliaRuleActionDtoWithDefaults() *AlgoliaRuleActionDto`

NewAlgoliaRuleActionDtoWithDefaults instantiates a new AlgoliaRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppId

`func (o *AlgoliaRuleActionDto) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *AlgoliaRuleActionDto) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *AlgoliaRuleActionDto) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetApiKey

`func (o *AlgoliaRuleActionDto) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AlgoliaRuleActionDto) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AlgoliaRuleActionDto) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetIndexName

`func (o *AlgoliaRuleActionDto) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *AlgoliaRuleActionDto) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *AlgoliaRuleActionDto) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetDocument

`func (o *AlgoliaRuleActionDto) GetDocument() string`

GetDocument returns the Document field if non-nil, zero value otherwise.

### GetDocumentOk

`func (o *AlgoliaRuleActionDto) GetDocumentOk() (*string, bool)`

GetDocumentOk returns a tuple with the Document field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocument

`func (o *AlgoliaRuleActionDto) SetDocument(v string)`

SetDocument sets Document field to given value.

### HasDocument

`func (o *AlgoliaRuleActionDto) HasDocument() bool`

HasDocument returns a boolean if a field has been set.

### SetDocumentNil

`func (o *AlgoliaRuleActionDto) SetDocumentNil(b bool)`

 SetDocumentNil sets the value for Document to be an explicit nil

### UnsetDocument
`func (o *AlgoliaRuleActionDto) UnsetDocument()`

UnsetDocument ensures that no value is present for Document, not even an explicit nil
### GetDelete

`func (o *AlgoliaRuleActionDto) GetDelete() string`

GetDelete returns the Delete field if non-nil, zero value otherwise.

### GetDeleteOk

`func (o *AlgoliaRuleActionDto) GetDeleteOk() (*string, bool)`

GetDeleteOk returns a tuple with the Delete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelete

`func (o *AlgoliaRuleActionDto) SetDelete(v string)`

SetDelete sets Delete field to given value.

### HasDelete

`func (o *AlgoliaRuleActionDto) HasDelete() bool`

HasDelete returns a boolean if a field has been set.

### SetDeleteNil

`func (o *AlgoliaRuleActionDto) SetDeleteNil(b bool)`

 SetDeleteNil sets the value for Delete to be an explicit nil

### UnsetDelete
`func (o *AlgoliaRuleActionDto) UnsetDelete()`

UnsetDelete ensures that no value is present for Delete, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CreateContentRuleActionDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The content data. | 
**Schema** | **string** | The name of the schema. | 
**Client** | Pointer to **NullableString** | An optional client name. | [optional] 
**Publish** | Pointer to **bool** | Publish the content. | [optional] 

## Methods

### NewCreateContentRuleActionDto

`func NewCreateContentRuleActionDto(data string, schema string, ) *CreateContentRuleActionDto`

NewCreateContentRuleActionDto instantiates a new CreateContentRuleActionDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateContentRuleActionDtoWithDefaults

`func NewCreateContentRuleActionDtoWithDefaults() *CreateContentRuleActionDto`

NewCreateContentRuleActionDtoWithDefaults instantiates a new CreateContentRuleActionDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CreateContentRuleActionDto) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CreateContentRuleActionDto) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CreateContentRuleActionDto) SetData(v string)`

SetData sets Data field to given value.


### GetSchema

`func (o *CreateContentRuleActionDto) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *CreateContentRuleActionDto) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *CreateContentRuleActionDto) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetClient

`func (o *CreateContentRuleActionDto) GetClient() string`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *CreateContentRuleActionDto) GetClientOk() (*string, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *CreateContentRuleActionDto) SetClient(v string)`

SetClient sets Client field to given value.

### HasClient

`func (o *CreateContentRuleActionDto) HasClient() bool`

HasClient returns a boolean if a field has been set.

### SetClientNil

`func (o *CreateContentRuleActionDto) SetClientNil(b bool)`

 SetClientNil sets the value for Client to be an explicit nil

### UnsetClient
`func (o *CreateContentRuleActionDto) UnsetClient()`

UnsetClient ensures that no value is present for Client, not even an explicit nil
### GetPublish

`func (o *CreateContentRuleActionDto) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *CreateContentRuleActionDto) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *CreateContentRuleActionDto) SetPublish(v bool)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *CreateContentRuleActionDto) HasPublish() bool`

HasPublish returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



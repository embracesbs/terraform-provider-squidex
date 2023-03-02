# CreateSchemaDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the schema. | 
**Type** | Pointer to [**SchemaType**](SchemaType.md) |  | [optional] 
**IsSingleton** | Pointer to **bool** | Set to true to allow a single content item only. | [optional] 

## Methods

### NewCreateSchemaDtoAllOf

`func NewCreateSchemaDtoAllOf(name string, ) *CreateSchemaDtoAllOf`

NewCreateSchemaDtoAllOf instantiates a new CreateSchemaDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSchemaDtoAllOfWithDefaults

`func NewCreateSchemaDtoAllOfWithDefaults() *CreateSchemaDtoAllOf`

NewCreateSchemaDtoAllOfWithDefaults instantiates a new CreateSchemaDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateSchemaDtoAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSchemaDtoAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSchemaDtoAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CreateSchemaDtoAllOf) GetType() SchemaType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSchemaDtoAllOf) GetTypeOk() (*SchemaType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSchemaDtoAllOf) SetType(v SchemaType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSchemaDtoAllOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIsSingleton

`func (o *CreateSchemaDtoAllOf) GetIsSingleton() bool`

GetIsSingleton returns the IsSingleton field if non-nil, zero value otherwise.

### GetIsSingletonOk

`func (o *CreateSchemaDtoAllOf) GetIsSingletonOk() (*bool, bool)`

GetIsSingletonOk returns a tuple with the IsSingleton field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSingleton

`func (o *CreateSchemaDtoAllOf) SetIsSingleton(v bool)`

SetIsSingleton sets IsSingleton field to given value.

### HasIsSingleton

`func (o *CreateSchemaDtoAllOf) HasIsSingleton() bool`

HasIsSingleton returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



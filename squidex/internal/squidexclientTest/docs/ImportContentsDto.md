# ImportContentsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datas** | [**[]map[string]map[string]interface{}**](map[string]map[string]interface{}.md) | The data to import. | 
**Publish** | Pointer to **bool** | True to automatically publish the content. | [optional] 
**DoNotScript** | Pointer to **bool** | True to turn off scripting for faster inserts. Default: true. | [optional] 
**OptimizeValidation** | Pointer to **bool** | True to turn off costly validation: Unique checks, asset checks and reference checks. Default: true. | [optional] 

## Methods

### NewImportContentsDto

`func NewImportContentsDto(datas []map[string]map[string]interface{}, ) *ImportContentsDto`

NewImportContentsDto instantiates a new ImportContentsDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImportContentsDtoWithDefaults

`func NewImportContentsDtoWithDefaults() *ImportContentsDto`

NewImportContentsDtoWithDefaults instantiates a new ImportContentsDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatas

`func (o *ImportContentsDto) GetDatas() []map[string]map[string]interface{}`

GetDatas returns the Datas field if non-nil, zero value otherwise.

### GetDatasOk

`func (o *ImportContentsDto) GetDatasOk() (*[]map[string]map[string]interface{}, bool)`

GetDatasOk returns a tuple with the Datas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatas

`func (o *ImportContentsDto) SetDatas(v []map[string]map[string]interface{})`

SetDatas sets Datas field to given value.


### GetPublish

`func (o *ImportContentsDto) GetPublish() bool`

GetPublish returns the Publish field if non-nil, zero value otherwise.

### GetPublishOk

`func (o *ImportContentsDto) GetPublishOk() (*bool, bool)`

GetPublishOk returns a tuple with the Publish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublish

`func (o *ImportContentsDto) SetPublish(v bool)`

SetPublish sets Publish field to given value.

### HasPublish

`func (o *ImportContentsDto) HasPublish() bool`

HasPublish returns a boolean if a field has been set.

### GetDoNotScript

`func (o *ImportContentsDto) GetDoNotScript() bool`

GetDoNotScript returns the DoNotScript field if non-nil, zero value otherwise.

### GetDoNotScriptOk

`func (o *ImportContentsDto) GetDoNotScriptOk() (*bool, bool)`

GetDoNotScriptOk returns a tuple with the DoNotScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotScript

`func (o *ImportContentsDto) SetDoNotScript(v bool)`

SetDoNotScript sets DoNotScript field to given value.

### HasDoNotScript

`func (o *ImportContentsDto) HasDoNotScript() bool`

HasDoNotScript returns a boolean if a field has been set.

### GetOptimizeValidation

`func (o *ImportContentsDto) GetOptimizeValidation() bool`

GetOptimizeValidation returns the OptimizeValidation field if non-nil, zero value otherwise.

### GetOptimizeValidationOk

`func (o *ImportContentsDto) GetOptimizeValidationOk() (*bool, bool)`

GetOptimizeValidationOk returns a tuple with the OptimizeValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptimizeValidation

`func (o *ImportContentsDto) SetOptimizeValidation(v bool)`

SetOptimizeValidation sets OptimizeValidation field to given value.

### HasOptimizeValidation

`func (o *ImportContentsDto) HasOptimizeValidation() bool`

HasOptimizeValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



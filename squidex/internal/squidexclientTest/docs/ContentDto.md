# ContentDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Id** | Pointer to **string** | The if of the content item. | [optional] 
**CreatedBy** | **string** | The user that has created the content item. | 
**LastModifiedBy** | **string** | The user that has updated the content item. | 
**Data** | **interface{}** | The data of the content item. | 
**ReferenceData** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Created** | Pointer to **time.Time** | The date and time when the content item has been created. | [optional] 
**LastModified** | Pointer to **time.Time** | The date and time when the content item has been modified last. | [optional] 
**Status** | Pointer to **string** | The status of the content. | [optional] 
**NewStatus** | Pointer to **NullableString** | The new status of the content. | [optional] 
**StatusColor** | Pointer to **string** | The color of the status. | [optional] 
**NewStatusColor** | Pointer to **NullableString** | The color of the new status. | [optional] 
**EditToken** | Pointer to **NullableString** | The UI token. | [optional] 
**ScheduleJob** | Pointer to [**ScheduleJobDto**](ScheduleJobDto.md) |  | [optional] 
**SchemaId** | Pointer to **string** | The id of the schema. | [optional] 
**SchemaName** | Pointer to **NullableString** | The name of the schema. | [optional] 
**SchemaDisplayName** | Pointer to **NullableString** | The display name of the schema. | [optional] 
**ReferenceFields** | Pointer to [**[]FieldDto**](FieldDto.md) | The reference fields. | [optional] 
**IsDeleted** | Pointer to **bool** | Indicates whether the content is deleted. | [optional] 
**Version** | Pointer to **int64** | The version of the content. | [optional] 

## Methods

### NewContentDto

`func NewContentDto(links map[string]ResourceLink, createdBy string, lastModifiedBy string, data interface{}, ) *ContentDto`

NewContentDto instantiates a new ContentDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentDtoWithDefaults

`func NewContentDtoWithDefaults() *ContentDto`

NewContentDtoWithDefaults instantiates a new ContentDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ContentDto) GetLinks() map[string]ResourceLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ContentDto) GetLinksOk() (*map[string]ResourceLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ContentDto) SetLinks(v map[string]ResourceLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *ContentDto) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentDto) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentDto) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ContentDto) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ContentDto) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ContentDto) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ContentDto) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetLastModifiedBy

`func (o *ContentDto) GetLastModifiedBy() string`

GetLastModifiedBy returns the LastModifiedBy field if non-nil, zero value otherwise.

### GetLastModifiedByOk

`func (o *ContentDto) GetLastModifiedByOk() (*string, bool)`

GetLastModifiedByOk returns a tuple with the LastModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModifiedBy

`func (o *ContentDto) SetLastModifiedBy(v string)`

SetLastModifiedBy sets LastModifiedBy field to given value.


### GetData

`func (o *ContentDto) GetData() interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContentDto) GetDataOk() (*interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContentDto) SetData(v interface{})`

SetData sets Data field to given value.


### SetDataNil

`func (o *ContentDto) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *ContentDto) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetReferenceData

`func (o *ContentDto) GetReferenceData() map[string]map[string]interface{}`

GetReferenceData returns the ReferenceData field if non-nil, zero value otherwise.

### GetReferenceDataOk

`func (o *ContentDto) GetReferenceDataOk() (*map[string]map[string]interface{}, bool)`

GetReferenceDataOk returns a tuple with the ReferenceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceData

`func (o *ContentDto) SetReferenceData(v map[string]map[string]interface{})`

SetReferenceData sets ReferenceData field to given value.

### HasReferenceData

`func (o *ContentDto) HasReferenceData() bool`

HasReferenceData returns a boolean if a field has been set.

### GetCreated

`func (o *ContentDto) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ContentDto) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ContentDto) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ContentDto) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastModified

`func (o *ContentDto) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *ContentDto) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *ContentDto) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.

### HasLastModified

`func (o *ContentDto) HasLastModified() bool`

HasLastModified returns a boolean if a field has been set.

### GetStatus

`func (o *ContentDto) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ContentDto) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ContentDto) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ContentDto) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetNewStatus

`func (o *ContentDto) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *ContentDto) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *ContentDto) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.

### HasNewStatus

`func (o *ContentDto) HasNewStatus() bool`

HasNewStatus returns a boolean if a field has been set.

### SetNewStatusNil

`func (o *ContentDto) SetNewStatusNil(b bool)`

 SetNewStatusNil sets the value for NewStatus to be an explicit nil

### UnsetNewStatus
`func (o *ContentDto) UnsetNewStatus()`

UnsetNewStatus ensures that no value is present for NewStatus, not even an explicit nil
### GetStatusColor

`func (o *ContentDto) GetStatusColor() string`

GetStatusColor returns the StatusColor field if non-nil, zero value otherwise.

### GetStatusColorOk

`func (o *ContentDto) GetStatusColorOk() (*string, bool)`

GetStatusColorOk returns a tuple with the StatusColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusColor

`func (o *ContentDto) SetStatusColor(v string)`

SetStatusColor sets StatusColor field to given value.

### HasStatusColor

`func (o *ContentDto) HasStatusColor() bool`

HasStatusColor returns a boolean if a field has been set.

### GetNewStatusColor

`func (o *ContentDto) GetNewStatusColor() string`

GetNewStatusColor returns the NewStatusColor field if non-nil, zero value otherwise.

### GetNewStatusColorOk

`func (o *ContentDto) GetNewStatusColorOk() (*string, bool)`

GetNewStatusColorOk returns a tuple with the NewStatusColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatusColor

`func (o *ContentDto) SetNewStatusColor(v string)`

SetNewStatusColor sets NewStatusColor field to given value.

### HasNewStatusColor

`func (o *ContentDto) HasNewStatusColor() bool`

HasNewStatusColor returns a boolean if a field has been set.

### SetNewStatusColorNil

`func (o *ContentDto) SetNewStatusColorNil(b bool)`

 SetNewStatusColorNil sets the value for NewStatusColor to be an explicit nil

### UnsetNewStatusColor
`func (o *ContentDto) UnsetNewStatusColor()`

UnsetNewStatusColor ensures that no value is present for NewStatusColor, not even an explicit nil
### GetEditToken

`func (o *ContentDto) GetEditToken() string`

GetEditToken returns the EditToken field if non-nil, zero value otherwise.

### GetEditTokenOk

`func (o *ContentDto) GetEditTokenOk() (*string, bool)`

GetEditTokenOk returns a tuple with the EditToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditToken

`func (o *ContentDto) SetEditToken(v string)`

SetEditToken sets EditToken field to given value.

### HasEditToken

`func (o *ContentDto) HasEditToken() bool`

HasEditToken returns a boolean if a field has been set.

### SetEditTokenNil

`func (o *ContentDto) SetEditTokenNil(b bool)`

 SetEditTokenNil sets the value for EditToken to be an explicit nil

### UnsetEditToken
`func (o *ContentDto) UnsetEditToken()`

UnsetEditToken ensures that no value is present for EditToken, not even an explicit nil
### GetScheduleJob

`func (o *ContentDto) GetScheduleJob() ScheduleJobDto`

GetScheduleJob returns the ScheduleJob field if non-nil, zero value otherwise.

### GetScheduleJobOk

`func (o *ContentDto) GetScheduleJobOk() (*ScheduleJobDto, bool)`

GetScheduleJobOk returns a tuple with the ScheduleJob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleJob

`func (o *ContentDto) SetScheduleJob(v ScheduleJobDto)`

SetScheduleJob sets ScheduleJob field to given value.

### HasScheduleJob

`func (o *ContentDto) HasScheduleJob() bool`

HasScheduleJob returns a boolean if a field has been set.

### GetSchemaId

`func (o *ContentDto) GetSchemaId() string`

GetSchemaId returns the SchemaId field if non-nil, zero value otherwise.

### GetSchemaIdOk

`func (o *ContentDto) GetSchemaIdOk() (*string, bool)`

GetSchemaIdOk returns a tuple with the SchemaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaId

`func (o *ContentDto) SetSchemaId(v string)`

SetSchemaId sets SchemaId field to given value.

### HasSchemaId

`func (o *ContentDto) HasSchemaId() bool`

HasSchemaId returns a boolean if a field has been set.

### GetSchemaName

`func (o *ContentDto) GetSchemaName() string`

GetSchemaName returns the SchemaName field if non-nil, zero value otherwise.

### GetSchemaNameOk

`func (o *ContentDto) GetSchemaNameOk() (*string, bool)`

GetSchemaNameOk returns a tuple with the SchemaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaName

`func (o *ContentDto) SetSchemaName(v string)`

SetSchemaName sets SchemaName field to given value.

### HasSchemaName

`func (o *ContentDto) HasSchemaName() bool`

HasSchemaName returns a boolean if a field has been set.

### SetSchemaNameNil

`func (o *ContentDto) SetSchemaNameNil(b bool)`

 SetSchemaNameNil sets the value for SchemaName to be an explicit nil

### UnsetSchemaName
`func (o *ContentDto) UnsetSchemaName()`

UnsetSchemaName ensures that no value is present for SchemaName, not even an explicit nil
### GetSchemaDisplayName

`func (o *ContentDto) GetSchemaDisplayName() string`

GetSchemaDisplayName returns the SchemaDisplayName field if non-nil, zero value otherwise.

### GetSchemaDisplayNameOk

`func (o *ContentDto) GetSchemaDisplayNameOk() (*string, bool)`

GetSchemaDisplayNameOk returns a tuple with the SchemaDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaDisplayName

`func (o *ContentDto) SetSchemaDisplayName(v string)`

SetSchemaDisplayName sets SchemaDisplayName field to given value.

### HasSchemaDisplayName

`func (o *ContentDto) HasSchemaDisplayName() bool`

HasSchemaDisplayName returns a boolean if a field has been set.

### SetSchemaDisplayNameNil

`func (o *ContentDto) SetSchemaDisplayNameNil(b bool)`

 SetSchemaDisplayNameNil sets the value for SchemaDisplayName to be an explicit nil

### UnsetSchemaDisplayName
`func (o *ContentDto) UnsetSchemaDisplayName()`

UnsetSchemaDisplayName ensures that no value is present for SchemaDisplayName, not even an explicit nil
### GetReferenceFields

`func (o *ContentDto) GetReferenceFields() []FieldDto`

GetReferenceFields returns the ReferenceFields field if non-nil, zero value otherwise.

### GetReferenceFieldsOk

`func (o *ContentDto) GetReferenceFieldsOk() (*[]FieldDto, bool)`

GetReferenceFieldsOk returns a tuple with the ReferenceFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceFields

`func (o *ContentDto) SetReferenceFields(v []FieldDto)`

SetReferenceFields sets ReferenceFields field to given value.

### HasReferenceFields

`func (o *ContentDto) HasReferenceFields() bool`

HasReferenceFields returns a boolean if a field has been set.

### SetReferenceFieldsNil

`func (o *ContentDto) SetReferenceFieldsNil(b bool)`

 SetReferenceFieldsNil sets the value for ReferenceFields to be an explicit nil

### UnsetReferenceFields
`func (o *ContentDto) UnsetReferenceFields()`

UnsetReferenceFields ensures that no value is present for ReferenceFields, not even an explicit nil
### GetIsDeleted

`func (o *ContentDto) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *ContentDto) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *ContentDto) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *ContentDto) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetVersion

`func (o *ContentDto) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ContentDto) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ContentDto) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ContentDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FeaturesDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | [**[]FeatureDto**](FeatureDto.md) | The latest features. | 
**Version** | Pointer to **int32** | The recent version. | [optional] 

## Methods

### NewFeaturesDto

`func NewFeaturesDto(features []FeatureDto, ) *FeaturesDto`

NewFeaturesDto instantiates a new FeaturesDto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturesDtoWithDefaults

`func NewFeaturesDtoWithDefaults() *FeaturesDto`

NewFeaturesDtoWithDefaults instantiates a new FeaturesDto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *FeaturesDto) GetFeatures() []FeatureDto`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *FeaturesDto) GetFeaturesOk() (*[]FeatureDto, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *FeaturesDto) SetFeatures(v []FeatureDto)`

SetFeatures sets Features field to given value.


### GetVersion

`func (o *FeaturesDto) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FeaturesDto) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FeaturesDto) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *FeaturesDto) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



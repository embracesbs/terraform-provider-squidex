# AssetsFieldPropertiesDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreviewMode** | Pointer to [**AssetPreviewMode**](AssetPreviewMode.md) |  | [optional] 
**DefaultValues** | Pointer to **map[string][]string** |  | [optional] 
**DefaultValue** | Pointer to **[]string** | The default value as a list of asset ids. | [optional] 
**FolderId** | Pointer to **NullableString** | The initial id to the folder. | [optional] 
**MinItems** | Pointer to **NullableInt32** | The minimum allowed items for the field value. | [optional] 
**MaxItems** | Pointer to **NullableInt32** | The maximum allowed items for the field value. | [optional] 
**MinSize** | Pointer to **NullableInt32** | The minimum file size in bytes. | [optional] 
**MaxSize** | Pointer to **NullableInt32** | The maximum file size in bytes. | [optional] 
**MinWidth** | Pointer to **NullableInt32** | The minimum image width in pixels. | [optional] 
**MaxWidth** | Pointer to **NullableInt32** | The maximum image width in pixels. | [optional] 
**MinHeight** | Pointer to **NullableInt32** | The minimum image height in pixels. | [optional] 
**MaxHeight** | Pointer to **NullableInt32** | The maximum image height in pixels. | [optional] 
**AspectWidth** | Pointer to **NullableInt32** | The image aspect width in pixels. | [optional] 
**AspectHeight** | Pointer to **NullableInt32** | The image aspect height in pixels. | [optional] 
**ExpectedType** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**ResolveFirst** | Pointer to **bool** | True to resolve first asset in the content list. | [optional] 
**MustBeImage** | Pointer to **bool** | True to resolve first image in the content list. | [optional] 
**ResolveImage** | Pointer to **bool** | True to resolve first image in the content list. | [optional] 
**AllowedExtensions** | Pointer to **[]string** | The allowed file extensions. | [optional] 
**AllowDuplicates** | Pointer to **bool** | True, if duplicate values are allowed. | [optional] 

## Methods

### NewAssetsFieldPropertiesDtoAllOf

`func NewAssetsFieldPropertiesDtoAllOf() *AssetsFieldPropertiesDtoAllOf`

NewAssetsFieldPropertiesDtoAllOf instantiates a new AssetsFieldPropertiesDtoAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetsFieldPropertiesDtoAllOfWithDefaults

`func NewAssetsFieldPropertiesDtoAllOfWithDefaults() *AssetsFieldPropertiesDtoAllOf`

NewAssetsFieldPropertiesDtoAllOfWithDefaults instantiates a new AssetsFieldPropertiesDtoAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreviewMode

`func (o *AssetsFieldPropertiesDtoAllOf) GetPreviewMode() AssetPreviewMode`

GetPreviewMode returns the PreviewMode field if non-nil, zero value otherwise.

### GetPreviewModeOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetPreviewModeOk() (*AssetPreviewMode, bool)`

GetPreviewModeOk returns a tuple with the PreviewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewMode

`func (o *AssetsFieldPropertiesDtoAllOf) SetPreviewMode(v AssetPreviewMode)`

SetPreviewMode sets PreviewMode field to given value.

### HasPreviewMode

`func (o *AssetsFieldPropertiesDtoAllOf) HasPreviewMode() bool`

HasPreviewMode returns a boolean if a field has been set.

### GetDefaultValues

`func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValues() map[string][]string`

GetDefaultValues returns the DefaultValues field if non-nil, zero value otherwise.

### GetDefaultValuesOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string][]string, bool)`

GetDefaultValuesOk returns a tuple with the DefaultValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValues

`func (o *AssetsFieldPropertiesDtoAllOf) SetDefaultValues(v map[string][]string)`

SetDefaultValues sets DefaultValues field to given value.

### HasDefaultValues

`func (o *AssetsFieldPropertiesDtoAllOf) HasDefaultValues() bool`

HasDefaultValues returns a boolean if a field has been set.

### GetDefaultValue

`func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValue() []string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValueOk() (*[]string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AssetsFieldPropertiesDtoAllOf) SetDefaultValue(v []string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AssetsFieldPropertiesDtoAllOf) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetFolderId

`func (o *AssetsFieldPropertiesDtoAllOf) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *AssetsFieldPropertiesDtoAllOf) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *AssetsFieldPropertiesDtoAllOf) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### SetFolderIdNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetFolderIdNil(b bool)`

 SetFolderIdNil sets the value for FolderId to be an explicit nil

### UnsetFolderId
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetFolderId()`

UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
### GetMinItems

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinItems() int32`

GetMinItems returns the MinItems field if non-nil, zero value otherwise.

### GetMinItemsOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinItemsOk() (*int32, bool)`

GetMinItemsOk returns a tuple with the MinItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinItems

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinItems(v int32)`

SetMinItems sets MinItems field to given value.

### HasMinItems

`func (o *AssetsFieldPropertiesDtoAllOf) HasMinItems() bool`

HasMinItems returns a boolean if a field has been set.

### SetMinItemsNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinItemsNil(b bool)`

 SetMinItemsNil sets the value for MinItems to be an explicit nil

### UnsetMinItems
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinItems()`

UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
### GetMaxItems

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxItems() int32`

GetMaxItems returns the MaxItems field if non-nil, zero value otherwise.

### GetMaxItemsOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxItemsOk() (*int32, bool)`

GetMaxItemsOk returns a tuple with the MaxItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxItems

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxItems(v int32)`

SetMaxItems sets MaxItems field to given value.

### HasMaxItems

`func (o *AssetsFieldPropertiesDtoAllOf) HasMaxItems() bool`

HasMaxItems returns a boolean if a field has been set.

### SetMaxItemsNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxItemsNil(b bool)`

 SetMaxItemsNil sets the value for MaxItems to be an explicit nil

### UnsetMaxItems
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxItems()`

UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
### GetMinSize

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinSize() int32`

GetMinSize returns the MinSize field if non-nil, zero value otherwise.

### GetMinSizeOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinSizeOk() (*int32, bool)`

GetMinSizeOk returns a tuple with the MinSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSize

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinSize(v int32)`

SetMinSize sets MinSize field to given value.

### HasMinSize

`func (o *AssetsFieldPropertiesDtoAllOf) HasMinSize() bool`

HasMinSize returns a boolean if a field has been set.

### SetMinSizeNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinSizeNil(b bool)`

 SetMinSizeNil sets the value for MinSize to be an explicit nil

### UnsetMinSize
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinSize()`

UnsetMinSize ensures that no value is present for MinSize, not even an explicit nil
### GetMaxSize

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *AssetsFieldPropertiesDtoAllOf) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.

### SetMaxSizeNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxSizeNil(b bool)`

 SetMaxSizeNil sets the value for MaxSize to be an explicit nil

### UnsetMaxSize
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxSize()`

UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
### GetMinWidth

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinWidth() int32`

GetMinWidth returns the MinWidth field if non-nil, zero value otherwise.

### GetMinWidthOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinWidthOk() (*int32, bool)`

GetMinWidthOk returns a tuple with the MinWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinWidth

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinWidth(v int32)`

SetMinWidth sets MinWidth field to given value.

### HasMinWidth

`func (o *AssetsFieldPropertiesDtoAllOf) HasMinWidth() bool`

HasMinWidth returns a boolean if a field has been set.

### SetMinWidthNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinWidthNil(b bool)`

 SetMinWidthNil sets the value for MinWidth to be an explicit nil

### UnsetMinWidth
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinWidth()`

UnsetMinWidth ensures that no value is present for MinWidth, not even an explicit nil
### GetMaxWidth

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxWidth() int32`

GetMaxWidth returns the MaxWidth field if non-nil, zero value otherwise.

### GetMaxWidthOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxWidthOk() (*int32, bool)`

GetMaxWidthOk returns a tuple with the MaxWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWidth

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxWidth(v int32)`

SetMaxWidth sets MaxWidth field to given value.

### HasMaxWidth

`func (o *AssetsFieldPropertiesDtoAllOf) HasMaxWidth() bool`

HasMaxWidth returns a boolean if a field has been set.

### SetMaxWidthNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxWidthNil(b bool)`

 SetMaxWidthNil sets the value for MaxWidth to be an explicit nil

### UnsetMaxWidth
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxWidth()`

UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
### GetMinHeight

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinHeight() int32`

GetMinHeight returns the MinHeight field if non-nil, zero value otherwise.

### GetMinHeightOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMinHeightOk() (*int32, bool)`

GetMinHeightOk returns a tuple with the MinHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinHeight

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinHeight(v int32)`

SetMinHeight sets MinHeight field to given value.

### HasMinHeight

`func (o *AssetsFieldPropertiesDtoAllOf) HasMinHeight() bool`

HasMinHeight returns a boolean if a field has been set.

### SetMinHeightNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMinHeightNil(b bool)`

 SetMinHeightNil sets the value for MinHeight to be an explicit nil

### UnsetMinHeight
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinHeight()`

UnsetMinHeight ensures that no value is present for MinHeight, not even an explicit nil
### GetMaxHeight

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxHeight() int32`

GetMaxHeight returns the MaxHeight field if non-nil, zero value otherwise.

### GetMaxHeightOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMaxHeightOk() (*int32, bool)`

GetMaxHeightOk returns a tuple with the MaxHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHeight

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxHeight(v int32)`

SetMaxHeight sets MaxHeight field to given value.

### HasMaxHeight

`func (o *AssetsFieldPropertiesDtoAllOf) HasMaxHeight() bool`

HasMaxHeight returns a boolean if a field has been set.

### SetMaxHeightNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetMaxHeightNil(b bool)`

 SetMaxHeightNil sets the value for MaxHeight to be an explicit nil

### UnsetMaxHeight
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxHeight()`

UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
### GetAspectWidth

`func (o *AssetsFieldPropertiesDtoAllOf) GetAspectWidth() int32`

GetAspectWidth returns the AspectWidth field if non-nil, zero value otherwise.

### GetAspectWidthOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetAspectWidthOk() (*int32, bool)`

GetAspectWidthOk returns a tuple with the AspectWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectWidth

`func (o *AssetsFieldPropertiesDtoAllOf) SetAspectWidth(v int32)`

SetAspectWidth sets AspectWidth field to given value.

### HasAspectWidth

`func (o *AssetsFieldPropertiesDtoAllOf) HasAspectWidth() bool`

HasAspectWidth returns a boolean if a field has been set.

### SetAspectWidthNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetAspectWidthNil(b bool)`

 SetAspectWidthNil sets the value for AspectWidth to be an explicit nil

### UnsetAspectWidth
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetAspectWidth()`

UnsetAspectWidth ensures that no value is present for AspectWidth, not even an explicit nil
### GetAspectHeight

`func (o *AssetsFieldPropertiesDtoAllOf) GetAspectHeight() int32`

GetAspectHeight returns the AspectHeight field if non-nil, zero value otherwise.

### GetAspectHeightOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetAspectHeightOk() (*int32, bool)`

GetAspectHeightOk returns a tuple with the AspectHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectHeight

`func (o *AssetsFieldPropertiesDtoAllOf) SetAspectHeight(v int32)`

SetAspectHeight sets AspectHeight field to given value.

### HasAspectHeight

`func (o *AssetsFieldPropertiesDtoAllOf) HasAspectHeight() bool`

HasAspectHeight returns a boolean if a field has been set.

### SetAspectHeightNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetAspectHeightNil(b bool)`

 SetAspectHeightNil sets the value for AspectHeight to be an explicit nil

### UnsetAspectHeight
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetAspectHeight()`

UnsetAspectHeight ensures that no value is present for AspectHeight, not even an explicit nil
### GetExpectedType

`func (o *AssetsFieldPropertiesDtoAllOf) GetExpectedType() AssetType`

GetExpectedType returns the ExpectedType field if non-nil, zero value otherwise.

### GetExpectedTypeOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetExpectedTypeOk() (*AssetType, bool)`

GetExpectedTypeOk returns a tuple with the ExpectedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedType

`func (o *AssetsFieldPropertiesDtoAllOf) SetExpectedType(v AssetType)`

SetExpectedType sets ExpectedType field to given value.

### HasExpectedType

`func (o *AssetsFieldPropertiesDtoAllOf) HasExpectedType() bool`

HasExpectedType returns a boolean if a field has been set.

### GetResolveFirst

`func (o *AssetsFieldPropertiesDtoAllOf) GetResolveFirst() bool`

GetResolveFirst returns the ResolveFirst field if non-nil, zero value otherwise.

### GetResolveFirstOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetResolveFirstOk() (*bool, bool)`

GetResolveFirstOk returns a tuple with the ResolveFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveFirst

`func (o *AssetsFieldPropertiesDtoAllOf) SetResolveFirst(v bool)`

SetResolveFirst sets ResolveFirst field to given value.

### HasResolveFirst

`func (o *AssetsFieldPropertiesDtoAllOf) HasResolveFirst() bool`

HasResolveFirst returns a boolean if a field has been set.

### GetMustBeImage

`func (o *AssetsFieldPropertiesDtoAllOf) GetMustBeImage() bool`

GetMustBeImage returns the MustBeImage field if non-nil, zero value otherwise.

### GetMustBeImageOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetMustBeImageOk() (*bool, bool)`

GetMustBeImageOk returns a tuple with the MustBeImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustBeImage

`func (o *AssetsFieldPropertiesDtoAllOf) SetMustBeImage(v bool)`

SetMustBeImage sets MustBeImage field to given value.

### HasMustBeImage

`func (o *AssetsFieldPropertiesDtoAllOf) HasMustBeImage() bool`

HasMustBeImage returns a boolean if a field has been set.

### GetResolveImage

`func (o *AssetsFieldPropertiesDtoAllOf) GetResolveImage() bool`

GetResolveImage returns the ResolveImage field if non-nil, zero value otherwise.

### GetResolveImageOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetResolveImageOk() (*bool, bool)`

GetResolveImageOk returns a tuple with the ResolveImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolveImage

`func (o *AssetsFieldPropertiesDtoAllOf) SetResolveImage(v bool)`

SetResolveImage sets ResolveImage field to given value.

### HasResolveImage

`func (o *AssetsFieldPropertiesDtoAllOf) HasResolveImage() bool`

HasResolveImage returns a boolean if a field has been set.

### GetAllowedExtensions

`func (o *AssetsFieldPropertiesDtoAllOf) GetAllowedExtensions() []string`

GetAllowedExtensions returns the AllowedExtensions field if non-nil, zero value otherwise.

### GetAllowedExtensionsOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetAllowedExtensionsOk() (*[]string, bool)`

GetAllowedExtensionsOk returns a tuple with the AllowedExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedExtensions

`func (o *AssetsFieldPropertiesDtoAllOf) SetAllowedExtensions(v []string)`

SetAllowedExtensions sets AllowedExtensions field to given value.

### HasAllowedExtensions

`func (o *AssetsFieldPropertiesDtoAllOf) HasAllowedExtensions() bool`

HasAllowedExtensions returns a boolean if a field has been set.

### SetAllowedExtensionsNil

`func (o *AssetsFieldPropertiesDtoAllOf) SetAllowedExtensionsNil(b bool)`

 SetAllowedExtensionsNil sets the value for AllowedExtensions to be an explicit nil

### UnsetAllowedExtensions
`func (o *AssetsFieldPropertiesDtoAllOf) UnsetAllowedExtensions()`

UnsetAllowedExtensions ensures that no value is present for AllowedExtensions, not even an explicit nil
### GetAllowDuplicates

`func (o *AssetsFieldPropertiesDtoAllOf) GetAllowDuplicates() bool`

GetAllowDuplicates returns the AllowDuplicates field if non-nil, zero value otherwise.

### GetAllowDuplicatesOk

`func (o *AssetsFieldPropertiesDtoAllOf) GetAllowDuplicatesOk() (*bool, bool)`

GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuplicates

`func (o *AssetsFieldPropertiesDtoAllOf) SetAllowDuplicates(v bool)`

SetAllowDuplicates sets AllowDuplicates field to given value.

### HasAllowDuplicates

`func (o *AssetsFieldPropertiesDtoAllOf) HasAllowDuplicates() bool`

HasAllowDuplicates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



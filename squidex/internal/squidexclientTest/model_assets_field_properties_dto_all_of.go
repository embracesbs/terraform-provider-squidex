/*
Squidex API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 7.0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package squidexclient

import (
	"encoding/json"
)

// AssetsFieldPropertiesDtoAllOf struct for AssetsFieldPropertiesDtoAllOf
type AssetsFieldPropertiesDtoAllOf struct {
	PreviewMode *AssetPreviewMode `json:"previewMode,omitempty"`
	DefaultValues *map[string][]string `json:"defaultValues,omitempty"`
	// The default value as a list of asset ids.
	DefaultValue []string `json:"defaultValue,omitempty"`
	// The initial id to the folder.
	FolderId NullableString `json:"folderId,omitempty"`
	// The minimum allowed items for the field value.
	MinItems NullableInt32 `json:"minItems,omitempty"`
	// The maximum allowed items for the field value.
	MaxItems NullableInt32 `json:"maxItems,omitempty"`
	// The minimum file size in bytes.
	MinSize NullableInt32 `json:"minSize,omitempty"`
	// The maximum file size in bytes.
	MaxSize NullableInt32 `json:"maxSize,omitempty"`
	// The minimum image width in pixels.
	MinWidth NullableInt32 `json:"minWidth,omitempty"`
	// The maximum image width in pixels.
	MaxWidth NullableInt32 `json:"maxWidth,omitempty"`
	// The minimum image height in pixels.
	MinHeight NullableInt32 `json:"minHeight,omitempty"`
	// The maximum image height in pixels.
	MaxHeight NullableInt32 `json:"maxHeight,omitempty"`
	// The image aspect width in pixels.
	AspectWidth NullableInt32 `json:"aspectWidth,omitempty"`
	// The image aspect height in pixels.
	AspectHeight NullableInt32 `json:"aspectHeight,omitempty"`
	ExpectedType *AssetType `json:"expectedType,omitempty"`
	// True to resolve first asset in the content list.
	ResolveFirst *bool `json:"resolveFirst,omitempty"`
	// True to resolve first image in the content list.
	// Deprecated
	MustBeImage *bool `json:"mustBeImage,omitempty"`
	// True to resolve first image in the content list.
	// Deprecated
	ResolveImage *bool `json:"resolveImage,omitempty"`
	// The allowed file extensions.
	AllowedExtensions []string `json:"allowedExtensions,omitempty"`
	// True, if duplicate values are allowed.
	AllowDuplicates *bool `json:"allowDuplicates,omitempty"`
}

// NewAssetsFieldPropertiesDtoAllOf instantiates a new AssetsFieldPropertiesDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetsFieldPropertiesDtoAllOf() *AssetsFieldPropertiesDtoAllOf {
	this := AssetsFieldPropertiesDtoAllOf{}
	return &this
}

// NewAssetsFieldPropertiesDtoAllOfWithDefaults instantiates a new AssetsFieldPropertiesDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetsFieldPropertiesDtoAllOfWithDefaults() *AssetsFieldPropertiesDtoAllOf {
	this := AssetsFieldPropertiesDtoAllOf{}
	return &this
}

// GetPreviewMode returns the PreviewMode field value if set, zero value otherwise.
func (o *AssetsFieldPropertiesDtoAllOf) GetPreviewMode() AssetPreviewMode {
	if o == nil || isNil(o.PreviewMode) {
		var ret AssetPreviewMode
		return ret
	}
	return *o.PreviewMode
}

// GetPreviewModeOk returns a tuple with the PreviewMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsFieldPropertiesDtoAllOf) GetPreviewModeOk() (*AssetPreviewMode, bool) {
	if o == nil || isNil(o.PreviewMode) {
    return nil, false
	}
	return o.PreviewMode, true
}

// HasPreviewMode returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasPreviewMode() bool {
	if o != nil && !isNil(o.PreviewMode) {
		return true
	}

	return false
}

// SetPreviewMode gets a reference to the given AssetPreviewMode and assigns it to the PreviewMode field.
func (o *AssetsFieldPropertiesDtoAllOf) SetPreviewMode(v AssetPreviewMode) {
	o.PreviewMode = &v
}

// GetDefaultValues returns the DefaultValues field value if set, zero value otherwise.
func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValues() map[string][]string {
	if o == nil || isNil(o.DefaultValues) {
		var ret map[string][]string
		return ret
	}
	return *o.DefaultValues
}

// GetDefaultValuesOk returns a tuple with the DefaultValues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValuesOk() (*map[string][]string, bool) {
	if o == nil || isNil(o.DefaultValues) {
    return nil, false
	}
	return o.DefaultValues, true
}

// HasDefaultValues returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasDefaultValues() bool {
	if o != nil && !isNil(o.DefaultValues) {
		return true
	}

	return false
}

// SetDefaultValues gets a reference to the given map[string][]string and assigns it to the DefaultValues field.
func (o *AssetsFieldPropertiesDtoAllOf) SetDefaultValues(v map[string][]string) {
	o.DefaultValues = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValue() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetDefaultValueOk() ([]string, bool) {
	if o == nil || isNil(o.DefaultValue) {
    return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasDefaultValue() bool {
	if o != nil && isNil(o.DefaultValue) {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given []string and assigns it to the DefaultValue field.
func (o *AssetsFieldPropertiesDtoAllOf) SetDefaultValue(v []string) {
	o.DefaultValue = v
}

// GetFolderId returns the FolderId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetFolderId() string {
	if o == nil || isNil(o.FolderId.Get()) {
		var ret string
		return ret
	}
	return *o.FolderId.Get()
}

// GetFolderIdOk returns a tuple with the FolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetFolderIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return o.FolderId.Get(), o.FolderId.IsSet()
}

// HasFolderId returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasFolderId() bool {
	if o != nil && o.FolderId.IsSet() {
		return true
	}

	return false
}

// SetFolderId gets a reference to the given NullableString and assigns it to the FolderId field.
func (o *AssetsFieldPropertiesDtoAllOf) SetFolderId(v string) {
	o.FolderId.Set(&v)
}
// SetFolderIdNil sets the value for FolderId to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetFolderIdNil() {
	o.FolderId.Set(nil)
}

// UnsetFolderId ensures that no value is present for FolderId, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetFolderId() {
	o.FolderId.Unset()
}

// GetMinItems returns the MinItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMinItems() int32 {
	if o == nil || isNil(o.MinItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MinItems.Get()
}

// GetMinItemsOk returns a tuple with the MinItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMinItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinItems.Get(), o.MinItems.IsSet()
}

// HasMinItems returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMinItems() bool {
	if o != nil && o.MinItems.IsSet() {
		return true
	}

	return false
}

// SetMinItems gets a reference to the given NullableInt32 and assigns it to the MinItems field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMinItems(v int32) {
	o.MinItems.Set(&v)
}
// SetMinItemsNil sets the value for MinItems to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMinItemsNil() {
	o.MinItems.Set(nil)
}

// UnsetMinItems ensures that no value is present for MinItems, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinItems() {
	o.MinItems.Unset()
}

// GetMaxItems returns the MaxItems field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxItems() int32 {
	if o == nil || isNil(o.MaxItems.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxItems.Get()
}

// GetMaxItemsOk returns a tuple with the MaxItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxItemsOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxItems.Get(), o.MaxItems.IsSet()
}

// HasMaxItems returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMaxItems() bool {
	if o != nil && o.MaxItems.IsSet() {
		return true
	}

	return false
}

// SetMaxItems gets a reference to the given NullableInt32 and assigns it to the MaxItems field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxItems(v int32) {
	o.MaxItems.Set(&v)
}
// SetMaxItemsNil sets the value for MaxItems to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxItemsNil() {
	o.MaxItems.Set(nil)
}

// UnsetMaxItems ensures that no value is present for MaxItems, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxItems() {
	o.MaxItems.Unset()
}

// GetMinSize returns the MinSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMinSize() int32 {
	if o == nil || isNil(o.MinSize.Get()) {
		var ret int32
		return ret
	}
	return *o.MinSize.Get()
}

// GetMinSizeOk returns a tuple with the MinSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMinSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinSize.Get(), o.MinSize.IsSet()
}

// HasMinSize returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMinSize() bool {
	if o != nil && o.MinSize.IsSet() {
		return true
	}

	return false
}

// SetMinSize gets a reference to the given NullableInt32 and assigns it to the MinSize field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMinSize(v int32) {
	o.MinSize.Set(&v)
}
// SetMinSizeNil sets the value for MinSize to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMinSizeNil() {
	o.MinSize.Set(nil)
}

// UnsetMinSize ensures that no value is present for MinSize, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinSize() {
	o.MinSize.Unset()
}

// GetMaxSize returns the MaxSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxSize() int32 {
	if o == nil || isNil(o.MaxSize.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxSize.Get()
}

// GetMaxSizeOk returns a tuple with the MaxSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxSizeOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxSize.Get(), o.MaxSize.IsSet()
}

// HasMaxSize returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMaxSize() bool {
	if o != nil && o.MaxSize.IsSet() {
		return true
	}

	return false
}

// SetMaxSize gets a reference to the given NullableInt32 and assigns it to the MaxSize field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxSize(v int32) {
	o.MaxSize.Set(&v)
}
// SetMaxSizeNil sets the value for MaxSize to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxSizeNil() {
	o.MaxSize.Set(nil)
}

// UnsetMaxSize ensures that no value is present for MaxSize, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxSize() {
	o.MaxSize.Unset()
}

// GetMinWidth returns the MinWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMinWidth() int32 {
	if o == nil || isNil(o.MinWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.MinWidth.Get()
}

// GetMinWidthOk returns a tuple with the MinWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMinWidthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinWidth.Get(), o.MinWidth.IsSet()
}

// HasMinWidth returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMinWidth() bool {
	if o != nil && o.MinWidth.IsSet() {
		return true
	}

	return false
}

// SetMinWidth gets a reference to the given NullableInt32 and assigns it to the MinWidth field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMinWidth(v int32) {
	o.MinWidth.Set(&v)
}
// SetMinWidthNil sets the value for MinWidth to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMinWidthNil() {
	o.MinWidth.Set(nil)
}

// UnsetMinWidth ensures that no value is present for MinWidth, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinWidth() {
	o.MinWidth.Unset()
}

// GetMaxWidth returns the MaxWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxWidth() int32 {
	if o == nil || isNil(o.MaxWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxWidth.Get()
}

// GetMaxWidthOk returns a tuple with the MaxWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxWidthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxWidth.Get(), o.MaxWidth.IsSet()
}

// HasMaxWidth returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMaxWidth() bool {
	if o != nil && o.MaxWidth.IsSet() {
		return true
	}

	return false
}

// SetMaxWidth gets a reference to the given NullableInt32 and assigns it to the MaxWidth field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxWidth(v int32) {
	o.MaxWidth.Set(&v)
}
// SetMaxWidthNil sets the value for MaxWidth to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxWidthNil() {
	o.MaxWidth.Set(nil)
}

// UnsetMaxWidth ensures that no value is present for MaxWidth, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxWidth() {
	o.MaxWidth.Unset()
}

// GetMinHeight returns the MinHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMinHeight() int32 {
	if o == nil || isNil(o.MinHeight.Get()) {
		var ret int32
		return ret
	}
	return *o.MinHeight.Get()
}

// GetMinHeightOk returns a tuple with the MinHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMinHeightOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MinHeight.Get(), o.MinHeight.IsSet()
}

// HasMinHeight returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMinHeight() bool {
	if o != nil && o.MinHeight.IsSet() {
		return true
	}

	return false
}

// SetMinHeight gets a reference to the given NullableInt32 and assigns it to the MinHeight field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMinHeight(v int32) {
	o.MinHeight.Set(&v)
}
// SetMinHeightNil sets the value for MinHeight to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMinHeightNil() {
	o.MinHeight.Set(nil)
}

// UnsetMinHeight ensures that no value is present for MinHeight, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMinHeight() {
	o.MinHeight.Unset()
}

// GetMaxHeight returns the MaxHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxHeight() int32 {
	if o == nil || isNil(o.MaxHeight.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxHeight.Get()
}

// GetMaxHeightOk returns a tuple with the MaxHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetMaxHeightOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.MaxHeight.Get(), o.MaxHeight.IsSet()
}

// HasMaxHeight returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMaxHeight() bool {
	if o != nil && o.MaxHeight.IsSet() {
		return true
	}

	return false
}

// SetMaxHeight gets a reference to the given NullableInt32 and assigns it to the MaxHeight field.
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxHeight(v int32) {
	o.MaxHeight.Set(&v)
}
// SetMaxHeightNil sets the value for MaxHeight to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetMaxHeightNil() {
	o.MaxHeight.Set(nil)
}

// UnsetMaxHeight ensures that no value is present for MaxHeight, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetMaxHeight() {
	o.MaxHeight.Unset()
}

// GetAspectWidth returns the AspectWidth field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetAspectWidth() int32 {
	if o == nil || isNil(o.AspectWidth.Get()) {
		var ret int32
		return ret
	}
	return *o.AspectWidth.Get()
}

// GetAspectWidthOk returns a tuple with the AspectWidth field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetAspectWidthOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.AspectWidth.Get(), o.AspectWidth.IsSet()
}

// HasAspectWidth returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasAspectWidth() bool {
	if o != nil && o.AspectWidth.IsSet() {
		return true
	}

	return false
}

// SetAspectWidth gets a reference to the given NullableInt32 and assigns it to the AspectWidth field.
func (o *AssetsFieldPropertiesDtoAllOf) SetAspectWidth(v int32) {
	o.AspectWidth.Set(&v)
}
// SetAspectWidthNil sets the value for AspectWidth to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetAspectWidthNil() {
	o.AspectWidth.Set(nil)
}

// UnsetAspectWidth ensures that no value is present for AspectWidth, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetAspectWidth() {
	o.AspectWidth.Unset()
}

// GetAspectHeight returns the AspectHeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetAspectHeight() int32 {
	if o == nil || isNil(o.AspectHeight.Get()) {
		var ret int32
		return ret
	}
	return *o.AspectHeight.Get()
}

// GetAspectHeightOk returns a tuple with the AspectHeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetAspectHeightOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return o.AspectHeight.Get(), o.AspectHeight.IsSet()
}

// HasAspectHeight returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasAspectHeight() bool {
	if o != nil && o.AspectHeight.IsSet() {
		return true
	}

	return false
}

// SetAspectHeight gets a reference to the given NullableInt32 and assigns it to the AspectHeight field.
func (o *AssetsFieldPropertiesDtoAllOf) SetAspectHeight(v int32) {
	o.AspectHeight.Set(&v)
}
// SetAspectHeightNil sets the value for AspectHeight to be an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) SetAspectHeightNil() {
	o.AspectHeight.Set(nil)
}

// UnsetAspectHeight ensures that no value is present for AspectHeight, not even an explicit nil
func (o *AssetsFieldPropertiesDtoAllOf) UnsetAspectHeight() {
	o.AspectHeight.Unset()
}

// GetExpectedType returns the ExpectedType field value if set, zero value otherwise.
func (o *AssetsFieldPropertiesDtoAllOf) GetExpectedType() AssetType {
	if o == nil || isNil(o.ExpectedType) {
		var ret AssetType
		return ret
	}
	return *o.ExpectedType
}

// GetExpectedTypeOk returns a tuple with the ExpectedType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsFieldPropertiesDtoAllOf) GetExpectedTypeOk() (*AssetType, bool) {
	if o == nil || isNil(o.ExpectedType) {
    return nil, false
	}
	return o.ExpectedType, true
}

// HasExpectedType returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasExpectedType() bool {
	if o != nil && !isNil(o.ExpectedType) {
		return true
	}

	return false
}

// SetExpectedType gets a reference to the given AssetType and assigns it to the ExpectedType field.
func (o *AssetsFieldPropertiesDtoAllOf) SetExpectedType(v AssetType) {
	o.ExpectedType = &v
}

// GetResolveFirst returns the ResolveFirst field value if set, zero value otherwise.
func (o *AssetsFieldPropertiesDtoAllOf) GetResolveFirst() bool {
	if o == nil || isNil(o.ResolveFirst) {
		var ret bool
		return ret
	}
	return *o.ResolveFirst
}

// GetResolveFirstOk returns a tuple with the ResolveFirst field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsFieldPropertiesDtoAllOf) GetResolveFirstOk() (*bool, bool) {
	if o == nil || isNil(o.ResolveFirst) {
    return nil, false
	}
	return o.ResolveFirst, true
}

// HasResolveFirst returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasResolveFirst() bool {
	if o != nil && !isNil(o.ResolveFirst) {
		return true
	}

	return false
}

// SetResolveFirst gets a reference to the given bool and assigns it to the ResolveFirst field.
func (o *AssetsFieldPropertiesDtoAllOf) SetResolveFirst(v bool) {
	o.ResolveFirst = &v
}

// GetMustBeImage returns the MustBeImage field value if set, zero value otherwise.
// Deprecated
func (o *AssetsFieldPropertiesDtoAllOf) GetMustBeImage() bool {
	if o == nil || isNil(o.MustBeImage) {
		var ret bool
		return ret
	}
	return *o.MustBeImage
}

// GetMustBeImageOk returns a tuple with the MustBeImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AssetsFieldPropertiesDtoAllOf) GetMustBeImageOk() (*bool, bool) {
	if o == nil || isNil(o.MustBeImage) {
    return nil, false
	}
	return o.MustBeImage, true
}

// HasMustBeImage returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasMustBeImage() bool {
	if o != nil && !isNil(o.MustBeImage) {
		return true
	}

	return false
}

// SetMustBeImage gets a reference to the given bool and assigns it to the MustBeImage field.
// Deprecated
func (o *AssetsFieldPropertiesDtoAllOf) SetMustBeImage(v bool) {
	o.MustBeImage = &v
}

// GetResolveImage returns the ResolveImage field value if set, zero value otherwise.
// Deprecated
func (o *AssetsFieldPropertiesDtoAllOf) GetResolveImage() bool {
	if o == nil || isNil(o.ResolveImage) {
		var ret bool
		return ret
	}
	return *o.ResolveImage
}

// GetResolveImageOk returns a tuple with the ResolveImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *AssetsFieldPropertiesDtoAllOf) GetResolveImageOk() (*bool, bool) {
	if o == nil || isNil(o.ResolveImage) {
    return nil, false
	}
	return o.ResolveImage, true
}

// HasResolveImage returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasResolveImage() bool {
	if o != nil && !isNil(o.ResolveImage) {
		return true
	}

	return false
}

// SetResolveImage gets a reference to the given bool and assigns it to the ResolveImage field.
// Deprecated
func (o *AssetsFieldPropertiesDtoAllOf) SetResolveImage(v bool) {
	o.ResolveImage = &v
}

// GetAllowedExtensions returns the AllowedExtensions field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AssetsFieldPropertiesDtoAllOf) GetAllowedExtensions() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.AllowedExtensions
}

// GetAllowedExtensionsOk returns a tuple with the AllowedExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AssetsFieldPropertiesDtoAllOf) GetAllowedExtensionsOk() ([]string, bool) {
	if o == nil || isNil(o.AllowedExtensions) {
    return nil, false
	}
	return o.AllowedExtensions, true
}

// HasAllowedExtensions returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasAllowedExtensions() bool {
	if o != nil && isNil(o.AllowedExtensions) {
		return true
	}

	return false
}

// SetAllowedExtensions gets a reference to the given []string and assigns it to the AllowedExtensions field.
func (o *AssetsFieldPropertiesDtoAllOf) SetAllowedExtensions(v []string) {
	o.AllowedExtensions = v
}

// GetAllowDuplicates returns the AllowDuplicates field value if set, zero value otherwise.
func (o *AssetsFieldPropertiesDtoAllOf) GetAllowDuplicates() bool {
	if o == nil || isNil(o.AllowDuplicates) {
		var ret bool
		return ret
	}
	return *o.AllowDuplicates
}

// GetAllowDuplicatesOk returns a tuple with the AllowDuplicates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetsFieldPropertiesDtoAllOf) GetAllowDuplicatesOk() (*bool, bool) {
	if o == nil || isNil(o.AllowDuplicates) {
    return nil, false
	}
	return o.AllowDuplicates, true
}

// HasAllowDuplicates returns a boolean if a field has been set.
func (o *AssetsFieldPropertiesDtoAllOf) HasAllowDuplicates() bool {
	if o != nil && !isNil(o.AllowDuplicates) {
		return true
	}

	return false
}

// SetAllowDuplicates gets a reference to the given bool and assigns it to the AllowDuplicates field.
func (o *AssetsFieldPropertiesDtoAllOf) SetAllowDuplicates(v bool) {
	o.AllowDuplicates = &v
}

func (o AssetsFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.PreviewMode) {
		toSerialize["previewMode"] = o.PreviewMode
	}
	if !isNil(o.DefaultValues) {
		toSerialize["defaultValues"] = o.DefaultValues
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if o.FolderId.IsSet() {
		toSerialize["folderId"] = o.FolderId.Get()
	}
	if o.MinItems.IsSet() {
		toSerialize["minItems"] = o.MinItems.Get()
	}
	if o.MaxItems.IsSet() {
		toSerialize["maxItems"] = o.MaxItems.Get()
	}
	if o.MinSize.IsSet() {
		toSerialize["minSize"] = o.MinSize.Get()
	}
	if o.MaxSize.IsSet() {
		toSerialize["maxSize"] = o.MaxSize.Get()
	}
	if o.MinWidth.IsSet() {
		toSerialize["minWidth"] = o.MinWidth.Get()
	}
	if o.MaxWidth.IsSet() {
		toSerialize["maxWidth"] = o.MaxWidth.Get()
	}
	if o.MinHeight.IsSet() {
		toSerialize["minHeight"] = o.MinHeight.Get()
	}
	if o.MaxHeight.IsSet() {
		toSerialize["maxHeight"] = o.MaxHeight.Get()
	}
	if o.AspectWidth.IsSet() {
		toSerialize["aspectWidth"] = o.AspectWidth.Get()
	}
	if o.AspectHeight.IsSet() {
		toSerialize["aspectHeight"] = o.AspectHeight.Get()
	}
	if !isNil(o.ExpectedType) {
		toSerialize["expectedType"] = o.ExpectedType
	}
	if !isNil(o.ResolveFirst) {
		toSerialize["resolveFirst"] = o.ResolveFirst
	}
	if !isNil(o.MustBeImage) {
		toSerialize["mustBeImage"] = o.MustBeImage
	}
	if !isNil(o.ResolveImage) {
		toSerialize["resolveImage"] = o.ResolveImage
	}
	if o.AllowedExtensions != nil {
		toSerialize["allowedExtensions"] = o.AllowedExtensions
	}
	if !isNil(o.AllowDuplicates) {
		toSerialize["allowDuplicates"] = o.AllowDuplicates
	}
	return json.Marshal(toSerialize)
}

type NullableAssetsFieldPropertiesDtoAllOf struct {
	value *AssetsFieldPropertiesDtoAllOf
	isSet bool
}

func (v NullableAssetsFieldPropertiesDtoAllOf) Get() *AssetsFieldPropertiesDtoAllOf {
	return v.value
}

func (v *NullableAssetsFieldPropertiesDtoAllOf) Set(val *AssetsFieldPropertiesDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetsFieldPropertiesDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetsFieldPropertiesDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetsFieldPropertiesDtoAllOf(val *AssetsFieldPropertiesDtoAllOf) *NullableAssetsFieldPropertiesDtoAllOf {
	return &NullableAssetsFieldPropertiesDtoAllOf{value: val, isSet: true}
}

func (v NullableAssetsFieldPropertiesDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetsFieldPropertiesDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


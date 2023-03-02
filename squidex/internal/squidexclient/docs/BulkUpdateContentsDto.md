# BulkUpdateContentsDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]BulkUpdateContentsJobDto**](BulkUpdateContentsJobDto.md) | The contents to update or insert. | 
**Publish** | **bool** | True to automatically publish the content. | [optional] 
**DoNotScript** | **bool** | True to turn off scripting for faster inserts. Default: true. | [optional] 
**DoNotValidate** | **bool** | True to turn off validation for faster inserts. Default: false. | [optional] 
**DoNotValidateWorkflow** | **bool** | True to turn off validation of workflow rules. Default: false. | [optional] 
**CheckReferrers** | **bool** | True to check referrers of deleted contents. | [optional] 
**OptimizeValidation** | **bool** | True to turn off costly validation: Unique checks, asset checks and reference checks. Default: true. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



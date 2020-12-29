# BackupJobDtoAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The id of the backup job. | [optional] 
**Started** | [**time.Time**](time.Time.md) | The time when the job has been started. | [optional] 
**Stopped** | Pointer to [**time.Time**](time.Time.md) | The time when the job has been stopped. | [optional] 
**HandledEvents** | **int32** | The number of handled events. | [optional] 
**HandledAssets** | **int32** | The number of handled assets. | [optional] 
**Status** | [**OneOfJobStatus**](oneOf&lt;JobStatus&gt;.md) | The status of the operation. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



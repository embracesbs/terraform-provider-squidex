# \BackupsApi

All URIs are relative to *https://squidex-embracecloudte.features.embracecloud.io/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BackupContentGetBackupContent**](BackupsApi.md#BackupContentGetBackupContent) | **Get** /apps/{app}/backups/{id} | Get the backup content.
[**BackupsDeleteBackup**](BackupsApi.md#BackupsDeleteBackup) | **Delete** /apps/{app}/backups/{id} | Delete a backup.
[**BackupsGetBackups**](BackupsApi.md#BackupsGetBackups) | **Get** /apps/{app}/backups | Get all backup jobs.
[**BackupsPostBackup**](BackupsApi.md#BackupsPostBackup) | **Post** /apps/{app}/backups | Start a new backup.
[**RestoreGetRestoreJob**](BackupsApi.md#RestoreGetRestoreJob) | **Get** /apps/restore | Get current restore status.
[**RestorePostRestoreJob**](BackupsApi.md#RestorePostRestoreJob) | **Post** /apps/restore | Restore a backup.



## BackupContentGetBackupContent

> *os.File BackupContentGetBackupContent(ctx, app, id)

Get the backup content.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the asset. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/octet-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupsDeleteBackup

> BackupsDeleteBackup(ctx, app, id)

Delete a backup.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 
**id** | **string**| The id of the backup to delete. | 

### Return type

 (empty response body)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupsGetBackups

> BackupJobsDto BackupsGetBackups(ctx, app)

Get all backup jobs.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

[**BackupJobsDto**](BackupJobsDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BackupsPostBackup

> BackupsPostBackup(ctx, app)

Start a new backup.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**app** | **string**| The name of the app. | 

### Return type

 (empty response body)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreGetRestoreJob

> RestoreJobDto RestoreGetRestoreJob(ctx, )

Get current restore status.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**RestoreJobDto**](RestoreJobDto.md)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestorePostRestoreJob

> RestorePostRestoreJob(ctx, restoreRequestDto)

Restore a backup.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**restoreRequestDto** | [**RestoreRequestDto**](RestoreRequestDto.md)| The backup to restore. | 

### Return type

 (empty response body)

### Authorization

[squidex-oauth-auth](../README.md#squidex-oauth-auth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


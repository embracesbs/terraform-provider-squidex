# AppDto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**map[string]ResourceLink**](ResourceLink.md) | The links. | 
**Name** | **string** | The name of the app. | 
**Label** | Pointer to **string** | The optional label of the app. | [optional] 
**Description** | Pointer to **string** | The optional description of the app. | [optional] 
**Version** | **int64** | The version of the app. | [optional] 
**Id** | **string** | The id of the app. | [optional] 
**Created** | [**time.Time**](time.Time.md) | The timestamp when the app has been created. | [optional] 
**LastModified** | [**time.Time**](time.Time.md) | The timestamp when the app has been modified last. | [optional] 
**Permissions** | **[]string** | The permission level of the user. | [optional] 
**CanAccessApi** | **bool** | Indicates if the user can access the api. | [optional] 
**CanAccessContent** | **bool** | Indicates if the user can access at least one content. | [optional] 
**RoleName** | Pointer to **string** | The role name of the user. | [optional] 
**RoleProperties** | [**interface{}**](.md) | The properties from the role. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



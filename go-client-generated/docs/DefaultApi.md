# {{classname}}

All URIs are relative to *localhost:30010*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsersFoldersFilesFileDelete**](DefaultApi.md#UsersFoldersFilesFileDelete) | **Delete** /users/folders/files/file | 
[**UsersFoldersFilesFileGet**](DefaultApi.md#UsersFoldersFilesFileGet) | **Get** /users/folders/files/file | 
[**UsersFoldersFilesFilePost**](DefaultApi.md#UsersFoldersFilesFilePost) | **Post** /users/folders/files/file | 
[**UsersFoldersFilesFilePut**](DefaultApi.md#UsersFoldersFilesFilePut) | **Put** /users/folders/files/file | 
[**UsersFoldersFilesGet**](DefaultApi.md#UsersFoldersFilesGet) | **Get** /users/folders/files | 
[**UsersFoldersGet**](DefaultApi.md#UsersFoldersGet) | **Get** /users/folders | 
[**UsersFoldersPost**](DefaultApi.md#UsersFoldersPost) | **Post** /users/folders | 
[**UsersLoginPost**](DefaultApi.md#UsersLoginPost) | **Post** /users/login | 
[**UsersSettingsDelete**](DefaultApi.md#UsersSettingsDelete) | **Delete** /users/settings | 
[**UsersSettingsGet**](DefaultApi.md#UsersSettingsGet) | **Get** /users/settings | 
[**UsersSettingsPut**](DefaultApi.md#UsersSettingsPut) | **Put** /users/settings | 
[**UsersSignupPost**](DefaultApi.md#UsersSignupPost) | **Post** /users/signup | 

# **UsersFoldersFilesFileDelete**
> UsersFoldersFilesFileDelete(ctx, )


deletes user's file (either input or output) from the database and removes its data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersFoldersFilesFileGet**
> UsersFoldersFilesFileGet(ctx, )


fetches file's info from the database and displays them along with its data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersFoldersFilesFilePost**
> UsersFoldersFilesFilePost(ctx, optional)


sorts the displayed input file and saves result in a newly created output file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUsersFoldersFilesFilePostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersFoldersFilesFilePostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of OutputFile**](OutputFile.md)| contains user&#x27;s output file name, and the column to sort the currently displayed input file by. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersFoldersFilesFilePut**
> UsersFoldersFilesFilePut(ctx, optional)


modifies user's input file (file name, data) and updates the fields in the database along with the file's data on the disk.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUsersFoldersFilesFilePutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersFoldersFilesFilePutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InputFile**](InputFile.md)| contains both file name and data or one of them. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersFoldersFilesGet**
> UsersFoldersFilesGet(ctx, )


fetches the names of the files belonging to the chosen folder and displays them in a list.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersFoldersGet**
> UsersFoldersGet(ctx, )


displays a list of public folder, user's input and output folders. each folder contains csv files, public folder holds server csv input files, and user's input and output hold user's csv input and output files respectively.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersFoldersPost**
> UsersFoldersPost(ctx, optional)


creates a new user's input file.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUsersFoldersPostOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersFoldersPostOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of InputFile**](InputFile.md)| contains file name and file data (comma separated values). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersLoginPost**
> UsersLoginPost(ctx, body)


posts user credentials to server to check if the user exists in the database to log them in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Login**](Login.md)| contains user&#x27;s credentials (username and password). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSettingsDelete**
> UsersSettingsDelete(ctx, )


deletes user from the database and removes their account and data.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSettingsGet**
> UsersSettingsGet(ctx, )


fetches user's insensitive info (username, email) from the database and displays them.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSettingsPut**
> UsersSettingsPut(ctx, optional)


modifies user's info (username, email, password) and updates the fields in the database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiUsersSettingsPutOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiUsersSettingsPutOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of Settings**](Settings.md)| contains either all or some updated user info (username, email, password). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsersSignupPost**
> UsersSignupPost(ctx, body)


posts new user's info to server to create a new account for them in the database.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Signup**](Signup.md)| contains new user&#x27;s info (username, email, password, confirm_password). | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)


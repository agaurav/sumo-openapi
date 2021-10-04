# \FolderManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFolder**](FolderManagementApi.md#CreateFolder) | **Post** /v2/content/folders | Create a new folder.
[**GetAdminRecommendedFolderAsync**](FolderManagementApi.md#GetAdminRecommendedFolderAsync) | **Get** /v2/content/folders/adminRecommended | Schedule Admin Recommended folder job
[**GetAdminRecommendedFolderAsyncResult**](FolderManagementApi.md#GetAdminRecommendedFolderAsyncResult) | **Get** /v2/content/folders/adminRecommended/{jobId}/result | Get Admin Recommended folder job result
[**GetAdminRecommendedFolderAsyncStatus**](FolderManagementApi.md#GetAdminRecommendedFolderAsyncStatus) | **Get** /v2/content/folders/adminRecommended/{jobId}/status | Get Admin Recommended folder job status
[**GetFolder**](FolderManagementApi.md#GetFolder) | **Get** /v2/content/folders/{id} | Get a folder.
[**GetGlobalFolderAsync**](FolderManagementApi.md#GetGlobalFolderAsync) | **Get** /v2/content/folders/global | Schedule Global View job
[**GetGlobalFolderAsyncResult**](FolderManagementApi.md#GetGlobalFolderAsyncResult) | **Get** /v2/content/folders/global/{jobId}/result | Get Global View job result
[**GetGlobalFolderAsyncStatus**](FolderManagementApi.md#GetGlobalFolderAsyncStatus) | **Get** /v2/content/folders/global/{jobId}/status | Get Global View job status
[**GetPersonalFolder**](FolderManagementApi.md#GetPersonalFolder) | **Get** /v2/content/folders/personal | Get personal folder.
[**UpdateFolder**](FolderManagementApi.md#UpdateFolder) | **Put** /v2/content/folders/{id} | Update a folder.



## CreateFolder

> Folder CreateFolder(ctx).FolderDefinition(folderDefinition).IsAdminMode(isAdminMode).Execute()

Create a new folder.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    folderDefinition := *openapiclient.NewFolderDefinition("SampleFolder", "ParentId_example") // FolderDefinition | Information about the new folder.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.CreateFolder(context.Background()).FolderDefinition(folderDefinition).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.CreateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.CreateFolder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **folderDefinition** | [**FolderDefinition**](FolderDefinition.md) | Information about the new folder. | 
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminRecommendedFolderAsync

> BeginAsyncJobResponse GetAdminRecommendedFolderAsync(ctx).IsAdminMode(isAdminMode).Execute()

Schedule Admin Recommended folder job



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetAdminRecommendedFolderAsync(context.Background()).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetAdminRecommendedFolderAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminRecommendedFolderAsync`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetAdminRecommendedFolderAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRecommendedFolderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminRecommendedFolderAsyncResult

> Folder GetAdminRecommendedFolderAsyncResult(ctx, jobId).Execute()

Get Admin Recommended folder job result



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := "jobId_example" // string | The identifier of the asynchronous Admin Recommended folder job.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetAdminRecommendedFolderAsyncResult(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetAdminRecommendedFolderAsyncResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminRecommendedFolderAsyncResult`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetAdminRecommendedFolderAsyncResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Admin Recommended folder job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRecommendedFolderAsyncResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdminRecommendedFolderAsyncStatus

> AsyncJobStatus GetAdminRecommendedFolderAsyncStatus(ctx, jobId).Execute()

Get Admin Recommended folder job status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := "jobId_example" // string | The identifier of the asynchronous Admin Recommended folder job.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetAdminRecommendedFolderAsyncStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetAdminRecommendedFolderAsyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdminRecommendedFolderAsyncStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetAdminRecommendedFolderAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Admin Recommended folder job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdminRecommendedFolderAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFolder

> Folder GetFolder(ctx, id).IsAdminMode(isAdminMode).Execute()

Get a folder.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Identifier of the folder to fetch.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetFolder(context.Background(), id).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalFolderAsync

> BeginAsyncJobResponse GetGlobalFolderAsync(ctx).IsAdminMode(isAdminMode).Execute()

Schedule Global View job



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetGlobalFolderAsync(context.Background()).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetGlobalFolderAsync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalFolderAsync`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetGlobalFolderAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalFolderAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalFolderAsyncResult

> ContentList GetGlobalFolderAsyncResult(ctx, jobId).Execute()

Get Global View job result



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := "jobId_example" // string | The identifier of the asynchronous Global View job.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetGlobalFolderAsyncResult(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetGlobalFolderAsyncResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalFolderAsyncResult`: ContentList
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetGlobalFolderAsyncResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Global View job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalFolderAsyncResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentList**](ContentList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalFolderAsyncStatus

> AsyncJobStatus GetGlobalFolderAsyncStatus(ctx, jobId).Execute()

Get Global View job status



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    jobId := "jobId_example" // string | The identifier of the asynchronous Global View job.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetGlobalFolderAsyncStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetGlobalFolderAsyncStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalFolderAsyncStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetGlobalFolderAsyncStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The identifier of the asynchronous Global View job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalFolderAsyncStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncJobStatus**](AsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPersonalFolder

> Folder GetPersonalFolder(ctx).Execute()

Get personal folder.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.GetPersonalFolder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.GetPersonalFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPersonalFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.GetPersonalFolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPersonalFolderRequest struct via the builder pattern


### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFolder

> Folder UpdateFolder(ctx, id).UpdateFolderRequest(updateFolderRequest).IsAdminMode(isAdminMode).Execute()

Update a folder.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | Identifier of the folder to update.
    updateFolderRequest := *openapiclient.NewUpdateFolderRequest("SampleFolder") // UpdateFolderRequest | Information to update about the folder.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FolderManagementApi.UpdateFolder(context.Background(), id).UpdateFolderRequest(updateFolderRequest).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FolderManagementApi.UpdateFolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFolder`: Folder
    fmt.Fprintf(os.Stdout, "Response from `FolderManagementApi.UpdateFolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the folder to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFolderRequest** | [**UpdateFolderRequest**](UpdateFolderRequest.md) | Information to update about the folder. | 
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**Folder**](Folder.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


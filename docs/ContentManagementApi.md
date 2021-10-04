# \ContentManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AsyncCopyStatus**](ContentManagementApi.md#AsyncCopyStatus) | **Get** /v2/content/{id}/copy/{jobId}/status | Content copy job status.
[**BeginAsyncCopy**](ContentManagementApi.md#BeginAsyncCopy) | **Post** /v2/content/{id}/copy | Start a content copy job.
[**BeginAsyncDelete**](ContentManagementApi.md#BeginAsyncDelete) | **Delete** /v2/content/{id}/delete | Start a content deletion job.
[**BeginAsyncExport**](ContentManagementApi.md#BeginAsyncExport) | **Post** /v2/content/{id}/export | Start a content export job.
[**BeginAsyncImport**](ContentManagementApi.md#BeginAsyncImport) | **Post** /v2/content/folders/{folderId}/import | Start a content import job.
[**GetAsyncDeleteStatus**](ContentManagementApi.md#GetAsyncDeleteStatus) | **Get** /v2/content/{id}/delete/{jobId}/status | Content deletion job status.
[**GetAsyncExportResult**](ContentManagementApi.md#GetAsyncExportResult) | **Get** /v2/content/{contentId}/export/{jobId}/result | Content export job result.
[**GetAsyncExportStatus**](ContentManagementApi.md#GetAsyncExportStatus) | **Get** /v2/content/{contentId}/export/{jobId}/status | Content export job status.
[**GetAsyncImportStatus**](ContentManagementApi.md#GetAsyncImportStatus) | **Get** /v2/content/folders/{folderId}/import/{jobId}/status | Content import job status.
[**GetItemByPath**](ContentManagementApi.md#GetItemByPath) | **Get** /v2/content/path | Get content item by path.
[**GetPathById**](ContentManagementApi.md#GetPathById) | **Get** /v2/content/{contentId}/path | Get path of an item.
[**MoveItem**](ContentManagementApi.md#MoveItem) | **Post** /v2/content/{id}/move | Move an item.



## AsyncCopyStatus

> AsyncJobStatus AsyncCopyStatus(ctx, id, jobId).IsAdminMode(isAdminMode).Execute()

Content copy job status.



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
    id := "id_example" // string | The identifier of the content which was copied.
    jobId := "jobId_example" // string | The identifier of the asynchronous copy request job.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.AsyncCopyStatus(context.Background(), id, jobId).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.AsyncCopyStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AsyncCopyStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.AsyncCopyStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the content which was copied. | 
**jobId** | **string** | The identifier of the asynchronous copy request job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAsyncCopyStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

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


## BeginAsyncCopy

> BeginAsyncJobResponse BeginAsyncCopy(ctx, id).DestinationFolder(destinationFolder).IsAdminMode(isAdminMode).Execute()

Start a content copy job.



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
    id := "id_example" // string | The identifier of the content item to copy. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
    destinationFolder := "destinationFolder_example" // string | The identifier of the destination folder.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.BeginAsyncCopy(context.Background(), id).DestinationFolder(destinationFolder).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.BeginAsyncCopy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BeginAsyncCopy`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.BeginAsyncCopy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the content item to copy. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBeginAsyncCopyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationFolder** | **string** | The identifier of the destination folder. | 
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


## BeginAsyncDelete

> BeginAsyncJobResponse BeginAsyncDelete(ctx, id).IsAdminMode(isAdminMode).Execute()

Start a content deletion job.



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
    id := "id_example" // string | Identifier of the content to delete. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.BeginAsyncDelete(context.Background(), id).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.BeginAsyncDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BeginAsyncDelete`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.BeginAsyncDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the content to delete. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBeginAsyncDeleteRequest struct via the builder pattern


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


## BeginAsyncExport

> BeginAsyncJobResponse BeginAsyncExport(ctx, id).IsAdminMode(isAdminMode).Execute()

Start a content export job.



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
    id := "id_example" // string | The identifier of the content item to export. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.BeginAsyncExport(context.Background(), id).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.BeginAsyncExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BeginAsyncExport`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.BeginAsyncExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the content item to export. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBeginAsyncExportRequest struct via the builder pattern


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


## BeginAsyncImport

> BeginAsyncJobResponse BeginAsyncImport(ctx, folderId).ContentSyncDefinition(contentSyncDefinition).IsAdminMode(isAdminMode).Overwrite(overwrite).Execute()

Start a content import job.



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
    folderId := "folderId_example" // string | The identifier of the folder to import into. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format.
    contentSyncDefinition := *openapiclient.NewContentSyncDefinition("Type_example", "Name_example") // ContentSyncDefinition | The content to import.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)
    overwrite := true // bool | Set this to \"true\" to overwrite a content item if the name already exists. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.BeginAsyncImport(context.Background(), folderId).ContentSyncDefinition(contentSyncDefinition).IsAdminMode(isAdminMode).Overwrite(overwrite).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.BeginAsyncImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BeginAsyncImport`: BeginAsyncJobResponse
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.BeginAsyncImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The identifier of the folder to import into. Identifiers from the Library in the Sumo user interface are provided in decimal format which is incompatible with this API. The identifier needs to be in hexadecimal format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBeginAsyncImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentSyncDefinition** | [**ContentSyncDefinition**](ContentSyncDefinition.md) | The content to import. | 
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 
 **overwrite** | **bool** | Set this to \&quot;true\&quot; to overwrite a content item if the name already exists. | [default to false]

### Return type

[**BeginAsyncJobResponse**](BeginAsyncJobResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncDeleteStatus

> AsyncJobStatus GetAsyncDeleteStatus(ctx, id, jobId).IsAdminMode(isAdminMode).Execute()

Content deletion job status.



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
    id := "id_example" // string | Identifier of the content to delete.
    jobId := "jobId_example" // string | The identifier of the asynchronous job.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.GetAsyncDeleteStatus(context.Background(), id, jobId).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.GetAsyncDeleteStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncDeleteStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.GetAsyncDeleteStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the content to delete. | 
**jobId** | **string** | The identifier of the asynchronous job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncDeleteStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

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


## GetAsyncExportResult

> ContentSyncDefinition GetAsyncExportResult(ctx, contentId, jobId).IsAdminMode(isAdminMode).Execute()

Content export job result.



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
    contentId := "contentId_example" // string | The identifier of the exported content item.
    jobId := "jobId_example" // string | The identifier of the asynchronous job.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.GetAsyncExportResult(context.Background(), contentId, jobId).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.GetAsyncExportResult``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncExportResult`: ContentSyncDefinition
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.GetAsyncExportResult`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string** | The identifier of the exported content item. | 
**jobId** | **string** | The identifier of the asynchronous job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncExportResultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentSyncDefinition**](ContentSyncDefinition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAsyncExportStatus

> AsyncJobStatus GetAsyncExportStatus(ctx, contentId, jobId).IsAdminMode(isAdminMode).Execute()

Content export job status.



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
    contentId := "contentId_example" // string | The identifier of the exported content item.
    jobId := "jobId_example" // string | The identifier of the asynchronous export job.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.GetAsyncExportStatus(context.Background(), contentId, jobId).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.GetAsyncExportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncExportStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.GetAsyncExportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string** | The identifier of the exported content item. | 
**jobId** | **string** | The identifier of the asynchronous export job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncExportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

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


## GetAsyncImportStatus

> AsyncJobStatus GetAsyncImportStatus(ctx, folderId, jobId).IsAdminMode(isAdminMode).Execute()

Content import job status.



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
    folderId := "folderId_example" // string | The identifier of the folder to import into.
    jobId := "jobId_example" // string | The identifier of the import request.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.GetAsyncImportStatus(context.Background(), folderId, jobId).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.GetAsyncImportStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAsyncImportStatus`: AsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.GetAsyncImportStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**folderId** | **string** | The identifier of the folder to import into. | 
**jobId** | **string** | The identifier of the import request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAsyncImportStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

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


## GetItemByPath

> Content GetItemByPath(ctx).Path(path).Execute()

Get content item by path.



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
    path := "/Library/Users/user@sumo.com/SampleFolder" // string | Path of the content item to retrieve.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.GetItemByPath(context.Background()).Path(path).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.GetItemByPath``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetItemByPath`: Content
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.GetItemByPath`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetItemByPathRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **path** | **string** | Path of the content item to retrieve. | 

### Return type

[**Content**](Content.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPathById

> ContentPath GetPathById(ctx, contentId).Execute()

Get path of an item.



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
    contentId := "contentId_example" // string | Identifier of the content item to get the path.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.GetPathById(context.Background(), contentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.GetPathById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPathById`: ContentPath
    fmt.Fprintf(os.Stdout, "Response from `ContentManagementApi.GetPathById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contentId** | **string** | Identifier of the content item to get the path. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPathByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContentPath**](ContentPath.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveItem

> MoveItem(ctx, id).DestinationFolderId(destinationFolderId).IsAdminMode(isAdminMode).Execute()

Move an item.



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
    destinationFolderId := "destinationFolderId_example" // string | Identifier of the destination folder.
    id := "id_example" // string | Identifier of the item the user wants to move.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentManagementApi.MoveItem(context.Background(), id).DestinationFolderId(destinationFolderId).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentManagementApi.MoveItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the item the user wants to move. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **destinationFolderId** | **string** | Identifier of the destination folder. | 

 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


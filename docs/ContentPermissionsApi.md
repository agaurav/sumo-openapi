# \ContentPermissionsApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddContentPermissions**](ContentPermissionsApi.md#AddContentPermissions) | **Put** /v2/content/{id}/permissions/add | Add permissions to a content item.
[**GetContentPermissions**](ContentPermissionsApi.md#GetContentPermissions) | **Get** /v2/content/{id}/permissions | Get permissions of a content item
[**RemoveContentPermissions**](ContentPermissionsApi.md#RemoveContentPermissions) | **Put** /v2/content/{id}/permissions/remove | Remove permissions from a content item.



## AddContentPermissions

> ContentPermissionResult AddContentPermissions(ctx, id).ContentPermissionUpdateRequest(contentPermissionUpdateRequest).IsAdminMode(isAdminMode).Execute()

Add permissions to a content item.



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
    id := "id_example" // string | The identifier of the content item.
    contentPermissionUpdateRequest := *openapiclient.NewContentPermissionUpdateRequest([]openapiclient.ContentPermissionAssignment{*openapiclient.NewContentPermissionAssignment("PermissionName_example", "role", "SourceId_example", "ContentId_example")}, false, "NotificationMessage_example") // ContentPermissionUpdateRequest | New permissions to add to the content item with the given identifier.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPermissionsApi.AddContentPermissions(context.Background(), id).ContentPermissionUpdateRequest(contentPermissionUpdateRequest).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPermissionsApi.AddContentPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddContentPermissions`: ContentPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `ContentPermissionsApi.AddContentPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the content item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddContentPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentPermissionUpdateRequest** | [**ContentPermissionUpdateRequest**](ContentPermissionUpdateRequest.md) | New permissions to add to the content item with the given identifier. | 
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentPermissionResult**](ContentPermissionResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContentPermissions

> ContentPermissionResult GetContentPermissions(ctx, id).ExplicitOnly(explicitOnly).IsAdminMode(isAdminMode).Execute()

Get permissions of a content item



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
    id := "id_example" // string | The identifier of the content item.
    explicitOnly := true // bool | There are two permission types: explicit and implicit. Permissions specifically assigned to the content item are explicit. Permissions derived from a parent content item, like a folder are implicit. To return only explicit permissions set this to true. (optional) (default to false)
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPermissionsApi.GetContentPermissions(context.Background(), id).ExplicitOnly(explicitOnly).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPermissionsApi.GetContentPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetContentPermissions`: ContentPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `ContentPermissionsApi.GetContentPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the content item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContentPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **explicitOnly** | **bool** | There are two permission types: explicit and implicit. Permissions specifically assigned to the content item are explicit. Permissions derived from a parent content item, like a folder are implicit. To return only explicit permissions set this to true. | [default to false]
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentPermissionResult**](ContentPermissionResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveContentPermissions

> ContentPermissionResult RemoveContentPermissions(ctx, id).ContentPermissionUpdateRequest(contentPermissionUpdateRequest).IsAdminMode(isAdminMode).Execute()

Remove permissions from a content item.



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
    id := "id_example" // string | The identifier of the content item.
    contentPermissionUpdateRequest := *openapiclient.NewContentPermissionUpdateRequest([]openapiclient.ContentPermissionAssignment{*openapiclient.NewContentPermissionAssignment("PermissionName_example", "role", "SourceId_example", "ContentId_example")}, false, "NotificationMessage_example") // ContentPermissionUpdateRequest | Permissions to remove from a content item with the given identifier.
    isAdminMode := "isAdminMode_example" // string | Set this to \"true\" if you want to perform the request as a Content Administrator. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ContentPermissionsApi.RemoveContentPermissions(context.Background(), id).ContentPermissionUpdateRequest(contentPermissionUpdateRequest).IsAdminMode(isAdminMode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ContentPermissionsApi.RemoveContentPermissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveContentPermissions`: ContentPermissionResult
    fmt.Fprintf(os.Stdout, "Response from `ContentPermissionsApi.RemoveContentPermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The identifier of the content item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveContentPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentPermissionUpdateRequest** | [**ContentPermissionUpdateRequest**](ContentPermissionUpdateRequest.md) | Permissions to remove from a content item with the given identifier. | 
 **isAdminMode** | **string** | Set this to \&quot;true\&quot; if you want to perform the request as a Content Administrator. | 

### Return type

[**ContentPermissionResult**](ContentPermissionResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \AccessKeyManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessKey**](AccessKeyManagementApi.md#CreateAccessKey) | **Post** /v1/accessKeys | Create an access key.
[**DeleteAccessKey**](AccessKeyManagementApi.md#DeleteAccessKey) | **Delete** /v1/accessKeys/{id} | Delete an access key.
[**ListAccessKeys**](AccessKeyManagementApi.md#ListAccessKeys) | **Get** /v1/accessKeys | List all access keys.
[**ListPersonalAccessKeys**](AccessKeyManagementApi.md#ListPersonalAccessKeys) | **Get** /v1/accessKeys/personal | List personal keys.
[**UpdateAccessKey**](AccessKeyManagementApi.md#UpdateAccessKey) | **Put** /v1/accessKeys/{id} | Update an access key.



## CreateAccessKey

> AccessKey CreateAccessKey(ctx).AccessKeyCreateRequest(accessKeyCreateRequest).Execute()

Create an access key.



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
    accessKeyCreateRequest := *openapiclient.NewAccessKeyCreateRequest("automation access key") // AccessKeyCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyManagementApi.CreateAccessKey(context.Background()).AccessKeyCreateRequest(accessKeyCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyManagementApi.CreateAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessKey`: AccessKey
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyManagementApi.CreateAccessKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessKeyCreateRequest** | [**AccessKeyCreateRequest**](AccessKeyCreateRequest.md) |  | 

### Return type

[**AccessKey**](AccessKey.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessKey

> DeleteAccessKey(ctx, id).Execute()

Delete an access key.



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
    id := "id_example" // string | The accessId of the access key to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyManagementApi.DeleteAccessKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyManagementApi.DeleteAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The accessId of the access key to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListAccessKeys

> PaginatedListAccessKeysResult ListAccessKeys(ctx).Limit(limit).Token(token).Execute()

List all access keys.



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
    limit := int32(56) // int32 | Limit the number of access keys returned in the response. The number of access keys returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyManagementApi.ListAccessKeys(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyManagementApi.ListAccessKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAccessKeys`: PaginatedListAccessKeysResult
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyManagementApi.ListAccessKeys`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of access keys returned in the response. The number of access keys returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**PaginatedListAccessKeysResult**](PaginatedListAccessKeysResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPersonalAccessKeys

> ListAccessKeysResult ListPersonalAccessKeys(ctx).Execute()

List personal keys.



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
    resp, r, err := api_client.AccessKeyManagementApi.ListPersonalAccessKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyManagementApi.ListPersonalAccessKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPersonalAccessKeys`: ListAccessKeysResult
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyManagementApi.ListPersonalAccessKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPersonalAccessKeysRequest struct via the builder pattern


### Return type

[**ListAccessKeysResult**](ListAccessKeysResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAccessKey

> AccessKeyPublic UpdateAccessKey(ctx, id).AccessKeyUpdateRequest(accessKeyUpdateRequest).Execute()

Update an access key.



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
    id := "id_example" // string | The accessId of the access key to update.
    accessKeyUpdateRequest := *openapiclient.NewAccessKeyUpdateRequest(true) // AccessKeyUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccessKeyManagementApi.UpdateAccessKey(context.Background(), id).AccessKeyUpdateRequest(accessKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccessKeyManagementApi.UpdateAccessKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAccessKey`: AccessKeyPublic
    fmt.Fprintf(os.Stdout, "Response from `AccessKeyManagementApi.UpdateAccessKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The accessId of the access key to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accessKeyUpdateRequest** | [**AccessKeyUpdateRequest**](AccessKeyUpdateRequest.md) |  | 

### Return type

[**AccessKeyPublic**](AccessKeyPublic.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


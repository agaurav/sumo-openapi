# \TokensLibraryManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToken**](TokensLibraryManagementApi.md#CreateToken) | **Post** /v1/tokens | Create a token.
[**DeleteToken**](TokensLibraryManagementApi.md#DeleteToken) | **Delete** /v1/tokens/{id} | Delete a token.
[**GetToken**](TokensLibraryManagementApi.md#GetToken) | **Get** /v1/tokens/{id} | Get a token.
[**ListTokens**](TokensLibraryManagementApi.md#ListTokens) | **Get** /v1/tokens | Get a list of tokens.
[**UpdateToken**](TokensLibraryManagementApi.md#UpdateToken) | **Put** /v1/tokens/{id} | Update a token.



## CreateToken

> TokenBaseResponse CreateToken(ctx).TokenBaseDefinition(tokenBaseDefinition).Execute()

Create a token.



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
    tokenBaseDefinition := *openapiclient.NewTokenBaseDefinition("token-name", "Active", "CollectorRegistration") // TokenBaseDefinition | Information about the token to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensLibraryManagementApi.CreateToken(context.Background()).TokenBaseDefinition(tokenBaseDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensLibraryManagementApi.CreateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToken`: TokenBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensLibraryManagementApi.CreateToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tokenBaseDefinition** | [**TokenBaseDefinition**](TokenBaseDefinition.md) | Information about the token to create. | 

### Return type

[**TokenBaseResponse**](TokenBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteToken

> DeleteToken(ctx, id).Execute()

Delete a token.



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
    id := "id_example" // string | Identifier of the token to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensLibraryManagementApi.DeleteToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensLibraryManagementApi.DeleteToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the token to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTokenRequest struct via the builder pattern


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


## GetToken

> TokenBaseResponse GetToken(ctx, id).Execute()

Get a token.



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
    id := "id_example" // string | Identifier of the token to return.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensLibraryManagementApi.GetToken(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensLibraryManagementApi.GetToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToken`: TokenBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensLibraryManagementApi.GetToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the token to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TokenBaseResponse**](TokenBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTokens

> ListTokensBaseResponse ListTokens(ctx).Execute()

Get a list of tokens.



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
    resp, r, err := api_client.TokensLibraryManagementApi.ListTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensLibraryManagementApi.ListTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTokens`: ListTokensBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensLibraryManagementApi.ListTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTokensRequest struct via the builder pattern


### Return type

[**ListTokensBaseResponse**](ListTokensBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToken

> TokenBaseResponse UpdateToken(ctx, id).TokenBaseDefinitionUpdate(tokenBaseDefinitionUpdate).Execute()

Update a token.



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
    id := "id_example" // string | Identifier of the token to update.
    tokenBaseDefinitionUpdate := *openapiclient.NewTokenBaseDefinitionUpdate("token-name", "Active", "CollectorRegistration", int64(123)) // TokenBaseDefinitionUpdate | The token to update.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TokensLibraryManagementApi.UpdateToken(context.Background(), id).TokenBaseDefinitionUpdate(tokenBaseDefinitionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TokensLibraryManagementApi.UpdateToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateToken`: TokenBaseResponse
    fmt.Fprintf(os.Stdout, "Response from `TokensLibraryManagementApi.UpdateToken`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the token to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tokenBaseDefinitionUpdate** | [**TokenBaseDefinitionUpdate**](TokenBaseDefinitionUpdate.md) | The token to update. | 

### Return type

[**TokenBaseResponse**](TokenBaseResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ConnectionManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnection**](ConnectionManagementApi.md#CreateConnection) | **Post** /v1/connections | Create a new connection.
[**DeleteConnection**](ConnectionManagementApi.md#DeleteConnection) | **Delete** /v1/connections/{id} | Delete a connection.
[**GetConnection**](ConnectionManagementApi.md#GetConnection) | **Get** /v1/connections/{id} | Get a connection.
[**ListConnections**](ConnectionManagementApi.md#ListConnections) | **Get** /v1/connections | Get a list of connections.
[**TestConnection**](ConnectionManagementApi.md#TestConnection) | **Post** /v1/connections/test | Test a new connection url.
[**UpdateConnection**](ConnectionManagementApi.md#UpdateConnection) | **Put** /v1/connections/{id} | Update a connection.



## CreateConnection

> Connection CreateConnection(ctx).ConnectionDefinition(connectionDefinition).Execute()

Create a new connection.



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
    connectionDefinition := *openapiclient.NewConnectionDefinition("Type_example", "Name_example") // ConnectionDefinition | Information about the new connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionManagementApi.CreateConnection(context.Background()).ConnectionDefinition(connectionDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementApi.CreateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConnection`: Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementApi.CreateConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionDefinition** | [**ConnectionDefinition**](ConnectionDefinition.md) | Information about the new connection. | 

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnection

> DeleteConnection(ctx, id).Type_(type_).Execute()

Delete a connection.



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
    id := "id_example" // string | Identifier of the connection to delete.
    type_ := "type__example" // string | Type of connection to delete. Valid values are `WebhookConnection`, `ServiceNowConnection`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionManagementApi.DeleteConnection(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementApi.DeleteConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the connection to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of connection to delete. Valid values are &#x60;WebhookConnection&#x60;, &#x60;ServiceNowConnection&#x60;. | 

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


## GetConnection

> Connection GetConnection(ctx, id).Type_(type_).Execute()

Get a connection.



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
    id := "id_example" // string | Identifier of connection to return.
    type_ := "type__example" // string | Type of connection to return. Valid values are `WebhookConnection`, `ServiceNowConnection`. (default to "WebhookConnection")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionManagementApi.GetConnection(context.Background(), id).Type_(type_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementApi.GetConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnection`: Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementApi.GetConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of connection to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **string** | Type of connection to return. Valid values are &#x60;WebhookConnection&#x60;, &#x60;ServiceNowConnection&#x60;. | [default to &quot;WebhookConnection&quot;]

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnections

> ListConnectionsResponse ListConnections(ctx).Limit(limit).Token(token).Execute()

Get a list of connections.



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
    limit := int32(56) // int32 | Limit the number of connections returned in the response. The number of connections returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionManagementApi.ListConnections(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementApi.ListConnections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnections`: ListConnectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementApi.ListConnections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of connections returned in the response. The number of connections returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListConnectionsResponse**](ListConnectionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestConnection

> TestConnectionResponse TestConnection(ctx).ConnectionDefinition(connectionDefinition).Execute()

Test a new connection url.



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
    connectionDefinition := *openapiclient.NewConnectionDefinition("Type_example", "Name_example") // ConnectionDefinition | Information about the new connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionManagementApi.TestConnection(context.Background()).ConnectionDefinition(connectionDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementApi.TestConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestConnection`: TestConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementApi.TestConnection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectionDefinition** | [**ConnectionDefinition**](ConnectionDefinition.md) | Information about the new connection. | 

### Return type

[**TestConnectionResponse**](TestConnectionResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateConnection

> Connection UpdateConnection(ctx, id).ConnectionDefinition(connectionDefinition).Execute()

Update a connection.



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
    id := "id_example" // string | Identifier of the connection to update.
    connectionDefinition := *openapiclient.NewConnectionDefinition("Type_example", "Name_example") // ConnectionDefinition | Information to update about the connection.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ConnectionManagementApi.UpdateConnection(context.Background(), id).ConnectionDefinition(connectionDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConnectionManagementApi.UpdateConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateConnection`: Connection
    fmt.Fprintf(os.Stdout, "Response from `ConnectionManagementApi.UpdateConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the connection to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **connectionDefinition** | [**ConnectionDefinition**](ConnectionDefinition.md) | Information to update about the connection. | 

### Return type

[**Connection**](Connection.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \ScheduledViewManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScheduledView**](ScheduledViewManagementApi.md#CreateScheduledView) | **Post** /v1/scheduledViews | Create a new scheduled view.
[**DisableScheduledView**](ScheduledViewManagementApi.md#DisableScheduledView) | **Delete** /v1/scheduledViews/{id}/disable | Disable a scheduled view.
[**GetScheduledView**](ScheduledViewManagementApi.md#GetScheduledView) | **Get** /v1/scheduledViews/{id} | Get a scheduled view.
[**ListScheduledViews**](ScheduledViewManagementApi.md#ListScheduledViews) | **Get** /v1/scheduledViews | Get a list of scheduled views.
[**PauseScheduledView**](ScheduledViewManagementApi.md#PauseScheduledView) | **Post** /v1/scheduledViews/{id}/pause | Pause a scheduled view.
[**StartScheduledView**](ScheduledViewManagementApi.md#StartScheduledView) | **Post** /v1/scheduledViews/{id}/start | Start a scheduled view.
[**UpdateScheduledView**](ScheduledViewManagementApi.md#UpdateScheduledView) | **Put** /v1/scheduledViews/{id} | Update a scheduled view.



## CreateScheduledView

> ScheduledView CreateScheduledView(ctx).CreateScheduledViewDefinition(createScheduledViewDefinition).Execute()

Create a new scheduled view.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    createScheduledViewDefinition := *openapiclient.NewCreateScheduledViewDefinition("_sourcecategory=*/Apache", "TestScheduledView", time.Now()) // CreateScheduledViewDefinition | Information about the new scheduled view.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledViewManagementApi.CreateScheduledView(context.Background()).CreateScheduledViewDefinition(createScheduledViewDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledViewManagementApi.CreateScheduledView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateScheduledView`: ScheduledView
    fmt.Fprintf(os.Stdout, "Response from `ScheduledViewManagementApi.CreateScheduledView`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScheduledViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createScheduledViewDefinition** | [**CreateScheduledViewDefinition**](CreateScheduledViewDefinition.md) | Information about the new scheduled view. | 

### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableScheduledView

> DisableScheduledView(ctx, id).Execute()

Disable a scheduled view.



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
    id := "id_example" // string | Identifier of the scheduled view to disable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledViewManagementApi.DisableScheduledView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledViewManagementApi.DisableScheduledView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the scheduled view to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableScheduledViewRequest struct via the builder pattern


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


## GetScheduledView

> ScheduledView GetScheduledView(ctx, id).Execute()

Get a scheduled view.



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
    id := "id_example" // string | Identifier of the scheduled view to fetch.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledViewManagementApi.GetScheduledView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledViewManagementApi.GetScheduledView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetScheduledView`: ScheduledView
    fmt.Fprintf(os.Stdout, "Response from `ScheduledViewManagementApi.GetScheduledView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the scheduled view to fetch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScheduledViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListScheduledViews

> ListScheduledViewsResponse ListScheduledViews(ctx).Limit(limit).Token(token).Execute()

Get a list of scheduled views.



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
    limit := int32(56) // int32 | Limit the number of scheduled views returned in the response. The number of scheduled views returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledViewManagementApi.ListScheduledViews(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledViewManagementApi.ListScheduledViews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScheduledViews`: ListScheduledViewsResponse
    fmt.Fprintf(os.Stdout, "Response from `ScheduledViewManagementApi.ListScheduledViews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScheduledViewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of scheduled views returned in the response. The number of scheduled views returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListScheduledViewsResponse**](ListScheduledViewsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseScheduledView

> ScheduledView PauseScheduledView(ctx, id).Execute()

Pause a scheduled view.



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
    id := "id_example" // string | Identifier of the scheduled view to pause.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledViewManagementApi.PauseScheduledView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledViewManagementApi.PauseScheduledView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PauseScheduledView`: ScheduledView
    fmt.Fprintf(os.Stdout, "Response from `ScheduledViewManagementApi.PauseScheduledView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the scheduled view to pause. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseScheduledViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StartScheduledView

> ScheduledView StartScheduledView(ctx, id).Execute()

Start a scheduled view.



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
    id := "id_example" // string | Identifier of the scheduled view to start.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledViewManagementApi.StartScheduledView(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledViewManagementApi.StartScheduledView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `StartScheduledView`: ScheduledView
    fmt.Fprintf(os.Stdout, "Response from `ScheduledViewManagementApi.StartScheduledView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the scheduled view to start. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStartScheduledViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScheduledView

> ScheduledView UpdateScheduledView(ctx, id).UpdateScheduledViewDefinition(updateScheduledViewDefinition).Execute()

Update a scheduled view.



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
    id := "id_example" // string | Identifier of the scheduled view to update.
    updateScheduledViewDefinition := *openapiclient.NewUpdateScheduledViewDefinition() // UpdateScheduledViewDefinition | Information to update about the scheduled view.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ScheduledViewManagementApi.UpdateScheduledView(context.Background(), id).UpdateScheduledViewDefinition(updateScheduledViewDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScheduledViewManagementApi.UpdateScheduledView``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateScheduledView`: ScheduledView
    fmt.Fprintf(os.Stdout, "Response from `ScheduledViewManagementApi.UpdateScheduledView`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the scheduled view to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScheduledViewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScheduledViewDefinition** | [**UpdateScheduledViewDefinition**](UpdateScheduledViewDefinition.md) | Information to update about the scheduled view. | 

### Return type

[**ScheduledView**](ScheduledView.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


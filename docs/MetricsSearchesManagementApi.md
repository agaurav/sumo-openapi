# \MetricsSearchesManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMetricsSearch**](MetricsSearchesManagementApi.md#CreateMetricsSearch) | **Post** /v1/metricsSearches | Save a metrics search.
[**DeleteMetricsSearch**](MetricsSearchesManagementApi.md#DeleteMetricsSearch) | **Delete** /v1/metricsSearches/{id} | Deletes a metrics search.
[**GetMetricsSearch**](MetricsSearchesManagementApi.md#GetMetricsSearch) | **Get** /v1/metricsSearches/{id} | Get a metrics search.
[**UpdateMetricsSearch**](MetricsSearchesManagementApi.md#UpdateMetricsSearch) | **Put** /v1/metricsSearches/{id} | Updates a metrics search.



## CreateMetricsSearch

> MetricsSearchInstance CreateMetricsSearch(ctx).SaveMetricsSearchRequest(saveMetricsSearchRequest).Execute()

Save a metrics search.



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
    saveMetricsSearchRequest := *openapiclient.NewSaveMetricsSearchRequest("Short title", "Long and detailed description", *openapiclient.NewResolvableTimeRange("Type_example"), []openapiclient.MetricsSearchQuery{*openapiclient.NewMetricsSearchQuery("A", "my_metric | avg")}, "000000000000001A") // SaveMetricsSearchRequest | The definition of the metrics search.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsSearchesManagementApi.CreateMetricsSearch(context.Background()).SaveMetricsSearchRequest(saveMetricsSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementApi.CreateMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMetricsSearch`: MetricsSearchInstance
    fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementApi.CreateMetricsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMetricsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **saveMetricsSearchRequest** | [**SaveMetricsSearchRequest**](SaveMetricsSearchRequest.md) | The definition of the metrics search. | 

### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMetricsSearch

> DeleteMetricsSearch(ctx, id).Execute()

Deletes a metrics search.



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
    id := "id_example" // string | Identifier of the metrics search.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsSearchesManagementApi.DeleteMetricsSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementApi.DeleteMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metrics search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMetricsSearchRequest struct via the builder pattern


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


## GetMetricsSearch

> MetricsSearchInstance GetMetricsSearch(ctx, id).Execute()

Get a metrics search.



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
    id := "id_example" // string | Identifier of the metrics search.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsSearchesManagementApi.GetMetricsSearch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementApi.GetMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMetricsSearch`: MetricsSearchInstance
    fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementApi.GetMetricsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metrics search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMetricsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMetricsSearch

> MetricsSearchInstance UpdateMetricsSearch(ctx, id).MetricsSearchV1(metricsSearchV1).Execute()

Updates a metrics search.



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
    id := "id_example" // string | Identifier of the metrics search.
    metricsSearchV1 := *openapiclient.NewMetricsSearchV1("Short title", "Long and detailed description", *openapiclient.NewResolvableTimeRange("Type_example"), []openapiclient.MetricsSearchQuery{*openapiclient.NewMetricsSearchQuery("A", "my_metric | avg")}) // MetricsSearchV1 | An updated metrics search definition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsSearchesManagementApi.UpdateMetricsSearch(context.Background(), id).MetricsSearchV1(metricsSearchV1).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsSearchesManagementApi.UpdateMetricsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMetricsSearch`: MetricsSearchInstance
    fmt.Fprintf(os.Stdout, "Response from `MetricsSearchesManagementApi.UpdateMetricsSearch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the metrics search. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMetricsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **metricsSearchV1** | [**MetricsSearchV1**](MetricsSearchV1.md) | An updated metrics search definition. | 

### Return type

[**MetricsSearchInstance**](MetricsSearchInstance.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


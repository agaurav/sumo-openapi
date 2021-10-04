# \MetricsQueryApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunMetricsQueries**](MetricsQueryApi.md#RunMetricsQueries) | **Post** /v1/metricsQueries | Run metrics queries



## RunMetricsQueries

> MetricsQueryResponse RunMetricsQueries(ctx).MetricsQueryRequest(metricsQueryRequest).Execute()

Run metrics queries



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
    metricsQueryRequest := *openapiclient.NewMetricsQueryRequest([]openapiclient.MetricsQueryRow{*openapiclient.NewMetricsQueryRow("A", "metric=CPU_Idle")}, *openapiclient.NewResolvableTimeRange("Type_example")) // MetricsQueryRequest | The parameters for the metrics query.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MetricsQueryApi.RunMetricsQueries(context.Background()).MetricsQueryRequest(metricsQueryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricsQueryApi.RunMetricsQueries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunMetricsQueries`: MetricsQueryResponse
    fmt.Fprintf(os.Stdout, "Response from `MetricsQueryApi.RunMetricsQueries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunMetricsQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **metricsQueryRequest** | [**MetricsQueryRequest**](MetricsQueryRequest.md) | The parameters for the metrics query. | 

### Return type

[**MetricsQueryResponse**](MetricsQueryResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


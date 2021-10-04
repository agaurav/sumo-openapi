# \IngestBudgetManagementV2Api

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIngestBudgetV2**](IngestBudgetManagementV2Api.md#CreateIngestBudgetV2) | **Post** /v2/ingestBudgets | Create a new ingest budget.
[**DeleteIngestBudgetV2**](IngestBudgetManagementV2Api.md#DeleteIngestBudgetV2) | **Delete** /v2/ingestBudgets/{id} | Delete an ingest budget.
[**GetIngestBudgetV2**](IngestBudgetManagementV2Api.md#GetIngestBudgetV2) | **Get** /v2/ingestBudgets/{id} | Get an ingest budget.
[**ListIngestBudgetsV2**](IngestBudgetManagementV2Api.md#ListIngestBudgetsV2) | **Get** /v2/ingestBudgets | Get a list of ingest budgets.
[**ResetUsageV2**](IngestBudgetManagementV2Api.md#ResetUsageV2) | **Post** /v2/ingestBudgets/{id}/usage/reset | Reset usage.
[**UpdateIngestBudgetV2**](IngestBudgetManagementV2Api.md#UpdateIngestBudgetV2) | **Put** /v2/ingestBudgets/{id} | Update an ingest budget.



## CreateIngestBudgetV2

> IngestBudgetV2 CreateIngestBudgetV2(ctx).IngestBudgetDefinitionV2(ingestBudgetDefinitionV2).Execute()

Create a new ingest budget.



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
    ingestBudgetDefinitionV2 := *openapiclient.NewIngestBudgetDefinitionV2("Developer Budget", "_sourceCategory=*prod*nginx*", int64(1000), "America/Los_Angeles", "23:30", "stopCollecting") // IngestBudgetDefinitionV2 | Information about the new ingest budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV2Api.CreateIngestBudgetV2(context.Background()).IngestBudgetDefinitionV2(ingestBudgetDefinitionV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV2Api.CreateIngestBudgetV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIngestBudgetV2`: IngestBudgetV2
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV2Api.CreateIngestBudgetV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestBudgetV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestBudgetDefinitionV2** | [**IngestBudgetDefinitionV2**](IngestBudgetDefinitionV2.md) | Information about the new ingest budget. | 

### Return type

[**IngestBudgetV2**](IngestBudgetV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestBudgetV2

> DeleteIngestBudgetV2(ctx, id).Execute()

Delete an ingest budget.



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
    id := "id_example" // string | Identifier of the ingest budget to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV2Api.DeleteIngestBudgetV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV2Api.DeleteIngestBudgetV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingest budget to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIngestBudgetV2Request struct via the builder pattern


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


## GetIngestBudgetV2

> IngestBudgetV2 GetIngestBudgetV2(ctx, id).Execute()

Get an ingest budget.



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
    id := "id_example" // string | Identifier of ingest budget to return.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV2Api.GetIngestBudgetV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV2Api.GetIngestBudgetV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIngestBudgetV2`: IngestBudgetV2
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV2Api.GetIngestBudgetV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of ingest budget to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestBudgetV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngestBudgetV2**](IngestBudgetV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIngestBudgetsV2

> ListIngestBudgetsResponseV2 ListIngestBudgetsV2(ctx).Limit(limit).Token(token).Execute()

Get a list of ingest budgets.



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
    limit := int32(56) // int32 | Limit the number of budgets returned in the response. The number of budgets returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV2Api.ListIngestBudgetsV2(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV2Api.ListIngestBudgetsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIngestBudgetsV2`: ListIngestBudgetsResponseV2
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV2Api.ListIngestBudgetsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIngestBudgetsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of budgets returned in the response. The number of budgets returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListIngestBudgetsResponseV2**](ListIngestBudgetsResponseV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetUsageV2

> ResetUsageV2(ctx, id).Execute()

Reset usage.



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
    id := "id_example" // string | Identifier of the ingest budget to reset usage.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV2Api.ResetUsageV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV2Api.ResetUsageV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingest budget to reset usage. | 

### Other Parameters

Other parameters are passed through a pointer to a apiResetUsageV2Request struct via the builder pattern


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


## UpdateIngestBudgetV2

> IngestBudgetV2 UpdateIngestBudgetV2(ctx, id).IngestBudgetDefinitionV2(ingestBudgetDefinitionV2).Execute()

Update an ingest budget.



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
    id := "id_example" // string | Identifier of the ingest budget to update.
    ingestBudgetDefinitionV2 := *openapiclient.NewIngestBudgetDefinitionV2("Developer Budget", "_sourceCategory=*prod*nginx*", int64(1000), "America/Los_Angeles", "23:30", "stopCollecting") // IngestBudgetDefinitionV2 | Information to update about the ingest budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV2Api.UpdateIngestBudgetV2(context.Background(), id).IngestBudgetDefinitionV2(ingestBudgetDefinitionV2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV2Api.UpdateIngestBudgetV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIngestBudgetV2`: IngestBudgetV2
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV2Api.UpdateIngestBudgetV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingest budget to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestBudgetV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ingestBudgetDefinitionV2** | [**IngestBudgetDefinitionV2**](IngestBudgetDefinitionV2.md) | Information to update about the ingest budget. | 

### Return type

[**IngestBudgetV2**](IngestBudgetV2.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \IngestBudgetManagementV1Api

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignCollectorToBudget**](IngestBudgetManagementV1Api.md#AssignCollectorToBudget) | **Put** /v1/ingestBudgets/{id}/collectors/{collectorId} | Assign a Collector to a budget.
[**CreateIngestBudget**](IngestBudgetManagementV1Api.md#CreateIngestBudget) | **Post** /v1/ingestBudgets | Create a new ingest budget.
[**DeleteIngestBudget**](IngestBudgetManagementV1Api.md#DeleteIngestBudget) | **Delete** /v1/ingestBudgets/{id} | Delete an ingest budget.
[**GetAssignedCollectors**](IngestBudgetManagementV1Api.md#GetAssignedCollectors) | **Get** /v1/ingestBudgets/{id}/collectors | Get a list of Collectors.
[**GetIngestBudget**](IngestBudgetManagementV1Api.md#GetIngestBudget) | **Get** /v1/ingestBudgets/{id} | Get an ingest budget.
[**ListIngestBudgets**](IngestBudgetManagementV1Api.md#ListIngestBudgets) | **Get** /v1/ingestBudgets | Get a list of ingest budgets.
[**RemoveCollectorFromBudget**](IngestBudgetManagementV1Api.md#RemoveCollectorFromBudget) | **Delete** /v1/ingestBudgets/{id}/collectors/{collectorId} | Remove Collector from a budget.
[**ResetUsage**](IngestBudgetManagementV1Api.md#ResetUsage) | **Post** /v1/ingestBudgets/{id}/usage/reset | Reset usage.
[**UpdateIngestBudget**](IngestBudgetManagementV1Api.md#UpdateIngestBudget) | **Put** /v1/ingestBudgets/{id} | Update an ingest budget.



## AssignCollectorToBudget

> IngestBudget AssignCollectorToBudget(ctx, id, collectorId).Execute()

Assign a Collector to a budget.



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
    id := "id_example" // string | Identifier of the ingest budget to assign to the Collector.
    collectorId := "collectorId_example" // string | Identifier of the Collector to assign.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV1Api.AssignCollectorToBudget(context.Background(), id, collectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.AssignCollectorToBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignCollectorToBudget`: IngestBudget
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV1Api.AssignCollectorToBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingest budget to assign to the Collector. | 
**collectorId** | **string** | Identifier of the Collector to assign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignCollectorToBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIngestBudget

> IngestBudget CreateIngestBudget(ctx).IngestBudgetDefinition(ingestBudgetDefinition).Execute()

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
    ingestBudgetDefinition := *openapiclient.NewIngestBudgetDefinition("Developer Budget", "dev_30_gb", int64(1000), "America/Los_Angeles", "23:30", "stopCollecting") // IngestBudgetDefinition | Information about the new ingest budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV1Api.CreateIngestBudget(context.Background()).IngestBudgetDefinition(ingestBudgetDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.CreateIngestBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIngestBudget`: IngestBudget
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV1Api.CreateIngestBudget`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIngestBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ingestBudgetDefinition** | [**IngestBudgetDefinition**](IngestBudgetDefinition.md) | Information about the new ingest budget. | 

### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIngestBudget

> DeleteIngestBudget(ctx, id).Execute()

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
    resp, r, err := api_client.IngestBudgetManagementV1Api.DeleteIngestBudget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.DeleteIngestBudget``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIngestBudgetRequest struct via the builder pattern


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


## GetAssignedCollectors

> ListCollectorIdentitiesResponse GetAssignedCollectors(ctx, id).Limit(limit).Token(token).Execute()

Get a list of Collectors.



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
    id := "id_example" // string | Identifier of ingest budget to which Collectors are assigned.
    limit := int32(56) // int32 | Limit the number of Collectors returned in the response. The number of Collectors returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV1Api.GetAssignedCollectors(context.Background(), id).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.GetAssignedCollectors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssignedCollectors`: ListCollectorIdentitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV1Api.GetAssignedCollectors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of ingest budget to which Collectors are assigned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignedCollectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of Collectors returned in the response. The number of Collectors returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListCollectorIdentitiesResponse**](ListCollectorIdentitiesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIngestBudget

> IngestBudget GetIngestBudget(ctx, id).Execute()

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
    resp, r, err := api_client.IngestBudgetManagementV1Api.GetIngestBudget(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.GetIngestBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIngestBudget`: IngestBudget
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV1Api.GetIngestBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of ingest budget to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIngestBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListIngestBudgets

> ListIngestBudgetsResponse ListIngestBudgets(ctx).Limit(limit).Token(token).Execute()

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
    resp, r, err := api_client.IngestBudgetManagementV1Api.ListIngestBudgets(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.ListIngestBudgets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListIngestBudgets`: ListIngestBudgetsResponse
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV1Api.ListIngestBudgets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListIngestBudgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of budgets returned in the response. The number of budgets returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. | 

### Return type

[**ListIngestBudgetsResponse**](ListIngestBudgetsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCollectorFromBudget

> IngestBudget RemoveCollectorFromBudget(ctx, id, collectorId).Execute()

Remove Collector from a budget.



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
    id := "id_example" // string | Identifier of the ingest budget to unassign from the Collector.
    collectorId := "collectorId_example" // string | Identifier of the Collector to unassign.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV1Api.RemoveCollectorFromBudget(context.Background(), id, collectorId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.RemoveCollectorFromBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveCollectorFromBudget`: IngestBudget
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV1Api.RemoveCollectorFromBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingest budget to unassign from the Collector. | 
**collectorId** | **string** | Identifier of the Collector to unassign. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCollectorFromBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetUsage

> ResetUsage(ctx, id).Execute()

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
    resp, r, err := api_client.IngestBudgetManagementV1Api.ResetUsage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.ResetUsage``: %v\n", err)
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

Other parameters are passed through a pointer to a apiResetUsageRequest struct via the builder pattern


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


## UpdateIngestBudget

> IngestBudget UpdateIngestBudget(ctx, id).IngestBudgetDefinition(ingestBudgetDefinition).Execute()

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
    ingestBudgetDefinition := *openapiclient.NewIngestBudgetDefinition("Developer Budget", "dev_30_gb", int64(1000), "America/Los_Angeles", "23:30", "stopCollecting") // IngestBudgetDefinition | Information to update about the ingest budget.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.IngestBudgetManagementV1Api.UpdateIngestBudget(context.Background(), id).IngestBudgetDefinition(ingestBudgetDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `IngestBudgetManagementV1Api.UpdateIngestBudget``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIngestBudget`: IngestBudget
    fmt.Fprintf(os.Stdout, "Response from `IngestBudgetManagementV1Api.UpdateIngestBudget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the ingest budget to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIngestBudgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ingestBudgetDefinition** | [**IngestBudgetDefinition**](IngestBudgetDefinition.md) | Information to update about the ingest budget. | 

### Return type

[**IngestBudget**](IngestBudget.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


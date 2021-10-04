# \PartitionManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelRetentionUpdate**](PartitionManagementApi.md#CancelRetentionUpdate) | **Post** /v1/partitions/{id}/cancelRetentionUpdate | Cancel a retention update for a partition
[**CreatePartition**](PartitionManagementApi.md#CreatePartition) | **Post** /v1/partitions | Create a new partition.
[**DecommissionPartition**](PartitionManagementApi.md#DecommissionPartition) | **Post** /v1/partitions/{id}/decommission | Decommission a partition.
[**GetPartition**](PartitionManagementApi.md#GetPartition) | **Get** /v1/partitions/{id} | Get a partition.
[**ListPartitions**](PartitionManagementApi.md#ListPartitions) | **Get** /v1/partitions | Get a list of partitions.
[**UpdatePartition**](PartitionManagementApi.md#UpdatePartition) | **Put** /v1/partitions/{id} | Update a partition.



## CancelRetentionUpdate

> CancelRetentionUpdate(ctx, id).Execute()

Cancel a retention update for a partition



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
    id := "1" // string | Identifier of the partition to cancel the retention update for.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionManagementApi.CancelRetentionUpdate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionManagementApi.CancelRetentionUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the partition to cancel the retention update for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelRetentionUpdateRequest struct via the builder pattern


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


## CreatePartition

> Partition CreatePartition(ctx).CreatePartitionDefinition(createPartitionDefinition).Execute()

Create a new partition.



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
    createPartitionDefinition := *openapiclient.NewCreatePartitionDefinition("apache", "_sourcecategory=*/Apache") // CreatePartitionDefinition | Information about the new partition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionManagementApi.CreatePartition(context.Background()).CreatePartitionDefinition(createPartitionDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionManagementApi.CreatePartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartition`: Partition
    fmt.Fprintf(os.Stdout, "Response from `PartitionManagementApi.CreatePartition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPartitionDefinition** | [**CreatePartitionDefinition**](CreatePartitionDefinition.md) | Information about the new partition. | 

### Return type

[**Partition**](Partition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecommissionPartition

> DecommissionPartition(ctx, id).Execute()

Decommission a partition.



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
    id := "id_example" // string | Identifier of the partition to decommission.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionManagementApi.DecommissionPartition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionManagementApi.DecommissionPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the partition to decommission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDecommissionPartitionRequest struct via the builder pattern


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


## GetPartition

> Partition GetPartition(ctx, id).Execute()

Get a partition.



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
    id := "id_example" // string | Identifier of partition to return.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionManagementApi.GetPartition(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionManagementApi.GetPartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPartition`: Partition
    fmt.Fprintf(os.Stdout, "Response from `PartitionManagementApi.GetPartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of partition to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Partition**](Partition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPartitions

> ListPartitionsResponse ListPartitions(ctx).Limit(limit).Token(token).ViewTypes(viewTypes).Execute()

Get a list of partitions.



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
    limit := int32(56) // int32 | Limit the number of partitions returned in the response. The number of partitions returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)
    viewTypes := []string{"Inner_example"} // []string | The type of partitions to retrieve. Valid values are:   1. `DefaultView`: To get General Index partition.   2. `Partition`: To get user defined views/partitions.   3. `AuditIndex`: To get the internal audit indexes. Eg. sumologic_audit_events.  More than one type of partitions can be retrieved in same request. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionManagementApi.ListPartitions(context.Background()).Limit(limit).Token(token).ViewTypes(viewTypes).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionManagementApi.ListPartitions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPartitions`: ListPartitionsResponse
    fmt.Fprintf(os.Stdout, "Response from `PartitionManagementApi.ListPartitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPartitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of partitions returned in the response. The number of partitions returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 
 **viewTypes** | **[]string** | The type of partitions to retrieve. Valid values are:   1. &#x60;DefaultView&#x60;: To get General Index partition.   2. &#x60;Partition&#x60;: To get user defined views/partitions.   3. &#x60;AuditIndex&#x60;: To get the internal audit indexes. Eg. sumologic_audit_events.  More than one type of partitions can be retrieved in same request. | 

### Return type

[**ListPartitionsResponse**](ListPartitionsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePartition

> Partition UpdatePartition(ctx, id).UpdatePartitionDefinition(updatePartitionDefinition).Execute()

Update a partition.



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
    id := "id_example" // string | Identifier of the partition to update.
    updatePartitionDefinition := *openapiclient.NewUpdatePartitionDefinition() // UpdatePartitionDefinition | Information to update about the partition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartitionManagementApi.UpdatePartition(context.Background(), id).UpdatePartitionDefinition(updatePartitionDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartitionManagementApi.UpdatePartition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePartition`: Partition
    fmt.Fprintf(os.Stdout, "Response from `PartitionManagementApi.UpdatePartition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the partition to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePartitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePartitionDefinition** | [**UpdatePartitionDefinition**](UpdatePartitionDefinition.md) | Information to update about the partition. | 

### Return type

[**Partition**](Partition.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


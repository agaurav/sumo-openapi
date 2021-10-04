# \ArchiveManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateArchiveJob**](ArchiveManagementApi.md#CreateArchiveJob) | **Post** /v1/archive/{sourceId}/jobs | Create an ingestion job.
[**DeleteArchiveJob**](ArchiveManagementApi.md#DeleteArchiveJob) | **Delete** /v1/archive/{sourceId}/jobs/{id} | Delete an ingestion job.
[**ListArchiveJobsBySourceId**](ArchiveManagementApi.md#ListArchiveJobsBySourceId) | **Get** /v1/archive/{sourceId}/jobs | Get ingestion jobs for an Archive Source.
[**ListArchiveJobsCountPerSource**](ArchiveManagementApi.md#ListArchiveJobsCountPerSource) | **Get** /v1/archive/jobs/count | List ingestion jobs for all Archive Sources.



## CreateArchiveJob

> ArchiveJob CreateArchiveJob(ctx, sourceId).CreateArchiveJobRequest(createArchiveJobRequest).Execute()

Create an ingestion job.



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
    sourceId := "000000000606C009" // string | The identifier of the Archive Source for which the job is to be added.
    createArchiveJobRequest := *openapiclient.NewCreateArchiveJobRequest("Name_example", time.Now(), time.Now()) // CreateArchiveJobRequest | The definition of the ingestion job to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchiveManagementApi.CreateArchiveJob(context.Background(), sourceId).CreateArchiveJobRequest(createArchiveJobRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveManagementApi.CreateArchiveJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateArchiveJob`: ArchiveJob
    fmt.Fprintf(os.Stdout, "Response from `ArchiveManagementApi.CreateArchiveJob`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The identifier of the Archive Source for which the job is to be added. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateArchiveJobRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createArchiveJobRequest** | [**CreateArchiveJobRequest**](CreateArchiveJobRequest.md) | The definition of the ingestion job to create. | 

### Return type

[**ArchiveJob**](ArchiveJob.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteArchiveJob

> DeleteArchiveJob(ctx, sourceId, id).Execute()

Delete an ingestion job.



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
    sourceId := "sourceId_example" // string | The identifier of the Archive Source.
    id := "id_example" // string | The identifier of the ingestion job to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchiveManagementApi.DeleteArchiveJob(context.Background(), sourceId, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveManagementApi.DeleteArchiveJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The identifier of the Archive Source. | 
**id** | **string** | The identifier of the ingestion job to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteArchiveJobRequest struct via the builder pattern


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


## ListArchiveJobsBySourceId

> ListArchiveJobsResponse ListArchiveJobsBySourceId(ctx, sourceId).Limit(limit).Token(token).Execute()

Get ingestion jobs for an Archive Source.



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
    sourceId := "000000000606C009" // string | The identifier of an Archive Source.
    limit := int32(56) // int32 | Limit the number of jobs returned in the response. The number of jobs returned may be less than the `limit`. (optional) (default to 10)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ArchiveManagementApi.ListArchiveJobsBySourceId(context.Background(), sourceId).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveManagementApi.ListArchiveJobsBySourceId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchiveJobsBySourceId`: ListArchiveJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `ArchiveManagementApi.ListArchiveJobsBySourceId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sourceId** | **string** | The identifier of an Archive Source. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListArchiveJobsBySourceIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of jobs returned in the response. The number of jobs returned may be less than the &#x60;limit&#x60;. | [default to 10]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListArchiveJobsResponse**](ListArchiveJobsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListArchiveJobsCountPerSource

> ListArchiveJobsCount ListArchiveJobsCountPerSource(ctx).Execute()

List ingestion jobs for all Archive Sources.



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
    resp, r, err := api_client.ArchiveManagementApi.ListArchiveJobsCountPerSource(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ArchiveManagementApi.ListArchiveJobsCountPerSource``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListArchiveJobsCountPerSource`: ListArchiveJobsCount
    fmt.Fprintf(os.Stdout, "Response from `ArchiveManagementApi.ListArchiveJobsCountPerSource`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListArchiveJobsCountPerSourceRequest struct via the builder pattern


### Return type

[**ListArchiveJobsCount**](ListArchiveJobsCount.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


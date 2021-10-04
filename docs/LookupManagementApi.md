# \LookupManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTable**](LookupManagementApi.md#CreateTable) | **Post** /v1/lookupTables | Create a lookup table.
[**DeleteTable**](LookupManagementApi.md#DeleteTable) | **Delete** /v1/lookupTables/{id} | Delete a lookup table.
[**DeleteTableRow**](LookupManagementApi.md#DeleteTableRow) | **Put** /v1/lookupTables/{id}/deleteTableRow | Delete a lookup table row.
[**LookupTableById**](LookupManagementApi.md#LookupTableById) | **Get** /v1/lookupTables/{id} | Get a lookup table.
[**RequestJobStatus**](LookupManagementApi.md#RequestJobStatus) | **Get** /v1/lookupTables/jobs/{jobId}/status | Get the status of an async job.
[**TruncateTable**](LookupManagementApi.md#TruncateTable) | **Post** /v1/lookupTables/{id}/truncate | Empty a lookup table.
[**UpdateTable**](LookupManagementApi.md#UpdateTable) | **Put** /v1/lookupTables/{id} | Edit a lookup table.
[**UpdateTableRow**](LookupManagementApi.md#UpdateTableRow) | **Put** /v1/lookupTables/{id}/row | Insert or Update a lookup table row.
[**UploadFile**](LookupManagementApi.md#UploadFile) | **Post** /v1/lookupTables/{id}/upload | Upload a CSV file.



## CreateTable

> LookupTable CreateTable(ctx).LookupTableDefinition(lookupTableDefinition).Execute()

Create a lookup table.



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
    lookupTableDefinition := *openapiclient.NewLookupTableDefinition("This is a sample lookup table description.", []openapiclient.LookupTableField{*openapiclient.NewLookupTableField("FieldName1", "boolean")}, []string{"PrimaryKeys_example"}, "SampleLookupTable", "0000000001C41EE4") // LookupTableDefinition | The schema and configuration for the lookup table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.CreateTable(context.Background()).LookupTableDefinition(lookupTableDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.CreateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTable`: LookupTable
    fmt.Fprintf(os.Stdout, "Response from `LookupManagementApi.CreateTable`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lookupTableDefinition** | [**LookupTableDefinition**](LookupTableDefinition.md) | The schema and configuration for the lookup table. | 

### Return type

[**LookupTable**](LookupTable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteTable

> DeleteTable(ctx, id).Execute()

Delete a lookup table.



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
    id := "0000000001C41EE4" // string | Identifier of the lookup table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.DeleteTable(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.DeleteTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the lookup table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableRequest struct via the builder pattern


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


## DeleteTableRow

> DeleteTableRow(ctx, id).RowDeleteDefinition(rowDeleteDefinition).Execute()

Delete a lookup table row.



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
    id := "0000000001C41EE4" // string | Identifier of the lookup table.
    rowDeleteDefinition := *openapiclient.NewRowDeleteDefinition([]openapiclient.TableRow{*openapiclient.NewTableRow("user_id", "user1")}) // RowDeleteDefinition | Lookup table row delete definition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.DeleteTableRow(context.Background(), id).RowDeleteDefinition(rowDeleteDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.DeleteTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the lookup table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rowDeleteDefinition** | [**RowDeleteDefinition**](RowDeleteDefinition.md) | Lookup table row delete definition. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LookupTableById

> LookupTable LookupTableById(ctx, id).Execute()

Get a lookup table.



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
    id := "0000000001C41EE4" // string | Identifier of the lookup table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.LookupTableById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.LookupTableById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LookupTableById`: LookupTable
    fmt.Fprintf(os.Stdout, "Response from `LookupManagementApi.LookupTableById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the lookup table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiLookupTableByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LookupTable**](LookupTable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RequestJobStatus

> LookupAsyncJobStatus RequestJobStatus(ctx, jobId).Execute()

Get the status of an async job.



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
    jobId := "0000000001C41AA3" // string | An identifier returned in response to an asynchronous request.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.RequestJobStatus(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.RequestJobStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RequestJobStatus`: LookupAsyncJobStatus
    fmt.Fprintf(os.Stdout, "Response from `LookupManagementApi.RequestJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | An identifier returned in response to an asynchronous request. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRequestJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LookupAsyncJobStatus**](LookupAsyncJobStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TruncateTable

> LookupRequestToken TruncateTable(ctx, id).Execute()

Empty a lookup table.



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
    id := "0000000001C41EE4" // string | Identifier of the table to clear.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.TruncateTable(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.TruncateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TruncateTable`: LookupRequestToken
    fmt.Fprintf(os.Stdout, "Response from `LookupManagementApi.TruncateTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the table to clear. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTruncateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LookupRequestToken**](LookupRequestToken.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTable

> LookupTable UpdateTable(ctx, id).LookupUpdateDefinition(lookupUpdateDefinition).Execute()

Edit a lookup table.



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
    id := "0000000001C41EE4" // string | Identifier of the lookup table.
    lookupUpdateDefinition := *openapiclient.NewLookupUpdateDefinition(int32(100), "This is a sample lookup table description.") // LookupUpdateDefinition | The configuration changes for the lookup table.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.UpdateTable(context.Background(), id).LookupUpdateDefinition(lookupUpdateDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.UpdateTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTable`: LookupTable
    fmt.Fprintf(os.Stdout, "Response from `LookupManagementApi.UpdateTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the lookup table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lookupUpdateDefinition** | [**LookupUpdateDefinition**](LookupUpdateDefinition.md) | The configuration changes for the lookup table. | 

### Return type

[**LookupTable**](LookupTable.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTableRow

> UpdateTableRow(ctx, id).RowUpdateDefinition(rowUpdateDefinition).Execute()

Insert or Update a lookup table row.



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
    id := "0000000001C41EE4" // string | Identifier of the lookup table.
    rowUpdateDefinition := *openapiclient.NewRowUpdateDefinition([]openapiclient.TableRow{*openapiclient.NewTableRow("user_id", "user1")}) // RowUpdateDefinition | Lookup table row update definition.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.UpdateTableRow(context.Background(), id).RowUpdateDefinition(rowUpdateDefinition).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.UpdateTableRow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the lookup table. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTableRowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rowUpdateDefinition** | [**RowUpdateDefinition**](RowUpdateDefinition.md) | Lookup table row update definition. | 

### Return type

 (empty response body)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadFile

> LookupRequestToken UploadFile(ctx, id).File(file).Merge(merge).FileEncoding(fileEncoding).Execute()

Upload a CSV file.



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
    id := "0000000001C41EE4" // string | Identifier of the lookup table to populate.
    file := os.NewFile(1234, "some_file") // *os.File | The CSV file to upload.   - The size limit for the CSV file is 100MB.   - Use Unix format, with newlines (\\\"\\\\n\\\") separating rows.   - The first row should contain headers that match the lookup table schema. Matching is     case-insensitive.
    merge := true // bool | This indicates whether the file contents will be merged with existing data in the lookup table or not. If this is true then data with the same primary keys will be updated while the rest of the rows will be appended. By default, merge is false. The response includes a request identifier that you need to use in the [Request Status API](#operation/requestStatus) to track the status of the upload request. (optional) (default to false)
    fileEncoding := "UTF-16" // string | File encoding of file being uploaded. (optional) (default to "UTF-8")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.LookupManagementApi.UploadFile(context.Background(), id).File(file).Merge(merge).FileEncoding(fileEncoding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookupManagementApi.UploadFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UploadFile`: LookupRequestToken
    fmt.Fprintf(os.Stdout, "Response from `LookupManagementApi.UploadFile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the lookup table to populate. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUploadFileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The CSV file to upload.   - The size limit for the CSV file is 100MB.   - Use Unix format, with newlines (\\\&quot;\\\\n\\\&quot;) separating rows.   - The first row should contain headers that match the lookup table schema. Matching is     case-insensitive. | 
 **merge** | **bool** | This indicates whether the file contents will be merged with existing data in the lookup table or not. If this is true then data with the same primary keys will be updated while the rest of the rows will be appended. By default, merge is false. The response includes a request identifier that you need to use in the [Request Status API](#operation/requestStatus) to track the status of the upload request. | [default to false]
 **fileEncoding** | **string** | File encoding of file being uploaded. | [default to &quot;UTF-8&quot;]

### Return type

[**LookupRequestToken**](LookupRequestToken.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


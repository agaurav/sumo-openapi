# \FieldManagementV1Api

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateField**](FieldManagementV1Api.md#CreateField) | **Post** /v1/fields | Create a new field.
[**DeleteField**](FieldManagementV1Api.md#DeleteField) | **Delete** /v1/fields/{id} | Delete a custom field.
[**DisableField**](FieldManagementV1Api.md#DisableField) | **Delete** /v1/fields/{id}/disable | Disable a custom field.
[**EnableField**](FieldManagementV1Api.md#EnableField) | **Put** /v1/fields/{id}/enable | Enable custom field with a specified identifier.
[**GetBuiltInField**](FieldManagementV1Api.md#GetBuiltInField) | **Get** /v1/fields/builtin/{id} | Get a built-in field.
[**GetCustomField**](FieldManagementV1Api.md#GetCustomField) | **Get** /v1/fields/{id} | Get a custom field.
[**GetFieldQuota**](FieldManagementV1Api.md#GetFieldQuota) | **Get** /v1/fields/quota | Get capacity information.
[**ListBuiltInFields**](FieldManagementV1Api.md#ListBuiltInFields) | **Get** /v1/fields/builtin | Get a list of built-in fields.
[**ListCustomFields**](FieldManagementV1Api.md#ListCustomFields) | **Get** /v1/fields | Get a list of all custom fields.
[**ListDroppedFields**](FieldManagementV1Api.md#ListDroppedFields) | **Get** /v1/fields/dropped | Get a list of dropped fields.



## CreateField

> CustomField CreateField(ctx).FieldName(fieldName).Execute()

Create a new field.



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
    fieldName := *openapiclient.NewFieldName("hostIP") // FieldName | Name of a field to add. The name is used as the key in the key-value pair.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FieldManagementV1Api.CreateField(context.Background()).FieldName(fieldName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.CreateField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateField`: CustomField
    fmt.Fprintf(os.Stdout, "Response from `FieldManagementV1Api.CreateField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldName** | [**FieldName**](FieldName.md) | Name of a field to add. The name is used as the key in the key-value pair. | 

### Return type

[**CustomField**](CustomField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteField

> DeleteField(ctx, id).Execute()

Delete a custom field.



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
    id := "00000000031D02DA" // string | Identifier of a field to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FieldManagementV1Api.DeleteField(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.DeleteField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of a field to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldRequest struct via the builder pattern


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


## DisableField

> DisableField(ctx, id).Execute()

Disable a custom field.



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
    id := "00000000031D02DA" // string | Identifier of a field to disable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FieldManagementV1Api.DisableField(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.DisableField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of a field to disable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDisableFieldRequest struct via the builder pattern


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


## EnableField

> EnableField(ctx, id).Execute()

Enable custom field with a specified identifier.



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
    id := "00000000031D02DA" // string | Identifier of a field to enable.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FieldManagementV1Api.EnableField(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.EnableField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of a field to enable. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEnableFieldRequest struct via the builder pattern


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


## GetBuiltInField

> BuiltinField GetBuiltInField(ctx, id).Execute()

Get a built-in field.



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
    id := "000000000000000A" // string | Identifier of a built-in field.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FieldManagementV1Api.GetBuiltInField(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.GetBuiltInField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuiltInField`: BuiltinField
    fmt.Fprintf(os.Stdout, "Response from `FieldManagementV1Api.GetBuiltInField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of a built-in field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuiltInFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BuiltinField**](BuiltinField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomField

> CustomField GetCustomField(ctx, id).Execute()

Get a custom field.



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
    id := "00000000031D02DA" // string | Identifier of a field.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.FieldManagementV1Api.GetCustomField(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.GetCustomField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomField`: CustomField
    fmt.Fprintf(os.Stdout, "Response from `FieldManagementV1Api.GetCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of a field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomField**](CustomField.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldQuota

> FieldQuotaUsage GetFieldQuota(ctx).Execute()

Get capacity information.



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
    resp, r, err := api_client.FieldManagementV1Api.GetFieldQuota(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.GetFieldQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFieldQuota`: FieldQuotaUsage
    fmt.Fprintf(os.Stdout, "Response from `FieldManagementV1Api.GetFieldQuota`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldQuotaRequest struct via the builder pattern


### Return type

[**FieldQuotaUsage**](FieldQuotaUsage.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuiltInFields

> ListBuiltinFieldsResponse ListBuiltInFields(ctx).Execute()

Get a list of built-in fields.



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
    resp, r, err := api_client.FieldManagementV1Api.ListBuiltInFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.ListBuiltInFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuiltInFields`: ListBuiltinFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `FieldManagementV1Api.ListBuiltInFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuiltInFieldsRequest struct via the builder pattern


### Return type

[**ListBuiltinFieldsResponse**](ListBuiltinFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCustomFields

> ListCustomFieldsResponse ListCustomFields(ctx).Execute()

Get a list of all custom fields.



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
    resp, r, err := api_client.FieldManagementV1Api.ListCustomFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.ListCustomFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomFields`: ListCustomFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `FieldManagementV1Api.ListCustomFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomFieldsRequest struct via the builder pattern


### Return type

[**ListCustomFieldsResponse**](ListCustomFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDroppedFields

> ListDroppedFieldsResponse ListDroppedFields(ctx).Execute()

Get a list of dropped fields.



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
    resp, r, err := api_client.FieldManagementV1Api.ListDroppedFields(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldManagementV1Api.ListDroppedFields``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDroppedFields`: ListDroppedFieldsResponse
    fmt.Fprintf(os.Stdout, "Response from `FieldManagementV1Api.ListDroppedFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDroppedFieldsRequest struct via the builder pattern


### Return type

[**ListDroppedFieldsResponse**](ListDroppedFieldsResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


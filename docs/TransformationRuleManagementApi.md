# \TransformationRuleManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRule**](TransformationRuleManagementApi.md#CreateRule) | **Post** /v1/transformationRules | Create a new transformation rule.
[**DeleteRule**](TransformationRuleManagementApi.md#DeleteRule) | **Delete** /v1/transformationRules/{id} | Delete a transformation rule.
[**GetTransformationRule**](TransformationRuleManagementApi.md#GetTransformationRule) | **Get** /v1/transformationRules/{id} | Get a transformation rule.
[**GetTransformationRules**](TransformationRuleManagementApi.md#GetTransformationRules) | **Get** /v1/transformationRules | Get a list of transformation rules.
[**UpdateTransformationRule**](TransformationRuleManagementApi.md#UpdateTransformationRule) | **Put** /v1/transformationRules/{id} | Update a transformation rule.



## CreateRule

> TransformationRuleResponse CreateRule(ctx).TransformationRuleRequest(transformationRuleRequest).Execute()

Create a new transformation rule.



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
    transformationRuleRequest := *openapiclient.NewTransformationRuleRequest(*openapiclient.NewTransformationRuleDefinition("Transformation Rule 1", "_sourceCategory=metricsstore", int64(8)), true) // TransformationRuleRequest | The configuration of the transformation rule to create.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransformationRuleManagementApi.CreateRule(context.Background()).TransformationRuleRequest(transformationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationRuleManagementApi.CreateRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRule`: TransformationRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `TransformationRuleManagementApi.CreateRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transformationRuleRequest** | [**TransformationRuleRequest**](TransformationRuleRequest.md) | The configuration of the transformation rule to create. | 

### Return type

[**TransformationRuleResponse**](TransformationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRule

> DeleteRule(ctx, id).Execute()

Delete a transformation rule.



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
    id := "id_example" // string | Identifier of the transformation rule to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransformationRuleManagementApi.DeleteRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationRuleManagementApi.DeleteRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the transformation rule to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRuleRequest struct via the builder pattern


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


## GetTransformationRule

> TransformationRuleResponse GetTransformationRule(ctx, id).Execute()

Get a transformation rule.



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
    id := "id_example" // string | Identifier of transformation rule to return.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransformationRuleManagementApi.GetTransformationRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationRuleManagementApi.GetTransformationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransformationRule`: TransformationRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `TransformationRuleManagementApi.GetTransformationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of transformation rule to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransformationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransformationRuleResponse**](TransformationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransformationRules

> TransformationRulesResponse GetTransformationRules(ctx).Limit(limit).Token(token).Execute()

Get a list of transformation rules.



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
    limit := int32(10) // int32 | Limit the number of transformation rules returned in the response. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransformationRuleManagementApi.GetTransformationRules(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationRuleManagementApi.GetTransformationRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTransformationRules`: TransformationRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `TransformationRuleManagementApi.GetTransformationRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTransformationRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of transformation rules returned in the response. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**TransformationRulesResponse**](TransformationRulesResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTransformationRule

> TransformationRuleResponse UpdateTransformationRule(ctx, id).TransformationRuleRequest(transformationRuleRequest).Execute()

Update a transformation rule.



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
    id := "id_example" // string | Identifier of the transformation rule to update.
    transformationRuleRequest := *openapiclient.NewTransformationRuleRequest(*openapiclient.NewTransformationRuleDefinition("Transformation Rule 1", "_sourceCategory=metricsstore", int64(8)), true) // TransformationRuleRequest | Information to update about the transformation rule.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TransformationRuleManagementApi.UpdateTransformationRule(context.Background(), id).TransformationRuleRequest(transformationRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TransformationRuleManagementApi.UpdateTransformationRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTransformationRule`: TransformationRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `TransformationRuleManagementApi.UpdateTransformationRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the transformation rule to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTransformationRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transformationRuleRequest** | [**TransformationRuleRequest**](TransformationRuleRequest.md) | Information to update about the transformation rule. | 

### Return type

[**TransformationRuleResponse**](TransformationRuleResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


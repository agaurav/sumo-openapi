# \HealthEventsApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAllHealthEvents**](HealthEventsApi.md#ListAllHealthEvents) | **Get** /v1/healthEvents | Get a list of health events.
[**ListAllHealthEventsForResources**](HealthEventsApi.md#ListAllHealthEventsForResources) | **Post** /v1/healthEvents/resources | Health events for specific resources.



## ListAllHealthEvents

> ListHealthEventResponse ListAllHealthEvents(ctx).Limit(limit).Token(token).Execute()

Get a list of health events.



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
    limit := int32(56) // int32 | Limit the number of health events returned in the response. The number of health events returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthEventsApi.ListAllHealthEvents(context.Background()).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthEventsApi.ListAllHealthEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllHealthEvents`: ListHealthEventResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthEventsApi.ListAllHealthEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllHealthEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Limit the number of health events returned in the response. The number of health events returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListHealthEventResponse**](ListHealthEventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllHealthEventsForResources

> ListHealthEventResponse ListAllHealthEventsForResources(ctx).ResourceIdentities(resourceIdentities).Limit(limit).Token(token).Execute()

Health events for specific resources.



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
    resourceIdentities := *openapiclient.NewResourceIdentities([]openapiclient.ResourceIdentity{*openapiclient.NewResourceIdentity("C03E086C137F38B4", "Collector")}) // ResourceIdentities | Resource identifiers to request health events from.
    limit := int32(56) // int32 | Limit the number of health events returned in the response. The number of health events returned may be less than the `limit`. (optional) (default to 100)
    token := "token_example" // string | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. `token` is set to null when no more pages are left. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.HealthEventsApi.ListAllHealthEventsForResources(context.Background()).ResourceIdentities(resourceIdentities).Limit(limit).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthEventsApi.ListAllHealthEventsForResources``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllHealthEventsForResources`: ListHealthEventResponse
    fmt.Fprintf(os.Stdout, "Response from `HealthEventsApi.ListAllHealthEventsForResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllHealthEventsForResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceIdentities** | [**ResourceIdentities**](ResourceIdentities.md) | Resource identifiers to request health events from. | 
 **limit** | **int32** | Limit the number of health events returned in the response. The number of health events returned may be less than the &#x60;limit&#x60;. | [default to 100]
 **token** | **string** | Continuation token to get the next page of results. A page object with the next continuation token is returned in the response body. Subsequent GET requests should specify the continuation token to get the next page of results. &#x60;token&#x60; is set to null when no more pages are left. | 

### Return type

[**ListHealthEventResponse**](ListHealthEventResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


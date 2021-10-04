# \ServiceAllowlistManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAllowlistedCidrs**](ServiceAllowlistManagementApi.md#AddAllowlistedCidrs) | **Post** /v1/serviceAllowlist/addresses/add | Allowlist CIDRs/IP addresses.
[**DeleteAllowlistedCidrs**](ServiceAllowlistManagementApi.md#DeleteAllowlistedCidrs) | **Post** /v1/serviceAllowlist/addresses/remove | Remove allowlisted CIDRs/IP addresses.
[**DisableAllowlisting**](ServiceAllowlistManagementApi.md#DisableAllowlisting) | **Post** /v1/serviceAllowlist/disable | Disable service allowlisting.
[**EnableAllowlisting**](ServiceAllowlistManagementApi.md#EnableAllowlisting) | **Post** /v1/serviceAllowlist/enable | Enable service allowlisting.
[**GetAllowlistingStatus**](ServiceAllowlistManagementApi.md#GetAllowlistingStatus) | **Get** /v1/serviceAllowlist/status | Get the allowlisting status.
[**ListAllowlistedCidrs**](ServiceAllowlistManagementApi.md#ListAllowlistedCidrs) | **Get** /v1/serviceAllowlist/addresses | List all allowlisted CIDRs/IP addresses.



## AddAllowlistedCidrs

> CidrList AddAllowlistedCidrs(ctx).CidrList(cidrList).Execute()

Allowlist CIDRs/IP addresses.



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
    cidrList := *openapiclient.NewCidrList([]openapiclient.Cidr{*openapiclient.NewCidr("192.35.24.1")}) // CidrList | List of all CIDR notations and/or IP addresses to be added to the allowlist of the organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAllowlistManagementApi.AddAllowlistedCidrs(context.Background()).CidrList(cidrList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAllowlistManagementApi.AddAllowlistedCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddAllowlistedCidrs`: CidrList
    fmt.Fprintf(os.Stdout, "Response from `ServiceAllowlistManagementApi.AddAllowlistedCidrs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddAllowlistedCidrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cidrList** | [**CidrList**](CidrList.md) | List of all CIDR notations and/or IP addresses to be added to the allowlist of the organization. | 

### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllowlistedCidrs

> CidrList DeleteAllowlistedCidrs(ctx).CidrList(cidrList).Execute()

Remove allowlisted CIDRs/IP addresses.



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
    cidrList := *openapiclient.NewCidrList([]openapiclient.Cidr{*openapiclient.NewCidr("192.35.24.1")}) // CidrList | List of all CIDR notations and/or IP addresses to be removed from the allowlist of the organization.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAllowlistManagementApi.DeleteAllowlistedCidrs(context.Background()).CidrList(cidrList).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAllowlistManagementApi.DeleteAllowlistedCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAllowlistedCidrs`: CidrList
    fmt.Fprintf(os.Stdout, "Response from `ServiceAllowlistManagementApi.DeleteAllowlistedCidrs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowlistedCidrsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cidrList** | [**CidrList**](CidrList.md) | List of all CIDR notations and/or IP addresses to be removed from the allowlist of the organization. | 

### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisableAllowlisting

> DisableAllowlisting(ctx).AllowlistType(allowlistType).Execute()

Disable service allowlisting.



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
    allowlistType := "Login" // string | The type of allowlisting to be disabled. It can be one of: `Login`, `Content`, or `Both`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAllowlistManagementApi.DisableAllowlisting(context.Background()).AllowlistType(allowlistType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAllowlistManagementApi.DisableAllowlisting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisableAllowlistingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowlistType** | **string** | The type of allowlisting to be disabled. It can be one of: &#x60;Login&#x60;, &#x60;Content&#x60;, or &#x60;Both&#x60;. | 

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


## EnableAllowlisting

> EnableAllowlisting(ctx).AllowlistType(allowlistType).Execute()

Enable service allowlisting.



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
    allowlistType := "Login" // string | The type of allowlisting to be enabled. It can be one of: `Login`, `Content`, or `Both`.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ServiceAllowlistManagementApi.EnableAllowlisting(context.Background()).AllowlistType(allowlistType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAllowlistManagementApi.EnableAllowlisting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnableAllowlistingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **allowlistType** | **string** | The type of allowlisting to be enabled. It can be one of: &#x60;Login&#x60;, &#x60;Content&#x60;, or &#x60;Both&#x60;. | 

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


## GetAllowlistingStatus

> AllowlistingStatus GetAllowlistingStatus(ctx).Execute()

Get the allowlisting status.



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
    resp, r, err := api_client.ServiceAllowlistManagementApi.GetAllowlistingStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAllowlistManagementApi.GetAllowlistingStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllowlistingStatus`: AllowlistingStatus
    fmt.Fprintf(os.Stdout, "Response from `ServiceAllowlistManagementApi.GetAllowlistingStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowlistingStatusRequest struct via the builder pattern


### Return type

[**AllowlistingStatus**](AllowlistingStatus.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllowlistedCidrs

> CidrList ListAllowlistedCidrs(ctx).Execute()

List all allowlisted CIDRs/IP addresses.



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
    resp, r, err := api_client.ServiceAllowlistManagementApi.ListAllowlistedCidrs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ServiceAllowlistManagementApi.ListAllowlistedCidrs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllowlistedCidrs`: CidrList
    fmt.Fprintf(os.Stdout, "Response from `ServiceAllowlistManagementApi.ListAllowlistedCidrs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllowlistedCidrsRequest struct via the builder pattern


### Return type

[**CidrList**](CidrList.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


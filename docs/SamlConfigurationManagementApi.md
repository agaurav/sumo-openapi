# \SamlConfigurationManagementApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAllowlistedUser**](SamlConfigurationManagementApi.md#CreateAllowlistedUser) | **Post** /v1/saml/allowlistedUsers/{userId} | Allowlist a user.
[**CreateIdentityProvider**](SamlConfigurationManagementApi.md#CreateIdentityProvider) | **Post** /v1/saml/identityProviders | Create a new SAML configuration.
[**DeleteAllowlistedUser**](SamlConfigurationManagementApi.md#DeleteAllowlistedUser) | **Delete** /v1/saml/allowlistedUsers/{userId} | Remove an allowlisted user.
[**DeleteIdentityProvider**](SamlConfigurationManagementApi.md#DeleteIdentityProvider) | **Delete** /v1/saml/identityProviders/{id} | Delete a SAML configuration.
[**DisableSamlLockdown**](SamlConfigurationManagementApi.md#DisableSamlLockdown) | **Post** /v1/saml/lockdown/disable | Disable SAML lockdown.
[**EnableSamlLockdown**](SamlConfigurationManagementApi.md#EnableSamlLockdown) | **Post** /v1/saml/lockdown/enable | Require SAML for sign-in.
[**GetAllowlistedUsers**](SamlConfigurationManagementApi.md#GetAllowlistedUsers) | **Get** /v1/saml/allowlistedUsers | Get list of allowlisted users.
[**GetIdentityProviders**](SamlConfigurationManagementApi.md#GetIdentityProviders) | **Get** /v1/saml/identityProviders | Get a list of SAML configurations.
[**UpdateIdentityProvider**](SamlConfigurationManagementApi.md#UpdateIdentityProvider) | **Put** /v1/saml/identityProviders/{id} | Update a SAML configuration.



## CreateAllowlistedUser

> AllowlistedUserResult CreateAllowlistedUser(ctx, userId).Execute()

Allowlist a user.



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
    userId := "userId_example" // string | Identifier of the user.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SamlConfigurationManagementApi.CreateAllowlistedUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.CreateAllowlistedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAllowlistedUser`: AllowlistedUserResult
    fmt.Fprintf(os.Stdout, "Response from `SamlConfigurationManagementApi.CreateAllowlistedUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of the user. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAllowlistedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AllowlistedUserResult**](AllowlistedUserResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIdentityProvider

> SamlIdentityProvider CreateIdentityProvider(ctx).SamlIdentityProviderRequest(samlIdentityProviderRequest).Execute()

Create a new SAML configuration.



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
    samlIdentityProviderRequest := *openapiclient.NewSamlIdentityProviderRequest("SumoLogic", "http://www.okta.com/abxcseyuiwelflkdjh", "X509cert1_example") // SamlIdentityProviderRequest | The configuration of the SAML identity provider.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SamlConfigurationManagementApi.CreateIdentityProvider(context.Background()).SamlIdentityProviderRequest(samlIdentityProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.CreateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateIdentityProvider`: SamlIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `SamlConfigurationManagementApi.CreateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **samlIdentityProviderRequest** | [**SamlIdentityProviderRequest**](SamlIdentityProviderRequest.md) | The configuration of the SAML identity provider. | 

### Return type

[**SamlIdentityProvider**](SamlIdentityProvider.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAllowlistedUser

> DeleteAllowlistedUser(ctx, userId).Execute()

Remove an allowlisted user.



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
    userId := "userId_example" // string | Identifier of user that will no longer be allowlisted from SAML Lockdown.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SamlConfigurationManagementApi.DeleteAllowlistedUser(context.Background(), userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.DeleteAllowlistedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | Identifier of user that will no longer be allowlisted from SAML Lockdown. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllowlistedUserRequest struct via the builder pattern


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


## DeleteIdentityProvider

> DeleteIdentityProvider(ctx, id).Execute()

Delete a SAML configuration.



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
    id := "id_example" // string | Identifier of the SAML configuration to delete.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SamlConfigurationManagementApi.DeleteIdentityProvider(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.DeleteIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the SAML configuration to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIdentityProviderRequest struct via the builder pattern


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


## DisableSamlLockdown

> DisableSamlLockdown(ctx).Execute()

Disable SAML lockdown.



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
    resp, r, err := api_client.SamlConfigurationManagementApi.DisableSamlLockdown(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.DisableSamlLockdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDisableSamlLockdownRequest struct via the builder pattern


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


## EnableSamlLockdown

> EnableSamlLockdown(ctx).Execute()

Require SAML for sign-in.



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
    resp, r, err := api_client.SamlConfigurationManagementApi.EnableSamlLockdown(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.EnableSamlLockdown``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnableSamlLockdownRequest struct via the builder pattern


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


## GetAllowlistedUsers

> []AllowlistedUserResult GetAllowlistedUsers(ctx).Execute()

Get list of allowlisted users.



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
    resp, r, err := api_client.SamlConfigurationManagementApi.GetAllowlistedUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.GetAllowlistedUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllowlistedUsers`: []AllowlistedUserResult
    fmt.Fprintf(os.Stdout, "Response from `SamlConfigurationManagementApi.GetAllowlistedUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllowlistedUsersRequest struct via the builder pattern


### Return type

[**[]AllowlistedUserResult**](AllowlistedUserResult.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdentityProviders

> []SamlIdentityProvider GetIdentityProviders(ctx).Execute()

Get a list of SAML configurations.



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
    resp, r, err := api_client.SamlConfigurationManagementApi.GetIdentityProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.GetIdentityProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetIdentityProviders`: []SamlIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `SamlConfigurationManagementApi.GetIdentityProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIdentityProvidersRequest struct via the builder pattern


### Return type

[**[]SamlIdentityProvider**](SamlIdentityProvider.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIdentityProvider

> SamlIdentityProvider UpdateIdentityProvider(ctx, id).SamlIdentityProviderRequest(samlIdentityProviderRequest).Execute()

Update a SAML configuration.



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
    id := "id_example" // string | Identifier of the SAML configuration to update.
    samlIdentityProviderRequest := *openapiclient.NewSamlIdentityProviderRequest("SumoLogic", "http://www.okta.com/abxcseyuiwelflkdjh", "X509cert1_example") // SamlIdentityProviderRequest | Information to update in the SAML configuration.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.SamlConfigurationManagementApi.UpdateIdentityProvider(context.Background(), id).SamlIdentityProviderRequest(samlIdentityProviderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SamlConfigurationManagementApi.UpdateIdentityProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateIdentityProvider`: SamlIdentityProvider
    fmt.Fprintf(os.Stdout, "Response from `SamlConfigurationManagementApi.UpdateIdentityProvider`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifier of the SAML configuration to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIdentityProviderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **samlIdentityProviderRequest** | [**SamlIdentityProviderRequest**](SamlIdentityProviderRequest.md) | Information to update in the SAML configuration. | 

### Return type

[**SamlIdentityProvider**](SamlIdentityProvider.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


# \PasswordPolicyApi

All URIs are relative to *https://api.au.sumologic.com/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPasswordPolicy**](PasswordPolicyApi.md#GetPasswordPolicy) | **Get** /v1/passwordPolicy | Get the current password policy.
[**SetPasswordPolicy**](PasswordPolicyApi.md#SetPasswordPolicy) | **Put** /v1/passwordPolicy | Update password policy.



## GetPasswordPolicy

> PasswordPolicy GetPasswordPolicy(ctx).Execute()

Get the current password policy.



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
    resp, r, err := api_client.PasswordPolicyApi.GetPasswordPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.GetPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyApi.GetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPasswordPolicyRequest struct via the builder pattern


### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetPasswordPolicy

> PasswordPolicy SetPasswordPolicy(ctx).PasswordPolicy(passwordPolicy).Execute()

Update password policy.



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
    passwordPolicy := *openapiclient.NewPasswordPolicy() // PasswordPolicy | 

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PasswordPolicyApi.SetPasswordPolicy(context.Background()).PasswordPolicy(passwordPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PasswordPolicyApi.SetPasswordPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetPasswordPolicy`: PasswordPolicy
    fmt.Fprintf(os.Stdout, "Response from `PasswordPolicyApi.SetPasswordPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetPasswordPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordPolicy** | [**PasswordPolicy**](PasswordPolicy.md) |  | 

### Return type

[**PasswordPolicy**](PasswordPolicy.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


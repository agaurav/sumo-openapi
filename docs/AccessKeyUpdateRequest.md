# AccessKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | **bool** | Indicates whether the access key is disabled or not. | 
**CorsHeaders** | Pointer to **[]string** | An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don&#39;t match any entry in the allowlist. | [optional] 

## Methods

### NewAccessKeyUpdateRequest

`func NewAccessKeyUpdateRequest(disabled bool, ) *AccessKeyUpdateRequest`

NewAccessKeyUpdateRequest instantiates a new AccessKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyUpdateRequestWithDefaults

`func NewAccessKeyUpdateRequestWithDefaults() *AccessKeyUpdateRequest`

NewAccessKeyUpdateRequestWithDefaults instantiates a new AccessKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *AccessKeyUpdateRequest) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccessKeyUpdateRequest) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccessKeyUpdateRequest) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetCorsHeaders

`func (o *AccessKeyUpdateRequest) GetCorsHeaders() []string`

GetCorsHeaders returns the CorsHeaders field if non-nil, zero value otherwise.

### GetCorsHeadersOk

`func (o *AccessKeyUpdateRequest) GetCorsHeadersOk() (*[]string, bool)`

GetCorsHeadersOk returns a tuple with the CorsHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHeaders

`func (o *AccessKeyUpdateRequest) SetCorsHeaders(v []string)`

SetCorsHeaders sets CorsHeaders field to given value.

### HasCorsHeaders

`func (o *AccessKeyUpdateRequest) HasCorsHeaders() bool`

HasCorsHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



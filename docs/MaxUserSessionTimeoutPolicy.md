# MaxUserSessionTimeoutPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxUserSessionTimeout** | **string** | Maximum web session timeout users are able to configure within their user preferences. Valid values are: &#x60;5m&#x60;, &#x60;15m&#x60;, &#x60;30m&#x60;, &#x60;1h&#x60;, &#x60;2h&#x60;, &#x60;6h&#x60;, &#x60;12h&#x60;, &#x60;1d&#x60;, &#x60;2d&#x60;, &#x60;3d&#x60;, &#x60;5d&#x60;, or &#x60;7d&#x60; | 

## Methods

### NewMaxUserSessionTimeoutPolicy

`func NewMaxUserSessionTimeoutPolicy(maxUserSessionTimeout string, ) *MaxUserSessionTimeoutPolicy`

NewMaxUserSessionTimeoutPolicy instantiates a new MaxUserSessionTimeoutPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxUserSessionTimeoutPolicyWithDefaults

`func NewMaxUserSessionTimeoutPolicyWithDefaults() *MaxUserSessionTimeoutPolicy`

NewMaxUserSessionTimeoutPolicyWithDefaults instantiates a new MaxUserSessionTimeoutPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxUserSessionTimeout

`func (o *MaxUserSessionTimeoutPolicy) GetMaxUserSessionTimeout() string`

GetMaxUserSessionTimeout returns the MaxUserSessionTimeout field if non-nil, zero value otherwise.

### GetMaxUserSessionTimeoutOk

`func (o *MaxUserSessionTimeoutPolicy) GetMaxUserSessionTimeoutOk() (*string, bool)`

GetMaxUserSessionTimeoutOk returns a tuple with the MaxUserSessionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUserSessionTimeout

`func (o *MaxUserSessionTimeoutPolicy) SetMaxUserSessionTimeout(v string)`

SetMaxUserSessionTimeout sets MaxUserSessionTimeout field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



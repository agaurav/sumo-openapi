# UserConcurrentSessionsLimitPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the User Concurrent Sessions Limit policy is enabled. | 
**MaxConcurrentSessions** | Pointer to **int32** | Maximum number of concurrent sessions a user may have. | [optional] [default to 100]

## Methods

### NewUserConcurrentSessionsLimitPolicy

`func NewUserConcurrentSessionsLimitPolicy(enabled bool, ) *UserConcurrentSessionsLimitPolicy`

NewUserConcurrentSessionsLimitPolicy instantiates a new UserConcurrentSessionsLimitPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserConcurrentSessionsLimitPolicyWithDefaults

`func NewUserConcurrentSessionsLimitPolicyWithDefaults() *UserConcurrentSessionsLimitPolicy`

NewUserConcurrentSessionsLimitPolicyWithDefaults instantiates a new UserConcurrentSessionsLimitPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UserConcurrentSessionsLimitPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserConcurrentSessionsLimitPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserConcurrentSessionsLimitPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxConcurrentSessions

`func (o *UserConcurrentSessionsLimitPolicy) GetMaxConcurrentSessions() int32`

GetMaxConcurrentSessions returns the MaxConcurrentSessions field if non-nil, zero value otherwise.

### GetMaxConcurrentSessionsOk

`func (o *UserConcurrentSessionsLimitPolicy) GetMaxConcurrentSessionsOk() (*int32, bool)`

GetMaxConcurrentSessionsOk returns a tuple with the MaxConcurrentSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentSessions

`func (o *UserConcurrentSessionsLimitPolicy) SetMaxConcurrentSessions(v int32)`

SetMaxConcurrentSessions sets MaxConcurrentSessions field to given value.

### HasMaxConcurrentSessions

`func (o *UserConcurrentSessionsLimitPolicy) HasMaxConcurrentSessions() bool`

HasMaxConcurrentSessions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



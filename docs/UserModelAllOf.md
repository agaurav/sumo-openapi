# UserModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the user. | 
**IsActive** | Pointer to **bool** | True if the user is active. | [optional] 
**IsLocked** | Pointer to **bool** | This has the value &#x60;true&#x60; if the user&#39;s account has been locked. If a user tries to log into their account several times and fails, his or her account will be locked for security reasons. | [optional] 
**IsMfaEnabled** | Pointer to **bool** | True if multi factor authentication is enabled for the user. | [optional] 
**LastLoginTimestamp** | Pointer to **time.Time** | Timestamp of the last login for the user in UTC. Will be null if the user has never logged in. | [optional] 

## Methods

### NewUserModelAllOf

`func NewUserModelAllOf(id string, ) *UserModelAllOf`

NewUserModelAllOf instantiates a new UserModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserModelAllOfWithDefaults

`func NewUserModelAllOfWithDefaults() *UserModelAllOf`

NewUserModelAllOfWithDefaults instantiates a new UserModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserModelAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserModelAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserModelAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetIsActive

`func (o *UserModelAllOf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UserModelAllOf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UserModelAllOf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UserModelAllOf) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetIsLocked

`func (o *UserModelAllOf) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UserModelAllOf) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UserModelAllOf) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *UserModelAllOf) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsMfaEnabled

`func (o *UserModelAllOf) GetIsMfaEnabled() bool`

GetIsMfaEnabled returns the IsMfaEnabled field if non-nil, zero value otherwise.

### GetIsMfaEnabledOk

`func (o *UserModelAllOf) GetIsMfaEnabledOk() (*bool, bool)`

GetIsMfaEnabledOk returns a tuple with the IsMfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMfaEnabled

`func (o *UserModelAllOf) SetIsMfaEnabled(v bool)`

SetIsMfaEnabled sets IsMfaEnabled field to given value.

### HasIsMfaEnabled

`func (o *UserModelAllOf) HasIsMfaEnabled() bool`

HasIsMfaEnabled returns a boolean if a field has been set.

### GetLastLoginTimestamp

`func (o *UserModelAllOf) GetLastLoginTimestamp() time.Time`

GetLastLoginTimestamp returns the LastLoginTimestamp field if non-nil, zero value otherwise.

### GetLastLoginTimestampOk

`func (o *UserModelAllOf) GetLastLoginTimestampOk() (*time.Time, bool)`

GetLastLoginTimestampOk returns a tuple with the LastLoginTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginTimestamp

`func (o *UserModelAllOf) SetLastLoginTimestamp(v time.Time)`

SetLastLoginTimestamp sets LastLoginTimestamp field to given value.

### HasLastLoginTimestamp

`func (o *UserModelAllOf) HasLastLoginTimestamp() bool`

HasLastLoginTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



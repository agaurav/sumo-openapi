# OAuthRefreshFailedTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExceptionType** | Pointer to **string** | The type of exception received while attempting OAuth. | [optional] 
**ExceptionMessage** | Pointer to **string** | The error message received with the failed OAuth request. | [optional] 

## Methods

### NewOAuthRefreshFailedTracker

`func NewOAuthRefreshFailedTracker() *OAuthRefreshFailedTracker`

NewOAuthRefreshFailedTracker instantiates a new OAuthRefreshFailedTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuthRefreshFailedTrackerWithDefaults

`func NewOAuthRefreshFailedTrackerWithDefaults() *OAuthRefreshFailedTracker`

NewOAuthRefreshFailedTrackerWithDefaults instantiates a new OAuthRefreshFailedTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExceptionType

`func (o *OAuthRefreshFailedTracker) GetExceptionType() string`

GetExceptionType returns the ExceptionType field if non-nil, zero value otherwise.

### GetExceptionTypeOk

`func (o *OAuthRefreshFailedTracker) GetExceptionTypeOk() (*string, bool)`

GetExceptionTypeOk returns a tuple with the ExceptionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionType

`func (o *OAuthRefreshFailedTracker) SetExceptionType(v string)`

SetExceptionType sets ExceptionType field to given value.

### HasExceptionType

`func (o *OAuthRefreshFailedTracker) HasExceptionType() bool`

HasExceptionType returns a boolean if a field has been set.

### GetExceptionMessage

`func (o *OAuthRefreshFailedTracker) GetExceptionMessage() string`

GetExceptionMessage returns the ExceptionMessage field if non-nil, zero value otherwise.

### GetExceptionMessageOk

`func (o *OAuthRefreshFailedTracker) GetExceptionMessageOk() (*string, bool)`

GetExceptionMessageOk returns a tuple with the ExceptionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionMessage

`func (o *OAuthRefreshFailedTracker) SetExceptionMessage(v string)`

SetExceptionMessage sets ExceptionMessage field to given value.

### HasExceptionMessage

`func (o *OAuthRefreshFailedTracker) HasExceptionMessage() bool`

HasExceptionMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



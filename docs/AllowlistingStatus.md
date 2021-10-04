# AllowlistingStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentEnabled** | **bool** | Whether service allowlisting is enabled for Content. | 
**LoginEnabled** | **bool** | Whether service allowlisting is enabled for Login. | 

## Methods

### NewAllowlistingStatus

`func NewAllowlistingStatus(contentEnabled bool, loginEnabled bool, ) *AllowlistingStatus`

NewAllowlistingStatus instantiates a new AllowlistingStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowlistingStatusWithDefaults

`func NewAllowlistingStatusWithDefaults() *AllowlistingStatus`

NewAllowlistingStatusWithDefaults instantiates a new AllowlistingStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentEnabled

`func (o *AllowlistingStatus) GetContentEnabled() bool`

GetContentEnabled returns the ContentEnabled field if non-nil, zero value otherwise.

### GetContentEnabledOk

`func (o *AllowlistingStatus) GetContentEnabledOk() (*bool, bool)`

GetContentEnabledOk returns a tuple with the ContentEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentEnabled

`func (o *AllowlistingStatus) SetContentEnabled(v bool)`

SetContentEnabled sets ContentEnabled field to given value.


### GetLoginEnabled

`func (o *AllowlistingStatus) GetLoginEnabled() bool`

GetLoginEnabled returns the LoginEnabled field if non-nil, zero value otherwise.

### GetLoginEnabledOk

`func (o *AllowlistingStatus) GetLoginEnabledOk() (*bool, bool)`

GetLoginEnabledOk returns a tuple with the LoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginEnabled

`func (o *AllowlistingStatus) SetLoginEnabled(v bool)`

SetLoginEnabled sets LoginEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



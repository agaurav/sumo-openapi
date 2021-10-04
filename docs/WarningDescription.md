# WarningDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Description of the warning. | 
**Cause** | Pointer to **string** | An optional cause of this warning. | [optional] 

## Methods

### NewWarningDescription

`func NewWarningDescription(message string, ) *WarningDescription`

NewWarningDescription instantiates a new WarningDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWarningDescriptionWithDefaults

`func NewWarningDescriptionWithDefaults() *WarningDescription`

NewWarningDescriptionWithDefaults instantiates a new WarningDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *WarningDescription) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WarningDescription) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WarningDescription) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCause

`func (o *WarningDescription) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *WarningDescription) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *WarningDescription) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *WarningDescription) HasCause() bool`

HasCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



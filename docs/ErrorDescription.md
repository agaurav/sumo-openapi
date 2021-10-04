# ErrorDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | An error code describing the type of error. | 
**Message** | **string** | A short English-language description of the error. | 
**Detail** | Pointer to **string** | An optional fuller English-language description of the error. | [optional] 
**Meta** | Pointer to **map[string]interface{}** | An optional list of metadata about the error. | [optional] 

## Methods

### NewErrorDescription

`func NewErrorDescription(code string, message string, ) *ErrorDescription`

NewErrorDescription instantiates a new ErrorDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorDescriptionWithDefaults

`func NewErrorDescriptionWithDefaults() *ErrorDescription`

NewErrorDescriptionWithDefaults instantiates a new ErrorDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorDescription) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorDescription) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorDescription) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ErrorDescription) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorDescription) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorDescription) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetail

`func (o *ErrorDescription) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ErrorDescription) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ErrorDescription) SetDetail(v string)`

SetDetail sets Detail field to given value.

### HasDetail

`func (o *ErrorDescription) HasDetail() bool`

HasDetail returns a boolean if a field has been set.

### GetMeta

`func (o *ErrorDescription) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ErrorDescription) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ErrorDescription) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ErrorDescription) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



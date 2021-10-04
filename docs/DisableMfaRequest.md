# DisableMfaRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | Email of user whose mfa is being disabled. | 
**Password** | **string** | Password of user whose mfa is being disabled. | 

## Methods

### NewDisableMfaRequest

`func NewDisableMfaRequest(email string, password string, ) *DisableMfaRequest`

NewDisableMfaRequest instantiates a new DisableMfaRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableMfaRequestWithDefaults

`func NewDisableMfaRequestWithDefaults() *DisableMfaRequest`

NewDisableMfaRequestWithDefaults instantiates a new DisableMfaRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *DisableMfaRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DisableMfaRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DisableMfaRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *DisableMfaRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DisableMfaRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DisableMfaRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



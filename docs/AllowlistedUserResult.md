# AllowlistedUserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | Unique identifier of the user. | 
**FirstName** | **string** | First name of the user. | 
**LastName** | **string** | Last name of the user. | 
**Email** | **string** | Email of the user. | 
**CanManageSaml** | **bool** | If the user can manage SAML Configurations. | 
**IsActive** | **bool** | Checks if the user is active. | 
**LastLogin** | **time.Time** | Timestamp of the last login of the user. | 

## Methods

### NewAllowlistedUserResult

`func NewAllowlistedUserResult(userId string, firstName string, lastName string, email string, canManageSaml bool, isActive bool, lastLogin time.Time, ) *AllowlistedUserResult`

NewAllowlistedUserResult instantiates a new AllowlistedUserResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAllowlistedUserResultWithDefaults

`func NewAllowlistedUserResultWithDefaults() *AllowlistedUserResult`

NewAllowlistedUserResultWithDefaults instantiates a new AllowlistedUserResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *AllowlistedUserResult) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AllowlistedUserResult) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AllowlistedUserResult) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetFirstName

`func (o *AllowlistedUserResult) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AllowlistedUserResult) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AllowlistedUserResult) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *AllowlistedUserResult) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AllowlistedUserResult) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AllowlistedUserResult) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *AllowlistedUserResult) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AllowlistedUserResult) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AllowlistedUserResult) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCanManageSaml

`func (o *AllowlistedUserResult) GetCanManageSaml() bool`

GetCanManageSaml returns the CanManageSaml field if non-nil, zero value otherwise.

### GetCanManageSamlOk

`func (o *AllowlistedUserResult) GetCanManageSamlOk() (*bool, bool)`

GetCanManageSamlOk returns a tuple with the CanManageSaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanManageSaml

`func (o *AllowlistedUserResult) SetCanManageSaml(v bool)`

SetCanManageSaml sets CanManageSaml field to given value.


### GetIsActive

`func (o *AllowlistedUserResult) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *AllowlistedUserResult) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *AllowlistedUserResult) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastLogin

`func (o *AllowlistedUserResult) GetLastLogin() time.Time`

GetLastLogin returns the LastLogin field if non-nil, zero value otherwise.

### GetLastLoginOk

`func (o *AllowlistedUserResult) GetLastLoginOk() (*time.Time, bool)`

GetLastLoginOk returns a tuple with the LastLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogin

`func (o *AllowlistedUserResult) SetLastLogin(v time.Time)`

SetLastLogin sets LastLogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



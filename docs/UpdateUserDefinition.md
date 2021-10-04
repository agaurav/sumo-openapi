# UpdateUserDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** | First name of the user. | 
**LastName** | **string** | Last name of the user. | 
**IsActive** | **bool** | This has the value &#x60;true&#x60; if the user is active and &#x60;false&#x60; if they have been deactivated. | 
**RoleIds** | **[]string** | List of role identifiers associated with the user. | 

## Methods

### NewUpdateUserDefinition

`func NewUpdateUserDefinition(firstName string, lastName string, isActive bool, roleIds []string, ) *UpdateUserDefinition`

NewUpdateUserDefinition instantiates a new UpdateUserDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserDefinitionWithDefaults

`func NewUpdateUserDefinitionWithDefaults() *UpdateUserDefinition`

NewUpdateUserDefinitionWithDefaults instantiates a new UpdateUserDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUserDefinition) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUserDefinition) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUserDefinition) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UpdateUserDefinition) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUserDefinition) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUserDefinition) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetIsActive

`func (o *UpdateUserDefinition) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateUserDefinition) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateUserDefinition) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetRoleIds

`func (o *UpdateUserDefinition) GetRoleIds() []string`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UpdateUserDefinition) GetRoleIdsOk() (*[]string, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UpdateUserDefinition) SetRoleIds(v []string)`

SetRoleIds sets RoleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



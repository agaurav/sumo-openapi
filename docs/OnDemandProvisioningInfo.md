# OnDemandProvisioningInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstNameAttribute** | Pointer to **string** | First name attribute of the new user account. | [optional] [default to ""]
**LastNameAttribute** | Pointer to **string** | Last name attribute of the new user account. | [optional] [default to ""]
**OnDemandProvisioningRoles** | **[]string** | Sumo Logic RBAC roles to be assigned when user accounts are provisioned. | [default to []]

## Methods

### NewOnDemandProvisioningInfo

`func NewOnDemandProvisioningInfo(onDemandProvisioningRoles []string, ) *OnDemandProvisioningInfo`

NewOnDemandProvisioningInfo instantiates a new OnDemandProvisioningInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnDemandProvisioningInfoWithDefaults

`func NewOnDemandProvisioningInfoWithDefaults() *OnDemandProvisioningInfo`

NewOnDemandProvisioningInfoWithDefaults instantiates a new OnDemandProvisioningInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstNameAttribute

`func (o *OnDemandProvisioningInfo) GetFirstNameAttribute() string`

GetFirstNameAttribute returns the FirstNameAttribute field if non-nil, zero value otherwise.

### GetFirstNameAttributeOk

`func (o *OnDemandProvisioningInfo) GetFirstNameAttributeOk() (*string, bool)`

GetFirstNameAttributeOk returns a tuple with the FirstNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstNameAttribute

`func (o *OnDemandProvisioningInfo) SetFirstNameAttribute(v string)`

SetFirstNameAttribute sets FirstNameAttribute field to given value.

### HasFirstNameAttribute

`func (o *OnDemandProvisioningInfo) HasFirstNameAttribute() bool`

HasFirstNameAttribute returns a boolean if a field has been set.

### GetLastNameAttribute

`func (o *OnDemandProvisioningInfo) GetLastNameAttribute() string`

GetLastNameAttribute returns the LastNameAttribute field if non-nil, zero value otherwise.

### GetLastNameAttributeOk

`func (o *OnDemandProvisioningInfo) GetLastNameAttributeOk() (*string, bool)`

GetLastNameAttributeOk returns a tuple with the LastNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastNameAttribute

`func (o *OnDemandProvisioningInfo) SetLastNameAttribute(v string)`

SetLastNameAttribute sets LastNameAttribute field to given value.

### HasLastNameAttribute

`func (o *OnDemandProvisioningInfo) HasLastNameAttribute() bool`

HasLastNameAttribute returns a boolean if a field has been set.

### GetOnDemandProvisioningRoles

`func (o *OnDemandProvisioningInfo) GetOnDemandProvisioningRoles() []string`

GetOnDemandProvisioningRoles returns the OnDemandProvisioningRoles field if non-nil, zero value otherwise.

### GetOnDemandProvisioningRolesOk

`func (o *OnDemandProvisioningInfo) GetOnDemandProvisioningRolesOk() (*[]string, bool)`

GetOnDemandProvisioningRolesOk returns a tuple with the OnDemandProvisioningRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandProvisioningRoles

`func (o *OnDemandProvisioningInfo) SetOnDemandProvisioningRoles(v []string)`

SetOnDemandProvisioningRoles sets OnDemandProvisioningRoles field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



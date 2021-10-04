# RoleModelAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the role. | 
**SystemDefined** | Pointer to **bool** | Role is system or user defined. | [optional] 

## Methods

### NewRoleModelAllOf

`func NewRoleModelAllOf(id string, ) *RoleModelAllOf`

NewRoleModelAllOf instantiates a new RoleModelAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleModelAllOfWithDefaults

`func NewRoleModelAllOfWithDefaults() *RoleModelAllOf`

NewRoleModelAllOfWithDefaults instantiates a new RoleModelAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RoleModelAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RoleModelAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RoleModelAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetSystemDefined

`func (o *RoleModelAllOf) GetSystemDefined() bool`

GetSystemDefined returns the SystemDefined field if non-nil, zero value otherwise.

### GetSystemDefinedOk

`func (o *RoleModelAllOf) GetSystemDefinedOk() (*bool, bool)`

GetSystemDefinedOk returns a tuple with the SystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemDefined

`func (o *RoleModelAllOf) SetSystemDefined(v bool)`

SetSystemDefined sets SystemDefined field to given value.

### HasSystemDefined

`func (o *RoleModelAllOf) HasSystemDefined() bool`

HasSystemDefined returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



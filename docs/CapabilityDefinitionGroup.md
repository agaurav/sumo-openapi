# CapabilityDefinitionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The backend name for the capability group | 
**Label** | **string** | The label for the capability group | 
**ParentId** | Pointer to **string** | The ID of the parent capability group | [optional] 

## Methods

### NewCapabilityDefinitionGroup

`func NewCapabilityDefinitionGroup(id string, label string, ) *CapabilityDefinitionGroup`

NewCapabilityDefinitionGroup instantiates a new CapabilityDefinitionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDefinitionGroupWithDefaults

`func NewCapabilityDefinitionGroupWithDefaults() *CapabilityDefinitionGroup`

NewCapabilityDefinitionGroupWithDefaults instantiates a new CapabilityDefinitionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapabilityDefinitionGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityDefinitionGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityDefinitionGroup) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *CapabilityDefinitionGroup) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CapabilityDefinitionGroup) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CapabilityDefinitionGroup) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetParentId

`func (o *CapabilityDefinitionGroup) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *CapabilityDefinitionGroup) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *CapabilityDefinitionGroup) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *CapabilityDefinitionGroup) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CapabilityDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The name of the capability | 
**Label** | **string** | The UI label for the capability. | 
**DependsOn** | **[]string** | Any capabilities that are required for this capability to be enabled. | 
**Group** | [**CapabilityDefinitionGroup**](CapabilityDefinitionGroup.md) |  | 
**Message** | Pointer to **string** | Warning message that appears when this capability is enabled. | [optional] 

## Methods

### NewCapabilityDefinition

`func NewCapabilityDefinition(id string, label string, dependsOn []string, group CapabilityDefinitionGroup, ) *CapabilityDefinition`

NewCapabilityDefinition instantiates a new CapabilityDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDefinitionWithDefaults

`func NewCapabilityDefinitionWithDefaults() *CapabilityDefinition`

NewCapabilityDefinitionWithDefaults instantiates a new CapabilityDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CapabilityDefinition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CapabilityDefinition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CapabilityDefinition) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *CapabilityDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CapabilityDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CapabilityDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDependsOn

`func (o *CapabilityDefinition) GetDependsOn() []string`

GetDependsOn returns the DependsOn field if non-nil, zero value otherwise.

### GetDependsOnOk

`func (o *CapabilityDefinition) GetDependsOnOk() (*[]string, bool)`

GetDependsOnOk returns a tuple with the DependsOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOn

`func (o *CapabilityDefinition) SetDependsOn(v []string)`

SetDependsOn sets DependsOn field to given value.


### GetGroup

`func (o *CapabilityDefinition) GetGroup() CapabilityDefinitionGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CapabilityDefinition) GetGroupOk() (*CapabilityDefinitionGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CapabilityDefinition) SetGroup(v CapabilityDefinitionGroup)`

SetGroup sets Group field to given value.


### GetMessage

`func (o *CapabilityDefinition) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CapabilityDefinition) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CapabilityDefinition) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CapabilityDefinition) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



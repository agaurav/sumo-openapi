# TopologySearchLabel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Key of a topology label to search for. | 
**Value** | **string** | Value of a topology label to search for. | 
**IsRequired** | Pointer to **bool** | Whether the content item is required to contain this label in order to be matched. If true, content items without this label will not be matched. If false, content items without this label will be matched.  | [optional] 

## Methods

### NewTopologySearchLabel

`func NewTopologySearchLabel(key string, value string, ) *TopologySearchLabel`

NewTopologySearchLabel instantiates a new TopologySearchLabel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologySearchLabelWithDefaults

`func NewTopologySearchLabelWithDefaults() *TopologySearchLabel`

NewTopologySearchLabelWithDefaults instantiates a new TopologySearchLabel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *TopologySearchLabel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TopologySearchLabel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TopologySearchLabel) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *TopologySearchLabel) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *TopologySearchLabel) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *TopologySearchLabel) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsRequired

`func (o *TopologySearchLabel) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *TopologySearchLabel) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *TopologySearchLabel) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *TopologySearchLabel) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



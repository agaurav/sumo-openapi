# Capacity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float64** | The value of the entitlement in units. | 
**Unit** | **string** | The unit of the entitlement. Units are provided in &#x60;GB&#x60; or &#x60;DPM&#x60;(data points per minute). | 
**CapacityType** | Pointer to **string** | Type of capacity. Valid values are: 1) &#x60;Paid&#x60; : This means that the capacity is chargeable. 2) &#x60;Free&#x60; : This means that this capacity is not chargeable. | [optional] 

## Methods

### NewCapacity

`func NewCapacity(value float64, unit string, ) *Capacity`

NewCapacity instantiates a new Capacity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapacityWithDefaults

`func NewCapacityWithDefaults() *Capacity`

NewCapacityWithDefaults instantiates a new Capacity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Capacity) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Capacity) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Capacity) SetValue(v float64)`

SetValue sets Value field to given value.


### GetUnit

`func (o *Capacity) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Capacity) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Capacity) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetCapacityType

`func (o *Capacity) GetCapacityType() string`

GetCapacityType returns the CapacityType field if non-nil, zero value otherwise.

### GetCapacityTypeOk

`func (o *Capacity) GetCapacityTypeOk() (*string, bool)`

GetCapacityTypeOk returns a tuple with the CapacityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityType

`func (o *Capacity) SetCapacityType(v string)`

SetCapacityType sets CapacityType field to given value.

### HasCapacityType

`func (o *Capacity) HasCapacityType() bool`

HasCapacityType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



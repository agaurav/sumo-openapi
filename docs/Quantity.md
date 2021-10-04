# Quantity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **int64** | The value of the consumable in units. | 
**Unit** | **string** | The unit of the consumable. Units are provided in: 1. &#x60;GB&#x60; 2. &#x60;DPM&#x60;(Data Points Per Minute) 3. &#x60;Credits&#x60; 4. &#x60;Days&#x60;  | 

## Methods

### NewQuantity

`func NewQuantity(value int64, unit string, ) *Quantity`

NewQuantity instantiates a new Quantity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQuantityWithDefaults

`func NewQuantityWithDefaults() *Quantity`

NewQuantityWithDefaults instantiates a new Quantity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *Quantity) GetValue() int64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Quantity) GetValueOk() (*int64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Quantity) SetValue(v int64)`

SetValue sets Value field to given value.


### GetUnit

`func (o *Quantity) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Quantity) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Quantity) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



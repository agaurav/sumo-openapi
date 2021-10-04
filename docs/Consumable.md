# Consumable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConsumableId** | **string** | Unique identifier of the consumable. Valid values are: 1. &#x60;Storage&#x60; 2. &#x60;Metrics&#x60; 3. &#x60;Continuous&#x60; 4. &#x60;Credits&#x60;  | 
**Quantity** | [**Quantity**](Quantity.md) |  | 

## Methods

### NewConsumable

`func NewConsumable(consumableId string, quantity Quantity, ) *Consumable`

NewConsumable instantiates a new Consumable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumableWithDefaults

`func NewConsumableWithDefaults() *Consumable`

NewConsumableWithDefaults instantiates a new Consumable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumableId

`func (o *Consumable) GetConsumableId() string`

GetConsumableId returns the ConsumableId field if non-nil, zero value otherwise.

### GetConsumableIdOk

`func (o *Consumable) GetConsumableIdOk() (*string, bool)`

GetConsumableIdOk returns a tuple with the ConsumableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumableId

`func (o *Consumable) SetConsumableId(v string)`

SetConsumableId sets ConsumableId field to given value.


### GetQuantity

`func (o *Consumable) GetQuantity() Quantity`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Consumable) GetQuantityOk() (*Quantity, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Consumable) SetQuantity(v Quantity)`

SetQuantity sets Quantity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



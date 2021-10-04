# DataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **float64** | The value of the data point in units. | 
**Unit** | **string** | The unit of the entitlement, possible values are &#x60;GB&#x60;, &#x60;DPM&#x60;, or &#x60;Credits&#x60;. | 

## Methods

### NewDataValue

`func NewDataValue(value float64, unit string, ) *DataValue`

NewDataValue instantiates a new DataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataValueWithDefaults

`func NewDataValueWithDefaults() *DataValue`

NewDataValueWithDefaults instantiates a new DataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *DataValue) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataValue) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataValue) SetValue(v float64)`

SetValue sets Value field to given value.


### GetUnit

`func (o *DataValue) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *DataValue) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *DataValue) SetUnit(v string)`

SetUnit sets Unit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Operator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]DataValue**](DataValue.md) | An array of objects denoting the value and unit of the results. | 
**Name** | **string** | The name of the operator applied to the data. | 

## Methods

### NewOperator

`func NewOperator(values []DataValue, name string, ) *Operator`

NewOperator instantiates a new Operator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorWithDefaults

`func NewOperatorWithDefaults() *Operator`

NewOperatorWithDefaults instantiates a new Operator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *Operator) GetValues() []DataValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Operator) GetValuesOk() (*[]DataValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Operator) SetValues(v []DataValue)`

SetValues sets Values field to given value.


### GetName

`func (o *Operator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Operator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Operator) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



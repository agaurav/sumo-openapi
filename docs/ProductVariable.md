# ProductVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductVariableName** | **string** | Name of a product variable. | 
**ProductVariableId** | **string** | Unique Identifier of the product variable. | 
**Unit** | **string** | Unit of measure for the productvariable. | 
**PossibleValues** | **[]int64** | Possible values allowed for the productvariable. | 

## Methods

### NewProductVariable

`func NewProductVariable(productVariableName string, productVariableId string, unit string, possibleValues []int64, ) *ProductVariable`

NewProductVariable instantiates a new ProductVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductVariableWithDefaults

`func NewProductVariableWithDefaults() *ProductVariable`

NewProductVariableWithDefaults instantiates a new ProductVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductVariableName

`func (o *ProductVariable) GetProductVariableName() string`

GetProductVariableName returns the ProductVariableName field if non-nil, zero value otherwise.

### GetProductVariableNameOk

`func (o *ProductVariable) GetProductVariableNameOk() (*string, bool)`

GetProductVariableNameOk returns a tuple with the ProductVariableName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariableName

`func (o *ProductVariable) SetProductVariableName(v string)`

SetProductVariableName sets ProductVariableName field to given value.


### GetProductVariableId

`func (o *ProductVariable) GetProductVariableId() string`

GetProductVariableId returns the ProductVariableId field if non-nil, zero value otherwise.

### GetProductVariableIdOk

`func (o *ProductVariable) GetProductVariableIdOk() (*string, bool)`

GetProductVariableIdOk returns a tuple with the ProductVariableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariableId

`func (o *ProductVariable) SetProductVariableId(v string)`

SetProductVariableId sets ProductVariableId field to given value.


### GetUnit

`func (o *ProductVariable) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *ProductVariable) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *ProductVariable) SetUnit(v string)`

SetUnit sets Unit field to given value.


### GetPossibleValues

`func (o *ProductVariable) GetPossibleValues() []int64`

GetPossibleValues returns the PossibleValues field if non-nil, zero value otherwise.

### GetPossibleValuesOk

`func (o *ProductVariable) GetPossibleValuesOk() (*[]int64, bool)`

GetPossibleValuesOk returns a tuple with the PossibleValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleValues

`func (o *ProductVariable) SetPossibleValues(v []int64)`

SetPossibleValues sets PossibleValues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



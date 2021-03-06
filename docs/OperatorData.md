# OperatorData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OperatorName** | **string** | The name of the metrics operator. | 
**Parameters** | [**[]OperatorParameter**](OperatorParameter.md) | A list of operator parameters for the operator data. | 

## Methods

### NewOperatorData

`func NewOperatorData(operatorName string, parameters []OperatorParameter, ) *OperatorData`

NewOperatorData instantiates a new OperatorData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorDataWithDefaults

`func NewOperatorDataWithDefaults() *OperatorData`

NewOperatorDataWithDefaults instantiates a new OperatorData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOperatorName

`func (o *OperatorData) GetOperatorName() string`

GetOperatorName returns the OperatorName field if non-nil, zero value otherwise.

### GetOperatorNameOk

`func (o *OperatorData) GetOperatorNameOk() (*string, bool)`

GetOperatorNameOk returns a tuple with the OperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorName

`func (o *OperatorData) SetOperatorName(v string)`

SetOperatorName sets OperatorName field to given value.


### GetParameters

`func (o *OperatorData) GetParameters() []OperatorParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *OperatorData) GetParametersOk() (*[]OperatorParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *OperatorData) SetParameters(v []OperatorParameter)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



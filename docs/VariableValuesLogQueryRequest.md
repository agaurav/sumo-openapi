# VariableValuesLogQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The original log query of the variable. | 
**Field** | **string** | A field in log query to populate the variable values. | 
**VariablesValues** | Pointer to [**VariablesValuesData**](VariablesValuesData.md) |  | [optional] 

## Methods

### NewVariableValuesLogQueryRequest

`func NewVariableValuesLogQueryRequest(query string, field string, ) *VariableValuesLogQueryRequest`

NewVariableValuesLogQueryRequest instantiates a new VariableValuesLogQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableValuesLogQueryRequestWithDefaults

`func NewVariableValuesLogQueryRequestWithDefaults() *VariableValuesLogQueryRequest`

NewVariableValuesLogQueryRequestWithDefaults instantiates a new VariableValuesLogQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *VariableValuesLogQueryRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *VariableValuesLogQueryRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *VariableValuesLogQueryRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetField

`func (o *VariableValuesLogQueryRequest) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *VariableValuesLogQueryRequest) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *VariableValuesLogQueryRequest) SetField(v string)`

SetField sets Field field to given value.


### GetVariablesValues

`func (o *VariableValuesLogQueryRequest) GetVariablesValues() VariablesValuesData`

GetVariablesValues returns the VariablesValues field if non-nil, zero value otherwise.

### GetVariablesValuesOk

`func (o *VariableValuesLogQueryRequest) GetVariablesValuesOk() (*VariablesValuesData, bool)`

GetVariablesValuesOk returns a tuple with the VariablesValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariablesValues

`func (o *VariableValuesLogQueryRequest) SetVariablesValues(v VariablesValuesData)`

SetVariablesValues sets VariablesValues field to given value.

### HasVariablesValues

`func (o *VariableValuesLogQueryRequest) HasVariablesValues() bool`

HasVariablesValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# VariableValuesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariableValues** | **[]string** | Values for the variable. | 
**Status** | Pointer to [**DashboardSearchStatus**](DashboardSearchStatus.md) |  | [optional] 
**VariableType** | Pointer to **string** | The type of the variable. | [optional] 
**Errors** | Pointer to [**[]ErrorDescription**](ErrorDescription.md) | Generic errors returned by backend from downstream assemblies. More specific errors will be thrown in the future. | [optional] 

## Methods

### NewVariableValuesData

`func NewVariableValuesData(variableValues []string, ) *VariableValuesData`

NewVariableValuesData instantiates a new VariableValuesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableValuesDataWithDefaults

`func NewVariableValuesDataWithDefaults() *VariableValuesData`

NewVariableValuesDataWithDefaults instantiates a new VariableValuesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariableValues

`func (o *VariableValuesData) GetVariableValues() []string`

GetVariableValues returns the VariableValues field if non-nil, zero value otherwise.

### GetVariableValuesOk

`func (o *VariableValuesData) GetVariableValuesOk() (*[]string, bool)`

GetVariableValuesOk returns a tuple with the VariableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValues

`func (o *VariableValuesData) SetVariableValues(v []string)`

SetVariableValues sets VariableValues field to given value.


### GetStatus

`func (o *VariableValuesData) GetStatus() DashboardSearchStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VariableValuesData) GetStatusOk() (*DashboardSearchStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VariableValuesData) SetStatus(v DashboardSearchStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VariableValuesData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariableType

`func (o *VariableValuesData) GetVariableType() string`

GetVariableType returns the VariableType field if non-nil, zero value otherwise.

### GetVariableTypeOk

`func (o *VariableValuesData) GetVariableTypeOk() (*string, bool)`

GetVariableTypeOk returns a tuple with the VariableType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableType

`func (o *VariableValuesData) SetVariableType(v string)`

SetVariableType sets VariableType field to given value.

### HasVariableType

`func (o *VariableValuesData) HasVariableType() bool`

HasVariableType returns a boolean if a field has been set.

### GetErrors

`func (o *VariableValuesData) GetErrors() []ErrorDescription`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VariableValuesData) GetErrorsOk() (*[]ErrorDescription, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VariableValuesData) SetErrors(v []ErrorDescription)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *VariableValuesData) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



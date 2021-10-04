# VariablesValuesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string][]string** | Data for variable values. | [default to {}]
**RichData** | Pointer to [**map[string]VariableValuesData**](VariableValuesData.md) | A rich form of data for the variable search, including variable values, status and variable type. This field is different from &#x60;data&#x60; in that it includes an object instead of list as the value in the map. The &#x60;data&#x60; field is kept for backwards compatibility, please use &#x60;richData&#x60; for all usages going forward. | [optional] 

## Methods

### NewVariablesValuesData

`func NewVariablesValuesData(data map[string][]string, ) *VariablesValuesData`

NewVariablesValuesData instantiates a new VariablesValuesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariablesValuesDataWithDefaults

`func NewVariablesValuesDataWithDefaults() *VariablesValuesData`

NewVariablesValuesDataWithDefaults instantiates a new VariablesValuesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *VariablesValuesData) GetData() map[string][]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *VariablesValuesData) GetDataOk() (*map[string][]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *VariablesValuesData) SetData(v map[string][]string)`

SetData sets Data field to given value.


### GetRichData

`func (o *VariablesValuesData) GetRichData() map[string]VariableValuesData`

GetRichData returns the RichData field if non-nil, zero value otherwise.

### GetRichDataOk

`func (o *VariablesValuesData) GetRichDataOk() (*map[string]VariableValuesData, bool)`

GetRichDataOk returns a tuple with the RichData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRichData

`func (o *VariablesValuesData) SetRichData(v map[string]VariableValuesData)`

SetRichData sets RichData field to given value.

### HasRichData

`func (o *VariablesValuesData) HasRichData() bool`

HasRichData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



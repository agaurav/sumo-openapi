# QueryParameterSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the parameter. | 
**Label** | **string** | The label of the parameter. | 
**Description** | **string** | A description of the parameter. | 
**DataType** | **string** | The data type of the parameter. Supported values are:   1. &#x60;NUMBER&#x60;   2. &#x60;STRING&#x60;   3. &#x60;QUERY_FRAGMENT&#x60;   4. &#x60;SEARCH_KEYWORD&#x60; | 
**Value** | **string** | A value for the parameter. Should be compatible with the type set in dataType field. | 
**AutoComplete** | [**ParameterAutoCompleteSyncDefinition**](ParameterAutoCompleteSyncDefinition.md) |  | 

## Methods

### NewQueryParameterSyncDefinition

`func NewQueryParameterSyncDefinition(name string, label string, description string, dataType string, value string, autoComplete ParameterAutoCompleteSyncDefinition, ) *QueryParameterSyncDefinition`

NewQueryParameterSyncDefinition instantiates a new QueryParameterSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryParameterSyncDefinitionWithDefaults

`func NewQueryParameterSyncDefinitionWithDefaults() *QueryParameterSyncDefinition`

NewQueryParameterSyncDefinitionWithDefaults instantiates a new QueryParameterSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *QueryParameterSyncDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueryParameterSyncDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueryParameterSyncDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *QueryParameterSyncDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *QueryParameterSyncDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *QueryParameterSyncDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *QueryParameterSyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *QueryParameterSyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *QueryParameterSyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDataType

`func (o *QueryParameterSyncDefinition) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *QueryParameterSyncDefinition) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *QueryParameterSyncDefinition) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetValue

`func (o *QueryParameterSyncDefinition) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *QueryParameterSyncDefinition) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *QueryParameterSyncDefinition) SetValue(v string)`

SetValue sets Value field to given value.


### GetAutoComplete

`func (o *QueryParameterSyncDefinition) GetAutoComplete() ParameterAutoCompleteSyncDefinition`

GetAutoComplete returns the AutoComplete field if non-nil, zero value otherwise.

### GetAutoCompleteOk

`func (o *QueryParameterSyncDefinition) GetAutoCompleteOk() (*ParameterAutoCompleteSyncDefinition, bool)`

GetAutoCompleteOk returns a tuple with the AutoComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoComplete

`func (o *QueryParameterSyncDefinition) SetAutoComplete(v ParameterAutoCompleteSyncDefinition)`

SetAutoComplete sets AutoComplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



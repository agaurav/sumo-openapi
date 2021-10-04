# ParameterAutoCompleteSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoCompleteType** | **string** | The autocomplete parameter type. Supported values are:   1. &#x60;SKIP_AUTOCOMPLETE&#x60;   2. &#x60;CSV_AUTOCOMPLETE&#x60;   3. &#x60;AUTOCOMPLETE_KEY&#x60;   4. &#x60;VALUE_ONLY_AUTOCOMPLETE&#x60;   5. &#x60;VALUE_ONLY_LOOKUP_AUTOCOMPLETE&#x60;   6. &#x60;LABEL_VALUE_LOOKUP_AUTOCOMPLETE&#x60; | 
**AutoCompleteKey** | Pointer to **string** | The autocomplete key to be used to fetch autocomplete values. | [optional] 
**AutoCompleteValues** | Pointer to [**[]AutoCompleteValueSyncDefinition**](AutoCompleteValueSyncDefinition.md) | The array of values of the corresponding autocomplete parameter. | [optional] 
**LookupFileName** | Pointer to **string** | The lookup file to use as a source for autocomplete values. | [optional] 
**LookupLabelColumn** | Pointer to **string** | The column from the lookup file to use for autocomplete labels. | [optional] 
**LookupValueColumn** | Pointer to **string** | The column from the lookup file to fill the actual value when a particular label is selected. | [optional] 

## Methods

### NewParameterAutoCompleteSyncDefinition

`func NewParameterAutoCompleteSyncDefinition(autoCompleteType string, ) *ParameterAutoCompleteSyncDefinition`

NewParameterAutoCompleteSyncDefinition instantiates a new ParameterAutoCompleteSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParameterAutoCompleteSyncDefinitionWithDefaults

`func NewParameterAutoCompleteSyncDefinitionWithDefaults() *ParameterAutoCompleteSyncDefinition`

NewParameterAutoCompleteSyncDefinitionWithDefaults instantiates a new ParameterAutoCompleteSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoCompleteType

`func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteType() string`

GetAutoCompleteType returns the AutoCompleteType field if non-nil, zero value otherwise.

### GetAutoCompleteTypeOk

`func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteTypeOk() (*string, bool)`

GetAutoCompleteTypeOk returns a tuple with the AutoCompleteType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteType

`func (o *ParameterAutoCompleteSyncDefinition) SetAutoCompleteType(v string)`

SetAutoCompleteType sets AutoCompleteType field to given value.


### GetAutoCompleteKey

`func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteKey() string`

GetAutoCompleteKey returns the AutoCompleteKey field if non-nil, zero value otherwise.

### GetAutoCompleteKeyOk

`func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteKeyOk() (*string, bool)`

GetAutoCompleteKeyOk returns a tuple with the AutoCompleteKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteKey

`func (o *ParameterAutoCompleteSyncDefinition) SetAutoCompleteKey(v string)`

SetAutoCompleteKey sets AutoCompleteKey field to given value.

### HasAutoCompleteKey

`func (o *ParameterAutoCompleteSyncDefinition) HasAutoCompleteKey() bool`

HasAutoCompleteKey returns a boolean if a field has been set.

### GetAutoCompleteValues

`func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteValues() []AutoCompleteValueSyncDefinition`

GetAutoCompleteValues returns the AutoCompleteValues field if non-nil, zero value otherwise.

### GetAutoCompleteValuesOk

`func (o *ParameterAutoCompleteSyncDefinition) GetAutoCompleteValuesOk() (*[]AutoCompleteValueSyncDefinition, bool)`

GetAutoCompleteValuesOk returns a tuple with the AutoCompleteValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCompleteValues

`func (o *ParameterAutoCompleteSyncDefinition) SetAutoCompleteValues(v []AutoCompleteValueSyncDefinition)`

SetAutoCompleteValues sets AutoCompleteValues field to given value.

### HasAutoCompleteValues

`func (o *ParameterAutoCompleteSyncDefinition) HasAutoCompleteValues() bool`

HasAutoCompleteValues returns a boolean if a field has been set.

### GetLookupFileName

`func (o *ParameterAutoCompleteSyncDefinition) GetLookupFileName() string`

GetLookupFileName returns the LookupFileName field if non-nil, zero value otherwise.

### GetLookupFileNameOk

`func (o *ParameterAutoCompleteSyncDefinition) GetLookupFileNameOk() (*string, bool)`

GetLookupFileNameOk returns a tuple with the LookupFileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupFileName

`func (o *ParameterAutoCompleteSyncDefinition) SetLookupFileName(v string)`

SetLookupFileName sets LookupFileName field to given value.

### HasLookupFileName

`func (o *ParameterAutoCompleteSyncDefinition) HasLookupFileName() bool`

HasLookupFileName returns a boolean if a field has been set.

### GetLookupLabelColumn

`func (o *ParameterAutoCompleteSyncDefinition) GetLookupLabelColumn() string`

GetLookupLabelColumn returns the LookupLabelColumn field if non-nil, zero value otherwise.

### GetLookupLabelColumnOk

`func (o *ParameterAutoCompleteSyncDefinition) GetLookupLabelColumnOk() (*string, bool)`

GetLookupLabelColumnOk returns a tuple with the LookupLabelColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupLabelColumn

`func (o *ParameterAutoCompleteSyncDefinition) SetLookupLabelColumn(v string)`

SetLookupLabelColumn sets LookupLabelColumn field to given value.

### HasLookupLabelColumn

`func (o *ParameterAutoCompleteSyncDefinition) HasLookupLabelColumn() bool`

HasLookupLabelColumn returns a boolean if a field has been set.

### GetLookupValueColumn

`func (o *ParameterAutoCompleteSyncDefinition) GetLookupValueColumn() string`

GetLookupValueColumn returns the LookupValueColumn field if non-nil, zero value otherwise.

### GetLookupValueColumnOk

`func (o *ParameterAutoCompleteSyncDefinition) GetLookupValueColumnOk() (*string, bool)`

GetLookupValueColumnOk returns a tuple with the LookupValueColumn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookupValueColumn

`func (o *ParameterAutoCompleteSyncDefinition) SetLookupValueColumn(v string)`

SetLookupValueColumn sets LookupValueColumn field to given value.

### HasLookupValueColumn

`func (o *ParameterAutoCompleteSyncDefinition) HasLookupValueColumn() bool`

HasLookupValueColumn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



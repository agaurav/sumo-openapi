# RowUpdateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Row** | [**[]TableRow**](TableRow.md) | A list of all the field identifiers and their corresponding values. | 

## Methods

### NewRowUpdateDefinition

`func NewRowUpdateDefinition(row []TableRow, ) *RowUpdateDefinition`

NewRowUpdateDefinition instantiates a new RowUpdateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRowUpdateDefinitionWithDefaults

`func NewRowUpdateDefinitionWithDefaults() *RowUpdateDefinition`

NewRowUpdateDefinitionWithDefaults instantiates a new RowUpdateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRow

`func (o *RowUpdateDefinition) GetRow() []TableRow`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *RowUpdateDefinition) GetRowOk() (*[]TableRow, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *RowUpdateDefinition) SetRow(v []TableRow)`

SetRow sets Row field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



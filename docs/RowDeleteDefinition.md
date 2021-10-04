# RowDeleteDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryKey** | [**[]TableRow**](TableRow.md) | A list of all primary key field identifiers and their corresponding values. | 

## Methods

### NewRowDeleteDefinition

`func NewRowDeleteDefinition(primaryKey []TableRow, ) *RowDeleteDefinition`

NewRowDeleteDefinition instantiates a new RowDeleteDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRowDeleteDefinitionWithDefaults

`func NewRowDeleteDefinitionWithDefaults() *RowDeleteDefinition`

NewRowDeleteDefinitionWithDefaults instantiates a new RowDeleteDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryKey

`func (o *RowDeleteDefinition) GetPrimaryKey() []TableRow`

GetPrimaryKey returns the PrimaryKey field if non-nil, zero value otherwise.

### GetPrimaryKeyOk

`func (o *RowDeleteDefinition) GetPrimaryKeyOk() (*[]TableRow, bool)`

GetPrimaryKeyOk returns a tuple with the PrimaryKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKey

`func (o *RowDeleteDefinition) SetPrimaryKey(v []TableRow)`

SetPrimaryKey sets PrimaryKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



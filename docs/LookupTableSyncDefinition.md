# LookupTableSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the lookup table. | 
**Fields** | [**[]LookupTableField**](LookupTableField.md) | The list of fields in the lookup table. | 
**PrimaryKeys** | **[]string** | The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain. | 
**Ttl** | Pointer to **int32** | A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically. | [optional] [default to 0]
**SizeLimitAction** | Pointer to **string** | The action that needs to be taken when the size limit is reached for the table. The possible values can be &#x60;StopIncomingMessages&#x60; or &#x60;DeleteOldData&#x60;. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached. | [optional] [default to "StopIncomingMessages"]

## Methods

### NewLookupTableSyncDefinition

`func NewLookupTableSyncDefinition(description string, fields []LookupTableField, primaryKeys []string, ) *LookupTableSyncDefinition`

NewLookupTableSyncDefinition instantiates a new LookupTableSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTableSyncDefinitionWithDefaults

`func NewLookupTableSyncDefinitionWithDefaults() *LookupTableSyncDefinition`

NewLookupTableSyncDefinitionWithDefaults instantiates a new LookupTableSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LookupTableSyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LookupTableSyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LookupTableSyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFields

`func (o *LookupTableSyncDefinition) GetFields() []LookupTableField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LookupTableSyncDefinition) GetFieldsOk() (*[]LookupTableField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LookupTableSyncDefinition) SetFields(v []LookupTableField)`

SetFields sets Fields field to given value.


### GetPrimaryKeys

`func (o *LookupTableSyncDefinition) GetPrimaryKeys() []string`

GetPrimaryKeys returns the PrimaryKeys field if non-nil, zero value otherwise.

### GetPrimaryKeysOk

`func (o *LookupTableSyncDefinition) GetPrimaryKeysOk() (*[]string, bool)`

GetPrimaryKeysOk returns a tuple with the PrimaryKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeys

`func (o *LookupTableSyncDefinition) SetPrimaryKeys(v []string)`

SetPrimaryKeys sets PrimaryKeys field to given value.


### GetTtl

`func (o *LookupTableSyncDefinition) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LookupTableSyncDefinition) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LookupTableSyncDefinition) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *LookupTableSyncDefinition) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetSizeLimitAction

`func (o *LookupTableSyncDefinition) GetSizeLimitAction() string`

GetSizeLimitAction returns the SizeLimitAction field if non-nil, zero value otherwise.

### GetSizeLimitActionOk

`func (o *LookupTableSyncDefinition) GetSizeLimitActionOk() (*string, bool)`

GetSizeLimitActionOk returns a tuple with the SizeLimitAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimitAction

`func (o *LookupTableSyncDefinition) SetSizeLimitAction(v string)`

SetSizeLimitAction sets SizeLimitAction field to given value.

### HasSizeLimitAction

`func (o *LookupTableSyncDefinition) HasSizeLimitAction() bool`

HasSizeLimitAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



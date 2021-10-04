# LookupTableDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the lookup table. | 
**Fields** | [**[]LookupTableField**](LookupTableField.md) | The list of fields in the lookup table. | 
**PrimaryKeys** | **[]string** | The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain. | 
**Ttl** | Pointer to **int32** | A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically. | [optional] [default to 0]
**SizeLimitAction** | Pointer to **string** | The action that needs to be taken when the size limit is reached for the table. The possible values can be &#x60;StopIncomingMessages&#x60; or &#x60;DeleteOldData&#x60;. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached. | [optional] [default to "StopIncomingMessages"]
**Name** | **string** | The name of the lookup table. | 
**ParentFolderId** | **string** | The parent-folder-path identifier of the lookup table in the Library. | 

## Methods

### NewLookupTableDefinition

`func NewLookupTableDefinition(description string, fields []LookupTableField, primaryKeys []string, name string, parentFolderId string, ) *LookupTableDefinition`

NewLookupTableDefinition instantiates a new LookupTableDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTableDefinitionWithDefaults

`func NewLookupTableDefinitionWithDefaults() *LookupTableDefinition`

NewLookupTableDefinitionWithDefaults instantiates a new LookupTableDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *LookupTableDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LookupTableDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LookupTableDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFields

`func (o *LookupTableDefinition) GetFields() []LookupTableField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LookupTableDefinition) GetFieldsOk() (*[]LookupTableField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LookupTableDefinition) SetFields(v []LookupTableField)`

SetFields sets Fields field to given value.


### GetPrimaryKeys

`func (o *LookupTableDefinition) GetPrimaryKeys() []string`

GetPrimaryKeys returns the PrimaryKeys field if non-nil, zero value otherwise.

### GetPrimaryKeysOk

`func (o *LookupTableDefinition) GetPrimaryKeysOk() (*[]string, bool)`

GetPrimaryKeysOk returns a tuple with the PrimaryKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeys

`func (o *LookupTableDefinition) SetPrimaryKeys(v []string)`

SetPrimaryKeys sets PrimaryKeys field to given value.


### GetTtl

`func (o *LookupTableDefinition) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LookupTableDefinition) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LookupTableDefinition) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *LookupTableDefinition) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetSizeLimitAction

`func (o *LookupTableDefinition) GetSizeLimitAction() string`

GetSizeLimitAction returns the SizeLimitAction field if non-nil, zero value otherwise.

### GetSizeLimitActionOk

`func (o *LookupTableDefinition) GetSizeLimitActionOk() (*string, bool)`

GetSizeLimitActionOk returns a tuple with the SizeLimitAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimitAction

`func (o *LookupTableDefinition) SetSizeLimitAction(v string)`

SetSizeLimitAction sets SizeLimitAction field to given value.

### HasSizeLimitAction

`func (o *LookupTableDefinition) HasSizeLimitAction() bool`

HasSizeLimitAction returns a boolean if a field has been set.

### GetName

`func (o *LookupTableDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LookupTableDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LookupTableDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *LookupTableDefinition) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *LookupTableDefinition) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *LookupTableDefinition) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



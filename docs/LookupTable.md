# LookupTable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Description** | **string** | The description of the lookup table. | 
**Fields** | [**[]LookupTableField**](LookupTableField.md) | The list of fields in the lookup table. | 
**PrimaryKeys** | **[]string** | The names of the fields that make up the primary key for the lookup table. These will be a subset of the fields that the table will contain. | 
**Ttl** | Pointer to **int32** | A time to live for each entry in the lookup table (in minutes). 365 days is the maximum time to live for each entry that you can specify. Setting it to 0 means that the records will not expire automatically. | [optional] [default to 0]
**SizeLimitAction** | Pointer to **string** | The action that needs to be taken when the size limit is reached for the table. The possible values can be &#x60;StopIncomingMessages&#x60; or &#x60;DeleteOldData&#x60;. DeleteOldData will start deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached. | [optional] [default to "StopIncomingMessages"]
**Name** | **string** | The name of the lookup table. | 
**ParentFolderId** | **string** | The parent-folder-path identifier of the lookup table in the Library. | 
**Id** | **string** | Identifier of the lookup table as a content item. | 
**ContentPath** | Pointer to **string** | Address/path of the parent folder of this lookup table in content library. For example, a lookup table existing  in the personal/lookupTable folder for user johndoe would be: /Library/Users/johndoe@acme.com/lookupTable | [optional] 
**Size** | Pointer to **int64** | The current size of the lookup table in bytes | [optional] 

## Methods

### NewLookupTable

`func NewLookupTable(createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, description string, fields []LookupTableField, primaryKeys []string, name string, parentFolderId string, id string, ) *LookupTable`

NewLookupTable instantiates a new LookupTable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTableWithDefaults

`func NewLookupTableWithDefaults() *LookupTable`

NewLookupTableWithDefaults instantiates a new LookupTable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LookupTable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LookupTable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LookupTable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *LookupTable) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *LookupTable) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *LookupTable) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *LookupTable) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *LookupTable) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *LookupTable) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *LookupTable) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *LookupTable) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *LookupTable) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetDescription

`func (o *LookupTable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LookupTable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LookupTable) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetFields

`func (o *LookupTable) GetFields() []LookupTableField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *LookupTable) GetFieldsOk() (*[]LookupTableField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *LookupTable) SetFields(v []LookupTableField)`

SetFields sets Fields field to given value.


### GetPrimaryKeys

`func (o *LookupTable) GetPrimaryKeys() []string`

GetPrimaryKeys returns the PrimaryKeys field if non-nil, zero value otherwise.

### GetPrimaryKeysOk

`func (o *LookupTable) GetPrimaryKeysOk() (*[]string, bool)`

GetPrimaryKeysOk returns a tuple with the PrimaryKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryKeys

`func (o *LookupTable) SetPrimaryKeys(v []string)`

SetPrimaryKeys sets PrimaryKeys field to given value.


### GetTtl

`func (o *LookupTable) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LookupTable) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LookupTable) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *LookupTable) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetSizeLimitAction

`func (o *LookupTable) GetSizeLimitAction() string`

GetSizeLimitAction returns the SizeLimitAction field if non-nil, zero value otherwise.

### GetSizeLimitActionOk

`func (o *LookupTable) GetSizeLimitActionOk() (*string, bool)`

GetSizeLimitActionOk returns a tuple with the SizeLimitAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimitAction

`func (o *LookupTable) SetSizeLimitAction(v string)`

SetSizeLimitAction sets SizeLimitAction field to given value.

### HasSizeLimitAction

`func (o *LookupTable) HasSizeLimitAction() bool`

HasSizeLimitAction returns a boolean if a field has been set.

### GetName

`func (o *LookupTable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LookupTable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LookupTable) SetName(v string)`

SetName sets Name field to given value.


### GetParentFolderId

`func (o *LookupTable) GetParentFolderId() string`

GetParentFolderId returns the ParentFolderId field if non-nil, zero value otherwise.

### GetParentFolderIdOk

`func (o *LookupTable) GetParentFolderIdOk() (*string, bool)`

GetParentFolderIdOk returns a tuple with the ParentFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentFolderId

`func (o *LookupTable) SetParentFolderId(v string)`

SetParentFolderId sets ParentFolderId field to given value.


### GetId

`func (o *LookupTable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LookupTable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LookupTable) SetId(v string)`

SetId sets Id field to given value.


### GetContentPath

`func (o *LookupTable) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *LookupTable) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *LookupTable) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *LookupTable) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### GetSize

`func (o *LookupTable) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LookupTable) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LookupTable) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *LookupTable) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



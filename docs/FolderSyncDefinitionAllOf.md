# FolderSyncDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | An optional description for the folder. | [optional] 
**Children** | [**[]ContentSyncDefinition**](ContentSyncDefinition.md) | The items in the folder, a list of Dashboard and/or Folder items. | 

## Methods

### NewFolderSyncDefinitionAllOf

`func NewFolderSyncDefinitionAllOf(children []ContentSyncDefinition, ) *FolderSyncDefinitionAllOf`

NewFolderSyncDefinitionAllOf instantiates a new FolderSyncDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderSyncDefinitionAllOfWithDefaults

`func NewFolderSyncDefinitionAllOfWithDefaults() *FolderSyncDefinitionAllOf`

NewFolderSyncDefinitionAllOfWithDefaults instantiates a new FolderSyncDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FolderSyncDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FolderSyncDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FolderSyncDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FolderSyncDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChildren

`func (o *FolderSyncDefinitionAllOf) GetChildren() []ContentSyncDefinition`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *FolderSyncDefinitionAllOf) GetChildrenOk() (*[]ContentSyncDefinition, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *FolderSyncDefinitionAllOf) SetChildren(v []ContentSyncDefinition)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



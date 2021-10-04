# FolderDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the folder. | 
**Description** | Pointer to **string** | The description of the folder. | [optional] 
**ParentId** | **string** | The identifier of the parent folder. | 

## Methods

### NewFolderDefinition

`func NewFolderDefinition(name string, parentId string, ) *FolderDefinition`

NewFolderDefinition instantiates a new FolderDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderDefinitionWithDefaults

`func NewFolderDefinitionWithDefaults() *FolderDefinition`

NewFolderDefinitionWithDefaults instantiates a new FolderDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *FolderDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *FolderDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FolderDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FolderDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FolderDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParentId

`func (o *FolderDefinition) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *FolderDefinition) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *FolderDefinition) SetParentId(v string)`

SetParentId sets ParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



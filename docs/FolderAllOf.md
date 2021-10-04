# FolderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the folder. | [optional] 
**Children** | Pointer to [**[]Content**](Content.md) | A list of the content items. | [optional] 

## Methods

### NewFolderAllOf

`func NewFolderAllOf() *FolderAllOf`

NewFolderAllOf instantiates a new FolderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderAllOfWithDefaults

`func NewFolderAllOfWithDefaults() *FolderAllOf`

NewFolderAllOfWithDefaults instantiates a new FolderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FolderAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FolderAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FolderAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FolderAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetChildren

`func (o *FolderAllOf) GetChildren() []Content`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *FolderAllOf) GetChildrenOk() (*[]Content, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *FolderAllOf) SetChildren(v []Content)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *FolderAllOf) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



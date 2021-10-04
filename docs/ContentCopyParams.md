# ContentCopyParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentId** | **string** | Identifier of the parent folder to copy to. | 
**Name** | Pointer to **string** | Optionally provide a new name. | [optional] 
**Description** | Pointer to **string** | Optionally provide a new description. | [optional] 

## Methods

### NewContentCopyParams

`func NewContentCopyParams(parentId string, ) *ContentCopyParams`

NewContentCopyParams instantiates a new ContentCopyParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentCopyParamsWithDefaults

`func NewContentCopyParamsWithDefaults() *ContentCopyParams`

NewContentCopyParamsWithDefaults instantiates a new ContentCopyParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentId

`func (o *ContentCopyParams) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ContentCopyParams) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ContentCopyParams) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetName

`func (o *ContentCopyParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ContentCopyParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ContentCopyParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ContentCopyParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ContentCopyParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ContentCopyParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ContentCopyParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ContentCopyParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



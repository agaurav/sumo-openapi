# TokenBaseDefinitionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the token. | 
**Description** | Pointer to **string** | Description of the token. | [optional] 
**Status** | **string** | Status of the token. Can be &#x60;Active&#x60;, or &#x60;Inactive&#x60;. | 
**Type** | **string** | Type of the token. Valid values: 1) CollectorRegistration | 
**Version** | **int64** | Version of the token. | 

## Methods

### NewTokenBaseDefinitionUpdate

`func NewTokenBaseDefinitionUpdate(name string, status string, type_ string, version int64, ) *TokenBaseDefinitionUpdate`

NewTokenBaseDefinitionUpdate instantiates a new TokenBaseDefinitionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBaseDefinitionUpdateWithDefaults

`func NewTokenBaseDefinitionUpdateWithDefaults() *TokenBaseDefinitionUpdate`

NewTokenBaseDefinitionUpdateWithDefaults instantiates a new TokenBaseDefinitionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TokenBaseDefinitionUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenBaseDefinitionUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenBaseDefinitionUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TokenBaseDefinitionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenBaseDefinitionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenBaseDefinitionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenBaseDefinitionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *TokenBaseDefinitionUpdate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenBaseDefinitionUpdate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenBaseDefinitionUpdate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *TokenBaseDefinitionUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenBaseDefinitionUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenBaseDefinitionUpdate) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *TokenBaseDefinitionUpdate) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TokenBaseDefinitionUpdate) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TokenBaseDefinitionUpdate) SetVersion(v int64)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



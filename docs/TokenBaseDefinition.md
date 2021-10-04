# TokenBaseDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the token. | 
**Description** | Pointer to **string** | Description of the token. | [optional] 
**Status** | **string** | Status of the token. Can be &#x60;Active&#x60;, or &#x60;Inactive&#x60;. | 
**Type** | **string** | Type of the token. Valid values: 1) CollectorRegistration | 

## Methods

### NewTokenBaseDefinition

`func NewTokenBaseDefinition(name string, status string, type_ string, ) *TokenBaseDefinition`

NewTokenBaseDefinition instantiates a new TokenBaseDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenBaseDefinitionWithDefaults

`func NewTokenBaseDefinitionWithDefaults() *TokenBaseDefinition`

NewTokenBaseDefinitionWithDefaults instantiates a new TokenBaseDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TokenBaseDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TokenBaseDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TokenBaseDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TokenBaseDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TokenBaseDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TokenBaseDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TokenBaseDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *TokenBaseDefinition) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TokenBaseDefinition) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TokenBaseDefinition) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetType

`func (o *TokenBaseDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TokenBaseDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TokenBaseDefinition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



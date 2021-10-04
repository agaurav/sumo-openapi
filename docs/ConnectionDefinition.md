# ConnectionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of connection. Valid values are &#x60;WebhookDefinition&#x60;, &#x60;ServiceNowDefinition&#x60;. | 
**Name** | **string** | Name of connection. Name should be a valid alphanumeric value. | 
**Description** | Pointer to **string** | Description of the connection. | [optional] [default to ""]

## Methods

### NewConnectionDefinition

`func NewConnectionDefinition(type_ string, name string, ) *ConnectionDefinition`

NewConnectionDefinition instantiates a new ConnectionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionDefinitionWithDefaults

`func NewConnectionDefinitionWithDefaults() *ConnectionDefinition`

NewConnectionDefinitionWithDefaults instantiates a new ConnectionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ConnectionDefinition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ConnectionDefinition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ConnectionDefinition) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *ConnectionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ConnectionDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ConnectionDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ConnectionDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ConnectionDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Cidr

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cidr** | **string** | The string representation of the CIDR notation or IP address. | 
**Description** | Pointer to **string** | Description of the CIDR notation or IP address. | [optional] 

## Methods

### NewCidr

`func NewCidr(cidr string, ) *Cidr`

NewCidr instantiates a new Cidr object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCidrWithDefaults

`func NewCidrWithDefaults() *Cidr`

NewCidrWithDefaults instantiates a new Cidr object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidr

`func (o *Cidr) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *Cidr) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *Cidr) SetCidr(v string)`

SetCidr sets Cidr field to given value.


### GetDescription

`func (o *Cidr) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Cidr) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Cidr) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Cidr) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



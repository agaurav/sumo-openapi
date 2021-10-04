# CapabilityMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | [**map[string]CapabilityDefinition**](CapabilityDefinition.md) | Map of capabilities to their attributes | 

## Methods

### NewCapabilityMap

`func NewCapabilityMap(capabilities map[string]CapabilityDefinition, ) *CapabilityMap`

NewCapabilityMap instantiates a new CapabilityMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityMapWithDefaults

`func NewCapabilityMapWithDefaults() *CapabilityMap`

NewCapabilityMapWithDefaults instantiates a new CapabilityMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *CapabilityMap) GetCapabilities() map[string]CapabilityDefinition`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CapabilityMap) GetCapabilitiesOk() (*map[string]CapabilityDefinition, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CapabilityMap) SetCapabilities(v map[string]CapabilityDefinition)`

SetCapabilities sets Capabilities field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



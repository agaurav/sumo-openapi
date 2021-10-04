# DynamicRuleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the dynamic parsing rule. Use a name that makes it easy to identify the rule. | 
**Scope** | **string** | Scope of the dynamic parsing rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You&#39;ll use the Scope to run a search against the rule. | 
**Enabled** | **bool** | Is the dynamic parsing rule enabled. | [default to true]

## Methods

### NewDynamicRuleDefinition

`func NewDynamicRuleDefinition(name string, scope string, enabled bool, ) *DynamicRuleDefinition`

NewDynamicRuleDefinition instantiates a new DynamicRuleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicRuleDefinitionWithDefaults

`func NewDynamicRuleDefinitionWithDefaults() *DynamicRuleDefinition`

NewDynamicRuleDefinitionWithDefaults instantiates a new DynamicRuleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DynamicRuleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicRuleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicRuleDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *DynamicRuleDefinition) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *DynamicRuleDefinition) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *DynamicRuleDefinition) SetScope(v string)`

SetScope sets Scope field to given value.


### GetEnabled

`func (o *DynamicRuleDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DynamicRuleDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DynamicRuleDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



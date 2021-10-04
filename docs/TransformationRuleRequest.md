# TransformationRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleDefinition** | [**TransformationRuleDefinition**](TransformationRuleDefinition.md) |  | 
**Enabled** | **bool** | True if the rule is enabled. | 

## Methods

### NewTransformationRuleRequest

`func NewTransformationRuleRequest(ruleDefinition TransformationRuleDefinition, enabled bool, ) *TransformationRuleRequest`

NewTransformationRuleRequest instantiates a new TransformationRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformationRuleRequestWithDefaults

`func NewTransformationRuleRequestWithDefaults() *TransformationRuleRequest`

NewTransformationRuleRequestWithDefaults instantiates a new TransformationRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleDefinition

`func (o *TransformationRuleRequest) GetRuleDefinition() TransformationRuleDefinition`

GetRuleDefinition returns the RuleDefinition field if non-nil, zero value otherwise.

### GetRuleDefinitionOk

`func (o *TransformationRuleRequest) GetRuleDefinitionOk() (*TransformationRuleDefinition, bool)`

GetRuleDefinitionOk returns a tuple with the RuleDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleDefinition

`func (o *TransformationRuleRequest) SetRuleDefinition(v TransformationRuleDefinition)`

SetRuleDefinition sets RuleDefinition field to given value.


### GetEnabled

`func (o *TransformationRuleRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TransformationRuleRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TransformationRuleRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



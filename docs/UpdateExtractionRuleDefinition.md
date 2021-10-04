# UpdateExtractionRuleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field extraction rule. Use a name that makes it easy to identify the rule. | 
**Scope** | **string** | Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You&#39;ll use the Scope to run a search against the rule. | 
**ParseExpression** | **string** | Describes the fields to be parsed. | 
**Enabled** | **bool** | Is the field extraction rule enabled. | 

## Methods

### NewUpdateExtractionRuleDefinition

`func NewUpdateExtractionRuleDefinition(name string, scope string, parseExpression string, enabled bool, ) *UpdateExtractionRuleDefinition`

NewUpdateExtractionRuleDefinition instantiates a new UpdateExtractionRuleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateExtractionRuleDefinitionWithDefaults

`func NewUpdateExtractionRuleDefinitionWithDefaults() *UpdateExtractionRuleDefinition`

NewUpdateExtractionRuleDefinitionWithDefaults instantiates a new UpdateExtractionRuleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateExtractionRuleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateExtractionRuleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateExtractionRuleDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *UpdateExtractionRuleDefinition) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateExtractionRuleDefinition) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateExtractionRuleDefinition) SetScope(v string)`

SetScope sets Scope field to given value.


### GetParseExpression

`func (o *UpdateExtractionRuleDefinition) GetParseExpression() string`

GetParseExpression returns the ParseExpression field if non-nil, zero value otherwise.

### GetParseExpressionOk

`func (o *UpdateExtractionRuleDefinition) GetParseExpressionOk() (*string, bool)`

GetParseExpressionOk returns a tuple with the ParseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseExpression

`func (o *UpdateExtractionRuleDefinition) SetParseExpression(v string)`

SetParseExpression sets ParseExpression field to given value.


### GetEnabled

`func (o *UpdateExtractionRuleDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateExtractionRuleDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateExtractionRuleDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



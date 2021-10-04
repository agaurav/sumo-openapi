# ExtractionRuleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field extraction rule. Use a name that makes it easy to identify the rule. | 
**Scope** | **string** | Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You&#39;ll use the Scope to run a search against the rule. | 
**ParseExpression** | **string** | Describes the fields to be parsed. | 
**Enabled** | Pointer to **bool** | Is the field extraction rule enabled. | [optional] [default to true]

## Methods

### NewExtractionRuleDefinition

`func NewExtractionRuleDefinition(name string, scope string, parseExpression string, ) *ExtractionRuleDefinition`

NewExtractionRuleDefinition instantiates a new ExtractionRuleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionRuleDefinitionWithDefaults

`func NewExtractionRuleDefinitionWithDefaults() *ExtractionRuleDefinition`

NewExtractionRuleDefinitionWithDefaults instantiates a new ExtractionRuleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtractionRuleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtractionRuleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtractionRuleDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *ExtractionRuleDefinition) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ExtractionRuleDefinition) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ExtractionRuleDefinition) SetScope(v string)`

SetScope sets Scope field to given value.


### GetParseExpression

`func (o *ExtractionRuleDefinition) GetParseExpression() string`

GetParseExpression returns the ParseExpression field if non-nil, zero value otherwise.

### GetParseExpressionOk

`func (o *ExtractionRuleDefinition) GetParseExpressionOk() (*string, bool)`

GetParseExpressionOk returns a tuple with the ParseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseExpression

`func (o *ExtractionRuleDefinition) SetParseExpression(v string)`

SetParseExpression sets ParseExpression field to given value.


### GetEnabled

`func (o *ExtractionRuleDefinition) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtractionRuleDefinition) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtractionRuleDefinition) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExtractionRuleDefinition) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



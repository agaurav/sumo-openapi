# BaseExtractionRuleDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field extraction rule. Use a name that makes it easy to identify the rule. | 
**Scope** | **string** | Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You&#39;ll use the Scope to run a search against the rule. | 
**ParseExpression** | **string** | Describes the fields to be parsed. | 

## Methods

### NewBaseExtractionRuleDefinition

`func NewBaseExtractionRuleDefinition(name string, scope string, parseExpression string, ) *BaseExtractionRuleDefinition`

NewBaseExtractionRuleDefinition instantiates a new BaseExtractionRuleDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseExtractionRuleDefinitionWithDefaults

`func NewBaseExtractionRuleDefinitionWithDefaults() *BaseExtractionRuleDefinition`

NewBaseExtractionRuleDefinitionWithDefaults instantiates a new BaseExtractionRuleDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BaseExtractionRuleDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseExtractionRuleDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseExtractionRuleDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *BaseExtractionRuleDefinition) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *BaseExtractionRuleDefinition) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *BaseExtractionRuleDefinition) SetScope(v string)`

SetScope sets Scope field to given value.


### GetParseExpression

`func (o *BaseExtractionRuleDefinition) GetParseExpression() string`

GetParseExpression returns the ParseExpression field if non-nil, zero value otherwise.

### GetParseExpressionOk

`func (o *BaseExtractionRuleDefinition) GetParseExpressionOk() (*string, bool)`

GetParseExpressionOk returns a tuple with the ParseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseExpression

`func (o *BaseExtractionRuleDefinition) SetParseExpression(v string)`

SetParseExpression sets ParseExpression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



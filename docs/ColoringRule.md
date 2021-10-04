# ColoringRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | **string** | Regex string to match queries to apply coloring to. | 
**SingleSeriesAggregateFunction** | **string** | Function to aggregate one series into one single value. | 
**MultipleSeriesAggregateFunction** | **string** | Function to aggregate the aggregate values of multiple time series into one single value. | 
**ColorThresholds** | Pointer to [**[]ColoringThreshold**](ColoringThreshold.md) | Color thresholds. | [optional] 

## Methods

### NewColoringRule

`func NewColoringRule(scope string, singleSeriesAggregateFunction string, multipleSeriesAggregateFunction string, ) *ColoringRule`

NewColoringRule instantiates a new ColoringRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColoringRuleWithDefaults

`func NewColoringRuleWithDefaults() *ColoringRule`

NewColoringRuleWithDefaults instantiates a new ColoringRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *ColoringRule) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ColoringRule) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ColoringRule) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSingleSeriesAggregateFunction

`func (o *ColoringRule) GetSingleSeriesAggregateFunction() string`

GetSingleSeriesAggregateFunction returns the SingleSeriesAggregateFunction field if non-nil, zero value otherwise.

### GetSingleSeriesAggregateFunctionOk

`func (o *ColoringRule) GetSingleSeriesAggregateFunctionOk() (*string, bool)`

GetSingleSeriesAggregateFunctionOk returns a tuple with the SingleSeriesAggregateFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingleSeriesAggregateFunction

`func (o *ColoringRule) SetSingleSeriesAggregateFunction(v string)`

SetSingleSeriesAggregateFunction sets SingleSeriesAggregateFunction field to given value.


### GetMultipleSeriesAggregateFunction

`func (o *ColoringRule) GetMultipleSeriesAggregateFunction() string`

GetMultipleSeriesAggregateFunction returns the MultipleSeriesAggregateFunction field if non-nil, zero value otherwise.

### GetMultipleSeriesAggregateFunctionOk

`func (o *ColoringRule) GetMultipleSeriesAggregateFunctionOk() (*string, bool)`

GetMultipleSeriesAggregateFunctionOk returns a tuple with the MultipleSeriesAggregateFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSeriesAggregateFunction

`func (o *ColoringRule) SetMultipleSeriesAggregateFunction(v string)`

SetMultipleSeriesAggregateFunction sets MultipleSeriesAggregateFunction field to given value.


### GetColorThresholds

`func (o *ColoringRule) GetColorThresholds() []ColoringThreshold`

GetColorThresholds returns the ColorThresholds field if non-nil, zero value otherwise.

### GetColorThresholdsOk

`func (o *ColoringRule) GetColorThresholdsOk() (*[]ColoringThreshold, bool)`

GetColorThresholdsOk returns a tuple with the ColorThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorThresholds

`func (o *ColoringRule) SetColorThresholds(v []ColoringThreshold)`

SetColorThresholds sets ColorThresholds field to given value.

### HasColorThresholds

`func (o *ColoringRule) HasColorThresholds() bool`

HasColorThresholds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



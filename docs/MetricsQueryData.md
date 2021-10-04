# MetricsQueryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | **string** | The metric of the query. | 
**AggregationType** | Pointer to **string** | The type of aggregation. Can be &#x60;Count&#x60;, &#x60;Minimum&#x60;, &#x60;Maximum&#x60;, &#x60;Sum&#x60;, &#x60;Average&#x60; or &#x60;None&#x60;. | [optional] 
**GroupBy** | Pointer to **string** | The field to group the results by. | [optional] 
**Filters** | [**[]MetricsFilter**](MetricsFilter.md) | A list of filters for the metrics query. | 
**Operators** | Pointer to [**[]OperatorData**](OperatorData.md) | A list of operator data for the metrics query. | [optional] 

## Methods

### NewMetricsQueryData

`func NewMetricsQueryData(metric string, filters []MetricsFilter, ) *MetricsQueryData`

NewMetricsQueryData instantiates a new MetricsQueryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryDataWithDefaults

`func NewMetricsQueryDataWithDefaults() *MetricsQueryData`

NewMetricsQueryDataWithDefaults instantiates a new MetricsQueryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MetricsQueryData) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricsQueryData) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricsQueryData) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetAggregationType

`func (o *MetricsQueryData) GetAggregationType() string`

GetAggregationType returns the AggregationType field if non-nil, zero value otherwise.

### GetAggregationTypeOk

`func (o *MetricsQueryData) GetAggregationTypeOk() (*string, bool)`

GetAggregationTypeOk returns a tuple with the AggregationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationType

`func (o *MetricsQueryData) SetAggregationType(v string)`

SetAggregationType sets AggregationType field to given value.

### HasAggregationType

`func (o *MetricsQueryData) HasAggregationType() bool`

HasAggregationType returns a boolean if a field has been set.

### GetGroupBy

`func (o *MetricsQueryData) GetGroupBy() string`

GetGroupBy returns the GroupBy field if non-nil, zero value otherwise.

### GetGroupByOk

`func (o *MetricsQueryData) GetGroupByOk() (*string, bool)`

GetGroupByOk returns a tuple with the GroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupBy

`func (o *MetricsQueryData) SetGroupBy(v string)`

SetGroupBy sets GroupBy field to given value.

### HasGroupBy

`func (o *MetricsQueryData) HasGroupBy() bool`

HasGroupBy returns a boolean if a field has been set.

### GetFilters

`func (o *MetricsQueryData) GetFilters() []MetricsFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *MetricsQueryData) GetFiltersOk() (*[]MetricsFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *MetricsQueryData) SetFilters(v []MetricsFilter)`

SetFilters sets Filters field to given value.


### GetOperators

`func (o *MetricsQueryData) GetOperators() []OperatorData`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *MetricsQueryData) GetOperatorsOk() (*[]OperatorData, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *MetricsQueryData) SetOperators(v []OperatorData)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *MetricsQueryData) HasOperators() bool`

HasOperators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



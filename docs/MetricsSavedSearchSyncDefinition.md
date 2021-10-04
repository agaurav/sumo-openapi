# MetricsSavedSearchSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Item description in the content library. | [optional] 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**LogQuery** | Pointer to **string** | Query used to add an overlay to the chart. | [optional] 
**MetricsQueries** | [**[]MetricsSavedSearchQuerySyncDefinition**](MetricsSavedSearchQuerySyncDefinition.md) | Metrics queries. | 
**DesiredQuantizationInSecs** | **int32** | Desired quantization in seconds. | 
**Properties** | Pointer to **string** | Chart properties. This field is optional. | [optional] 

## Methods

### NewMetricsSavedSearchSyncDefinition

`func NewMetricsSavedSearchSyncDefinition(timeRange ResolvableTimeRange, metricsQueries []MetricsSavedSearchQuerySyncDefinition, desiredQuantizationInSecs int32, ) *MetricsSavedSearchSyncDefinition`

NewMetricsSavedSearchSyncDefinition instantiates a new MetricsSavedSearchSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSavedSearchSyncDefinitionWithDefaults

`func NewMetricsSavedSearchSyncDefinitionWithDefaults() *MetricsSavedSearchSyncDefinition`

NewMetricsSavedSearchSyncDefinitionWithDefaults instantiates a new MetricsSavedSearchSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MetricsSavedSearchSyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsSavedSearchSyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsSavedSearchSyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsSavedSearchSyncDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTimeRange

`func (o *MetricsSavedSearchSyncDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsSavedSearchSyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsSavedSearchSyncDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetLogQuery

`func (o *MetricsSavedSearchSyncDefinition) GetLogQuery() string`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *MetricsSavedSearchSyncDefinition) GetLogQueryOk() (*string, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *MetricsSavedSearchSyncDefinition) SetLogQuery(v string)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *MetricsSavedSearchSyncDefinition) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetMetricsQueries

`func (o *MetricsSavedSearchSyncDefinition) GetMetricsQueries() []MetricsSavedSearchQuerySyncDefinition`

GetMetricsQueries returns the MetricsQueries field if non-nil, zero value otherwise.

### GetMetricsQueriesOk

`func (o *MetricsSavedSearchSyncDefinition) GetMetricsQueriesOk() (*[]MetricsSavedSearchQuerySyncDefinition, bool)`

GetMetricsQueriesOk returns a tuple with the MetricsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsQueries

`func (o *MetricsSavedSearchSyncDefinition) SetMetricsQueries(v []MetricsSavedSearchQuerySyncDefinition)`

SetMetricsQueries sets MetricsQueries field to given value.


### GetDesiredQuantizationInSecs

`func (o *MetricsSavedSearchSyncDefinition) GetDesiredQuantizationInSecs() int32`

GetDesiredQuantizationInSecs returns the DesiredQuantizationInSecs field if non-nil, zero value otherwise.

### GetDesiredQuantizationInSecsOk

`func (o *MetricsSavedSearchSyncDefinition) GetDesiredQuantizationInSecsOk() (*int32, bool)`

GetDesiredQuantizationInSecsOk returns a tuple with the DesiredQuantizationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQuantizationInSecs

`func (o *MetricsSavedSearchSyncDefinition) SetDesiredQuantizationInSecs(v int32)`

SetDesiredQuantizationInSecs sets DesiredQuantizationInSecs field to given value.


### GetProperties

`func (o *MetricsSavedSearchSyncDefinition) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MetricsSavedSearchSyncDefinition) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MetricsSavedSearchSyncDefinition) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MetricsSavedSearchSyncDefinition) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



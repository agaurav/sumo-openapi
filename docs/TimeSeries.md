# TimeSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricDefinition** | [**MetricDefinition**](MetricDefinition.md) |  | 
**Points** | [**Points**](Points.md) |  | 

## Methods

### NewTimeSeries

`func NewTimeSeries(metricDefinition MetricDefinition, points Points, ) *TimeSeries`

NewTimeSeries instantiates a new TimeSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesWithDefaults

`func NewTimeSeriesWithDefaults() *TimeSeries`

NewTimeSeriesWithDefaults instantiates a new TimeSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricDefinition

`func (o *TimeSeries) GetMetricDefinition() MetricDefinition`

GetMetricDefinition returns the MetricDefinition field if non-nil, zero value otherwise.

### GetMetricDefinitionOk

`func (o *TimeSeries) GetMetricDefinitionOk() (*MetricDefinition, bool)`

GetMetricDefinitionOk returns a tuple with the MetricDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricDefinition

`func (o *TimeSeries) SetMetricDefinition(v MetricDefinition)`

SetMetricDefinition sets MetricDefinition field to given value.


### GetPoints

`func (o *TimeSeries) GetPoints() Points`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *TimeSeries) GetPointsOk() (*Points, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *TimeSeries) SetPoints(v Points)`

SetPoints sets Points field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



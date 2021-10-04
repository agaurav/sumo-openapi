# MetricsMissingDataCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TriggerSource** | **string** | Determines which time series from queries to use for Metrics MissingData and ResolvedMissingData triggers Valid values:   1. &#x60;AllTimeSeries&#x60;: Evaluate the condition against all time series. (NOTE: This option is only valid if monitorType is &#x60;Metrics&#x60;)   2. &#x60;AnyTimeSeries&#x60;: Evaluate the condition against any time series. (NOTE: This option is only valid if monitorType is &#x60;Metrics&#x60;)   3. &#x60;AllResults&#x60;: Evaluate the condition against results from all queries. (NOTE: This option is only valid if monitorType is &#x60;Logs&#x60;) | 
**TimeRange** | **string** | The relative time range of the monitor. | 

## Methods

### NewMetricsMissingDataCondition

`func NewMetricsMissingDataCondition(triggerSource string, timeRange string, ) *MetricsMissingDataCondition`

NewMetricsMissingDataCondition instantiates a new MetricsMissingDataCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsMissingDataConditionWithDefaults

`func NewMetricsMissingDataConditionWithDefaults() *MetricsMissingDataCondition`

NewMetricsMissingDataConditionWithDefaults instantiates a new MetricsMissingDataCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTriggerSource

`func (o *MetricsMissingDataCondition) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *MetricsMissingDataCondition) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *MetricsMissingDataCondition) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.


### GetTimeRange

`func (o *MetricsMissingDataCondition) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsMissingDataCondition) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsMissingDataCondition) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



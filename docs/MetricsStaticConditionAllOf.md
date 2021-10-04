# MetricsStaticConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** | The relative time range of the monitor. | 
**Threshold** | **float64** | The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [default to 0.0]
**ThresholdType** | **string** | The comparison type for the &#x60;threshold&#x60; evaluation. This defines how you want the data value compared. Valid values:   1. &#x60;LessThan&#x60;: Less than than the configured threshold.   2. &#x60;GreaterThan&#x60;: Greater than the configured threshold.   3. &#x60;LessThanOrEqual&#x60;: Less than or equal to the configured threshold.   4. &#x60;GreaterThanOrEqual&#x60;: Greater than or equal to the configured threshold. ThresholdType value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [default to "GreaterThanOrEqual"]
**OccurrenceType** | **string** | The criteria to evaluate the threshold and thresholdType in the given time range. Valid values:   1. &#x60;AtLeastOnce&#x60;: Trigger if the threshold is met at least once. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   2. &#x60;Always&#x60;: Trigger if the threshold is met continuously. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   3. &#x60;ResultCount&#x60;: Trigger if the threshold is met against the count of results. (NOTE: This is the only valid value if monitorType is &#x60;Logs&#x60;.)   4. &#x60;MissingData&#x60;: Trigger if the data is missing. (NOTE: This is valid for both &#x60;Logs&#x60; and &#x60;Metrics&#x60; monitorTypes) | 

## Methods

### NewMetricsStaticConditionAllOf

`func NewMetricsStaticConditionAllOf(timeRange string, threshold float64, thresholdType string, occurrenceType string, ) *MetricsStaticConditionAllOf`

NewMetricsStaticConditionAllOf instantiates a new MetricsStaticConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsStaticConditionAllOfWithDefaults

`func NewMetricsStaticConditionAllOfWithDefaults() *MetricsStaticConditionAllOf`

NewMetricsStaticConditionAllOfWithDefaults instantiates a new MetricsStaticConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *MetricsStaticConditionAllOf) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsStaticConditionAllOf) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsStaticConditionAllOf) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetThreshold

`func (o *MetricsStaticConditionAllOf) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricsStaticConditionAllOf) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricsStaticConditionAllOf) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetThresholdType

`func (o *MetricsStaticConditionAllOf) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *MetricsStaticConditionAllOf) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *MetricsStaticConditionAllOf) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.


### GetOccurrenceType

`func (o *MetricsStaticConditionAllOf) GetOccurrenceType() string`

GetOccurrenceType returns the OccurrenceType field if non-nil, zero value otherwise.

### GetOccurrenceTypeOk

`func (o *MetricsStaticConditionAllOf) GetOccurrenceTypeOk() (*string, bool)`

GetOccurrenceTypeOk returns a tuple with the OccurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceType

`func (o *MetricsStaticConditionAllOf) SetOccurrenceType(v string)`

SetOccurrenceType sets OccurrenceType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



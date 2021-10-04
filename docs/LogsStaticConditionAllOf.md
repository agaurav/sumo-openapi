# LogsStaticConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** | The relative time range of the monitor. | 
**Threshold** | **float64** | The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [default to 0.0]
**ThresholdType** | **string** | The comparison type for the &#x60;threshold&#x60; evaluation. This defines how you want the data value compared. Valid values:   1. &#x60;LessThan&#x60;: Less than than the configured threshold.   2. &#x60;GreaterThan&#x60;: Greater than the configured threshold.   3. &#x60;LessThanOrEqual&#x60;: Less than or equal to the configured threshold.   4. &#x60;GreaterThanOrEqual&#x60;: Greater than or equal to the configured threshold. ThresholdType value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [default to "GreaterThanOrEqual"]
**Field** | Pointer to **string** | The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If &#x60;field&#x60; is not specified, monitor would default to result count instead. | [optional] 

## Methods

### NewLogsStaticConditionAllOf

`func NewLogsStaticConditionAllOf(timeRange string, threshold float64, thresholdType string, ) *LogsStaticConditionAllOf`

NewLogsStaticConditionAllOf instantiates a new LogsStaticConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsStaticConditionAllOfWithDefaults

`func NewLogsStaticConditionAllOfWithDefaults() *LogsStaticConditionAllOf`

NewLogsStaticConditionAllOfWithDefaults instantiates a new LogsStaticConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *LogsStaticConditionAllOf) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *LogsStaticConditionAllOf) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *LogsStaticConditionAllOf) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetThreshold

`func (o *LogsStaticConditionAllOf) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *LogsStaticConditionAllOf) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *LogsStaticConditionAllOf) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.


### GetThresholdType

`func (o *LogsStaticConditionAllOf) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *LogsStaticConditionAllOf) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *LogsStaticConditionAllOf) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.


### GetField

`func (o *LogsStaticConditionAllOf) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *LogsStaticConditionAllOf) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *LogsStaticConditionAllOf) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *LogsStaticConditionAllOf) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# StaticCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | **string** | The relative time range of the monitor. | 
**Threshold** | Pointer to **float64** | The data value for the condition. This defines the threshold for when to trigger. Threshold value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [optional] [default to 0.0]
**ThresholdType** | Pointer to **string** | The comparison type for the &#x60;threshold&#x60; evaluation. This defines how you want the data value compared. Valid values:   1. &#x60;LessThan&#x60;: Less than than the configured threshold.   2. &#x60;GreaterThan&#x60;: Greater than the configured threshold.   3. &#x60;LessThanOrEqual&#x60;: Less than or equal to the configured threshold.   4. &#x60;GreaterThanOrEqual&#x60;: Greater than or equal to the configured threshold. ThresholdType value is not applicable for &#x60;MissingData&#x60; and &#x60;ResolvedMissingData&#x60; triggerTypes and will be ignored if specified. | [optional] [default to "GreaterThanOrEqual"]
**Field** | Pointer to **string** | The name of the field that the trigger condition will alert on. The trigger could compare the value of specified field with the threshold. If &#x60;field&#x60; is not specified, monitor would default to result count instead. | [optional] 
**OccurrenceType** | **string** | The criteria to evaluate the threshold and thresholdType in the given time range. Valid values:   1. &#x60;AtLeastOnce&#x60;: Trigger if the threshold is met at least once. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   2. &#x60;Always&#x60;: Trigger if the threshold is met continuously. (NOTE: This is the only valid value if monitorType is &#x60;Metrics&#x60;.)   3. &#x60;ResultCount&#x60;: Trigger if the threshold is met against the count of results. (NOTE: This is the only valid value if monitorType is &#x60;Logs&#x60;.)   4. &#x60;MissingData&#x60;: Trigger if the data is missing. (NOTE: This is valid for both &#x60;Logs&#x60; and &#x60;Metrics&#x60; monitorTypes) | 
**TriggerSource** | **string** | Determines which time series from queries to use for Metrics MissingData and ResolvedMissingData triggers Valid values:   1. &#x60;AllTimeSeries&#x60;: Evaluate the condition against all time series. (NOTE: This option is only valid if monitorType is &#x60;Metrics&#x60;)   2. &#x60;AnyTimeSeries&#x60;: Evaluate the condition against any time series. (NOTE: This option is only valid if monitorType is &#x60;Metrics&#x60;)   3. &#x60;AllResults&#x60;: Evaluate the condition against results from all queries. (NOTE: This option is only valid if monitorType is &#x60;Logs&#x60;) | 

## Methods

### NewStaticCondition

`func NewStaticCondition(timeRange string, occurrenceType string, triggerSource string, ) *StaticCondition`

NewStaticCondition instantiates a new StaticCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticConditionWithDefaults

`func NewStaticConditionWithDefaults() *StaticCondition`

NewStaticConditionWithDefaults instantiates a new StaticCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *StaticCondition) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *StaticCondition) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *StaticCondition) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetThreshold

`func (o *StaticCondition) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *StaticCondition) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *StaticCondition) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *StaticCondition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetThresholdType

`func (o *StaticCondition) GetThresholdType() string`

GetThresholdType returns the ThresholdType field if non-nil, zero value otherwise.

### GetThresholdTypeOk

`func (o *StaticCondition) GetThresholdTypeOk() (*string, bool)`

GetThresholdTypeOk returns a tuple with the ThresholdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdType

`func (o *StaticCondition) SetThresholdType(v string)`

SetThresholdType sets ThresholdType field to given value.

### HasThresholdType

`func (o *StaticCondition) HasThresholdType() bool`

HasThresholdType returns a boolean if a field has been set.

### GetField

`func (o *StaticCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *StaticCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *StaticCondition) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *StaticCondition) HasField() bool`

HasField returns a boolean if a field has been set.

### GetOccurrenceType

`func (o *StaticCondition) GetOccurrenceType() string`

GetOccurrenceType returns the OccurrenceType field if non-nil, zero value otherwise.

### GetOccurrenceTypeOk

`func (o *StaticCondition) GetOccurrenceTypeOk() (*string, bool)`

GetOccurrenceTypeOk returns a tuple with the OccurrenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrenceType

`func (o *StaticCondition) SetOccurrenceType(v string)`

SetOccurrenceType sets OccurrenceType field to given value.


### GetTriggerSource

`func (o *StaticCondition) GetTriggerSource() string`

GetTriggerSource returns the TriggerSource field if non-nil, zero value otherwise.

### GetTriggerSourceOk

`func (o *StaticCondition) GetTriggerSourceOk() (*string, bool)`

GetTriggerSourceOk returns a tuple with the TriggerSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerSource

`func (o *StaticCondition) SetTriggerSource(v string)`

SetTriggerSource sets TriggerSource field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



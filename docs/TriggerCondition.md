# TriggerCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DetectionMethod** | Pointer to **string** | Detection method of the trigger condition. Valid values:   1. &#x60;StaticCondition&#x60;: A condition that triggers based off of a static threshold.   2. &#x60;LogsStaticCondition&#x60;: A logs condition that triggers based off of a static threshold. Currently LogsStaticCondition is available in closed beta (Notify your Sumo Logic representative in order to get the early access).   3. &#x60;MetricsStaticCondition&#x60;: A metrics condition that triggers based off of a static threshold. Currently MetricsStaticCondition is available in closed beta (Notify your Sumo Logic representative in order to get the early access).   4. &#x60;LogsOutlierCondition&#x60;: A logs condition that triggers based off of a dynamic outlier threshold. Currently LogsOutlierCondition is available in closed beta (Notify your Sumo Logic representative in order to get the early access).   5. &#x60;MetricsOutlierCondition&#x60;: A metrics condition that triggers based off of a dynamic outlier threshold. Currently MetricsOutlierCondition is available in closed beta (Notify your Sumo Logic representative in order to get the early access).   6. &#x60;LogsMissingDataCondition&#x60;: A logs missing data condition that triggers based off of no data available. Currently LogsMissingDataCondition is available in closed beta (Notify your Sumo Logic representative in order to get the early access).   7. &#x60;MetricsMissingDataCondition&#x60;: A metrics missing data condition that triggers based off of no data available. Currently MetricsMissingDataCondition is available in closed beta (Notify your Sumo Logic representative in order to get the early access). | [optional] [default to "StaticCondition"]
**TriggerType** | **string** | The type of trigger condition. Valid values:   1. &#x60;Critical&#x60;: A critical condition to trigger on.   2. &#x60;Warning&#x60;: A warning condition to trigger on.   3. &#x60;MissingData&#x60;: A condition that indicates data is missing.   4. &#x60;ResolvedCritical&#x60;: A condition to resolve a Critical trigger on.   5. &#x60;ResolvedWarning&#x60;: A condition to resolve a Warning trigger on.   6. &#x60;ResolvedMissingData&#x60;: A condition to resolve a MissingData trigger. | 

## Methods

### NewTriggerCondition

`func NewTriggerCondition(triggerType string, ) *TriggerCondition`

NewTriggerCondition instantiates a new TriggerCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerConditionWithDefaults

`func NewTriggerConditionWithDefaults() *TriggerCondition`

NewTriggerConditionWithDefaults instantiates a new TriggerCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetectionMethod

`func (o *TriggerCondition) GetDetectionMethod() string`

GetDetectionMethod returns the DetectionMethod field if non-nil, zero value otherwise.

### GetDetectionMethodOk

`func (o *TriggerCondition) GetDetectionMethodOk() (*string, bool)`

GetDetectionMethodOk returns a tuple with the DetectionMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectionMethod

`func (o *TriggerCondition) SetDetectionMethod(v string)`

SetDetectionMethod sets DetectionMethod field to given value.

### HasDetectionMethod

`func (o *TriggerCondition) HasDetectionMethod() bool`

HasDetectionMethod returns a boolean if a field has been set.

### GetTriggerType

`func (o *TriggerCondition) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *TriggerCondition) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *TriggerCondition) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



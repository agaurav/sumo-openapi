# MetricsOutlierCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaselineWindow** | Pointer to **string** | The time range used to compute the baseline. | [optional] [default to "1d"]
**Direction** | Pointer to **string** | Specifies which direction should trigger violations. | [optional] [default to "Both"]
**Threshold** | Pointer to **float64** | How much should the indicator be different from the baseline for each datapoint. | [optional] [default to 3.0]

## Methods

### NewMetricsOutlierCondition

`func NewMetricsOutlierCondition() *MetricsOutlierCondition`

NewMetricsOutlierCondition instantiates a new MetricsOutlierCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsOutlierConditionWithDefaults

`func NewMetricsOutlierConditionWithDefaults() *MetricsOutlierCondition`

NewMetricsOutlierConditionWithDefaults instantiates a new MetricsOutlierCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaselineWindow

`func (o *MetricsOutlierCondition) GetBaselineWindow() string`

GetBaselineWindow returns the BaselineWindow field if non-nil, zero value otherwise.

### GetBaselineWindowOk

`func (o *MetricsOutlierCondition) GetBaselineWindowOk() (*string, bool)`

GetBaselineWindowOk returns a tuple with the BaselineWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindow

`func (o *MetricsOutlierCondition) SetBaselineWindow(v string)`

SetBaselineWindow sets BaselineWindow field to given value.

### HasBaselineWindow

`func (o *MetricsOutlierCondition) HasBaselineWindow() bool`

HasBaselineWindow returns a boolean if a field has been set.

### GetDirection

`func (o *MetricsOutlierCondition) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MetricsOutlierCondition) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MetricsOutlierCondition) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MetricsOutlierCondition) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetThreshold

`func (o *MetricsOutlierCondition) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricsOutlierCondition) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricsOutlierCondition) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MetricsOutlierCondition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



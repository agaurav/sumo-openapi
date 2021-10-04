# MetricsOutlierConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaselineWindow** | Pointer to **string** | The time range used to compute the baseline. | [optional] [default to "1d"]
**Direction** | Pointer to **string** | Specifies which direction should trigger violations. | [optional] [default to "Both"]
**Threshold** | Pointer to **float64** | How much should the indicator be different from the baseline for each datapoint. | [optional] [default to 3.0]

## Methods

### NewMetricsOutlierConditionAllOf

`func NewMetricsOutlierConditionAllOf() *MetricsOutlierConditionAllOf`

NewMetricsOutlierConditionAllOf instantiates a new MetricsOutlierConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsOutlierConditionAllOfWithDefaults

`func NewMetricsOutlierConditionAllOfWithDefaults() *MetricsOutlierConditionAllOf`

NewMetricsOutlierConditionAllOfWithDefaults instantiates a new MetricsOutlierConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaselineWindow

`func (o *MetricsOutlierConditionAllOf) GetBaselineWindow() string`

GetBaselineWindow returns the BaselineWindow field if non-nil, zero value otherwise.

### GetBaselineWindowOk

`func (o *MetricsOutlierConditionAllOf) GetBaselineWindowOk() (*string, bool)`

GetBaselineWindowOk returns a tuple with the BaselineWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindow

`func (o *MetricsOutlierConditionAllOf) SetBaselineWindow(v string)`

SetBaselineWindow sets BaselineWindow field to given value.

### HasBaselineWindow

`func (o *MetricsOutlierConditionAllOf) HasBaselineWindow() bool`

HasBaselineWindow returns a boolean if a field has been set.

### GetDirection

`func (o *MetricsOutlierConditionAllOf) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MetricsOutlierConditionAllOf) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MetricsOutlierConditionAllOf) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MetricsOutlierConditionAllOf) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetThreshold

`func (o *MetricsOutlierConditionAllOf) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricsOutlierConditionAllOf) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricsOutlierConditionAllOf) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MetricsOutlierConditionAllOf) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



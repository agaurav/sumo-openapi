# MetricsOutlier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrimmedQuery** | Pointer to **string** | The query string after trimming out the outlier clause. | [optional] 
**BaselineWindow** | Pointer to **string** | The time range used to compute the baseline. | [optional] [default to "5m"]
**BaselineTimeRangeWindow** | Pointer to [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] 
**Direction** | Pointer to **string** | Specifies which direction should trigger violations. Valid values:   1. &#x60;Both&#x60;: Both positive and negative deviations   2. &#x60;Up&#x60;: Positive deviations only   3. &#x60;Down&#x60;: Negative deviations only example: \&quot;Up\&quot; pattern: \&quot;^(Both|Up|Down)$\&quot; default: \&quot;Both\&quot; x-pattern-message: \&quot;should be one of the following: &#39;Both&#39;, &#39;Up&#39;, &#39;Down&#39;\&quot; | [optional] 
**Threshold** | Pointer to **float64** | How much should the indicator be different from the baseline for each datapoint. | [optional] [default to 3.0]

## Methods

### NewMetricsOutlier

`func NewMetricsOutlier() *MetricsOutlier`

NewMetricsOutlier instantiates a new MetricsOutlier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsOutlierWithDefaults

`func NewMetricsOutlierWithDefaults() *MetricsOutlier`

NewMetricsOutlierWithDefaults instantiates a new MetricsOutlier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrimmedQuery

`func (o *MetricsOutlier) GetTrimmedQuery() string`

GetTrimmedQuery returns the TrimmedQuery field if non-nil, zero value otherwise.

### GetTrimmedQueryOk

`func (o *MetricsOutlier) GetTrimmedQueryOk() (*string, bool)`

GetTrimmedQueryOk returns a tuple with the TrimmedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimmedQuery

`func (o *MetricsOutlier) SetTrimmedQuery(v string)`

SetTrimmedQuery sets TrimmedQuery field to given value.

### HasTrimmedQuery

`func (o *MetricsOutlier) HasTrimmedQuery() bool`

HasTrimmedQuery returns a boolean if a field has been set.

### GetBaselineWindow

`func (o *MetricsOutlier) GetBaselineWindow() string`

GetBaselineWindow returns the BaselineWindow field if non-nil, zero value otherwise.

### GetBaselineWindowOk

`func (o *MetricsOutlier) GetBaselineWindowOk() (*string, bool)`

GetBaselineWindowOk returns a tuple with the BaselineWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineWindow

`func (o *MetricsOutlier) SetBaselineWindow(v string)`

SetBaselineWindow sets BaselineWindow field to given value.

### HasBaselineWindow

`func (o *MetricsOutlier) HasBaselineWindow() bool`

HasBaselineWindow returns a boolean if a field has been set.

### GetBaselineTimeRangeWindow

`func (o *MetricsOutlier) GetBaselineTimeRangeWindow() ResolvableTimeRange`

GetBaselineTimeRangeWindow returns the BaselineTimeRangeWindow field if non-nil, zero value otherwise.

### GetBaselineTimeRangeWindowOk

`func (o *MetricsOutlier) GetBaselineTimeRangeWindowOk() (*ResolvableTimeRange, bool)`

GetBaselineTimeRangeWindowOk returns a tuple with the BaselineTimeRangeWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselineTimeRangeWindow

`func (o *MetricsOutlier) SetBaselineTimeRangeWindow(v ResolvableTimeRange)`

SetBaselineTimeRangeWindow sets BaselineTimeRangeWindow field to given value.

### HasBaselineTimeRangeWindow

`func (o *MetricsOutlier) HasBaselineTimeRangeWindow() bool`

HasBaselineTimeRangeWindow returns a boolean if a field has been set.

### GetDirection

`func (o *MetricsOutlier) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MetricsOutlier) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MetricsOutlier) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *MetricsOutlier) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetThreshold

`func (o *MetricsOutlier) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *MetricsOutlier) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *MetricsOutlier) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *MetricsOutlier) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



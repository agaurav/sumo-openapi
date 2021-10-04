# LogsOutlier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrimmedQuery** | Pointer to **string** | The query string after trimming out the outlier clause. | [optional] 
**Window** | Pointer to **int64** | Sets the trailing number of data points to calculate mean and sigma. | [optional] [default to 10]
**Consecutive** | Pointer to **int64** | Sets the required number of consecutive indicator data points (outliers) to trigger a violation. | [optional] [default to 1]
**Direction** | Pointer to **string** | Specifies which direction should trigger violations. Valid values:   1. &#x60;Both&#x60;: Both positive and negative deviations   2. &#x60;Up&#x60;: Positive deviations only   3. &#x60;Down&#x60;: Negative deviations only example: \&quot;Up\&quot; pattern: \&quot;^(Both|Up|Down)$\&quot; default: \&quot;Both\&quot; x-pattern-message: \&quot;should be one of the following: &#39;Both&#39;, &#39;Up&#39;, &#39;Down&#39;\&quot; | [optional] 
**Threshold** | Pointer to **float64** | Sets the number of standard deviations for calculating violations. | [optional] [default to 3.0]
**Field** | Pointer to **string** | The name of the field that the trigger condition will alert on. | [optional] 

## Methods

### NewLogsOutlier

`func NewLogsOutlier() *LogsOutlier`

NewLogsOutlier instantiates a new LogsOutlier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsOutlierWithDefaults

`func NewLogsOutlierWithDefaults() *LogsOutlier`

NewLogsOutlierWithDefaults instantiates a new LogsOutlier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrimmedQuery

`func (o *LogsOutlier) GetTrimmedQuery() string`

GetTrimmedQuery returns the TrimmedQuery field if non-nil, zero value otherwise.

### GetTrimmedQueryOk

`func (o *LogsOutlier) GetTrimmedQueryOk() (*string, bool)`

GetTrimmedQueryOk returns a tuple with the TrimmedQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrimmedQuery

`func (o *LogsOutlier) SetTrimmedQuery(v string)`

SetTrimmedQuery sets TrimmedQuery field to given value.

### HasTrimmedQuery

`func (o *LogsOutlier) HasTrimmedQuery() bool`

HasTrimmedQuery returns a boolean if a field has been set.

### GetWindow

`func (o *LogsOutlier) GetWindow() int64`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LogsOutlier) GetWindowOk() (*int64, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LogsOutlier) SetWindow(v int64)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *LogsOutlier) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetConsecutive

`func (o *LogsOutlier) GetConsecutive() int64`

GetConsecutive returns the Consecutive field if non-nil, zero value otherwise.

### GetConsecutiveOk

`func (o *LogsOutlier) GetConsecutiveOk() (*int64, bool)`

GetConsecutiveOk returns a tuple with the Consecutive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutive

`func (o *LogsOutlier) SetConsecutive(v int64)`

SetConsecutive sets Consecutive field to given value.

### HasConsecutive

`func (o *LogsOutlier) HasConsecutive() bool`

HasConsecutive returns a boolean if a field has been set.

### GetDirection

`func (o *LogsOutlier) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LogsOutlier) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LogsOutlier) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LogsOutlier) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetThreshold

`func (o *LogsOutlier) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *LogsOutlier) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *LogsOutlier) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *LogsOutlier) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetField

`func (o *LogsOutlier) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *LogsOutlier) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *LogsOutlier) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *LogsOutlier) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



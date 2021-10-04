# LogsOutlierCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Window** | Pointer to **int64** | Sets the trailing number of data points to calculate mean and sigma. | [optional] [default to 50]
**Consecutive** | Pointer to **int64** | Sets the required number of consecutive indicator data points (outliers) to trigger a violation. | [optional] [default to 1]
**Direction** | Pointer to **string** | Specifies which direction should trigger violations. | [optional] [default to "Both"]
**Threshold** | Pointer to **float64** | Sets the number of standard deviations for calculating violations. | [optional] [default to 3.0]
**Field** | Pointer to **string** | The name of the field that the trigger condition will alert on. | [optional] 

## Methods

### NewLogsOutlierCondition

`func NewLogsOutlierCondition() *LogsOutlierCondition`

NewLogsOutlierCondition instantiates a new LogsOutlierCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsOutlierConditionWithDefaults

`func NewLogsOutlierConditionWithDefaults() *LogsOutlierCondition`

NewLogsOutlierConditionWithDefaults instantiates a new LogsOutlierCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindow

`func (o *LogsOutlierCondition) GetWindow() int64`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LogsOutlierCondition) GetWindowOk() (*int64, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LogsOutlierCondition) SetWindow(v int64)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *LogsOutlierCondition) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetConsecutive

`func (o *LogsOutlierCondition) GetConsecutive() int64`

GetConsecutive returns the Consecutive field if non-nil, zero value otherwise.

### GetConsecutiveOk

`func (o *LogsOutlierCondition) GetConsecutiveOk() (*int64, bool)`

GetConsecutiveOk returns a tuple with the Consecutive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutive

`func (o *LogsOutlierCondition) SetConsecutive(v int64)`

SetConsecutive sets Consecutive field to given value.

### HasConsecutive

`func (o *LogsOutlierCondition) HasConsecutive() bool`

HasConsecutive returns a boolean if a field has been set.

### GetDirection

`func (o *LogsOutlierCondition) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LogsOutlierCondition) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LogsOutlierCondition) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LogsOutlierCondition) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetThreshold

`func (o *LogsOutlierCondition) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *LogsOutlierCondition) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *LogsOutlierCondition) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *LogsOutlierCondition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetField

`func (o *LogsOutlierCondition) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *LogsOutlierCondition) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *LogsOutlierCondition) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *LogsOutlierCondition) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



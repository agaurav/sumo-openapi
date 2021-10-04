# LogsOutlierConditionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Window** | Pointer to **int64** | Sets the trailing number of data points to calculate mean and sigma. | [optional] [default to 50]
**Consecutive** | Pointer to **int64** | Sets the required number of consecutive indicator data points (outliers) to trigger a violation. | [optional] [default to 1]
**Direction** | Pointer to **string** | Specifies which direction should trigger violations. | [optional] [default to "Both"]
**Threshold** | Pointer to **float64** | Sets the number of standard deviations for calculating violations. | [optional] [default to 3.0]
**Field** | Pointer to **string** | The name of the field that the trigger condition will alert on. | [optional] 

## Methods

### NewLogsOutlierConditionAllOf

`func NewLogsOutlierConditionAllOf() *LogsOutlierConditionAllOf`

NewLogsOutlierConditionAllOf instantiates a new LogsOutlierConditionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogsOutlierConditionAllOfWithDefaults

`func NewLogsOutlierConditionAllOfWithDefaults() *LogsOutlierConditionAllOf`

NewLogsOutlierConditionAllOfWithDefaults instantiates a new LogsOutlierConditionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWindow

`func (o *LogsOutlierConditionAllOf) GetWindow() int64`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *LogsOutlierConditionAllOf) GetWindowOk() (*int64, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *LogsOutlierConditionAllOf) SetWindow(v int64)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *LogsOutlierConditionAllOf) HasWindow() bool`

HasWindow returns a boolean if a field has been set.

### GetConsecutive

`func (o *LogsOutlierConditionAllOf) GetConsecutive() int64`

GetConsecutive returns the Consecutive field if non-nil, zero value otherwise.

### GetConsecutiveOk

`func (o *LogsOutlierConditionAllOf) GetConsecutiveOk() (*int64, bool)`

GetConsecutiveOk returns a tuple with the Consecutive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutive

`func (o *LogsOutlierConditionAllOf) SetConsecutive(v int64)`

SetConsecutive sets Consecutive field to given value.

### HasConsecutive

`func (o *LogsOutlierConditionAllOf) HasConsecutive() bool`

HasConsecutive returns a boolean if a field has been set.

### GetDirection

`func (o *LogsOutlierConditionAllOf) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *LogsOutlierConditionAllOf) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *LogsOutlierConditionAllOf) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *LogsOutlierConditionAllOf) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetThreshold

`func (o *LogsOutlierConditionAllOf) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *LogsOutlierConditionAllOf) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *LogsOutlierConditionAllOf) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *LogsOutlierConditionAllOf) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetField

`func (o *LogsOutlierConditionAllOf) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *LogsOutlierConditionAllOf) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *LogsOutlierConditionAllOf) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *LogsOutlierConditionAllOf) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



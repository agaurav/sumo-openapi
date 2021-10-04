# AlertChartMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AbnormalityStartTime** | Pointer to **int64** | The time stamp at which abnomarlity started. | [optional] 
**AbnormalityEndTime** | Pointer to **int64** | The time stamp at which abnomarlity ended. | [optional] 
**EvaluationDelay** | Pointer to **int64** | The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time. | [optional] 
**AlertCreatedAt** | Pointer to **int64** | The time stamp at which the alert response page is created. | [optional] 
**AlertResolvedAt** | Pointer to **int64** | The time stamp at which the alert response page is resolved. | [optional] 

## Methods

### NewAlertChartMetadata

`func NewAlertChartMetadata() *AlertChartMetadata`

NewAlertChartMetadata instantiates a new AlertChartMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChartMetadataWithDefaults

`func NewAlertChartMetadataWithDefaults() *AlertChartMetadata`

NewAlertChartMetadataWithDefaults instantiates a new AlertChartMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbnormalityStartTime

`func (o *AlertChartMetadata) GetAbnormalityStartTime() int64`

GetAbnormalityStartTime returns the AbnormalityStartTime field if non-nil, zero value otherwise.

### GetAbnormalityStartTimeOk

`func (o *AlertChartMetadata) GetAbnormalityStartTimeOk() (*int64, bool)`

GetAbnormalityStartTimeOk returns a tuple with the AbnormalityStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalityStartTime

`func (o *AlertChartMetadata) SetAbnormalityStartTime(v int64)`

SetAbnormalityStartTime sets AbnormalityStartTime field to given value.

### HasAbnormalityStartTime

`func (o *AlertChartMetadata) HasAbnormalityStartTime() bool`

HasAbnormalityStartTime returns a boolean if a field has been set.

### GetAbnormalityEndTime

`func (o *AlertChartMetadata) GetAbnormalityEndTime() int64`

GetAbnormalityEndTime returns the AbnormalityEndTime field if non-nil, zero value otherwise.

### GetAbnormalityEndTimeOk

`func (o *AlertChartMetadata) GetAbnormalityEndTimeOk() (*int64, bool)`

GetAbnormalityEndTimeOk returns a tuple with the AbnormalityEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalityEndTime

`func (o *AlertChartMetadata) SetAbnormalityEndTime(v int64)`

SetAbnormalityEndTime sets AbnormalityEndTime field to given value.

### HasAbnormalityEndTime

`func (o *AlertChartMetadata) HasAbnormalityEndTime() bool`

HasAbnormalityEndTime returns a boolean if a field has been set.

### GetEvaluationDelay

`func (o *AlertChartMetadata) GetEvaluationDelay() int64`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *AlertChartMetadata) GetEvaluationDelayOk() (*int64, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *AlertChartMetadata) SetEvaluationDelay(v int64)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *AlertChartMetadata) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### GetAlertCreatedAt

`func (o *AlertChartMetadata) GetAlertCreatedAt() int64`

GetAlertCreatedAt returns the AlertCreatedAt field if non-nil, zero value otherwise.

### GetAlertCreatedAtOk

`func (o *AlertChartMetadata) GetAlertCreatedAtOk() (*int64, bool)`

GetAlertCreatedAtOk returns a tuple with the AlertCreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertCreatedAt

`func (o *AlertChartMetadata) SetAlertCreatedAt(v int64)`

SetAlertCreatedAt sets AlertCreatedAt field to given value.

### HasAlertCreatedAt

`func (o *AlertChartMetadata) HasAlertCreatedAt() bool`

HasAlertCreatedAt returns a boolean if a field has been set.

### GetAlertResolvedAt

`func (o *AlertChartMetadata) GetAlertResolvedAt() int64`

GetAlertResolvedAt returns the AlertResolvedAt field if non-nil, zero value otherwise.

### GetAlertResolvedAtOk

`func (o *AlertChartMetadata) GetAlertResolvedAtOk() (*int64, bool)`

GetAlertResolvedAtOk returns a tuple with the AlertResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertResolvedAt

`func (o *AlertChartMetadata) SetAlertResolvedAt(v int64)`

SetAlertResolvedAt sets AlertResolvedAt field to given value.

### HasAlertResolvedAt

`func (o *AlertChartMetadata) HasAlertResolvedAt() bool`

HasAlertResolvedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



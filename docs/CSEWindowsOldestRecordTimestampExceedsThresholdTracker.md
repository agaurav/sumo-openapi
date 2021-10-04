# CSEWindowsOldestRecordTimestampExceedsThresholdTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**SensorId** | Pointer to **string** | The sensor ID. | [optional] 
**SensorHostname** | Pointer to **string** | The sensor&#39;s hostname. | [optional] 
**Source** | Pointer to **string** | The HostName + EventLog name for EventLogs and Domain name for Directory. | [optional] 
**LastErrorMessage** | Pointer to **string** | The last error message. | [optional] 
**NumberOfFilesPending** | Pointer to **string** | The number of files pending upload. | [optional] 
**OldestTimestampInQueue** | Pointer to **string** | The oldest timestamp in the queue. | [optional] 

## Methods

### NewCSEWindowsOldestRecordTimestampExceedsThresholdTracker

`func NewCSEWindowsOldestRecordTimestampExceedsThresholdTracker() *CSEWindowsOldestRecordTimestampExceedsThresholdTracker`

NewCSEWindowsOldestRecordTimestampExceedsThresholdTracker instantiates a new CSEWindowsOldestRecordTimestampExceedsThresholdTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSEWindowsOldestRecordTimestampExceedsThresholdTrackerWithDefaults

`func NewCSEWindowsOldestRecordTimestampExceedsThresholdTrackerWithDefaults() *CSEWindowsOldestRecordTimestampExceedsThresholdTracker`

NewCSEWindowsOldestRecordTimestampExceedsThresholdTrackerWithDefaults instantiates a new CSEWindowsOldestRecordTimestampExceedsThresholdTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSensorId

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetSensorId() string`

GetSensorId returns the SensorId field if non-nil, zero value otherwise.

### GetSensorIdOk

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetSensorIdOk() (*string, bool)`

GetSensorIdOk returns a tuple with the SensorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorId

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) SetSensorId(v string)`

SetSensorId sets SensorId field to given value.

### HasSensorId

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) HasSensorId() bool`

HasSensorId returns a boolean if a field has been set.

### GetSensorHostname

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetSensorHostname() string`

GetSensorHostname returns the SensorHostname field if non-nil, zero value otherwise.

### GetSensorHostnameOk

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetSensorHostnameOk() (*string, bool)`

GetSensorHostnameOk returns a tuple with the SensorHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorHostname

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) SetSensorHostname(v string)`

SetSensorHostname sets SensorHostname field to given value.

### HasSensorHostname

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) HasSensorHostname() bool`

HasSensorHostname returns a boolean if a field has been set.

### GetSource

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetLastErrorMessage

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetLastErrorMessage() string`

GetLastErrorMessage returns the LastErrorMessage field if non-nil, zero value otherwise.

### GetLastErrorMessageOk

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetLastErrorMessageOk() (*string, bool)`

GetLastErrorMessageOk returns a tuple with the LastErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorMessage

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) SetLastErrorMessage(v string)`

SetLastErrorMessage sets LastErrorMessage field to given value.

### HasLastErrorMessage

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) HasLastErrorMessage() bool`

HasLastErrorMessage returns a boolean if a field has been set.

### GetNumberOfFilesPending

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetNumberOfFilesPending() string`

GetNumberOfFilesPending returns the NumberOfFilesPending field if non-nil, zero value otherwise.

### GetNumberOfFilesPendingOk

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetNumberOfFilesPendingOk() (*string, bool)`

GetNumberOfFilesPendingOk returns a tuple with the NumberOfFilesPending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfFilesPending

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) SetNumberOfFilesPending(v string)`

SetNumberOfFilesPending sets NumberOfFilesPending field to given value.

### HasNumberOfFilesPending

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) HasNumberOfFilesPending() bool`

HasNumberOfFilesPending returns a boolean if a field has been set.

### GetOldestTimestampInQueue

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetOldestTimestampInQueue() string`

GetOldestTimestampInQueue returns the OldestTimestampInQueue field if non-nil, zero value otherwise.

### GetOldestTimestampInQueueOk

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) GetOldestTimestampInQueueOk() (*string, bool)`

GetOldestTimestampInQueueOk returns a tuple with the OldestTimestampInQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldestTimestampInQueue

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) SetOldestTimestampInQueue(v string)`

SetOldestTimestampInQueue sets OldestTimestampInQueue field to given value.

### HasOldestTimestampInQueue

`func (o *CSEWindowsOldestRecordTimestampExceedsThresholdTracker) HasOldestTimestampInQueue() bool`

HasOldestTimestampInQueue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**SensorId** | Pointer to **string** | The sensor ID. | [optional] 
**SensorHostname** | Pointer to **string** | The sensor&#39;s hostname. | [optional] 
**SensorUserName** | Pointer to **string** | The sensor&#39;s user name. | [optional] 
**FolderPath** | Pointer to **string** | The path of the folder. | [optional] 
**FilePath** | Pointer to **string** | The complete file path. | [optional] 
**Source** | Pointer to **string** | The HostName + EventLog name for EventLogs and Domain name for Directory.. | [optional] 

## Methods

### NewCSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker

`func NewCSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker() *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker`

NewCSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker instantiates a new CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSEWindowsWriteQueueFilesToSensorDirectoryFailedTrackerWithDefaults

`func NewCSEWindowsWriteQueueFilesToSensorDirectoryFailedTrackerWithDefaults() *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker`

NewCSEWindowsWriteQueueFilesToSensorDirectoryFailedTrackerWithDefaults instantiates a new CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSensorId

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSensorId() string`

GetSensorId returns the SensorId field if non-nil, zero value otherwise.

### GetSensorIdOk

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSensorIdOk() (*string, bool)`

GetSensorIdOk returns a tuple with the SensorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorId

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) SetSensorId(v string)`

SetSensorId sets SensorId field to given value.

### HasSensorId

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) HasSensorId() bool`

HasSensorId returns a boolean if a field has been set.

### GetSensorHostname

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSensorHostname() string`

GetSensorHostname returns the SensorHostname field if non-nil, zero value otherwise.

### GetSensorHostnameOk

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSensorHostnameOk() (*string, bool)`

GetSensorHostnameOk returns a tuple with the SensorHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorHostname

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) SetSensorHostname(v string)`

SetSensorHostname sets SensorHostname field to given value.

### HasSensorHostname

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) HasSensorHostname() bool`

HasSensorHostname returns a boolean if a field has been set.

### GetSensorUserName

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSensorUserName() string`

GetSensorUserName returns the SensorUserName field if non-nil, zero value otherwise.

### GetSensorUserNameOk

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSensorUserNameOk() (*string, bool)`

GetSensorUserNameOk returns a tuple with the SensorUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorUserName

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) SetSensorUserName(v string)`

SetSensorUserName sets SensorUserName field to given value.

### HasSensorUserName

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) HasSensorUserName() bool`

HasSensorUserName returns a boolean if a field has been set.

### GetFolderPath

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### GetFilePath

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetSource

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *CSEWindowsWriteQueueFilesToSensorDirectoryFailedTracker) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



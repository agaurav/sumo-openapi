# CSEWindowsStorageLimitApproachingTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**SensorId** | Pointer to **string** | The sensor ID. | [optional] 
**SensorHostname** | Pointer to **string** | The sensor&#39;s hostname. | [optional] 
**FolderPath** | Pointer to **string** | The path of the folder. | [optional] 
**FolderSizeLimit** | Pointer to **string** | The complete file path. | [optional] 
**CurrentFolderSize** | Pointer to **string** | Current size of the folder. | [optional] 
**PercentageAvailableDiskSpaceLimit** | Pointer to **string** | The percentage available disk space limit. | [optional] 
**CurrentPercentageAvailableDiskSpace** | Pointer to **string** | The current percentage available disk space. | [optional] 
**LastError** | Pointer to **string** | The last error. | [optional] 

## Methods

### NewCSEWindowsStorageLimitApproachingTracker

`func NewCSEWindowsStorageLimitApproachingTracker() *CSEWindowsStorageLimitApproachingTracker`

NewCSEWindowsStorageLimitApproachingTracker instantiates a new CSEWindowsStorageLimitApproachingTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSEWindowsStorageLimitApproachingTrackerWithDefaults

`func NewCSEWindowsStorageLimitApproachingTrackerWithDefaults() *CSEWindowsStorageLimitApproachingTracker`

NewCSEWindowsStorageLimitApproachingTrackerWithDefaults instantiates a new CSEWindowsStorageLimitApproachingTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CSEWindowsStorageLimitApproachingTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CSEWindowsStorageLimitApproachingTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CSEWindowsStorageLimitApproachingTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetSensorId

`func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorId() string`

GetSensorId returns the SensorId field if non-nil, zero value otherwise.

### GetSensorIdOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorIdOk() (*string, bool)`

GetSensorIdOk returns a tuple with the SensorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorId

`func (o *CSEWindowsStorageLimitApproachingTracker) SetSensorId(v string)`

SetSensorId sets SensorId field to given value.

### HasSensorId

`func (o *CSEWindowsStorageLimitApproachingTracker) HasSensorId() bool`

HasSensorId returns a boolean if a field has been set.

### GetSensorHostname

`func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorHostname() string`

GetSensorHostname returns the SensorHostname field if non-nil, zero value otherwise.

### GetSensorHostnameOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetSensorHostnameOk() (*string, bool)`

GetSensorHostnameOk returns a tuple with the SensorHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensorHostname

`func (o *CSEWindowsStorageLimitApproachingTracker) SetSensorHostname(v string)`

SetSensorHostname sets SensorHostname field to given value.

### HasSensorHostname

`func (o *CSEWindowsStorageLimitApproachingTracker) HasSensorHostname() bool`

HasSensorHostname returns a boolean if a field has been set.

### GetFolderPath

`func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderPath() string`

GetFolderPath returns the FolderPath field if non-nil, zero value otherwise.

### GetFolderPathOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderPathOk() (*string, bool)`

GetFolderPathOk returns a tuple with the FolderPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderPath

`func (o *CSEWindowsStorageLimitApproachingTracker) SetFolderPath(v string)`

SetFolderPath sets FolderPath field to given value.

### HasFolderPath

`func (o *CSEWindowsStorageLimitApproachingTracker) HasFolderPath() bool`

HasFolderPath returns a boolean if a field has been set.

### GetFolderSizeLimit

`func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderSizeLimit() string`

GetFolderSizeLimit returns the FolderSizeLimit field if non-nil, zero value otherwise.

### GetFolderSizeLimitOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetFolderSizeLimitOk() (*string, bool)`

GetFolderSizeLimitOk returns a tuple with the FolderSizeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderSizeLimit

`func (o *CSEWindowsStorageLimitApproachingTracker) SetFolderSizeLimit(v string)`

SetFolderSizeLimit sets FolderSizeLimit field to given value.

### HasFolderSizeLimit

`func (o *CSEWindowsStorageLimitApproachingTracker) HasFolderSizeLimit() bool`

HasFolderSizeLimit returns a boolean if a field has been set.

### GetCurrentFolderSize

`func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentFolderSize() string`

GetCurrentFolderSize returns the CurrentFolderSize field if non-nil, zero value otherwise.

### GetCurrentFolderSizeOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentFolderSizeOk() (*string, bool)`

GetCurrentFolderSizeOk returns a tuple with the CurrentFolderSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentFolderSize

`func (o *CSEWindowsStorageLimitApproachingTracker) SetCurrentFolderSize(v string)`

SetCurrentFolderSize sets CurrentFolderSize field to given value.

### HasCurrentFolderSize

`func (o *CSEWindowsStorageLimitApproachingTracker) HasCurrentFolderSize() bool`

HasCurrentFolderSize returns a boolean if a field has been set.

### GetPercentageAvailableDiskSpaceLimit

`func (o *CSEWindowsStorageLimitApproachingTracker) GetPercentageAvailableDiskSpaceLimit() string`

GetPercentageAvailableDiskSpaceLimit returns the PercentageAvailableDiskSpaceLimit field if non-nil, zero value otherwise.

### GetPercentageAvailableDiskSpaceLimitOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetPercentageAvailableDiskSpaceLimitOk() (*string, bool)`

GetPercentageAvailableDiskSpaceLimitOk returns a tuple with the PercentageAvailableDiskSpaceLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageAvailableDiskSpaceLimit

`func (o *CSEWindowsStorageLimitApproachingTracker) SetPercentageAvailableDiskSpaceLimit(v string)`

SetPercentageAvailableDiskSpaceLimit sets PercentageAvailableDiskSpaceLimit field to given value.

### HasPercentageAvailableDiskSpaceLimit

`func (o *CSEWindowsStorageLimitApproachingTracker) HasPercentageAvailableDiskSpaceLimit() bool`

HasPercentageAvailableDiskSpaceLimit returns a boolean if a field has been set.

### GetCurrentPercentageAvailableDiskSpace

`func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentPercentageAvailableDiskSpace() string`

GetCurrentPercentageAvailableDiskSpace returns the CurrentPercentageAvailableDiskSpace field if non-nil, zero value otherwise.

### GetCurrentPercentageAvailableDiskSpaceOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetCurrentPercentageAvailableDiskSpaceOk() (*string, bool)`

GetCurrentPercentageAvailableDiskSpaceOk returns a tuple with the CurrentPercentageAvailableDiskSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPercentageAvailableDiskSpace

`func (o *CSEWindowsStorageLimitApproachingTracker) SetCurrentPercentageAvailableDiskSpace(v string)`

SetCurrentPercentageAvailableDiskSpace sets CurrentPercentageAvailableDiskSpace field to given value.

### HasCurrentPercentageAvailableDiskSpace

`func (o *CSEWindowsStorageLimitApproachingTracker) HasCurrentPercentageAvailableDiskSpace() bool`

HasCurrentPercentageAvailableDiskSpace returns a boolean if a field has been set.

### GetLastError

`func (o *CSEWindowsStorageLimitApproachingTracker) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *CSEWindowsStorageLimitApproachingTracker) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *CSEWindowsStorageLimitApproachingTracker) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *CSEWindowsStorageLimitApproachingTracker) HasLastError() bool`

HasLastError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



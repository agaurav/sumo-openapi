# ArchiveJobsCount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceId** | **string** | Identifier for the archive source. | 
**Pending** | **int64** | The total number of archive jobs with pending status for the archive source. | 
**Scanning** | **int64** | The total number of archive jobs with scanning status for the archive source. | 
**Ingesting** | **int64** | The total number of archive jobs with ingesting status for the archive source. | 
**Failed** | **int64** | The total number of archive jobs with failed status for the archive source. | 
**Succeeded** | **int64** | The total number of archive jobs with succeeded status for the archive source. | 

## Methods

### NewArchiveJobsCount

`func NewArchiveJobsCount(sourceId string, pending int64, scanning int64, ingesting int64, failed int64, succeeded int64, ) *ArchiveJobsCount`

NewArchiveJobsCount instantiates a new ArchiveJobsCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveJobsCountWithDefaults

`func NewArchiveJobsCountWithDefaults() *ArchiveJobsCount`

NewArchiveJobsCountWithDefaults instantiates a new ArchiveJobsCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceId

`func (o *ArchiveJobsCount) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ArchiveJobsCount) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ArchiveJobsCount) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetPending

`func (o *ArchiveJobsCount) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *ArchiveJobsCount) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *ArchiveJobsCount) SetPending(v int64)`

SetPending sets Pending field to given value.


### GetScanning

`func (o *ArchiveJobsCount) GetScanning() int64`

GetScanning returns the Scanning field if non-nil, zero value otherwise.

### GetScanningOk

`func (o *ArchiveJobsCount) GetScanningOk() (*int64, bool)`

GetScanningOk returns a tuple with the Scanning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanning

`func (o *ArchiveJobsCount) SetScanning(v int64)`

SetScanning sets Scanning field to given value.


### GetIngesting

`func (o *ArchiveJobsCount) GetIngesting() int64`

GetIngesting returns the Ingesting field if non-nil, zero value otherwise.

### GetIngestingOk

`func (o *ArchiveJobsCount) GetIngestingOk() (*int64, bool)`

GetIngestingOk returns a tuple with the Ingesting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngesting

`func (o *ArchiveJobsCount) SetIngesting(v int64)`

SetIngesting sets Ingesting field to given value.


### GetFailed

`func (o *ArchiveJobsCount) GetFailed() int64`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *ArchiveJobsCount) GetFailedOk() (*int64, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *ArchiveJobsCount) SetFailed(v int64)`

SetFailed sets Failed field to given value.


### GetSucceeded

`func (o *ArchiveJobsCount) GetSucceeded() int64`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *ArchiveJobsCount) GetSucceededOk() (*int64, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *ArchiveJobsCount) SetSucceeded(v int64)`

SetSucceeded sets Succeeded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



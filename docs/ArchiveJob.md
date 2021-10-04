# ArchiveJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the ingestion job. | 
**StartTime** | **time.Time** | The starting timestamp of the ingestion job. | 
**EndTime** | **time.Time** | The ending timestamp of the ingestion job. | 
**Id** | **string** | The unique identifier of the ingestion job. | 
**TotalObjectsScanned** | **int64** | The total number of objects scanned by the ingestion job. | 
**TotalObjectsIngested** | **int64** | The total number of objects ingested by the ingestion job. | 
**TotalBytesIngested** | **int64** | The total bytes ingested by the ingestion job. | 
**Status** | **string** | The status of the ingestion job, either &#x60;Pending&#x60;,&#x60;Scanning&#x60;,&#x60;Ingesting&#x60;,&#x60;Failed&#x60;, or &#x60;Succeeded&#x60;. | 
**CreatedAt** | **time.Time** | The creation timestamp in UTC of the ingestion job. | 
**CreatedBy** | **string** | The identifier of the user who created the ingestion job. | 

## Methods

### NewArchiveJob

`func NewArchiveJob(name string, startTime time.Time, endTime time.Time, id string, totalObjectsScanned int64, totalObjectsIngested int64, totalBytesIngested int64, status string, createdAt time.Time, createdBy string, ) *ArchiveJob`

NewArchiveJob instantiates a new ArchiveJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveJobWithDefaults

`func NewArchiveJobWithDefaults() *ArchiveJob`

NewArchiveJobWithDefaults instantiates a new ArchiveJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArchiveJob) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArchiveJob) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArchiveJob) SetName(v string)`

SetName sets Name field to given value.


### GetStartTime

`func (o *ArchiveJob) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ArchiveJob) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ArchiveJob) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *ArchiveJob) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ArchiveJob) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ArchiveJob) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetId

`func (o *ArchiveJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveJob) SetId(v string)`

SetId sets Id field to given value.


### GetTotalObjectsScanned

`func (o *ArchiveJob) GetTotalObjectsScanned() int64`

GetTotalObjectsScanned returns the TotalObjectsScanned field if non-nil, zero value otherwise.

### GetTotalObjectsScannedOk

`func (o *ArchiveJob) GetTotalObjectsScannedOk() (*int64, bool)`

GetTotalObjectsScannedOk returns a tuple with the TotalObjectsScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjectsScanned

`func (o *ArchiveJob) SetTotalObjectsScanned(v int64)`

SetTotalObjectsScanned sets TotalObjectsScanned field to given value.


### GetTotalObjectsIngested

`func (o *ArchiveJob) GetTotalObjectsIngested() int64`

GetTotalObjectsIngested returns the TotalObjectsIngested field if non-nil, zero value otherwise.

### GetTotalObjectsIngestedOk

`func (o *ArchiveJob) GetTotalObjectsIngestedOk() (*int64, bool)`

GetTotalObjectsIngestedOk returns a tuple with the TotalObjectsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjectsIngested

`func (o *ArchiveJob) SetTotalObjectsIngested(v int64)`

SetTotalObjectsIngested sets TotalObjectsIngested field to given value.


### GetTotalBytesIngested

`func (o *ArchiveJob) GetTotalBytesIngested() int64`

GetTotalBytesIngested returns the TotalBytesIngested field if non-nil, zero value otherwise.

### GetTotalBytesIngestedOk

`func (o *ArchiveJob) GetTotalBytesIngestedOk() (*int64, bool)`

GetTotalBytesIngestedOk returns a tuple with the TotalBytesIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesIngested

`func (o *ArchiveJob) SetTotalBytesIngested(v int64)`

SetTotalBytesIngested sets TotalBytesIngested field to given value.


### GetStatus

`func (o *ArchiveJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchiveJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchiveJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ArchiveJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArchiveJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArchiveJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ArchiveJob) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArchiveJob) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArchiveJob) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



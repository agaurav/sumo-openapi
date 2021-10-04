# ArchiveJobAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the ingestion job. | 
**TotalObjectsScanned** | **int64** | The total number of objects scanned by the ingestion job. | 
**TotalObjectsIngested** | **int64** | The total number of objects ingested by the ingestion job. | 
**TotalBytesIngested** | **int64** | The total bytes ingested by the ingestion job. | 
**Status** | **string** | The status of the ingestion job, either &#x60;Pending&#x60;,&#x60;Scanning&#x60;,&#x60;Ingesting&#x60;,&#x60;Failed&#x60;, or &#x60;Succeeded&#x60;. | 
**CreatedAt** | **time.Time** | The creation timestamp in UTC of the ingestion job. | 
**CreatedBy** | **string** | The identifier of the user who created the ingestion job. | 

## Methods

### NewArchiveJobAllOf

`func NewArchiveJobAllOf(id string, totalObjectsScanned int64, totalObjectsIngested int64, totalBytesIngested int64, status string, createdAt time.Time, createdBy string, ) *ArchiveJobAllOf`

NewArchiveJobAllOf instantiates a new ArchiveJobAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveJobAllOfWithDefaults

`func NewArchiveJobAllOfWithDefaults() *ArchiveJobAllOf`

NewArchiveJobAllOfWithDefaults instantiates a new ArchiveJobAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveJobAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveJobAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveJobAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetTotalObjectsScanned

`func (o *ArchiveJobAllOf) GetTotalObjectsScanned() int64`

GetTotalObjectsScanned returns the TotalObjectsScanned field if non-nil, zero value otherwise.

### GetTotalObjectsScannedOk

`func (o *ArchiveJobAllOf) GetTotalObjectsScannedOk() (*int64, bool)`

GetTotalObjectsScannedOk returns a tuple with the TotalObjectsScanned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjectsScanned

`func (o *ArchiveJobAllOf) SetTotalObjectsScanned(v int64)`

SetTotalObjectsScanned sets TotalObjectsScanned field to given value.


### GetTotalObjectsIngested

`func (o *ArchiveJobAllOf) GetTotalObjectsIngested() int64`

GetTotalObjectsIngested returns the TotalObjectsIngested field if non-nil, zero value otherwise.

### GetTotalObjectsIngestedOk

`func (o *ArchiveJobAllOf) GetTotalObjectsIngestedOk() (*int64, bool)`

GetTotalObjectsIngestedOk returns a tuple with the TotalObjectsIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalObjectsIngested

`func (o *ArchiveJobAllOf) SetTotalObjectsIngested(v int64)`

SetTotalObjectsIngested sets TotalObjectsIngested field to given value.


### GetTotalBytesIngested

`func (o *ArchiveJobAllOf) GetTotalBytesIngested() int64`

GetTotalBytesIngested returns the TotalBytesIngested field if non-nil, zero value otherwise.

### GetTotalBytesIngestedOk

`func (o *ArchiveJobAllOf) GetTotalBytesIngestedOk() (*int64, bool)`

GetTotalBytesIngestedOk returns a tuple with the TotalBytesIngested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytesIngested

`func (o *ArchiveJobAllOf) SetTotalBytesIngested(v int64)`

SetTotalBytesIngested sets TotalBytesIngested field to given value.


### GetStatus

`func (o *ArchiveJobAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchiveJobAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchiveJobAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *ArchiveJobAllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ArchiveJobAllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ArchiveJobAllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ArchiveJobAllOf) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArchiveJobAllOf) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArchiveJobAllOf) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# IngestBudgetV2AllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the ingest budget. | 
**UsageBytes** | Pointer to **int64** | Current usage since the last reset, in bytes. | [optional] 
**UsageStatus** | Pointer to **string** | Status of the current usage. Can be &#x60;Normal&#x60;, &#x60;Approaching&#x60;, &#x60;Exceeded&#x60;, or &#x60;Unknown&#x60; (unable to retrieve usage). | [optional] 
**CreatedAt** | **time.Time** | The creation timestamp in UTC of the Ingest Budget. | 
**CreatedBy** | **string** | The identifier of the user who created the Ingest Budget. | 
**ModifiedAt** | **time.Time** | The modified timestamp in UTC of the Ingest Budget. | 
**ModifiedBy** | **string** | The identifier of the user who modified the Ingest Budget. | 
**BudgetVersion** | Pointer to **int32** | The version of the Ingest Budget | [optional] 

## Methods

### NewIngestBudgetV2AllOf

`func NewIngestBudgetV2AllOf(id string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, ) *IngestBudgetV2AllOf`

NewIngestBudgetV2AllOf instantiates a new IngestBudgetV2AllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestBudgetV2AllOfWithDefaults

`func NewIngestBudgetV2AllOfWithDefaults() *IngestBudgetV2AllOf`

NewIngestBudgetV2AllOfWithDefaults instantiates a new IngestBudgetV2AllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestBudgetV2AllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestBudgetV2AllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestBudgetV2AllOf) SetId(v string)`

SetId sets Id field to given value.


### GetUsageBytes

`func (o *IngestBudgetV2AllOf) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *IngestBudgetV2AllOf) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *IngestBudgetV2AllOf) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *IngestBudgetV2AllOf) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### GetUsageStatus

`func (o *IngestBudgetV2AllOf) GetUsageStatus() string`

GetUsageStatus returns the UsageStatus field if non-nil, zero value otherwise.

### GetUsageStatusOk

`func (o *IngestBudgetV2AllOf) GetUsageStatusOk() (*string, bool)`

GetUsageStatusOk returns a tuple with the UsageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStatus

`func (o *IngestBudgetV2AllOf) SetUsageStatus(v string)`

SetUsageStatus sets UsageStatus field to given value.

### HasUsageStatus

`func (o *IngestBudgetV2AllOf) HasUsageStatus() bool`

HasUsageStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IngestBudgetV2AllOf) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngestBudgetV2AllOf) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngestBudgetV2AllOf) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *IngestBudgetV2AllOf) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IngestBudgetV2AllOf) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IngestBudgetV2AllOf) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *IngestBudgetV2AllOf) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *IngestBudgetV2AllOf) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *IngestBudgetV2AllOf) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *IngestBudgetV2AllOf) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *IngestBudgetV2AllOf) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *IngestBudgetV2AllOf) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetBudgetVersion

`func (o *IngestBudgetV2AllOf) GetBudgetVersion() int32`

GetBudgetVersion returns the BudgetVersion field if non-nil, zero value otherwise.

### GetBudgetVersionOk

`func (o *IngestBudgetV2AllOf) GetBudgetVersionOk() (*int32, bool)`

GetBudgetVersionOk returns a tuple with the BudgetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetVersion

`func (o *IngestBudgetV2AllOf) SetBudgetVersion(v int32)`

SetBudgetVersion sets BudgetVersion field to given value.

### HasBudgetVersion

`func (o *IngestBudgetV2AllOf) HasBudgetVersion() bool`

HasBudgetVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



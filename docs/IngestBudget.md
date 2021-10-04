# IngestBudget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the ingest budget. | 
**FieldValue** | **string** | Custom field value that is used to assign Collectors to the ingest budget. | 
**CapacityBytes** | **int64** | Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit. | 
**Timezone** | **string** | Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 
**ResetTime** | **string** | Reset time of the ingest budget in HH:MM format. | 
**Description** | Pointer to **string** | Description of the ingest budget. | [optional] 
**Action** | **string** | Action to take when ingest budget&#39;s capacity is reached. All actions are audited. Supported values are:   * &#x60;stopCollecting&#x60;   * &#x60;keepCollecting&#x60; | 
**AuditThreshold** | Pointer to **int32** | The threshold as a percentage of when an ingest budget&#39;s capacity usage is logged in the Audit Index. | [optional] 
**CreatedAt** | **NullableTime** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedByUser** | [**UserInfo**](UserInfo.md) |  | 
**ModifiedAt** | **NullableTime** | Last modification timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**ModifiedByUser** | [**UserInfo**](UserInfo.md) |  | 
**Id** | **string** | Unique identifier for the ingest budget. | 
**UsageBytes** | Pointer to **int64** | Current usage since the last reset, in bytes. | [optional] 
**UsageStatus** | Pointer to **string** | Status of the current usage. Can be &#x60;Normal&#x60;, &#x60;Approaching&#x60;, &#x60;Exceeded&#x60;, or &#x60;Unknown&#x60; (unable to retrieve usage). | [optional] 
**NumberOfCollectors** | Pointer to **int64** | Number of collectors assigned to the ingest budget. | [optional] 

## Methods

### NewIngestBudget

`func NewIngestBudget(name string, fieldValue string, capacityBytes int64, timezone string, resetTime string, action string, createdAt NullableTime, createdByUser UserInfo, modifiedAt NullableTime, modifiedByUser UserInfo, id string, ) *IngestBudget`

NewIngestBudget instantiates a new IngestBudget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestBudgetWithDefaults

`func NewIngestBudgetWithDefaults() *IngestBudget`

NewIngestBudgetWithDefaults instantiates a new IngestBudget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IngestBudget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestBudget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestBudget) SetName(v string)`

SetName sets Name field to given value.


### GetFieldValue

`func (o *IngestBudget) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *IngestBudget) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *IngestBudget) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.


### GetCapacityBytes

`func (o *IngestBudget) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *IngestBudget) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *IngestBudget) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.


### GetTimezone

`func (o *IngestBudget) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IngestBudget) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IngestBudget) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetResetTime

`func (o *IngestBudget) GetResetTime() string`

GetResetTime returns the ResetTime field if non-nil, zero value otherwise.

### GetResetTimeOk

`func (o *IngestBudget) GetResetTimeOk() (*string, bool)`

GetResetTimeOk returns a tuple with the ResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTime

`func (o *IngestBudget) SetResetTime(v string)`

SetResetTime sets ResetTime field to given value.


### GetDescription

`func (o *IngestBudget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestBudget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestBudget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestBudget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *IngestBudget) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IngestBudget) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IngestBudget) SetAction(v string)`

SetAction sets Action field to given value.


### GetAuditThreshold

`func (o *IngestBudget) GetAuditThreshold() int32`

GetAuditThreshold returns the AuditThreshold field if non-nil, zero value otherwise.

### GetAuditThresholdOk

`func (o *IngestBudget) GetAuditThresholdOk() (*int32, bool)`

GetAuditThresholdOk returns a tuple with the AuditThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditThreshold

`func (o *IngestBudget) SetAuditThreshold(v int32)`

SetAuditThreshold sets AuditThreshold field to given value.

### HasAuditThreshold

`func (o *IngestBudget) HasAuditThreshold() bool`

HasAuditThreshold returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IngestBudget) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngestBudget) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngestBudget) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### SetCreatedAtNil

`func (o *IngestBudget) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *IngestBudget) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetCreatedByUser

`func (o *IngestBudget) GetCreatedByUser() UserInfo`

GetCreatedByUser returns the CreatedByUser field if non-nil, zero value otherwise.

### GetCreatedByUserOk

`func (o *IngestBudget) GetCreatedByUserOk() (*UserInfo, bool)`

GetCreatedByUserOk returns a tuple with the CreatedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUser

`func (o *IngestBudget) SetCreatedByUser(v UserInfo)`

SetCreatedByUser sets CreatedByUser field to given value.


### GetModifiedAt

`func (o *IngestBudget) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *IngestBudget) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *IngestBudget) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### SetModifiedAtNil

`func (o *IngestBudget) SetModifiedAtNil(b bool)`

 SetModifiedAtNil sets the value for ModifiedAt to be an explicit nil

### UnsetModifiedAt
`func (o *IngestBudget) UnsetModifiedAt()`

UnsetModifiedAt ensures that no value is present for ModifiedAt, not even an explicit nil
### GetModifiedByUser

`func (o *IngestBudget) GetModifiedByUser() UserInfo`

GetModifiedByUser returns the ModifiedByUser field if non-nil, zero value otherwise.

### GetModifiedByUserOk

`func (o *IngestBudget) GetModifiedByUserOk() (*UserInfo, bool)`

GetModifiedByUserOk returns a tuple with the ModifiedByUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedByUser

`func (o *IngestBudget) SetModifiedByUser(v UserInfo)`

SetModifiedByUser sets ModifiedByUser field to given value.


### GetId

`func (o *IngestBudget) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestBudget) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestBudget) SetId(v string)`

SetId sets Id field to given value.


### GetUsageBytes

`func (o *IngestBudget) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *IngestBudget) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *IngestBudget) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *IngestBudget) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### GetUsageStatus

`func (o *IngestBudget) GetUsageStatus() string`

GetUsageStatus returns the UsageStatus field if non-nil, zero value otherwise.

### GetUsageStatusOk

`func (o *IngestBudget) GetUsageStatusOk() (*string, bool)`

GetUsageStatusOk returns a tuple with the UsageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStatus

`func (o *IngestBudget) SetUsageStatus(v string)`

SetUsageStatus sets UsageStatus field to given value.

### HasUsageStatus

`func (o *IngestBudget) HasUsageStatus() bool`

HasUsageStatus returns a boolean if a field has been set.

### GetNumberOfCollectors

`func (o *IngestBudget) GetNumberOfCollectors() int64`

GetNumberOfCollectors returns the NumberOfCollectors field if non-nil, zero value otherwise.

### GetNumberOfCollectorsOk

`func (o *IngestBudget) GetNumberOfCollectorsOk() (*int64, bool)`

GetNumberOfCollectorsOk returns a tuple with the NumberOfCollectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCollectors

`func (o *IngestBudget) SetNumberOfCollectors(v int64)`

SetNumberOfCollectors sets NumberOfCollectors field to given value.

### HasNumberOfCollectors

`func (o *IngestBudget) HasNumberOfCollectors() bool`

HasNumberOfCollectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



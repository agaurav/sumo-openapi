# IngestBudgetV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the ingest budget. | 
**Scope** | **string** | A scope is a constraint that will be used to identify the messages on which budget needs to be applied. A scope is consists of key and value separated by &#x3D;. The field must be enabled in the fields table. Value supports wildcard. e.g. _sourceCategory&#x3D;*prod*payment*, cluster&#x3D;kafka. If the scope is defined _sourceCategory&#x3D;*nginx* in this budget will be applied on messages having fields _sourceCategory&#x3D;prod/nginx, _sourceCategory&#x3D;dev/nginx, or _sourceCategory&#x3D;dev/nginx/error | 
**CapacityBytes** | **int64** | Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit. | 
**Timezone** | **string** | Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List). | 
**ResetTime** | **string** | Reset time of the ingest budget in HH:MM format. | 
**Description** | Pointer to **string** | Description of the ingest budget. | [optional] 
**Action** | **string** | Action to take when ingest budget&#39;s capacity is reached. All actions are audited. Supported values are:   * &#x60;stopCollecting&#x60;   * &#x60;keepCollecting&#x60; | 
**AuditThreshold** | Pointer to **int32** | The threshold as a percentage of when an ingest budget&#39;s capacity usage is logged in the Audit Index. | [optional] 
**Id** | **string** | Unique identifier for the ingest budget. | 
**UsageBytes** | Pointer to **int64** | Current usage since the last reset, in bytes. | [optional] 
**UsageStatus** | Pointer to **string** | Status of the current usage. Can be &#x60;Normal&#x60;, &#x60;Approaching&#x60;, &#x60;Exceeded&#x60;, or &#x60;Unknown&#x60; (unable to retrieve usage). | [optional] 
**CreatedAt** | **time.Time** | The creation timestamp in UTC of the Ingest Budget. | 
**CreatedBy** | **string** | The identifier of the user who created the Ingest Budget. | 
**ModifiedAt** | **time.Time** | The modified timestamp in UTC of the Ingest Budget. | 
**ModifiedBy** | **string** | The identifier of the user who modified the Ingest Budget. | 
**BudgetVersion** | Pointer to **int32** | The version of the Ingest Budget | [optional] 

## Methods

### NewIngestBudgetV2

`func NewIngestBudgetV2(name string, scope string, capacityBytes int64, timezone string, resetTime string, action string, id string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, ) *IngestBudgetV2`

NewIngestBudgetV2 instantiates a new IngestBudgetV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestBudgetV2WithDefaults

`func NewIngestBudgetV2WithDefaults() *IngestBudgetV2`

NewIngestBudgetV2WithDefaults instantiates a new IngestBudgetV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IngestBudgetV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestBudgetV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestBudgetV2) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *IngestBudgetV2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IngestBudgetV2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IngestBudgetV2) SetScope(v string)`

SetScope sets Scope field to given value.


### GetCapacityBytes

`func (o *IngestBudgetV2) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *IngestBudgetV2) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *IngestBudgetV2) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.


### GetTimezone

`func (o *IngestBudgetV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IngestBudgetV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IngestBudgetV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetResetTime

`func (o *IngestBudgetV2) GetResetTime() string`

GetResetTime returns the ResetTime field if non-nil, zero value otherwise.

### GetResetTimeOk

`func (o *IngestBudgetV2) GetResetTimeOk() (*string, bool)`

GetResetTimeOk returns a tuple with the ResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTime

`func (o *IngestBudgetV2) SetResetTime(v string)`

SetResetTime sets ResetTime field to given value.


### GetDescription

`func (o *IngestBudgetV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestBudgetV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestBudgetV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestBudgetV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *IngestBudgetV2) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IngestBudgetV2) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IngestBudgetV2) SetAction(v string)`

SetAction sets Action field to given value.


### GetAuditThreshold

`func (o *IngestBudgetV2) GetAuditThreshold() int32`

GetAuditThreshold returns the AuditThreshold field if non-nil, zero value otherwise.

### GetAuditThresholdOk

`func (o *IngestBudgetV2) GetAuditThresholdOk() (*int32, bool)`

GetAuditThresholdOk returns a tuple with the AuditThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditThreshold

`func (o *IngestBudgetV2) SetAuditThreshold(v int32)`

SetAuditThreshold sets AuditThreshold field to given value.

### HasAuditThreshold

`func (o *IngestBudgetV2) HasAuditThreshold() bool`

HasAuditThreshold returns a boolean if a field has been set.

### GetId

`func (o *IngestBudgetV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestBudgetV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestBudgetV2) SetId(v string)`

SetId sets Id field to given value.


### GetUsageBytes

`func (o *IngestBudgetV2) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *IngestBudgetV2) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *IngestBudgetV2) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *IngestBudgetV2) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### GetUsageStatus

`func (o *IngestBudgetV2) GetUsageStatus() string`

GetUsageStatus returns the UsageStatus field if non-nil, zero value otherwise.

### GetUsageStatusOk

`func (o *IngestBudgetV2) GetUsageStatusOk() (*string, bool)`

GetUsageStatusOk returns a tuple with the UsageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStatus

`func (o *IngestBudgetV2) SetUsageStatus(v string)`

SetUsageStatus sets UsageStatus field to given value.

### HasUsageStatus

`func (o *IngestBudgetV2) HasUsageStatus() bool`

HasUsageStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IngestBudgetV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IngestBudgetV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IngestBudgetV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *IngestBudgetV2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *IngestBudgetV2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *IngestBudgetV2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *IngestBudgetV2) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *IngestBudgetV2) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *IngestBudgetV2) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *IngestBudgetV2) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *IngestBudgetV2) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *IngestBudgetV2) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetBudgetVersion

`func (o *IngestBudgetV2) GetBudgetVersion() int32`

GetBudgetVersion returns the BudgetVersion field if non-nil, zero value otherwise.

### GetBudgetVersionOk

`func (o *IngestBudgetV2) GetBudgetVersionOk() (*int32, bool)`

GetBudgetVersionOk returns a tuple with the BudgetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetVersion

`func (o *IngestBudgetV2) SetBudgetVersion(v int32)`

SetBudgetVersion sets BudgetVersion field to given value.

### HasBudgetVersion

`func (o *IngestBudgetV2) HasBudgetVersion() bool`

HasBudgetVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



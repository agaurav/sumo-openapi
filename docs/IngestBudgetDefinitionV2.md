# IngestBudgetDefinitionV2

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

## Methods

### NewIngestBudgetDefinitionV2

`func NewIngestBudgetDefinitionV2(name string, scope string, capacityBytes int64, timezone string, resetTime string, action string, ) *IngestBudgetDefinitionV2`

NewIngestBudgetDefinitionV2 instantiates a new IngestBudgetDefinitionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestBudgetDefinitionV2WithDefaults

`func NewIngestBudgetDefinitionV2WithDefaults() *IngestBudgetDefinitionV2`

NewIngestBudgetDefinitionV2WithDefaults instantiates a new IngestBudgetDefinitionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IngestBudgetDefinitionV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestBudgetDefinitionV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestBudgetDefinitionV2) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *IngestBudgetDefinitionV2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IngestBudgetDefinitionV2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IngestBudgetDefinitionV2) SetScope(v string)`

SetScope sets Scope field to given value.


### GetCapacityBytes

`func (o *IngestBudgetDefinitionV2) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *IngestBudgetDefinitionV2) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *IngestBudgetDefinitionV2) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.


### GetTimezone

`func (o *IngestBudgetDefinitionV2) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IngestBudgetDefinitionV2) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IngestBudgetDefinitionV2) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetResetTime

`func (o *IngestBudgetDefinitionV2) GetResetTime() string`

GetResetTime returns the ResetTime field if non-nil, zero value otherwise.

### GetResetTimeOk

`func (o *IngestBudgetDefinitionV2) GetResetTimeOk() (*string, bool)`

GetResetTimeOk returns a tuple with the ResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTime

`func (o *IngestBudgetDefinitionV2) SetResetTime(v string)`

SetResetTime sets ResetTime field to given value.


### GetDescription

`func (o *IngestBudgetDefinitionV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestBudgetDefinitionV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestBudgetDefinitionV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestBudgetDefinitionV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *IngestBudgetDefinitionV2) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IngestBudgetDefinitionV2) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IngestBudgetDefinitionV2) SetAction(v string)`

SetAction sets Action field to given value.


### GetAuditThreshold

`func (o *IngestBudgetDefinitionV2) GetAuditThreshold() int32`

GetAuditThreshold returns the AuditThreshold field if non-nil, zero value otherwise.

### GetAuditThresholdOk

`func (o *IngestBudgetDefinitionV2) GetAuditThresholdOk() (*int32, bool)`

GetAuditThresholdOk returns a tuple with the AuditThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditThreshold

`func (o *IngestBudgetDefinitionV2) SetAuditThreshold(v int32)`

SetAuditThreshold sets AuditThreshold field to given value.

### HasAuditThreshold

`func (o *IngestBudgetDefinitionV2) HasAuditThreshold() bool`

HasAuditThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



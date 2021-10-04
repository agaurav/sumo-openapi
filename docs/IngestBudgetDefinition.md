# IngestBudgetDefinition

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

## Methods

### NewIngestBudgetDefinition

`func NewIngestBudgetDefinition(name string, fieldValue string, capacityBytes int64, timezone string, resetTime string, action string, ) *IngestBudgetDefinition`

NewIngestBudgetDefinition instantiates a new IngestBudgetDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestBudgetDefinitionWithDefaults

`func NewIngestBudgetDefinitionWithDefaults() *IngestBudgetDefinition`

NewIngestBudgetDefinitionWithDefaults instantiates a new IngestBudgetDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IngestBudgetDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IngestBudgetDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IngestBudgetDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetFieldValue

`func (o *IngestBudgetDefinition) GetFieldValue() string`

GetFieldValue returns the FieldValue field if non-nil, zero value otherwise.

### GetFieldValueOk

`func (o *IngestBudgetDefinition) GetFieldValueOk() (*string, bool)`

GetFieldValueOk returns a tuple with the FieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValue

`func (o *IngestBudgetDefinition) SetFieldValue(v string)`

SetFieldValue sets FieldValue field to given value.


### GetCapacityBytes

`func (o *IngestBudgetDefinition) GetCapacityBytes() int64`

GetCapacityBytes returns the CapacityBytes field if non-nil, zero value otherwise.

### GetCapacityBytesOk

`func (o *IngestBudgetDefinition) GetCapacityBytesOk() (*int64, bool)`

GetCapacityBytesOk returns a tuple with the CapacityBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacityBytes

`func (o *IngestBudgetDefinition) SetCapacityBytes(v int64)`

SetCapacityBytes sets CapacityBytes field to given value.


### GetTimezone

`func (o *IngestBudgetDefinition) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *IngestBudgetDefinition) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *IngestBudgetDefinition) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetResetTime

`func (o *IngestBudgetDefinition) GetResetTime() string`

GetResetTime returns the ResetTime field if non-nil, zero value otherwise.

### GetResetTimeOk

`func (o *IngestBudgetDefinition) GetResetTimeOk() (*string, bool)`

GetResetTimeOk returns a tuple with the ResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetTime

`func (o *IngestBudgetDefinition) SetResetTime(v string)`

SetResetTime sets ResetTime field to given value.


### GetDescription

`func (o *IngestBudgetDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IngestBudgetDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IngestBudgetDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IngestBudgetDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAction

`func (o *IngestBudgetDefinition) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *IngestBudgetDefinition) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *IngestBudgetDefinition) SetAction(v string)`

SetAction sets Action field to given value.


### GetAuditThreshold

`func (o *IngestBudgetDefinition) GetAuditThreshold() int32`

GetAuditThreshold returns the AuditThreshold field if non-nil, zero value otherwise.

### GetAuditThresholdOk

`func (o *IngestBudgetDefinition) GetAuditThresholdOk() (*int32, bool)`

GetAuditThresholdOk returns a tuple with the AuditThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditThreshold

`func (o *IngestBudgetDefinition) SetAuditThreshold(v int32)`

SetAuditThreshold sets AuditThreshold field to given value.

### HasAuditThreshold

`func (o *IngestBudgetDefinition) HasAuditThreshold() bool`

HasAuditThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



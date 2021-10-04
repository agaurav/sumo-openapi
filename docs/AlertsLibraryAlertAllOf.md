# AlertsLibraryAlertAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorId** | Pointer to **string** | The Id of the associated monitor. | [optional] 
**ResolvedAt** | Pointer to **NullableTime** | The time at which the alert was resolved. | [optional] 
**AbnormalityStartTime** | Pointer to **time.Time** | The time at which the incident started. | [optional] 
**AlertType** | Pointer to **string** | The severity of the Alert. Valid values:   1. &#x60;Critical&#x60;   2. &#x60;Warning&#x60;   3. &#x60;MissingData&#x60; | [optional] 
**Status** | Pointer to **string** | The status of the Alert. Valid values:   1. &#x60;Triggered&#x60;   2. &#x60;Resolved&#x60; | [optional] 
**MonitorQueries** | Pointer to [**[]AlertMonitorQuery**](AlertMonitorQuery.md) | All queries from the monitor relevant to the alert. | [optional] 
**TriggerQueries** | Pointer to [**[]AlertMonitorQuery**](AlertMonitorQuery.md) | All queries from the monitor relevant to the alert with triggered time series filters. | [optional] 
**MonitorUrl** | Pointer to **string** | URL for this monitor&#39;s view page | [optional] 
**TriggerQueryUrl** | Pointer to **string** | A link to search with the triggering data and time range | [optional] 
**TriggerConditions** | Pointer to [**[]TriggerCondition**](TriggerCondition.md) | Trigger conditions which were breached to create this Alert. | [optional] 
**TriggerValue** | Pointer to **float64** | The of the query result which breached the trigger condition. | [optional] 
**MonitorType** | Pointer to **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor. | [optional] 
**EntityIds** | Pointer to **[]string** | One or more entity identifiers involved in this Alert. | [optional] 
**Notes** | Pointer to **string** |  | [optional] 
**ExtraDetails** | Pointer to [**ExtraDetails**](ExtraDetails.md) |  | [optional] 

## Methods

### NewAlertsLibraryAlertAllOf

`func NewAlertsLibraryAlertAllOf() *AlertsLibraryAlertAllOf`

NewAlertsLibraryAlertAllOf instantiates a new AlertsLibraryAlertAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsLibraryAlertAllOfWithDefaults

`func NewAlertsLibraryAlertAllOfWithDefaults() *AlertsLibraryAlertAllOf`

NewAlertsLibraryAlertAllOfWithDefaults instantiates a new AlertsLibraryAlertAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *AlertsLibraryAlertAllOf) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *AlertsLibraryAlertAllOf) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *AlertsLibraryAlertAllOf) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *AlertsLibraryAlertAllOf) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetResolvedAt

`func (o *AlertsLibraryAlertAllOf) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *AlertsLibraryAlertAllOf) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *AlertsLibraryAlertAllOf) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *AlertsLibraryAlertAllOf) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *AlertsLibraryAlertAllOf) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *AlertsLibraryAlertAllOf) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetAbnormalityStartTime

`func (o *AlertsLibraryAlertAllOf) GetAbnormalityStartTime() time.Time`

GetAbnormalityStartTime returns the AbnormalityStartTime field if non-nil, zero value otherwise.

### GetAbnormalityStartTimeOk

`func (o *AlertsLibraryAlertAllOf) GetAbnormalityStartTimeOk() (*time.Time, bool)`

GetAbnormalityStartTimeOk returns a tuple with the AbnormalityStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalityStartTime

`func (o *AlertsLibraryAlertAllOf) SetAbnormalityStartTime(v time.Time)`

SetAbnormalityStartTime sets AbnormalityStartTime field to given value.

### HasAbnormalityStartTime

`func (o *AlertsLibraryAlertAllOf) HasAbnormalityStartTime() bool`

HasAbnormalityStartTime returns a boolean if a field has been set.

### GetAlertType

`func (o *AlertsLibraryAlertAllOf) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertsLibraryAlertAllOf) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertsLibraryAlertAllOf) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *AlertsLibraryAlertAllOf) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetStatus

`func (o *AlertsLibraryAlertAllOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertsLibraryAlertAllOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertsLibraryAlertAllOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertsLibraryAlertAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMonitorQueries

`func (o *AlertsLibraryAlertAllOf) GetMonitorQueries() []AlertMonitorQuery`

GetMonitorQueries returns the MonitorQueries field if non-nil, zero value otherwise.

### GetMonitorQueriesOk

`func (o *AlertsLibraryAlertAllOf) GetMonitorQueriesOk() (*[]AlertMonitorQuery, bool)`

GetMonitorQueriesOk returns a tuple with the MonitorQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorQueries

`func (o *AlertsLibraryAlertAllOf) SetMonitorQueries(v []AlertMonitorQuery)`

SetMonitorQueries sets MonitorQueries field to given value.

### HasMonitorQueries

`func (o *AlertsLibraryAlertAllOf) HasMonitorQueries() bool`

HasMonitorQueries returns a boolean if a field has been set.

### GetTriggerQueries

`func (o *AlertsLibraryAlertAllOf) GetTriggerQueries() []AlertMonitorQuery`

GetTriggerQueries returns the TriggerQueries field if non-nil, zero value otherwise.

### GetTriggerQueriesOk

`func (o *AlertsLibraryAlertAllOf) GetTriggerQueriesOk() (*[]AlertMonitorQuery, bool)`

GetTriggerQueriesOk returns a tuple with the TriggerQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerQueries

`func (o *AlertsLibraryAlertAllOf) SetTriggerQueries(v []AlertMonitorQuery)`

SetTriggerQueries sets TriggerQueries field to given value.

### HasTriggerQueries

`func (o *AlertsLibraryAlertAllOf) HasTriggerQueries() bool`

HasTriggerQueries returns a boolean if a field has been set.

### GetMonitorUrl

`func (o *AlertsLibraryAlertAllOf) GetMonitorUrl() string`

GetMonitorUrl returns the MonitorUrl field if non-nil, zero value otherwise.

### GetMonitorUrlOk

`func (o *AlertsLibraryAlertAllOf) GetMonitorUrlOk() (*string, bool)`

GetMonitorUrlOk returns a tuple with the MonitorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUrl

`func (o *AlertsLibraryAlertAllOf) SetMonitorUrl(v string)`

SetMonitorUrl sets MonitorUrl field to given value.

### HasMonitorUrl

`func (o *AlertsLibraryAlertAllOf) HasMonitorUrl() bool`

HasMonitorUrl returns a boolean if a field has been set.

### GetTriggerQueryUrl

`func (o *AlertsLibraryAlertAllOf) GetTriggerQueryUrl() string`

GetTriggerQueryUrl returns the TriggerQueryUrl field if non-nil, zero value otherwise.

### GetTriggerQueryUrlOk

`func (o *AlertsLibraryAlertAllOf) GetTriggerQueryUrlOk() (*string, bool)`

GetTriggerQueryUrlOk returns a tuple with the TriggerQueryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerQueryUrl

`func (o *AlertsLibraryAlertAllOf) SetTriggerQueryUrl(v string)`

SetTriggerQueryUrl sets TriggerQueryUrl field to given value.

### HasTriggerQueryUrl

`func (o *AlertsLibraryAlertAllOf) HasTriggerQueryUrl() bool`

HasTriggerQueryUrl returns a boolean if a field has been set.

### GetTriggerConditions

`func (o *AlertsLibraryAlertAllOf) GetTriggerConditions() []TriggerCondition`

GetTriggerConditions returns the TriggerConditions field if non-nil, zero value otherwise.

### GetTriggerConditionsOk

`func (o *AlertsLibraryAlertAllOf) GetTriggerConditionsOk() (*[]TriggerCondition, bool)`

GetTriggerConditionsOk returns a tuple with the TriggerConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerConditions

`func (o *AlertsLibraryAlertAllOf) SetTriggerConditions(v []TriggerCondition)`

SetTriggerConditions sets TriggerConditions field to given value.

### HasTriggerConditions

`func (o *AlertsLibraryAlertAllOf) HasTriggerConditions() bool`

HasTriggerConditions returns a boolean if a field has been set.

### GetTriggerValue

`func (o *AlertsLibraryAlertAllOf) GetTriggerValue() float64`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AlertsLibraryAlertAllOf) GetTriggerValueOk() (*float64, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AlertsLibraryAlertAllOf) SetTriggerValue(v float64)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *AlertsLibraryAlertAllOf) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetMonitorType

`func (o *AlertsLibraryAlertAllOf) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *AlertsLibraryAlertAllOf) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *AlertsLibraryAlertAllOf) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *AlertsLibraryAlertAllOf) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetEntityIds

`func (o *AlertsLibraryAlertAllOf) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *AlertsLibraryAlertAllOf) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *AlertsLibraryAlertAllOf) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *AlertsLibraryAlertAllOf) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetNotes

`func (o *AlertsLibraryAlertAllOf) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AlertsLibraryAlertAllOf) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AlertsLibraryAlertAllOf) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AlertsLibraryAlertAllOf) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetExtraDetails

`func (o *AlertsLibraryAlertAllOf) GetExtraDetails() ExtraDetails`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *AlertsLibraryAlertAllOf) GetExtraDetailsOk() (*ExtraDetails, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *AlertsLibraryAlertAllOf) SetExtraDetails(v ExtraDetails)`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *AlertsLibraryAlertAllOf) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



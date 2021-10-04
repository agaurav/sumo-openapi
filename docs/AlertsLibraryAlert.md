# AlertsLibraryAlert

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

### NewAlertsLibraryAlert

`func NewAlertsLibraryAlert() *AlertsLibraryAlert`

NewAlertsLibraryAlert instantiates a new AlertsLibraryAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsLibraryAlertWithDefaults

`func NewAlertsLibraryAlertWithDefaults() *AlertsLibraryAlert`

NewAlertsLibraryAlertWithDefaults instantiates a new AlertsLibraryAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorId

`func (o *AlertsLibraryAlert) GetMonitorId() string`

GetMonitorId returns the MonitorId field if non-nil, zero value otherwise.

### GetMonitorIdOk

`func (o *AlertsLibraryAlert) GetMonitorIdOk() (*string, bool)`

GetMonitorIdOk returns a tuple with the MonitorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorId

`func (o *AlertsLibraryAlert) SetMonitorId(v string)`

SetMonitorId sets MonitorId field to given value.

### HasMonitorId

`func (o *AlertsLibraryAlert) HasMonitorId() bool`

HasMonitorId returns a boolean if a field has been set.

### GetResolvedAt

`func (o *AlertsLibraryAlert) GetResolvedAt() time.Time`

GetResolvedAt returns the ResolvedAt field if non-nil, zero value otherwise.

### GetResolvedAtOk

`func (o *AlertsLibraryAlert) GetResolvedAtOk() (*time.Time, bool)`

GetResolvedAtOk returns a tuple with the ResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedAt

`func (o *AlertsLibraryAlert) SetResolvedAt(v time.Time)`

SetResolvedAt sets ResolvedAt field to given value.

### HasResolvedAt

`func (o *AlertsLibraryAlert) HasResolvedAt() bool`

HasResolvedAt returns a boolean if a field has been set.

### SetResolvedAtNil

`func (o *AlertsLibraryAlert) SetResolvedAtNil(b bool)`

 SetResolvedAtNil sets the value for ResolvedAt to be an explicit nil

### UnsetResolvedAt
`func (o *AlertsLibraryAlert) UnsetResolvedAt()`

UnsetResolvedAt ensures that no value is present for ResolvedAt, not even an explicit nil
### GetAbnormalityStartTime

`func (o *AlertsLibraryAlert) GetAbnormalityStartTime() time.Time`

GetAbnormalityStartTime returns the AbnormalityStartTime field if non-nil, zero value otherwise.

### GetAbnormalityStartTimeOk

`func (o *AlertsLibraryAlert) GetAbnormalityStartTimeOk() (*time.Time, bool)`

GetAbnormalityStartTimeOk returns a tuple with the AbnormalityStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbnormalityStartTime

`func (o *AlertsLibraryAlert) SetAbnormalityStartTime(v time.Time)`

SetAbnormalityStartTime sets AbnormalityStartTime field to given value.

### HasAbnormalityStartTime

`func (o *AlertsLibraryAlert) HasAbnormalityStartTime() bool`

HasAbnormalityStartTime returns a boolean if a field has been set.

### GetAlertType

`func (o *AlertsLibraryAlert) GetAlertType() string`

GetAlertType returns the AlertType field if non-nil, zero value otherwise.

### GetAlertTypeOk

`func (o *AlertsLibraryAlert) GetAlertTypeOk() (*string, bool)`

GetAlertTypeOk returns a tuple with the AlertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertType

`func (o *AlertsLibraryAlert) SetAlertType(v string)`

SetAlertType sets AlertType field to given value.

### HasAlertType

`func (o *AlertsLibraryAlert) HasAlertType() bool`

HasAlertType returns a boolean if a field has been set.

### GetStatus

`func (o *AlertsLibraryAlert) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AlertsLibraryAlert) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AlertsLibraryAlert) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AlertsLibraryAlert) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMonitorQueries

`func (o *AlertsLibraryAlert) GetMonitorQueries() []AlertMonitorQuery`

GetMonitorQueries returns the MonitorQueries field if non-nil, zero value otherwise.

### GetMonitorQueriesOk

`func (o *AlertsLibraryAlert) GetMonitorQueriesOk() (*[]AlertMonitorQuery, bool)`

GetMonitorQueriesOk returns a tuple with the MonitorQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorQueries

`func (o *AlertsLibraryAlert) SetMonitorQueries(v []AlertMonitorQuery)`

SetMonitorQueries sets MonitorQueries field to given value.

### HasMonitorQueries

`func (o *AlertsLibraryAlert) HasMonitorQueries() bool`

HasMonitorQueries returns a boolean if a field has been set.

### GetTriggerQueries

`func (o *AlertsLibraryAlert) GetTriggerQueries() []AlertMonitorQuery`

GetTriggerQueries returns the TriggerQueries field if non-nil, zero value otherwise.

### GetTriggerQueriesOk

`func (o *AlertsLibraryAlert) GetTriggerQueriesOk() (*[]AlertMonitorQuery, bool)`

GetTriggerQueriesOk returns a tuple with the TriggerQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerQueries

`func (o *AlertsLibraryAlert) SetTriggerQueries(v []AlertMonitorQuery)`

SetTriggerQueries sets TriggerQueries field to given value.

### HasTriggerQueries

`func (o *AlertsLibraryAlert) HasTriggerQueries() bool`

HasTriggerQueries returns a boolean if a field has been set.

### GetMonitorUrl

`func (o *AlertsLibraryAlert) GetMonitorUrl() string`

GetMonitorUrl returns the MonitorUrl field if non-nil, zero value otherwise.

### GetMonitorUrlOk

`func (o *AlertsLibraryAlert) GetMonitorUrlOk() (*string, bool)`

GetMonitorUrlOk returns a tuple with the MonitorUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUrl

`func (o *AlertsLibraryAlert) SetMonitorUrl(v string)`

SetMonitorUrl sets MonitorUrl field to given value.

### HasMonitorUrl

`func (o *AlertsLibraryAlert) HasMonitorUrl() bool`

HasMonitorUrl returns a boolean if a field has been set.

### GetTriggerQueryUrl

`func (o *AlertsLibraryAlert) GetTriggerQueryUrl() string`

GetTriggerQueryUrl returns the TriggerQueryUrl field if non-nil, zero value otherwise.

### GetTriggerQueryUrlOk

`func (o *AlertsLibraryAlert) GetTriggerQueryUrlOk() (*string, bool)`

GetTriggerQueryUrlOk returns a tuple with the TriggerQueryUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerQueryUrl

`func (o *AlertsLibraryAlert) SetTriggerQueryUrl(v string)`

SetTriggerQueryUrl sets TriggerQueryUrl field to given value.

### HasTriggerQueryUrl

`func (o *AlertsLibraryAlert) HasTriggerQueryUrl() bool`

HasTriggerQueryUrl returns a boolean if a field has been set.

### GetTriggerConditions

`func (o *AlertsLibraryAlert) GetTriggerConditions() []TriggerCondition`

GetTriggerConditions returns the TriggerConditions field if non-nil, zero value otherwise.

### GetTriggerConditionsOk

`func (o *AlertsLibraryAlert) GetTriggerConditionsOk() (*[]TriggerCondition, bool)`

GetTriggerConditionsOk returns a tuple with the TriggerConditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerConditions

`func (o *AlertsLibraryAlert) SetTriggerConditions(v []TriggerCondition)`

SetTriggerConditions sets TriggerConditions field to given value.

### HasTriggerConditions

`func (o *AlertsLibraryAlert) HasTriggerConditions() bool`

HasTriggerConditions returns a boolean if a field has been set.

### GetTriggerValue

`func (o *AlertsLibraryAlert) GetTriggerValue() float64`

GetTriggerValue returns the TriggerValue field if non-nil, zero value otherwise.

### GetTriggerValueOk

`func (o *AlertsLibraryAlert) GetTriggerValueOk() (*float64, bool)`

GetTriggerValueOk returns a tuple with the TriggerValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerValue

`func (o *AlertsLibraryAlert) SetTriggerValue(v float64)`

SetTriggerValue sets TriggerValue field to given value.

### HasTriggerValue

`func (o *AlertsLibraryAlert) HasTriggerValue() bool`

HasTriggerValue returns a boolean if a field has been set.

### GetMonitorType

`func (o *AlertsLibraryAlert) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *AlertsLibraryAlert) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *AlertsLibraryAlert) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *AlertsLibraryAlert) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetEntityIds

`func (o *AlertsLibraryAlert) GetEntityIds() []string`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *AlertsLibraryAlert) GetEntityIdsOk() (*[]string, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *AlertsLibraryAlert) SetEntityIds(v []string)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *AlertsLibraryAlert) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetNotes

`func (o *AlertsLibraryAlert) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *AlertsLibraryAlert) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *AlertsLibraryAlert) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *AlertsLibraryAlert) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetExtraDetails

`func (o *AlertsLibraryAlert) GetExtraDetails() ExtraDetails`

GetExtraDetails returns the ExtraDetails field if non-nil, zero value otherwise.

### GetExtraDetailsOk

`func (o *AlertsLibraryAlert) GetExtraDetailsOk() (*ExtraDetails, bool)`

GetExtraDetailsOk returns a tuple with the ExtraDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraDetails

`func (o *AlertsLibraryAlert) SetExtraDetails(v ExtraDetails)`

SetExtraDetails sets ExtraDetails field to given value.

### HasExtraDetails

`func (o *AlertsLibraryAlert) HasExtraDetails() bool`

HasExtraDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



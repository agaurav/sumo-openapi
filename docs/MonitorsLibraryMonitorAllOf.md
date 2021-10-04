# MonitorsLibraryMonitorAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor. | 
**EvaluationDelay** | Pointer to **string** | The delay duration for evaluating the monitor (relative to current time). The timerange of monitor will be shifted in the past by this delay time. | [optional] [default to "0m"]
**NotificationGroupFields** | Pointer to **[]string** | The set of fields to be used to group alert notifications for a monitor. The value of this field will be considered only when &#39;groupNotifications&#39; is true. | [optional] 
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | 
**Triggers** | [**[]TriggerCondition**](TriggerCondition.md) | Defines the conditions of when to send notifications. | 
**Notifications** | Pointer to [**[]MonitorNotification**](MonitorNotification.md) | The notifications the monitor will send when the respective trigger condition is met. | [optional] [default to []]
**IsDisabled** | Pointer to **bool** | Whether or not the monitor is disabled. Disabled monitors will not run, and will not generate or send notifications. | [optional] [default to false]
**GroupNotifications** | Pointer to **bool** | Whether or not to group notifications for individual items that meet the trigger condition. | [optional] [default to true]
**Playbook** | Pointer to **string** | Notes such as links and instruction to help you resolve alerts triggered by this monitor. {{Markdown}} supported. It will be enabled only if available for your organization. Please contact your Sumo Logic account team to learn more. | [optional] [default to ""]

## Methods

### NewMonitorsLibraryMonitorAllOf

`func NewMonitorsLibraryMonitorAllOf(monitorType string, queries []MonitorQuery, triggers []TriggerCondition, ) *MonitorsLibraryMonitorAllOf`

NewMonitorsLibraryMonitorAllOf instantiates a new MonitorsLibraryMonitorAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorsLibraryMonitorAllOfWithDefaults

`func NewMonitorsLibraryMonitorAllOfWithDefaults() *MonitorsLibraryMonitorAllOf`

NewMonitorsLibraryMonitorAllOfWithDefaults instantiates a new MonitorsLibraryMonitorAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorsLibraryMonitorAllOf) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorsLibraryMonitorAllOf) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorsLibraryMonitorAllOf) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.


### GetEvaluationDelay

`func (o *MonitorsLibraryMonitorAllOf) GetEvaluationDelay() string`

GetEvaluationDelay returns the EvaluationDelay field if non-nil, zero value otherwise.

### GetEvaluationDelayOk

`func (o *MonitorsLibraryMonitorAllOf) GetEvaluationDelayOk() (*string, bool)`

GetEvaluationDelayOk returns a tuple with the EvaluationDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluationDelay

`func (o *MonitorsLibraryMonitorAllOf) SetEvaluationDelay(v string)`

SetEvaluationDelay sets EvaluationDelay field to given value.

### HasEvaluationDelay

`func (o *MonitorsLibraryMonitorAllOf) HasEvaluationDelay() bool`

HasEvaluationDelay returns a boolean if a field has been set.

### GetNotificationGroupFields

`func (o *MonitorsLibraryMonitorAllOf) GetNotificationGroupFields() []string`

GetNotificationGroupFields returns the NotificationGroupFields field if non-nil, zero value otherwise.

### GetNotificationGroupFieldsOk

`func (o *MonitorsLibraryMonitorAllOf) GetNotificationGroupFieldsOk() (*[]string, bool)`

GetNotificationGroupFieldsOk returns a tuple with the NotificationGroupFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationGroupFields

`func (o *MonitorsLibraryMonitorAllOf) SetNotificationGroupFields(v []string)`

SetNotificationGroupFields sets NotificationGroupFields field to given value.

### HasNotificationGroupFields

`func (o *MonitorsLibraryMonitorAllOf) HasNotificationGroupFields() bool`

HasNotificationGroupFields returns a boolean if a field has been set.

### GetQueries

`func (o *MonitorsLibraryMonitorAllOf) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorsLibraryMonitorAllOf) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorsLibraryMonitorAllOf) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.


### GetTriggers

`func (o *MonitorsLibraryMonitorAllOf) GetTriggers() []TriggerCondition`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *MonitorsLibraryMonitorAllOf) GetTriggersOk() (*[]TriggerCondition, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *MonitorsLibraryMonitorAllOf) SetTriggers(v []TriggerCondition)`

SetTriggers sets Triggers field to given value.


### GetNotifications

`func (o *MonitorsLibraryMonitorAllOf) GetNotifications() []MonitorNotification`

GetNotifications returns the Notifications field if non-nil, zero value otherwise.

### GetNotificationsOk

`func (o *MonitorsLibraryMonitorAllOf) GetNotificationsOk() (*[]MonitorNotification, bool)`

GetNotificationsOk returns a tuple with the Notifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifications

`func (o *MonitorsLibraryMonitorAllOf) SetNotifications(v []MonitorNotification)`

SetNotifications sets Notifications field to given value.

### HasNotifications

`func (o *MonitorsLibraryMonitorAllOf) HasNotifications() bool`

HasNotifications returns a boolean if a field has been set.

### GetIsDisabled

`func (o *MonitorsLibraryMonitorAllOf) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *MonitorsLibraryMonitorAllOf) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *MonitorsLibraryMonitorAllOf) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *MonitorsLibraryMonitorAllOf) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetGroupNotifications

`func (o *MonitorsLibraryMonitorAllOf) GetGroupNotifications() bool`

GetGroupNotifications returns the GroupNotifications field if non-nil, zero value otherwise.

### GetGroupNotificationsOk

`func (o *MonitorsLibraryMonitorAllOf) GetGroupNotificationsOk() (*bool, bool)`

GetGroupNotificationsOk returns a tuple with the GroupNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupNotifications

`func (o *MonitorsLibraryMonitorAllOf) SetGroupNotifications(v bool)`

SetGroupNotifications sets GroupNotifications field to given value.

### HasGroupNotifications

`func (o *MonitorsLibraryMonitorAllOf) HasGroupNotifications() bool`

HasGroupNotifications returns a boolean if a field has been set.

### GetPlaybook

`func (o *MonitorsLibraryMonitorAllOf) GetPlaybook() string`

GetPlaybook returns the Playbook field if non-nil, zero value otherwise.

### GetPlaybookOk

`func (o *MonitorsLibraryMonitorAllOf) GetPlaybookOk() (*string, bool)`

GetPlaybookOk returns a tuple with the Playbook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaybook

`func (o *MonitorsLibraryMonitorAllOf) SetPlaybook(v string)`

SetPlaybook sets Playbook field to given value.

### HasPlaybook

`func (o *MonitorsLibraryMonitorAllOf) HasPlaybook() bool`

HasPlaybook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



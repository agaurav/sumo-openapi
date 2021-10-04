# SearchScheduleSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CronExpression** | Pointer to **string** | Cron-like expression specifying the search&#39;s schedule. Field scheduleType must be set to \&quot;Custom\&quot;, otherwise, scheduleType takes precedence over cronExpression. | [optional] 
**DisplayableTimeRange** | Pointer to **string** | A human-friendly text describing the query time range. For e.g. \&quot;-2h\&quot;, \&quot;last three days\&quot;, \&quot;team default time\&quot; | [optional] 
**ParseableTimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**TimeZone** | **string** | Time zone identifier for time specification. Either an abbreviation such as \&quot;PST\&quot;, a full name such as \&quot;America/Los_Angeles\&quot;, or a custom ID such as \&quot;GMT-8:00\&quot;. Note that the support of abbreviations is for JDK 1.1.x compatibility only and full names should be used. | 
**Threshold** | Pointer to [**NotificationThresholdSyncDefinition**](NotificationThresholdSyncDefinition.md) |  | [optional] 
**Notification** | [**ScheduleNotificationSyncDefinition**](ScheduleNotificationSyncDefinition.md) |  | 
**ScheduleType** | **string** | Run schedule of the scheduled search. Set to \&quot;Custom\&quot; to specify the schedule with a CRON expression. Possible schedule types are:   - &#x60;RealTime&#x60;   - &#x60;15Minutes&#x60;   - &#x60;1Hour&#x60;   - &#x60;2Hours&#x60;   - &#x60;4Hours&#x60;   - &#x60;6Hours&#x60;   - &#x60;8Hours&#x60;   - &#x60;12Hours&#x60;   - &#x60;1Day&#x60;   - &#x60;1Week&#x60;   - &#x60;Custom&#x60; | 
**MuteErrorEmails** | Pointer to **bool** | If enabled, emails are not sent out in case of errors with the search. | [optional] 
**Parameters** | [**[]ScheduleSearchParameterSyncDefinition**](ScheduleSearchParameterSyncDefinition.md) | A list of scheduled search parameters. | 

## Methods

### NewSearchScheduleSyncDefinition

`func NewSearchScheduleSyncDefinition(parseableTimeRange ResolvableTimeRange, timeZone string, notification ScheduleNotificationSyncDefinition, scheduleType string, parameters []ScheduleSearchParameterSyncDefinition, ) *SearchScheduleSyncDefinition`

NewSearchScheduleSyncDefinition instantiates a new SearchScheduleSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchScheduleSyncDefinitionWithDefaults

`func NewSearchScheduleSyncDefinitionWithDefaults() *SearchScheduleSyncDefinition`

NewSearchScheduleSyncDefinitionWithDefaults instantiates a new SearchScheduleSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCronExpression

`func (o *SearchScheduleSyncDefinition) GetCronExpression() string`

GetCronExpression returns the CronExpression field if non-nil, zero value otherwise.

### GetCronExpressionOk

`func (o *SearchScheduleSyncDefinition) GetCronExpressionOk() (*string, bool)`

GetCronExpressionOk returns a tuple with the CronExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCronExpression

`func (o *SearchScheduleSyncDefinition) SetCronExpression(v string)`

SetCronExpression sets CronExpression field to given value.

### HasCronExpression

`func (o *SearchScheduleSyncDefinition) HasCronExpression() bool`

HasCronExpression returns a boolean if a field has been set.

### GetDisplayableTimeRange

`func (o *SearchScheduleSyncDefinition) GetDisplayableTimeRange() string`

GetDisplayableTimeRange returns the DisplayableTimeRange field if non-nil, zero value otherwise.

### GetDisplayableTimeRangeOk

`func (o *SearchScheduleSyncDefinition) GetDisplayableTimeRangeOk() (*string, bool)`

GetDisplayableTimeRangeOk returns a tuple with the DisplayableTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayableTimeRange

`func (o *SearchScheduleSyncDefinition) SetDisplayableTimeRange(v string)`

SetDisplayableTimeRange sets DisplayableTimeRange field to given value.

### HasDisplayableTimeRange

`func (o *SearchScheduleSyncDefinition) HasDisplayableTimeRange() bool`

HasDisplayableTimeRange returns a boolean if a field has been set.

### GetParseableTimeRange

`func (o *SearchScheduleSyncDefinition) GetParseableTimeRange() ResolvableTimeRange`

GetParseableTimeRange returns the ParseableTimeRange field if non-nil, zero value otherwise.

### GetParseableTimeRangeOk

`func (o *SearchScheduleSyncDefinition) GetParseableTimeRangeOk() (*ResolvableTimeRange, bool)`

GetParseableTimeRangeOk returns a tuple with the ParseableTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseableTimeRange

`func (o *SearchScheduleSyncDefinition) SetParseableTimeRange(v ResolvableTimeRange)`

SetParseableTimeRange sets ParseableTimeRange field to given value.


### GetTimeZone

`func (o *SearchScheduleSyncDefinition) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *SearchScheduleSyncDefinition) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *SearchScheduleSyncDefinition) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.


### GetThreshold

`func (o *SearchScheduleSyncDefinition) GetThreshold() NotificationThresholdSyncDefinition`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *SearchScheduleSyncDefinition) GetThresholdOk() (*NotificationThresholdSyncDefinition, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *SearchScheduleSyncDefinition) SetThreshold(v NotificationThresholdSyncDefinition)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *SearchScheduleSyncDefinition) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetNotification

`func (o *SearchScheduleSyncDefinition) GetNotification() ScheduleNotificationSyncDefinition`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *SearchScheduleSyncDefinition) GetNotificationOk() (*ScheduleNotificationSyncDefinition, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *SearchScheduleSyncDefinition) SetNotification(v ScheduleNotificationSyncDefinition)`

SetNotification sets Notification field to given value.


### GetScheduleType

`func (o *SearchScheduleSyncDefinition) GetScheduleType() string`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *SearchScheduleSyncDefinition) GetScheduleTypeOk() (*string, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *SearchScheduleSyncDefinition) SetScheduleType(v string)`

SetScheduleType sets ScheduleType field to given value.


### GetMuteErrorEmails

`func (o *SearchScheduleSyncDefinition) GetMuteErrorEmails() bool`

GetMuteErrorEmails returns the MuteErrorEmails field if non-nil, zero value otherwise.

### GetMuteErrorEmailsOk

`func (o *SearchScheduleSyncDefinition) GetMuteErrorEmailsOk() (*bool, bool)`

GetMuteErrorEmailsOk returns a tuple with the MuteErrorEmails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuteErrorEmails

`func (o *SearchScheduleSyncDefinition) SetMuteErrorEmails(v bool)`

SetMuteErrorEmails sets MuteErrorEmails field to given value.

### HasMuteErrorEmails

`func (o *SearchScheduleSyncDefinition) HasMuteErrorEmails() bool`

HasMuteErrorEmails returns a boolean if a field has been set.

### GetParameters

`func (o *SearchScheduleSyncDefinition) GetParameters() []ScheduleSearchParameterSyncDefinition`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *SearchScheduleSyncDefinition) GetParametersOk() (*[]ScheduleSearchParameterSyncDefinition, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *SearchScheduleSyncDefinition) SetParameters(v []ScheduleSearchParameterSyncDefinition)`

SetParameters sets Parameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



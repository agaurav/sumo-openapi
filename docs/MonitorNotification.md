# MonitorNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Notification** | [**Action**](Action.md) |  | 
**RunForTriggerTypes** | **[]string** | The trigger types assigned to send this notification. | 

## Methods

### NewMonitorNotification

`func NewMonitorNotification(notification Action, runForTriggerTypes []string, ) *MonitorNotification`

NewMonitorNotification instantiates a new MonitorNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorNotificationWithDefaults

`func NewMonitorNotificationWithDefaults() *MonitorNotification`

NewMonitorNotificationWithDefaults instantiates a new MonitorNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNotification

`func (o *MonitorNotification) GetNotification() Action`

GetNotification returns the Notification field if non-nil, zero value otherwise.

### GetNotificationOk

`func (o *MonitorNotification) GetNotificationOk() (*Action, bool)`

GetNotificationOk returns a tuple with the Notification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotification

`func (o *MonitorNotification) SetNotification(v Action)`

SetNotification sets Notification field to given value.


### GetRunForTriggerTypes

`func (o *MonitorNotification) GetRunForTriggerTypes() []string`

GetRunForTriggerTypes returns the RunForTriggerTypes field if non-nil, zero value otherwise.

### GetRunForTriggerTypesOk

`func (o *MonitorNotification) GetRunForTriggerTypesOk() (*[]string, bool)`

GetRunForTriggerTypesOk returns a tuple with the RunForTriggerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunForTriggerTypes

`func (o *MonitorNotification) SetRunForTriggerTypes(v []string)`

SetRunForTriggerTypes sets RunForTriggerTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



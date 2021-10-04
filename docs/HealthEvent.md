# HealthEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventId** | **string** | The unique identifier of the event. | 
**EventName** | **string** | The name of the event. | 
**Details** | [**TrackerIdentity**](TrackerIdentity.md) |  | 
**ResourceIdentity** | [**ResourceIdentity**](ResourceIdentity.md) |  | 
**EventTime** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**Subsystem** | **string** | The product area of the event. | 
**SeverityLevel** | **string** | The criticality of the event. It is either &#x60;Error&#x60; or &#x60;Warning&#x60; | 

## Methods

### NewHealthEvent

`func NewHealthEvent(eventId string, eventName string, details TrackerIdentity, resourceIdentity ResourceIdentity, eventTime time.Time, subsystem string, severityLevel string, ) *HealthEvent`

NewHealthEvent instantiates a new HealthEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthEventWithDefaults

`func NewHealthEventWithDefaults() *HealthEvent`

NewHealthEventWithDefaults instantiates a new HealthEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventId

`func (o *HealthEvent) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *HealthEvent) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *HealthEvent) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetEventName

`func (o *HealthEvent) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *HealthEvent) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *HealthEvent) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetDetails

`func (o *HealthEvent) GetDetails() TrackerIdentity`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HealthEvent) GetDetailsOk() (*TrackerIdentity, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HealthEvent) SetDetails(v TrackerIdentity)`

SetDetails sets Details field to given value.


### GetResourceIdentity

`func (o *HealthEvent) GetResourceIdentity() ResourceIdentity`

GetResourceIdentity returns the ResourceIdentity field if non-nil, zero value otherwise.

### GetResourceIdentityOk

`func (o *HealthEvent) GetResourceIdentityOk() (*ResourceIdentity, bool)`

GetResourceIdentityOk returns a tuple with the ResourceIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceIdentity

`func (o *HealthEvent) SetResourceIdentity(v ResourceIdentity)`

SetResourceIdentity sets ResourceIdentity field to given value.


### GetEventTime

`func (o *HealthEvent) GetEventTime() time.Time`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *HealthEvent) GetEventTimeOk() (*time.Time, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *HealthEvent) SetEventTime(v time.Time)`

SetEventTime sets EventTime field to given value.


### GetSubsystem

`func (o *HealthEvent) GetSubsystem() string`

GetSubsystem returns the Subsystem field if non-nil, zero value otherwise.

### GetSubsystemOk

`func (o *HealthEvent) GetSubsystemOk() (*string, bool)`

GetSubsystemOk returns a tuple with the Subsystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubsystem

`func (o *HealthEvent) SetSubsystem(v string)`

SetSubsystem sets Subsystem field to given value.


### GetSeverityLevel

`func (o *HealthEvent) GetSeverityLevel() string`

GetSeverityLevel returns the SeverityLevel field if non-nil, zero value otherwise.

### GetSeverityLevelOk

`func (o *HealthEvent) GetSeverityLevelOk() (*string, bool)`

GetSeverityLevelOk returns a tuple with the SeverityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverityLevel

`func (o *HealthEvent) SetSeverityLevel(v string)`

SetSeverityLevel sets SeverityLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



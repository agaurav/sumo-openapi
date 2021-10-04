# UpdateScheduledViewDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataForwardingId** | Pointer to **string** | An optional ID of a data forwarding configuration to be used by the scheduled view. | [optional] 
**RetentionPeriod** | Pointer to **int32** | The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled. | [optional] [default to -1]
**ReduceRetentionPeriodImmediately** | Pointer to **bool** | This is required if the newly specified &#x60;retentionPeriod&#x60; is less than the existing retention period.  In such a situation, a value of &#x60;true&#x60; says that data between the existing retention period and the new retention period should be deleted immediately; if &#x60;false&#x60;, such data will be deleted after seven days. This property is optional and ignored if the specified &#x60;retentionPeriod&#x60; is greater than or equal to the current retention period. | [optional] [default to false]

## Methods

### NewUpdateScheduledViewDefinition

`func NewUpdateScheduledViewDefinition() *UpdateScheduledViewDefinition`

NewUpdateScheduledViewDefinition instantiates a new UpdateScheduledViewDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateScheduledViewDefinitionWithDefaults

`func NewUpdateScheduledViewDefinitionWithDefaults() *UpdateScheduledViewDefinition`

NewUpdateScheduledViewDefinitionWithDefaults instantiates a new UpdateScheduledViewDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataForwardingId

`func (o *UpdateScheduledViewDefinition) GetDataForwardingId() string`

GetDataForwardingId returns the DataForwardingId field if non-nil, zero value otherwise.

### GetDataForwardingIdOk

`func (o *UpdateScheduledViewDefinition) GetDataForwardingIdOk() (*string, bool)`

GetDataForwardingIdOk returns a tuple with the DataForwardingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwardingId

`func (o *UpdateScheduledViewDefinition) SetDataForwardingId(v string)`

SetDataForwardingId sets DataForwardingId field to given value.

### HasDataForwardingId

`func (o *UpdateScheduledViewDefinition) HasDataForwardingId() bool`

HasDataForwardingId returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *UpdateScheduledViewDefinition) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *UpdateScheduledViewDefinition) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *UpdateScheduledViewDefinition) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *UpdateScheduledViewDefinition) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetReduceRetentionPeriodImmediately

`func (o *UpdateScheduledViewDefinition) GetReduceRetentionPeriodImmediately() bool`

GetReduceRetentionPeriodImmediately returns the ReduceRetentionPeriodImmediately field if non-nil, zero value otherwise.

### GetReduceRetentionPeriodImmediatelyOk

`func (o *UpdateScheduledViewDefinition) GetReduceRetentionPeriodImmediatelyOk() (*bool, bool)`

GetReduceRetentionPeriodImmediatelyOk returns a tuple with the ReduceRetentionPeriodImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceRetentionPeriodImmediately

`func (o *UpdateScheduledViewDefinition) SetReduceRetentionPeriodImmediately(v bool)`

SetReduceRetentionPeriodImmediately sets ReduceRetentionPeriodImmediately field to given value.

### HasReduceRetentionPeriodImmediately

`func (o *UpdateScheduledViewDefinition) HasReduceRetentionPeriodImmediately() bool`

HasReduceRetentionPeriodImmediately returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



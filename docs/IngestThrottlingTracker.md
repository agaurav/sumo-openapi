# IngestThrottlingTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**DataType** | Pointer to **string** | The type of data for which the rate limit was enabled. The possible values are &#x60;LogIngest&#x60; and &#x60;MetricsIngest&#x60;. | [optional] 

## Methods

### NewIngestThrottlingTracker

`func NewIngestThrottlingTracker() *IngestThrottlingTracker`

NewIngestThrottlingTracker instantiates a new IngestThrottlingTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestThrottlingTrackerWithDefaults

`func NewIngestThrottlingTrackerWithDefaults() *IngestThrottlingTracker`

NewIngestThrottlingTrackerWithDefaults instantiates a new IngestThrottlingTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *IngestThrottlingTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *IngestThrottlingTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *IngestThrottlingTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *IngestThrottlingTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetDataType

`func (o *IngestThrottlingTracker) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *IngestThrottlingTracker) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *IngestThrottlingTracker) SetDataType(v string)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *IngestThrottlingTracker) HasDataType() bool`

HasDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



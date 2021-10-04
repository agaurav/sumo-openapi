# SignalsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalType** | **string** | The type of the signal to compute. Can be &#x60;LogFluctuation&#x60;, &#x60;DimensionalityExplanation&#x60;, &#x60;GisBenchmark&#x60; or &#x60;Anomalies&#x60;  | 
**SignalId** | **string** | The id for the signal result in hex format. | 
**StartTime** | **time.Time** | Start time of the signal. | 
**EndTime** | **time.Time** | End time of the signal. | 
**Summary** | **string** | Description of the payload. | 
**Payload** | **string** | Json string for computed signal. | 
**OpenInQueries** | [**[]OpenInQuery**](OpenInQuery.md) | Raw data queries for the computed signal. | 

## Methods

### NewSignalsResponse

`func NewSignalsResponse(signalType string, signalId string, startTime time.Time, endTime time.Time, summary string, payload string, openInQueries []OpenInQuery, ) *SignalsResponse`

NewSignalsResponse instantiates a new SignalsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalsResponseWithDefaults

`func NewSignalsResponseWithDefaults() *SignalsResponse`

NewSignalsResponseWithDefaults instantiates a new SignalsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalType

`func (o *SignalsResponse) GetSignalType() string`

GetSignalType returns the SignalType field if non-nil, zero value otherwise.

### GetSignalTypeOk

`func (o *SignalsResponse) GetSignalTypeOk() (*string, bool)`

GetSignalTypeOk returns a tuple with the SignalType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalType

`func (o *SignalsResponse) SetSignalType(v string)`

SetSignalType sets SignalType field to given value.


### GetSignalId

`func (o *SignalsResponse) GetSignalId() string`

GetSignalId returns the SignalId field if non-nil, zero value otherwise.

### GetSignalIdOk

`func (o *SignalsResponse) GetSignalIdOk() (*string, bool)`

GetSignalIdOk returns a tuple with the SignalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalId

`func (o *SignalsResponse) SetSignalId(v string)`

SetSignalId sets SignalId field to given value.


### GetStartTime

`func (o *SignalsResponse) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SignalsResponse) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SignalsResponse) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetEndTime

`func (o *SignalsResponse) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *SignalsResponse) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *SignalsResponse) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.


### GetSummary

`func (o *SignalsResponse) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *SignalsResponse) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *SignalsResponse) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetPayload

`func (o *SignalsResponse) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SignalsResponse) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SignalsResponse) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetOpenInQueries

`func (o *SignalsResponse) GetOpenInQueries() []OpenInQuery`

GetOpenInQueries returns the OpenInQueries field if non-nil, zero value otherwise.

### GetOpenInQueriesOk

`func (o *SignalsResponse) GetOpenInQueriesOk() (*[]OpenInQuery, bool)`

GetOpenInQueriesOk returns a tuple with the OpenInQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenInQueries

`func (o *SignalsResponse) SetOpenInQueries(v []OpenInQuery)`

SetOpenInQueries sets OpenInQueries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



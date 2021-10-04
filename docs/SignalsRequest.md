# SignalsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalTypes** | **[]string** | A list of signal types to compute. Can be &#x60;LogFluctuation&#x60;, &#x60;DimensionalityExplanation&#x60;, &#x60;GisBenchmark&#x60; or &#x60;Anomalies&#x60;  | 
**SignalContext** | [**SignalContext**](SignalContext.md) |  | 

## Methods

### NewSignalsRequest

`func NewSignalsRequest(signalTypes []string, signalContext SignalContext, ) *SignalsRequest`

NewSignalsRequest instantiates a new SignalsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalsRequestWithDefaults

`func NewSignalsRequestWithDefaults() *SignalsRequest`

NewSignalsRequestWithDefaults instantiates a new SignalsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalTypes

`func (o *SignalsRequest) GetSignalTypes() []string`

GetSignalTypes returns the SignalTypes field if non-nil, zero value otherwise.

### GetSignalTypesOk

`func (o *SignalsRequest) GetSignalTypesOk() (*[]string, bool)`

GetSignalTypesOk returns a tuple with the SignalTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalTypes

`func (o *SignalsRequest) SetSignalTypes(v []string)`

SetSignalTypes sets SignalTypes field to given value.


### GetSignalContext

`func (o *SignalsRequest) GetSignalContext() SignalContext`

GetSignalContext returns the SignalContext field if non-nil, zero value otherwise.

### GetSignalContextOk

`func (o *SignalsRequest) GetSignalContextOk() (*SignalContext, bool)`

GetSignalContextOk returns a tuple with the SignalContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalContext

`func (o *SignalsRequest) SetSignalContext(v SignalContext)`

SetSignalContext sets SignalContext field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



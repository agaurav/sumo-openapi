# SignalsJobResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsComplete** | **bool** | Whether the signal computing job finished. | 
**Signals** | [**[]SignalsResponse**](SignalsResponse.md) | Sequence of computed signals. | 
**Warnings** | [**[]WarningDetails**](WarningDetails.md) | List of warnings while computing signals. | 

## Methods

### NewSignalsJobResult

`func NewSignalsJobResult(isComplete bool, signals []SignalsResponse, warnings []WarningDetails, ) *SignalsJobResult`

NewSignalsJobResult instantiates a new SignalsJobResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalsJobResultWithDefaults

`func NewSignalsJobResultWithDefaults() *SignalsJobResult`

NewSignalsJobResultWithDefaults instantiates a new SignalsJobResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsComplete

`func (o *SignalsJobResult) GetIsComplete() bool`

GetIsComplete returns the IsComplete field if non-nil, zero value otherwise.

### GetIsCompleteOk

`func (o *SignalsJobResult) GetIsCompleteOk() (*bool, bool)`

GetIsCompleteOk returns a tuple with the IsComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsComplete

`func (o *SignalsJobResult) SetIsComplete(v bool)`

SetIsComplete sets IsComplete field to given value.


### GetSignals

`func (o *SignalsJobResult) GetSignals() []SignalsResponse`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *SignalsJobResult) GetSignalsOk() (*[]SignalsResponse, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *SignalsJobResult) SetSignals(v []SignalsResponse)`

SetSignals sets Signals field to given value.


### GetWarnings

`func (o *SignalsJobResult) GetWarnings() []WarningDetails`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *SignalsJobResult) GetWarningsOk() (*[]WarningDetails, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *SignalsJobResult) SetWarnings(v []WarningDetails)`

SetWarnings sets Warnings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



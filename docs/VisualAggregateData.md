# VisualAggregateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | **float64** | The maximum value in the series. | 
**Min** | **float64** | The minimum value in the series. | 
**Avg** | **float64** | The average value in the series. | 
**Sum** | **float64** | The sum of all the values in the series. | 
**Latest** | **float64** | The last value in the series. | 
**Count** | Pointer to **float64** | The number of values in the series. | [optional] 

## Methods

### NewVisualAggregateData

`func NewVisualAggregateData(max float64, min float64, avg float64, sum float64, latest float64, ) *VisualAggregateData`

NewVisualAggregateData instantiates a new VisualAggregateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVisualAggregateDataWithDefaults

`func NewVisualAggregateDataWithDefaults() *VisualAggregateData`

NewVisualAggregateDataWithDefaults instantiates a new VisualAggregateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *VisualAggregateData) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *VisualAggregateData) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *VisualAggregateData) SetMax(v float64)`

SetMax sets Max field to given value.


### GetMin

`func (o *VisualAggregateData) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *VisualAggregateData) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *VisualAggregateData) SetMin(v float64)`

SetMin sets Min field to given value.


### GetAvg

`func (o *VisualAggregateData) GetAvg() float64`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *VisualAggregateData) GetAvgOk() (*float64, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *VisualAggregateData) SetAvg(v float64)`

SetAvg sets Avg field to given value.


### GetSum

`func (o *VisualAggregateData) GetSum() float64`

GetSum returns the Sum field if non-nil, zero value otherwise.

### GetSumOk

`func (o *VisualAggregateData) GetSumOk() (*float64, bool)`

GetSumOk returns a tuple with the Sum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSum

`func (o *VisualAggregateData) SetSum(v float64)`

SetSum sets Sum field to given value.


### GetLatest

`func (o *VisualAggregateData) GetLatest() float64`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *VisualAggregateData) GetLatestOk() (*float64, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *VisualAggregateData) SetLatest(v float64)`

SetLatest sets Latest field to given value.


### GetCount

`func (o *VisualAggregateData) GetCount() float64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VisualAggregateData) GetCountOk() (*float64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VisualAggregateData) SetCount(v float64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VisualAggregateData) HasCount() bool`

HasCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



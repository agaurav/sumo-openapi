# ChartDataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | Pointer to [**[]SeriesData**](SeriesData.md) | List of time series of the monitor chart data. | [optional] 

## Methods

### NewChartDataResult

`func NewChartDataResult() *ChartDataResult`

NewChartDataResult instantiates a new ChartDataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChartDataResultWithDefaults

`func NewChartDataResultWithDefaults() *ChartDataResult`

NewChartDataResultWithDefaults instantiates a new ChartDataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *ChartDataResult) GetSeries() []SeriesData`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *ChartDataResult) GetSeriesOk() (*[]SeriesData, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *ChartDataResult) SetSeries(v []SeriesData)`

SetSeries sets Series field to given value.

### HasSeries

`func (o *ChartDataResult) HasSeries() bool`

HasSeries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



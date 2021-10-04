# AlertChartDataResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Series** | [**[]SeriesData**](SeriesData.md) | List of time series of the alert chart data. | 
**Metadata** | [**AlertChartMetadata**](AlertChartMetadata.md) |  | 

## Methods

### NewAlertChartDataResult

`func NewAlertChartDataResult(series []SeriesData, metadata AlertChartMetadata, ) *AlertChartDataResult`

NewAlertChartDataResult instantiates a new AlertChartDataResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertChartDataResultWithDefaults

`func NewAlertChartDataResultWithDefaults() *AlertChartDataResult`

NewAlertChartDataResultWithDefaults instantiates a new AlertChartDataResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSeries

`func (o *AlertChartDataResult) GetSeries() []SeriesData`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *AlertChartDataResult) GetSeriesOk() (*[]SeriesData, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *AlertChartDataResult) SetSeries(v []SeriesData)`

SetSeries sets Series field to given value.


### GetMetadata

`func (o *AlertChartDataResult) GetMetadata() AlertChartMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AlertChartDataResult) GetMetadataOk() (*AlertChartMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AlertChartDataResult) SetMetadata(v AlertChartMetadata)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



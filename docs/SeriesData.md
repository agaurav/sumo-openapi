# SeriesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the series. | 
**DataPoints** | [**[]DataPoint**](DataPoint.md) | Data points of the series. | 
**SeriesAxisRange** | [**SeriesAxisRange**](SeriesAxisRange.md) |  | 
**AggregateInfo** | Pointer to [**VisualAggregateData**](VisualAggregateData.md) |  | [optional] 
**SeriesMetadata** | Pointer to [**SeriesMetadata**](SeriesMetadata.md) |  | [optional] 

## Methods

### NewSeriesData

`func NewSeriesData(name string, dataPoints []DataPoint, seriesAxisRange SeriesAxisRange, ) *SeriesData`

NewSeriesData instantiates a new SeriesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSeriesDataWithDefaults

`func NewSeriesDataWithDefaults() *SeriesData`

NewSeriesDataWithDefaults instantiates a new SeriesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SeriesData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SeriesData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SeriesData) SetName(v string)`

SetName sets Name field to given value.


### GetDataPoints

`func (o *SeriesData) GetDataPoints() []DataPoint`

GetDataPoints returns the DataPoints field if non-nil, zero value otherwise.

### GetDataPointsOk

`func (o *SeriesData) GetDataPointsOk() (*[]DataPoint, bool)`

GetDataPointsOk returns a tuple with the DataPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPoints

`func (o *SeriesData) SetDataPoints(v []DataPoint)`

SetDataPoints sets DataPoints field to given value.


### GetSeriesAxisRange

`func (o *SeriesData) GetSeriesAxisRange() SeriesAxisRange`

GetSeriesAxisRange returns the SeriesAxisRange field if non-nil, zero value otherwise.

### GetSeriesAxisRangeOk

`func (o *SeriesData) GetSeriesAxisRangeOk() (*SeriesAxisRange, bool)`

GetSeriesAxisRangeOk returns a tuple with the SeriesAxisRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesAxisRange

`func (o *SeriesData) SetSeriesAxisRange(v SeriesAxisRange)`

SetSeriesAxisRange sets SeriesAxisRange field to given value.


### GetAggregateInfo

`func (o *SeriesData) GetAggregateInfo() VisualAggregateData`

GetAggregateInfo returns the AggregateInfo field if non-nil, zero value otherwise.

### GetAggregateInfoOk

`func (o *SeriesData) GetAggregateInfoOk() (*VisualAggregateData, bool)`

GetAggregateInfoOk returns a tuple with the AggregateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregateInfo

`func (o *SeriesData) SetAggregateInfo(v VisualAggregateData)`

SetAggregateInfo sets AggregateInfo field to given value.

### HasAggregateInfo

`func (o *SeriesData) HasAggregateInfo() bool`

HasAggregateInfo returns a boolean if a field has been set.

### GetSeriesMetadata

`func (o *SeriesData) GetSeriesMetadata() SeriesMetadata`

GetSeriesMetadata returns the SeriesMetadata field if non-nil, zero value otherwise.

### GetSeriesMetadataOk

`func (o *SeriesData) GetSeriesMetadataOk() (*SeriesMetadata, bool)`

GetSeriesMetadataOk returns a tuple with the SeriesMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeriesMetadata

`func (o *SeriesData) SetSeriesMetadata(v SeriesMetadata)`

SetSeriesMetadata sets SeriesMetadata field to given value.

### HasSeriesMetadata

`func (o *SeriesData) HasSeriesMetadata() bool`

HasSeriesMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



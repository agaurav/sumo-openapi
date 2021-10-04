# StaticSeriesDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **int64** | Epoch unix time stamp. | 
**Y** | **float64** | The value of the data point. | 

## Methods

### NewStaticSeriesDataPoint

`func NewStaticSeriesDataPoint(x int64, y float64, ) *StaticSeriesDataPoint`

NewStaticSeriesDataPoint instantiates a new StaticSeriesDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticSeriesDataPointWithDefaults

`func NewStaticSeriesDataPointWithDefaults() *StaticSeriesDataPoint`

NewStaticSeriesDataPointWithDefaults instantiates a new StaticSeriesDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *StaticSeriesDataPoint) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *StaticSeriesDataPoint) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *StaticSeriesDataPoint) SetX(v int64)`

SetX sets X field to given value.


### GetY

`func (o *StaticSeriesDataPoint) GetY() float64`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *StaticSeriesDataPoint) GetYOk() (*float64, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *StaticSeriesDataPoint) SetY(v float64)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OutlierSeriesDataPoint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **int64** | Epoch unix time stamp. | 
**Y** | [**OutlierDataValue**](OutlierDataValue.md) |  | 

## Methods

### NewOutlierSeriesDataPoint

`func NewOutlierSeriesDataPoint(x int64, y OutlierDataValue, ) *OutlierSeriesDataPoint`

NewOutlierSeriesDataPoint instantiates a new OutlierSeriesDataPoint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlierSeriesDataPointWithDefaults

`func NewOutlierSeriesDataPointWithDefaults() *OutlierSeriesDataPoint`

NewOutlierSeriesDataPointWithDefaults instantiates a new OutlierSeriesDataPoint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *OutlierSeriesDataPoint) GetX() int64`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *OutlierSeriesDataPoint) GetXOk() (*int64, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *OutlierSeriesDataPoint) SetX(v int64)`

SetX sets X field to given value.


### GetY

`func (o *OutlierSeriesDataPoint) GetY() OutlierDataValue`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *OutlierSeriesDataPoint) GetYOk() (*OutlierDataValue, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *OutlierSeriesDataPoint) SetY(v OutlierDataValue)`

SetY sets Y field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



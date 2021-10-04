# DataPoints

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | [**BeginBoundedTimeRange**](BeginBoundedTimeRange.md) |  | 
**Values** | Pointer to [**[]DataValue**](DataValue.md) | An array of objects denoting the value and unit of the data points. | [optional] 

## Methods

### NewDataPoints

`func NewDataPoints(timeRange BeginBoundedTimeRange, ) *DataPoints`

NewDataPoints instantiates a new DataPoints object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataPointsWithDefaults

`func NewDataPointsWithDefaults() *DataPoints`

NewDataPointsWithDefaults instantiates a new DataPoints object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *DataPoints) GetTimeRange() BeginBoundedTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DataPoints) GetTimeRangeOk() (*BeginBoundedTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DataPoints) SetTimeRange(v BeginBoundedTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetValues

`func (o *DataPoints) GetValues() []DataValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DataPoints) GetValuesOk() (*[]DataValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DataPoints) SetValues(v []DataValue)`

SetValues sets Values field to given value.

### HasValues

`func (o *DataPoints) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



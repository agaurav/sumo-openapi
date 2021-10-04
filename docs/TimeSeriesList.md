# TimeSeriesList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeSeries** | [**[]TimeSeries**](TimeSeries.md) | A list of timeseries returned by corresponding query. | 
**Unit** | Pointer to **string** | Unit of the query. | [optional] 
**TimeShiftLabel** | Pointer to **string** | Time shift value if specified in request in human readable format. | [optional] 

## Methods

### NewTimeSeriesList

`func NewTimeSeriesList(timeSeries []TimeSeries, ) *TimeSeriesList`

NewTimeSeriesList instantiates a new TimeSeriesList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesListWithDefaults

`func NewTimeSeriesListWithDefaults() *TimeSeriesList`

NewTimeSeriesListWithDefaults instantiates a new TimeSeriesList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeSeries

`func (o *TimeSeriesList) GetTimeSeries() []TimeSeries`

GetTimeSeries returns the TimeSeries field if non-nil, zero value otherwise.

### GetTimeSeriesOk

`func (o *TimeSeriesList) GetTimeSeriesOk() (*[]TimeSeries, bool)`

GetTimeSeriesOk returns a tuple with the TimeSeries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeries

`func (o *TimeSeriesList) SetTimeSeries(v []TimeSeries)`

SetTimeSeries sets TimeSeries field to given value.


### GetUnit

`func (o *TimeSeriesList) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *TimeSeriesList) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *TimeSeriesList) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *TimeSeriesList) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetTimeShiftLabel

`func (o *TimeSeriesList) GetTimeShiftLabel() string`

GetTimeShiftLabel returns the TimeShiftLabel field if non-nil, zero value otherwise.

### GetTimeShiftLabelOk

`func (o *TimeSeriesList) GetTimeShiftLabelOk() (*string, bool)`

GetTimeShiftLabelOk returns a tuple with the TimeShiftLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeShiftLabel

`func (o *TimeSeriesList) SetTimeShiftLabel(v string)`

SetTimeShiftLabel sets TimeShiftLabel field to given value.

### HasTimeShiftLabel

`func (o *TimeSeriesList) HasTimeShiftLabel() bool`

HasTimeShiftLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



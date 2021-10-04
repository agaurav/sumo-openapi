# TimeSeriesRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | Row id for the query row as specified in the request. | 
**TimeSeriesList** | [**TimeSeriesList**](TimeSeriesList.md) |  | 

## Methods

### NewTimeSeriesRow

`func NewTimeSeriesRow(rowId string, timeSeriesList TimeSeriesList, ) *TimeSeriesRow`

NewTimeSeriesRow instantiates a new TimeSeriesRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeSeriesRowWithDefaults

`func NewTimeSeriesRowWithDefaults() *TimeSeriesRow`

NewTimeSeriesRowWithDefaults instantiates a new TimeSeriesRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *TimeSeriesRow) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *TimeSeriesRow) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *TimeSeriesRow) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetTimeSeriesList

`func (o *TimeSeriesRow) GetTimeSeriesList() TimeSeriesList`

GetTimeSeriesList returns the TimeSeriesList field if non-nil, zero value otherwise.

### GetTimeSeriesListOk

`func (o *TimeSeriesRow) GetTimeSeriesListOk() (*TimeSeriesList, bool)`

GetTimeSeriesListOk returns a tuple with the TimeSeriesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeSeriesList

`func (o *TimeSeriesRow) SetTimeSeriesList(v TimeSeriesList)`

SetTimeSeriesList sets TimeSeriesList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



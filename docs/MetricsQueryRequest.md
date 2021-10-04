# MetricsQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]MetricsQueryRow**](MetricsQueryRow.md) | A list of metrics queries. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 

## Methods

### NewMetricsQueryRequest

`func NewMetricsQueryRequest(queries []MetricsQueryRow, timeRange ResolvableTimeRange, ) *MetricsQueryRequest`

NewMetricsQueryRequest instantiates a new MetricsQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryRequestWithDefaults

`func NewMetricsQueryRequestWithDefaults() *MetricsQueryRequest`

NewMetricsQueryRequestWithDefaults instantiates a new MetricsQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *MetricsQueryRequest) GetQueries() []MetricsQueryRow`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MetricsQueryRequest) GetQueriesOk() (*[]MetricsQueryRow, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MetricsQueryRequest) SetQueries(v []MetricsQueryRow)`

SetQueries sets Queries field to given value.


### GetTimeRange

`func (o *MetricsQueryRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsQueryRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsQueryRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



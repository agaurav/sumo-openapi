# MetricsQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryResult** | Pointer to [**[]TimeSeriesRow**](TimeSeriesRow.md) | A list of the time series returned by metric query. | [optional] 
**Errors** | [**ErrorResponse**](ErrorResponse.md) |  | 

## Methods

### NewMetricsQueryResponse

`func NewMetricsQueryResponse(errors ErrorResponse, ) *MetricsQueryResponse`

NewMetricsQueryResponse instantiates a new MetricsQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryResponseWithDefaults

`func NewMetricsQueryResponseWithDefaults() *MetricsQueryResponse`

NewMetricsQueryResponseWithDefaults instantiates a new MetricsQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryResult

`func (o *MetricsQueryResponse) GetQueryResult() []TimeSeriesRow`

GetQueryResult returns the QueryResult field if non-nil, zero value otherwise.

### GetQueryResultOk

`func (o *MetricsQueryResponse) GetQueryResultOk() (*[]TimeSeriesRow, bool)`

GetQueryResultOk returns a tuple with the QueryResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryResult

`func (o *MetricsQueryResponse) SetQueryResult(v []TimeSeriesRow)`

SetQueryResult sets QueryResult field to given value.

### HasQueryResult

`func (o *MetricsQueryResponse) HasQueryResult() bool`

HasQueryResult returns a boolean if a field has been set.

### GetErrors

`func (o *MetricsQueryResponse) GetErrors() ErrorResponse`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *MetricsQueryResponse) GetErrorsOk() (*ErrorResponse, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *MetricsQueryResponse) SetErrors(v ErrorResponse)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# MetricsQueryRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | Row id for the query row, A to Z letter. | 
**Query** | **string** | A metric query consists of a metric, one or more filters and optionally, one or more [Metrics Operators](https://help.sumologic.com/?cid&#x3D;10144). Strictly speaking, both filters and operators are optional.  Most of the [Metrics Operators](https://help.sumologic.com/?cid&#x3D;10144) are allowed in the query string except &#x60;fillmissing&#x60;, &#x60;outlier&#x60;, &#x60;quantize&#x60; and &#x60;timeshift&#x60;.    * &#x60;fillmissing&#x60;: Not supported in API.   * &#x60;outlier&#x60;: Not supported in API.   * &#x60;quantize&#x60;: Only supported through &#x60;quantization&#x60; param.   * &#x60;timeshift&#x60;: Only supported through &#x60;timeshift&#x60; param.   In practice, your metric queries will almost always contain filters that narrow the scope of your query. For more information about the query language see [Metrics Queries](http://help.sumologic.com/?cid&#x3D;1079). | 
**Quantization** | Pointer to **int64** | Segregates time series data by time period. This allows you to create aggregated results in buckets of fixed intervals (for example, 5-minute intervals). The value is in milliseconds. | [optional] 
**Rollup** | Pointer to **string** | We use the term rollup to refer to the aggregation function Sumo Logic uses when quantizing metrics. Can be &#x60;Avg&#x60;, &#x60;Sum&#x60;, &#x60;Min&#x60;, &#x60;Max&#x60;, &#x60;Count&#x60; or &#x60;None&#x60;. | [optional] 
**Timeshift** | Pointer to **int64** | Shifts the time series from your metrics query by the specified amount of time. This can help when comparing a time series across multiple time periods. Specified as a signed duration in milliseconds. | [optional] 
**Transient** | Pointer to **bool** | Determines if the row should be returned in the response. Can be used in conjunction with a join, if only the result of the join is needed, and not the intermediate rows. | [optional] [default to false]

## Methods

### NewMetricsQueryRow

`func NewMetricsQueryRow(rowId string, query string, ) *MetricsQueryRow`

NewMetricsQueryRow instantiates a new MetricsQueryRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQueryRowWithDefaults

`func NewMetricsQueryRowWithDefaults() *MetricsQueryRow`

NewMetricsQueryRowWithDefaults instantiates a new MetricsQueryRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *MetricsQueryRow) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *MetricsQueryRow) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *MetricsQueryRow) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetQuery

`func (o *MetricsQueryRow) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricsQueryRow) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricsQueryRow) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetQuantization

`func (o *MetricsQueryRow) GetQuantization() int64`

GetQuantization returns the Quantization field if non-nil, zero value otherwise.

### GetQuantizationOk

`func (o *MetricsQueryRow) GetQuantizationOk() (*int64, bool)`

GetQuantizationOk returns a tuple with the Quantization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantization

`func (o *MetricsQueryRow) SetQuantization(v int64)`

SetQuantization sets Quantization field to given value.

### HasQuantization

`func (o *MetricsQueryRow) HasQuantization() bool`

HasQuantization returns a boolean if a field has been set.

### GetRollup

`func (o *MetricsQueryRow) GetRollup() string`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *MetricsQueryRow) GetRollupOk() (*string, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *MetricsQueryRow) SetRollup(v string)`

SetRollup sets Rollup field to given value.

### HasRollup

`func (o *MetricsQueryRow) HasRollup() bool`

HasRollup returns a boolean if a field has been set.

### GetTimeshift

`func (o *MetricsQueryRow) GetTimeshift() int64`

GetTimeshift returns the Timeshift field if non-nil, zero value otherwise.

### GetTimeshiftOk

`func (o *MetricsQueryRow) GetTimeshiftOk() (*int64, bool)`

GetTimeshiftOk returns a tuple with the Timeshift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeshift

`func (o *MetricsQueryRow) SetTimeshift(v int64)`

SetTimeshift sets Timeshift field to given value.

### HasTimeshift

`func (o *MetricsQueryRow) HasTimeshift() bool`

HasTimeshift returns a boolean if a field has been set.

### GetTransient

`func (o *MetricsQueryRow) GetTransient() bool`

GetTransient returns the Transient field if non-nil, zero value otherwise.

### GetTransientOk

`func (o *MetricsQueryRow) GetTransientOk() (*bool, bool)`

GetTransientOk returns a tuple with the Transient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransient

`func (o *MetricsQueryRow) SetTransient(v bool)`

SetTransient sets Transient field to given value.

### HasTransient

`func (o *MetricsQueryRow) HasTransient() bool`

HasTransient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



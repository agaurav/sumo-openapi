# MetricsSearchQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | Row identifier. All row IDs are represented by subsequent upper case letters starting with &#x60;A&#x60;. | 
**Query** | **string** | Metrics query. | 

## Methods

### NewMetricsSearchQuery

`func NewMetricsSearchQuery(rowId string, query string, ) *MetricsSearchQuery`

NewMetricsSearchQuery instantiates a new MetricsSearchQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSearchQueryWithDefaults

`func NewMetricsSearchQueryWithDefaults() *MetricsSearchQuery`

NewMetricsSearchQueryWithDefaults instantiates a new MetricsSearchQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *MetricsSearchQuery) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *MetricsSearchQuery) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *MetricsSearchQuery) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetQuery

`func (o *MetricsSearchQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricsSearchQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricsSearchQuery) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



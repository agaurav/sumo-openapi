# MetricsQuerySyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The text of a metrics query. | 
**RowId** | **string** | A label referring to the query; used if other metrics queries reference this one. | 

## Methods

### NewMetricsQuerySyncDefinition

`func NewMetricsQuerySyncDefinition(query string, rowId string, ) *MetricsQuerySyncDefinition`

NewMetricsQuerySyncDefinition instantiates a new MetricsQuerySyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsQuerySyncDefinitionWithDefaults

`func NewMetricsQuerySyncDefinitionWithDefaults() *MetricsQuerySyncDefinition`

NewMetricsQuerySyncDefinitionWithDefaults instantiates a new MetricsQuerySyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *MetricsQuerySyncDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MetricsQuerySyncDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MetricsQuerySyncDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetRowId

`func (o *MetricsQuerySyncDefinition) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *MetricsQuerySyncDefinition) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *MetricsQuerySyncDefinition) SetRowId(v string)`

SetRowId sets RowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



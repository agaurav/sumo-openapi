# MonitorQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | The unique identifier of the row. Defaults to sequential capital letters, &#x60;A&#x60;, &#x60;B&#x60;, &#x60;C&#x60;, etc. | 
**Query** | **string** | The logs or metrics query that defines the stream of data the monitor runs on. | 

## Methods

### NewMonitorQuery

`func NewMonitorQuery(rowId string, query string, ) *MonitorQuery`

NewMonitorQuery instantiates a new MonitorQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorQueryWithDefaults

`func NewMonitorQueryWithDefaults() *MonitorQuery`

NewMonitorQueryWithDefaults instantiates a new MonitorQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *MonitorQuery) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *MonitorQuery) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *MonitorQuery) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetQuery

`func (o *MonitorQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *MonitorQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *MonitorQuery) SetQuery(v string)`

SetQuery sets Query field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



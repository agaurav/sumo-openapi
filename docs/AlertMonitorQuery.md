# AlertMonitorQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RowId** | **string** | The unique identifier of the row. Defaults to sequential capital letters, &#x60;A&#x60;, &#x60;B&#x60;, &#x60;C&#x60;, etc. | 
**Query** | **string** | The logs or metrics query that defines the stream of data the monitor runs on. | 
**IsTriggerRow** | **bool** | Indicates whether the current row is the trigger (final) row. | 

## Methods

### NewAlertMonitorQuery

`func NewAlertMonitorQuery(rowId string, query string, isTriggerRow bool, ) *AlertMonitorQuery`

NewAlertMonitorQuery instantiates a new AlertMonitorQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertMonitorQueryWithDefaults

`func NewAlertMonitorQueryWithDefaults() *AlertMonitorQuery`

NewAlertMonitorQueryWithDefaults instantiates a new AlertMonitorQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRowId

`func (o *AlertMonitorQuery) GetRowId() string`

GetRowId returns the RowId field if non-nil, zero value otherwise.

### GetRowIdOk

`func (o *AlertMonitorQuery) GetRowIdOk() (*string, bool)`

GetRowIdOk returns a tuple with the RowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowId

`func (o *AlertMonitorQuery) SetRowId(v string)`

SetRowId sets RowId field to given value.


### GetQuery

`func (o *AlertMonitorQuery) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AlertMonitorQuery) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AlertMonitorQuery) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetIsTriggerRow

`func (o *AlertMonitorQuery) GetIsTriggerRow() bool`

GetIsTriggerRow returns the IsTriggerRow field if non-nil, zero value otherwise.

### GetIsTriggerRowOk

`func (o *AlertMonitorQuery) GetIsTriggerRowOk() (*bool, bool)`

GetIsTriggerRowOk returns a tuple with the IsTriggerRow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTriggerRow

`func (o *AlertMonitorQuery) SetIsTriggerRow(v bool)`

SetIsTriggerRow sets IsTriggerRow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



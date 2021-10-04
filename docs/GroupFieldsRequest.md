# GroupFieldsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Queries** | [**[]MonitorQuery**](MonitorQuery.md) | All queries from the monitor. | 
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor. | 

## Methods

### NewGroupFieldsRequest

`func NewGroupFieldsRequest(queries []MonitorQuery, monitorType string, ) *GroupFieldsRequest`

NewGroupFieldsRequest instantiates a new GroupFieldsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupFieldsRequestWithDefaults

`func NewGroupFieldsRequestWithDefaults() *GroupFieldsRequest`

NewGroupFieldsRequestWithDefaults instantiates a new GroupFieldsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueries

`func (o *GroupFieldsRequest) GetQueries() []MonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *GroupFieldsRequest) GetQueriesOk() (*[]MonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *GroupFieldsRequest) SetQueries(v []MonitorQuery)`

SetQueries sets Queries field to given value.


### GetMonitorType

`func (o *GroupFieldsRequest) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *GroupFieldsRequest) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *GroupFieldsRequest) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



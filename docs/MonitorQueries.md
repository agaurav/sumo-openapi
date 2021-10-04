# MonitorQueries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | **string** | The type of monitor. Valid values:   1. &#x60;Logs&#x60;: A logs query monitor.   2. &#x60;Metrics&#x60;: A metrics query monitor. | 
**TimeRange** | **string** | The relative time range of the monitor. | 
**Queries** | [**[]UnvalidatedMonitorQuery**](UnvalidatedMonitorQuery.md) | Queries to be validated. | 

## Methods

### NewMonitorQueries

`func NewMonitorQueries(monitorType string, timeRange string, queries []UnvalidatedMonitorQuery, ) *MonitorQueries`

NewMonitorQueries instantiates a new MonitorQueries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorQueriesWithDefaults

`func NewMonitorQueriesWithDefaults() *MonitorQueries`

NewMonitorQueriesWithDefaults instantiates a new MonitorQueries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorQueries) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorQueries) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorQueries) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.


### GetTimeRange

`func (o *MonitorQueries) GetTimeRange() string`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MonitorQueries) GetTimeRangeOk() (*string, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MonitorQueries) SetTimeRange(v string)`

SetTimeRange sets TimeRange field to given value.


### GetQueries

`func (o *MonitorQueries) GetQueries() []UnvalidatedMonitorQuery`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MonitorQueries) GetQueriesOk() (*[]UnvalidatedMonitorQuery, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MonitorQueries) SetQueries(v []UnvalidatedMonitorQuery)`

SetQueries sets Queries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



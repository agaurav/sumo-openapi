# MonitorUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MonitorType** | Pointer to **string** | The type of monitor usage info (Logs or Metrics). | [optional] 
**Usage** | Pointer to **int32** | Current number of active Logs/Metrics monitors. | [optional] 
**Limit** | Pointer to **int32** | The limit of active Logs/Metrics monitors. | [optional] 
**Total** | Pointer to **int32** | The total number of monitors created. (Including both active and disabled Logs/Metrics monitors) | [optional] 

## Methods

### NewMonitorUsage

`func NewMonitorUsage() *MonitorUsage`

NewMonitorUsage instantiates a new MonitorUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorUsageWithDefaults

`func NewMonitorUsageWithDefaults() *MonitorUsage`

NewMonitorUsageWithDefaults instantiates a new MonitorUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMonitorType

`func (o *MonitorUsage) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *MonitorUsage) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *MonitorUsage) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *MonitorUsage) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetUsage

`func (o *MonitorUsage) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MonitorUsage) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MonitorUsage) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *MonitorUsage) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetLimit

`func (o *MonitorUsage) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MonitorUsage) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MonitorUsage) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MonitorUsage) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTotal

`func (o *MonitorUsage) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MonitorUsage) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MonitorUsage) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *MonitorUsage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



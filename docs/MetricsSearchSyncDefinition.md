# MetricsSearchSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**Description** | Pointer to **string** | Description of the metrics search page. | [optional] 
**Queries** | [**[]Query**](Query.md) | Queries of the metrics search page. | 
**VisualSettings** | Pointer to **string** | Visual settings of the metrics search page. Must be a string representing a valid JSON object.  | [optional] 

## Methods

### NewMetricsSearchSyncDefinition

`func NewMetricsSearchSyncDefinition(timeRange ResolvableTimeRange, queries []Query, ) *MetricsSearchSyncDefinition`

NewMetricsSearchSyncDefinition instantiates a new MetricsSearchSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSearchSyncDefinitionWithDefaults

`func NewMetricsSearchSyncDefinitionWithDefaults() *MetricsSearchSyncDefinition`

NewMetricsSearchSyncDefinitionWithDefaults instantiates a new MetricsSearchSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeRange

`func (o *MetricsSearchSyncDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsSearchSyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsSearchSyncDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetDescription

`func (o *MetricsSearchSyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsSearchSyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsSearchSyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricsSearchSyncDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQueries

`func (o *MetricsSearchSyncDefinition) GetQueries() []Query`

GetQueries returns the Queries field if non-nil, zero value otherwise.

### GetQueriesOk

`func (o *MetricsSearchSyncDefinition) GetQueriesOk() (*[]Query, bool)`

GetQueriesOk returns a tuple with the Queries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueries

`func (o *MetricsSearchSyncDefinition) SetQueries(v []Query)`

SetQueries sets Queries field to given value.


### GetVisualSettings

`func (o *MetricsSearchSyncDefinition) GetVisualSettings() string`

GetVisualSettings returns the VisualSettings field if non-nil, zero value otherwise.

### GetVisualSettingsOk

`func (o *MetricsSearchSyncDefinition) GetVisualSettingsOk() (*string, bool)`

GetVisualSettingsOk returns a tuple with the VisualSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualSettings

`func (o *MetricsSearchSyncDefinition) SetVisualSettings(v string)`

SetVisualSettings sets VisualSettings field to given value.

### HasVisualSettings

`func (o *MetricsSearchSyncDefinition) HasVisualSettings() bool`

HasVisualSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



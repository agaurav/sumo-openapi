# SaveMetricsSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Item title in the content library. | 
**Description** | **string** | Item description in the content library. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**LogQuery** | Pointer to **string** | Log query used to add an overlay to the chart. | [optional] 
**MetricsQueries** | [**[]MetricsSearchQuery**](MetricsSearchQuery.md) | Metrics queries, up to the maximum of six. | 
**DesiredQuantizationInSecs** | Pointer to **int32** | Desired quantization in seconds. | [optional] [default to 0]
**Properties** | Pointer to **string** | Chart properties, like line width, color palette, and the fill missing data method. Leave this field empty to use the defaults. This property contains JSON object encoded as a string.  | [optional] 
**ParentId** | **string** | Identifier of a folder to which the metrics search should be added. | 

## Methods

### NewSaveMetricsSearchRequest

`func NewSaveMetricsSearchRequest(title string, description string, timeRange ResolvableTimeRange, metricsQueries []MetricsSearchQuery, parentId string, ) *SaveMetricsSearchRequest`

NewSaveMetricsSearchRequest instantiates a new SaveMetricsSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveMetricsSearchRequestWithDefaults

`func NewSaveMetricsSearchRequestWithDefaults() *SaveMetricsSearchRequest`

NewSaveMetricsSearchRequestWithDefaults instantiates a new SaveMetricsSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SaveMetricsSearchRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SaveMetricsSearchRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SaveMetricsSearchRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *SaveMetricsSearchRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SaveMetricsSearchRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SaveMetricsSearchRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTimeRange

`func (o *SaveMetricsSearchRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *SaveMetricsSearchRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *SaveMetricsSearchRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetLogQuery

`func (o *SaveMetricsSearchRequest) GetLogQuery() string`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *SaveMetricsSearchRequest) GetLogQueryOk() (*string, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *SaveMetricsSearchRequest) SetLogQuery(v string)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *SaveMetricsSearchRequest) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetMetricsQueries

`func (o *SaveMetricsSearchRequest) GetMetricsQueries() []MetricsSearchQuery`

GetMetricsQueries returns the MetricsQueries field if non-nil, zero value otherwise.

### GetMetricsQueriesOk

`func (o *SaveMetricsSearchRequest) GetMetricsQueriesOk() (*[]MetricsSearchQuery, bool)`

GetMetricsQueriesOk returns a tuple with the MetricsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsQueries

`func (o *SaveMetricsSearchRequest) SetMetricsQueries(v []MetricsSearchQuery)`

SetMetricsQueries sets MetricsQueries field to given value.


### GetDesiredQuantizationInSecs

`func (o *SaveMetricsSearchRequest) GetDesiredQuantizationInSecs() int32`

GetDesiredQuantizationInSecs returns the DesiredQuantizationInSecs field if non-nil, zero value otherwise.

### GetDesiredQuantizationInSecsOk

`func (o *SaveMetricsSearchRequest) GetDesiredQuantizationInSecsOk() (*int32, bool)`

GetDesiredQuantizationInSecsOk returns a tuple with the DesiredQuantizationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQuantizationInSecs

`func (o *SaveMetricsSearchRequest) SetDesiredQuantizationInSecs(v int32)`

SetDesiredQuantizationInSecs sets DesiredQuantizationInSecs field to given value.

### HasDesiredQuantizationInSecs

`func (o *SaveMetricsSearchRequest) HasDesiredQuantizationInSecs() bool`

HasDesiredQuantizationInSecs returns a boolean if a field has been set.

### GetProperties

`func (o *SaveMetricsSearchRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SaveMetricsSearchRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SaveMetricsSearchRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *SaveMetricsSearchRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetParentId

`func (o *SaveMetricsSearchRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *SaveMetricsSearchRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *SaveMetricsSearchRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



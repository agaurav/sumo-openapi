# MetricsSearchInstance

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
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | Identifier of the metrics search. | 
**ParentId** | Pointer to **string** | Identifier of the parent element in the content library, such as folder. | [optional] 

## Methods

### NewMetricsSearchInstance

`func NewMetricsSearchInstance(title string, description string, timeRange ResolvableTimeRange, metricsQueries []MetricsSearchQuery, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, ) *MetricsSearchInstance`

NewMetricsSearchInstance instantiates a new MetricsSearchInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsSearchInstanceWithDefaults

`func NewMetricsSearchInstanceWithDefaults() *MetricsSearchInstance`

NewMetricsSearchInstanceWithDefaults instantiates a new MetricsSearchInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *MetricsSearchInstance) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MetricsSearchInstance) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MetricsSearchInstance) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *MetricsSearchInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricsSearchInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricsSearchInstance) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetTimeRange

`func (o *MetricsSearchInstance) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MetricsSearchInstance) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MetricsSearchInstance) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetLogQuery

`func (o *MetricsSearchInstance) GetLogQuery() string`

GetLogQuery returns the LogQuery field if non-nil, zero value otherwise.

### GetLogQueryOk

`func (o *MetricsSearchInstance) GetLogQueryOk() (*string, bool)`

GetLogQueryOk returns a tuple with the LogQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogQuery

`func (o *MetricsSearchInstance) SetLogQuery(v string)`

SetLogQuery sets LogQuery field to given value.

### HasLogQuery

`func (o *MetricsSearchInstance) HasLogQuery() bool`

HasLogQuery returns a boolean if a field has been set.

### GetMetricsQueries

`func (o *MetricsSearchInstance) GetMetricsQueries() []MetricsSearchQuery`

GetMetricsQueries returns the MetricsQueries field if non-nil, zero value otherwise.

### GetMetricsQueriesOk

`func (o *MetricsSearchInstance) GetMetricsQueriesOk() (*[]MetricsSearchQuery, bool)`

GetMetricsQueriesOk returns a tuple with the MetricsQueries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricsQueries

`func (o *MetricsSearchInstance) SetMetricsQueries(v []MetricsSearchQuery)`

SetMetricsQueries sets MetricsQueries field to given value.


### GetDesiredQuantizationInSecs

`func (o *MetricsSearchInstance) GetDesiredQuantizationInSecs() int32`

GetDesiredQuantizationInSecs returns the DesiredQuantizationInSecs field if non-nil, zero value otherwise.

### GetDesiredQuantizationInSecsOk

`func (o *MetricsSearchInstance) GetDesiredQuantizationInSecsOk() (*int32, bool)`

GetDesiredQuantizationInSecsOk returns a tuple with the DesiredQuantizationInSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredQuantizationInSecs

`func (o *MetricsSearchInstance) SetDesiredQuantizationInSecs(v int32)`

SetDesiredQuantizationInSecs sets DesiredQuantizationInSecs field to given value.

### HasDesiredQuantizationInSecs

`func (o *MetricsSearchInstance) HasDesiredQuantizationInSecs() bool`

HasDesiredQuantizationInSecs returns a boolean if a field has been set.

### GetProperties

`func (o *MetricsSearchInstance) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *MetricsSearchInstance) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *MetricsSearchInstance) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *MetricsSearchInstance) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MetricsSearchInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MetricsSearchInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MetricsSearchInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *MetricsSearchInstance) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *MetricsSearchInstance) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *MetricsSearchInstance) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *MetricsSearchInstance) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *MetricsSearchInstance) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *MetricsSearchInstance) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *MetricsSearchInstance) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *MetricsSearchInstance) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *MetricsSearchInstance) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *MetricsSearchInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MetricsSearchInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MetricsSearchInstance) SetId(v string)`

SetId sets Id field to given value.


### GetParentId

`func (o *MetricsSearchInstance) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *MetricsSearchInstance) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *MetricsSearchInstance) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *MetricsSearchInstance) HasParentId() bool`

HasParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LogSearchEstimatedUsageByTierDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | Query to perform. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**RunByReceiptTime** | Pointer to **bool** | This has the value &#x60;true&#x60; if the search is to be run by receipt time and &#x60;false&#x60; if it is to be run by message time. | [optional] [default to false]
**QueryParameters** | Pointer to [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | Definition of the query parameters. | [optional] 
**Timezone** | **string** | Time zone to get the estimated usage details. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).  | 
**EstimatedUsageDetails** | [**[]EstimatedUsageDetailsWithTier**](EstimatedUsageDetailsWithTier.md) |  | 

## Methods

### NewLogSearchEstimatedUsageByTierDefinition

`func NewLogSearchEstimatedUsageByTierDefinition(queryString string, timeRange ResolvableTimeRange, timezone string, estimatedUsageDetails []EstimatedUsageDetailsWithTier, ) *LogSearchEstimatedUsageByTierDefinition`

NewLogSearchEstimatedUsageByTierDefinition instantiates a new LogSearchEstimatedUsageByTierDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchEstimatedUsageByTierDefinitionWithDefaults

`func NewLogSearchEstimatedUsageByTierDefinitionWithDefaults() *LogSearchEstimatedUsageByTierDefinition`

NewLogSearchEstimatedUsageByTierDefinitionWithDefaults instantiates a new LogSearchEstimatedUsageByTierDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *LogSearchEstimatedUsageByTierDefinition) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogSearchEstimatedUsageByTierDefinition) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogSearchEstimatedUsageByTierDefinition) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetTimeRange

`func (o *LogSearchEstimatedUsageByTierDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *LogSearchEstimatedUsageByTierDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *LogSearchEstimatedUsageByTierDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetRunByReceiptTime

`func (o *LogSearchEstimatedUsageByTierDefinition) GetRunByReceiptTime() bool`

GetRunByReceiptTime returns the RunByReceiptTime field if non-nil, zero value otherwise.

### GetRunByReceiptTimeOk

`func (o *LogSearchEstimatedUsageByTierDefinition) GetRunByReceiptTimeOk() (*bool, bool)`

GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByReceiptTime

`func (o *LogSearchEstimatedUsageByTierDefinition) SetRunByReceiptTime(v bool)`

SetRunByReceiptTime sets RunByReceiptTime field to given value.

### HasRunByReceiptTime

`func (o *LogSearchEstimatedUsageByTierDefinition) HasRunByReceiptTime() bool`

HasRunByReceiptTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *LogSearchEstimatedUsageByTierDefinition) GetQueryParameters() []QueryParameterSyncDefinition`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *LogSearchEstimatedUsageByTierDefinition) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *LogSearchEstimatedUsageByTierDefinition) SetQueryParameters(v []QueryParameterSyncDefinition)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *LogSearchEstimatedUsageByTierDefinition) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.

### GetTimezone

`func (o *LogSearchEstimatedUsageByTierDefinition) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *LogSearchEstimatedUsageByTierDefinition) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *LogSearchEstimatedUsageByTierDefinition) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.


### GetEstimatedUsageDetails

`func (o *LogSearchEstimatedUsageByTierDefinition) GetEstimatedUsageDetails() []EstimatedUsageDetailsWithTier`

GetEstimatedUsageDetails returns the EstimatedUsageDetails field if non-nil, zero value otherwise.

### GetEstimatedUsageDetailsOk

`func (o *LogSearchEstimatedUsageByTierDefinition) GetEstimatedUsageDetailsOk() (*[]EstimatedUsageDetailsWithTier, bool)`

GetEstimatedUsageDetailsOk returns a tuple with the EstimatedUsageDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedUsageDetails

`func (o *LogSearchEstimatedUsageByTierDefinition) SetEstimatedUsageDetails(v []EstimatedUsageDetailsWithTier)`

SetEstimatedUsageDetails sets EstimatedUsageDetails field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



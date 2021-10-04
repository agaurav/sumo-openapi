# LogSearchQueryTimeRangeBaseExceptParsingMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | **string** | Query to perform. | 
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**RunByReceiptTime** | Pointer to **bool** | This has the value &#x60;true&#x60; if the search is to be run by receipt time and &#x60;false&#x60; if it is to be run by message time. | [optional] [default to false]
**QueryParameters** | Pointer to [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | Definition of the query parameters. | [optional] 

## Methods

### NewLogSearchQueryTimeRangeBaseExceptParsingMode

`func NewLogSearchQueryTimeRangeBaseExceptParsingMode(queryString string, timeRange ResolvableTimeRange, ) *LogSearchQueryTimeRangeBaseExceptParsingMode`

NewLogSearchQueryTimeRangeBaseExceptParsingMode instantiates a new LogSearchQueryTimeRangeBaseExceptParsingMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchQueryTimeRangeBaseExceptParsingModeWithDefaults

`func NewLogSearchQueryTimeRangeBaseExceptParsingModeWithDefaults() *LogSearchQueryTimeRangeBaseExceptParsingMode`

NewLogSearchQueryTimeRangeBaseExceptParsingModeWithDefaults instantiates a new LogSearchQueryTimeRangeBaseExceptParsingMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.


### GetTimeRange

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetRunByReceiptTime

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetRunByReceiptTime() bool`

GetRunByReceiptTime returns the RunByReceiptTime field if non-nil, zero value otherwise.

### GetRunByReceiptTimeOk

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetRunByReceiptTimeOk() (*bool, bool)`

GetRunByReceiptTimeOk returns a tuple with the RunByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunByReceiptTime

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetRunByReceiptTime(v bool)`

SetRunByReceiptTime sets RunByReceiptTime field to given value.

### HasRunByReceiptTime

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) HasRunByReceiptTime() bool`

HasRunByReceiptTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryParameters() []QueryParameterSyncDefinition`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) SetQueryParameters(v []QueryParameterSyncDefinition)`

SetQueryParameters sets QueryParameters field to given value.

### HasQueryParameters

`func (o *LogSearchQueryTimeRangeBaseExceptParsingMode) HasQueryParameters() bool`

HasQueryParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



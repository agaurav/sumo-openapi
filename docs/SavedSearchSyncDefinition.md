# SavedSearchSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryText** | **string** | The text of a Sumo Logic query. | 
**DefaultTimeRange** | **string** | Default time range for the search. Possible types of time ranges are:   - relative time range: e.g. \&quot;-1d -12h\&quot; represents a time range from one day ago to 12 hours ago.   - absolute time range: e.g. \&quot;01-04-2017 20:32:00 to 01-04-2017 20:35:00\&quot; represents a time range     from April 1st, 2017 at 8:32 PM until April 1st, 2017 at 8:35 PM. | 
**ByReceiptTime** | **bool** | Set it to true to run the search using receipt time. By default, searches do not run by receipt time. | [default to false]
**ViewName** | Pointer to **string** | The name of the Scheduled View that has indexed the data you want to search. | [optional] 
**ViewStartTime** | Pointer to **time.Time** | Start timestamp of the Scheduled View in UTC format. | [optional] 
**QueryParameters** | [**[]QueryParameterSyncDefinition**](QueryParameterSyncDefinition.md) | An array of search query parameter objects. | 
**ParsingMode** | Pointer to **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to "Manual"]

## Methods

### NewSavedSearchSyncDefinition

`func NewSavedSearchSyncDefinition(queryText string, defaultTimeRange string, byReceiptTime bool, queryParameters []QueryParameterSyncDefinition, ) *SavedSearchSyncDefinition`

NewSavedSearchSyncDefinition instantiates a new SavedSearchSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSavedSearchSyncDefinitionWithDefaults

`func NewSavedSearchSyncDefinitionWithDefaults() *SavedSearchSyncDefinition`

NewSavedSearchSyncDefinitionWithDefaults instantiates a new SavedSearchSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryText

`func (o *SavedSearchSyncDefinition) GetQueryText() string`

GetQueryText returns the QueryText field if non-nil, zero value otherwise.

### GetQueryTextOk

`func (o *SavedSearchSyncDefinition) GetQueryTextOk() (*string, bool)`

GetQueryTextOk returns a tuple with the QueryText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryText

`func (o *SavedSearchSyncDefinition) SetQueryText(v string)`

SetQueryText sets QueryText field to given value.


### GetDefaultTimeRange

`func (o *SavedSearchSyncDefinition) GetDefaultTimeRange() string`

GetDefaultTimeRange returns the DefaultTimeRange field if non-nil, zero value otherwise.

### GetDefaultTimeRangeOk

`func (o *SavedSearchSyncDefinition) GetDefaultTimeRangeOk() (*string, bool)`

GetDefaultTimeRangeOk returns a tuple with the DefaultTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTimeRange

`func (o *SavedSearchSyncDefinition) SetDefaultTimeRange(v string)`

SetDefaultTimeRange sets DefaultTimeRange field to given value.


### GetByReceiptTime

`func (o *SavedSearchSyncDefinition) GetByReceiptTime() bool`

GetByReceiptTime returns the ByReceiptTime field if non-nil, zero value otherwise.

### GetByReceiptTimeOk

`func (o *SavedSearchSyncDefinition) GetByReceiptTimeOk() (*bool, bool)`

GetByReceiptTimeOk returns a tuple with the ByReceiptTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByReceiptTime

`func (o *SavedSearchSyncDefinition) SetByReceiptTime(v bool)`

SetByReceiptTime sets ByReceiptTime field to given value.


### GetViewName

`func (o *SavedSearchSyncDefinition) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *SavedSearchSyncDefinition) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *SavedSearchSyncDefinition) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *SavedSearchSyncDefinition) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetViewStartTime

`func (o *SavedSearchSyncDefinition) GetViewStartTime() time.Time`

GetViewStartTime returns the ViewStartTime field if non-nil, zero value otherwise.

### GetViewStartTimeOk

`func (o *SavedSearchSyncDefinition) GetViewStartTimeOk() (*time.Time, bool)`

GetViewStartTimeOk returns a tuple with the ViewStartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStartTime

`func (o *SavedSearchSyncDefinition) SetViewStartTime(v time.Time)`

SetViewStartTime sets ViewStartTime field to given value.

### HasViewStartTime

`func (o *SavedSearchSyncDefinition) HasViewStartTime() bool`

HasViewStartTime returns a boolean if a field has been set.

### GetQueryParameters

`func (o *SavedSearchSyncDefinition) GetQueryParameters() []QueryParameterSyncDefinition`

GetQueryParameters returns the QueryParameters field if non-nil, zero value otherwise.

### GetQueryParametersOk

`func (o *SavedSearchSyncDefinition) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool)`

GetQueryParametersOk returns a tuple with the QueryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryParameters

`func (o *SavedSearchSyncDefinition) SetQueryParameters(v []QueryParameterSyncDefinition)`

SetQueryParameters sets QueryParameters field to given value.


### GetParsingMode

`func (o *SavedSearchSyncDefinition) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *SavedSearchSyncDefinition) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *SavedSearchSyncDefinition) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *SavedSearchSyncDefinition) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



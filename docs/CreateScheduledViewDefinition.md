# CreateScheduledViewDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | **string** | The query that defines the data to be included in the scheduled view. | 
**IndexName** | **string** | Name of the index for the scheduled view. | 
**StartTime** | **time.Time** | Start timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**RetentionPeriod** | Pointer to **int32** | The number of days to retain data in the scheduled view, or -1 to use the default value for your account.  Only relevant if your account has multi-retention. enabled. | [optional] [default to -1]
**DataForwardingId** | Pointer to **string** | An optional ID of a data forwarding configuration to be used by the scheduled view. | [optional] 
**ParsingMode** | Pointer to **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to "Manual"]

## Methods

### NewCreateScheduledViewDefinition

`func NewCreateScheduledViewDefinition(query string, indexName string, startTime time.Time, ) *CreateScheduledViewDefinition`

NewCreateScheduledViewDefinition instantiates a new CreateScheduledViewDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateScheduledViewDefinitionWithDefaults

`func NewCreateScheduledViewDefinitionWithDefaults() *CreateScheduledViewDefinition`

NewCreateScheduledViewDefinitionWithDefaults instantiates a new CreateScheduledViewDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CreateScheduledViewDefinition) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CreateScheduledViewDefinition) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CreateScheduledViewDefinition) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetIndexName

`func (o *CreateScheduledViewDefinition) GetIndexName() string`

GetIndexName returns the IndexName field if non-nil, zero value otherwise.

### GetIndexNameOk

`func (o *CreateScheduledViewDefinition) GetIndexNameOk() (*string, bool)`

GetIndexNameOk returns a tuple with the IndexName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexName

`func (o *CreateScheduledViewDefinition) SetIndexName(v string)`

SetIndexName sets IndexName field to given value.


### GetStartTime

`func (o *CreateScheduledViewDefinition) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CreateScheduledViewDefinition) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CreateScheduledViewDefinition) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetRetentionPeriod

`func (o *CreateScheduledViewDefinition) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *CreateScheduledViewDefinition) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *CreateScheduledViewDefinition) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *CreateScheduledViewDefinition) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetDataForwardingId

`func (o *CreateScheduledViewDefinition) GetDataForwardingId() string`

GetDataForwardingId returns the DataForwardingId field if non-nil, zero value otherwise.

### GetDataForwardingIdOk

`func (o *CreateScheduledViewDefinition) GetDataForwardingIdOk() (*string, bool)`

GetDataForwardingIdOk returns a tuple with the DataForwardingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwardingId

`func (o *CreateScheduledViewDefinition) SetDataForwardingId(v string)`

SetDataForwardingId sets DataForwardingId field to given value.

### HasDataForwardingId

`func (o *CreateScheduledViewDefinition) HasDataForwardingId() bool`

HasDataForwardingId returns a boolean if a field has been set.

### GetParsingMode

`func (o *CreateScheduledViewDefinition) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *CreateScheduledViewDefinition) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *CreateScheduledViewDefinition) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *CreateScheduledViewDefinition) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LogSearchQueryTimeRangeBaseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParsingMode** | Pointer to **string** | Define the parsing mode to scan the JSON format log messages. Possible values are:   1. &#x60;AutoParse&#x60;   2. &#x60;Manual&#x60; In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid&#x3D;0011). | [optional] [default to "Manual"]

## Methods

### NewLogSearchQueryTimeRangeBaseAllOf

`func NewLogSearchQueryTimeRangeBaseAllOf() *LogSearchQueryTimeRangeBaseAllOf`

NewLogSearchQueryTimeRangeBaseAllOf instantiates a new LogSearchQueryTimeRangeBaseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogSearchQueryTimeRangeBaseAllOfWithDefaults

`func NewLogSearchQueryTimeRangeBaseAllOfWithDefaults() *LogSearchQueryTimeRangeBaseAllOf`

NewLogSearchQueryTimeRangeBaseAllOfWithDefaults instantiates a new LogSearchQueryTimeRangeBaseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParsingMode

`func (o *LogSearchQueryTimeRangeBaseAllOf) GetParsingMode() string`

GetParsingMode returns the ParsingMode field if non-nil, zero value otherwise.

### GetParsingModeOk

`func (o *LogSearchQueryTimeRangeBaseAllOf) GetParsingModeOk() (*string, bool)`

GetParsingModeOk returns a tuple with the ParsingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParsingMode

`func (o *LogSearchQueryTimeRangeBaseAllOf) SetParsingMode(v string)`

SetParsingMode sets ParsingMode field to given value.

### HasParsingMode

`func (o *LogSearchQueryTimeRangeBaseAllOf) HasParsingMode() bool`

HasParsingMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# CompleteLiteralTimeRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RangeName** | **string** | Name of the complete time range. Possible values are: - &#x60;today&#x60;, - &#x60;yesterday&#x60;, - &#x60;previous_week&#x60;, - &#x60;previous_month&#x60;. | 

## Methods

### NewCompleteLiteralTimeRange

`func NewCompleteLiteralTimeRange(rangeName string, ) *CompleteLiteralTimeRange`

NewCompleteLiteralTimeRange instantiates a new CompleteLiteralTimeRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteLiteralTimeRangeWithDefaults

`func NewCompleteLiteralTimeRangeWithDefaults() *CompleteLiteralTimeRange`

NewCompleteLiteralTimeRangeWithDefaults instantiates a new CompleteLiteralTimeRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRangeName

`func (o *CompleteLiteralTimeRange) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *CompleteLiteralTimeRange) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *CompleteLiteralTimeRange) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



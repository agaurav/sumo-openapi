# TimeRangeBoundary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of the time range boundary. Value must be from list: - &#x60;RelativeTimeRangeBoundary&#x60;, - &#x60;EpochTimeRangeBoundary&#x60;, - &#x60;Iso8601TimeRangeBoundary&#x60;, - &#x60;LiteralTimeRangeBoundary&#x60;. | 

## Methods

### NewTimeRangeBoundary

`func NewTimeRangeBoundary(type_ string, ) *TimeRangeBoundary`

NewTimeRangeBoundary instantiates a new TimeRangeBoundary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimeRangeBoundaryWithDefaults

`func NewTimeRangeBoundaryWithDefaults() *TimeRangeBoundary`

NewTimeRangeBoundaryWithDefaults instantiates a new TimeRangeBoundary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimeRangeBoundary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimeRangeBoundary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimeRangeBoundary) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



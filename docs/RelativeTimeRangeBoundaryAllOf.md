# RelativeTimeRangeBoundaryAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelativeTime** | **string** | Relative time as a string consisting of following elements: - &#x60;-&#x60; (optional): minus sign indicates time in the past, - &#x60;&lt;number&gt;&#x60;: number of time units, - &#x60;&lt;time_unit&gt;&#x60;: time unit; possible values are: &#x60;w&#x60; (week), &#x60;d&#x60; (day), &#x60;h&#x60; (hour), &#x60;m&#x60; (minute), &#x60;s&#x60; (second). Multiple pairs of &#x60;&lt;number&gt;&lt;time_unit&gt;&#x60; may be provided, and they may be in any order. For example, &#x60;-2w5d3h&#x60; points to the moment in time 2 weeks, 5 days and 3 hours ago. | 

## Methods

### NewRelativeTimeRangeBoundaryAllOf

`func NewRelativeTimeRangeBoundaryAllOf(relativeTime string, ) *RelativeTimeRangeBoundaryAllOf`

NewRelativeTimeRangeBoundaryAllOf instantiates a new RelativeTimeRangeBoundaryAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelativeTimeRangeBoundaryAllOfWithDefaults

`func NewRelativeTimeRangeBoundaryAllOfWithDefaults() *RelativeTimeRangeBoundaryAllOf`

NewRelativeTimeRangeBoundaryAllOfWithDefaults instantiates a new RelativeTimeRangeBoundaryAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelativeTime

`func (o *RelativeTimeRangeBoundaryAllOf) GetRelativeTime() string`

GetRelativeTime returns the RelativeTime field if non-nil, zero value otherwise.

### GetRelativeTimeOk

`func (o *RelativeTimeRangeBoundaryAllOf) GetRelativeTimeOk() (*string, bool)`

GetRelativeTimeOk returns a tuple with the RelativeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativeTime

`func (o *RelativeTimeRangeBoundaryAllOf) SetRelativeTime(v string)`

SetRelativeTime sets RelativeTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



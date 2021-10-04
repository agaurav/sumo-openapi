# Points

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timestamps** | **[]int64** | Array of timestamps of datapoints in milliseconds. | 
**Values** | **[]float64** | Array of values of datapoints corresponding to timestamp array. | 

## Methods

### NewPoints

`func NewPoints(timestamps []int64, values []float64, ) *Points`

NewPoints instantiates a new Points object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPointsWithDefaults

`func NewPointsWithDefaults() *Points`

NewPointsWithDefaults instantiates a new Points object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimestamps

`func (o *Points) GetTimestamps() []int64`

GetTimestamps returns the Timestamps field if non-nil, zero value otherwise.

### GetTimestampsOk

`func (o *Points) GetTimestampsOk() (*[]int64, bool)`

GetTimestampsOk returns a tuple with the Timestamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamps

`func (o *Points) SetTimestamps(v []int64)`

SetTimestamps sets Timestamps field to given value.


### GetValues

`func (o *Points) GetValues() []float64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Points) GetValuesOk() (*[]float64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Points) SetValues(v []float64)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



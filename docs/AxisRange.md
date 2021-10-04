# AxisRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Min** | Pointer to **int64** | minimum limit of x or y axis. | [optional] 
**Max** | Pointer to **int64** | maximum limit of x or y axis. | [optional] 

## Methods

### NewAxisRange

`func NewAxisRange() *AxisRange`

NewAxisRange instantiates a new AxisRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAxisRangeWithDefaults

`func NewAxisRangeWithDefaults() *AxisRange`

NewAxisRangeWithDefaults instantiates a new AxisRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMin

`func (o *AxisRange) GetMin() int64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *AxisRange) GetMinOk() (*int64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *AxisRange) SetMin(v int64)`

SetMin sets Min field to given value.

### HasMin

`func (o *AxisRange) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *AxisRange) GetMax() int64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *AxisRange) GetMaxOk() (*int64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *AxisRange) SetMax(v int64)`

SetMax sets Max field to given value.

### HasMax

`func (o *AxisRange) HasMax() bool`

HasMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



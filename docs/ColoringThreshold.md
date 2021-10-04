# ColoringThreshold

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** | Color for the threshold. | 
**Min** | Pointer to **float64** | Absolute inclusive threshold to color by. | [optional] 
**Max** | Pointer to **float64** | Absolute exclusive threshold to color by. | [optional] 

## Methods

### NewColoringThreshold

`func NewColoringThreshold(color string, ) *ColoringThreshold`

NewColoringThreshold instantiates a new ColoringThreshold object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewColoringThresholdWithDefaults

`func NewColoringThresholdWithDefaults() *ColoringThreshold`

NewColoringThresholdWithDefaults instantiates a new ColoringThreshold object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ColoringThreshold) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ColoringThreshold) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ColoringThreshold) SetColor(v string)`

SetColor sets Color field to given value.


### GetMin

`func (o *ColoringThreshold) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *ColoringThreshold) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *ColoringThreshold) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *ColoringThreshold) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMax

`func (o *ColoringThreshold) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *ColoringThreshold) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *ColoringThreshold) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *ColoringThreshold) HasMax() bool`

HasMax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



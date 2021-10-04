# DimensionKeyValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the metric dimension. | [optional] 
**Value** | Pointer to **string** | The value of the metric dimension. | [optional] 

## Methods

### NewDimensionKeyValue

`func NewDimensionKeyValue() *DimensionKeyValue`

NewDimensionKeyValue instantiates a new DimensionKeyValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDimensionKeyValueWithDefaults

`func NewDimensionKeyValueWithDefaults() *DimensionKeyValue`

NewDimensionKeyValueWithDefaults instantiates a new DimensionKeyValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *DimensionKeyValue) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *DimensionKeyValue) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *DimensionKeyValue) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *DimensionKeyValue) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *DimensionKeyValue) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DimensionKeyValue) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DimensionKeyValue) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *DimensionKeyValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



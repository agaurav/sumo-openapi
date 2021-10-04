# MetricsFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the metrics filter. | [optional] 
**Value** | **string** | The value of the metrics filter. | 
**Negation** | Pointer to **bool** | Whether or not the metrics filter is negated. | [optional] 

## Methods

### NewMetricsFilter

`func NewMetricsFilter(value string, ) *MetricsFilter`

NewMetricsFilter instantiates a new MetricsFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricsFilterWithDefaults

`func NewMetricsFilterWithDefaults() *MetricsFilter`

NewMetricsFilterWithDefaults instantiates a new MetricsFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *MetricsFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricsFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricsFilter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricsFilter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *MetricsFilter) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricsFilter) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricsFilter) SetValue(v string)`

SetValue sets Value field to given value.


### GetNegation

`func (o *MetricsFilter) GetNegation() bool`

GetNegation returns the Negation field if non-nil, zero value otherwise.

### GetNegationOk

`func (o *MetricsFilter) GetNegationOk() (*bool, bool)`

GetNegationOk returns a tuple with the Negation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegation

`func (o *MetricsFilter) SetNegation(v bool)`

SetNegation sets Negation field to given value.

### HasNegation

`func (o *MetricsFilter) HasNegation() bool`

HasNegation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



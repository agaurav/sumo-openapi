# MetricDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metric** | Pointer to **string** | Name of the metric returning the timeseries. | [optional] 
**Dimensions** | Pointer to **map[string]string** | Metric dimensions / metadata related to each timeseries. | [optional] 

## Methods

### NewMetricDefinition

`func NewMetricDefinition() *MetricDefinition`

NewMetricDefinition instantiates a new MetricDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDefinitionWithDefaults

`func NewMetricDefinitionWithDefaults() *MetricDefinition`

NewMetricDefinitionWithDefaults instantiates a new MetricDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetric

`func (o *MetricDefinition) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *MetricDefinition) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *MetricDefinition) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *MetricDefinition) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetDimensions

`func (o *MetricDefinition) GetDimensions() map[string]string`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *MetricDefinition) GetDimensionsOk() (*map[string]string, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *MetricDefinition) SetDimensions(v map[string]string)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *MetricDefinition) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Collector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CollectorId** | **string** | Identifier of a collector. | 
**CollectorName** | **string** | Name of a collector. | 

## Methods

### NewCollector

`func NewCollector(collectorId string, collectorName string, ) *Collector`

NewCollector instantiates a new Collector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectorWithDefaults

`func NewCollectorWithDefaults() *Collector`

NewCollectorWithDefaults instantiates a new Collector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollectorId

`func (o *Collector) GetCollectorId() string`

GetCollectorId returns the CollectorId field if non-nil, zero value otherwise.

### GetCollectorIdOk

`func (o *Collector) GetCollectorIdOk() (*string, bool)`

GetCollectorIdOk returns a tuple with the CollectorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorId

`func (o *Collector) SetCollectorId(v string)`

SetCollectorId sets CollectorId field to given value.


### GetCollectorName

`func (o *Collector) GetCollectorName() string`

GetCollectorName returns the CollectorName field if non-nil, zero value otherwise.

### GetCollectorNameOk

`func (o *Collector) GetCollectorNameOk() (*string, bool)`

GetCollectorNameOk returns a tuple with the CollectorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorName

`func (o *Collector) SetCollectorName(v string)`

SetCollectorName sets CollectorName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



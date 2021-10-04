# Baselines

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContinuousIngest** | Pointer to **int64** | The amount of continuous logs ingest to allocate to the organization, in GBs. | [optional] [default to 0]
**ContinuousStorage** | Pointer to **int64** | Number of days of continuous logs storage to allocate to the organization, in Days. | [optional] [default to 30]
**FrequentIngest** | Pointer to **int64** | The amount of frequent logs ingest to allocate to the organization, in GBs. | [optional] [default to 0]
**FrequentStorage** | Pointer to **int64** | Number of days of frequent logs storage to allocate to the organization, in Days. | [optional] [default to 30]
**InfrequentIngest** | Pointer to **int64** | The amount of infrequent logs ingest to allocate to the organization, in GBs. | [optional] [default to 0]
**InfrequentStorage** | Pointer to **int64** | The amount of infrequent logs storage to allocate to the organization, in Days. | [optional] [default to 30]
**InfrequentScan** | Pointer to **int64** | The amount of infrequent logs scan to allocate to the organization, in GBs. | [optional] [default to 0]
**Metrics** | Pointer to **int64** | The amount of Metrics usage to allocate to the organization, in DPMs (Data Points per Minute). | [optional] [default to 0]
**CseIngest** | Pointer to **int64** | The amount of CSE ingest to allocate to the organization, in GBs. | [optional] [default to 0]
**CseStorage** | Pointer to **int64** | The amount of CSE storage to allocate to the organization, in GBs. | [optional] [default to 0]
**TracingIngest** | Pointer to **int64** | The amount of tracing data ingest to allocate to the organization, in GBs. | [optional] [default to 0]

## Methods

### NewBaselines

`func NewBaselines() *Baselines`

NewBaselines instantiates a new Baselines object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaselinesWithDefaults

`func NewBaselinesWithDefaults() *Baselines`

NewBaselinesWithDefaults instantiates a new Baselines object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContinuousIngest

`func (o *Baselines) GetContinuousIngest() int64`

GetContinuousIngest returns the ContinuousIngest field if non-nil, zero value otherwise.

### GetContinuousIngestOk

`func (o *Baselines) GetContinuousIngestOk() (*int64, bool)`

GetContinuousIngestOk returns a tuple with the ContinuousIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousIngest

`func (o *Baselines) SetContinuousIngest(v int64)`

SetContinuousIngest sets ContinuousIngest field to given value.

### HasContinuousIngest

`func (o *Baselines) HasContinuousIngest() bool`

HasContinuousIngest returns a boolean if a field has been set.

### GetContinuousStorage

`func (o *Baselines) GetContinuousStorage() int64`

GetContinuousStorage returns the ContinuousStorage field if non-nil, zero value otherwise.

### GetContinuousStorageOk

`func (o *Baselines) GetContinuousStorageOk() (*int64, bool)`

GetContinuousStorageOk returns a tuple with the ContinuousStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContinuousStorage

`func (o *Baselines) SetContinuousStorage(v int64)`

SetContinuousStorage sets ContinuousStorage field to given value.

### HasContinuousStorage

`func (o *Baselines) HasContinuousStorage() bool`

HasContinuousStorage returns a boolean if a field has been set.

### GetFrequentIngest

`func (o *Baselines) GetFrequentIngest() int64`

GetFrequentIngest returns the FrequentIngest field if non-nil, zero value otherwise.

### GetFrequentIngestOk

`func (o *Baselines) GetFrequentIngestOk() (*int64, bool)`

GetFrequentIngestOk returns a tuple with the FrequentIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentIngest

`func (o *Baselines) SetFrequentIngest(v int64)`

SetFrequentIngest sets FrequentIngest field to given value.

### HasFrequentIngest

`func (o *Baselines) HasFrequentIngest() bool`

HasFrequentIngest returns a boolean if a field has been set.

### GetFrequentStorage

`func (o *Baselines) GetFrequentStorage() int64`

GetFrequentStorage returns the FrequentStorage field if non-nil, zero value otherwise.

### GetFrequentStorageOk

`func (o *Baselines) GetFrequentStorageOk() (*int64, bool)`

GetFrequentStorageOk returns a tuple with the FrequentStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequentStorage

`func (o *Baselines) SetFrequentStorage(v int64)`

SetFrequentStorage sets FrequentStorage field to given value.

### HasFrequentStorage

`func (o *Baselines) HasFrequentStorage() bool`

HasFrequentStorage returns a boolean if a field has been set.

### GetInfrequentIngest

`func (o *Baselines) GetInfrequentIngest() int64`

GetInfrequentIngest returns the InfrequentIngest field if non-nil, zero value otherwise.

### GetInfrequentIngestOk

`func (o *Baselines) GetInfrequentIngestOk() (*int64, bool)`

GetInfrequentIngestOk returns a tuple with the InfrequentIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrequentIngest

`func (o *Baselines) SetInfrequentIngest(v int64)`

SetInfrequentIngest sets InfrequentIngest field to given value.

### HasInfrequentIngest

`func (o *Baselines) HasInfrequentIngest() bool`

HasInfrequentIngest returns a boolean if a field has been set.

### GetInfrequentStorage

`func (o *Baselines) GetInfrequentStorage() int64`

GetInfrequentStorage returns the InfrequentStorage field if non-nil, zero value otherwise.

### GetInfrequentStorageOk

`func (o *Baselines) GetInfrequentStorageOk() (*int64, bool)`

GetInfrequentStorageOk returns a tuple with the InfrequentStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrequentStorage

`func (o *Baselines) SetInfrequentStorage(v int64)`

SetInfrequentStorage sets InfrequentStorage field to given value.

### HasInfrequentStorage

`func (o *Baselines) HasInfrequentStorage() bool`

HasInfrequentStorage returns a boolean if a field has been set.

### GetInfrequentScan

`func (o *Baselines) GetInfrequentScan() int64`

GetInfrequentScan returns the InfrequentScan field if non-nil, zero value otherwise.

### GetInfrequentScanOk

`func (o *Baselines) GetInfrequentScanOk() (*int64, bool)`

GetInfrequentScanOk returns a tuple with the InfrequentScan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfrequentScan

`func (o *Baselines) SetInfrequentScan(v int64)`

SetInfrequentScan sets InfrequentScan field to given value.

### HasInfrequentScan

`func (o *Baselines) HasInfrequentScan() bool`

HasInfrequentScan returns a boolean if a field has been set.

### GetMetrics

`func (o *Baselines) GetMetrics() int64`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *Baselines) GetMetricsOk() (*int64, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *Baselines) SetMetrics(v int64)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *Baselines) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetCseIngest

`func (o *Baselines) GetCseIngest() int64`

GetCseIngest returns the CseIngest field if non-nil, zero value otherwise.

### GetCseIngestOk

`func (o *Baselines) GetCseIngestOk() (*int64, bool)`

GetCseIngestOk returns a tuple with the CseIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseIngest

`func (o *Baselines) SetCseIngest(v int64)`

SetCseIngest sets CseIngest field to given value.

### HasCseIngest

`func (o *Baselines) HasCseIngest() bool`

HasCseIngest returns a boolean if a field has been set.

### GetCseStorage

`func (o *Baselines) GetCseStorage() int64`

GetCseStorage returns the CseStorage field if non-nil, zero value otherwise.

### GetCseStorageOk

`func (o *Baselines) GetCseStorageOk() (*int64, bool)`

GetCseStorageOk returns a tuple with the CseStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCseStorage

`func (o *Baselines) SetCseStorage(v int64)`

SetCseStorage sets CseStorage field to given value.

### HasCseStorage

`func (o *Baselines) HasCseStorage() bool`

HasCseStorage returns a boolean if a field has been set.

### GetTracingIngest

`func (o *Baselines) GetTracingIngest() int64`

GetTracingIngest returns the TracingIngest field if non-nil, zero value otherwise.

### GetTracingIngestOk

`func (o *Baselines) GetTracingIngestOk() (*int64, bool)`

GetTracingIngestOk returns a tuple with the TracingIngest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracingIngest

`func (o *Baselines) SetTracingIngest(v int64)`

SetTracingIngest sets TracingIngest field to given value.

### HasTracingIngest

`func (o *Baselines) HasTracingIngest() bool`

HasTracingIngest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



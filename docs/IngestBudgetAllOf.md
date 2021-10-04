# IngestBudgetAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the ingest budget. | 
**UsageBytes** | Pointer to **int64** | Current usage since the last reset, in bytes. | [optional] 
**UsageStatus** | Pointer to **string** | Status of the current usage. Can be &#x60;Normal&#x60;, &#x60;Approaching&#x60;, &#x60;Exceeded&#x60;, or &#x60;Unknown&#x60; (unable to retrieve usage). | [optional] 
**NumberOfCollectors** | Pointer to **int64** | Number of collectors assigned to the ingest budget. | [optional] 

## Methods

### NewIngestBudgetAllOf

`func NewIngestBudgetAllOf(id string, ) *IngestBudgetAllOf`

NewIngestBudgetAllOf instantiates a new IngestBudgetAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestBudgetAllOfWithDefaults

`func NewIngestBudgetAllOfWithDefaults() *IngestBudgetAllOf`

NewIngestBudgetAllOfWithDefaults instantiates a new IngestBudgetAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IngestBudgetAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IngestBudgetAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IngestBudgetAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetUsageBytes

`func (o *IngestBudgetAllOf) GetUsageBytes() int64`

GetUsageBytes returns the UsageBytes field if non-nil, zero value otherwise.

### GetUsageBytesOk

`func (o *IngestBudgetAllOf) GetUsageBytesOk() (*int64, bool)`

GetUsageBytesOk returns a tuple with the UsageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageBytes

`func (o *IngestBudgetAllOf) SetUsageBytes(v int64)`

SetUsageBytes sets UsageBytes field to given value.

### HasUsageBytes

`func (o *IngestBudgetAllOf) HasUsageBytes() bool`

HasUsageBytes returns a boolean if a field has been set.

### GetUsageStatus

`func (o *IngestBudgetAllOf) GetUsageStatus() string`

GetUsageStatus returns the UsageStatus field if non-nil, zero value otherwise.

### GetUsageStatusOk

`func (o *IngestBudgetAllOf) GetUsageStatusOk() (*string, bool)`

GetUsageStatusOk returns a tuple with the UsageStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageStatus

`func (o *IngestBudgetAllOf) SetUsageStatus(v string)`

SetUsageStatus sets UsageStatus field to given value.

### HasUsageStatus

`func (o *IngestBudgetAllOf) HasUsageStatus() bool`

HasUsageStatus returns a boolean if a field has been set.

### GetNumberOfCollectors

`func (o *IngestBudgetAllOf) GetNumberOfCollectors() int64`

GetNumberOfCollectors returns the NumberOfCollectors field if non-nil, zero value otherwise.

### GetNumberOfCollectorsOk

`func (o *IngestBudgetAllOf) GetNumberOfCollectorsOk() (*int64, bool)`

GetNumberOfCollectorsOk returns a tuple with the NumberOfCollectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfCollectors

`func (o *IngestBudgetAllOf) SetNumberOfCollectors(v int64)`

SetNumberOfCollectors sets NumberOfCollectors field to given value.

### HasNumberOfCollectors

`func (o *IngestBudgetAllOf) HasNumberOfCollectors() bool`

HasNumberOfCollectors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



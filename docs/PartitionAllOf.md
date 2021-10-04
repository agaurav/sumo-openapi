# PartitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the partition. | 
**TotalBytes** | **int64** | Size of data in partition in bytes. | 
**IsActive** | Pointer to **bool** | This has the value &#x60;true&#x60; if the partition is active and &#x60;false&#x60; if it has been decommissioned. | [optional] 
**NewRetentionPeriod** | Pointer to **int32** | If the retentionPeriod is scheduled to be updated in the future (i.e., if retentionPeriod is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retentionPeriod while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time. | [optional] 
**IndexType** | Pointer to **string** | This has the value &#x60;DefaultIndex&#x60;, &#x60;AuditIndex&#x60;or &#x60;Partition&#x60; depending upon the type of partition. | [optional] 
**RetentionEffectiveAt** | Pointer to **time.Time** | When the newRetentionPeriod will become effective in UTC format. | [optional] 
**DataForwardingId** | Pointer to **string** | Id of the data forwarding configuration to be used by the partition. | [optional] 

## Methods

### NewPartitionAllOf

`func NewPartitionAllOf(id string, totalBytes int64, ) *PartitionAllOf`

NewPartitionAllOf instantiates a new PartitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionAllOfWithDefaults

`func NewPartitionAllOfWithDefaults() *PartitionAllOf`

NewPartitionAllOfWithDefaults instantiates a new PartitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PartitionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PartitionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PartitionAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetTotalBytes

`func (o *PartitionAllOf) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *PartitionAllOf) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *PartitionAllOf) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.


### GetIsActive

`func (o *PartitionAllOf) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *PartitionAllOf) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *PartitionAllOf) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *PartitionAllOf) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetNewRetentionPeriod

`func (o *PartitionAllOf) GetNewRetentionPeriod() int32`

GetNewRetentionPeriod returns the NewRetentionPeriod field if non-nil, zero value otherwise.

### GetNewRetentionPeriodOk

`func (o *PartitionAllOf) GetNewRetentionPeriodOk() (*int32, bool)`

GetNewRetentionPeriodOk returns a tuple with the NewRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRetentionPeriod

`func (o *PartitionAllOf) SetNewRetentionPeriod(v int32)`

SetNewRetentionPeriod sets NewRetentionPeriod field to given value.

### HasNewRetentionPeriod

`func (o *PartitionAllOf) HasNewRetentionPeriod() bool`

HasNewRetentionPeriod returns a boolean if a field has been set.

### GetIndexType

`func (o *PartitionAllOf) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *PartitionAllOf) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *PartitionAllOf) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *PartitionAllOf) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetRetentionEffectiveAt

`func (o *PartitionAllOf) GetRetentionEffectiveAt() time.Time`

GetRetentionEffectiveAt returns the RetentionEffectiveAt field if non-nil, zero value otherwise.

### GetRetentionEffectiveAtOk

`func (o *PartitionAllOf) GetRetentionEffectiveAtOk() (*time.Time, bool)`

GetRetentionEffectiveAtOk returns a tuple with the RetentionEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionEffectiveAt

`func (o *PartitionAllOf) SetRetentionEffectiveAt(v time.Time)`

SetRetentionEffectiveAt sets RetentionEffectiveAt field to given value.

### HasRetentionEffectiveAt

`func (o *PartitionAllOf) HasRetentionEffectiveAt() bool`

HasRetentionEffectiveAt returns a boolean if a field has been set.

### GetDataForwardingId

`func (o *PartitionAllOf) GetDataForwardingId() string`

GetDataForwardingId returns the DataForwardingId field if non-nil, zero value otherwise.

### GetDataForwardingIdOk

`func (o *PartitionAllOf) GetDataForwardingIdOk() (*string, bool)`

GetDataForwardingIdOk returns a tuple with the DataForwardingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwardingId

`func (o *PartitionAllOf) SetDataForwardingId(v string)`

SetDataForwardingId sets DataForwardingId field to given value.

### HasDataForwardingId

`func (o *PartitionAllOf) HasDataForwardingId() bool`

HasDataForwardingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Partition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the partition. | 
**RoutingExpression** | **string** | The query that defines the data to be included in the partition. | 
**AnalyticsTier** | Pointer to **string** | The Data Tier where the data in the partition will reside. Possible values are:               1. &#x60;continuous&#x60;               2. &#x60;frequent&#x60;               3. &#x60;infrequent&#x60; Note: The \&quot;infrequent\&quot; and \&quot;frequent\&quot; tiers are only to available Cloud Flex Credits Enterprise Suite accounts. | [optional] [default to "continuous"]
**RetentionPeriod** | Pointer to **int32** | The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled. | [optional] [default to -1]
**IsCompliant** | Pointer to **bool** | Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later. | [optional] [default to false]
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | Unique identifier for the partition. | 
**TotalBytes** | **int64** | Size of data in partition in bytes. | 
**IsActive** | Pointer to **bool** | This has the value &#x60;true&#x60; if the partition is active and &#x60;false&#x60; if it has been decommissioned. | [optional] 
**NewRetentionPeriod** | Pointer to **int32** | If the retentionPeriod is scheduled to be updated in the future (i.e., if retentionPeriod is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retentionPeriod while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time. | [optional] 
**IndexType** | Pointer to **string** | This has the value &#x60;DefaultIndex&#x60;, &#x60;AuditIndex&#x60;or &#x60;Partition&#x60; depending upon the type of partition. | [optional] 
**RetentionEffectiveAt** | Pointer to **time.Time** | When the newRetentionPeriod will become effective in UTC format. | [optional] 
**DataForwardingId** | Pointer to **string** | Id of the data forwarding configuration to be used by the partition. | [optional] 

## Methods

### NewPartition

`func NewPartition(name string, routingExpression string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, totalBytes int64, ) *Partition`

NewPartition instantiates a new Partition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionWithDefaults

`func NewPartitionWithDefaults() *Partition`

NewPartitionWithDefaults instantiates a new Partition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Partition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Partition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Partition) SetName(v string)`

SetName sets Name field to given value.


### GetRoutingExpression

`func (o *Partition) GetRoutingExpression() string`

GetRoutingExpression returns the RoutingExpression field if non-nil, zero value otherwise.

### GetRoutingExpressionOk

`func (o *Partition) GetRoutingExpressionOk() (*string, bool)`

GetRoutingExpressionOk returns a tuple with the RoutingExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingExpression

`func (o *Partition) SetRoutingExpression(v string)`

SetRoutingExpression sets RoutingExpression field to given value.


### GetAnalyticsTier

`func (o *Partition) GetAnalyticsTier() string`

GetAnalyticsTier returns the AnalyticsTier field if non-nil, zero value otherwise.

### GetAnalyticsTierOk

`func (o *Partition) GetAnalyticsTierOk() (*string, bool)`

GetAnalyticsTierOk returns a tuple with the AnalyticsTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTier

`func (o *Partition) SetAnalyticsTier(v string)`

SetAnalyticsTier sets AnalyticsTier field to given value.

### HasAnalyticsTier

`func (o *Partition) HasAnalyticsTier() bool`

HasAnalyticsTier returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *Partition) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *Partition) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *Partition) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *Partition) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetIsCompliant

`func (o *Partition) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *Partition) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *Partition) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *Partition) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Partition) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Partition) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Partition) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Partition) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Partition) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Partition) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *Partition) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *Partition) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *Partition) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *Partition) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *Partition) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *Partition) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *Partition) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Partition) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Partition) SetId(v string)`

SetId sets Id field to given value.


### GetTotalBytes

`func (o *Partition) GetTotalBytes() int64`

GetTotalBytes returns the TotalBytes field if non-nil, zero value otherwise.

### GetTotalBytesOk

`func (o *Partition) GetTotalBytesOk() (*int64, bool)`

GetTotalBytesOk returns a tuple with the TotalBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalBytes

`func (o *Partition) SetTotalBytes(v int64)`

SetTotalBytes sets TotalBytes field to given value.


### GetIsActive

`func (o *Partition) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Partition) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Partition) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Partition) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetNewRetentionPeriod

`func (o *Partition) GetNewRetentionPeriod() int32`

GetNewRetentionPeriod returns the NewRetentionPeriod field if non-nil, zero value otherwise.

### GetNewRetentionPeriodOk

`func (o *Partition) GetNewRetentionPeriodOk() (*int32, bool)`

GetNewRetentionPeriodOk returns a tuple with the NewRetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRetentionPeriod

`func (o *Partition) SetNewRetentionPeriod(v int32)`

SetNewRetentionPeriod sets NewRetentionPeriod field to given value.

### HasNewRetentionPeriod

`func (o *Partition) HasNewRetentionPeriod() bool`

HasNewRetentionPeriod returns a boolean if a field has been set.

### GetIndexType

`func (o *Partition) GetIndexType() string`

GetIndexType returns the IndexType field if non-nil, zero value otherwise.

### GetIndexTypeOk

`func (o *Partition) GetIndexTypeOk() (*string, bool)`

GetIndexTypeOk returns a tuple with the IndexType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndexType

`func (o *Partition) SetIndexType(v string)`

SetIndexType sets IndexType field to given value.

### HasIndexType

`func (o *Partition) HasIndexType() bool`

HasIndexType returns a boolean if a field has been set.

### GetRetentionEffectiveAt

`func (o *Partition) GetRetentionEffectiveAt() time.Time`

GetRetentionEffectiveAt returns the RetentionEffectiveAt field if non-nil, zero value otherwise.

### GetRetentionEffectiveAtOk

`func (o *Partition) GetRetentionEffectiveAtOk() (*time.Time, bool)`

GetRetentionEffectiveAtOk returns a tuple with the RetentionEffectiveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionEffectiveAt

`func (o *Partition) SetRetentionEffectiveAt(v time.Time)`

SetRetentionEffectiveAt sets RetentionEffectiveAt field to given value.

### HasRetentionEffectiveAt

`func (o *Partition) HasRetentionEffectiveAt() bool`

HasRetentionEffectiveAt returns a boolean if a field has been set.

### GetDataForwardingId

`func (o *Partition) GetDataForwardingId() string`

GetDataForwardingId returns the DataForwardingId field if non-nil, zero value otherwise.

### GetDataForwardingIdOk

`func (o *Partition) GetDataForwardingIdOk() (*string, bool)`

GetDataForwardingIdOk returns a tuple with the DataForwardingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataForwardingId

`func (o *Partition) SetDataForwardingId(v string)`

SetDataForwardingId sets DataForwardingId field to given value.

### HasDataForwardingId

`func (o *Partition) HasDataForwardingId() bool`

HasDataForwardingId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



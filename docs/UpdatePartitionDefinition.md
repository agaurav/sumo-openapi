# UpdatePartitionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetentionPeriod** | Pointer to **int32** | The number of days to retain data in the partition, or -1 to use the default value for your account. Only relevant if your account has variable retention enabled. | [optional] 
**ReduceRetentionPeriodImmediately** | Pointer to **bool** | This is required if the newly specified &#x60;retentionPeriod&#x60; is less than the existing retention period.  In such a situation, a value of &#x60;true&#x60; says that data between the existing retention period and the new retention period should be deleted immediately; if &#x60;false&#x60;, such data will be deleted after seven days. This property is optional and ignored if the specified &#x60;retentionPeriod&#x60; is greater than or equal to the current retention period. | [optional] [default to false]
**IsCompliant** | Pointer to **bool** | Whether to mark a partition as compliant. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition marked as compliant. A partition once marked compliant, cannot be marked non-compliant later. | [optional] [default to false]
**RoutingExpression** | Pointer to **string** | The query that defines the data to be included in the partition. | [optional] 

## Methods

### NewUpdatePartitionDefinition

`func NewUpdatePartitionDefinition() *UpdatePartitionDefinition`

NewUpdatePartitionDefinition instantiates a new UpdatePartitionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePartitionDefinitionWithDefaults

`func NewUpdatePartitionDefinitionWithDefaults() *UpdatePartitionDefinition`

NewUpdatePartitionDefinitionWithDefaults instantiates a new UpdatePartitionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetentionPeriod

`func (o *UpdatePartitionDefinition) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *UpdatePartitionDefinition) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *UpdatePartitionDefinition) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *UpdatePartitionDefinition) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetReduceRetentionPeriodImmediately

`func (o *UpdatePartitionDefinition) GetReduceRetentionPeriodImmediately() bool`

GetReduceRetentionPeriodImmediately returns the ReduceRetentionPeriodImmediately field if non-nil, zero value otherwise.

### GetReduceRetentionPeriodImmediatelyOk

`func (o *UpdatePartitionDefinition) GetReduceRetentionPeriodImmediatelyOk() (*bool, bool)`

GetReduceRetentionPeriodImmediatelyOk returns a tuple with the ReduceRetentionPeriodImmediately field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReduceRetentionPeriodImmediately

`func (o *UpdatePartitionDefinition) SetReduceRetentionPeriodImmediately(v bool)`

SetReduceRetentionPeriodImmediately sets ReduceRetentionPeriodImmediately field to given value.

### HasReduceRetentionPeriodImmediately

`func (o *UpdatePartitionDefinition) HasReduceRetentionPeriodImmediately() bool`

HasReduceRetentionPeriodImmediately returns a boolean if a field has been set.

### GetIsCompliant

`func (o *UpdatePartitionDefinition) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *UpdatePartitionDefinition) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *UpdatePartitionDefinition) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *UpdatePartitionDefinition) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.

### GetRoutingExpression

`func (o *UpdatePartitionDefinition) GetRoutingExpression() string`

GetRoutingExpression returns the RoutingExpression field if non-nil, zero value otherwise.

### GetRoutingExpressionOk

`func (o *UpdatePartitionDefinition) GetRoutingExpressionOk() (*string, bool)`

GetRoutingExpressionOk returns a tuple with the RoutingExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingExpression

`func (o *UpdatePartitionDefinition) SetRoutingExpression(v string)`

SetRoutingExpression sets RoutingExpression field to given value.

### HasRoutingExpression

`func (o *UpdatePartitionDefinition) HasRoutingExpression() bool`

HasRoutingExpression returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



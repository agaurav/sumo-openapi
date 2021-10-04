# CreatePartitionDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the partition. | 
**RoutingExpression** | **string** | The query that defines the data to be included in the partition. | 
**AnalyticsTier** | Pointer to **string** | The Data Tier where the data in the partition will reside. Possible values are:               1. &#x60;continuous&#x60;               2. &#x60;frequent&#x60;               3. &#x60;infrequent&#x60; Note: The \&quot;infrequent\&quot; and \&quot;frequent\&quot; tiers are only to available Cloud Flex Credits Enterprise Suite accounts. | [optional] [default to "continuous"]
**RetentionPeriod** | Pointer to **int32** | The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled. | [optional] [default to -1]
**IsCompliant** | Pointer to **bool** | Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later. | [optional] [default to false]

## Methods

### NewCreatePartitionDefinition

`func NewCreatePartitionDefinition(name string, routingExpression string, ) *CreatePartitionDefinition`

NewCreatePartitionDefinition instantiates a new CreatePartitionDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePartitionDefinitionWithDefaults

`func NewCreatePartitionDefinitionWithDefaults() *CreatePartitionDefinition`

NewCreatePartitionDefinitionWithDefaults instantiates a new CreatePartitionDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreatePartitionDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreatePartitionDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreatePartitionDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetRoutingExpression

`func (o *CreatePartitionDefinition) GetRoutingExpression() string`

GetRoutingExpression returns the RoutingExpression field if non-nil, zero value otherwise.

### GetRoutingExpressionOk

`func (o *CreatePartitionDefinition) GetRoutingExpressionOk() (*string, bool)`

GetRoutingExpressionOk returns a tuple with the RoutingExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingExpression

`func (o *CreatePartitionDefinition) SetRoutingExpression(v string)`

SetRoutingExpression sets RoutingExpression field to given value.


### GetAnalyticsTier

`func (o *CreatePartitionDefinition) GetAnalyticsTier() string`

GetAnalyticsTier returns the AnalyticsTier field if non-nil, zero value otherwise.

### GetAnalyticsTierOk

`func (o *CreatePartitionDefinition) GetAnalyticsTierOk() (*string, bool)`

GetAnalyticsTierOk returns a tuple with the AnalyticsTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyticsTier

`func (o *CreatePartitionDefinition) SetAnalyticsTier(v string)`

SetAnalyticsTier sets AnalyticsTier field to given value.

### HasAnalyticsTier

`func (o *CreatePartitionDefinition) HasAnalyticsTier() bool`

HasAnalyticsTier returns a boolean if a field has been set.

### GetRetentionPeriod

`func (o *CreatePartitionDefinition) GetRetentionPeriod() int32`

GetRetentionPeriod returns the RetentionPeriod field if non-nil, zero value otherwise.

### GetRetentionPeriodOk

`func (o *CreatePartitionDefinition) GetRetentionPeriodOk() (*int32, bool)`

GetRetentionPeriodOk returns a tuple with the RetentionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPeriod

`func (o *CreatePartitionDefinition) SetRetentionPeriod(v int32)`

SetRetentionPeriod sets RetentionPeriod field to given value.

### HasRetentionPeriod

`func (o *CreatePartitionDefinition) HasRetentionPeriod() bool`

HasRetentionPeriod returns a boolean if a field has been set.

### GetIsCompliant

`func (o *CreatePartitionDefinition) GetIsCompliant() bool`

GetIsCompliant returns the IsCompliant field if non-nil, zero value otherwise.

### GetIsCompliantOk

`func (o *CreatePartitionDefinition) GetIsCompliantOk() (*bool, bool)`

GetIsCompliantOk returns a tuple with the IsCompliant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCompliant

`func (o *CreatePartitionDefinition) SetIsCompliant(v bool)`

SetIsCompliant sets IsCompliant field to given value.

### HasIsCompliant

`func (o *CreatePartitionDefinition) HasIsCompliant() bool`

HasIsCompliant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



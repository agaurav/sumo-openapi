# LookupTablesLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TablesCreated** | Pointer to **int32** | Number of lookup tables currently created. | [optional] 
**TableCapacityRemaining** | Pointer to **int32** | Remaining count of lookup tables that can be created. | [optional] 
**TotalTableCapacity** | Pointer to **int32** | Total capacity of lookup tables that can be created for the given org id. | [optional] 

## Methods

### NewLookupTablesLimits

`func NewLookupTablesLimits() *LookupTablesLimits`

NewLookupTablesLimits instantiates a new LookupTablesLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTablesLimitsWithDefaults

`func NewLookupTablesLimitsWithDefaults() *LookupTablesLimits`

NewLookupTablesLimitsWithDefaults instantiates a new LookupTablesLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTablesCreated

`func (o *LookupTablesLimits) GetTablesCreated() int32`

GetTablesCreated returns the TablesCreated field if non-nil, zero value otherwise.

### GetTablesCreatedOk

`func (o *LookupTablesLimits) GetTablesCreatedOk() (*int32, bool)`

GetTablesCreatedOk returns a tuple with the TablesCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTablesCreated

`func (o *LookupTablesLimits) SetTablesCreated(v int32)`

SetTablesCreated sets TablesCreated field to given value.

### HasTablesCreated

`func (o *LookupTablesLimits) HasTablesCreated() bool`

HasTablesCreated returns a boolean if a field has been set.

### GetTableCapacityRemaining

`func (o *LookupTablesLimits) GetTableCapacityRemaining() int32`

GetTableCapacityRemaining returns the TableCapacityRemaining field if non-nil, zero value otherwise.

### GetTableCapacityRemainingOk

`func (o *LookupTablesLimits) GetTableCapacityRemainingOk() (*int32, bool)`

GetTableCapacityRemainingOk returns a tuple with the TableCapacityRemaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTableCapacityRemaining

`func (o *LookupTablesLimits) SetTableCapacityRemaining(v int32)`

SetTableCapacityRemaining sets TableCapacityRemaining field to given value.

### HasTableCapacityRemaining

`func (o *LookupTablesLimits) HasTableCapacityRemaining() bool`

HasTableCapacityRemaining returns a boolean if a field has been set.

### GetTotalTableCapacity

`func (o *LookupTablesLimits) GetTotalTableCapacity() int32`

GetTotalTableCapacity returns the TotalTableCapacity field if non-nil, zero value otherwise.

### GetTotalTableCapacityOk

`func (o *LookupTablesLimits) GetTotalTableCapacityOk() (*int32, bool)`

GetTotalTableCapacityOk returns a tuple with the TotalTableCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTableCapacity

`func (o *LookupTablesLimits) SetTotalTableCapacity(v int32)`

SetTotalTableCapacity sets TotalTableCapacity field to given value.

### HasTotalTableCapacity

`func (o *LookupTablesLimits) HasTotalTableCapacity() bool`

HasTotalTableCapacity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# FieldQuotaUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Quota** | **int32** | Maximum number of fields available. | 
**Remaining** | **int32** | Current number of fields available. | 

## Methods

### NewFieldQuotaUsage

`func NewFieldQuotaUsage(quota int32, remaining int32, ) *FieldQuotaUsage`

NewFieldQuotaUsage instantiates a new FieldQuotaUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldQuotaUsageWithDefaults

`func NewFieldQuotaUsageWithDefaults() *FieldQuotaUsage`

NewFieldQuotaUsageWithDefaults instantiates a new FieldQuotaUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuota

`func (o *FieldQuotaUsage) GetQuota() int32`

GetQuota returns the Quota field if non-nil, zero value otherwise.

### GetQuotaOk

`func (o *FieldQuotaUsage) GetQuotaOk() (*int32, bool)`

GetQuotaOk returns a tuple with the Quota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuota

`func (o *FieldQuotaUsage) SetQuota(v int32)`

SetQuota sets Quota field to given value.


### GetRemaining

`func (o *FieldQuotaUsage) GetRemaining() int32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *FieldQuotaUsage) GetRemainingOk() (*int32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *FieldQuotaUsage) SetRemaining(v int32)`

SetRemaining sets Remaining field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



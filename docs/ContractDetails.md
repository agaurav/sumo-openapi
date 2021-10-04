# ContractDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrgId** | **string** | Organization identifier of the account. | 
**PlanType** | **string** | Plan name of the account. | 
**Entitlements** | [**[]Entitlements**](Entitlements.md) | List of the entitlements of the account. Entitlements of the account are the list of products subscribed by the user. | 
**SharedBuckets** | Pointer to [**[]SharedBucket**](SharedBucket.md) | Contains list of buckets. Bucket means shared pool from which multiple entitlements can use capacity. | [optional] 
**ContractPeriod** | [**ContractPeriod**](ContractPeriod.md) |  | 
**CurrentBillingPeriod** | [**CurrentBillingPeriod**](CurrentBillingPeriod.md) |  | 

## Methods

### NewContractDetails

`func NewContractDetails(orgId string, planType string, entitlements []Entitlements, contractPeriod ContractPeriod, currentBillingPeriod CurrentBillingPeriod, ) *ContractDetails`

NewContractDetails instantiates a new ContractDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContractDetailsWithDefaults

`func NewContractDetailsWithDefaults() *ContractDetails`

NewContractDetailsWithDefaults instantiates a new ContractDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrgId

`func (o *ContractDetails) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ContractDetails) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ContractDetails) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetPlanType

`func (o *ContractDetails) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *ContractDetails) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *ContractDetails) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetEntitlements

`func (o *ContractDetails) GetEntitlements() []Entitlements`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *ContractDetails) GetEntitlementsOk() (*[]Entitlements, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *ContractDetails) SetEntitlements(v []Entitlements)`

SetEntitlements sets Entitlements field to given value.


### GetSharedBuckets

`func (o *ContractDetails) GetSharedBuckets() []SharedBucket`

GetSharedBuckets returns the SharedBuckets field if non-nil, zero value otherwise.

### GetSharedBucketsOk

`func (o *ContractDetails) GetSharedBucketsOk() (*[]SharedBucket, bool)`

GetSharedBucketsOk returns a tuple with the SharedBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedBuckets

`func (o *ContractDetails) SetSharedBuckets(v []SharedBucket)`

SetSharedBuckets sets SharedBuckets field to given value.

### HasSharedBuckets

`func (o *ContractDetails) HasSharedBuckets() bool`

HasSharedBuckets returns a boolean if a field has been set.

### GetContractPeriod

`func (o *ContractDetails) GetContractPeriod() ContractPeriod`

GetContractPeriod returns the ContractPeriod field if non-nil, zero value otherwise.

### GetContractPeriodOk

`func (o *ContractDetails) GetContractPeriodOk() (*ContractPeriod, bool)`

GetContractPeriodOk returns a tuple with the ContractPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractPeriod

`func (o *ContractDetails) SetContractPeriod(v ContractPeriod)`

SetContractPeriod sets ContractPeriod field to given value.


### GetCurrentBillingPeriod

`func (o *ContractDetails) GetCurrentBillingPeriod() CurrentBillingPeriod`

GetCurrentBillingPeriod returns the CurrentBillingPeriod field if non-nil, zero value otherwise.

### GetCurrentBillingPeriodOk

`func (o *ContractDetails) GetCurrentBillingPeriodOk() (*CurrentBillingPeriod, bool)`

GetCurrentBillingPeriodOk returns a tuple with the CurrentBillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBillingPeriod

`func (o *ContractDetails) SetCurrentBillingPeriod(v CurrentBillingPeriod)`

SetCurrentBillingPeriod sets CurrentBillingPeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



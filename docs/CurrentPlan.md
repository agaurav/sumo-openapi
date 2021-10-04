# CurrentPlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Unique identifier of the product in current plan. Valid values are: 1. &#x60;Free&#x60; 2. &#x60;Trial&#x60; 3. &#x60;Essentials&#x60; 4. &#x60;EnterpriseOps&#x60; 5. &#x60;EnterpriseSec&#x60; 6. &#x60;EnterpriseSuite&#x60;  | 
**PlanCost** | **float64** | Cost incurred for the current plan. | 
**BillingFrequency** | **string** | Billing frequency for the current plan. Valid values are: 1. &#x60;Monthly&#x60; 2. &#x60;Annually&#x60;  | 
**Consumables** | Pointer to [**[]Consumable**](Consumable.md) | Consumables in the current plan. | [optional] 
**PlanType** | Pointer to **string** | Whether the account is &#x60;Free&#x60;/&#x60;Trial&#x60;/&#x60;Paid&#x60; | [optional] 
**PlanName** | Pointer to **string** | The plan name for the product being used. | [optional] 
**DiscountAmount** | Pointer to **int32** | The discount offered for the given contract period. | [optional] 
**ContractPeriod** | Pointer to [**ContractPeriod**](ContractPeriod.md) |  | [optional] 
**CurrentBillingPeriod** | Pointer to [**CurrentBillingPeriod**](CurrentBillingPeriod.md) |  | [optional] 
**Credits** | Pointer to **int64** | Numerical value of the amount of credits | [optional] 
**Baselines** | Pointer to [**Baselines**](Baselines.md) |  | [optional] 

## Methods

### NewCurrentPlan

`func NewCurrentPlan(productId string, planCost float64, billingFrequency string, ) *CurrentPlan`

NewCurrentPlan instantiates a new CurrentPlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentPlanWithDefaults

`func NewCurrentPlanWithDefaults() *CurrentPlan`

NewCurrentPlanWithDefaults instantiates a new CurrentPlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *CurrentPlan) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *CurrentPlan) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *CurrentPlan) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetPlanCost

`func (o *CurrentPlan) GetPlanCost() float64`

GetPlanCost returns the PlanCost field if non-nil, zero value otherwise.

### GetPlanCostOk

`func (o *CurrentPlan) GetPlanCostOk() (*float64, bool)`

GetPlanCostOk returns a tuple with the PlanCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCost

`func (o *CurrentPlan) SetPlanCost(v float64)`

SetPlanCost sets PlanCost field to given value.


### GetBillingFrequency

`func (o *CurrentPlan) GetBillingFrequency() string`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *CurrentPlan) GetBillingFrequencyOk() (*string, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *CurrentPlan) SetBillingFrequency(v string)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetConsumables

`func (o *CurrentPlan) GetConsumables() []Consumable`

GetConsumables returns the Consumables field if non-nil, zero value otherwise.

### GetConsumablesOk

`func (o *CurrentPlan) GetConsumablesOk() (*[]Consumable, bool)`

GetConsumablesOk returns a tuple with the Consumables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumables

`func (o *CurrentPlan) SetConsumables(v []Consumable)`

SetConsumables sets Consumables field to given value.

### HasConsumables

`func (o *CurrentPlan) HasConsumables() bool`

HasConsumables returns a boolean if a field has been set.

### GetPlanType

`func (o *CurrentPlan) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *CurrentPlan) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *CurrentPlan) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.

### HasPlanType

`func (o *CurrentPlan) HasPlanType() bool`

HasPlanType returns a boolean if a field has been set.

### GetPlanName

`func (o *CurrentPlan) GetPlanName() string`

GetPlanName returns the PlanName field if non-nil, zero value otherwise.

### GetPlanNameOk

`func (o *CurrentPlan) GetPlanNameOk() (*string, bool)`

GetPlanNameOk returns a tuple with the PlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanName

`func (o *CurrentPlan) SetPlanName(v string)`

SetPlanName sets PlanName field to given value.

### HasPlanName

`func (o *CurrentPlan) HasPlanName() bool`

HasPlanName returns a boolean if a field has been set.

### GetDiscountAmount

`func (o *CurrentPlan) GetDiscountAmount() int32`

GetDiscountAmount returns the DiscountAmount field if non-nil, zero value otherwise.

### GetDiscountAmountOk

`func (o *CurrentPlan) GetDiscountAmountOk() (*int32, bool)`

GetDiscountAmountOk returns a tuple with the DiscountAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountAmount

`func (o *CurrentPlan) SetDiscountAmount(v int32)`

SetDiscountAmount sets DiscountAmount field to given value.

### HasDiscountAmount

`func (o *CurrentPlan) HasDiscountAmount() bool`

HasDiscountAmount returns a boolean if a field has been set.

### GetContractPeriod

`func (o *CurrentPlan) GetContractPeriod() ContractPeriod`

GetContractPeriod returns the ContractPeriod field if non-nil, zero value otherwise.

### GetContractPeriodOk

`func (o *CurrentPlan) GetContractPeriodOk() (*ContractPeriod, bool)`

GetContractPeriodOk returns a tuple with the ContractPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractPeriod

`func (o *CurrentPlan) SetContractPeriod(v ContractPeriod)`

SetContractPeriod sets ContractPeriod field to given value.

### HasContractPeriod

`func (o *CurrentPlan) HasContractPeriod() bool`

HasContractPeriod returns a boolean if a field has been set.

### GetCurrentBillingPeriod

`func (o *CurrentPlan) GetCurrentBillingPeriod() CurrentBillingPeriod`

GetCurrentBillingPeriod returns the CurrentBillingPeriod field if non-nil, zero value otherwise.

### GetCurrentBillingPeriodOk

`func (o *CurrentPlan) GetCurrentBillingPeriodOk() (*CurrentBillingPeriod, bool)`

GetCurrentBillingPeriodOk returns a tuple with the CurrentBillingPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBillingPeriod

`func (o *CurrentPlan) SetCurrentBillingPeriod(v CurrentBillingPeriod)`

SetCurrentBillingPeriod sets CurrentBillingPeriod field to given value.

### HasCurrentBillingPeriod

`func (o *CurrentPlan) HasCurrentBillingPeriod() bool`

HasCurrentBillingPeriod returns a boolean if a field has been set.

### GetCredits

`func (o *CurrentPlan) GetCredits() int64`

GetCredits returns the Credits field if non-nil, zero value otherwise.

### GetCreditsOk

`func (o *CurrentPlan) GetCreditsOk() (*int64, bool)`

GetCreditsOk returns a tuple with the Credits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredits

`func (o *CurrentPlan) SetCredits(v int64)`

SetCredits sets Credits field to given value.

### HasCredits

`func (o *CurrentPlan) HasCredits() bool`

HasCredits returns a boolean if a field has been set.

### GetBaselines

`func (o *CurrentPlan) GetBaselines() Baselines`

GetBaselines returns the Baselines field if non-nil, zero value otherwise.

### GetBaselinesOk

`func (o *CurrentPlan) GetBaselinesOk() (*Baselines, bool)`

GetBaselinesOk returns a tuple with the Baselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselines

`func (o *CurrentPlan) SetBaselines(v Baselines)`

SetBaselines sets Baselines field to given value.

### HasBaselines

`func (o *CurrentPlan) HasBaselines() bool`

HasBaselines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



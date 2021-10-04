# AccountStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PricingModel** | **string** | Whether the account is &#x60;cloudflex&#x60; or &#x60;credits&#x60; | 
**CanUpdatePlan** | **bool** | If the plan can be updated by the given user | 
**PlanType** | **string** | Whether the account is &#x60;Free&#x60;/&#x60;Trial&#x60;/&#x60;Paid&#x60; | 
**PlanExpirationDays** | Pointer to **int32** | The number of days in which the plan will expire | [optional] 
**ApplicationUse** | **string** | The current usage of the application. | 
**AccountActivated** | Pointer to **bool** | If the account is activated or not | [optional] 

## Methods

### NewAccountStatusResponse

`func NewAccountStatusResponse(pricingModel string, canUpdatePlan bool, planType string, applicationUse string, ) *AccountStatusResponse`

NewAccountStatusResponse instantiates a new AccountStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountStatusResponseWithDefaults

`func NewAccountStatusResponseWithDefaults() *AccountStatusResponse`

NewAccountStatusResponseWithDefaults instantiates a new AccountStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPricingModel

`func (o *AccountStatusResponse) GetPricingModel() string`

GetPricingModel returns the PricingModel field if non-nil, zero value otherwise.

### GetPricingModelOk

`func (o *AccountStatusResponse) GetPricingModelOk() (*string, bool)`

GetPricingModelOk returns a tuple with the PricingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingModel

`func (o *AccountStatusResponse) SetPricingModel(v string)`

SetPricingModel sets PricingModel field to given value.


### GetCanUpdatePlan

`func (o *AccountStatusResponse) GetCanUpdatePlan() bool`

GetCanUpdatePlan returns the CanUpdatePlan field if non-nil, zero value otherwise.

### GetCanUpdatePlanOk

`func (o *AccountStatusResponse) GetCanUpdatePlanOk() (*bool, bool)`

GetCanUpdatePlanOk returns a tuple with the CanUpdatePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanUpdatePlan

`func (o *AccountStatusResponse) SetCanUpdatePlan(v bool)`

SetCanUpdatePlan sets CanUpdatePlan field to given value.


### GetPlanType

`func (o *AccountStatusResponse) GetPlanType() string`

GetPlanType returns the PlanType field if non-nil, zero value otherwise.

### GetPlanTypeOk

`func (o *AccountStatusResponse) GetPlanTypeOk() (*string, bool)`

GetPlanTypeOk returns a tuple with the PlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanType

`func (o *AccountStatusResponse) SetPlanType(v string)`

SetPlanType sets PlanType field to given value.


### GetPlanExpirationDays

`func (o *AccountStatusResponse) GetPlanExpirationDays() int32`

GetPlanExpirationDays returns the PlanExpirationDays field if non-nil, zero value otherwise.

### GetPlanExpirationDaysOk

`func (o *AccountStatusResponse) GetPlanExpirationDaysOk() (*int32, bool)`

GetPlanExpirationDaysOk returns a tuple with the PlanExpirationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanExpirationDays

`func (o *AccountStatusResponse) SetPlanExpirationDays(v int32)`

SetPlanExpirationDays sets PlanExpirationDays field to given value.

### HasPlanExpirationDays

`func (o *AccountStatusResponse) HasPlanExpirationDays() bool`

HasPlanExpirationDays returns a boolean if a field has been set.

### GetApplicationUse

`func (o *AccountStatusResponse) GetApplicationUse() string`

GetApplicationUse returns the ApplicationUse field if non-nil, zero value otherwise.

### GetApplicationUseOk

`func (o *AccountStatusResponse) GetApplicationUseOk() (*string, bool)`

GetApplicationUseOk returns a tuple with the ApplicationUse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUse

`func (o *AccountStatusResponse) SetApplicationUse(v string)`

SetApplicationUse sets ApplicationUse field to given value.


### GetAccountActivated

`func (o *AccountStatusResponse) GetAccountActivated() bool`

GetAccountActivated returns the AccountActivated field if non-nil, zero value otherwise.

### GetAccountActivatedOk

`func (o *AccountStatusResponse) GetAccountActivatedOk() (*bool, bool)`

GetAccountActivatedOk returns a tuple with the AccountActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountActivated

`func (o *AccountStatusResponse) SetAccountActivated(v bool)`

SetAccountActivated sets AccountActivated field to given value.

### HasAccountActivated

`func (o *AccountStatusResponse) HasAccountActivated() bool`

HasAccountActivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# PlanUpdateEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailId** | **string** | email id on which support team will contact on | 
**PhoneNumber** | Pointer to **string** | contact number on which support team can call user | [optional] 
**BillingFrequency** | Pointer to **string** | The frequency with with the customer needs to be billed at. The current supported values are Monthly and Annually | [optional] 
**Baselines** | [**SelfServiceCreditsBaselines**](SelfServiceCreditsBaselines.md) |  | 
**Details** | Pointer to **string** | option details the user might want to inform | [optional] 

## Methods

### NewPlanUpdateEmail

`func NewPlanUpdateEmail(emailId string, baselines SelfServiceCreditsBaselines, ) *PlanUpdateEmail`

NewPlanUpdateEmail instantiates a new PlanUpdateEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanUpdateEmailWithDefaults

`func NewPlanUpdateEmailWithDefaults() *PlanUpdateEmail`

NewPlanUpdateEmailWithDefaults instantiates a new PlanUpdateEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailId

`func (o *PlanUpdateEmail) GetEmailId() string`

GetEmailId returns the EmailId field if non-nil, zero value otherwise.

### GetEmailIdOk

`func (o *PlanUpdateEmail) GetEmailIdOk() (*string, bool)`

GetEmailIdOk returns a tuple with the EmailId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailId

`func (o *PlanUpdateEmail) SetEmailId(v string)`

SetEmailId sets EmailId field to given value.


### GetPhoneNumber

`func (o *PlanUpdateEmail) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PlanUpdateEmail) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PlanUpdateEmail) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PlanUpdateEmail) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetBillingFrequency

`func (o *PlanUpdateEmail) GetBillingFrequency() string`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *PlanUpdateEmail) GetBillingFrequencyOk() (*string, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *PlanUpdateEmail) SetBillingFrequency(v string)`

SetBillingFrequency sets BillingFrequency field to given value.

### HasBillingFrequency

`func (o *PlanUpdateEmail) HasBillingFrequency() bool`

HasBillingFrequency returns a boolean if a field has been set.

### GetBaselines

`func (o *PlanUpdateEmail) GetBaselines() SelfServiceCreditsBaselines`

GetBaselines returns the Baselines field if non-nil, zero value otherwise.

### GetBaselinesOk

`func (o *PlanUpdateEmail) GetBaselinesOk() (*SelfServiceCreditsBaselines, bool)`

GetBaselinesOk returns a tuple with the Baselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselines

`func (o *PlanUpdateEmail) SetBaselines(v SelfServiceCreditsBaselines)`

SetBaselines sets Baselines field to given value.


### GetDetails

`func (o *PlanUpdateEmail) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *PlanUpdateEmail) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *PlanUpdateEmail) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *PlanUpdateEmail) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



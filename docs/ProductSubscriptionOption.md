# ProductSubscriptionOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingFrequency** | **string** | Identifier for the plans billing term. Valid values are:  1. Monthly  2. Annually  | 
**DiscountPercentage** | **int32** | Discount percentage for this plan&#39;s subscription. | 

## Methods

### NewProductSubscriptionOption

`func NewProductSubscriptionOption(billingFrequency string, discountPercentage int32, ) *ProductSubscriptionOption`

NewProductSubscriptionOption instantiates a new ProductSubscriptionOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductSubscriptionOptionWithDefaults

`func NewProductSubscriptionOptionWithDefaults() *ProductSubscriptionOption`

NewProductSubscriptionOptionWithDefaults instantiates a new ProductSubscriptionOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingFrequency

`func (o *ProductSubscriptionOption) GetBillingFrequency() string`

GetBillingFrequency returns the BillingFrequency field if non-nil, zero value otherwise.

### GetBillingFrequencyOk

`func (o *ProductSubscriptionOption) GetBillingFrequencyOk() (*string, bool)`

GetBillingFrequencyOk returns a tuple with the BillingFrequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingFrequency

`func (o *ProductSubscriptionOption) SetBillingFrequency(v string)`

SetBillingFrequency sets BillingFrequency field to given value.


### GetDiscountPercentage

`func (o *ProductSubscriptionOption) GetDiscountPercentage() int32`

GetDiscountPercentage returns the DiscountPercentage field if non-nil, zero value otherwise.

### GetDiscountPercentageOk

`func (o *ProductSubscriptionOption) GetDiscountPercentageOk() (*int32, bool)`

GetDiscountPercentageOk returns a tuple with the DiscountPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscountPercentage

`func (o *ProductSubscriptionOption) SetDiscountPercentage(v int32)`

SetDiscountPercentage sets DiscountPercentage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



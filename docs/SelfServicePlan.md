# SelfServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | **string** | Unique identifier of the product in current plan. Valid values are: 1. &#x60;Free&#x60; 2. &#x60;Trial&#x60; 3. &#x60;Essentials&#x60; 4. &#x60;EnterpriseOps&#x60; 5. &#x60;EnterpriseSec&#x60; 6. &#x60;EnterpriseSuite&#x60;  | 
**ProductName** | **string** | Name for the product. | 
**ProductGroups** | [**[]ProductGroup**](ProductGroup.md) | A list of product group for preview. | 
**ProductSubscriptionOptions** | [**[]ProductSubscriptionOption**](ProductSubscriptionOption.md) | A list of product subscription option. | 

## Methods

### NewSelfServicePlan

`func NewSelfServicePlan(productId string, productName string, productGroups []ProductGroup, productSubscriptionOptions []ProductSubscriptionOption, ) *SelfServicePlan`

NewSelfServicePlan instantiates a new SelfServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfServicePlanWithDefaults

`func NewSelfServicePlanWithDefaults() *SelfServicePlan`

NewSelfServicePlanWithDefaults instantiates a new SelfServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *SelfServicePlan) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *SelfServicePlan) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *SelfServicePlan) SetProductId(v string)`

SetProductId sets ProductId field to given value.


### GetProductName

`func (o *SelfServicePlan) GetProductName() string`

GetProductName returns the ProductName field if non-nil, zero value otherwise.

### GetProductNameOk

`func (o *SelfServicePlan) GetProductNameOk() (*string, bool)`

GetProductNameOk returns a tuple with the ProductName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductName

`func (o *SelfServicePlan) SetProductName(v string)`

SetProductName sets ProductName field to given value.


### GetProductGroups

`func (o *SelfServicePlan) GetProductGroups() []ProductGroup`

GetProductGroups returns the ProductGroups field if non-nil, zero value otherwise.

### GetProductGroupsOk

`func (o *SelfServicePlan) GetProductGroupsOk() (*[]ProductGroup, bool)`

GetProductGroupsOk returns a tuple with the ProductGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroups

`func (o *SelfServicePlan) SetProductGroups(v []ProductGroup)`

SetProductGroups sets ProductGroups field to given value.


### GetProductSubscriptionOptions

`func (o *SelfServicePlan) GetProductSubscriptionOptions() []ProductSubscriptionOption`

GetProductSubscriptionOptions returns the ProductSubscriptionOptions field if non-nil, zero value otherwise.

### GetProductSubscriptionOptionsOk

`func (o *SelfServicePlan) GetProductSubscriptionOptionsOk() (*[]ProductSubscriptionOption, bool)`

GetProductSubscriptionOptionsOk returns a tuple with the ProductSubscriptionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductSubscriptionOptions

`func (o *SelfServicePlan) SetProductSubscriptionOptions(v []ProductSubscriptionOption)`

SetProductSubscriptionOptions sets ProductSubscriptionOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



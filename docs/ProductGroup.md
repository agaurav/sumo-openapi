# ProductGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductGroupName** | **string** | Name of the Product group:  | 
**ProductVariables** | [**[]ProductVariable**](ProductVariable.md) | Different product variables of the product group | 

## Methods

### NewProductGroup

`func NewProductGroup(productGroupName string, productVariables []ProductVariable, ) *ProductGroup`

NewProductGroup instantiates a new ProductGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductGroupWithDefaults

`func NewProductGroupWithDefaults() *ProductGroup`

NewProductGroupWithDefaults instantiates a new ProductGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductGroupName

`func (o *ProductGroup) GetProductGroupName() string`

GetProductGroupName returns the ProductGroupName field if non-nil, zero value otherwise.

### GetProductGroupNameOk

`func (o *ProductGroup) GetProductGroupNameOk() (*string, bool)`

GetProductGroupNameOk returns a tuple with the ProductGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductGroupName

`func (o *ProductGroup) SetProductGroupName(v string)`

SetProductGroupName sets ProductGroupName field to given value.


### GetProductVariables

`func (o *ProductGroup) GetProductVariables() []ProductVariable`

GetProductVariables returns the ProductVariables field if non-nil, zero value otherwise.

### GetProductVariablesOk

`func (o *ProductGroup) GetProductVariablesOk() (*[]ProductVariable, bool)`

GetProductVariablesOk returns a tuple with the ProductVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductVariables

`func (o *ProductGroup) SetProductVariables(v []ProductVariable)`

SetProductVariables sets ProductVariables field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



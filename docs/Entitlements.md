# Entitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractType** | **string** | Details of the contract type. &#x60;AnnualBucket&#x60; are contracts that buy and consume ingest on yearly basis. &#x60;Credits&#x60; are contracts that buy a single unit called credits for all our features. &#x60;DailyAverage&#x60; are contracts that buy and consume ingest on a monthly basis. | 
**EntitlementType** | **string** | Text denoting the type of entitlement. - &#x60;continuous&#x60; for Continuous Analytics, - &#x60;frequent&#x60; for Frequent Analytics, - &#x60;storage&#x60; for Total Storage, - &#x60;metrics&#x60; for Metrics. | 
**Label** | **string** | The label of an entitlement is the plan name displayed on the accounts page in our user interface. | 
**Capacity** | [**Capacity**](Capacity.md) |  | 
**Capacities** | Pointer to [**[]Capacity**](Capacity.md) | Contains the capacities that were part of the contract. | [optional] 

## Methods

### NewEntitlements

`func NewEntitlements(contractType string, entitlementType string, label string, capacity Capacity, ) *Entitlements`

NewEntitlements instantiates a new Entitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsWithDefaults

`func NewEntitlementsWithDefaults() *Entitlements`

NewEntitlementsWithDefaults instantiates a new Entitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractType

`func (o *Entitlements) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *Entitlements) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *Entitlements) SetContractType(v string)`

SetContractType sets ContractType field to given value.


### GetEntitlementType

`func (o *Entitlements) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *Entitlements) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *Entitlements) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.


### GetLabel

`func (o *Entitlements) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Entitlements) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Entitlements) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCapacity

`func (o *Entitlements) GetCapacity() Capacity`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *Entitlements) GetCapacityOk() (*Capacity, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *Entitlements) SetCapacity(v Capacity)`

SetCapacity sets Capacity field to given value.


### GetCapacities

`func (o *Entitlements) GetCapacities() []Capacity`

GetCapacities returns the Capacities field if non-nil, zero value otherwise.

### GetCapacitiesOk

`func (o *Entitlements) GetCapacitiesOk() (*[]Capacity, bool)`

GetCapacitiesOk returns a tuple with the Capacities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacities

`func (o *Entitlements) SetCapacities(v []Capacity)`

SetCapacities sets Capacities field to given value.

### HasCapacities

`func (o *Entitlements) HasCapacities() bool`

HasCapacities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



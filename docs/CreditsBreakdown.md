# CreditsBreakdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeploymentChargeCredits** | **float64** | The total credits deducted from the parent organization in the form of deployment charge. | 
**AllocatedCredits** | **float64** | The total useable credits allocated to the child organization. | 

## Methods

### NewCreditsBreakdown

`func NewCreditsBreakdown(deploymentChargeCredits float64, allocatedCredits float64, ) *CreditsBreakdown`

NewCreditsBreakdown instantiates a new CreditsBreakdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreditsBreakdownWithDefaults

`func NewCreditsBreakdownWithDefaults() *CreditsBreakdown`

NewCreditsBreakdownWithDefaults instantiates a new CreditsBreakdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeploymentChargeCredits

`func (o *CreditsBreakdown) GetDeploymentChargeCredits() float64`

GetDeploymentChargeCredits returns the DeploymentChargeCredits field if non-nil, zero value otherwise.

### GetDeploymentChargeCreditsOk

`func (o *CreditsBreakdown) GetDeploymentChargeCreditsOk() (*float64, bool)`

GetDeploymentChargeCreditsOk returns a tuple with the DeploymentChargeCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentChargeCredits

`func (o *CreditsBreakdown) SetDeploymentChargeCredits(v float64)`

SetDeploymentChargeCredits sets DeploymentChargeCredits field to given value.


### GetAllocatedCredits

`func (o *CreditsBreakdown) GetAllocatedCredits() float64`

GetAllocatedCredits returns the AllocatedCredits field if non-nil, zero value otherwise.

### GetAllocatedCreditsOk

`func (o *CreditsBreakdown) GetAllocatedCreditsOk() (*float64, bool)`

GetAllocatedCreditsOk returns a tuple with the AllocatedCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocatedCredits

`func (o *CreditsBreakdown) SetAllocatedCredits(v float64)`

SetAllocatedCredits sets AllocatedCredits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



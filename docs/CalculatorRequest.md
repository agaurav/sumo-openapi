# CalculatorRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentDeploymentId** | Pointer to **string** | Identifier of the deployment in which the parent org is present. | [optional] 
**DeploymentId** | Pointer to **string** | Identifier of the deployment in which the child org is present. | [optional] 
**TrialPlanPeriod** | Pointer to **int32** | length of the trial period. | [optional] 
**Baselines** | Pointer to [**Baselines**](Baselines.md) |  | [optional] 

## Methods

### NewCalculatorRequest

`func NewCalculatorRequest() *CalculatorRequest`

NewCalculatorRequest instantiates a new CalculatorRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCalculatorRequestWithDefaults

`func NewCalculatorRequestWithDefaults() *CalculatorRequest`

NewCalculatorRequestWithDefaults instantiates a new CalculatorRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentDeploymentId

`func (o *CalculatorRequest) GetParentDeploymentId() string`

GetParentDeploymentId returns the ParentDeploymentId field if non-nil, zero value otherwise.

### GetParentDeploymentIdOk

`func (o *CalculatorRequest) GetParentDeploymentIdOk() (*string, bool)`

GetParentDeploymentIdOk returns a tuple with the ParentDeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDeploymentId

`func (o *CalculatorRequest) SetParentDeploymentId(v string)`

SetParentDeploymentId sets ParentDeploymentId field to given value.

### HasParentDeploymentId

`func (o *CalculatorRequest) HasParentDeploymentId() bool`

HasParentDeploymentId returns a boolean if a field has been set.

### GetDeploymentId

`func (o *CalculatorRequest) GetDeploymentId() string`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *CalculatorRequest) GetDeploymentIdOk() (*string, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *CalculatorRequest) SetDeploymentId(v string)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *CalculatorRequest) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetTrialPlanPeriod

`func (o *CalculatorRequest) GetTrialPlanPeriod() int32`

GetTrialPlanPeriod returns the TrialPlanPeriod field if non-nil, zero value otherwise.

### GetTrialPlanPeriodOk

`func (o *CalculatorRequest) GetTrialPlanPeriodOk() (*int32, bool)`

GetTrialPlanPeriodOk returns a tuple with the TrialPlanPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialPlanPeriod

`func (o *CalculatorRequest) SetTrialPlanPeriod(v int32)`

SetTrialPlanPeriod sets TrialPlanPeriod field to given value.

### HasTrialPlanPeriod

`func (o *CalculatorRequest) HasTrialPlanPeriod() bool`

HasTrialPlanPeriod returns a boolean if a field has been set.

### GetBaselines

`func (o *CalculatorRequest) GetBaselines() Baselines`

GetBaselines returns the Baselines field if non-nil, zero value otherwise.

### GetBaselinesOk

`func (o *CalculatorRequest) GetBaselinesOk() (*Baselines, bool)`

GetBaselinesOk returns a tuple with the Baselines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaselines

`func (o *CalculatorRequest) SetBaselines(v Baselines)`

SetBaselines sets Baselines field to given value.

### HasBaselines

`func (o *CalculatorRequest) HasBaselines() bool`

HasBaselines returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



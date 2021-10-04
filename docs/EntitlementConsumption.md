# EntitlementConsumption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementType** | **string** | String value denoting the type of entitlement. - &#x60;continuous&#x60; for Continuous Analytics, - &#x60;frequent&#x60; for Frequent Analytics, - &#x60;storage&#x60; for Total Storage, - &#x60;metrics&#x60; for Metrics. | 
**Datapoints** | Pointer to [**[]DataPoints**](DataPoints.md) | Array of data points of the entitlement with their respective date range. | [optional] 
**Operators** | [**[]Operator**](Operator.md) | Operators used on the data. Available operators are &#x60;sum&#x60;, &#x60;average&#x60;, &#x60;usagePercentage&#x60;, &#x60;forecastValue&#x60;, &#x60;forecastPercentage&#x60;, and &#x60;forecastRemainingDays&#x60;. sum - Returns the sum of the usages. average - Returns the average of the usages. usagePercentage - Returns percentage of total capacity used for the startDate and endDate. forecastValue - Returns expected usage value assuming current usage behavior continues. forecastPercentage - Returns expected percentage usage by the endDate assuming current usage behavior continues. forecastRemainingDays- Returns the number of expected days, from today, that consumption will last assuming current usage behavior continues. | 
**ContractType** | **string** | Consumption model of the entitlements, available values are &#x60;DailyAverage&#x60;, &#x60;AnnualBucket&#x60;, and &#x60;Credits&#x60;. | 

## Methods

### NewEntitlementConsumption

`func NewEntitlementConsumption(entitlementType string, operators []Operator, contractType string, ) *EntitlementConsumption`

NewEntitlementConsumption instantiates a new EntitlementConsumption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementConsumptionWithDefaults

`func NewEntitlementConsumptionWithDefaults() *EntitlementConsumption`

NewEntitlementConsumptionWithDefaults instantiates a new EntitlementConsumption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementType

`func (o *EntitlementConsumption) GetEntitlementType() string`

GetEntitlementType returns the EntitlementType field if non-nil, zero value otherwise.

### GetEntitlementTypeOk

`func (o *EntitlementConsumption) GetEntitlementTypeOk() (*string, bool)`

GetEntitlementTypeOk returns a tuple with the EntitlementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementType

`func (o *EntitlementConsumption) SetEntitlementType(v string)`

SetEntitlementType sets EntitlementType field to given value.


### GetDatapoints

`func (o *EntitlementConsumption) GetDatapoints() []DataPoints`

GetDatapoints returns the Datapoints field if non-nil, zero value otherwise.

### GetDatapointsOk

`func (o *EntitlementConsumption) GetDatapointsOk() (*[]DataPoints, bool)`

GetDatapointsOk returns a tuple with the Datapoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatapoints

`func (o *EntitlementConsumption) SetDatapoints(v []DataPoints)`

SetDatapoints sets Datapoints field to given value.

### HasDatapoints

`func (o *EntitlementConsumption) HasDatapoints() bool`

HasDatapoints returns a boolean if a field has been set.

### GetOperators

`func (o *EntitlementConsumption) GetOperators() []Operator`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *EntitlementConsumption) GetOperatorsOk() (*[]Operator, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *EntitlementConsumption) SetOperators(v []Operator)`

SetOperators sets Operators field to given value.


### GetContractType

`func (o *EntitlementConsumption) GetContractType() string`

GetContractType returns the ContractType field if non-nil, zero value otherwise.

### GetContractTypeOk

`func (o *EntitlementConsumption) GetContractTypeOk() (*string, bool)`

GetContractTypeOk returns a tuple with the ContractType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractType

`func (o *EntitlementConsumption) SetContractType(v string)`

SetContractType sets ContractType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



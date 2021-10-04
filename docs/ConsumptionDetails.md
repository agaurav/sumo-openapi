# ConsumptionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntitlementConsumptions** | [**[]EntitlementConsumption**](EntitlementConsumption.md) | An array of entitlements. | 
**StartDate** | **string** | Start date of the data usage. | 
**EndDate** | **string** | End date of the data usage. | 

## Methods

### NewConsumptionDetails

`func NewConsumptionDetails(entitlementConsumptions []EntitlementConsumption, startDate string, endDate string, ) *ConsumptionDetails`

NewConsumptionDetails instantiates a new ConsumptionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConsumptionDetailsWithDefaults

`func NewConsumptionDetailsWithDefaults() *ConsumptionDetails`

NewConsumptionDetailsWithDefaults instantiates a new ConsumptionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlementConsumptions

`func (o *ConsumptionDetails) GetEntitlementConsumptions() []EntitlementConsumption`

GetEntitlementConsumptions returns the EntitlementConsumptions field if non-nil, zero value otherwise.

### GetEntitlementConsumptionsOk

`func (o *ConsumptionDetails) GetEntitlementConsumptionsOk() (*[]EntitlementConsumption, bool)`

GetEntitlementConsumptionsOk returns a tuple with the EntitlementConsumptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementConsumptions

`func (o *ConsumptionDetails) SetEntitlementConsumptions(v []EntitlementConsumption)`

SetEntitlementConsumptions sets EntitlementConsumptions field to given value.


### GetStartDate

`func (o *ConsumptionDetails) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ConsumptionDetails) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ConsumptionDetails) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.


### GetEndDate

`func (o *ConsumptionDetails) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ConsumptionDetails) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ConsumptionDetails) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



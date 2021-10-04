# TotalCredits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TotalCredits** | **float64** | Numerical value of the amount of credits | 
**Breakdown** | Pointer to [**CreditsBreakdown**](CreditsBreakdown.md) |  | [optional] 

## Methods

### NewTotalCredits

`func NewTotalCredits(totalCredits float64, ) *TotalCredits`

NewTotalCredits instantiates a new TotalCredits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotalCreditsWithDefaults

`func NewTotalCreditsWithDefaults() *TotalCredits`

NewTotalCreditsWithDefaults instantiates a new TotalCredits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotalCredits

`func (o *TotalCredits) GetTotalCredits() float64`

GetTotalCredits returns the TotalCredits field if non-nil, zero value otherwise.

### GetTotalCreditsOk

`func (o *TotalCredits) GetTotalCreditsOk() (*float64, bool)`

GetTotalCreditsOk returns a tuple with the TotalCredits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCredits

`func (o *TotalCredits) SetTotalCredits(v float64)`

SetTotalCredits sets TotalCredits field to given value.


### GetBreakdown

`func (o *TotalCredits) GetBreakdown() CreditsBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *TotalCredits) GetBreakdownOk() (*CreditsBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *TotalCredits) SetBreakdown(v CreditsBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *TotalCredits) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# OutlierDataValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Baseline** | Pointer to [**OutlierBound**](OutlierBound.md) |  | [optional] 
**Critical** | Pointer to [**OutlierBound**](OutlierBound.md) |  | [optional] 
**Warning** | Pointer to [**OutlierBound**](OutlierBound.md) |  | [optional] 
**Value** | Pointer to **float64** | The value of outlier data point. | [optional] 
**Violation** | Pointer to **string** | The type of violation. | [optional] 

## Methods

### NewOutlierDataValue

`func NewOutlierDataValue() *OutlierDataValue`

NewOutlierDataValue instantiates a new OutlierDataValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutlierDataValueWithDefaults

`func NewOutlierDataValueWithDefaults() *OutlierDataValue`

NewOutlierDataValueWithDefaults instantiates a new OutlierDataValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseline

`func (o *OutlierDataValue) GetBaseline() OutlierBound`

GetBaseline returns the Baseline field if non-nil, zero value otherwise.

### GetBaselineOk

`func (o *OutlierDataValue) GetBaselineOk() (*OutlierBound, bool)`

GetBaselineOk returns a tuple with the Baseline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseline

`func (o *OutlierDataValue) SetBaseline(v OutlierBound)`

SetBaseline sets Baseline field to given value.

### HasBaseline

`func (o *OutlierDataValue) HasBaseline() bool`

HasBaseline returns a boolean if a field has been set.

### GetCritical

`func (o *OutlierDataValue) GetCritical() OutlierBound`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *OutlierDataValue) GetCriticalOk() (*OutlierBound, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *OutlierDataValue) SetCritical(v OutlierBound)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *OutlierDataValue) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetWarning

`func (o *OutlierDataValue) GetWarning() OutlierBound`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *OutlierDataValue) GetWarningOk() (*OutlierBound, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *OutlierDataValue) SetWarning(v OutlierBound)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *OutlierDataValue) HasWarning() bool`

HasWarning returns a boolean if a field has been set.

### GetValue

`func (o *OutlierDataValue) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *OutlierDataValue) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *OutlierDataValue) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *OutlierDataValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetViolation

`func (o *OutlierDataValue) GetViolation() string`

GetViolation returns the Violation field if non-nil, zero value otherwise.

### GetViolationOk

`func (o *OutlierDataValue) GetViolationOk() (*string, bool)`

GetViolationOk returns a tuple with the Violation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViolation

`func (o *OutlierDataValue) SetViolation(v string)`

SetViolation sets Violation field to given value.

### HasViolation

`func (o *OutlierDataValue) HasViolation() bool`

HasViolation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



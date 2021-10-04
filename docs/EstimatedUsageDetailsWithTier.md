# EstimatedUsageDetailsWithTier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | Pointer to **string** | Name of the data tier. Supported Values are Continuous, Frequent, Infrequent | [optional] 
**DataScannedInBytes** | Pointer to **int64** | Amount of data scanned in bytes, to run the query. | [optional] 

## Methods

### NewEstimatedUsageDetailsWithTier

`func NewEstimatedUsageDetailsWithTier() *EstimatedUsageDetailsWithTier`

NewEstimatedUsageDetailsWithTier instantiates a new EstimatedUsageDetailsWithTier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEstimatedUsageDetailsWithTierWithDefaults

`func NewEstimatedUsageDetailsWithTierWithDefaults() *EstimatedUsageDetailsWithTier`

NewEstimatedUsageDetailsWithTierWithDefaults instantiates a new EstimatedUsageDetailsWithTier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *EstimatedUsageDetailsWithTier) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *EstimatedUsageDetailsWithTier) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *EstimatedUsageDetailsWithTier) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *EstimatedUsageDetailsWithTier) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetDataScannedInBytes

`func (o *EstimatedUsageDetailsWithTier) GetDataScannedInBytes() int64`

GetDataScannedInBytes returns the DataScannedInBytes field if non-nil, zero value otherwise.

### GetDataScannedInBytesOk

`func (o *EstimatedUsageDetailsWithTier) GetDataScannedInBytesOk() (*int64, bool)`

GetDataScannedInBytesOk returns a tuple with the DataScannedInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataScannedInBytes

`func (o *EstimatedUsageDetailsWithTier) SetDataScannedInBytes(v int64)`

SetDataScannedInBytes sets DataScannedInBytes field to given value.

### HasDataScannedInBytes

`func (o *EstimatedUsageDetailsWithTier) HasDataScannedInBytes() bool`

HasDataScannedInBytes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



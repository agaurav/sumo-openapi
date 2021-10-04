# IngestBudgetResourceIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IngestBudgetFieldValue** | Pointer to **string** | The unique field value of the ingest budget v1. This will be empty for v2 budgets. | [optional] [default to "Unknown"]
**Scope** | Pointer to **string** | The scope of the ingest budget v2. This will be empty for v1 budgets. | [optional] [default to "Unknown"]

## Methods

### NewIngestBudgetResourceIdentity

`func NewIngestBudgetResourceIdentity() *IngestBudgetResourceIdentity`

NewIngestBudgetResourceIdentity instantiates a new IngestBudgetResourceIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIngestBudgetResourceIdentityWithDefaults

`func NewIngestBudgetResourceIdentityWithDefaults() *IngestBudgetResourceIdentity`

NewIngestBudgetResourceIdentityWithDefaults instantiates a new IngestBudgetResourceIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIngestBudgetFieldValue

`func (o *IngestBudgetResourceIdentity) GetIngestBudgetFieldValue() string`

GetIngestBudgetFieldValue returns the IngestBudgetFieldValue field if non-nil, zero value otherwise.

### GetIngestBudgetFieldValueOk

`func (o *IngestBudgetResourceIdentity) GetIngestBudgetFieldValueOk() (*string, bool)`

GetIngestBudgetFieldValueOk returns a tuple with the IngestBudgetFieldValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngestBudgetFieldValue

`func (o *IngestBudgetResourceIdentity) SetIngestBudgetFieldValue(v string)`

SetIngestBudgetFieldValue sets IngestBudgetFieldValue field to given value.

### HasIngestBudgetFieldValue

`func (o *IngestBudgetResourceIdentity) HasIngestBudgetFieldValue() bool`

HasIngestBudgetFieldValue returns a boolean if a field has been set.

### GetScope

`func (o *IngestBudgetResourceIdentity) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *IngestBudgetResourceIdentity) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *IngestBudgetResourceIdentity) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *IngestBudgetResourceIdentity) HasScope() bool`

HasScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



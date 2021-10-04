# ListIngestBudgetsResponseV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]IngestBudgetV2**](IngestBudgetV2.md) | List of ingest budgets. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewListIngestBudgetsResponseV2

`func NewListIngestBudgetsResponseV2(data []IngestBudgetV2, ) *ListIngestBudgetsResponseV2`

NewListIngestBudgetsResponseV2 instantiates a new ListIngestBudgetsResponseV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIngestBudgetsResponseV2WithDefaults

`func NewListIngestBudgetsResponseV2WithDefaults() *ListIngestBudgetsResponseV2`

NewListIngestBudgetsResponseV2WithDefaults instantiates a new ListIngestBudgetsResponseV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListIngestBudgetsResponseV2) GetData() []IngestBudgetV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListIngestBudgetsResponseV2) GetDataOk() (*[]IngestBudgetV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListIngestBudgetsResponseV2) SetData(v []IngestBudgetV2)`

SetData sets Data field to given value.


### GetNext

`func (o *ListIngestBudgetsResponseV2) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListIngestBudgetsResponseV2) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListIngestBudgetsResponseV2) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListIngestBudgetsResponseV2) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



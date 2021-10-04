# ListDynamicRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]DynamicRule**](DynamicRule.md) | List of dynamic parsing rules. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewListDynamicRulesResponse

`func NewListDynamicRulesResponse(data []DynamicRule, ) *ListDynamicRulesResponse`

NewListDynamicRulesResponse instantiates a new ListDynamicRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDynamicRulesResponseWithDefaults

`func NewListDynamicRulesResponseWithDefaults() *ListDynamicRulesResponse`

NewListDynamicRulesResponseWithDefaults instantiates a new ListDynamicRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListDynamicRulesResponse) GetData() []DynamicRule`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListDynamicRulesResponse) GetDataOk() (*[]DynamicRule, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListDynamicRulesResponse) SetData(v []DynamicRule)`

SetData sets Data field to given value.


### GetNext

`func (o *ListDynamicRulesResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListDynamicRulesResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListDynamicRulesResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListDynamicRulesResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



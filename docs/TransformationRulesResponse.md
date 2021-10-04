# TransformationRulesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TransformationRuleResponse**](TransformationRuleResponse.md) | List of transformation rules. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewTransformationRulesResponse

`func NewTransformationRulesResponse(data []TransformationRuleResponse, ) *TransformationRulesResponse`

NewTransformationRulesResponse instantiates a new TransformationRulesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransformationRulesResponseWithDefaults

`func NewTransformationRulesResponseWithDefaults() *TransformationRulesResponse`

NewTransformationRulesResponseWithDefaults instantiates a new TransformationRulesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TransformationRulesResponse) GetData() []TransformationRuleResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransformationRulesResponse) GetDataOk() (*[]TransformationRuleResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransformationRulesResponse) SetData(v []TransformationRuleResponse)`

SetData sets Data field to given value.


### GetNext

`func (o *TransformationRulesResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *TransformationRulesResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *TransformationRulesResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *TransformationRulesResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



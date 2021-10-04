# ListAccessKeysResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AccessKeyPublic**](AccessKeyPublic.md) | An array of access keys. | 

## Methods

### NewListAccessKeysResult

`func NewListAccessKeysResult(data []AccessKeyPublic, ) *ListAccessKeysResult`

NewListAccessKeysResult instantiates a new ListAccessKeysResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccessKeysResultWithDefaults

`func NewListAccessKeysResultWithDefaults() *ListAccessKeysResult`

NewListAccessKeysResultWithDefaults instantiates a new ListAccessKeysResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListAccessKeysResult) GetData() []AccessKeyPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListAccessKeysResult) GetDataOk() (*[]AccessKeyPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListAccessKeysResult) SetData(v []AccessKeyPublic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



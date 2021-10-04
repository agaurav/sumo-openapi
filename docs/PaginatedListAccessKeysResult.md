# PaginatedListAccessKeysResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AccessKeyPublic**](AccessKeyPublic.md) | An array of access keys. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewPaginatedListAccessKeysResult

`func NewPaginatedListAccessKeysResult(data []AccessKeyPublic, ) *PaginatedListAccessKeysResult`

NewPaginatedListAccessKeysResult instantiates a new PaginatedListAccessKeysResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedListAccessKeysResultWithDefaults

`func NewPaginatedListAccessKeysResultWithDefaults() *PaginatedListAccessKeysResult`

NewPaginatedListAccessKeysResultWithDefaults instantiates a new PaginatedListAccessKeysResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PaginatedListAccessKeysResult) GetData() []AccessKeyPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaginatedListAccessKeysResult) GetDataOk() (*[]AccessKeyPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaginatedListAccessKeysResult) SetData(v []AccessKeyPublic)`

SetData sets Data field to given value.


### GetNext

`func (o *PaginatedListAccessKeysResult) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PaginatedListAccessKeysResult) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PaginatedListAccessKeysResult) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PaginatedListAccessKeysResult) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



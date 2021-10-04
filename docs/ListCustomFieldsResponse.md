# ListCustomFieldsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CustomField**](CustomField.md) | List of custom fields. | 

## Methods

### NewListCustomFieldsResponse

`func NewListCustomFieldsResponse(data []CustomField, ) *ListCustomFieldsResponse`

NewListCustomFieldsResponse instantiates a new ListCustomFieldsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomFieldsResponseWithDefaults

`func NewListCustomFieldsResponseWithDefaults() *ListCustomFieldsResponse`

NewListCustomFieldsResponseWithDefaults instantiates a new ListCustomFieldsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCustomFieldsResponse) GetData() []CustomField`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCustomFieldsResponse) GetDataOk() (*[]CustomField, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCustomFieldsResponse) SetData(v []CustomField)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



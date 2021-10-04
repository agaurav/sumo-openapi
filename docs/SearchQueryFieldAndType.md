# SearchQueryFieldAndType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** | Log field parsed from log search query. | [optional] 
**FieldType** | Pointer to **string** | The type of the field inferred from log results and explicit configuration. Valid values:   1. &#x60;NumericValue&#x60;: A field with a numerical type.   2. &#x60;DistinctCount&#x60;: A field with a dimensional type. | [optional] 
**IsImplicitField** | Pointer to **bool** | Indicates if the field is implicit or user defined. | [optional] 

## Methods

### NewSearchQueryFieldAndType

`func NewSearchQueryFieldAndType() *SearchQueryFieldAndType`

NewSearchQueryFieldAndType instantiates a new SearchQueryFieldAndType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchQueryFieldAndTypeWithDefaults

`func NewSearchQueryFieldAndTypeWithDefaults() *SearchQueryFieldAndType`

NewSearchQueryFieldAndTypeWithDefaults instantiates a new SearchQueryFieldAndType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *SearchQueryFieldAndType) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *SearchQueryFieldAndType) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *SearchQueryFieldAndType) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *SearchQueryFieldAndType) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldType

`func (o *SearchQueryFieldAndType) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *SearchQueryFieldAndType) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *SearchQueryFieldAndType) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.

### HasFieldType

`func (o *SearchQueryFieldAndType) HasFieldType() bool`

HasFieldType returns a boolean if a field has been set.

### GetIsImplicitField

`func (o *SearchQueryFieldAndType) GetIsImplicitField() bool`

GetIsImplicitField returns the IsImplicitField field if non-nil, zero value otherwise.

### GetIsImplicitFieldOk

`func (o *SearchQueryFieldAndType) GetIsImplicitFieldOk() (*bool, bool)`

GetIsImplicitFieldOk returns a tuple with the IsImplicitField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsImplicitField

`func (o *SearchQueryFieldAndType) SetIsImplicitField(v bool)`

SetIsImplicitField sets IsImplicitField field to given value.

### HasIsImplicitField

`func (o *SearchQueryFieldAndType) HasIsImplicitField() bool`

HasIsImplicitField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



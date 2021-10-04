# LookupTableField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | The name of the field. | 
**FieldType** | **string** | The data type of the field. Supported types:   - &#x60;boolean&#x60;   - &#x60;int&#x60;   - &#x60;long&#x60;   - &#x60;double&#x60;   - &#x60;string&#x60; | 

## Methods

### NewLookupTableField

`func NewLookupTableField(fieldName string, fieldType string, ) *LookupTableField`

NewLookupTableField instantiates a new LookupTableField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTableFieldWithDefaults

`func NewLookupTableFieldWithDefaults() *LookupTableField`

NewLookupTableFieldWithDefaults instantiates a new LookupTableField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *LookupTableField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *LookupTableField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *LookupTableField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldType

`func (o *LookupTableField) GetFieldType() string`

GetFieldType returns the FieldType field if non-nil, zero value otherwise.

### GetFieldTypeOk

`func (o *LookupTableField) GetFieldTypeOk() (*string, bool)`

GetFieldTypeOk returns a tuple with the FieldType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldType

`func (o *LookupTableField) SetFieldType(v string)`

SetFieldType sets FieldType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



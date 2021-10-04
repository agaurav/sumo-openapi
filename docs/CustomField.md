# CustomField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Field name. | 
**FieldId** | **string** | Identifier of the field. | 
**DataType** | **string** | Field type. Possible values are &#x60;String&#x60;, &#x60;Long&#x60;, &#x60;Int&#x60;, &#x60;Double&#x60;, and &#x60;Boolean&#x60;. | 
**State** | **string** | Indicates whether the field is enabled and its values are being accepted. Possible values are &#x60;Enabled&#x60; and &#x60;Disabled&#x60;. | 

## Methods

### NewCustomField

`func NewCustomField(fieldName string, fieldId string, dataType string, state string, ) *CustomField`

NewCustomField instantiates a new CustomField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldWithDefaults

`func NewCustomFieldWithDefaults() *CustomField`

NewCustomFieldWithDefaults instantiates a new CustomField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *CustomField) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *CustomField) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *CustomField) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldId

`func (o *CustomField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *CustomField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *CustomField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetDataType

`func (o *CustomField) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CustomField) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CustomField) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetState

`func (o *CustomField) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomField) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomField) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



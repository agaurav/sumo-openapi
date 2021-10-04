# BuiltinFieldUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | Field name. | 
**FieldId** | **string** | Identifier of the field. | 
**DataType** | **string** | Field type. Possible values are &#x60;String&#x60;, &#x60;Long&#x60;, &#x60;Int&#x60;, &#x60;Double&#x60;, &#x60;Boolean&#x60;. | 
**State** | **string** | Indicates whether the field is enabled and its values are being accepted. Possible values are &#x60;Enabled&#x60; and &#x60;Disabled&#x60;. | 
**FieldExtractionRules** | Pointer to **[]string** | An array of hexadecimal identifiers of field extraction rules which use this field. | [optional] 
**Roles** | Pointer to **[]string** | An array of hexadecimal identifiers of roles which use this field in the search filter. | [optional] 
**Partitions** | Pointer to **[]string** | An array of hexadecimal identifiers of partitions which use this field in the routing expression. | [optional] 

## Methods

### NewBuiltinFieldUsage

`func NewBuiltinFieldUsage(fieldName string, fieldId string, dataType string, state string, ) *BuiltinFieldUsage`

NewBuiltinFieldUsage instantiates a new BuiltinFieldUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuiltinFieldUsageWithDefaults

`func NewBuiltinFieldUsageWithDefaults() *BuiltinFieldUsage`

NewBuiltinFieldUsageWithDefaults instantiates a new BuiltinFieldUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *BuiltinFieldUsage) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *BuiltinFieldUsage) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *BuiltinFieldUsage) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldId

`func (o *BuiltinFieldUsage) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *BuiltinFieldUsage) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *BuiltinFieldUsage) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetDataType

`func (o *BuiltinFieldUsage) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *BuiltinFieldUsage) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *BuiltinFieldUsage) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetState

`func (o *BuiltinFieldUsage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BuiltinFieldUsage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BuiltinFieldUsage) SetState(v string)`

SetState sets State field to given value.


### GetFieldExtractionRules

`func (o *BuiltinFieldUsage) GetFieldExtractionRules() []string`

GetFieldExtractionRules returns the FieldExtractionRules field if non-nil, zero value otherwise.

### GetFieldExtractionRulesOk

`func (o *BuiltinFieldUsage) GetFieldExtractionRulesOk() (*[]string, bool)`

GetFieldExtractionRulesOk returns a tuple with the FieldExtractionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldExtractionRules

`func (o *BuiltinFieldUsage) SetFieldExtractionRules(v []string)`

SetFieldExtractionRules sets FieldExtractionRules field to given value.

### HasFieldExtractionRules

`func (o *BuiltinFieldUsage) HasFieldExtractionRules() bool`

HasFieldExtractionRules returns a boolean if a field has been set.

### GetRoles

`func (o *BuiltinFieldUsage) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BuiltinFieldUsage) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BuiltinFieldUsage) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BuiltinFieldUsage) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPartitions

`func (o *BuiltinFieldUsage) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *BuiltinFieldUsage) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *BuiltinFieldUsage) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *BuiltinFieldUsage) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



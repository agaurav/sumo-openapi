# BuiltinFieldUsageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** | Identifier of the field. | 
**DataType** | **string** | Field type. Possible values are &#x60;String&#x60;, &#x60;Long&#x60;, &#x60;Int&#x60;, &#x60;Double&#x60;, &#x60;Boolean&#x60;. | 
**State** | **string** | Indicates whether the field is enabled and its values are being accepted. Possible values are &#x60;Enabled&#x60; and &#x60;Disabled&#x60;. | 
**FieldExtractionRules** | Pointer to **[]string** | An array of hexadecimal identifiers of field extraction rules which use this field. | [optional] 
**Roles** | Pointer to **[]string** | An array of hexadecimal identifiers of roles which use this field in the search filter. | [optional] 
**Partitions** | Pointer to **[]string** | An array of hexadecimal identifiers of partitions which use this field in the routing expression. | [optional] 

## Methods

### NewBuiltinFieldUsageAllOf

`func NewBuiltinFieldUsageAllOf(fieldId string, dataType string, state string, ) *BuiltinFieldUsageAllOf`

NewBuiltinFieldUsageAllOf instantiates a new BuiltinFieldUsageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuiltinFieldUsageAllOfWithDefaults

`func NewBuiltinFieldUsageAllOfWithDefaults() *BuiltinFieldUsageAllOf`

NewBuiltinFieldUsageAllOfWithDefaults instantiates a new BuiltinFieldUsageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *BuiltinFieldUsageAllOf) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *BuiltinFieldUsageAllOf) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *BuiltinFieldUsageAllOf) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetDataType

`func (o *BuiltinFieldUsageAllOf) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *BuiltinFieldUsageAllOf) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *BuiltinFieldUsageAllOf) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetState

`func (o *BuiltinFieldUsageAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BuiltinFieldUsageAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BuiltinFieldUsageAllOf) SetState(v string)`

SetState sets State field to given value.


### GetFieldExtractionRules

`func (o *BuiltinFieldUsageAllOf) GetFieldExtractionRules() []string`

GetFieldExtractionRules returns the FieldExtractionRules field if non-nil, zero value otherwise.

### GetFieldExtractionRulesOk

`func (o *BuiltinFieldUsageAllOf) GetFieldExtractionRulesOk() (*[]string, bool)`

GetFieldExtractionRulesOk returns a tuple with the FieldExtractionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldExtractionRules

`func (o *BuiltinFieldUsageAllOf) SetFieldExtractionRules(v []string)`

SetFieldExtractionRules sets FieldExtractionRules field to given value.

### HasFieldExtractionRules

`func (o *BuiltinFieldUsageAllOf) HasFieldExtractionRules() bool`

HasFieldExtractionRules returns a boolean if a field has been set.

### GetRoles

`func (o *BuiltinFieldUsageAllOf) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *BuiltinFieldUsageAllOf) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *BuiltinFieldUsageAllOf) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *BuiltinFieldUsageAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPartitions

`func (o *BuiltinFieldUsageAllOf) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *BuiltinFieldUsageAllOf) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *BuiltinFieldUsageAllOf) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *BuiltinFieldUsageAllOf) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



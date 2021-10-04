# CustomFieldUsage

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
**CollectorsCount** | Pointer to **int32** | Total number of collectors using this field. | [optional] 
**SourcesCount** | Pointer to **int32** | Total number of sources using this field. | [optional] 

## Methods

### NewCustomFieldUsage

`func NewCustomFieldUsage(fieldName string, fieldId string, dataType string, state string, ) *CustomFieldUsage`

NewCustomFieldUsage instantiates a new CustomFieldUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldUsageWithDefaults

`func NewCustomFieldUsageWithDefaults() *CustomFieldUsage`

NewCustomFieldUsageWithDefaults instantiates a new CustomFieldUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *CustomFieldUsage) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *CustomFieldUsage) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *CustomFieldUsage) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetFieldId

`func (o *CustomFieldUsage) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *CustomFieldUsage) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *CustomFieldUsage) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetDataType

`func (o *CustomFieldUsage) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CustomFieldUsage) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CustomFieldUsage) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetState

`func (o *CustomFieldUsage) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomFieldUsage) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomFieldUsage) SetState(v string)`

SetState sets State field to given value.


### GetFieldExtractionRules

`func (o *CustomFieldUsage) GetFieldExtractionRules() []string`

GetFieldExtractionRules returns the FieldExtractionRules field if non-nil, zero value otherwise.

### GetFieldExtractionRulesOk

`func (o *CustomFieldUsage) GetFieldExtractionRulesOk() (*[]string, bool)`

GetFieldExtractionRulesOk returns a tuple with the FieldExtractionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldExtractionRules

`func (o *CustomFieldUsage) SetFieldExtractionRules(v []string)`

SetFieldExtractionRules sets FieldExtractionRules field to given value.

### HasFieldExtractionRules

`func (o *CustomFieldUsage) HasFieldExtractionRules() bool`

HasFieldExtractionRules returns a boolean if a field has been set.

### GetRoles

`func (o *CustomFieldUsage) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CustomFieldUsage) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CustomFieldUsage) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CustomFieldUsage) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPartitions

`func (o *CustomFieldUsage) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *CustomFieldUsage) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *CustomFieldUsage) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *CustomFieldUsage) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetCollectorsCount

`func (o *CustomFieldUsage) GetCollectorsCount() int32`

GetCollectorsCount returns the CollectorsCount field if non-nil, zero value otherwise.

### GetCollectorsCountOk

`func (o *CustomFieldUsage) GetCollectorsCountOk() (*int32, bool)`

GetCollectorsCountOk returns a tuple with the CollectorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorsCount

`func (o *CustomFieldUsage) SetCollectorsCount(v int32)`

SetCollectorsCount sets CollectorsCount field to given value.

### HasCollectorsCount

`func (o *CustomFieldUsage) HasCollectorsCount() bool`

HasCollectorsCount returns a boolean if a field has been set.

### GetSourcesCount

`func (o *CustomFieldUsage) GetSourcesCount() int32`

GetSourcesCount returns the SourcesCount field if non-nil, zero value otherwise.

### GetSourcesCountOk

`func (o *CustomFieldUsage) GetSourcesCountOk() (*int32, bool)`

GetSourcesCountOk returns a tuple with the SourcesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesCount

`func (o *CustomFieldUsage) SetSourcesCount(v int32)`

SetSourcesCount sets SourcesCount field to given value.

### HasSourcesCount

`func (o *CustomFieldUsage) HasSourcesCount() bool`

HasSourcesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



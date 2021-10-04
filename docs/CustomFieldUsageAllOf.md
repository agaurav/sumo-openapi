# CustomFieldUsageAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** | Identifier of the field. | 
**DataType** | **string** | Field type. Possible values are &#x60;String&#x60;, &#x60;Long&#x60;, &#x60;Int&#x60;, &#x60;Double&#x60;, &#x60;Boolean&#x60;. | 
**State** | **string** | Indicates whether the field is enabled and its values are being accepted. Possible values are &#x60;Enabled&#x60; and &#x60;Disabled&#x60;. | 
**FieldExtractionRules** | Pointer to **[]string** | An array of hexadecimal identifiers of field extraction rules which use this field. | [optional] 
**Roles** | Pointer to **[]string** | An array of hexadecimal identifiers of roles which use this field in the search filter. | [optional] 
**Partitions** | Pointer to **[]string** | An array of hexadecimal identifiers of partitions which use this field in the routing expression. | [optional] 
**CollectorsCount** | Pointer to **int32** | Total number of collectors using this field. | [optional] 
**SourcesCount** | Pointer to **int32** | Total number of sources using this field. | [optional] 

## Methods

### NewCustomFieldUsageAllOf

`func NewCustomFieldUsageAllOf(fieldId string, dataType string, state string, ) *CustomFieldUsageAllOf`

NewCustomFieldUsageAllOf instantiates a new CustomFieldUsageAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldUsageAllOfWithDefaults

`func NewCustomFieldUsageAllOfWithDefaults() *CustomFieldUsageAllOf`

NewCustomFieldUsageAllOfWithDefaults instantiates a new CustomFieldUsageAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *CustomFieldUsageAllOf) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *CustomFieldUsageAllOf) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *CustomFieldUsageAllOf) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetDataType

`func (o *CustomFieldUsageAllOf) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *CustomFieldUsageAllOf) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *CustomFieldUsageAllOf) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetState

`func (o *CustomFieldUsageAllOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CustomFieldUsageAllOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CustomFieldUsageAllOf) SetState(v string)`

SetState sets State field to given value.


### GetFieldExtractionRules

`func (o *CustomFieldUsageAllOf) GetFieldExtractionRules() []string`

GetFieldExtractionRules returns the FieldExtractionRules field if non-nil, zero value otherwise.

### GetFieldExtractionRulesOk

`func (o *CustomFieldUsageAllOf) GetFieldExtractionRulesOk() (*[]string, bool)`

GetFieldExtractionRulesOk returns a tuple with the FieldExtractionRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldExtractionRules

`func (o *CustomFieldUsageAllOf) SetFieldExtractionRules(v []string)`

SetFieldExtractionRules sets FieldExtractionRules field to given value.

### HasFieldExtractionRules

`func (o *CustomFieldUsageAllOf) HasFieldExtractionRules() bool`

HasFieldExtractionRules returns a boolean if a field has been set.

### GetRoles

`func (o *CustomFieldUsageAllOf) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *CustomFieldUsageAllOf) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *CustomFieldUsageAllOf) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *CustomFieldUsageAllOf) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetPartitions

`func (o *CustomFieldUsageAllOf) GetPartitions() []string`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *CustomFieldUsageAllOf) GetPartitionsOk() (*[]string, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *CustomFieldUsageAllOf) SetPartitions(v []string)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *CustomFieldUsageAllOf) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetCollectorsCount

`func (o *CustomFieldUsageAllOf) GetCollectorsCount() int32`

GetCollectorsCount returns the CollectorsCount field if non-nil, zero value otherwise.

### GetCollectorsCountOk

`func (o *CustomFieldUsageAllOf) GetCollectorsCountOk() (*int32, bool)`

GetCollectorsCountOk returns a tuple with the CollectorsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectorsCount

`func (o *CustomFieldUsageAllOf) SetCollectorsCount(v int32)`

SetCollectorsCount sets CollectorsCount field to given value.

### HasCollectorsCount

`func (o *CustomFieldUsageAllOf) HasCollectorsCount() bool`

HasCollectorsCount returns a boolean if a field has been set.

### GetSourcesCount

`func (o *CustomFieldUsageAllOf) GetSourcesCount() int32`

GetSourcesCount returns the SourcesCount field if non-nil, zero value otherwise.

### GetSourcesCountOk

`func (o *CustomFieldUsageAllOf) GetSourcesCountOk() (*int32, bool)`

GetSourcesCountOk returns a tuple with the SourcesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourcesCount

`func (o *CustomFieldUsageAllOf) SetSourcesCount(v int32)`

SetSourcesCount sets SourcesCount field to given value.

### HasSourcesCount

`func (o *CustomFieldUsageAllOf) HasSourcesCount() bool`

HasSourcesCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



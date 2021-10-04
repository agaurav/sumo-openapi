# ExtractionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the field extraction rule. Use a name that makes it easy to identify the rule. | 
**Scope** | **string** | Scope of the field extraction rule. This could be a sourceCategory, sourceHost, or any other metadata that describes the data you want to extract from. Think of the Scope as the first portion of an ad hoc search, before the first pipe ( | ). You&#39;ll use the Scope to run a search against the rule. | 
**ParseExpression** | **string** | Describes the fields to be parsed. | 
**Enabled** | Pointer to **bool** | Is the field extraction rule enabled. | [optional] [default to true]
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | Unique identifier for the field extraction rule. | 
**FieldNames** | Pointer to **[]string** | List of extracted fields from \&quot;parseExpression\&quot;. | [optional] 

## Methods

### NewExtractionRule

`func NewExtractionRule(name string, scope string, parseExpression string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, ) *ExtractionRule`

NewExtractionRule instantiates a new ExtractionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionRuleWithDefaults

`func NewExtractionRuleWithDefaults() *ExtractionRule`

NewExtractionRuleWithDefaults instantiates a new ExtractionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExtractionRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExtractionRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExtractionRule) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *ExtractionRule) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ExtractionRule) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ExtractionRule) SetScope(v string)`

SetScope sets Scope field to given value.


### GetParseExpression

`func (o *ExtractionRule) GetParseExpression() string`

GetParseExpression returns the ParseExpression field if non-nil, zero value otherwise.

### GetParseExpressionOk

`func (o *ExtractionRule) GetParseExpressionOk() (*string, bool)`

GetParseExpressionOk returns a tuple with the ParseExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseExpression

`func (o *ExtractionRule) SetParseExpression(v string)`

SetParseExpression sets ParseExpression field to given value.


### GetEnabled

`func (o *ExtractionRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ExtractionRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ExtractionRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ExtractionRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExtractionRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExtractionRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExtractionRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *ExtractionRule) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ExtractionRule) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ExtractionRule) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *ExtractionRule) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *ExtractionRule) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *ExtractionRule) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *ExtractionRule) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *ExtractionRule) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *ExtractionRule) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *ExtractionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtractionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtractionRule) SetId(v string)`

SetId sets Id field to given value.


### GetFieldNames

`func (o *ExtractionRule) GetFieldNames() []string`

GetFieldNames returns the FieldNames field if non-nil, zero value otherwise.

### GetFieldNamesOk

`func (o *ExtractionRule) GetFieldNamesOk() (*[]string, bool)`

GetFieldNamesOk returns a tuple with the FieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNames

`func (o *ExtractionRule) SetFieldNames(v []string)`

SetFieldNames sets FieldNames field to given value.

### HasFieldNames

`func (o *ExtractionRule) HasFieldNames() bool`

HasFieldNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ExtractionRuleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the field extraction rule. | 
**FieldNames** | Pointer to **[]string** | List of extracted fields from \&quot;parseExpression\&quot;. | [optional] 

## Methods

### NewExtractionRuleAllOf

`func NewExtractionRuleAllOf(id string, ) *ExtractionRuleAllOf`

NewExtractionRuleAllOf instantiates a new ExtractionRuleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtractionRuleAllOfWithDefaults

`func NewExtractionRuleAllOfWithDefaults() *ExtractionRuleAllOf`

NewExtractionRuleAllOfWithDefaults instantiates a new ExtractionRuleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExtractionRuleAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExtractionRuleAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExtractionRuleAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetFieldNames

`func (o *ExtractionRuleAllOf) GetFieldNames() []string`

GetFieldNames returns the FieldNames field if non-nil, zero value otherwise.

### GetFieldNamesOk

`func (o *ExtractionRuleAllOf) GetFieldNamesOk() (*[]string, bool)`

GetFieldNamesOk returns a tuple with the FieldNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldNames

`func (o *ExtractionRuleAllOf) SetFieldNames(v []string)`

SetFieldNames sets FieldNames field to given value.

### HasFieldNames

`func (o *ExtractionRuleAllOf) HasFieldNames() bool`

HasFieldNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



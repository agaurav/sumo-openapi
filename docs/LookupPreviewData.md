# LookupPreviewData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldProperties** | Pointer to [**[]PreviewLookupTableField**](PreviewLookupTableField.md) | The field properties of the lookup table. This includes the field name, field description, and an identifier associated with each field. | [optional] 
**FieldValueMapList** | Pointer to **[]map[string]string** | The data of the lookup table as a list of field identifiers mapped to their values. | [optional] 

## Methods

### NewLookupPreviewData

`func NewLookupPreviewData() *LookupPreviewData`

NewLookupPreviewData instantiates a new LookupPreviewData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupPreviewDataWithDefaults

`func NewLookupPreviewDataWithDefaults() *LookupPreviewData`

NewLookupPreviewDataWithDefaults instantiates a new LookupPreviewData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldProperties

`func (o *LookupPreviewData) GetFieldProperties() []PreviewLookupTableField`

GetFieldProperties returns the FieldProperties field if non-nil, zero value otherwise.

### GetFieldPropertiesOk

`func (o *LookupPreviewData) GetFieldPropertiesOk() (*[]PreviewLookupTableField, bool)`

GetFieldPropertiesOk returns a tuple with the FieldProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldProperties

`func (o *LookupPreviewData) SetFieldProperties(v []PreviewLookupTableField)`

SetFieldProperties sets FieldProperties field to given value.

### HasFieldProperties

`func (o *LookupPreviewData) HasFieldProperties() bool`

HasFieldProperties returns a boolean if a field has been set.

### GetFieldValueMapList

`func (o *LookupPreviewData) GetFieldValueMapList() []map[string]string`

GetFieldValueMapList returns the FieldValueMapList field if non-nil, zero value otherwise.

### GetFieldValueMapListOk

`func (o *LookupPreviewData) GetFieldValueMapListOk() (*[]map[string]string, bool)`

GetFieldValueMapListOk returns a tuple with the FieldValueMapList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldValueMapList

`func (o *LookupPreviewData) SetFieldValueMapList(v []map[string]string)`

SetFieldValueMapList sets FieldValueMapList field to given value.

### HasFieldValueMapList

`func (o *LookupPreviewData) HasFieldValueMapList() bool`

HasFieldValueMapList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



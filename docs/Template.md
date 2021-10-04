# Template

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateType** | **string** | The type of template. &#x60;DashboardTemplate&#x60; provides a snapshot view of the exported dashboard. &#x60;DashboardReportModeTemplate&#x60; provides a printer-friendly view of the exported dashboard. New templates may be supported in the future. | 

## Methods

### NewTemplate

`func NewTemplate(templateType string, ) *Template`

NewTemplate instantiates a new Template object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplateWithDefaults

`func NewTemplateWithDefaults() *Template`

NewTemplateWithDefaults instantiates a new Template object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateType

`func (o *Template) GetTemplateType() string`

GetTemplateType returns the TemplateType field if non-nil, zero value otherwise.

### GetTemplateTypeOk

`func (o *Template) GetTemplateTypeOk() (*string, bool)`

GetTemplateTypeOk returns a tuple with the TemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateType

`func (o *Template) SetTemplateType(v string)`

SetTemplateType sets TemplateType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



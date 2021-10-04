# ReportFilterSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | **string** | The name af the field being filtered on, as listed in PanelField. | 
**Label** | **string** | The name of the field being filtered on, as displayed to the user. | 
**DefaultValue** | Pointer to **string** | The default value of the parameter. | [optional] 
**FilterType** | **string** | Type of filter. Can only be &#x60;numeric&#x60; or &#x60;textbox&#x60;. | 
**Properties** | **string** | Visual settings for the panel. | 
**PanelIds** | **[]string** | A list of panel identifiers that the filter applies to. | 

## Methods

### NewReportFilterSyncDefinition

`func NewReportFilterSyncDefinition(fieldName string, label string, filterType string, properties string, panelIds []string, ) *ReportFilterSyncDefinition`

NewReportFilterSyncDefinition instantiates a new ReportFilterSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportFilterSyncDefinitionWithDefaults

`func NewReportFilterSyncDefinitionWithDefaults() *ReportFilterSyncDefinition`

NewReportFilterSyncDefinitionWithDefaults instantiates a new ReportFilterSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *ReportFilterSyncDefinition) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ReportFilterSyncDefinition) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ReportFilterSyncDefinition) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetLabel

`func (o *ReportFilterSyncDefinition) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ReportFilterSyncDefinition) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ReportFilterSyncDefinition) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDefaultValue

`func (o *ReportFilterSyncDefinition) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ReportFilterSyncDefinition) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ReportFilterSyncDefinition) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ReportFilterSyncDefinition) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetFilterType

`func (o *ReportFilterSyncDefinition) GetFilterType() string`

GetFilterType returns the FilterType field if non-nil, zero value otherwise.

### GetFilterTypeOk

`func (o *ReportFilterSyncDefinition) GetFilterTypeOk() (*string, bool)`

GetFilterTypeOk returns a tuple with the FilterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterType

`func (o *ReportFilterSyncDefinition) SetFilterType(v string)`

SetFilterType sets FilterType field to given value.


### GetProperties

`func (o *ReportFilterSyncDefinition) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ReportFilterSyncDefinition) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ReportFilterSyncDefinition) SetProperties(v string)`

SetProperties sets Properties field to given value.


### GetPanelIds

`func (o *ReportFilterSyncDefinition) GetPanelIds() []string`

GetPanelIds returns the PanelIds field if non-nil, zero value otherwise.

### GetPanelIdsOk

`func (o *ReportFilterSyncDefinition) GetPanelIdsOk() (*[]string, bool)`

GetPanelIdsOk returns a tuple with the PanelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelIds

`func (o *ReportFilterSyncDefinition) SetPanelIds(v []string)`

SetPanelIds sets PanelIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



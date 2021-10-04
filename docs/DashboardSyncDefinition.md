# DashboardSyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of the dashboard. | 
**DetailLevel** | **int32** | Supported values are:   - &#x60;1&#x60; for small   - &#x60;2&#x60; for medium   - &#x60;3&#x60; for large | 
**Properties** | **string** | Visual settings for the panel. | 
**Panels** | [**[]ReportPanelSyncDefinition**](ReportPanelSyncDefinition.md) | The panels of the dashboard. _Dashboard links are not supported._ | 
**Filters** | [**[]ReportFilterSyncDefinition**](ReportFilterSyncDefinition.md) | The filters for the dashboard. Filters allow you to control the amount of information displayed in your dashboards. | 

## Methods

### NewDashboardSyncDefinition

`func NewDashboardSyncDefinition(description string, detailLevel int32, properties string, panels []ReportPanelSyncDefinition, filters []ReportFilterSyncDefinition, ) *DashboardSyncDefinition`

NewDashboardSyncDefinition instantiates a new DashboardSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSyncDefinitionWithDefaults

`func NewDashboardSyncDefinitionWithDefaults() *DashboardSyncDefinition`

NewDashboardSyncDefinitionWithDefaults instantiates a new DashboardSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DashboardSyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardSyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardSyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDetailLevel

`func (o *DashboardSyncDefinition) GetDetailLevel() int32`

GetDetailLevel returns the DetailLevel field if non-nil, zero value otherwise.

### GetDetailLevelOk

`func (o *DashboardSyncDefinition) GetDetailLevelOk() (*int32, bool)`

GetDetailLevelOk returns a tuple with the DetailLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailLevel

`func (o *DashboardSyncDefinition) SetDetailLevel(v int32)`

SetDetailLevel sets DetailLevel field to given value.


### GetProperties

`func (o *DashboardSyncDefinition) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DashboardSyncDefinition) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DashboardSyncDefinition) SetProperties(v string)`

SetProperties sets Properties field to given value.


### GetPanels

`func (o *DashboardSyncDefinition) GetPanels() []ReportPanelSyncDefinition`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *DashboardSyncDefinition) GetPanelsOk() (*[]ReportPanelSyncDefinition, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *DashboardSyncDefinition) SetPanels(v []ReportPanelSyncDefinition)`

SetPanels sets Panels field to given value.


### GetFilters

`func (o *DashboardSyncDefinition) GetFilters() []ReportFilterSyncDefinition`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DashboardSyncDefinition) GetFiltersOk() (*[]ReportFilterSyncDefinition, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DashboardSyncDefinition) SetFilters(v []ReportFilterSyncDefinition)`

SetFilters sets Filters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



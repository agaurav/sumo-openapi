# DashboardV2SyncDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A description of the dashboard. | [optional] 
**Title** | **string** | The title of the dashboard. | 
**RootPanel** | Pointer to [**ContainerPanel**](ContainerPanel.md) |  | [optional] 
**Theme** | Pointer to **string** | Theme for the dashboard. Must be &#x60;light&#x60; or &#x60;dark&#x60;. | [optional] [default to "light"]
**TopologyLabelMap** | Pointer to [**TopologyLabelMap**](TopologyLabelMap.md) |  | [optional] 
**RefreshInterval** | Pointer to **int32** | Interval of time (in seconds) to automatically refresh the dashboard. A value of 0 means we never automatically refresh the dashboard. | [optional] 
**TimeRange** | Pointer to [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | [optional] 
**Layout** | Pointer to [**Layout**](Layout.md) |  | [optional] 
**Panels** | Pointer to [**[]Panel**](Panel.md) | Children panels that the container panel contains. | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables that could be applied to the panel&#39;s children. | [optional] 
**ColoringRules** | Pointer to [**[]ColoringRule**](ColoringRule.md) | Coloring rules to color the panel/data with. | [optional] 

## Methods

### NewDashboardV2SyncDefinition

`func NewDashboardV2SyncDefinition(title string, ) *DashboardV2SyncDefinition`

NewDashboardV2SyncDefinition instantiates a new DashboardV2SyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardV2SyncDefinitionWithDefaults

`func NewDashboardV2SyncDefinitionWithDefaults() *DashboardV2SyncDefinition`

NewDashboardV2SyncDefinitionWithDefaults instantiates a new DashboardV2SyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DashboardV2SyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardV2SyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardV2SyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardV2SyncDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *DashboardV2SyncDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardV2SyncDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardV2SyncDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRootPanel

`func (o *DashboardV2SyncDefinition) GetRootPanel() ContainerPanel`

GetRootPanel returns the RootPanel field if non-nil, zero value otherwise.

### GetRootPanelOk

`func (o *DashboardV2SyncDefinition) GetRootPanelOk() (*ContainerPanel, bool)`

GetRootPanelOk returns a tuple with the RootPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPanel

`func (o *DashboardV2SyncDefinition) SetRootPanel(v ContainerPanel)`

SetRootPanel sets RootPanel field to given value.

### HasRootPanel

`func (o *DashboardV2SyncDefinition) HasRootPanel() bool`

HasRootPanel returns a boolean if a field has been set.

### GetTheme

`func (o *DashboardV2SyncDefinition) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *DashboardV2SyncDefinition) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *DashboardV2SyncDefinition) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *DashboardV2SyncDefinition) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetTopologyLabelMap

`func (o *DashboardV2SyncDefinition) GetTopologyLabelMap() TopologyLabelMap`

GetTopologyLabelMap returns the TopologyLabelMap field if non-nil, zero value otherwise.

### GetTopologyLabelMapOk

`func (o *DashboardV2SyncDefinition) GetTopologyLabelMapOk() (*TopologyLabelMap, bool)`

GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyLabelMap

`func (o *DashboardV2SyncDefinition) SetTopologyLabelMap(v TopologyLabelMap)`

SetTopologyLabelMap sets TopologyLabelMap field to given value.

### HasTopologyLabelMap

`func (o *DashboardV2SyncDefinition) HasTopologyLabelMap() bool`

HasTopologyLabelMap returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *DashboardV2SyncDefinition) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *DashboardV2SyncDefinition) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *DashboardV2SyncDefinition) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *DashboardV2SyncDefinition) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetTimeRange

`func (o *DashboardV2SyncDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardV2SyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardV2SyncDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *DashboardV2SyncDefinition) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetLayout

`func (o *DashboardV2SyncDefinition) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *DashboardV2SyncDefinition) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *DashboardV2SyncDefinition) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *DashboardV2SyncDefinition) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPanels

`func (o *DashboardV2SyncDefinition) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *DashboardV2SyncDefinition) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *DashboardV2SyncDefinition) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *DashboardV2SyncDefinition) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetVariables

`func (o *DashboardV2SyncDefinition) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DashboardV2SyncDefinition) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DashboardV2SyncDefinition) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DashboardV2SyncDefinition) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetColoringRules

`func (o *DashboardV2SyncDefinition) GetColoringRules() []ColoringRule`

GetColoringRules returns the ColoringRules field if non-nil, zero value otherwise.

### GetColoringRulesOk

`func (o *DashboardV2SyncDefinition) GetColoringRulesOk() (*[]ColoringRule, bool)`

GetColoringRulesOk returns a tuple with the ColoringRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColoringRules

`func (o *DashboardV2SyncDefinition) SetColoringRules(v []ColoringRule)`

SetColoringRules sets ColoringRules field to given value.

### HasColoringRules

`func (o *DashboardV2SyncDefinition) HasColoringRules() bool`

HasColoringRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



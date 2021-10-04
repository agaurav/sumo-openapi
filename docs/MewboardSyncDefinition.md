# MewboardSyncDefinition

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

### NewMewboardSyncDefinition

`func NewMewboardSyncDefinition(title string, ) *MewboardSyncDefinition`

NewMewboardSyncDefinition instantiates a new MewboardSyncDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMewboardSyncDefinitionWithDefaults

`func NewMewboardSyncDefinitionWithDefaults() *MewboardSyncDefinition`

NewMewboardSyncDefinitionWithDefaults instantiates a new MewboardSyncDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MewboardSyncDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MewboardSyncDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MewboardSyncDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MewboardSyncDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *MewboardSyncDefinition) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MewboardSyncDefinition) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MewboardSyncDefinition) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRootPanel

`func (o *MewboardSyncDefinition) GetRootPanel() ContainerPanel`

GetRootPanel returns the RootPanel field if non-nil, zero value otherwise.

### GetRootPanelOk

`func (o *MewboardSyncDefinition) GetRootPanelOk() (*ContainerPanel, bool)`

GetRootPanelOk returns a tuple with the RootPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPanel

`func (o *MewboardSyncDefinition) SetRootPanel(v ContainerPanel)`

SetRootPanel sets RootPanel field to given value.

### HasRootPanel

`func (o *MewboardSyncDefinition) HasRootPanel() bool`

HasRootPanel returns a boolean if a field has been set.

### GetTheme

`func (o *MewboardSyncDefinition) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MewboardSyncDefinition) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MewboardSyncDefinition) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MewboardSyncDefinition) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetTopologyLabelMap

`func (o *MewboardSyncDefinition) GetTopologyLabelMap() TopologyLabelMap`

GetTopologyLabelMap returns the TopologyLabelMap field if non-nil, zero value otherwise.

### GetTopologyLabelMapOk

`func (o *MewboardSyncDefinition) GetTopologyLabelMapOk() (*TopologyLabelMap, bool)`

GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyLabelMap

`func (o *MewboardSyncDefinition) SetTopologyLabelMap(v TopologyLabelMap)`

SetTopologyLabelMap sets TopologyLabelMap field to given value.

### HasTopologyLabelMap

`func (o *MewboardSyncDefinition) HasTopologyLabelMap() bool`

HasTopologyLabelMap returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *MewboardSyncDefinition) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *MewboardSyncDefinition) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *MewboardSyncDefinition) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *MewboardSyncDefinition) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetTimeRange

`func (o *MewboardSyncDefinition) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MewboardSyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MewboardSyncDefinition) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *MewboardSyncDefinition) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetLayout

`func (o *MewboardSyncDefinition) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *MewboardSyncDefinition) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *MewboardSyncDefinition) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *MewboardSyncDefinition) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPanels

`func (o *MewboardSyncDefinition) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *MewboardSyncDefinition) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *MewboardSyncDefinition) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *MewboardSyncDefinition) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetVariables

`func (o *MewboardSyncDefinition) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *MewboardSyncDefinition) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *MewboardSyncDefinition) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *MewboardSyncDefinition) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetColoringRules

`func (o *MewboardSyncDefinition) GetColoringRules() []ColoringRule`

GetColoringRules returns the ColoringRules field if non-nil, zero value otherwise.

### GetColoringRulesOk

`func (o *MewboardSyncDefinition) GetColoringRulesOk() (*[]ColoringRule, bool)`

GetColoringRulesOk returns a tuple with the ColoringRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColoringRules

`func (o *MewboardSyncDefinition) SetColoringRules(v []ColoringRule)`

SetColoringRules sets ColoringRules field to given value.

### HasColoringRules

`func (o *MewboardSyncDefinition) HasColoringRules() bool`

HasColoringRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



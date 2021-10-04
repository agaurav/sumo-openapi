# MewboardSyncDefinitionAllOf

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

### NewMewboardSyncDefinitionAllOf

`func NewMewboardSyncDefinitionAllOf(title string, ) *MewboardSyncDefinitionAllOf`

NewMewboardSyncDefinitionAllOf instantiates a new MewboardSyncDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMewboardSyncDefinitionAllOfWithDefaults

`func NewMewboardSyncDefinitionAllOfWithDefaults() *MewboardSyncDefinitionAllOf`

NewMewboardSyncDefinitionAllOfWithDefaults instantiates a new MewboardSyncDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *MewboardSyncDefinitionAllOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MewboardSyncDefinitionAllOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MewboardSyncDefinitionAllOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MewboardSyncDefinitionAllOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTitle

`func (o *MewboardSyncDefinitionAllOf) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MewboardSyncDefinitionAllOf) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MewboardSyncDefinitionAllOf) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetRootPanel

`func (o *MewboardSyncDefinitionAllOf) GetRootPanel() ContainerPanel`

GetRootPanel returns the RootPanel field if non-nil, zero value otherwise.

### GetRootPanelOk

`func (o *MewboardSyncDefinitionAllOf) GetRootPanelOk() (*ContainerPanel, bool)`

GetRootPanelOk returns a tuple with the RootPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootPanel

`func (o *MewboardSyncDefinitionAllOf) SetRootPanel(v ContainerPanel)`

SetRootPanel sets RootPanel field to given value.

### HasRootPanel

`func (o *MewboardSyncDefinitionAllOf) HasRootPanel() bool`

HasRootPanel returns a boolean if a field has been set.

### GetTheme

`func (o *MewboardSyncDefinitionAllOf) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *MewboardSyncDefinitionAllOf) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *MewboardSyncDefinitionAllOf) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *MewboardSyncDefinitionAllOf) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetTopologyLabelMap

`func (o *MewboardSyncDefinitionAllOf) GetTopologyLabelMap() TopologyLabelMap`

GetTopologyLabelMap returns the TopologyLabelMap field if non-nil, zero value otherwise.

### GetTopologyLabelMapOk

`func (o *MewboardSyncDefinitionAllOf) GetTopologyLabelMapOk() (*TopologyLabelMap, bool)`

GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyLabelMap

`func (o *MewboardSyncDefinitionAllOf) SetTopologyLabelMap(v TopologyLabelMap)`

SetTopologyLabelMap sets TopologyLabelMap field to given value.

### HasTopologyLabelMap

`func (o *MewboardSyncDefinitionAllOf) HasTopologyLabelMap() bool`

HasTopologyLabelMap returns a boolean if a field has been set.

### GetRefreshInterval

`func (o *MewboardSyncDefinitionAllOf) GetRefreshInterval() int32`

GetRefreshInterval returns the RefreshInterval field if non-nil, zero value otherwise.

### GetRefreshIntervalOk

`func (o *MewboardSyncDefinitionAllOf) GetRefreshIntervalOk() (*int32, bool)`

GetRefreshIntervalOk returns a tuple with the RefreshInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshInterval

`func (o *MewboardSyncDefinitionAllOf) SetRefreshInterval(v int32)`

SetRefreshInterval sets RefreshInterval field to given value.

### HasRefreshInterval

`func (o *MewboardSyncDefinitionAllOf) HasRefreshInterval() bool`

HasRefreshInterval returns a boolean if a field has been set.

### GetTimeRange

`func (o *MewboardSyncDefinitionAllOf) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *MewboardSyncDefinitionAllOf) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *MewboardSyncDefinitionAllOf) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.

### HasTimeRange

`func (o *MewboardSyncDefinitionAllOf) HasTimeRange() bool`

HasTimeRange returns a boolean if a field has been set.

### GetLayout

`func (o *MewboardSyncDefinitionAllOf) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *MewboardSyncDefinitionAllOf) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *MewboardSyncDefinitionAllOf) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *MewboardSyncDefinitionAllOf) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPanels

`func (o *MewboardSyncDefinitionAllOf) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *MewboardSyncDefinitionAllOf) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *MewboardSyncDefinitionAllOf) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *MewboardSyncDefinitionAllOf) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetVariables

`func (o *MewboardSyncDefinitionAllOf) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *MewboardSyncDefinitionAllOf) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *MewboardSyncDefinitionAllOf) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *MewboardSyncDefinitionAllOf) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetColoringRules

`func (o *MewboardSyncDefinitionAllOf) GetColoringRules() []ColoringRule`

GetColoringRules returns the ColoringRules field if non-nil, zero value otherwise.

### GetColoringRulesOk

`func (o *MewboardSyncDefinitionAllOf) GetColoringRulesOk() (*[]ColoringRule, bool)`

GetColoringRulesOk returns a tuple with the ColoringRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColoringRules

`func (o *MewboardSyncDefinitionAllOf) SetColoringRules(v []ColoringRule)`

SetColoringRules sets ColoringRules field to given value.

### HasColoringRules

`func (o *MewboardSyncDefinitionAllOf) HasColoringRules() bool`

HasColoringRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



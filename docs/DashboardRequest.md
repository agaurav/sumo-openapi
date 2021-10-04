# DashboardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Title of the dashboard. | 
**Description** | Pointer to **string** | Description of the dashboard. | [optional] 
**FolderId** | Pointer to **string** | The identifier of the folder to save the dashboard in. By default it is saved in your personal folder.  | [optional] 
**TopologyLabelMap** | Pointer to [**TopologyLabelMap**](TopologyLabelMap.md) |  | [optional] 
**Domain** | Pointer to **string** | If set denotes that the dashboard concerns a given domain (e.g. &#x60;aws&#x60;, &#x60;k8s&#x60;, &#x60;app&#x60;). | [optional] [default to ""]
**TimeRange** | [**ResolvableTimeRange**](ResolvableTimeRange.md) |  | 
**Panels** | Pointer to [**[]Panel**](Panel.md) | Panels in the dashboard. | [optional] 
**Layout** | Pointer to [**Layout**](Layout.md) |  | [optional] 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables to apply to the panels. | [optional] 
**Theme** | Pointer to **string** | Theme for the dashboard. Either &#x60;Light&#x60; or &#x60;Dark&#x60;. | [optional] [default to "Light"]

## Methods

### NewDashboardRequest

`func NewDashboardRequest(title string, timeRange ResolvableTimeRange, ) *DashboardRequest`

NewDashboardRequest instantiates a new DashboardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardRequestWithDefaults

`func NewDashboardRequestWithDefaults() *DashboardRequest`

NewDashboardRequestWithDefaults instantiates a new DashboardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *DashboardRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *DashboardRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *DashboardRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *DashboardRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DashboardRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DashboardRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DashboardRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFolderId

`func (o *DashboardRequest) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *DashboardRequest) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *DashboardRequest) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *DashboardRequest) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTopologyLabelMap

`func (o *DashboardRequest) GetTopologyLabelMap() TopologyLabelMap`

GetTopologyLabelMap returns the TopologyLabelMap field if non-nil, zero value otherwise.

### GetTopologyLabelMapOk

`func (o *DashboardRequest) GetTopologyLabelMapOk() (*TopologyLabelMap, bool)`

GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyLabelMap

`func (o *DashboardRequest) SetTopologyLabelMap(v TopologyLabelMap)`

SetTopologyLabelMap sets TopologyLabelMap field to given value.

### HasTopologyLabelMap

`func (o *DashboardRequest) HasTopologyLabelMap() bool`

HasTopologyLabelMap returns a boolean if a field has been set.

### GetDomain

`func (o *DashboardRequest) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DashboardRequest) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DashboardRequest) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DashboardRequest) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimeRange

`func (o *DashboardRequest) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *DashboardRequest) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *DashboardRequest) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetPanels

`func (o *DashboardRequest) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *DashboardRequest) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *DashboardRequest) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *DashboardRequest) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetLayout

`func (o *DashboardRequest) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *DashboardRequest) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *DashboardRequest) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *DashboardRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetVariables

`func (o *DashboardRequest) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DashboardRequest) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DashboardRequest) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DashboardRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetTheme

`func (o *DashboardRequest) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *DashboardRequest) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *DashboardRequest) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *DashboardRequest) HasTheme() bool`

HasTheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# Dashboard

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
**Id** | Pointer to **string** | Unique identifier for the dashboard. | [optional] 

## Methods

### NewDashboard

`func NewDashboard(title string, timeRange ResolvableTimeRange, ) *Dashboard`

NewDashboard instantiates a new Dashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardWithDefaults

`func NewDashboardWithDefaults() *Dashboard`

NewDashboardWithDefaults instantiates a new Dashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *Dashboard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Dashboard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Dashboard) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *Dashboard) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Dashboard) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Dashboard) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Dashboard) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFolderId

`func (o *Dashboard) GetFolderId() string`

GetFolderId returns the FolderId field if non-nil, zero value otherwise.

### GetFolderIdOk

`func (o *Dashboard) GetFolderIdOk() (*string, bool)`

GetFolderIdOk returns a tuple with the FolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolderId

`func (o *Dashboard) SetFolderId(v string)`

SetFolderId sets FolderId field to given value.

### HasFolderId

`func (o *Dashboard) HasFolderId() bool`

HasFolderId returns a boolean if a field has been set.

### GetTopologyLabelMap

`func (o *Dashboard) GetTopologyLabelMap() TopologyLabelMap`

GetTopologyLabelMap returns the TopologyLabelMap field if non-nil, zero value otherwise.

### GetTopologyLabelMapOk

`func (o *Dashboard) GetTopologyLabelMapOk() (*TopologyLabelMap, bool)`

GetTopologyLabelMapOk returns a tuple with the TopologyLabelMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopologyLabelMap

`func (o *Dashboard) SetTopologyLabelMap(v TopologyLabelMap)`

SetTopologyLabelMap sets TopologyLabelMap field to given value.

### HasTopologyLabelMap

`func (o *Dashboard) HasTopologyLabelMap() bool`

HasTopologyLabelMap returns a boolean if a field has been set.

### GetDomain

`func (o *Dashboard) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *Dashboard) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *Dashboard) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *Dashboard) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetTimeRange

`func (o *Dashboard) GetTimeRange() ResolvableTimeRange`

GetTimeRange returns the TimeRange field if non-nil, zero value otherwise.

### GetTimeRangeOk

`func (o *Dashboard) GetTimeRangeOk() (*ResolvableTimeRange, bool)`

GetTimeRangeOk returns a tuple with the TimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRange

`func (o *Dashboard) SetTimeRange(v ResolvableTimeRange)`

SetTimeRange sets TimeRange field to given value.


### GetPanels

`func (o *Dashboard) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *Dashboard) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *Dashboard) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *Dashboard) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetLayout

`func (o *Dashboard) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Dashboard) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Dashboard) SetLayout(v Layout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *Dashboard) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetVariables

`func (o *Dashboard) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *Dashboard) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *Dashboard) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *Dashboard) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetTheme

`func (o *Dashboard) GetTheme() string`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *Dashboard) GetThemeOk() (*string, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *Dashboard) SetTheme(v string)`

SetTheme sets Theme field to given value.

### HasTheme

`func (o *Dashboard) HasTheme() bool`

HasTheme returns a boolean if a field has been set.

### GetId

`func (o *Dashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Dashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Dashboard) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Dashboard) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



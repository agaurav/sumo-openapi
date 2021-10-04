# AppListItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemType** | **string** | Type of the item. Can be &#x60;Dashboard&#x60;, &#x60;Report&#x60;, &#x60;Search&#x60;, &#x60;ScheduledSearch&#x60;, &#x60;MetricsSearch&#x60; or &#x60;Folder&#x60;. | 
**Name** | **string** | Name of the item. | 
**Description** | Pointer to **string** | Description of the item. | [optional] 
**Query** | Pointer to **string** | Search query for the item. Applicable only for &#x60;Search&#x60;, &#x60;ScheduledSearch&#x60; and &#x60;MetricsSearch&#x60; itemType. | [optional] 
**ScreenshotUrl** | Pointer to **string** | URL for the screenshot of the item. Applicable only for &#x60;Dashboard&#x60; and &#x60;Report&#x60; itemType. | [optional] 
**Panels** | Pointer to [**[]PanelItem**](PanelItem.md) | Panels associated with the item. Applicable only for &#x60;Dashboard&#x60; and &#x60;Report&#x60; itemType. | [optional] 
**Children** | Pointer to [**[]AppListItem**](AppListItem.md) | Child content items. Applicable only for &#x60;Folder&#x60; itemType. | [optional] 

## Methods

### NewAppListItem

`func NewAppListItem(itemType string, name string, ) *AppListItem`

NewAppListItem instantiates a new AppListItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppListItemWithDefaults

`func NewAppListItemWithDefaults() *AppListItem`

NewAppListItemWithDefaults instantiates a new AppListItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemType

`func (o *AppListItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *AppListItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *AppListItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.


### GetName

`func (o *AppListItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppListItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppListItem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppListItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppListItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppListItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AppListItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuery

`func (o *AppListItem) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *AppListItem) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *AppListItem) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *AppListItem) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetScreenshotUrl

`func (o *AppListItem) GetScreenshotUrl() string`

GetScreenshotUrl returns the ScreenshotUrl field if non-nil, zero value otherwise.

### GetScreenshotUrlOk

`func (o *AppListItem) GetScreenshotUrlOk() (*string, bool)`

GetScreenshotUrlOk returns a tuple with the ScreenshotUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotUrl

`func (o *AppListItem) SetScreenshotUrl(v string)`

SetScreenshotUrl sets ScreenshotUrl field to given value.

### HasScreenshotUrl

`func (o *AppListItem) HasScreenshotUrl() bool`

HasScreenshotUrl returns a boolean if a field has been set.

### GetPanels

`func (o *AppListItem) GetPanels() []PanelItem`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *AppListItem) GetPanelsOk() (*[]PanelItem, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *AppListItem) SetPanels(v []PanelItem)`

SetPanels sets Panels field to given value.

### HasPanels

`func (o *AppListItem) HasPanels() bool`

HasPanels returns a boolean if a field has been set.

### GetChildren

`func (o *AppListItem) GetChildren() []AppListItem`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AppListItem) GetChildrenOk() (*[]AppListItem, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AppListItem) SetChildren(v []AppListItem)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *AppListItem) HasChildren() bool`

HasChildren returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



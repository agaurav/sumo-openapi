# Panel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the panel. | [optional] 
**Key** | **string** | Key for the panel. Used to create searches for the queries in the panel and configure the layout of the panel in the dashboard.  | 
**Title** | Pointer to **string** | Title of the panel. | [optional] 
**VisualSettings** | Pointer to **string** | Visual settings of the panel. | [optional] 
**KeepVisualSettingsConsistentWithParent** | Pointer to **bool** | Keeps the visual settings, like series colors, consistent with the settings of the parent panel. | [optional] [default to true]
**PanelType** | **string** | Type of panel. | 

## Methods

### NewPanel

`func NewPanel(key string, panelType string, ) *Panel`

NewPanel instantiates a new Panel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPanelWithDefaults

`func NewPanelWithDefaults() *Panel`

NewPanelWithDefaults instantiates a new Panel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Panel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Panel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Panel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Panel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *Panel) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Panel) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Panel) SetKey(v string)`

SetKey sets Key field to given value.


### GetTitle

`func (o *Panel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Panel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Panel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Panel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVisualSettings

`func (o *Panel) GetVisualSettings() string`

GetVisualSettings returns the VisualSettings field if non-nil, zero value otherwise.

### GetVisualSettingsOk

`func (o *Panel) GetVisualSettingsOk() (*string, bool)`

GetVisualSettingsOk returns a tuple with the VisualSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualSettings

`func (o *Panel) SetVisualSettings(v string)`

SetVisualSettings sets VisualSettings field to given value.

### HasVisualSettings

`func (o *Panel) HasVisualSettings() bool`

HasVisualSettings returns a boolean if a field has been set.

### GetKeepVisualSettingsConsistentWithParent

`func (o *Panel) GetKeepVisualSettingsConsistentWithParent() bool`

GetKeepVisualSettingsConsistentWithParent returns the KeepVisualSettingsConsistentWithParent field if non-nil, zero value otherwise.

### GetKeepVisualSettingsConsistentWithParentOk

`func (o *Panel) GetKeepVisualSettingsConsistentWithParentOk() (*bool, bool)`

GetKeepVisualSettingsConsistentWithParentOk returns a tuple with the KeepVisualSettingsConsistentWithParent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepVisualSettingsConsistentWithParent

`func (o *Panel) SetKeepVisualSettingsConsistentWithParent(v bool)`

SetKeepVisualSettingsConsistentWithParent sets KeepVisualSettingsConsistentWithParent field to given value.

### HasKeepVisualSettingsConsistentWithParent

`func (o *Panel) HasKeepVisualSettingsConsistentWithParent() bool`

HasKeepVisualSettingsConsistentWithParent returns a boolean if a field has been set.

### GetPanelType

`func (o *Panel) GetPanelType() string`

GetPanelType returns the PanelType field if non-nil, zero value otherwise.

### GetPanelTypeOk

`func (o *Panel) GetPanelTypeOk() (*string, bool)`

GetPanelTypeOk returns a tuple with the PanelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanelType

`func (o *Panel) SetPanelType(v string)`

SetPanelType sets PanelType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



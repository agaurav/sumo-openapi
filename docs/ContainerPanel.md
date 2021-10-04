# ContainerPanel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | [**Layout**](Layout.md) |  | 
**Panels** | [**[]Panel**](Panel.md) | Children panels that the container panel contains. | 
**Variables** | Pointer to [**[]Variable**](Variable.md) | Variables to apply to the panel children. | [optional] 
**ColoringRules** | Pointer to [**[]ColoringRule**](ColoringRule.md) | Rules to set the color of data. | [optional] 

## Methods

### NewContainerPanel

`func NewContainerPanel(layout Layout, panels []Panel, ) *ContainerPanel`

NewContainerPanel instantiates a new ContainerPanel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainerPanelWithDefaults

`func NewContainerPanelWithDefaults() *ContainerPanel`

NewContainerPanelWithDefaults instantiates a new ContainerPanel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *ContainerPanel) GetLayout() Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ContainerPanel) GetLayoutOk() (*Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ContainerPanel) SetLayout(v Layout)`

SetLayout sets Layout field to given value.


### GetPanels

`func (o *ContainerPanel) GetPanels() []Panel`

GetPanels returns the Panels field if non-nil, zero value otherwise.

### GetPanelsOk

`func (o *ContainerPanel) GetPanelsOk() (*[]Panel, bool)`

GetPanelsOk returns a tuple with the Panels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPanels

`func (o *ContainerPanel) SetPanels(v []Panel)`

SetPanels sets Panels field to given value.


### GetVariables

`func (o *ContainerPanel) GetVariables() []Variable`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ContainerPanel) GetVariablesOk() (*[]Variable, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ContainerPanel) SetVariables(v []Variable)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ContainerPanel) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetColoringRules

`func (o *ContainerPanel) GetColoringRules() []ColoringRule`

GetColoringRules returns the ColoringRules field if non-nil, zero value otherwise.

### GetColoringRulesOk

`func (o *ContainerPanel) GetColoringRulesOk() (*[]ColoringRule, bool)`

GetColoringRulesOk returns a tuple with the ColoringRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColoringRules

`func (o *ContainerPanel) SetColoringRules(v []ColoringRule)`

SetColoringRules sets ColoringRules field to given value.

### HasColoringRules

`func (o *ContainerPanel) HasColoringRules() bool`

HasColoringRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



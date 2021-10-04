# Layout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LayoutType** | **string** | The type of panel layout on the Dashboard. For example, Grid, Tabs, or Hierarchical. Currently supports &#x60;Grid&#x60; only. | 
**LayoutStructures** | [**[]LayoutStructure**](LayoutStructure.md) | Layout structures for the panel childen. | 

## Methods

### NewLayout

`func NewLayout(layoutType string, layoutStructures []LayoutStructure, ) *Layout`

NewLayout instantiates a new Layout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutWithDefaults

`func NewLayoutWithDefaults() *Layout`

NewLayoutWithDefaults instantiates a new Layout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayoutType

`func (o *Layout) GetLayoutType() string`

GetLayoutType returns the LayoutType field if non-nil, zero value otherwise.

### GetLayoutTypeOk

`func (o *Layout) GetLayoutTypeOk() (*string, bool)`

GetLayoutTypeOk returns a tuple with the LayoutType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutType

`func (o *Layout) SetLayoutType(v string)`

SetLayoutType sets LayoutType field to given value.


### GetLayoutStructures

`func (o *Layout) GetLayoutStructures() []LayoutStructure`

GetLayoutStructures returns the LayoutStructures field if non-nil, zero value otherwise.

### GetLayoutStructuresOk

`func (o *Layout) GetLayoutStructuresOk() (*[]LayoutStructure, bool)`

GetLayoutStructuresOk returns a tuple with the LayoutStructures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutStructures

`func (o *Layout) SetLayoutStructures(v []LayoutStructure)`

SetLayoutStructures sets LayoutStructures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



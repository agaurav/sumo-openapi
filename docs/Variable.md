# Variable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the variable. | [optional] 
**Name** | **string** | Name of the variable. The variable name is case-insensitive. Only alphanumeric, and underscores are allowed in the variable name.  | 
**DisplayName** | Pointer to **string** | Display name of the variable shown in the UI. If this field is empty, the name field will be used. The display name is case-insensitive. Only numbers, and underscores are allowed in the variable name. This field is not yet supported by the UI.  | [optional] 
**DefaultValue** | Pointer to **string** | Default value of the variable. | [optional] 
**SourceDefinition** | [**VariableSourceDefinition**](VariableSourceDefinition.md) |  | 
**AllowMultiSelect** | Pointer to **bool** | Allow multiple selections in the values dropdown. | [optional] [default to false]
**IncludeAllOption** | Pointer to **bool** | Include an \&quot;All\&quot; option at the top of the variable&#39;s values dropdown. | [optional] [default to true]
**HideFromUI** | Pointer to **bool** | Hide the variable in the dashboard UI. | [optional] [default to false]

## Methods

### NewVariable

`func NewVariable(name string, sourceDefinition VariableSourceDefinition, ) *Variable`

NewVariable instantiates a new Variable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableWithDefaults

`func NewVariableWithDefaults() *Variable`

NewVariableWithDefaults instantiates a new Variable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Variable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Variable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Variable) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Variable) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Variable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Variable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Variable) SetName(v string)`

SetName sets Name field to given value.


### GetDisplayName

`func (o *Variable) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Variable) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Variable) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Variable) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDefaultValue

`func (o *Variable) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *Variable) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *Variable) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *Variable) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetSourceDefinition

`func (o *Variable) GetSourceDefinition() VariableSourceDefinition`

GetSourceDefinition returns the SourceDefinition field if non-nil, zero value otherwise.

### GetSourceDefinitionOk

`func (o *Variable) GetSourceDefinitionOk() (*VariableSourceDefinition, bool)`

GetSourceDefinitionOk returns a tuple with the SourceDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDefinition

`func (o *Variable) SetSourceDefinition(v VariableSourceDefinition)`

SetSourceDefinition sets SourceDefinition field to given value.


### GetAllowMultiSelect

`func (o *Variable) GetAllowMultiSelect() bool`

GetAllowMultiSelect returns the AllowMultiSelect field if non-nil, zero value otherwise.

### GetAllowMultiSelectOk

`func (o *Variable) GetAllowMultiSelectOk() (*bool, bool)`

GetAllowMultiSelectOk returns a tuple with the AllowMultiSelect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiSelect

`func (o *Variable) SetAllowMultiSelect(v bool)`

SetAllowMultiSelect sets AllowMultiSelect field to given value.

### HasAllowMultiSelect

`func (o *Variable) HasAllowMultiSelect() bool`

HasAllowMultiSelect returns a boolean if a field has been set.

### GetIncludeAllOption

`func (o *Variable) GetIncludeAllOption() bool`

GetIncludeAllOption returns the IncludeAllOption field if non-nil, zero value otherwise.

### GetIncludeAllOptionOk

`func (o *Variable) GetIncludeAllOptionOk() (*bool, bool)`

GetIncludeAllOptionOk returns a tuple with the IncludeAllOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllOption

`func (o *Variable) SetIncludeAllOption(v bool)`

SetIncludeAllOption sets IncludeAllOption field to given value.

### HasIncludeAllOption

`func (o *Variable) HasIncludeAllOption() bool`

HasIncludeAllOption returns a boolean if a field has been set.

### GetHideFromUI

`func (o *Variable) GetHideFromUI() bool`

GetHideFromUI returns the HideFromUI field if non-nil, zero value otherwise.

### GetHideFromUIOk

`func (o *Variable) GetHideFromUIOk() (*bool, bool)`

GetHideFromUIOk returns a tuple with the HideFromUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFromUI

`func (o *Variable) SetHideFromUI(v bool)`

SetHideFromUI sets HideFromUI field to given value.

### HasHideFromUI

`func (o *Variable) HasHideFromUI() bool`

HasHideFromUI returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



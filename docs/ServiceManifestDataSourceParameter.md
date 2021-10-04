# ServiceManifestDataSourceParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParameterType** | **string** | Parameter type. | 
**ParameterId** | **string** | Parameter identifier. | 
**DataSourceType** | Pointer to **string** | Data source type. | [optional] 
**Label** | Pointer to **string** | Label. | [optional] 
**Description** | Pointer to **string** | Description. | [optional] 
**Example** | Pointer to **string** | Example. | [optional] 
**Hidden** | Pointer to **bool** | Should the UI display? | [optional] [default to false]

## Methods

### NewServiceManifestDataSourceParameter

`func NewServiceManifestDataSourceParameter(parameterType string, parameterId string, ) *ServiceManifestDataSourceParameter`

NewServiceManifestDataSourceParameter instantiates a new ServiceManifestDataSourceParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceManifestDataSourceParameterWithDefaults

`func NewServiceManifestDataSourceParameterWithDefaults() *ServiceManifestDataSourceParameter`

NewServiceManifestDataSourceParameterWithDefaults instantiates a new ServiceManifestDataSourceParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameterType

`func (o *ServiceManifestDataSourceParameter) GetParameterType() string`

GetParameterType returns the ParameterType field if non-nil, zero value otherwise.

### GetParameterTypeOk

`func (o *ServiceManifestDataSourceParameter) GetParameterTypeOk() (*string, bool)`

GetParameterTypeOk returns a tuple with the ParameterType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterType

`func (o *ServiceManifestDataSourceParameter) SetParameterType(v string)`

SetParameterType sets ParameterType field to given value.


### GetParameterId

`func (o *ServiceManifestDataSourceParameter) GetParameterId() string`

GetParameterId returns the ParameterId field if non-nil, zero value otherwise.

### GetParameterIdOk

`func (o *ServiceManifestDataSourceParameter) GetParameterIdOk() (*string, bool)`

GetParameterIdOk returns a tuple with the ParameterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterId

`func (o *ServiceManifestDataSourceParameter) SetParameterId(v string)`

SetParameterId sets ParameterId field to given value.


### GetDataSourceType

`func (o *ServiceManifestDataSourceParameter) GetDataSourceType() string`

GetDataSourceType returns the DataSourceType field if non-nil, zero value otherwise.

### GetDataSourceTypeOk

`func (o *ServiceManifestDataSourceParameter) GetDataSourceTypeOk() (*string, bool)`

GetDataSourceTypeOk returns a tuple with the DataSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceType

`func (o *ServiceManifestDataSourceParameter) SetDataSourceType(v string)`

SetDataSourceType sets DataSourceType field to given value.

### HasDataSourceType

`func (o *ServiceManifestDataSourceParameter) HasDataSourceType() bool`

HasDataSourceType returns a boolean if a field has been set.

### GetLabel

`func (o *ServiceManifestDataSourceParameter) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ServiceManifestDataSourceParameter) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ServiceManifestDataSourceParameter) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ServiceManifestDataSourceParameter) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetDescription

`func (o *ServiceManifestDataSourceParameter) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceManifestDataSourceParameter) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceManifestDataSourceParameter) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceManifestDataSourceParameter) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExample

`func (o *ServiceManifestDataSourceParameter) GetExample() string`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ServiceManifestDataSourceParameter) GetExampleOk() (*string, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ServiceManifestDataSourceParameter) SetExample(v string)`

SetExample sets Example field to given value.

### HasExample

`func (o *ServiceManifestDataSourceParameter) HasExample() bool`

HasExample returns a boolean if a field has been set.

### GetHidden

`func (o *ServiceManifestDataSourceParameter) GetHidden() bool`

GetHidden returns the Hidden field if non-nil, zero value otherwise.

### GetHiddenOk

`func (o *ServiceManifestDataSourceParameter) GetHiddenOk() (*bool, bool)`

GetHiddenOk returns a tuple with the Hidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHidden

`func (o *ServiceManifestDataSourceParameter) SetHidden(v bool)`

SetHidden sets Hidden field to given value.

### HasHidden

`func (o *ServiceManifestDataSourceParameter) HasHidden() bool`

HasHidden returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



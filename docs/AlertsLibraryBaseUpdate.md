# AlertsLibraryBaseUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the alert or folder. | 
**Description** | Pointer to **string** | The description of the alert or folder. | [optional] [default to ""]
**Version** | **int64** | The version of the alert or folder. | 
**Type** | **string** | Type of the object model. | 

## Methods

### NewAlertsLibraryBaseUpdate

`func NewAlertsLibraryBaseUpdate(name string, version int64, type_ string, ) *AlertsLibraryBaseUpdate`

NewAlertsLibraryBaseUpdate instantiates a new AlertsLibraryBaseUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsLibraryBaseUpdateWithDefaults

`func NewAlertsLibraryBaseUpdateWithDefaults() *AlertsLibraryBaseUpdate`

NewAlertsLibraryBaseUpdateWithDefaults instantiates a new AlertsLibraryBaseUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertsLibraryBaseUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertsLibraryBaseUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertsLibraryBaseUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertsLibraryBaseUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertsLibraryBaseUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertsLibraryBaseUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertsLibraryBaseUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *AlertsLibraryBaseUpdate) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AlertsLibraryBaseUpdate) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AlertsLibraryBaseUpdate) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetType

`func (o *AlertsLibraryBaseUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertsLibraryBaseUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertsLibraryBaseUpdate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



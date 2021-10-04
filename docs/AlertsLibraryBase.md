# AlertsLibraryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the alert or folder. | 
**Description** | Pointer to **string** | Description of the alert or folder. | [optional] [default to ""]
**Type** | **string** | Type of the object model. Valid values:   1) AlertsLibraryAlert   2) AlertsLibraryFolder | 
**IsLocked** | Pointer to **bool** | Locking/Unlocking requires the &#x60;LockAlerts&#x60; capability. Locked objects can only be &#x60;Localized&#x60;. Updating or moving requires unlocking the object. Locking/Unlocking recursively locks all of the objects children. All children of a locked object must be locked. | [optional] [default to false]

## Methods

### NewAlertsLibraryBase

`func NewAlertsLibraryBase(name string, type_ string, ) *AlertsLibraryBase`

NewAlertsLibraryBase instantiates a new AlertsLibraryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsLibraryBaseWithDefaults

`func NewAlertsLibraryBaseWithDefaults() *AlertsLibraryBase`

NewAlertsLibraryBaseWithDefaults instantiates a new AlertsLibraryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AlertsLibraryBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertsLibraryBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertsLibraryBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertsLibraryBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertsLibraryBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertsLibraryBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AlertsLibraryBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AlertsLibraryBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertsLibraryBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertsLibraryBase) SetType(v string)`

SetType sets Type field to given value.


### GetIsLocked

`func (o *AlertsLibraryBase) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AlertsLibraryBase) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AlertsLibraryBase) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AlertsLibraryBase) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



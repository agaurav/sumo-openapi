# MonitorsLibraryBaseUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the monitor or folder. | 
**Description** | Pointer to **string** | The description of the monitor or folder. | [optional] [default to ""]
**Version** | **int64** | The version of the monitor or folder. | 
**Type** | **string** | Type of the object model. | 

## Methods

### NewMonitorsLibraryBaseUpdate

`func NewMonitorsLibraryBaseUpdate(name string, version int64, type_ string, ) *MonitorsLibraryBaseUpdate`

NewMonitorsLibraryBaseUpdate instantiates a new MonitorsLibraryBaseUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorsLibraryBaseUpdateWithDefaults

`func NewMonitorsLibraryBaseUpdateWithDefaults() *MonitorsLibraryBaseUpdate`

NewMonitorsLibraryBaseUpdateWithDefaults instantiates a new MonitorsLibraryBaseUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitorsLibraryBaseUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorsLibraryBaseUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorsLibraryBaseUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MonitorsLibraryBaseUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorsLibraryBaseUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorsLibraryBaseUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorsLibraryBaseUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *MonitorsLibraryBaseUpdate) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MonitorsLibraryBaseUpdate) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MonitorsLibraryBaseUpdate) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetType

`func (o *MonitorsLibraryBaseUpdate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorsLibraryBaseUpdate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorsLibraryBaseUpdate) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



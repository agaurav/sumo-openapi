# MonitorsLibraryBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the monitor or folder. | 
**Description** | Pointer to **string** | Description of the monitor or folder. | [optional] [default to ""]
**Type** | **string** | Type of the object model. Valid values:   1) MonitorsLibraryMonitor   2) MonitorsLibraryFolder | 

## Methods

### NewMonitorsLibraryBase

`func NewMonitorsLibraryBase(name string, type_ string, ) *MonitorsLibraryBase`

NewMonitorsLibraryBase instantiates a new MonitorsLibraryBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorsLibraryBaseWithDefaults

`func NewMonitorsLibraryBaseWithDefaults() *MonitorsLibraryBase`

NewMonitorsLibraryBaseWithDefaults instantiates a new MonitorsLibraryBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MonitorsLibraryBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MonitorsLibraryBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MonitorsLibraryBase) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *MonitorsLibraryBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MonitorsLibraryBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MonitorsLibraryBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MonitorsLibraryBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *MonitorsLibraryBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MonitorsLibraryBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MonitorsLibraryBase) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



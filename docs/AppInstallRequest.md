# AppInstallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Preferred name of the app to be installed. This will be the name of the app in the selected installation folder. | 
**Description** | **string** | Preferred description of the app to be installed. This will be displayed as the app description in the selected installation folder. | 
**DestinationFolderId** | **string** | Identifier of the folder in which the app will be installed in hexadecimal format. | 
**DataSourceValues** | Pointer to **map[string]string** | Dictionary of properties specifying log-source name and value. | [optional] 

## Methods

### NewAppInstallRequest

`func NewAppInstallRequest(name string, description string, destinationFolderId string, ) *AppInstallRequest`

NewAppInstallRequest instantiates a new AppInstallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppInstallRequestWithDefaults

`func NewAppInstallRequestWithDefaults() *AppInstallRequest`

NewAppInstallRequestWithDefaults instantiates a new AppInstallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AppInstallRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppInstallRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppInstallRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppInstallRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppInstallRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppInstallRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDestinationFolderId

`func (o *AppInstallRequest) GetDestinationFolderId() string`

GetDestinationFolderId returns the DestinationFolderId field if non-nil, zero value otherwise.

### GetDestinationFolderIdOk

`func (o *AppInstallRequest) GetDestinationFolderIdOk() (*string, bool)`

GetDestinationFolderIdOk returns a tuple with the DestinationFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationFolderId

`func (o *AppInstallRequest) SetDestinationFolderId(v string)`

SetDestinationFolderId sets DestinationFolderId field to given value.


### GetDataSourceValues

`func (o *AppInstallRequest) GetDataSourceValues() map[string]string`

GetDataSourceValues returns the DataSourceValues field if non-nil, zero value otherwise.

### GetDataSourceValuesOk

`func (o *AppInstallRequest) GetDataSourceValuesOk() (*map[string]string, bool)`

GetDataSourceValuesOk returns a tuple with the DataSourceValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSourceValues

`func (o *AppInstallRequest) SetDataSourceValues(v map[string]string)`

SetDataSourceValues sets DataSourceValues field to given value.

### HasDataSourceValues

`func (o *AppInstallRequest) HasDataSourceValues() bool`

HasDataSourceValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



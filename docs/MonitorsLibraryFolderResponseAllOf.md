# MonitorsLibraryFolderResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** | Aggregated permission summary for the calling user. If detailed permission statements are required, please call list permissions endpoint. | 
**Children** | [**[]MonitorsLibraryBaseResponse**](MonitorsLibraryBaseResponse.md) | Children of the folder. NOTE: Permissions field will not be filled (empty list) for children. | 

## Methods

### NewMonitorsLibraryFolderResponseAllOf

`func NewMonitorsLibraryFolderResponseAllOf(permissions []string, children []MonitorsLibraryBaseResponse, ) *MonitorsLibraryFolderResponseAllOf`

NewMonitorsLibraryFolderResponseAllOf instantiates a new MonitorsLibraryFolderResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitorsLibraryFolderResponseAllOfWithDefaults

`func NewMonitorsLibraryFolderResponseAllOfWithDefaults() *MonitorsLibraryFolderResponseAllOf`

NewMonitorsLibraryFolderResponseAllOfWithDefaults instantiates a new MonitorsLibraryFolderResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *MonitorsLibraryFolderResponseAllOf) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *MonitorsLibraryFolderResponseAllOf) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *MonitorsLibraryFolderResponseAllOf) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetChildren

`func (o *MonitorsLibraryFolderResponseAllOf) GetChildren() []MonitorsLibraryBaseResponse`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *MonitorsLibraryFolderResponseAllOf) GetChildrenOk() (*[]MonitorsLibraryBaseResponse, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *MonitorsLibraryFolderResponseAllOf) SetChildren(v []MonitorsLibraryBaseResponse)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AlertsLibraryFolderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | [**[]AlertsLibraryBaseResponse**](AlertsLibraryBaseResponse.md) | Children of the folder. NOTE: Permissions field will not be filled (empty list) for children. | 

## Methods

### NewAlertsLibraryFolderResponse

`func NewAlertsLibraryFolderResponse(children []AlertsLibraryBaseResponse, ) *AlertsLibraryFolderResponse`

NewAlertsLibraryFolderResponse instantiates a new AlertsLibraryFolderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsLibraryFolderResponseWithDefaults

`func NewAlertsLibraryFolderResponseWithDefaults() *AlertsLibraryFolderResponse`

NewAlertsLibraryFolderResponseWithDefaults instantiates a new AlertsLibraryFolderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *AlertsLibraryFolderResponse) GetChildren() []AlertsLibraryBaseResponse`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *AlertsLibraryFolderResponse) GetChildrenOk() (*[]AlertsLibraryBaseResponse, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *AlertsLibraryFolderResponse) SetChildren(v []AlertsLibraryBaseResponse)`

SetChildren sets Children field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



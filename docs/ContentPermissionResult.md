# ContentPermissionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExplicitPermissions** | [**[]ContentPermissionAssignment**](ContentPermissionAssignment.md) | Explicitly assigned content permissions. | 
**ImplicitPermissions** | Pointer to [**[]ContentPermissionAssignment**](ContentPermissionAssignment.md) | Implicitly inherited content permissions. | [optional] 

## Methods

### NewContentPermissionResult

`func NewContentPermissionResult(explicitPermissions []ContentPermissionAssignment, ) *ContentPermissionResult`

NewContentPermissionResult instantiates a new ContentPermissionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPermissionResultWithDefaults

`func NewContentPermissionResultWithDefaults() *ContentPermissionResult`

NewContentPermissionResultWithDefaults instantiates a new ContentPermissionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplicitPermissions

`func (o *ContentPermissionResult) GetExplicitPermissions() []ContentPermissionAssignment`

GetExplicitPermissions returns the ExplicitPermissions field if non-nil, zero value otherwise.

### GetExplicitPermissionsOk

`func (o *ContentPermissionResult) GetExplicitPermissionsOk() (*[]ContentPermissionAssignment, bool)`

GetExplicitPermissionsOk returns a tuple with the ExplicitPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPermissions

`func (o *ContentPermissionResult) SetExplicitPermissions(v []ContentPermissionAssignment)`

SetExplicitPermissions sets ExplicitPermissions field to given value.


### GetImplicitPermissions

`func (o *ContentPermissionResult) GetImplicitPermissions() []ContentPermissionAssignment`

GetImplicitPermissions returns the ImplicitPermissions field if non-nil, zero value otherwise.

### GetImplicitPermissionsOk

`func (o *ContentPermissionResult) GetImplicitPermissionsOk() (*[]ContentPermissionAssignment, bool)`

GetImplicitPermissionsOk returns a tuple with the ImplicitPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImplicitPermissions

`func (o *ContentPermissionResult) SetImplicitPermissions(v []ContentPermissionAssignment)`

SetImplicitPermissions sets ImplicitPermissions field to given value.

### HasImplicitPermissions

`func (o *ContentPermissionResult) HasImplicitPermissions() bool`

HasImplicitPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



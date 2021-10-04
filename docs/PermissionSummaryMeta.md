# PermissionSummaryMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the permission. Example values are: &#x60;Read&#x60;, &#x60;Update&#x60;, &#x60;Create&#x60;, etc. | 
**IsInherited** | **bool** | A true value implies that the permission is inherited from some ancestors of the resource. A false value implies that the permission is explicitly assigned to the resource. | 
**IsExplicit** | **bool** | A true value implies that the permission is explicitly assigned to the resource. A false value implies that the permission is not explicitly assigned to the resource. | 
**IsRevoked** | **bool** | A true value implies that the capability required for this permission has been revoked. | 
**IsRecursive** | **bool** | A true value implies that the permission is recursively cascaded down to all the direct and indirect children of the resource. | 
**IsSystemDefined** | **bool** | A true value implies that the permission is defined by the system on the resource and can not be modified by the user. A false value implies that the permission is defined by the user on the resource and can be modified by the user. | 

## Methods

### NewPermissionSummaryMeta

`func NewPermissionSummaryMeta(name string, isInherited bool, isExplicit bool, isRevoked bool, isRecursive bool, isSystemDefined bool, ) *PermissionSummaryMeta`

NewPermissionSummaryMeta instantiates a new PermissionSummaryMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSummaryMetaWithDefaults

`func NewPermissionSummaryMetaWithDefaults() *PermissionSummaryMeta`

NewPermissionSummaryMetaWithDefaults instantiates a new PermissionSummaryMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PermissionSummaryMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionSummaryMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionSummaryMeta) SetName(v string)`

SetName sets Name field to given value.


### GetIsInherited

`func (o *PermissionSummaryMeta) GetIsInherited() bool`

GetIsInherited returns the IsInherited field if non-nil, zero value otherwise.

### GetIsInheritedOk

`func (o *PermissionSummaryMeta) GetIsInheritedOk() (*bool, bool)`

GetIsInheritedOk returns a tuple with the IsInherited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInherited

`func (o *PermissionSummaryMeta) SetIsInherited(v bool)`

SetIsInherited sets IsInherited field to given value.


### GetIsExplicit

`func (o *PermissionSummaryMeta) GetIsExplicit() bool`

GetIsExplicit returns the IsExplicit field if non-nil, zero value otherwise.

### GetIsExplicitOk

`func (o *PermissionSummaryMeta) GetIsExplicitOk() (*bool, bool)`

GetIsExplicitOk returns a tuple with the IsExplicit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsExplicit

`func (o *PermissionSummaryMeta) SetIsExplicit(v bool)`

SetIsExplicit sets IsExplicit field to given value.


### GetIsRevoked

`func (o *PermissionSummaryMeta) GetIsRevoked() bool`

GetIsRevoked returns the IsRevoked field if non-nil, zero value otherwise.

### GetIsRevokedOk

`func (o *PermissionSummaryMeta) GetIsRevokedOk() (*bool, bool)`

GetIsRevokedOk returns a tuple with the IsRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRevoked

`func (o *PermissionSummaryMeta) SetIsRevoked(v bool)`

SetIsRevoked sets IsRevoked field to given value.


### GetIsRecursive

`func (o *PermissionSummaryMeta) GetIsRecursive() bool`

GetIsRecursive returns the IsRecursive field if non-nil, zero value otherwise.

### GetIsRecursiveOk

`func (o *PermissionSummaryMeta) GetIsRecursiveOk() (*bool, bool)`

GetIsRecursiveOk returns a tuple with the IsRecursive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecursive

`func (o *PermissionSummaryMeta) SetIsRecursive(v bool)`

SetIsRecursive sets IsRecursive field to given value.


### GetIsSystemDefined

`func (o *PermissionSummaryMeta) GetIsSystemDefined() bool`

GetIsSystemDefined returns the IsSystemDefined field if non-nil, zero value otherwise.

### GetIsSystemDefinedOk

`func (o *PermissionSummaryMeta) GetIsSystemDefinedOk() (*bool, bool)`

GetIsSystemDefinedOk returns a tuple with the IsSystemDefined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystemDefined

`func (o *PermissionSummaryMeta) SetIsSystemDefined(v bool)`

SetIsSystemDefined sets IsSystemDefined field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



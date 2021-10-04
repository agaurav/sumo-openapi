# PermissionStatementDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | **[]string** | List of permissions. | 
**SubjectType** | **string** | Type of subject for the permission. Valid values are: &#x60;role&#x60;. | 
**SubjectId** | **string** | The identifier that belongs to the subject type chosen above. For e.g. if the subjectType is set to &#x60;role&#x60;, subjectId should be the identifier of a role. | 
**TargetId** | **string** | The identifier that belongs to the resource this permission assignment applies to. | 

## Methods

### NewPermissionStatementDefinition

`func NewPermissionStatementDefinition(permissions []string, subjectType string, subjectId string, targetId string, ) *PermissionStatementDefinition`

NewPermissionStatementDefinition instantiates a new PermissionStatementDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionStatementDefinitionWithDefaults

`func NewPermissionStatementDefinitionWithDefaults() *PermissionStatementDefinition`

NewPermissionStatementDefinitionWithDefaults instantiates a new PermissionStatementDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *PermissionStatementDefinition) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionStatementDefinition) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionStatementDefinition) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetSubjectType

`func (o *PermissionStatementDefinition) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *PermissionStatementDefinition) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *PermissionStatementDefinition) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectId

`func (o *PermissionStatementDefinition) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *PermissionStatementDefinition) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *PermissionStatementDefinition) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetTargetId

`func (o *PermissionStatementDefinition) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *PermissionStatementDefinition) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *PermissionStatementDefinition) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



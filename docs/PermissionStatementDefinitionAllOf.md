# PermissionStatementDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectType** | **string** | Type of subject for the permission. Valid values are: &#x60;role&#x60;. | 
**SubjectId** | **string** | The identifier that belongs to the subject type chosen above. For e.g. if the subjectType is set to &#x60;role&#x60;, subjectId should be the identifier of a role. | 
**TargetId** | **string** | The identifier that belongs to the resource this permission assignment applies to. | 

## Methods

### NewPermissionStatementDefinitionAllOf

`func NewPermissionStatementDefinitionAllOf(subjectType string, subjectId string, targetId string, ) *PermissionStatementDefinitionAllOf`

NewPermissionStatementDefinitionAllOf instantiates a new PermissionStatementDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionStatementDefinitionAllOfWithDefaults

`func NewPermissionStatementDefinitionAllOfWithDefaults() *PermissionStatementDefinitionAllOf`

NewPermissionStatementDefinitionAllOfWithDefaults instantiates a new PermissionStatementDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectType

`func (o *PermissionStatementDefinitionAllOf) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *PermissionStatementDefinitionAllOf) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *PermissionStatementDefinitionAllOf) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectId

`func (o *PermissionStatementDefinitionAllOf) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *PermissionStatementDefinitionAllOf) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *PermissionStatementDefinitionAllOf) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetTargetId

`func (o *PermissionStatementDefinitionAllOf) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *PermissionStatementDefinitionAllOf) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *PermissionStatementDefinitionAllOf) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



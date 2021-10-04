# PermissionSubject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectType** | **string** | Type of subject for the permission. Valid values are: &#x60;user&#x60; or &#x60;role&#x60;. | 
**SubjectId** | **string** | The identifier that belongs to the subject type chosen above. For e.g. if the subjectType is set to &#x60;user&#x60;, subjectId should be the identifier of a user (same goes for &#x60;role&#x60; subjectType). | 

## Methods

### NewPermissionSubject

`func NewPermissionSubject(subjectType string, subjectId string, ) *PermissionSubject`

NewPermissionSubject instantiates a new PermissionSubject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSubjectWithDefaults

`func NewPermissionSubjectWithDefaults() *PermissionSubject`

NewPermissionSubjectWithDefaults instantiates a new PermissionSubject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectType

`func (o *PermissionSubject) GetSubjectType() string`

GetSubjectType returns the SubjectType field if non-nil, zero value otherwise.

### GetSubjectTypeOk

`func (o *PermissionSubject) GetSubjectTypeOk() (*string, bool)`

GetSubjectTypeOk returns a tuple with the SubjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectType

`func (o *PermissionSubject) SetSubjectType(v string)`

SetSubjectType sets SubjectType field to given value.


### GetSubjectId

`func (o *PermissionSubject) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *PermissionSubject) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *PermissionSubject) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



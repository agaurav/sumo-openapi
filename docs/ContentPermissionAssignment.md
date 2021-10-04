# ContentPermissionAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PermissionName** | **string** | Content permission name. Valid values are: &#x60;View&#x60;, &#x60;GrantView&#x60;, &#x60;Edit&#x60;, &#x60;GrantEdit&#x60;, &#x60;Manage&#x60;, and &#x60;GrantManage&#x60;. | 
**SourceType** | **string** | Type of source for the permission. Valid values are: &#x60;user&#x60;, &#x60;role&#x60;, and &#x60;org&#x60;. | 
**SourceId** | **string** | An identifier that belongs to the source type chosen above. For e.g. if the sourceType is set to \&quot;user\&quot;, sourceId should be identifier of a user (same goes for &#x60;role&#x60; and &#x60;org&#x60; sourceType) | 
**ContentId** | **string** | Unique identifier for the content item. | 

## Methods

### NewContentPermissionAssignment

`func NewContentPermissionAssignment(permissionName string, sourceType string, sourceId string, contentId string, ) *ContentPermissionAssignment`

NewContentPermissionAssignment instantiates a new ContentPermissionAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPermissionAssignmentWithDefaults

`func NewContentPermissionAssignmentWithDefaults() *ContentPermissionAssignment`

NewContentPermissionAssignmentWithDefaults instantiates a new ContentPermissionAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissionName

`func (o *ContentPermissionAssignment) GetPermissionName() string`

GetPermissionName returns the PermissionName field if non-nil, zero value otherwise.

### GetPermissionNameOk

`func (o *ContentPermissionAssignment) GetPermissionNameOk() (*string, bool)`

GetPermissionNameOk returns a tuple with the PermissionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionName

`func (o *ContentPermissionAssignment) SetPermissionName(v string)`

SetPermissionName sets PermissionName field to given value.


### GetSourceType

`func (o *ContentPermissionAssignment) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ContentPermissionAssignment) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ContentPermissionAssignment) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.


### GetSourceId

`func (o *ContentPermissionAssignment) GetSourceId() string`

GetSourceId returns the SourceId field if non-nil, zero value otherwise.

### GetSourceIdOk

`func (o *ContentPermissionAssignment) GetSourceIdOk() (*string, bool)`

GetSourceIdOk returns a tuple with the SourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceId

`func (o *ContentPermissionAssignment) SetSourceId(v string)`

SetSourceId sets SourceId field to given value.


### GetContentId

`func (o *ContentPermissionAssignment) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *ContentPermissionAssignment) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *ContentPermissionAssignment) SetContentId(v string)`

SetContentId sets ContentId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



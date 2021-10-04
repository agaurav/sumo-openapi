# ContentPermissionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentPermissionAssignments** | [**[]ContentPermissionAssignment**](ContentPermissionAssignment.md) | Content permissions to be updated. | 
**NotifyRecipients** | **bool** | Set this to \&quot;true\&quot; to notify the users who had a permission update. | 
**NotificationMessage** | **string** | The notification message sent to the users who had a permission update. | 

## Methods

### NewContentPermissionUpdateRequest

`func NewContentPermissionUpdateRequest(contentPermissionAssignments []ContentPermissionAssignment, notifyRecipients bool, notificationMessage string, ) *ContentPermissionUpdateRequest`

NewContentPermissionUpdateRequest instantiates a new ContentPermissionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentPermissionUpdateRequestWithDefaults

`func NewContentPermissionUpdateRequestWithDefaults() *ContentPermissionUpdateRequest`

NewContentPermissionUpdateRequestWithDefaults instantiates a new ContentPermissionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentPermissionAssignments

`func (o *ContentPermissionUpdateRequest) GetContentPermissionAssignments() []ContentPermissionAssignment`

GetContentPermissionAssignments returns the ContentPermissionAssignments field if non-nil, zero value otherwise.

### GetContentPermissionAssignmentsOk

`func (o *ContentPermissionUpdateRequest) GetContentPermissionAssignmentsOk() (*[]ContentPermissionAssignment, bool)`

GetContentPermissionAssignmentsOk returns a tuple with the ContentPermissionAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPermissionAssignments

`func (o *ContentPermissionUpdateRequest) SetContentPermissionAssignments(v []ContentPermissionAssignment)`

SetContentPermissionAssignments sets ContentPermissionAssignments field to given value.


### GetNotifyRecipients

`func (o *ContentPermissionUpdateRequest) GetNotifyRecipients() bool`

GetNotifyRecipients returns the NotifyRecipients field if non-nil, zero value otherwise.

### GetNotifyRecipientsOk

`func (o *ContentPermissionUpdateRequest) GetNotifyRecipientsOk() (*bool, bool)`

GetNotifyRecipientsOk returns a tuple with the NotifyRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyRecipients

`func (o *ContentPermissionUpdateRequest) SetNotifyRecipients(v bool)`

SetNotifyRecipients sets NotifyRecipients field to given value.


### GetNotificationMessage

`func (o *ContentPermissionUpdateRequest) GetNotificationMessage() string`

GetNotificationMessage returns the NotificationMessage field if non-nil, zero value otherwise.

### GetNotificationMessageOk

`func (o *ContentPermissionUpdateRequest) GetNotificationMessageOk() (*string, bool)`

GetNotificationMessageOk returns a tuple with the NotificationMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationMessage

`func (o *ContentPermissionUpdateRequest) SetNotificationMessage(v string)`

SetNotificationMessage sets NotificationMessage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



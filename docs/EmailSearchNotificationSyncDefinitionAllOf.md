# EmailSearchNotificationSyncDefinitionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ToList** | **[]string** | A list of email recipients. | 
**SubjectTemplate** | Pointer to **string** | If the notification is scheduled with a threshold, the default subject template will be \&quot;Search Alert: {{AlertCondition}} results found for {{SearchName}}\&quot;. For email notifications without a threshold, the default subject template is \&quot;Search Results: {{SearchName}}\&quot;. | [optional] 
**IncludeQuery** | Pointer to **bool** | A boolean value to indicate if the search query should be included in the notification email. | [optional] [default to true]
**IncludeResultSet** | Pointer to **bool** | A boolean value to indicate if the search result set should be included in the notification email. | [optional] [default to true]
**IncludeHistogram** | Pointer to **bool** | A boolean value to indicate if the search result histogram should be included in the notification email. | [optional] [default to true]
**IncludeCsvAttachment** | Pointer to **bool** | A boolean value to indicate if the search results should be included in the notification email as a CSV attachment. | [optional] [default to false]

## Methods

### NewEmailSearchNotificationSyncDefinitionAllOf

`func NewEmailSearchNotificationSyncDefinitionAllOf(toList []string, ) *EmailSearchNotificationSyncDefinitionAllOf`

NewEmailSearchNotificationSyncDefinitionAllOf instantiates a new EmailSearchNotificationSyncDefinitionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailSearchNotificationSyncDefinitionAllOfWithDefaults

`func NewEmailSearchNotificationSyncDefinitionAllOfWithDefaults() *EmailSearchNotificationSyncDefinitionAllOf`

NewEmailSearchNotificationSyncDefinitionAllOfWithDefaults instantiates a new EmailSearchNotificationSyncDefinitionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToList

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetToList() []string`

GetToList returns the ToList field if non-nil, zero value otherwise.

### GetToListOk

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetToListOk() (*[]string, bool)`

GetToListOk returns a tuple with the ToList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToList

`func (o *EmailSearchNotificationSyncDefinitionAllOf) SetToList(v []string)`

SetToList sets ToList field to given value.


### GetSubjectTemplate

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetSubjectTemplate() string`

GetSubjectTemplate returns the SubjectTemplate field if non-nil, zero value otherwise.

### GetSubjectTemplateOk

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetSubjectTemplateOk() (*string, bool)`

GetSubjectTemplateOk returns a tuple with the SubjectTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTemplate

`func (o *EmailSearchNotificationSyncDefinitionAllOf) SetSubjectTemplate(v string)`

SetSubjectTemplate sets SubjectTemplate field to given value.

### HasSubjectTemplate

`func (o *EmailSearchNotificationSyncDefinitionAllOf) HasSubjectTemplate() bool`

HasSubjectTemplate returns a boolean if a field has been set.

### GetIncludeQuery

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeQuery() bool`

GetIncludeQuery returns the IncludeQuery field if non-nil, zero value otherwise.

### GetIncludeQueryOk

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeQueryOk() (*bool, bool)`

GetIncludeQueryOk returns a tuple with the IncludeQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeQuery

`func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeQuery(v bool)`

SetIncludeQuery sets IncludeQuery field to given value.

### HasIncludeQuery

`func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeQuery() bool`

HasIncludeQuery returns a boolean if a field has been set.

### GetIncludeResultSet

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeResultSet() bool`

GetIncludeResultSet returns the IncludeResultSet field if non-nil, zero value otherwise.

### GetIncludeResultSetOk

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeResultSetOk() (*bool, bool)`

GetIncludeResultSetOk returns a tuple with the IncludeResultSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeResultSet

`func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeResultSet(v bool)`

SetIncludeResultSet sets IncludeResultSet field to given value.

### HasIncludeResultSet

`func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeResultSet() bool`

HasIncludeResultSet returns a boolean if a field has been set.

### GetIncludeHistogram

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeHistogram() bool`

GetIncludeHistogram returns the IncludeHistogram field if non-nil, zero value otherwise.

### GetIncludeHistogramOk

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeHistogramOk() (*bool, bool)`

GetIncludeHistogramOk returns a tuple with the IncludeHistogram field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeHistogram

`func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeHistogram(v bool)`

SetIncludeHistogram sets IncludeHistogram field to given value.

### HasIncludeHistogram

`func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeHistogram() bool`

HasIncludeHistogram returns a boolean if a field has been set.

### GetIncludeCsvAttachment

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeCsvAttachment() bool`

GetIncludeCsvAttachment returns the IncludeCsvAttachment field if non-nil, zero value otherwise.

### GetIncludeCsvAttachmentOk

`func (o *EmailSearchNotificationSyncDefinitionAllOf) GetIncludeCsvAttachmentOk() (*bool, bool)`

GetIncludeCsvAttachmentOk returns a tuple with the IncludeCsvAttachment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeCsvAttachment

`func (o *EmailSearchNotificationSyncDefinitionAllOf) SetIncludeCsvAttachment(v bool)`

SetIncludeCsvAttachment sets IncludeCsvAttachment field to given value.

### HasIncludeCsvAttachment

`func (o *EmailSearchNotificationSyncDefinitionAllOf) HasIncludeCsvAttachment() bool`

HasIncludeCsvAttachment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



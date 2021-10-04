# CollectionS3GetObjectAccessDeniedTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**BucketName** | Pointer to **string** | The bucket name of the associated Source. | [optional] 
**AccessKey** | Pointer to **string** | The access key used to make the request. In the case of IAM roles, this is the temporary key used for authentication. | [optional] 

## Methods

### NewCollectionS3GetObjectAccessDeniedTracker

`func NewCollectionS3GetObjectAccessDeniedTracker() *CollectionS3GetObjectAccessDeniedTracker`

NewCollectionS3GetObjectAccessDeniedTracker instantiates a new CollectionS3GetObjectAccessDeniedTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionS3GetObjectAccessDeniedTrackerWithDefaults

`func NewCollectionS3GetObjectAccessDeniedTrackerWithDefaults() *CollectionS3GetObjectAccessDeniedTracker`

NewCollectionS3GetObjectAccessDeniedTrackerWithDefaults instantiates a new CollectionS3GetObjectAccessDeniedTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CollectionS3GetObjectAccessDeniedTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CollectionS3GetObjectAccessDeniedTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CollectionS3GetObjectAccessDeniedTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CollectionS3GetObjectAccessDeniedTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetBucketName

`func (o *CollectionS3GetObjectAccessDeniedTracker) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CollectionS3GetObjectAccessDeniedTracker) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CollectionS3GetObjectAccessDeniedTracker) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *CollectionS3GetObjectAccessDeniedTracker) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetAccessKey

`func (o *CollectionS3GetObjectAccessDeniedTracker) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CollectionS3GetObjectAccessDeniedTracker) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CollectionS3GetObjectAccessDeniedTracker) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CollectionS3GetObjectAccessDeniedTracker) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



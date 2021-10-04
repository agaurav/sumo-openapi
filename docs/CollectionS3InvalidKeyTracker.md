# CollectionS3InvalidKeyTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**AccessKey** | Pointer to **string** | The access key used to make the request. In the case of IAM roles, this is the temporary key used for authentication. | [optional] 

## Methods

### NewCollectionS3InvalidKeyTracker

`func NewCollectionS3InvalidKeyTracker() *CollectionS3InvalidKeyTracker`

NewCollectionS3InvalidKeyTracker instantiates a new CollectionS3InvalidKeyTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionS3InvalidKeyTrackerWithDefaults

`func NewCollectionS3InvalidKeyTrackerWithDefaults() *CollectionS3InvalidKeyTracker`

NewCollectionS3InvalidKeyTrackerWithDefaults instantiates a new CollectionS3InvalidKeyTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CollectionS3InvalidKeyTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CollectionS3InvalidKeyTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CollectionS3InvalidKeyTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CollectionS3InvalidKeyTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetAccessKey

`func (o *CollectionS3InvalidKeyTracker) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CollectionS3InvalidKeyTracker) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CollectionS3InvalidKeyTracker) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CollectionS3InvalidKeyTracker) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



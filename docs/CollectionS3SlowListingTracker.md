# CollectionS3SlowListingTracker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** | Event type. | [optional] 
**BucketName** | Pointer to **string** | The bucket name of the associated Source. | [optional] 
**FlaggedAfterMinutes** | Pointer to **string** | The number of minutes elapsed in scanning after which this incident was created. | [optional] 

## Methods

### NewCollectionS3SlowListingTracker

`func NewCollectionS3SlowListingTracker() *CollectionS3SlowListingTracker`

NewCollectionS3SlowListingTracker instantiates a new CollectionS3SlowListingTracker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionS3SlowListingTrackerWithDefaults

`func NewCollectionS3SlowListingTrackerWithDefaults() *CollectionS3SlowListingTracker`

NewCollectionS3SlowListingTrackerWithDefaults instantiates a new CollectionS3SlowListingTracker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *CollectionS3SlowListingTracker) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *CollectionS3SlowListingTracker) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *CollectionS3SlowListingTracker) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *CollectionS3SlowListingTracker) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetBucketName

`func (o *CollectionS3SlowListingTracker) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *CollectionS3SlowListingTracker) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *CollectionS3SlowListingTracker) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *CollectionS3SlowListingTracker) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetFlaggedAfterMinutes

`func (o *CollectionS3SlowListingTracker) GetFlaggedAfterMinutes() string`

GetFlaggedAfterMinutes returns the FlaggedAfterMinutes field if non-nil, zero value otherwise.

### GetFlaggedAfterMinutesOk

`func (o *CollectionS3SlowListingTracker) GetFlaggedAfterMinutesOk() (*string, bool)`

GetFlaggedAfterMinutesOk returns a tuple with the FlaggedAfterMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedAfterMinutes

`func (o *CollectionS3SlowListingTracker) SetFlaggedAfterMinutes(v string)`

SetFlaggedAfterMinutes sets FlaggedAfterMinutes field to given value.

### HasFlaggedAfterMinutes

`func (o *CollectionS3SlowListingTracker) HasFlaggedAfterMinutes() bool`

HasFlaggedAfterMinutes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



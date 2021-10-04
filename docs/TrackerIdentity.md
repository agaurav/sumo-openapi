# TrackerIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrackerId** | **string** | Name that uniquely identifies the health event. It focuses on what happened rather than why. | 
**Error** | **string** | Description of the underlying reason for the event change. | 
**Description** | **string** | A more elaborate description of why the event occurred. | 

## Methods

### NewTrackerIdentity

`func NewTrackerIdentity(trackerId string, error_ string, description string, ) *TrackerIdentity`

NewTrackerIdentity instantiates a new TrackerIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrackerIdentityWithDefaults

`func NewTrackerIdentityWithDefaults() *TrackerIdentity`

NewTrackerIdentityWithDefaults instantiates a new TrackerIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrackerId

`func (o *TrackerIdentity) GetTrackerId() string`

GetTrackerId returns the TrackerId field if non-nil, zero value otherwise.

### GetTrackerIdOk

`func (o *TrackerIdentity) GetTrackerIdOk() (*string, bool)`

GetTrackerIdOk returns a tuple with the TrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackerId

`func (o *TrackerIdentity) SetTrackerId(v string)`

SetTrackerId sets TrackerId field to given value.


### GetError

`func (o *TrackerIdentity) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *TrackerIdentity) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *TrackerIdentity) SetError(v string)`

SetError sets Error field to given value.


### GetDescription

`func (o *TrackerIdentity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TrackerIdentity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TrackerIdentity) SetDescription(v string)`

SetDescription sets Description field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



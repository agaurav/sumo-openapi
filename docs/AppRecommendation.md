# AppRecommendation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | **string** | Unique identifier for the app. | 
**Name** | **string** | Name of the app. | 
**Description** | **string** | Description of the app. | 
**IconURL** | **string** | URL of the app icon. | 
**Confidence** | **float64** | Percentage relevance of recommendation. | 

## Methods

### NewAppRecommendation

`func NewAppRecommendation(uuid string, name string, description string, iconURL string, confidence float64, ) *AppRecommendation`

NewAppRecommendation instantiates a new AppRecommendation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppRecommendationWithDefaults

`func NewAppRecommendationWithDefaults() *AppRecommendation`

NewAppRecommendationWithDefaults instantiates a new AppRecommendation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *AppRecommendation) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppRecommendation) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppRecommendation) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppRecommendation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppRecommendation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppRecommendation) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AppRecommendation) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppRecommendation) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppRecommendation) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetIconURL

`func (o *AppRecommendation) GetIconURL() string`

GetIconURL returns the IconURL field if non-nil, zero value otherwise.

### GetIconURLOk

`func (o *AppRecommendation) GetIconURLOk() (*string, bool)`

GetIconURLOk returns a tuple with the IconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconURL

`func (o *AppRecommendation) SetIconURL(v string)`

SetIconURL sets IconURL field to given value.


### GetConfidence

`func (o *AppRecommendation) GetConfidence() float64`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *AppRecommendation) GetConfidenceOk() (*float64, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *AppRecommendation) SetConfidence(v float64)`

SetConfidence sets Confidence field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# RelatedAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alert** | Pointer to [**AlertsLibraryAlertResponse**](AlertsLibraryAlertResponse.md) |  | [optional] 
**Relations** | Pointer to **[]string** | Tags describing the relationship between the two alerts. | [optional] 

## Methods

### NewRelatedAlert

`func NewRelatedAlert() *RelatedAlert`

NewRelatedAlert instantiates a new RelatedAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelatedAlertWithDefaults

`func NewRelatedAlertWithDefaults() *RelatedAlert`

NewRelatedAlertWithDefaults instantiates a new RelatedAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlert

`func (o *RelatedAlert) GetAlert() AlertsLibraryAlertResponse`

GetAlert returns the Alert field if non-nil, zero value otherwise.

### GetAlertOk

`func (o *RelatedAlert) GetAlertOk() (*AlertsLibraryAlertResponse, bool)`

GetAlertOk returns a tuple with the Alert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlert

`func (o *RelatedAlert) SetAlert(v AlertsLibraryAlertResponse)`

SetAlert sets Alert field to given value.

### HasAlert

`func (o *RelatedAlert) HasAlert() bool`

HasAlert returns a boolean if a field has been set.

### GetRelations

`func (o *RelatedAlert) GetRelations() []string`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *RelatedAlert) GetRelationsOk() (*[]string, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *RelatedAlert) SetRelations(v []string)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *RelatedAlert) HasRelations() bool`

HasRelations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



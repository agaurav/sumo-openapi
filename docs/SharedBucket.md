# SharedBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the bucket. | 
**Label** | **string** | The text to be displayed on UI for this bucket. | 
**LinkedEntitlementTypes** | **[]string** | List of entitlement types which can consume from this bucket. | 
**Capacitites** | Pointer to [**[]Capacity**](Capacity.md) | List of capacities alloted. | [optional] 

## Methods

### NewSharedBucket

`func NewSharedBucket(name string, label string, linkedEntitlementTypes []string, ) *SharedBucket`

NewSharedBucket instantiates a new SharedBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedBucketWithDefaults

`func NewSharedBucketWithDefaults() *SharedBucket`

NewSharedBucketWithDefaults instantiates a new SharedBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SharedBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SharedBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SharedBucket) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *SharedBucket) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SharedBucket) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SharedBucket) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLinkedEntitlementTypes

`func (o *SharedBucket) GetLinkedEntitlementTypes() []string`

GetLinkedEntitlementTypes returns the LinkedEntitlementTypes field if non-nil, zero value otherwise.

### GetLinkedEntitlementTypesOk

`func (o *SharedBucket) GetLinkedEntitlementTypesOk() (*[]string, bool)`

GetLinkedEntitlementTypesOk returns a tuple with the LinkedEntitlementTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedEntitlementTypes

`func (o *SharedBucket) SetLinkedEntitlementTypes(v []string)`

SetLinkedEntitlementTypes sets LinkedEntitlementTypes field to given value.


### GetCapacitites

`func (o *SharedBucket) GetCapacitites() []Capacity`

GetCapacitites returns the Capacitites field if non-nil, zero value otherwise.

### GetCapacititesOk

`func (o *SharedBucket) GetCapacititesOk() (*[]Capacity, bool)`

GetCapacititesOk returns a tuple with the Capacitites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacitites

`func (o *SharedBucket) SetCapacitites(v []Capacity)`

SetCapacitites sets Capacitites field to given value.

### HasCapacitites

`func (o *SharedBucket) HasCapacitites() bool`

HasCapacitites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



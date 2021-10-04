# Path

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PathItems** | [**[]PathItem**](PathItem.md) | Elements of the path. | 
**Path** | **string** | String representation of the path. | 

## Methods

### NewPath

`func NewPath(pathItems []PathItem, path string, ) *Path`

NewPath instantiates a new Path object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathWithDefaults

`func NewPathWithDefaults() *Path`

NewPathWithDefaults instantiates a new Path object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPathItems

`func (o *Path) GetPathItems() []PathItem`

GetPathItems returns the PathItems field if non-nil, zero value otherwise.

### GetPathItemsOk

`func (o *Path) GetPathItemsOk() (*[]PathItem, bool)`

GetPathItemsOk returns a tuple with the PathItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathItems

`func (o *Path) SetPathItems(v []PathItem)`

SetPathItems sets PathItems field to given value.


### GetPath

`func (o *Path) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Path) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Path) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



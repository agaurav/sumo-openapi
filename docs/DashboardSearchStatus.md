# DashboardSearchStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** | Current state of the search. | 
**PercentCompleted** | Pointer to **int32** | Percentage of search completed. | [optional] 

## Methods

### NewDashboardSearchStatus

`func NewDashboardSearchStatus(state string, ) *DashboardSearchStatus`

NewDashboardSearchStatus instantiates a new DashboardSearchStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardSearchStatusWithDefaults

`func NewDashboardSearchStatusWithDefaults() *DashboardSearchStatus`

NewDashboardSearchStatusWithDefaults instantiates a new DashboardSearchStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *DashboardSearchStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DashboardSearchStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DashboardSearchStatus) SetState(v string)`

SetState sets State field to given value.


### GetPercentCompleted

`func (o *DashboardSearchStatus) GetPercentCompleted() int32`

GetPercentCompleted returns the PercentCompleted field if non-nil, zero value otherwise.

### GetPercentCompletedOk

`func (o *DashboardSearchStatus) GetPercentCompletedOk() (*int32, bool)`

GetPercentCompletedOk returns a tuple with the PercentCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentCompleted

`func (o *DashboardSearchStatus) SetPercentCompleted(v int32)`

SetPercentCompleted sets PercentCompleted field to given value.

### HasPercentCompleted

`func (o *DashboardSearchStatus) HasPercentCompleted() bool`

HasPercentCompleted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# LinkedDashboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the linked dashboard. | 
**RelativePath** | Pointer to **string** | Relative path of the linked dashboard to the dashboard of the linking panel. | [optional] 
**IncludeTimeRange** | Pointer to **bool** | Include time range from the current dashboard to the linked dashboard. | [optional] [default to true]
**IncludeVariables** | Pointer to **bool** | Include variables from the current dashboard to the linked dashboard. | [optional] [default to true]

## Methods

### NewLinkedDashboard

`func NewLinkedDashboard(id string, ) *LinkedDashboard`

NewLinkedDashboard instantiates a new LinkedDashboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedDashboardWithDefaults

`func NewLinkedDashboardWithDefaults() *LinkedDashboard`

NewLinkedDashboardWithDefaults instantiates a new LinkedDashboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LinkedDashboard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LinkedDashboard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LinkedDashboard) SetId(v string)`

SetId sets Id field to given value.


### GetRelativePath

`func (o *LinkedDashboard) GetRelativePath() string`

GetRelativePath returns the RelativePath field if non-nil, zero value otherwise.

### GetRelativePathOk

`func (o *LinkedDashboard) GetRelativePathOk() (*string, bool)`

GetRelativePathOk returns a tuple with the RelativePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePath

`func (o *LinkedDashboard) SetRelativePath(v string)`

SetRelativePath sets RelativePath field to given value.

### HasRelativePath

`func (o *LinkedDashboard) HasRelativePath() bool`

HasRelativePath returns a boolean if a field has been set.

### GetIncludeTimeRange

`func (o *LinkedDashboard) GetIncludeTimeRange() bool`

GetIncludeTimeRange returns the IncludeTimeRange field if non-nil, zero value otherwise.

### GetIncludeTimeRangeOk

`func (o *LinkedDashboard) GetIncludeTimeRangeOk() (*bool, bool)`

GetIncludeTimeRangeOk returns a tuple with the IncludeTimeRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTimeRange

`func (o *LinkedDashboard) SetIncludeTimeRange(v bool)`

SetIncludeTimeRange sets IncludeTimeRange field to given value.

### HasIncludeTimeRange

`func (o *LinkedDashboard) HasIncludeTimeRange() bool`

HasIncludeTimeRange returns a boolean if a field has been set.

### GetIncludeVariables

`func (o *LinkedDashboard) GetIncludeVariables() bool`

GetIncludeVariables returns the IncludeVariables field if non-nil, zero value otherwise.

### GetIncludeVariablesOk

`func (o *LinkedDashboard) GetIncludeVariablesOk() (*bool, bool)`

GetIncludeVariablesOk returns a tuple with the IncludeVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeVariables

`func (o *LinkedDashboard) SetIncludeVariables(v bool)`

SetIncludeVariables sets IncludeVariables field to given value.

### HasIncludeVariables

`func (o *LinkedDashboard) HasIncludeVariables() bool`

HasIncludeVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListRoleModelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RoleModel**](RoleModel.md) | List of roles. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewListRoleModelsResponse

`func NewListRoleModelsResponse(data []RoleModel, ) *ListRoleModelsResponse`

NewListRoleModelsResponse instantiates a new ListRoleModelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoleModelsResponseWithDefaults

`func NewListRoleModelsResponseWithDefaults() *ListRoleModelsResponse`

NewListRoleModelsResponseWithDefaults instantiates a new ListRoleModelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListRoleModelsResponse) GetData() []RoleModel`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListRoleModelsResponse) GetDataOk() (*[]RoleModel, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListRoleModelsResponse) SetData(v []RoleModel)`

SetData sets Data field to given value.


### GetNext

`func (o *ListRoleModelsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListRoleModelsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListRoleModelsResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListRoleModelsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



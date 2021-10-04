# ListConnectionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Connection**](Connection.md) | List of connections. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewListConnectionsResponse

`func NewListConnectionsResponse(data []Connection, ) *ListConnectionsResponse`

NewListConnectionsResponse instantiates a new ListConnectionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectionsResponseWithDefaults

`func NewListConnectionsResponseWithDefaults() *ListConnectionsResponse`

NewListConnectionsResponseWithDefaults instantiates a new ListConnectionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListConnectionsResponse) GetData() []Connection`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListConnectionsResponse) GetDataOk() (*[]Connection, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListConnectionsResponse) SetData(v []Connection)`

SetData sets Data field to given value.


### GetNext

`func (o *ListConnectionsResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListConnectionsResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListConnectionsResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListConnectionsResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ListCollectorIdentitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CollectorIdentity**](CollectorIdentity.md) | List of Collector identities. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewListCollectorIdentitiesResponse

`func NewListCollectorIdentitiesResponse(data []CollectorIdentity, ) *ListCollectorIdentitiesResponse`

NewListCollectorIdentitiesResponse instantiates a new ListCollectorIdentitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCollectorIdentitiesResponseWithDefaults

`func NewListCollectorIdentitiesResponseWithDefaults() *ListCollectorIdentitiesResponse`

NewListCollectorIdentitiesResponseWithDefaults instantiates a new ListCollectorIdentitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListCollectorIdentitiesResponse) GetData() []CollectorIdentity`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCollectorIdentitiesResponse) GetDataOk() (*[]CollectorIdentity, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCollectorIdentitiesResponse) SetData(v []CollectorIdentity)`

SetData sets Data field to given value.


### GetNext

`func (o *ListCollectorIdentitiesResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListCollectorIdentitiesResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListCollectorIdentitiesResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListCollectorIdentitiesResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



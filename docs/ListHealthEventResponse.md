# ListHealthEventResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]HealthEvent**](HealthEvent.md) | List of health events. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewListHealthEventResponse

`func NewListHealthEventResponse(data []HealthEvent, ) *ListHealthEventResponse`

NewListHealthEventResponse instantiates a new ListHealthEventResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealthEventResponseWithDefaults

`func NewListHealthEventResponseWithDefaults() *ListHealthEventResponse`

NewListHealthEventResponseWithDefaults instantiates a new ListHealthEventResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListHealthEventResponse) GetData() []HealthEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListHealthEventResponse) GetDataOk() (*[]HealthEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListHealthEventResponse) SetData(v []HealthEvent)`

SetData sets Data field to given value.


### GetNext

`func (o *ListHealthEventResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *ListHealthEventResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *ListHealthEventResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *ListHealthEventResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GetCollectorsUsageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Collector**](Collector.md) | List of collectors. | 
**Next** | Pointer to **string** | Next continuation token. | [optional] 

## Methods

### NewGetCollectorsUsageResponse

`func NewGetCollectorsUsageResponse(data []Collector, ) *GetCollectorsUsageResponse`

NewGetCollectorsUsageResponse instantiates a new GetCollectorsUsageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCollectorsUsageResponseWithDefaults

`func NewGetCollectorsUsageResponseWithDefaults() *GetCollectorsUsageResponse`

NewGetCollectorsUsageResponseWithDefaults instantiates a new GetCollectorsUsageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetCollectorsUsageResponse) GetData() []Collector`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetCollectorsUsageResponse) GetDataOk() (*[]Collector, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetCollectorsUsageResponse) SetData(v []Collector)`

SetData sets Data field to given value.


### GetNext

`func (o *GetCollectorsUsageResponse) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *GetCollectorsUsageResponse) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *GetCollectorsUsageResponse) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *GetCollectorsUsageResponse) HasNext() bool`

HasNext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



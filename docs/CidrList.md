# CidrList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Cidr**](Cidr.md) | An array of CIDR notations and/or IP addresses. | 

## Methods

### NewCidrList

`func NewCidrList(data []Cidr, ) *CidrList`

NewCidrList instantiates a new CidrList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCidrListWithDefaults

`func NewCidrListWithDefaults() *CidrList`

NewCidrListWithDefaults instantiates a new CidrList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CidrList) GetData() []Cidr`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CidrList) GetDataOk() (*[]Cidr, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CidrList) SetData(v []Cidr)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ServiceNowAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The identifier of the connection. | 
**ConnectionSubtype** | Pointer to **string** | The subtype of the connection. Valid values are &#x60;Event&#x60; or &#x60;Incident&#x60;. | [optional] 
**PayloadOverride** | Pointer to **string** | The override of the default JSON payload of the connection. Should be in JSON format. | [optional] 

## Methods

### NewServiceNowAllOf

`func NewServiceNowAllOf(connectionId string, ) *ServiceNowAllOf`

NewServiceNowAllOf instantiates a new ServiceNowAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceNowAllOfWithDefaults

`func NewServiceNowAllOfWithDefaults() *ServiceNowAllOf`

NewServiceNowAllOfWithDefaults instantiates a new ServiceNowAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *ServiceNowAllOf) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ServiceNowAllOf) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ServiceNowAllOf) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetConnectionSubtype

`func (o *ServiceNowAllOf) GetConnectionSubtype() string`

GetConnectionSubtype returns the ConnectionSubtype field if non-nil, zero value otherwise.

### GetConnectionSubtypeOk

`func (o *ServiceNowAllOf) GetConnectionSubtypeOk() (*string, bool)`

GetConnectionSubtypeOk returns a tuple with the ConnectionSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSubtype

`func (o *ServiceNowAllOf) SetConnectionSubtype(v string)`

SetConnectionSubtype sets ConnectionSubtype field to given value.

### HasConnectionSubtype

`func (o *ServiceNowAllOf) HasConnectionSubtype() bool`

HasConnectionSubtype returns a boolean if a field has been set.

### GetPayloadOverride

`func (o *ServiceNowAllOf) GetPayloadOverride() string`

GetPayloadOverride returns the PayloadOverride field if non-nil, zero value otherwise.

### GetPayloadOverrideOk

`func (o *ServiceNowAllOf) GetPayloadOverrideOk() (*string, bool)`

GetPayloadOverrideOk returns a tuple with the PayloadOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadOverride

`func (o *ServiceNowAllOf) SetPayloadOverride(v string)`

SetPayloadOverride sets PayloadOverride field to given value.

### HasPayloadOverride

`func (o *ServiceNowAllOf) HasPayloadOverride() bool`

HasPayloadOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



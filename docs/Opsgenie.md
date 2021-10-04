# Opsgenie

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | **string** | The identifier of the connection. | 
**PayloadOverride** | Pointer to **string** | The override of the default JSON payload of the connection. Should be in JSON format. | [optional] 

## Methods

### NewOpsgenie

`func NewOpsgenie(connectionId string, ) *Opsgenie`

NewOpsgenie instantiates a new Opsgenie object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpsgenieWithDefaults

`func NewOpsgenieWithDefaults() *Opsgenie`

NewOpsgenieWithDefaults instantiates a new Opsgenie object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *Opsgenie) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *Opsgenie) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *Opsgenie) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetPayloadOverride

`func (o *Opsgenie) GetPayloadOverride() string`

GetPayloadOverride returns the PayloadOverride field if non-nil, zero value otherwise.

### GetPayloadOverrideOk

`func (o *Opsgenie) GetPayloadOverrideOk() (*string, bool)`

GetPayloadOverrideOk returns a tuple with the PayloadOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadOverride

`func (o *Opsgenie) SetPayloadOverride(v string)`

SetPayloadOverride sets PayloadOverride field to given value.

### HasPayloadOverride

`func (o *Opsgenie) HasPayloadOverride() bool`

HasPayloadOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



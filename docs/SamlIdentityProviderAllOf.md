# SamlIdentityProviderAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the SAML Identity Provider. | 
**AssertionConsumerUrl** | Pointer to **string** | The URL on Sumo Logic where the IdP will redirect to with its authentication response. | [optional] [default to ""]
**EntityId** | Pointer to **string** | A unique identifier that is the intended audience of the SAML assertion. | [optional] [default to ""]

## Methods

### NewSamlIdentityProviderAllOf

`func NewSamlIdentityProviderAllOf(id string, ) *SamlIdentityProviderAllOf`

NewSamlIdentityProviderAllOf instantiates a new SamlIdentityProviderAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlIdentityProviderAllOfWithDefaults

`func NewSamlIdentityProviderAllOfWithDefaults() *SamlIdentityProviderAllOf`

NewSamlIdentityProviderAllOfWithDefaults instantiates a new SamlIdentityProviderAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SamlIdentityProviderAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlIdentityProviderAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SamlIdentityProviderAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetAssertionConsumerUrl

`func (o *SamlIdentityProviderAllOf) GetAssertionConsumerUrl() string`

GetAssertionConsumerUrl returns the AssertionConsumerUrl field if non-nil, zero value otherwise.

### GetAssertionConsumerUrlOk

`func (o *SamlIdentityProviderAllOf) GetAssertionConsumerUrlOk() (*string, bool)`

GetAssertionConsumerUrlOk returns a tuple with the AssertionConsumerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionConsumerUrl

`func (o *SamlIdentityProviderAllOf) SetAssertionConsumerUrl(v string)`

SetAssertionConsumerUrl sets AssertionConsumerUrl field to given value.

### HasAssertionConsumerUrl

`func (o *SamlIdentityProviderAllOf) HasAssertionConsumerUrl() bool`

HasAssertionConsumerUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *SamlIdentityProviderAllOf) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SamlIdentityProviderAllOf) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SamlIdentityProviderAllOf) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SamlIdentityProviderAllOf) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



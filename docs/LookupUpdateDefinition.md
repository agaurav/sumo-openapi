# LookupUpdateDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | **int32** | A time to live for each entry in the lookup table (in minutes). 0 is a special value. A TTL of 0 implies entry will never be deleted from the table. | [default to 0]
**Description** | **string** | The description of the lookup table. The description cannot be blank. | 
**SizeLimitAction** | Pointer to **string** | The action that needs to be taken when the size limit is reached for the table. The possible values can be &#x60;StopIncomingMessages&#x60; or &#x60;DeleteOldData&#x60;. DeleteOldData will starting deleting old data once size limit is reached whereas StopIncomingMessages will discard all the updates made to the lookup table once size limit is reached. | [optional] [default to "StopIncomingMessages"]

## Methods

### NewLookupUpdateDefinition

`func NewLookupUpdateDefinition(ttl int32, description string, ) *LookupUpdateDefinition`

NewLookupUpdateDefinition instantiates a new LookupUpdateDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupUpdateDefinitionWithDefaults

`func NewLookupUpdateDefinitionWithDefaults() *LookupUpdateDefinition`

NewLookupUpdateDefinitionWithDefaults instantiates a new LookupUpdateDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *LookupUpdateDefinition) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *LookupUpdateDefinition) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *LookupUpdateDefinition) SetTtl(v int32)`

SetTtl sets Ttl field to given value.


### GetDescription

`func (o *LookupUpdateDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LookupUpdateDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LookupUpdateDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetSizeLimitAction

`func (o *LookupUpdateDefinition) GetSizeLimitAction() string`

GetSizeLimitAction returns the SizeLimitAction field if non-nil, zero value otherwise.

### GetSizeLimitActionOk

`func (o *LookupUpdateDefinition) GetSizeLimitActionOk() (*string, bool)`

GetSizeLimitActionOk returns a tuple with the SizeLimitAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeLimitAction

`func (o *LookupUpdateDefinition) SetSizeLimitAction(v string)`

SetSizeLimitAction sets SizeLimitAction field to given value.

### HasSizeLimitAction

`func (o *LookupUpdateDefinition) HasSizeLimitAction() bool`

HasSizeLimitAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# ResourceIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of the resource. | 
**Name** | Pointer to **string** | The name of the resource. | [optional] [default to "Unknown"]
**Type** | **string** | -&gt; Resource type. Supported types are - &#x60;Collector&#x60;, &#x60;Source&#x60;, &#x60;IngestBudget&#x60; and &#x60;Organisation&#x60;. | 

## Methods

### NewResourceIdentity

`func NewResourceIdentity(id string, type_ string, ) *ResourceIdentity`

NewResourceIdentity instantiates a new ResourceIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceIdentityWithDefaults

`func NewResourceIdentityWithDefaults() *ResourceIdentity`

NewResourceIdentityWithDefaults instantiates a new ResourceIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceIdentity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ResourceIdentity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceIdentity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceIdentity) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceIdentity) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *ResourceIdentity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceIdentity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceIdentity) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



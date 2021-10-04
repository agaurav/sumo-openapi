# LookupTableAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifier of the lookup table as a content item. | [optional] 
**ContentPath** | Pointer to **string** | Address/path of the parent folder of this lookup table in content library. For example, a lookup table existing  in the personal/lookupTable folder for user johndoe would be: /Library/Users/johndoe@acme.com/lookupTable | [optional] 
**Size** | Pointer to **int64** | The current size of the lookup table in bytes | [optional] 

## Methods

### NewLookupTableAllOf

`func NewLookupTableAllOf() *LookupTableAllOf`

NewLookupTableAllOf instantiates a new LookupTableAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLookupTableAllOfWithDefaults

`func NewLookupTableAllOfWithDefaults() *LookupTableAllOf`

NewLookupTableAllOfWithDefaults instantiates a new LookupTableAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LookupTableAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LookupTableAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LookupTableAllOf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LookupTableAllOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContentPath

`func (o *LookupTableAllOf) GetContentPath() string`

GetContentPath returns the ContentPath field if non-nil, zero value otherwise.

### GetContentPathOk

`func (o *LookupTableAllOf) GetContentPathOk() (*string, bool)`

GetContentPathOk returns a tuple with the ContentPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPath

`func (o *LookupTableAllOf) SetContentPath(v string)`

SetContentPath sets ContentPath field to given value.

### HasContentPath

`func (o *LookupTableAllOf) HasContentPath() bool`

HasContentPath returns a boolean if a field has been set.

### GetSize

`func (o *LookupTableAllOf) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *LookupTableAllOf) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *LookupTableAllOf) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *LookupTableAllOf) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



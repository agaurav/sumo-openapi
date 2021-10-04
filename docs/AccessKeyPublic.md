# AccessKeyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the access key. | 
**Label** | **string** | The name of the access key. | 
**CorsHeaders** | Pointer to **[]string** | An array of domains for which the access key is valid. Whether Sumo Logic accepts or rejects an API request depends on whether it contains an ORIGIN header and the entries in the allowlist. Sumo Logic will reject:   1. Requests with an ORIGIN header but the allowlist is empty.   2. Requests with an ORIGIN header that don&#39;t match any entry in the allowlist. | [optional] 
**Disabled** | **bool** | Indicates whether the access key is disabled or not. | 
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the access key. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 

## Methods

### NewAccessKeyPublic

`func NewAccessKeyPublic(id string, label string, disabled bool, createdAt time.Time, createdBy string, modifiedAt time.Time, ) *AccessKeyPublic`

NewAccessKeyPublic instantiates a new AccessKeyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessKeyPublicWithDefaults

`func NewAccessKeyPublicWithDefaults() *AccessKeyPublic`

NewAccessKeyPublicWithDefaults instantiates a new AccessKeyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AccessKeyPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessKeyPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessKeyPublic) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *AccessKeyPublic) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AccessKeyPublic) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AccessKeyPublic) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCorsHeaders

`func (o *AccessKeyPublic) GetCorsHeaders() []string`

GetCorsHeaders returns the CorsHeaders field if non-nil, zero value otherwise.

### GetCorsHeadersOk

`func (o *AccessKeyPublic) GetCorsHeadersOk() (*[]string, bool)`

GetCorsHeadersOk returns a tuple with the CorsHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorsHeaders

`func (o *AccessKeyPublic) SetCorsHeaders(v []string)`

SetCorsHeaders sets CorsHeaders field to given value.

### HasCorsHeaders

`func (o *AccessKeyPublic) HasCorsHeaders() bool`

HasCorsHeaders returns a boolean if a field has been set.

### GetDisabled

`func (o *AccessKeyPublic) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *AccessKeyPublic) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *AccessKeyPublic) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.


### GetCreatedAt

`func (o *AccessKeyPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AccessKeyPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AccessKeyPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AccessKeyPublic) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AccessKeyPublic) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AccessKeyPublic) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *AccessKeyPublic) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AccessKeyPublic) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AccessKeyPublic) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



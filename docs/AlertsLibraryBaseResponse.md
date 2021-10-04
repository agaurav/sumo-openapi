# AlertsLibraryBaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the alert or folder. | 
**Name** | **string** | Identifier of the alert or folder. | 
**Description** | **string** | Description of the alert or folder. | 
**Version** | **int64** | Version of the alert or folder. | 
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**ParentId** | **string** | Identifier of the parent folder. | 
**ContentType** | **string** | Type of the content. Valid values:   1) Alert   2) Folder | 
**Type** | **string** | Type of the object model. | 
**IsLocked** | Pointer to **bool** | Whether the object is locked. | [optional] 
**IsSystem** | **bool** | System objects are objects provided by Sumo Logic. System objects can only be localized. Non-local fields can&#39;t be updated. | 
**IsMutable** | **bool** | Immutable objects are \&quot;READ-ONLY\&quot;. | 

## Methods

### NewAlertsLibraryBaseResponse

`func NewAlertsLibraryBaseResponse(id string, name string, description string, version int64, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, parentId string, contentType string, type_ string, isSystem bool, isMutable bool, ) *AlertsLibraryBaseResponse`

NewAlertsLibraryBaseResponse instantiates a new AlertsLibraryBaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsLibraryBaseResponseWithDefaults

`func NewAlertsLibraryBaseResponseWithDefaults() *AlertsLibraryBaseResponse`

NewAlertsLibraryBaseResponseWithDefaults instantiates a new AlertsLibraryBaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsLibraryBaseResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsLibraryBaseResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsLibraryBaseResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AlertsLibraryBaseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AlertsLibraryBaseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AlertsLibraryBaseResponse) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AlertsLibraryBaseResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AlertsLibraryBaseResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AlertsLibraryBaseResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetVersion

`func (o *AlertsLibraryBaseResponse) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AlertsLibraryBaseResponse) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AlertsLibraryBaseResponse) SetVersion(v int64)`

SetVersion sets Version field to given value.


### GetCreatedAt

`func (o *AlertsLibraryBaseResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AlertsLibraryBaseResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AlertsLibraryBaseResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *AlertsLibraryBaseResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AlertsLibraryBaseResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AlertsLibraryBaseResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *AlertsLibraryBaseResponse) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *AlertsLibraryBaseResponse) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *AlertsLibraryBaseResponse) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *AlertsLibraryBaseResponse) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *AlertsLibraryBaseResponse) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *AlertsLibraryBaseResponse) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetParentId

`func (o *AlertsLibraryBaseResponse) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *AlertsLibraryBaseResponse) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *AlertsLibraryBaseResponse) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetContentType

`func (o *AlertsLibraryBaseResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AlertsLibraryBaseResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AlertsLibraryBaseResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetType

`func (o *AlertsLibraryBaseResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AlertsLibraryBaseResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AlertsLibraryBaseResponse) SetType(v string)`

SetType sets Type field to given value.


### GetIsLocked

`func (o *AlertsLibraryBaseResponse) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AlertsLibraryBaseResponse) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AlertsLibraryBaseResponse) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AlertsLibraryBaseResponse) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsSystem

`func (o *AlertsLibraryBaseResponse) GetIsSystem() bool`

GetIsSystem returns the IsSystem field if non-nil, zero value otherwise.

### GetIsSystemOk

`func (o *AlertsLibraryBaseResponse) GetIsSystemOk() (*bool, bool)`

GetIsSystemOk returns a tuple with the IsSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsSystem

`func (o *AlertsLibraryBaseResponse) SetIsSystem(v bool)`

SetIsSystem sets IsSystem field to given value.


### GetIsMutable

`func (o *AlertsLibraryBaseResponse) GetIsMutable() bool`

GetIsMutable returns the IsMutable field if non-nil, zero value otherwise.

### GetIsMutableOk

`func (o *AlertsLibraryBaseResponse) GetIsMutableOk() (*bool, bool)`

GetIsMutableOk returns a tuple with the IsMutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMutable

`func (o *AlertsLibraryBaseResponse) SetIsMutable(v bool)`

SetIsMutable sets IsMutable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



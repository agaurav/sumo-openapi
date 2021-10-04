# AppDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentId** | **string** | Content identifier of the app in hexadecimal format. | 
**Uuid** | **string** | Unique identifier for the app. | 
**Name** | **string** | Name of the app. | 
**AppVersion** | **string** | Version of the app. | 
**Preview** | Pointer to **bool** | Indicates whether the app is in preview or not. | [optional] 
**ManifestVersion** | Pointer to **string** | Manifest version of the app | [optional] 

## Methods

### NewAppDefinition

`func NewAppDefinition(contentId string, uuid string, name string, appVersion string, ) *AppDefinition`

NewAppDefinition instantiates a new AppDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDefinitionWithDefaults

`func NewAppDefinitionWithDefaults() *AppDefinition`

NewAppDefinitionWithDefaults instantiates a new AppDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentId

`func (o *AppDefinition) GetContentId() string`

GetContentId returns the ContentId field if non-nil, zero value otherwise.

### GetContentIdOk

`func (o *AppDefinition) GetContentIdOk() (*string, bool)`

GetContentIdOk returns a tuple with the ContentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentId

`func (o *AppDefinition) SetContentId(v string)`

SetContentId sets ContentId field to given value.


### GetUuid

`func (o *AppDefinition) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AppDefinition) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AppDefinition) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetName

`func (o *AppDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetAppVersion

`func (o *AppDefinition) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *AppDefinition) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *AppDefinition) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.


### GetPreview

`func (o *AppDefinition) GetPreview() bool`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *AppDefinition) GetPreviewOk() (*bool, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *AppDefinition) SetPreview(v bool)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *AppDefinition) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetManifestVersion

`func (o *AppDefinition) GetManifestVersion() string`

GetManifestVersion returns the ManifestVersion field if non-nil, zero value otherwise.

### GetManifestVersionOk

`func (o *AppDefinition) GetManifestVersionOk() (*string, bool)`

GetManifestVersionOk returns a tuple with the ManifestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestVersion

`func (o *AppDefinition) SetManifestVersion(v string)`

SetManifestVersion sets ManifestVersion field to given value.

### HasManifestVersion

`func (o *AppDefinition) HasManifestVersion() bool`

HasManifestVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



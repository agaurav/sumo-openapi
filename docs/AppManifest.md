# AppManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Family** | Pointer to **string** | The app family | [optional] 
**Description** | **string** | Description of the app. | 
**Categories** | Pointer to **[]string** | Categories that the app belongs to. | [optional] 
**HoverText** | **string** | Text to be displayed when hovered over in UI. | 
**IconURL** | **string** | App icon URL. | 
**ScreenshotURLs** | Pointer to **[]string** | App screenshot URLs. | [optional] 
**HelpURL** | Pointer to **string** | App help page URL. | [optional] 
**HelpDocIdMap** | Pointer to **map[string]string** | the IDs of the docs pages for this app | [optional] 
**CommunityURL** | Pointer to **string** | App community page URL. | [optional] 
**Requirements** | Pointer to **[]string** | Requirements for the app. | [optional] 
**AccountTypes** | Pointer to **[]string** | Account types that are allowed to install the app | [optional] 
**RequiresInstallationInstructions** | Pointer to **bool** | Indicates whether installation instructions are required or not. | [optional] 
**InstallationInstructions** | Pointer to **string** | Installation instructions for the app. | [optional] 
**Parameters** | Pointer to [**[]ServiceManifestDataSourceParameter**](ServiceManifestDataSourceParameter.md) | Content identifier of the app. | [optional] 
**Author** | Pointer to **string** | App author. | [optional] 
**AuthorWebsite** | Pointer to **string** | App author website URL. | [optional] 

## Methods

### NewAppManifest

`func NewAppManifest(description string, hoverText string, iconURL string, ) *AppManifest`

NewAppManifest instantiates a new AppManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppManifestWithDefaults

`func NewAppManifestWithDefaults() *AppManifest`

NewAppManifestWithDefaults instantiates a new AppManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFamily

`func (o *AppManifest) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AppManifest) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AppManifest) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *AppManifest) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetDescription

`func (o *AppManifest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AppManifest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AppManifest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCategories

`func (o *AppManifest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AppManifest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AppManifest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AppManifest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetHoverText

`func (o *AppManifest) GetHoverText() string`

GetHoverText returns the HoverText field if non-nil, zero value otherwise.

### GetHoverTextOk

`func (o *AppManifest) GetHoverTextOk() (*string, bool)`

GetHoverTextOk returns a tuple with the HoverText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoverText

`func (o *AppManifest) SetHoverText(v string)`

SetHoverText sets HoverText field to given value.


### GetIconURL

`func (o *AppManifest) GetIconURL() string`

GetIconURL returns the IconURL field if non-nil, zero value otherwise.

### GetIconURLOk

`func (o *AppManifest) GetIconURLOk() (*string, bool)`

GetIconURLOk returns a tuple with the IconURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconURL

`func (o *AppManifest) SetIconURL(v string)`

SetIconURL sets IconURL field to given value.


### GetScreenshotURLs

`func (o *AppManifest) GetScreenshotURLs() []string`

GetScreenshotURLs returns the ScreenshotURLs field if non-nil, zero value otherwise.

### GetScreenshotURLsOk

`func (o *AppManifest) GetScreenshotURLsOk() (*[]string, bool)`

GetScreenshotURLsOk returns a tuple with the ScreenshotURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreenshotURLs

`func (o *AppManifest) SetScreenshotURLs(v []string)`

SetScreenshotURLs sets ScreenshotURLs field to given value.

### HasScreenshotURLs

`func (o *AppManifest) HasScreenshotURLs() bool`

HasScreenshotURLs returns a boolean if a field has been set.

### GetHelpURL

`func (o *AppManifest) GetHelpURL() string`

GetHelpURL returns the HelpURL field if non-nil, zero value otherwise.

### GetHelpURLOk

`func (o *AppManifest) GetHelpURLOk() (*string, bool)`

GetHelpURLOk returns a tuple with the HelpURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpURL

`func (o *AppManifest) SetHelpURL(v string)`

SetHelpURL sets HelpURL field to given value.

### HasHelpURL

`func (o *AppManifest) HasHelpURL() bool`

HasHelpURL returns a boolean if a field has been set.

### GetHelpDocIdMap

`func (o *AppManifest) GetHelpDocIdMap() map[string]string`

GetHelpDocIdMap returns the HelpDocIdMap field if non-nil, zero value otherwise.

### GetHelpDocIdMapOk

`func (o *AppManifest) GetHelpDocIdMapOk() (*map[string]string, bool)`

GetHelpDocIdMapOk returns a tuple with the HelpDocIdMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpDocIdMap

`func (o *AppManifest) SetHelpDocIdMap(v map[string]string)`

SetHelpDocIdMap sets HelpDocIdMap field to given value.

### HasHelpDocIdMap

`func (o *AppManifest) HasHelpDocIdMap() bool`

HasHelpDocIdMap returns a boolean if a field has been set.

### GetCommunityURL

`func (o *AppManifest) GetCommunityURL() string`

GetCommunityURL returns the CommunityURL field if non-nil, zero value otherwise.

### GetCommunityURLOk

`func (o *AppManifest) GetCommunityURLOk() (*string, bool)`

GetCommunityURLOk returns a tuple with the CommunityURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunityURL

`func (o *AppManifest) SetCommunityURL(v string)`

SetCommunityURL sets CommunityURL field to given value.

### HasCommunityURL

`func (o *AppManifest) HasCommunityURL() bool`

HasCommunityURL returns a boolean if a field has been set.

### GetRequirements

`func (o *AppManifest) GetRequirements() []string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *AppManifest) GetRequirementsOk() (*[]string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *AppManifest) SetRequirements(v []string)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *AppManifest) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetAccountTypes

`func (o *AppManifest) GetAccountTypes() []string`

GetAccountTypes returns the AccountTypes field if non-nil, zero value otherwise.

### GetAccountTypesOk

`func (o *AppManifest) GetAccountTypesOk() (*[]string, bool)`

GetAccountTypesOk returns a tuple with the AccountTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountTypes

`func (o *AppManifest) SetAccountTypes(v []string)`

SetAccountTypes sets AccountTypes field to given value.

### HasAccountTypes

`func (o *AppManifest) HasAccountTypes() bool`

HasAccountTypes returns a boolean if a field has been set.

### GetRequiresInstallationInstructions

`func (o *AppManifest) GetRequiresInstallationInstructions() bool`

GetRequiresInstallationInstructions returns the RequiresInstallationInstructions field if non-nil, zero value otherwise.

### GetRequiresInstallationInstructionsOk

`func (o *AppManifest) GetRequiresInstallationInstructionsOk() (*bool, bool)`

GetRequiresInstallationInstructionsOk returns a tuple with the RequiresInstallationInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresInstallationInstructions

`func (o *AppManifest) SetRequiresInstallationInstructions(v bool)`

SetRequiresInstallationInstructions sets RequiresInstallationInstructions field to given value.

### HasRequiresInstallationInstructions

`func (o *AppManifest) HasRequiresInstallationInstructions() bool`

HasRequiresInstallationInstructions returns a boolean if a field has been set.

### GetInstallationInstructions

`func (o *AppManifest) GetInstallationInstructions() string`

GetInstallationInstructions returns the InstallationInstructions field if non-nil, zero value otherwise.

### GetInstallationInstructionsOk

`func (o *AppManifest) GetInstallationInstructionsOk() (*string, bool)`

GetInstallationInstructionsOk returns a tuple with the InstallationInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallationInstructions

`func (o *AppManifest) SetInstallationInstructions(v string)`

SetInstallationInstructions sets InstallationInstructions field to given value.

### HasInstallationInstructions

`func (o *AppManifest) HasInstallationInstructions() bool`

HasInstallationInstructions returns a boolean if a field has been set.

### GetParameters

`func (o *AppManifest) GetParameters() []ServiceManifestDataSourceParameter`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *AppManifest) GetParametersOk() (*[]ServiceManifestDataSourceParameter, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *AppManifest) SetParameters(v []ServiceManifestDataSourceParameter)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *AppManifest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetAuthor

`func (o *AppManifest) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *AppManifest) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *AppManifest) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *AppManifest) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetAuthorWebsite

`func (o *AppManifest) GetAuthorWebsite() string`

GetAuthorWebsite returns the AuthorWebsite field if non-nil, zero value otherwise.

### GetAuthorWebsiteOk

`func (o *AppManifest) GetAuthorWebsiteOk() (*string, bool)`

GetAuthorWebsiteOk returns a tuple with the AuthorWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorWebsite

`func (o *AppManifest) SetAuthorWebsite(v string)`

SetAuthorWebsite sets AuthorWebsite field to given value.

### HasAuthorWebsite

`func (o *AppManifest) HasAuthorWebsite() bool`

HasAuthorWebsite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



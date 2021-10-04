/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs see our [API home page](https://help.sumologic.com/APIs). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/Manage/Security/Access-Keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-301-Error-Moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-403-Error-This-operation-is-not-allowed-for-your-account-type) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// AppManifest struct for AppManifest
type AppManifest struct {
	// The app family
	Family *string `json:"family,omitempty"`
	// Description of the app.
	Description string `json:"description"`
	// Categories that the app belongs to.
	Categories *[]string `json:"categories,omitempty"`
	// Text to be displayed when hovered over in UI.
	HoverText string `json:"hoverText"`
	// App icon URL.
	IconURL string `json:"iconURL"`
	// App screenshot URLs.
	ScreenshotURLs *[]string `json:"screenshotURLs,omitempty"`
	// App help page URL.
	HelpURL *string `json:"helpURL,omitempty"`
	// the IDs of the docs pages for this app
	HelpDocIdMap *map[string]string `json:"helpDocIdMap,omitempty"`
	// App community page URL.
	CommunityURL *string `json:"communityURL,omitempty"`
	// Requirements for the app.
	Requirements *[]string `json:"requirements,omitempty"`
	// Account types that are allowed to install the app
	AccountTypes *[]string `json:"accountTypes,omitempty"`
	// Indicates whether installation instructions are required or not.
	RequiresInstallationInstructions *bool `json:"requiresInstallationInstructions,omitempty"`
	// Installation instructions for the app.
	InstallationInstructions *string `json:"installationInstructions,omitempty"`
	// Content identifier of the app.
	Parameters *[]ServiceManifestDataSourceParameter `json:"parameters,omitempty"`
	// App author.
	Author *string `json:"author,omitempty"`
	// App author website URL.
	AuthorWebsite *string `json:"authorWebsite,omitempty"`
}

// NewAppManifest instantiates a new AppManifest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppManifest(description string, hoverText string, iconURL string) *AppManifest {
	this := AppManifest{}
	this.Description = description
	this.HoverText = hoverText
	this.IconURL = iconURL
	return &this
}

// NewAppManifestWithDefaults instantiates a new AppManifest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppManifestWithDefaults() *AppManifest {
	this := AppManifest{}
	return &this
}

// GetFamily returns the Family field value if set, zero value otherwise.
func (o *AppManifest) GetFamily() string {
	if o == nil || o.Family == nil {
		var ret string
		return ret
	}
	return *o.Family
}

// GetFamilyOk returns a tuple with the Family field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetFamilyOk() (*string, bool) {
	if o == nil || o.Family == nil {
		return nil, false
	}
	return o.Family, true
}

// HasFamily returns a boolean if a field has been set.
func (o *AppManifest) HasFamily() bool {
	if o != nil && o.Family != nil {
		return true
	}

	return false
}

// SetFamily gets a reference to the given string and assigns it to the Family field.
func (o *AppManifest) SetFamily(v string) {
	o.Family = &v
}

// GetDescription returns the Description field value
func (o *AppManifest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *AppManifest) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *AppManifest) SetDescription(v string) {
	o.Description = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *AppManifest) GetCategories() []string {
	if o == nil || o.Categories == nil {
		var ret []string
		return ret
	}
	return *o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetCategoriesOk() (*[]string, bool) {
	if o == nil || o.Categories == nil {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *AppManifest) HasCategories() bool {
	if o != nil && o.Categories != nil {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []string and assigns it to the Categories field.
func (o *AppManifest) SetCategories(v []string) {
	o.Categories = &v
}

// GetHoverText returns the HoverText field value
func (o *AppManifest) GetHoverText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HoverText
}

// GetHoverTextOk returns a tuple with the HoverText field value
// and a boolean to check if the value has been set.
func (o *AppManifest) GetHoverTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HoverText, true
}

// SetHoverText sets field value
func (o *AppManifest) SetHoverText(v string) {
	o.HoverText = v
}

// GetIconURL returns the IconURL field value
func (o *AppManifest) GetIconURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconURL
}

// GetIconURLOk returns a tuple with the IconURL field value
// and a boolean to check if the value has been set.
func (o *AppManifest) GetIconURLOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IconURL, true
}

// SetIconURL sets field value
func (o *AppManifest) SetIconURL(v string) {
	o.IconURL = v
}

// GetScreenshotURLs returns the ScreenshotURLs field value if set, zero value otherwise.
func (o *AppManifest) GetScreenshotURLs() []string {
	if o == nil || o.ScreenshotURLs == nil {
		var ret []string
		return ret
	}
	return *o.ScreenshotURLs
}

// GetScreenshotURLsOk returns a tuple with the ScreenshotURLs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetScreenshotURLsOk() (*[]string, bool) {
	if o == nil || o.ScreenshotURLs == nil {
		return nil, false
	}
	return o.ScreenshotURLs, true
}

// HasScreenshotURLs returns a boolean if a field has been set.
func (o *AppManifest) HasScreenshotURLs() bool {
	if o != nil && o.ScreenshotURLs != nil {
		return true
	}

	return false
}

// SetScreenshotURLs gets a reference to the given []string and assigns it to the ScreenshotURLs field.
func (o *AppManifest) SetScreenshotURLs(v []string) {
	o.ScreenshotURLs = &v
}

// GetHelpURL returns the HelpURL field value if set, zero value otherwise.
func (o *AppManifest) GetHelpURL() string {
	if o == nil || o.HelpURL == nil {
		var ret string
		return ret
	}
	return *o.HelpURL
}

// GetHelpURLOk returns a tuple with the HelpURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetHelpURLOk() (*string, bool) {
	if o == nil || o.HelpURL == nil {
		return nil, false
	}
	return o.HelpURL, true
}

// HasHelpURL returns a boolean if a field has been set.
func (o *AppManifest) HasHelpURL() bool {
	if o != nil && o.HelpURL != nil {
		return true
	}

	return false
}

// SetHelpURL gets a reference to the given string and assigns it to the HelpURL field.
func (o *AppManifest) SetHelpURL(v string) {
	o.HelpURL = &v
}

// GetHelpDocIdMap returns the HelpDocIdMap field value if set, zero value otherwise.
func (o *AppManifest) GetHelpDocIdMap() map[string]string {
	if o == nil || o.HelpDocIdMap == nil {
		var ret map[string]string
		return ret
	}
	return *o.HelpDocIdMap
}

// GetHelpDocIdMapOk returns a tuple with the HelpDocIdMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetHelpDocIdMapOk() (*map[string]string, bool) {
	if o == nil || o.HelpDocIdMap == nil {
		return nil, false
	}
	return o.HelpDocIdMap, true
}

// HasHelpDocIdMap returns a boolean if a field has been set.
func (o *AppManifest) HasHelpDocIdMap() bool {
	if o != nil && o.HelpDocIdMap != nil {
		return true
	}

	return false
}

// SetHelpDocIdMap gets a reference to the given map[string]string and assigns it to the HelpDocIdMap field.
func (o *AppManifest) SetHelpDocIdMap(v map[string]string) {
	o.HelpDocIdMap = &v
}

// GetCommunityURL returns the CommunityURL field value if set, zero value otherwise.
func (o *AppManifest) GetCommunityURL() string {
	if o == nil || o.CommunityURL == nil {
		var ret string
		return ret
	}
	return *o.CommunityURL
}

// GetCommunityURLOk returns a tuple with the CommunityURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetCommunityURLOk() (*string, bool) {
	if o == nil || o.CommunityURL == nil {
		return nil, false
	}
	return o.CommunityURL, true
}

// HasCommunityURL returns a boolean if a field has been set.
func (o *AppManifest) HasCommunityURL() bool {
	if o != nil && o.CommunityURL != nil {
		return true
	}

	return false
}

// SetCommunityURL gets a reference to the given string and assigns it to the CommunityURL field.
func (o *AppManifest) SetCommunityURL(v string) {
	o.CommunityURL = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *AppManifest) GetRequirements() []string {
	if o == nil || o.Requirements == nil {
		var ret []string
		return ret
	}
	return *o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetRequirementsOk() (*[]string, bool) {
	if o == nil || o.Requirements == nil {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *AppManifest) HasRequirements() bool {
	if o != nil && o.Requirements != nil {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given []string and assigns it to the Requirements field.
func (o *AppManifest) SetRequirements(v []string) {
	o.Requirements = &v
}

// GetAccountTypes returns the AccountTypes field value if set, zero value otherwise.
func (o *AppManifest) GetAccountTypes() []string {
	if o == nil || o.AccountTypes == nil {
		var ret []string
		return ret
	}
	return *o.AccountTypes
}

// GetAccountTypesOk returns a tuple with the AccountTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetAccountTypesOk() (*[]string, bool) {
	if o == nil || o.AccountTypes == nil {
		return nil, false
	}
	return o.AccountTypes, true
}

// HasAccountTypes returns a boolean if a field has been set.
func (o *AppManifest) HasAccountTypes() bool {
	if o != nil && o.AccountTypes != nil {
		return true
	}

	return false
}

// SetAccountTypes gets a reference to the given []string and assigns it to the AccountTypes field.
func (o *AppManifest) SetAccountTypes(v []string) {
	o.AccountTypes = &v
}

// GetRequiresInstallationInstructions returns the RequiresInstallationInstructions field value if set, zero value otherwise.
func (o *AppManifest) GetRequiresInstallationInstructions() bool {
	if o == nil || o.RequiresInstallationInstructions == nil {
		var ret bool
		return ret
	}
	return *o.RequiresInstallationInstructions
}

// GetRequiresInstallationInstructionsOk returns a tuple with the RequiresInstallationInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetRequiresInstallationInstructionsOk() (*bool, bool) {
	if o == nil || o.RequiresInstallationInstructions == nil {
		return nil, false
	}
	return o.RequiresInstallationInstructions, true
}

// HasRequiresInstallationInstructions returns a boolean if a field has been set.
func (o *AppManifest) HasRequiresInstallationInstructions() bool {
	if o != nil && o.RequiresInstallationInstructions != nil {
		return true
	}

	return false
}

// SetRequiresInstallationInstructions gets a reference to the given bool and assigns it to the RequiresInstallationInstructions field.
func (o *AppManifest) SetRequiresInstallationInstructions(v bool) {
	o.RequiresInstallationInstructions = &v
}

// GetInstallationInstructions returns the InstallationInstructions field value if set, zero value otherwise.
func (o *AppManifest) GetInstallationInstructions() string {
	if o == nil || o.InstallationInstructions == nil {
		var ret string
		return ret
	}
	return *o.InstallationInstructions
}

// GetInstallationInstructionsOk returns a tuple with the InstallationInstructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetInstallationInstructionsOk() (*string, bool) {
	if o == nil || o.InstallationInstructions == nil {
		return nil, false
	}
	return o.InstallationInstructions, true
}

// HasInstallationInstructions returns a boolean if a field has been set.
func (o *AppManifest) HasInstallationInstructions() bool {
	if o != nil && o.InstallationInstructions != nil {
		return true
	}

	return false
}

// SetInstallationInstructions gets a reference to the given string and assigns it to the InstallationInstructions field.
func (o *AppManifest) SetInstallationInstructions(v string) {
	o.InstallationInstructions = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *AppManifest) GetParameters() []ServiceManifestDataSourceParameter {
	if o == nil || o.Parameters == nil {
		var ret []ServiceManifestDataSourceParameter
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetParametersOk() (*[]ServiceManifestDataSourceParameter, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *AppManifest) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given []ServiceManifestDataSourceParameter and assigns it to the Parameters field.
func (o *AppManifest) SetParameters(v []ServiceManifestDataSourceParameter) {
	o.Parameters = &v
}

// GetAuthor returns the Author field value if set, zero value otherwise.
func (o *AppManifest) GetAuthor() string {
	if o == nil || o.Author == nil {
		var ret string
		return ret
	}
	return *o.Author
}

// GetAuthorOk returns a tuple with the Author field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetAuthorOk() (*string, bool) {
	if o == nil || o.Author == nil {
		return nil, false
	}
	return o.Author, true
}

// HasAuthor returns a boolean if a field has been set.
func (o *AppManifest) HasAuthor() bool {
	if o != nil && o.Author != nil {
		return true
	}

	return false
}

// SetAuthor gets a reference to the given string and assigns it to the Author field.
func (o *AppManifest) SetAuthor(v string) {
	o.Author = &v
}

// GetAuthorWebsite returns the AuthorWebsite field value if set, zero value otherwise.
func (o *AppManifest) GetAuthorWebsite() string {
	if o == nil || o.AuthorWebsite == nil {
		var ret string
		return ret
	}
	return *o.AuthorWebsite
}

// GetAuthorWebsiteOk returns a tuple with the AuthorWebsite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppManifest) GetAuthorWebsiteOk() (*string, bool) {
	if o == nil || o.AuthorWebsite == nil {
		return nil, false
	}
	return o.AuthorWebsite, true
}

// HasAuthorWebsite returns a boolean if a field has been set.
func (o *AppManifest) HasAuthorWebsite() bool {
	if o != nil && o.AuthorWebsite != nil {
		return true
	}

	return false
}

// SetAuthorWebsite gets a reference to the given string and assigns it to the AuthorWebsite field.
func (o *AppManifest) SetAuthorWebsite(v string) {
	o.AuthorWebsite = &v
}

func (o AppManifest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Family != nil {
		toSerialize["family"] = o.Family
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if o.Categories != nil {
		toSerialize["categories"] = o.Categories
	}
	if true {
		toSerialize["hoverText"] = o.HoverText
	}
	if true {
		toSerialize["iconURL"] = o.IconURL
	}
	if o.ScreenshotURLs != nil {
		toSerialize["screenshotURLs"] = o.ScreenshotURLs
	}
	if o.HelpURL != nil {
		toSerialize["helpURL"] = o.HelpURL
	}
	if o.HelpDocIdMap != nil {
		toSerialize["helpDocIdMap"] = o.HelpDocIdMap
	}
	if o.CommunityURL != nil {
		toSerialize["communityURL"] = o.CommunityURL
	}
	if o.Requirements != nil {
		toSerialize["requirements"] = o.Requirements
	}
	if o.AccountTypes != nil {
		toSerialize["accountTypes"] = o.AccountTypes
	}
	if o.RequiresInstallationInstructions != nil {
		toSerialize["requiresInstallationInstructions"] = o.RequiresInstallationInstructions
	}
	if o.InstallationInstructions != nil {
		toSerialize["installationInstructions"] = o.InstallationInstructions
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Author != nil {
		toSerialize["author"] = o.Author
	}
	if o.AuthorWebsite != nil {
		toSerialize["authorWebsite"] = o.AuthorWebsite
	}
	return json.Marshal(toSerialize)
}

type NullableAppManifest struct {
	value *AppManifest
	isSet bool
}

func (v NullableAppManifest) Get() *AppManifest {
	return v.value
}

func (v *NullableAppManifest) Set(val *AppManifest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppManifest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppManifest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppManifest(val *AppManifest) *NullableAppManifest {
	return &NullableAppManifest{value: val, isSet: true}
}

func (v NullableAppManifest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppManifest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



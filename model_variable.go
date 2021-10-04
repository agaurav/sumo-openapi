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

// Variable struct for Variable
type Variable struct {
	// Unique identifier for the variable.
	Id *string `json:"id,omitempty"`
	// Name of the variable. The variable name is case-insensitive. Only alphanumeric, and underscores are allowed in the variable name. 
	Name string `json:"name"`
	// Display name of the variable shown in the UI. If this field is empty, the name field will be used. The display name is case-insensitive. Only numbers, and underscores are allowed in the variable name. This field is not yet supported by the UI. 
	DisplayName *string `json:"displayName,omitempty"`
	// Default value of the variable.
	DefaultValue *string `json:"defaultValue,omitempty"`
	SourceDefinition VariableSourceDefinition `json:"sourceDefinition"`
	// Allow multiple selections in the values dropdown.
	AllowMultiSelect *bool `json:"allowMultiSelect,omitempty"`
	// Include an \"All\" option at the top of the variable's values dropdown.
	IncludeAllOption *bool `json:"includeAllOption,omitempty"`
	// Hide the variable in the dashboard UI.
	HideFromUI *bool `json:"hideFromUI,omitempty"`
}

// NewVariable instantiates a new Variable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariable(name string, sourceDefinition VariableSourceDefinition) *Variable {
	this := Variable{}
	this.Name = name
	this.SourceDefinition = sourceDefinition
	var allowMultiSelect bool = false
	this.AllowMultiSelect = &allowMultiSelect
	var includeAllOption bool = true
	this.IncludeAllOption = &includeAllOption
	var hideFromUI bool = false
	this.HideFromUI = &hideFromUI
	return &this
}

// NewVariableWithDefaults instantiates a new Variable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableWithDefaults() *Variable {
	this := Variable{}
	var allowMultiSelect bool = false
	this.AllowMultiSelect = &allowMultiSelect
	var includeAllOption bool = true
	this.IncludeAllOption = &includeAllOption
	var hideFromUI bool = false
	this.HideFromUI = &hideFromUI
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Variable) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Variable) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Variable) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *Variable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Variable) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Variable) SetName(v string) {
	o.Name = v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *Variable) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *Variable) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *Variable) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetDefaultValue returns the DefaultValue field value if set, zero value otherwise.
func (o *Variable) GetDefaultValue() string {
	if o == nil || o.DefaultValue == nil {
		var ret string
		return ret
	}
	return *o.DefaultValue
}

// GetDefaultValueOk returns a tuple with the DefaultValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetDefaultValueOk() (*string, bool) {
	if o == nil || o.DefaultValue == nil {
		return nil, false
	}
	return o.DefaultValue, true
}

// HasDefaultValue returns a boolean if a field has been set.
func (o *Variable) HasDefaultValue() bool {
	if o != nil && o.DefaultValue != nil {
		return true
	}

	return false
}

// SetDefaultValue gets a reference to the given string and assigns it to the DefaultValue field.
func (o *Variable) SetDefaultValue(v string) {
	o.DefaultValue = &v
}

// GetSourceDefinition returns the SourceDefinition field value
func (o *Variable) GetSourceDefinition() VariableSourceDefinition {
	if o == nil {
		var ret VariableSourceDefinition
		return ret
	}

	return o.SourceDefinition
}

// GetSourceDefinitionOk returns a tuple with the SourceDefinition field value
// and a boolean to check if the value has been set.
func (o *Variable) GetSourceDefinitionOk() (*VariableSourceDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SourceDefinition, true
}

// SetSourceDefinition sets field value
func (o *Variable) SetSourceDefinition(v VariableSourceDefinition) {
	o.SourceDefinition = v
}

// GetAllowMultiSelect returns the AllowMultiSelect field value if set, zero value otherwise.
func (o *Variable) GetAllowMultiSelect() bool {
	if o == nil || o.AllowMultiSelect == nil {
		var ret bool
		return ret
	}
	return *o.AllowMultiSelect
}

// GetAllowMultiSelectOk returns a tuple with the AllowMultiSelect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetAllowMultiSelectOk() (*bool, bool) {
	if o == nil || o.AllowMultiSelect == nil {
		return nil, false
	}
	return o.AllowMultiSelect, true
}

// HasAllowMultiSelect returns a boolean if a field has been set.
func (o *Variable) HasAllowMultiSelect() bool {
	if o != nil && o.AllowMultiSelect != nil {
		return true
	}

	return false
}

// SetAllowMultiSelect gets a reference to the given bool and assigns it to the AllowMultiSelect field.
func (o *Variable) SetAllowMultiSelect(v bool) {
	o.AllowMultiSelect = &v
}

// GetIncludeAllOption returns the IncludeAllOption field value if set, zero value otherwise.
func (o *Variable) GetIncludeAllOption() bool {
	if o == nil || o.IncludeAllOption == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAllOption
}

// GetIncludeAllOptionOk returns a tuple with the IncludeAllOption field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetIncludeAllOptionOk() (*bool, bool) {
	if o == nil || o.IncludeAllOption == nil {
		return nil, false
	}
	return o.IncludeAllOption, true
}

// HasIncludeAllOption returns a boolean if a field has been set.
func (o *Variable) HasIncludeAllOption() bool {
	if o != nil && o.IncludeAllOption != nil {
		return true
	}

	return false
}

// SetIncludeAllOption gets a reference to the given bool and assigns it to the IncludeAllOption field.
func (o *Variable) SetIncludeAllOption(v bool) {
	o.IncludeAllOption = &v
}

// GetHideFromUI returns the HideFromUI field value if set, zero value otherwise.
func (o *Variable) GetHideFromUI() bool {
	if o == nil || o.HideFromUI == nil {
		var ret bool
		return ret
	}
	return *o.HideFromUI
}

// GetHideFromUIOk returns a tuple with the HideFromUI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Variable) GetHideFromUIOk() (*bool, bool) {
	if o == nil || o.HideFromUI == nil {
		return nil, false
	}
	return o.HideFromUI, true
}

// HasHideFromUI returns a boolean if a field has been set.
func (o *Variable) HasHideFromUI() bool {
	if o != nil && o.HideFromUI != nil {
		return true
	}

	return false
}

// SetHideFromUI gets a reference to the given bool and assigns it to the HideFromUI field.
func (o *Variable) SetHideFromUI(v bool) {
	o.HideFromUI = &v
}

func (o Variable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if o.DefaultValue != nil {
		toSerialize["defaultValue"] = o.DefaultValue
	}
	if true {
		toSerialize["sourceDefinition"] = o.SourceDefinition
	}
	if o.AllowMultiSelect != nil {
		toSerialize["allowMultiSelect"] = o.AllowMultiSelect
	}
	if o.IncludeAllOption != nil {
		toSerialize["includeAllOption"] = o.IncludeAllOption
	}
	if o.HideFromUI != nil {
		toSerialize["hideFromUI"] = o.HideFromUI
	}
	return json.Marshal(toSerialize)
}

type NullableVariable struct {
	value *Variable
	isSet bool
}

func (v NullableVariable) Get() *Variable {
	return v.value
}

func (v *NullableVariable) Set(val *Variable) {
	v.value = val
	v.isSet = true
}

func (v NullableVariable) IsSet() bool {
	return v.isSet
}

func (v *NullableVariable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariable(val *Variable) *NullableVariable {
	return &NullableVariable{value: val, isSet: true}
}

func (v NullableVariable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



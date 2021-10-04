/*
Sumo Logic API

# Getting Started Welcome to the Sumo Logic API reference. You can use these APIs to interact with the Sumo Logic platform. For information on the collector and search APIs see our [API home page](https://help.sumologic.com/APIs). ## API Endpoints Sumo Logic has several deployments in different geographic locations. You'll need to use the Sumo Logic API endpoint corresponding to your geographic location. See the table below for the different API endpoints by deployment. For details determining your account's deployment see [API endpoints](https://help.sumologic.com/?cid=3011).    <table>     <tr>       <td> <strong>Deployment</strong> </td>       <td> <strong>Endpoint</strong> </td>     </tr>     <tr>       <td> AU </td>       <td> https://api.au.sumologic.com/api/ </td>     </tr>     <tr>       <td> CA </td>       <td> https://api.ca.sumologic.com/api/ </td>     </tr>     <tr>       <td> DE </td>       <td> https://api.de.sumologic.com/api/ </td>     </tr>     <tr>       <td> EU </td>       <td> https://api.eu.sumologic.com/api/ </td>     </tr>     <tr>       <td> FED </td>       <td> https://api.fed.sumologic.com/api/ </td>     </tr>     <tr>       <td> IN </td>       <td> https://api.in.sumologic.com/api/ </td>     </tr>     <tr>       <td> JP </td>       <td> https://api.jp.sumologic.com/api/ </td>     </tr>     <tr>       <td> US1 </td>       <td> https://api.sumologic.com/api/ </td>     </tr>     <tr>       <td> US2 </td>       <td> https://api.us2.sumologic.com/api/ </td>     </tr>   </table>  ## Authentication Sumo Logic supports the following options for API authentication: - Access ID and Access Key - Base64 encoded Access ID and Access Key  See [Access Keys](https://help.sumologic.com/Manage/Security/Access-Keys) to generate an Access Key. Make sure to copy the key you create, because it is displayed only once. When you have an Access ID and Access Key you can execute requests such as the following:   ```bash   curl -u \"<accessId>:<accessKey>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```  Where `deployment` is either `au`, `ca`, `de`, `eu`, `fed`, `in`, `jp`, `us1`, or `us2`. See [API endpoints](#section/API-Endpoints) for details.  If you prefer to use basic access authentication, you can do a Base64 encoding of your `<accessId>:<accessKey>` to authenticate your HTTPS request. The following is an example request, replace the placeholder `<encoded>` with your encoded Access ID and Access Key string:   ```bash   curl -H \"Authorization: Basic <encoded>\" -X GET https://api.<deployment>.sumologic.com/api/v1/users   ```   Refer to [API Authentication](https://help.sumologic.com/?cid=3012) for a Base64 example.  ## Status Codes Generic status codes that apply to all our APIs. See the [HTTP status code registry](https://www.iana.org/assignments/http-status-codes/http-status-codes.xhtml) for reference.   <table>     <tr>       <td> <strong>HTTP Status Code</strong> </td>       <td> <strong>Error Code</strong> </td>       <td> <strong>Description</strong> </td>     </tr>     <tr>       <td> 301 </td>       <td> moved </td>       <td> The requested resource SHOULD be accessed through returned URI in Location Header. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-301-Error-Moved) for details.</td>     </tr>     <tr>       <td> 401 </td>       <td> unauthorized </td>       <td> Credential could not be verified.</td>     </tr>     <tr>       <td> 403 </td>       <td> forbidden </td>       <td> This operation is not allowed for your account type or the user doesn't have the role capability to perform this action. See [troubleshooting](https://help.sumologic.com/APIs/Troubleshooting-APIs/API-403-Error-This-operation-is-not-allowed-for-your-account-type) for details.</td>     </tr>     <tr>       <td> 404 </td>       <td> notfound </td>       <td> Requested resource could not be found. </td>     </tr>     <tr>       <td> 405 </td>       <td> method.unsupported </td>       <td> Unsupported method for URL. </td>     </tr>     <tr>       <td> 415 </td>       <td> contenttype.invalid </td>       <td> Invalid content type. </td>     </tr>     <tr>       <td> 429 </td>       <td> rate.limit.exceeded </td>       <td> The API request rate is higher than 4 request per second or inflight API requests are higher than 10 request per second. </td>     </tr>     <tr>       <td> 500 </td>       <td> internal.error </td>       <td> Internal server error. </td>     </tr>     <tr>       <td> 503 </td>       <td> service.unavailable </td>       <td> Service is currently unavailable. </td>     </tr>   </table>  ## Filtering Some API endpoints support filtering results on a specified set of fields. Each endpoint that supports filtering will list the fields that can be filtered. Multiple fields can be combined by using an ampersand `&` character.  For example, to get 20 users whose `firstName` is `John` and `lastName` is `Doe`:   ```bash   api.sumologic.com/v1/users?limit=20&firstName=John&lastName=Doe   ```  ## Sorting Some API endpoints support sorting fields by using the `sortBy` query parameter. The default sort order is ascending. Prefix the field with a minus sign `-` to sort in descending order.  For example, to get 20 users sorted by their `email` in descending order:   ```bash   api.sumologic.com/v1/users?limit=20&sort=-email   ```  ## Asynchronous Request Asynchronous requests do not wait for results, instead they immediately respond back with a job identifier while the job runs in the background. You can use the job identifier to track the status of the asynchronous job request. Here is a typical flow for an asynchronous request. 1. Start an asynchronous job. On success, a job identifier is returned. The job identifier uniquely identifies   your asynchronous job.  2. Once started, use the job identifier from step 1 to track the status of your asynchronous job. An asynchronous   request will typically provide an endpoint to poll for the status of asynchronous job. A successful response   from the status endpoint will have the following structure:   ```json   {       \"status\": \"Status of asynchronous request\",       \"statusMessage\": \"Optional message with additional information in case request succeeds\",       \"error\": \"Error object in case request fails\"   }   ```   The `status` field can have one of the following values:     1. `Success`: The job succeeded. The `statusMessage` field might have additional information.     2. `InProgress`: The job is still running.     3. `Failed`: The job failed. The `error` field in the response will have more information about the failure.  3. Some asynchronous APIs may provide a third endpoint (like [export result](#operation/getAsyncExportResult))   to fetch the result of an asynchronous job.   ### Example Let's say we want to export a folder with the identifier `0000000006A2E86F`. We will use the [async export](#operation/beginAsyncExport) API to export all the content under the folder with `id=0000000006A2E86F`. 1. Start an export job for the folder   ```bash   curl -X POST -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export   ```   See [authentication section](#section/Authentication) for more details about `accessId`, `accessKey`, and   `deployment`.   On success, you will get back a job identifier. In the response below, `C03E086C137F38B4` is the job identifier.   ```bash   {       \"id\": \"C03E086C137F38B4\"   }   ```  2. Now poll for the status of the asynchronous job with the [status](#operation/getAsyncExportStatus) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/status   ```   You may get a response like   ```json   {       \"status\": \"InProgress\",       \"statusMessage\": null,       \"error\": null   }   ```   It implies the job is still in progress. Keep polling till the status is either `Success` or `Failed`.  3. When the asynchronous job completes (`status != \"InProgress\"`), you can fetch the results with the   [export result](#operation/getAsyncExportResult) endpoint.   ```bash   curl -X GET -u \"<accessId>:<accessKey>\" https://api.<deployment>.sumologic.com/api/v2/content/0000000006A2E86F/export/C03E086C137F38B4/result   ```    The asynchronous job may fail (`status == \"Failed\"`). You can look at the `error` field for more details.   ```json   {       \"status\": \"Failed\",       \"errors\": {           \"code\": \"content1:too_many_items\",           \"message\": \"Too many objects: object count(1100) was greater than limit 1000\"       }   }   ```   ## Rate Limiting * A rate limit of four API requests per second (240 requests per minute) applies to all API calls from a user. * A rate limit of 10 concurrent requests to any API endpoint applies to an access key.  If a rate is exceeded, a rate limit exceeded 429 status code is returned.  ## Generating Clients You can use [OpenAPI Generator](https://openapi-generator.tech) to generate clients from the YAML file to access the API.  ### Using [NPM](https://www.npmjs.com/get-npm) 1. Install [NPM package wrapper](https://github.com/openapitools/openapi-generator-cli) globally, exposing the CLI   on the command line:   ```bash   npm install @openapitools/openapi-generator-cli -g   ```   You can see detailed instructions [here](https://openapi-generator.tech/docs/installation#npm).  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ```   ### Using [Homebrew](https://brew.sh/) 1. Install OpenAPI Generator   ```bash   brew install openapi-generator   ```  2. Download the [YAML file](/docs/sumologic-api.yaml) and save it locally. Let's say the file is saved as `sumologic-api.yaml`. 3. Use the following command to generate `python` client side code inside the `sumo/client/python` directory:   ```bash   openapi-generator generate -i sumologic-api.yaml -g python -o sumo/client/python   ``` 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// SavedSearchSyncDefinition struct for SavedSearchSyncDefinition
type SavedSearchSyncDefinition struct {
	// The text of a Sumo Logic query.
	QueryText string `json:"queryText"`
	// Default time range for the search. Possible types of time ranges are:   - relative time range: e.g. \"-1d -12h\" represents a time range from one day ago to 12 hours ago.   - absolute time range: e.g. \"01-04-2017 20:32:00 to 01-04-2017 20:35:00\" represents a time range     from April 1st, 2017 at 8:32 PM until April 1st, 2017 at 8:35 PM.
	DefaultTimeRange string `json:"defaultTimeRange"`
	// Set it to true to run the search using receipt time. By default, searches do not run by receipt time.
	ByReceiptTime bool `json:"byReceiptTime"`
	// The name of the Scheduled View that has indexed the data you want to search.
	ViewName *string `json:"viewName,omitempty"`
	// Start timestamp of the Scheduled View in UTC format.
	ViewStartTime *time.Time `json:"viewStartTime,omitempty"`
	// An array of search query parameter objects.
	QueryParameters []QueryParameterSyncDefinition `json:"queryParameters"`
	// Define the parsing mode to scan the JSON format log messages. Possible values are:   1. `AutoParse`   2. `Manual` In AutoParse mode, the system automatically figures out fields to parse based on the search query. While in the Manual mode, no fields are parsed out automatically. For more information see [Dynamic Parsing](https://help.sumologic.com/?cid=0011).
	ParsingMode *string `json:"parsingMode,omitempty"`
}

// NewSavedSearchSyncDefinition instantiates a new SavedSearchSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSavedSearchSyncDefinition(queryText string, defaultTimeRange string, byReceiptTime bool, queryParameters []QueryParameterSyncDefinition) *SavedSearchSyncDefinition {
	this := SavedSearchSyncDefinition{}
	this.QueryText = queryText
	this.DefaultTimeRange = defaultTimeRange
	this.ByReceiptTime = byReceiptTime
	this.QueryParameters = queryParameters
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	return &this
}

// NewSavedSearchSyncDefinitionWithDefaults instantiates a new SavedSearchSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSavedSearchSyncDefinitionWithDefaults() *SavedSearchSyncDefinition {
	this := SavedSearchSyncDefinition{}
	var byReceiptTime bool = false
	this.ByReceiptTime = byReceiptTime
	var parsingMode string = "Manual"
	this.ParsingMode = &parsingMode
	return &this
}

// GetQueryText returns the QueryText field value
func (o *SavedSearchSyncDefinition) GetQueryText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryText
}

// GetQueryTextOk returns a tuple with the QueryText field value
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinition) GetQueryTextOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QueryText, true
}

// SetQueryText sets field value
func (o *SavedSearchSyncDefinition) SetQueryText(v string) {
	o.QueryText = v
}

// GetDefaultTimeRange returns the DefaultTimeRange field value
func (o *SavedSearchSyncDefinition) GetDefaultTimeRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultTimeRange
}

// GetDefaultTimeRangeOk returns a tuple with the DefaultTimeRange field value
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinition) GetDefaultTimeRangeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultTimeRange, true
}

// SetDefaultTimeRange sets field value
func (o *SavedSearchSyncDefinition) SetDefaultTimeRange(v string) {
	o.DefaultTimeRange = v
}

// GetByReceiptTime returns the ByReceiptTime field value
func (o *SavedSearchSyncDefinition) GetByReceiptTime() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ByReceiptTime
}

// GetByReceiptTimeOk returns a tuple with the ByReceiptTime field value
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinition) GetByReceiptTimeOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ByReceiptTime, true
}

// SetByReceiptTime sets field value
func (o *SavedSearchSyncDefinition) SetByReceiptTime(v bool) {
	o.ByReceiptTime = v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *SavedSearchSyncDefinition) GetViewName() string {
	if o == nil || o.ViewName == nil {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinition) GetViewNameOk() (*string, bool) {
	if o == nil || o.ViewName == nil {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *SavedSearchSyncDefinition) HasViewName() bool {
	if o != nil && o.ViewName != nil {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *SavedSearchSyncDefinition) SetViewName(v string) {
	o.ViewName = &v
}

// GetViewStartTime returns the ViewStartTime field value if set, zero value otherwise.
func (o *SavedSearchSyncDefinition) GetViewStartTime() time.Time {
	if o == nil || o.ViewStartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.ViewStartTime
}

// GetViewStartTimeOk returns a tuple with the ViewStartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinition) GetViewStartTimeOk() (*time.Time, bool) {
	if o == nil || o.ViewStartTime == nil {
		return nil, false
	}
	return o.ViewStartTime, true
}

// HasViewStartTime returns a boolean if a field has been set.
func (o *SavedSearchSyncDefinition) HasViewStartTime() bool {
	if o != nil && o.ViewStartTime != nil {
		return true
	}

	return false
}

// SetViewStartTime gets a reference to the given time.Time and assigns it to the ViewStartTime field.
func (o *SavedSearchSyncDefinition) SetViewStartTime(v time.Time) {
	o.ViewStartTime = &v
}

// GetQueryParameters returns the QueryParameters field value
func (o *SavedSearchSyncDefinition) GetQueryParameters() []QueryParameterSyncDefinition {
	if o == nil {
		var ret []QueryParameterSyncDefinition
		return ret
	}

	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinition) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QueryParameters, true
}

// SetQueryParameters sets field value
func (o *SavedSearchSyncDefinition) SetQueryParameters(v []QueryParameterSyncDefinition) {
	o.QueryParameters = v
}

// GetParsingMode returns the ParsingMode field value if set, zero value otherwise.
func (o *SavedSearchSyncDefinition) GetParsingMode() string {
	if o == nil || o.ParsingMode == nil {
		var ret string
		return ret
	}
	return *o.ParsingMode
}

// GetParsingModeOk returns a tuple with the ParsingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SavedSearchSyncDefinition) GetParsingModeOk() (*string, bool) {
	if o == nil || o.ParsingMode == nil {
		return nil, false
	}
	return o.ParsingMode, true
}

// HasParsingMode returns a boolean if a field has been set.
func (o *SavedSearchSyncDefinition) HasParsingMode() bool {
	if o != nil && o.ParsingMode != nil {
		return true
	}

	return false
}

// SetParsingMode gets a reference to the given string and assigns it to the ParsingMode field.
func (o *SavedSearchSyncDefinition) SetParsingMode(v string) {
	o.ParsingMode = &v
}

func (o SavedSearchSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["queryText"] = o.QueryText
	}
	if true {
		toSerialize["defaultTimeRange"] = o.DefaultTimeRange
	}
	if true {
		toSerialize["byReceiptTime"] = o.ByReceiptTime
	}
	if o.ViewName != nil {
		toSerialize["viewName"] = o.ViewName
	}
	if o.ViewStartTime != nil {
		toSerialize["viewStartTime"] = o.ViewStartTime
	}
	if true {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if o.ParsingMode != nil {
		toSerialize["parsingMode"] = o.ParsingMode
	}
	return json.Marshal(toSerialize)
}

type NullableSavedSearchSyncDefinition struct {
	value *SavedSearchSyncDefinition
	isSet bool
}

func (v NullableSavedSearchSyncDefinition) Get() *SavedSearchSyncDefinition {
	return v.value
}

func (v *NullableSavedSearchSyncDefinition) Set(val *SavedSearchSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSavedSearchSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSavedSearchSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSavedSearchSyncDefinition(val *SavedSearchSyncDefinition) *NullableSavedSearchSyncDefinition {
	return &NullableSavedSearchSyncDefinition{value: val, isSet: true}
}

func (v NullableSavedSearchSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSavedSearchSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



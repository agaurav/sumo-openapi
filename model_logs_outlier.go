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

// LogsOutlier The parameters extracted from the logs outlier query.
type LogsOutlier struct {
	// The query string after trimming out the outlier clause.
	TrimmedQuery *string `json:"trimmedQuery,omitempty"`
	// Sets the trailing number of data points to calculate mean and sigma.
	Window *int64 `json:"window,omitempty"`
	// Sets the required number of consecutive indicator data points (outliers) to trigger a violation.
	Consecutive *int64 `json:"consecutive,omitempty"`
	// Specifies which direction should trigger violations. Valid values:   1. `Both`: Both positive and negative deviations   2. `Up`: Positive deviations only   3. `Down`: Negative deviations only example: \"Up\" pattern: \"^(Both|Up|Down)$\" default: \"Both\" x-pattern-message: \"should be one of the following: 'Both', 'Up', 'Down'\"
	Direction *string `json:"direction,omitempty"`
	// Sets the number of standard deviations for calculating violations.
	Threshold *float64 `json:"threshold,omitempty"`
	// The name of the field that the trigger condition will alert on.
	Field *string `json:"field,omitempty"`
}

// NewLogsOutlier instantiates a new LogsOutlier object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLogsOutlier() *LogsOutlier {
	this := LogsOutlier{}
	var window int64 = 10
	this.Window = &window
	var consecutive int64 = 1
	this.Consecutive = &consecutive
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// NewLogsOutlierWithDefaults instantiates a new LogsOutlier object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLogsOutlierWithDefaults() *LogsOutlier {
	this := LogsOutlier{}
	var window int64 = 10
	this.Window = &window
	var consecutive int64 = 1
	this.Consecutive = &consecutive
	var threshold float64 = 3.0
	this.Threshold = &threshold
	return &this
}

// GetTrimmedQuery returns the TrimmedQuery field value if set, zero value otherwise.
func (o *LogsOutlier) GetTrimmedQuery() string {
	if o == nil || o.TrimmedQuery == nil {
		var ret string
		return ret
	}
	return *o.TrimmedQuery
}

// GetTrimmedQueryOk returns a tuple with the TrimmedQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsOutlier) GetTrimmedQueryOk() (*string, bool) {
	if o == nil || o.TrimmedQuery == nil {
		return nil, false
	}
	return o.TrimmedQuery, true
}

// HasTrimmedQuery returns a boolean if a field has been set.
func (o *LogsOutlier) HasTrimmedQuery() bool {
	if o != nil && o.TrimmedQuery != nil {
		return true
	}

	return false
}

// SetTrimmedQuery gets a reference to the given string and assigns it to the TrimmedQuery field.
func (o *LogsOutlier) SetTrimmedQuery(v string) {
	o.TrimmedQuery = &v
}

// GetWindow returns the Window field value if set, zero value otherwise.
func (o *LogsOutlier) GetWindow() int64 {
	if o == nil || o.Window == nil {
		var ret int64
		return ret
	}
	return *o.Window
}

// GetWindowOk returns a tuple with the Window field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsOutlier) GetWindowOk() (*int64, bool) {
	if o == nil || o.Window == nil {
		return nil, false
	}
	return o.Window, true
}

// HasWindow returns a boolean if a field has been set.
func (o *LogsOutlier) HasWindow() bool {
	if o != nil && o.Window != nil {
		return true
	}

	return false
}

// SetWindow gets a reference to the given int64 and assigns it to the Window field.
func (o *LogsOutlier) SetWindow(v int64) {
	o.Window = &v
}

// GetConsecutive returns the Consecutive field value if set, zero value otherwise.
func (o *LogsOutlier) GetConsecutive() int64 {
	if o == nil || o.Consecutive == nil {
		var ret int64
		return ret
	}
	return *o.Consecutive
}

// GetConsecutiveOk returns a tuple with the Consecutive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsOutlier) GetConsecutiveOk() (*int64, bool) {
	if o == nil || o.Consecutive == nil {
		return nil, false
	}
	return o.Consecutive, true
}

// HasConsecutive returns a boolean if a field has been set.
func (o *LogsOutlier) HasConsecutive() bool {
	if o != nil && o.Consecutive != nil {
		return true
	}

	return false
}

// SetConsecutive gets a reference to the given int64 and assigns it to the Consecutive field.
func (o *LogsOutlier) SetConsecutive(v int64) {
	o.Consecutive = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *LogsOutlier) GetDirection() string {
	if o == nil || o.Direction == nil {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsOutlier) GetDirectionOk() (*string, bool) {
	if o == nil || o.Direction == nil {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *LogsOutlier) HasDirection() bool {
	if o != nil && o.Direction != nil {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *LogsOutlier) SetDirection(v string) {
	o.Direction = &v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *LogsOutlier) GetThreshold() float64 {
	if o == nil || o.Threshold == nil {
		var ret float64
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsOutlier) GetThresholdOk() (*float64, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *LogsOutlier) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given float64 and assigns it to the Threshold field.
func (o *LogsOutlier) SetThreshold(v float64) {
	o.Threshold = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *LogsOutlier) GetField() string {
	if o == nil || o.Field == nil {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LogsOutlier) GetFieldOk() (*string, bool) {
	if o == nil || o.Field == nil {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *LogsOutlier) HasField() bool {
	if o != nil && o.Field != nil {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *LogsOutlier) SetField(v string) {
	o.Field = &v
}

func (o LogsOutlier) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.TrimmedQuery != nil {
		toSerialize["trimmedQuery"] = o.TrimmedQuery
	}
	if o.Window != nil {
		toSerialize["window"] = o.Window
	}
	if o.Consecutive != nil {
		toSerialize["consecutive"] = o.Consecutive
	}
	if o.Direction != nil {
		toSerialize["direction"] = o.Direction
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if o.Field != nil {
		toSerialize["field"] = o.Field
	}
	return json.Marshal(toSerialize)
}

type NullableLogsOutlier struct {
	value *LogsOutlier
	isSet bool
}

func (v NullableLogsOutlier) Get() *LogsOutlier {
	return v.value
}

func (v *NullableLogsOutlier) Set(val *LogsOutlier) {
	v.value = val
	v.isSet = true
}

func (v NullableLogsOutlier) IsSet() bool {
	return v.isSet
}

func (v *NullableLogsOutlier) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLogsOutlier(val *LogsOutlier) *NullableLogsOutlier {
	return &NullableLogsOutlier{value: val, isSet: true}
}

func (v NullableLogsOutlier) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLogsOutlier) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



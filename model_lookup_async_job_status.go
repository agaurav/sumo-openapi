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

// LookupAsyncJobStatus Lookup table async job status.
type LookupAsyncJobStatus struct {
	// An identifier returned in response to an asynchronous request.
	JobId string `json:"jobId"`
	// Whether or not the request is pending (`Pending`), in progress (`InProgress`), has completed successfully (`Success`), has completed partially with warnings (`PartialSuccess`) or has completed with an error (`Failed`).
	Status string `json:"status"`
	// Additional status messages generated if any if the status is `Success`.
	StatusMessages *[]string `json:"statusMessages,omitempty"`
	// More information about the failures, if the status is `Failed`.
	Errors *[]ErrorDescription `json:"errors,omitempty"`
	// More information about the warnings, if the status is `PartialSuccess`.
	Warnings *[]WarningDescription `json:"warnings,omitempty"`
	// Content id of lookup table on which this operation was performed.
	LookupContentId string `json:"lookupContentId"`
	// Name of lookup table on which this operation was performed.
	LookupName string `json:"lookupName"`
	// Content path of lookup table on which this operation was performed.
	LookupContentPath string `json:"lookupContentPath"`
	// Type of asynchronous request made:   - `BulkMerge`   - `BulkReplace`   - `Truncate`
	RequestType *string `json:"requestType,omitempty"`
	// User id of user who initiated this operation.
	UserId string `json:"userId"`
	// Creation time of this job in UTC.
	CreatedAt time.Time `json:"createdAt"`
	// Timestamp in UTC when status was last updated.
	ModifiedAt time.Time `json:"modifiedAt"`
}

// NewLookupAsyncJobStatus instantiates a new LookupAsyncJobStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLookupAsyncJobStatus(jobId string, status string, lookupContentId string, lookupName string, lookupContentPath string, userId string, createdAt time.Time, modifiedAt time.Time) *LookupAsyncJobStatus {
	this := LookupAsyncJobStatus{}
	this.JobId = jobId
	this.Status = status
	this.LookupContentId = lookupContentId
	this.LookupName = lookupName
	this.LookupContentPath = lookupContentPath
	this.UserId = userId
	this.CreatedAt = createdAt
	this.ModifiedAt = modifiedAt
	return &this
}

// NewLookupAsyncJobStatusWithDefaults instantiates a new LookupAsyncJobStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLookupAsyncJobStatusWithDefaults() *LookupAsyncJobStatus {
	this := LookupAsyncJobStatus{}
	return &this
}

// GetJobId returns the JobId field value
func (o *LookupAsyncJobStatus) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetJobIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *LookupAsyncJobStatus) SetJobId(v string) {
	o.JobId = v
}

// GetStatus returns the Status field value
func (o *LookupAsyncJobStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *LookupAsyncJobStatus) SetStatus(v string) {
	o.Status = v
}

// GetStatusMessages returns the StatusMessages field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetStatusMessages() []string {
	if o == nil || o.StatusMessages == nil {
		var ret []string
		return ret
	}
	return *o.StatusMessages
}

// GetStatusMessagesOk returns a tuple with the StatusMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetStatusMessagesOk() (*[]string, bool) {
	if o == nil || o.StatusMessages == nil {
		return nil, false
	}
	return o.StatusMessages, true
}

// HasStatusMessages returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasStatusMessages() bool {
	if o != nil && o.StatusMessages != nil {
		return true
	}

	return false
}

// SetStatusMessages gets a reference to the given []string and assigns it to the StatusMessages field.
func (o *LookupAsyncJobStatus) SetStatusMessages(v []string) {
	o.StatusMessages = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetErrors() []ErrorDescription {
	if o == nil || o.Errors == nil {
		var ret []ErrorDescription
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetErrorsOk() (*[]ErrorDescription, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ErrorDescription and assigns it to the Errors field.
func (o *LookupAsyncJobStatus) SetErrors(v []ErrorDescription) {
	o.Errors = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetWarnings() []WarningDescription {
	if o == nil || o.Warnings == nil {
		var ret []WarningDescription
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetWarningsOk() (*[]WarningDescription, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningDescription and assigns it to the Warnings field.
func (o *LookupAsyncJobStatus) SetWarnings(v []WarningDescription) {
	o.Warnings = &v
}

// GetLookupContentId returns the LookupContentId field value
func (o *LookupAsyncJobStatus) GetLookupContentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupContentId
}

// GetLookupContentIdOk returns a tuple with the LookupContentId field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetLookupContentIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LookupContentId, true
}

// SetLookupContentId sets field value
func (o *LookupAsyncJobStatus) SetLookupContentId(v string) {
	o.LookupContentId = v
}

// GetLookupName returns the LookupName field value
func (o *LookupAsyncJobStatus) GetLookupName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupName
}

// GetLookupNameOk returns a tuple with the LookupName field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetLookupNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LookupName, true
}

// SetLookupName sets field value
func (o *LookupAsyncJobStatus) SetLookupName(v string) {
	o.LookupName = v
}

// GetLookupContentPath returns the LookupContentPath field value
func (o *LookupAsyncJobStatus) GetLookupContentPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LookupContentPath
}

// GetLookupContentPathOk returns a tuple with the LookupContentPath field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetLookupContentPathOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LookupContentPath, true
}

// SetLookupContentPath sets field value
func (o *LookupAsyncJobStatus) SetLookupContentPath(v string) {
	o.LookupContentPath = v
}

// GetRequestType returns the RequestType field value if set, zero value otherwise.
func (o *LookupAsyncJobStatus) GetRequestType() string {
	if o == nil || o.RequestType == nil {
		var ret string
		return ret
	}
	return *o.RequestType
}

// GetRequestTypeOk returns a tuple with the RequestType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetRequestTypeOk() (*string, bool) {
	if o == nil || o.RequestType == nil {
		return nil, false
	}
	return o.RequestType, true
}

// HasRequestType returns a boolean if a field has been set.
func (o *LookupAsyncJobStatus) HasRequestType() bool {
	if o != nil && o.RequestType != nil {
		return true
	}

	return false
}

// SetRequestType gets a reference to the given string and assigns it to the RequestType field.
func (o *LookupAsyncJobStatus) SetRequestType(v string) {
	o.RequestType = &v
}

// GetUserId returns the UserId field value
func (o *LookupAsyncJobStatus) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetUserIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *LookupAsyncJobStatus) SetUserId(v string) {
	o.UserId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *LookupAsyncJobStatus) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *LookupAsyncJobStatus) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *LookupAsyncJobStatus) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *LookupAsyncJobStatus) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *LookupAsyncJobStatus) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

func (o LookupAsyncJobStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["jobId"] = o.JobId
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if o.StatusMessages != nil {
		toSerialize["statusMessages"] = o.StatusMessages
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if true {
		toSerialize["lookupContentId"] = o.LookupContentId
	}
	if true {
		toSerialize["lookupName"] = o.LookupName
	}
	if true {
		toSerialize["lookupContentPath"] = o.LookupContentPath
	}
	if o.RequestType != nil {
		toSerialize["requestType"] = o.RequestType
	}
	if true {
		toSerialize["userId"] = o.UserId
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	return json.Marshal(toSerialize)
}

type NullableLookupAsyncJobStatus struct {
	value *LookupAsyncJobStatus
	isSet bool
}

func (v NullableLookupAsyncJobStatus) Get() *LookupAsyncJobStatus {
	return v.value
}

func (v *NullableLookupAsyncJobStatus) Set(val *LookupAsyncJobStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableLookupAsyncJobStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableLookupAsyncJobStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLookupAsyncJobStatus(val *LookupAsyncJobStatus) *NullableLookupAsyncJobStatus {
	return &NullableLookupAsyncJobStatus{value: val, isSet: true}
}

func (v NullableLookupAsyncJobStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLookupAsyncJobStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



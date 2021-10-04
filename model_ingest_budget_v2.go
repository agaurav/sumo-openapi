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

// IngestBudgetV2 struct for IngestBudgetV2
type IngestBudgetV2 struct {
	// Display name of the ingest budget.
	Name string `json:"name"`
	// A scope is a constraint that will be used to identify the messages on which budget needs to be applied. A scope is consists of key and value separated by =. The field must be enabled in the fields table. Value supports wildcard. e.g. _sourceCategory=*prod*payment*, cluster=kafka. If the scope is defined _sourceCategory=*nginx* in this budget will be applied on messages having fields _sourceCategory=prod/nginx, _sourceCategory=dev/nginx, or _sourceCategory=dev/nginx/error
	Scope string `json:"scope"`
	// Capacity of the ingest budget, in bytes. It takes a few minutes for Collectors to stop collecting when capacity is reached. We recommend setting a soft limit that is lower than your needed hard limit.
	CapacityBytes int64 `json:"capacityBytes"`
	// Time zone of the reset time for the ingest budget. Follow the format in the [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones#List).
	Timezone string `json:"timezone"`
	// Reset time of the ingest budget in HH:MM format.
	ResetTime string `json:"resetTime"`
	// Description of the ingest budget.
	Description *string `json:"description,omitempty"`
	// Action to take when ingest budget's capacity is reached. All actions are audited. Supported values are:   * `stopCollecting`   * `keepCollecting`
	Action string `json:"action"`
	// The threshold as a percentage of when an ingest budget's capacity usage is logged in the Audit Index.
	AuditThreshold *int32 `json:"auditThreshold,omitempty"`
	// Unique identifier for the ingest budget.
	Id string `json:"id"`
	// Current usage since the last reset, in bytes.
	UsageBytes *int64 `json:"usageBytes,omitempty"`
	// Status of the current usage. Can be `Normal`, `Approaching`, `Exceeded`, or `Unknown` (unable to retrieve usage).
	UsageStatus *string `json:"usageStatus,omitempty"`
	// The creation timestamp in UTC of the Ingest Budget.
	CreatedAt time.Time `json:"createdAt"`
	// The identifier of the user who created the Ingest Budget.
	CreatedBy string `json:"createdBy"`
	// The modified timestamp in UTC of the Ingest Budget.
	ModifiedAt time.Time `json:"modifiedAt"`
	// The identifier of the user who modified the Ingest Budget.
	ModifiedBy string `json:"modifiedBy"`
	// The version of the Ingest Budget
	BudgetVersion *int32 `json:"budgetVersion,omitempty"`
}

// NewIngestBudgetV2 instantiates a new IngestBudgetV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngestBudgetV2(name string, scope string, capacityBytes int64, timezone string, resetTime string, action string, id string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string) *IngestBudgetV2 {
	this := IngestBudgetV2{}
	this.Name = name
	this.Scope = scope
	this.CapacityBytes = capacityBytes
	this.Timezone = timezone
	this.ResetTime = resetTime
	this.Action = action
	this.Id = id
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	return &this
}

// NewIngestBudgetV2WithDefaults instantiates a new IngestBudgetV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngestBudgetV2WithDefaults() *IngestBudgetV2 {
	this := IngestBudgetV2{}
	return &this
}

// GetName returns the Name field value
func (o *IngestBudgetV2) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IngestBudgetV2) SetName(v string) {
	o.Name = v
}

// GetScope returns the Scope field value
func (o *IngestBudgetV2) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetScopeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *IngestBudgetV2) SetScope(v string) {
	o.Scope = v
}

// GetCapacityBytes returns the CapacityBytes field value
func (o *IngestBudgetV2) GetCapacityBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CapacityBytes
}

// GetCapacityBytesOk returns a tuple with the CapacityBytes field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetCapacityBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CapacityBytes, true
}

// SetCapacityBytes sets field value
func (o *IngestBudgetV2) SetCapacityBytes(v int64) {
	o.CapacityBytes = v
}

// GetTimezone returns the Timezone field value
func (o *IngestBudgetV2) GetTimezone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetTimezoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timezone, true
}

// SetTimezone sets field value
func (o *IngestBudgetV2) SetTimezone(v string) {
	o.Timezone = v
}

// GetResetTime returns the ResetTime field value
func (o *IngestBudgetV2) GetResetTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResetTime
}

// GetResetTimeOk returns a tuple with the ResetTime field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetResetTimeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResetTime, true
}

// SetResetTime sets field value
func (o *IngestBudgetV2) SetResetTime(v string) {
	o.ResetTime = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IngestBudgetV2) SetDescription(v string) {
	o.Description = &v
}

// GetAction returns the Action field value
func (o *IngestBudgetV2) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetActionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *IngestBudgetV2) SetAction(v string) {
	o.Action = v
}

// GetAuditThreshold returns the AuditThreshold field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetAuditThreshold() int32 {
	if o == nil || o.AuditThreshold == nil {
		var ret int32
		return ret
	}
	return *o.AuditThreshold
}

// GetAuditThresholdOk returns a tuple with the AuditThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetAuditThresholdOk() (*int32, bool) {
	if o == nil || o.AuditThreshold == nil {
		return nil, false
	}
	return o.AuditThreshold, true
}

// HasAuditThreshold returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasAuditThreshold() bool {
	if o != nil && o.AuditThreshold != nil {
		return true
	}

	return false
}

// SetAuditThreshold gets a reference to the given int32 and assigns it to the AuditThreshold field.
func (o *IngestBudgetV2) SetAuditThreshold(v int32) {
	o.AuditThreshold = &v
}

// GetId returns the Id field value
func (o *IngestBudgetV2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *IngestBudgetV2) SetId(v string) {
	o.Id = v
}

// GetUsageBytes returns the UsageBytes field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetUsageBytes() int64 {
	if o == nil || o.UsageBytes == nil {
		var ret int64
		return ret
	}
	return *o.UsageBytes
}

// GetUsageBytesOk returns a tuple with the UsageBytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetUsageBytesOk() (*int64, bool) {
	if o == nil || o.UsageBytes == nil {
		return nil, false
	}
	return o.UsageBytes, true
}

// HasUsageBytes returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasUsageBytes() bool {
	if o != nil && o.UsageBytes != nil {
		return true
	}

	return false
}

// SetUsageBytes gets a reference to the given int64 and assigns it to the UsageBytes field.
func (o *IngestBudgetV2) SetUsageBytes(v int64) {
	o.UsageBytes = &v
}

// GetUsageStatus returns the UsageStatus field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetUsageStatus() string {
	if o == nil || o.UsageStatus == nil {
		var ret string
		return ret
	}
	return *o.UsageStatus
}

// GetUsageStatusOk returns a tuple with the UsageStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetUsageStatusOk() (*string, bool) {
	if o == nil || o.UsageStatus == nil {
		return nil, false
	}
	return o.UsageStatus, true
}

// HasUsageStatus returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasUsageStatus() bool {
	if o != nil && o.UsageStatus != nil {
		return true
	}

	return false
}

// SetUsageStatus gets a reference to the given string and assigns it to the UsageStatus field.
func (o *IngestBudgetV2) SetUsageStatus(v string) {
	o.UsageStatus = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *IngestBudgetV2) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IngestBudgetV2) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *IngestBudgetV2) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *IngestBudgetV2) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *IngestBudgetV2) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *IngestBudgetV2) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *IngestBudgetV2) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *IngestBudgetV2) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetBudgetVersion returns the BudgetVersion field value if set, zero value otherwise.
func (o *IngestBudgetV2) GetBudgetVersion() int32 {
	if o == nil || o.BudgetVersion == nil {
		var ret int32
		return ret
	}
	return *o.BudgetVersion
}

// GetBudgetVersionOk returns a tuple with the BudgetVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IngestBudgetV2) GetBudgetVersionOk() (*int32, bool) {
	if o == nil || o.BudgetVersion == nil {
		return nil, false
	}
	return o.BudgetVersion, true
}

// HasBudgetVersion returns a boolean if a field has been set.
func (o *IngestBudgetV2) HasBudgetVersion() bool {
	if o != nil && o.BudgetVersion != nil {
		return true
	}

	return false
}

// SetBudgetVersion gets a reference to the given int32 and assigns it to the BudgetVersion field.
func (o *IngestBudgetV2) SetBudgetVersion(v int32) {
	o.BudgetVersion = &v
}

func (o IngestBudgetV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["capacityBytes"] = o.CapacityBytes
	}
	if true {
		toSerialize["timezone"] = o.Timezone
	}
	if true {
		toSerialize["resetTime"] = o.ResetTime
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.AuditThreshold != nil {
		toSerialize["auditThreshold"] = o.AuditThreshold
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.UsageBytes != nil {
		toSerialize["usageBytes"] = o.UsageBytes
	}
	if o.UsageStatus != nil {
		toSerialize["usageStatus"] = o.UsageStatus
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if true {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if true {
		toSerialize["modifiedAt"] = o.ModifiedAt
	}
	if true {
		toSerialize["modifiedBy"] = o.ModifiedBy
	}
	if o.BudgetVersion != nil {
		toSerialize["budgetVersion"] = o.BudgetVersion
	}
	return json.Marshal(toSerialize)
}

type NullableIngestBudgetV2 struct {
	value *IngestBudgetV2
	isSet bool
}

func (v NullableIngestBudgetV2) Get() *IngestBudgetV2 {
	return v.value
}

func (v *NullableIngestBudgetV2) Set(val *IngestBudgetV2) {
	v.value = val
	v.isSet = true
}

func (v NullableIngestBudgetV2) IsSet() bool {
	return v.isSet
}

func (v *NullableIngestBudgetV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngestBudgetV2(val *IngestBudgetV2) *NullableIngestBudgetV2 {
	return &NullableIngestBudgetV2{value: val, isSet: true}
}

func (v NullableIngestBudgetV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngestBudgetV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



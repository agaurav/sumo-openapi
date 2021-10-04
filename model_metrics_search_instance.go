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

// MetricsSearchInstance struct for MetricsSearchInstance
type MetricsSearchInstance struct {
	// Item title in the content library.
	Title string `json:"title"`
	// Item description in the content library.
	Description string `json:"description"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// Log query used to add an overlay to the chart.
	LogQuery *string `json:"logQuery,omitempty"`
	// Metrics queries, up to the maximum of six.
	MetricsQueries []MetricsSearchQuery `json:"metricsQueries"`
	// Desired quantization in seconds.
	DesiredQuantizationInSecs *int32 `json:"desiredQuantizationInSecs,omitempty"`
	// Chart properties, like line width, color palette, and the fill missing data method. Leave this field empty to use the defaults. This property contains JSON object encoded as a string. 
	Properties *string `json:"properties,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Identifier of the metrics search.
	Id string `json:"id"`
	// Identifier of the parent element in the content library, such as folder.
	ParentId *string `json:"parentId,omitempty"`
}

// NewMetricsSearchInstance instantiates a new MetricsSearchInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsSearchInstance(title string, description string, timeRange ResolvableTimeRange, metricsQueries []MetricsSearchQuery, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *MetricsSearchInstance {
	this := MetricsSearchInstance{}
	this.Title = title
	this.Description = description
	this.TimeRange = timeRange
	this.MetricsQueries = metricsQueries
	var desiredQuantizationInSecs int32 = 0
	this.DesiredQuantizationInSecs = &desiredQuantizationInSecs
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	return &this
}

// NewMetricsSearchInstanceWithDefaults instantiates a new MetricsSearchInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsSearchInstanceWithDefaults() *MetricsSearchInstance {
	this := MetricsSearchInstance{}
	var desiredQuantizationInSecs int32 = 0
	this.DesiredQuantizationInSecs = &desiredQuantizationInSecs
	return &this
}

// GetTitle returns the Title field value
func (o *MetricsSearchInstance) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetTitleOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *MetricsSearchInstance) SetTitle(v string) {
	o.Title = v
}

// GetDescription returns the Description field value
func (o *MetricsSearchInstance) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *MetricsSearchInstance) SetDescription(v string) {
	o.Description = v
}

// GetTimeRange returns the TimeRange field value
func (o *MetricsSearchInstance) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *MetricsSearchInstance) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *MetricsSearchInstance) GetLogQuery() string {
	if o == nil || o.LogQuery == nil {
		var ret string
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetLogQueryOk() (*string, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *MetricsSearchInstance) HasLogQuery() bool {
	if o != nil && o.LogQuery != nil {
		return true
	}

	return false
}

// SetLogQuery gets a reference to the given string and assigns it to the LogQuery field.
func (o *MetricsSearchInstance) SetLogQuery(v string) {
	o.LogQuery = &v
}

// GetMetricsQueries returns the MetricsQueries field value
func (o *MetricsSearchInstance) GetMetricsQueries() []MetricsSearchQuery {
	if o == nil {
		var ret []MetricsSearchQuery
		return ret
	}

	return o.MetricsQueries
}

// GetMetricsQueriesOk returns a tuple with the MetricsQueries field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetMetricsQueriesOk() (*[]MetricsSearchQuery, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricsQueries, true
}

// SetMetricsQueries sets field value
func (o *MetricsSearchInstance) SetMetricsQueries(v []MetricsSearchQuery) {
	o.MetricsQueries = v
}

// GetDesiredQuantizationInSecs returns the DesiredQuantizationInSecs field value if set, zero value otherwise.
func (o *MetricsSearchInstance) GetDesiredQuantizationInSecs() int32 {
	if o == nil || o.DesiredQuantizationInSecs == nil {
		var ret int32
		return ret
	}
	return *o.DesiredQuantizationInSecs
}

// GetDesiredQuantizationInSecsOk returns a tuple with the DesiredQuantizationInSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetDesiredQuantizationInSecsOk() (*int32, bool) {
	if o == nil || o.DesiredQuantizationInSecs == nil {
		return nil, false
	}
	return o.DesiredQuantizationInSecs, true
}

// HasDesiredQuantizationInSecs returns a boolean if a field has been set.
func (o *MetricsSearchInstance) HasDesiredQuantizationInSecs() bool {
	if o != nil && o.DesiredQuantizationInSecs != nil {
		return true
	}

	return false
}

// SetDesiredQuantizationInSecs gets a reference to the given int32 and assigns it to the DesiredQuantizationInSecs field.
func (o *MetricsSearchInstance) SetDesiredQuantizationInSecs(v int32) {
	o.DesiredQuantizationInSecs = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *MetricsSearchInstance) GetProperties() string {
	if o == nil || o.Properties == nil {
		var ret string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetPropertiesOk() (*string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *MetricsSearchInstance) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given string and assigns it to the Properties field.
func (o *MetricsSearchInstance) SetProperties(v string) {
	o.Properties = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *MetricsSearchInstance) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *MetricsSearchInstance) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *MetricsSearchInstance) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *MetricsSearchInstance) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *MetricsSearchInstance) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *MetricsSearchInstance) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *MetricsSearchInstance) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *MetricsSearchInstance) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *MetricsSearchInstance) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MetricsSearchInstance) SetId(v string) {
	o.Id = v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *MetricsSearchInstance) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsSearchInstance) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *MetricsSearchInstance) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *MetricsSearchInstance) SetParentId(v string) {
	o.ParentId = &v
}

func (o MetricsSearchInstance) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["title"] = o.Title
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["timeRange"] = o.TimeRange
	}
	if o.LogQuery != nil {
		toSerialize["logQuery"] = o.LogQuery
	}
	if true {
		toSerialize["metricsQueries"] = o.MetricsQueries
	}
	if o.DesiredQuantizationInSecs != nil {
		toSerialize["desiredQuantizationInSecs"] = o.DesiredQuantizationInSecs
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
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
	if true {
		toSerialize["id"] = o.Id
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	return json.Marshal(toSerialize)
}

type NullableMetricsSearchInstance struct {
	value *MetricsSearchInstance
	isSet bool
}

func (v NullableMetricsSearchInstance) Get() *MetricsSearchInstance {
	return v.value
}

func (v *NullableMetricsSearchInstance) Set(val *MetricsSearchInstance) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsSearchInstance) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsSearchInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsSearchInstance(val *MetricsSearchInstance) *NullableMetricsSearchInstance {
	return &NullableMetricsSearchInstance{value: val, isSet: true}
}

func (v NullableMetricsSearchInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsSearchInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



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

// Partition struct for Partition
type Partition struct {
	// The name of the partition.
	Name string `json:"name"`
	// The query that defines the data to be included in the partition.
	RoutingExpression string `json:"routingExpression"`
	// The Data Tier where the data in the partition will reside. Possible values are:               1. `continuous`               2. `frequent`               3. `infrequent` Note: The \"infrequent\" and \"frequent\" tiers are only to available Cloud Flex Credits Enterprise Suite accounts.
	AnalyticsTier *string `json:"analyticsTier,omitempty"`
	// The number of days to retain data in the partition, or -1 to use the default value for your account.  Only relevant if your account has variable retention enabled.
	RetentionPeriod *int32 `json:"retentionPeriod,omitempty"`
	// Whether the partition is compliant or not. Mark a partition as compliant if it contains data used for compliance or audit purpose. Retention for a compliant partition can only be increased and cannot be reduced after the partition is marked compliant. A partition once marked compliant, cannot be marked non-compliant later.
	IsCompliant *bool `json:"isCompliant,omitempty"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier for the partition.
	Id string `json:"id"`
	// Size of data in partition in bytes.
	TotalBytes int64 `json:"totalBytes"`
	// This has the value `true` if the partition is active and `false` if it has been decommissioned.
	IsActive *bool `json:"isActive,omitempty"`
	// If the retentionPeriod is scheduled to be updated in the future (i.e., if retentionPeriod is previously reduced with value of reduceRetentionPeriodImmediately as false), this property gives the future value of retentionPeriod while retentionPeriod gives the current value. retentionPeriod will take up the value of newRetentionPeriod after the scheduled time.
	NewRetentionPeriod *int32 `json:"newRetentionPeriod,omitempty"`
	// This has the value `DefaultIndex`, `AuditIndex`or `Partition` depending upon the type of partition.
	IndexType *string `json:"indexType,omitempty"`
	// When the newRetentionPeriod will become effective in UTC format.
	RetentionEffectiveAt *time.Time `json:"retentionEffectiveAt,omitempty"`
	// Id of the data forwarding configuration to be used by the partition.
	DataForwardingId *string `json:"dataForwardingId,omitempty"`
}

// NewPartition instantiates a new Partition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPartition(name string, routingExpression string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, totalBytes int64) *Partition {
	this := Partition{}
	this.Name = name
	this.RoutingExpression = routingExpression
	var analyticsTier string = "continuous"
	this.AnalyticsTier = &analyticsTier
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var isCompliant bool = false
	this.IsCompliant = &isCompliant
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	this.TotalBytes = totalBytes
	return &this
}

// NewPartitionWithDefaults instantiates a new Partition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPartitionWithDefaults() *Partition {
	this := Partition{}
	var analyticsTier string = "continuous"
	this.AnalyticsTier = &analyticsTier
	var retentionPeriod int32 = -1
	this.RetentionPeriod = &retentionPeriod
	var isCompliant bool = false
	this.IsCompliant = &isCompliant
	return &this
}

// GetName returns the Name field value
func (o *Partition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Partition) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Partition) SetName(v string) {
	o.Name = v
}

// GetRoutingExpression returns the RoutingExpression field value
func (o *Partition) GetRoutingExpression() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoutingExpression
}

// GetRoutingExpressionOk returns a tuple with the RoutingExpression field value
// and a boolean to check if the value has been set.
func (o *Partition) GetRoutingExpressionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RoutingExpression, true
}

// SetRoutingExpression sets field value
func (o *Partition) SetRoutingExpression(v string) {
	o.RoutingExpression = v
}

// GetAnalyticsTier returns the AnalyticsTier field value if set, zero value otherwise.
func (o *Partition) GetAnalyticsTier() string {
	if o == nil || o.AnalyticsTier == nil {
		var ret string
		return ret
	}
	return *o.AnalyticsTier
}

// GetAnalyticsTierOk returns a tuple with the AnalyticsTier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetAnalyticsTierOk() (*string, bool) {
	if o == nil || o.AnalyticsTier == nil {
		return nil, false
	}
	return o.AnalyticsTier, true
}

// HasAnalyticsTier returns a boolean if a field has been set.
func (o *Partition) HasAnalyticsTier() bool {
	if o != nil && o.AnalyticsTier != nil {
		return true
	}

	return false
}

// SetAnalyticsTier gets a reference to the given string and assigns it to the AnalyticsTier field.
func (o *Partition) SetAnalyticsTier(v string) {
	o.AnalyticsTier = &v
}

// GetRetentionPeriod returns the RetentionPeriod field value if set, zero value otherwise.
func (o *Partition) GetRetentionPeriod() int32 {
	if o == nil || o.RetentionPeriod == nil {
		var ret int32
		return ret
	}
	return *o.RetentionPeriod
}

// GetRetentionPeriodOk returns a tuple with the RetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetRetentionPeriodOk() (*int32, bool) {
	if o == nil || o.RetentionPeriod == nil {
		return nil, false
	}
	return o.RetentionPeriod, true
}

// HasRetentionPeriod returns a boolean if a field has been set.
func (o *Partition) HasRetentionPeriod() bool {
	if o != nil && o.RetentionPeriod != nil {
		return true
	}

	return false
}

// SetRetentionPeriod gets a reference to the given int32 and assigns it to the RetentionPeriod field.
func (o *Partition) SetRetentionPeriod(v int32) {
	o.RetentionPeriod = &v
}

// GetIsCompliant returns the IsCompliant field value if set, zero value otherwise.
func (o *Partition) GetIsCompliant() bool {
	if o == nil || o.IsCompliant == nil {
		var ret bool
		return ret
	}
	return *o.IsCompliant
}

// GetIsCompliantOk returns a tuple with the IsCompliant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetIsCompliantOk() (*bool, bool) {
	if o == nil || o.IsCompliant == nil {
		return nil, false
	}
	return o.IsCompliant, true
}

// HasIsCompliant returns a boolean if a field has been set.
func (o *Partition) HasIsCompliant() bool {
	if o != nil && o.IsCompliant != nil {
		return true
	}

	return false
}

// SetIsCompliant gets a reference to the given bool and assigns it to the IsCompliant field.
func (o *Partition) SetIsCompliant(v bool) {
	o.IsCompliant = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Partition) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Partition) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Partition) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Partition) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Partition) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Partition) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *Partition) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *Partition) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *Partition) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *Partition) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *Partition) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *Partition) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *Partition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Partition) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Partition) SetId(v string) {
	o.Id = v
}

// GetTotalBytes returns the TotalBytes field value
func (o *Partition) GetTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *Partition) GetTotalBytesOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value
func (o *Partition) SetTotalBytes(v int64) {
	o.TotalBytes = v
}

// GetIsActive returns the IsActive field value if set, zero value otherwise.
func (o *Partition) GetIsActive() bool {
	if o == nil || o.IsActive == nil {
		var ret bool
		return ret
	}
	return *o.IsActive
}

// GetIsActiveOk returns a tuple with the IsActive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetIsActiveOk() (*bool, bool) {
	if o == nil || o.IsActive == nil {
		return nil, false
	}
	return o.IsActive, true
}

// HasIsActive returns a boolean if a field has been set.
func (o *Partition) HasIsActive() bool {
	if o != nil && o.IsActive != nil {
		return true
	}

	return false
}

// SetIsActive gets a reference to the given bool and assigns it to the IsActive field.
func (o *Partition) SetIsActive(v bool) {
	o.IsActive = &v
}

// GetNewRetentionPeriod returns the NewRetentionPeriod field value if set, zero value otherwise.
func (o *Partition) GetNewRetentionPeriod() int32 {
	if o == nil || o.NewRetentionPeriod == nil {
		var ret int32
		return ret
	}
	return *o.NewRetentionPeriod
}

// GetNewRetentionPeriodOk returns a tuple with the NewRetentionPeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetNewRetentionPeriodOk() (*int32, bool) {
	if o == nil || o.NewRetentionPeriod == nil {
		return nil, false
	}
	return o.NewRetentionPeriod, true
}

// HasNewRetentionPeriod returns a boolean if a field has been set.
func (o *Partition) HasNewRetentionPeriod() bool {
	if o != nil && o.NewRetentionPeriod != nil {
		return true
	}

	return false
}

// SetNewRetentionPeriod gets a reference to the given int32 and assigns it to the NewRetentionPeriod field.
func (o *Partition) SetNewRetentionPeriod(v int32) {
	o.NewRetentionPeriod = &v
}

// GetIndexType returns the IndexType field value if set, zero value otherwise.
func (o *Partition) GetIndexType() string {
	if o == nil || o.IndexType == nil {
		var ret string
		return ret
	}
	return *o.IndexType
}

// GetIndexTypeOk returns a tuple with the IndexType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetIndexTypeOk() (*string, bool) {
	if o == nil || o.IndexType == nil {
		return nil, false
	}
	return o.IndexType, true
}

// HasIndexType returns a boolean if a field has been set.
func (o *Partition) HasIndexType() bool {
	if o != nil && o.IndexType != nil {
		return true
	}

	return false
}

// SetIndexType gets a reference to the given string and assigns it to the IndexType field.
func (o *Partition) SetIndexType(v string) {
	o.IndexType = &v
}

// GetRetentionEffectiveAt returns the RetentionEffectiveAt field value if set, zero value otherwise.
func (o *Partition) GetRetentionEffectiveAt() time.Time {
	if o == nil || o.RetentionEffectiveAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RetentionEffectiveAt
}

// GetRetentionEffectiveAtOk returns a tuple with the RetentionEffectiveAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetRetentionEffectiveAtOk() (*time.Time, bool) {
	if o == nil || o.RetentionEffectiveAt == nil {
		return nil, false
	}
	return o.RetentionEffectiveAt, true
}

// HasRetentionEffectiveAt returns a boolean if a field has been set.
func (o *Partition) HasRetentionEffectiveAt() bool {
	if o != nil && o.RetentionEffectiveAt != nil {
		return true
	}

	return false
}

// SetRetentionEffectiveAt gets a reference to the given time.Time and assigns it to the RetentionEffectiveAt field.
func (o *Partition) SetRetentionEffectiveAt(v time.Time) {
	o.RetentionEffectiveAt = &v
}

// GetDataForwardingId returns the DataForwardingId field value if set, zero value otherwise.
func (o *Partition) GetDataForwardingId() string {
	if o == nil || o.DataForwardingId == nil {
		var ret string
		return ret
	}
	return *o.DataForwardingId
}

// GetDataForwardingIdOk returns a tuple with the DataForwardingId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Partition) GetDataForwardingIdOk() (*string, bool) {
	if o == nil || o.DataForwardingId == nil {
		return nil, false
	}
	return o.DataForwardingId, true
}

// HasDataForwardingId returns a boolean if a field has been set.
func (o *Partition) HasDataForwardingId() bool {
	if o != nil && o.DataForwardingId != nil {
		return true
	}

	return false
}

// SetDataForwardingId gets a reference to the given string and assigns it to the DataForwardingId field.
func (o *Partition) SetDataForwardingId(v string) {
	o.DataForwardingId = &v
}

func (o Partition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["routingExpression"] = o.RoutingExpression
	}
	if o.AnalyticsTier != nil {
		toSerialize["analyticsTier"] = o.AnalyticsTier
	}
	if o.RetentionPeriod != nil {
		toSerialize["retentionPeriod"] = o.RetentionPeriod
	}
	if o.IsCompliant != nil {
		toSerialize["isCompliant"] = o.IsCompliant
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
	if true {
		toSerialize["totalBytes"] = o.TotalBytes
	}
	if o.IsActive != nil {
		toSerialize["isActive"] = o.IsActive
	}
	if o.NewRetentionPeriod != nil {
		toSerialize["newRetentionPeriod"] = o.NewRetentionPeriod
	}
	if o.IndexType != nil {
		toSerialize["indexType"] = o.IndexType
	}
	if o.RetentionEffectiveAt != nil {
		toSerialize["retentionEffectiveAt"] = o.RetentionEffectiveAt
	}
	if o.DataForwardingId != nil {
		toSerialize["dataForwardingId"] = o.DataForwardingId
	}
	return json.Marshal(toSerialize)
}

type NullablePartition struct {
	value *Partition
	isSet bool
}

func (v NullablePartition) Get() *Partition {
	return v.value
}

func (v *NullablePartition) Set(val *Partition) {
	v.value = val
	v.isSet = true
}

func (v NullablePartition) IsSet() bool {
	return v.isSet
}

func (v *NullablePartition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartition(val *Partition) *NullablePartition {
	return &NullablePartition{value: val, isSet: true}
}

func (v NullablePartition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



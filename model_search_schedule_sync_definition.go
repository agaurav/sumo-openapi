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

// SearchScheduleSyncDefinition struct for SearchScheduleSyncDefinition
type SearchScheduleSyncDefinition struct {
	// Cron-like expression specifying the search's schedule. Field scheduleType must be set to \"Custom\", otherwise, scheduleType takes precedence over cronExpression.
	CronExpression *string `json:"cronExpression,omitempty"`
	// A human-friendly text describing the query time range. For e.g. \"-2h\", \"last three days\", \"team default time\"
	DisplayableTimeRange *string `json:"displayableTimeRange,omitempty"`
	ParseableTimeRange ResolvableTimeRange `json:"parseableTimeRange"`
	// Time zone identifier for time specification. Either an abbreviation such as \"PST\", a full name such as \"America/Los_Angeles\", or a custom ID such as \"GMT-8:00\". Note that the support of abbreviations is for JDK 1.1.x compatibility only and full names should be used.
	TimeZone string `json:"timeZone"`
	Threshold *NotificationThresholdSyncDefinition `json:"threshold,omitempty"`
	Notification ScheduleNotificationSyncDefinition `json:"notification"`
	// Run schedule of the scheduled search. Set to \"Custom\" to specify the schedule with a CRON expression. Possible schedule types are:   - `RealTime`   - `15Minutes`   - `1Hour`   - `2Hours`   - `4Hours`   - `6Hours`   - `8Hours`   - `12Hours`   - `1Day`   - `1Week`   - `Custom`
	ScheduleType string `json:"scheduleType"`
	// If enabled, emails are not sent out in case of errors with the search.
	MuteErrorEmails *bool `json:"muteErrorEmails,omitempty"`
	// A list of scheduled search parameters.
	Parameters []ScheduleSearchParameterSyncDefinition `json:"parameters"`
}

// NewSearchScheduleSyncDefinition instantiates a new SearchScheduleSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchScheduleSyncDefinition(parseableTimeRange ResolvableTimeRange, timeZone string, notification ScheduleNotificationSyncDefinition, scheduleType string, parameters []ScheduleSearchParameterSyncDefinition) *SearchScheduleSyncDefinition {
	this := SearchScheduleSyncDefinition{}
	this.ParseableTimeRange = parseableTimeRange
	this.TimeZone = timeZone
	this.Notification = notification
	this.ScheduleType = scheduleType
	this.Parameters = parameters
	return &this
}

// NewSearchScheduleSyncDefinitionWithDefaults instantiates a new SearchScheduleSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchScheduleSyncDefinitionWithDefaults() *SearchScheduleSyncDefinition {
	this := SearchScheduleSyncDefinition{}
	return &this
}

// GetCronExpression returns the CronExpression field value if set, zero value otherwise.
func (o *SearchScheduleSyncDefinition) GetCronExpression() string {
	if o == nil || o.CronExpression == nil {
		var ret string
		return ret
	}
	return *o.CronExpression
}

// GetCronExpressionOk returns a tuple with the CronExpression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetCronExpressionOk() (*string, bool) {
	if o == nil || o.CronExpression == nil {
		return nil, false
	}
	return o.CronExpression, true
}

// HasCronExpression returns a boolean if a field has been set.
func (o *SearchScheduleSyncDefinition) HasCronExpression() bool {
	if o != nil && o.CronExpression != nil {
		return true
	}

	return false
}

// SetCronExpression gets a reference to the given string and assigns it to the CronExpression field.
func (o *SearchScheduleSyncDefinition) SetCronExpression(v string) {
	o.CronExpression = &v
}

// GetDisplayableTimeRange returns the DisplayableTimeRange field value if set, zero value otherwise.
func (o *SearchScheduleSyncDefinition) GetDisplayableTimeRange() string {
	if o == nil || o.DisplayableTimeRange == nil {
		var ret string
		return ret
	}
	return *o.DisplayableTimeRange
}

// GetDisplayableTimeRangeOk returns a tuple with the DisplayableTimeRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetDisplayableTimeRangeOk() (*string, bool) {
	if o == nil || o.DisplayableTimeRange == nil {
		return nil, false
	}
	return o.DisplayableTimeRange, true
}

// HasDisplayableTimeRange returns a boolean if a field has been set.
func (o *SearchScheduleSyncDefinition) HasDisplayableTimeRange() bool {
	if o != nil && o.DisplayableTimeRange != nil {
		return true
	}

	return false
}

// SetDisplayableTimeRange gets a reference to the given string and assigns it to the DisplayableTimeRange field.
func (o *SearchScheduleSyncDefinition) SetDisplayableTimeRange(v string) {
	o.DisplayableTimeRange = &v
}

// GetParseableTimeRange returns the ParseableTimeRange field value
func (o *SearchScheduleSyncDefinition) GetParseableTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.ParseableTimeRange
}

// GetParseableTimeRangeOk returns a tuple with the ParseableTimeRange field value
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetParseableTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ParseableTimeRange, true
}

// SetParseableTimeRange sets field value
func (o *SearchScheduleSyncDefinition) SetParseableTimeRange(v ResolvableTimeRange) {
	o.ParseableTimeRange = v
}

// GetTimeZone returns the TimeZone field value
func (o *SearchScheduleSyncDefinition) GetTimeZone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetTimeZoneOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeZone, true
}

// SetTimeZone sets field value
func (o *SearchScheduleSyncDefinition) SetTimeZone(v string) {
	o.TimeZone = v
}

// GetThreshold returns the Threshold field value if set, zero value otherwise.
func (o *SearchScheduleSyncDefinition) GetThreshold() NotificationThresholdSyncDefinition {
	if o == nil || o.Threshold == nil {
		var ret NotificationThresholdSyncDefinition
		return ret
	}
	return *o.Threshold
}

// GetThresholdOk returns a tuple with the Threshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetThresholdOk() (*NotificationThresholdSyncDefinition, bool) {
	if o == nil || o.Threshold == nil {
		return nil, false
	}
	return o.Threshold, true
}

// HasThreshold returns a boolean if a field has been set.
func (o *SearchScheduleSyncDefinition) HasThreshold() bool {
	if o != nil && o.Threshold != nil {
		return true
	}

	return false
}

// SetThreshold gets a reference to the given NotificationThresholdSyncDefinition and assigns it to the Threshold field.
func (o *SearchScheduleSyncDefinition) SetThreshold(v NotificationThresholdSyncDefinition) {
	o.Threshold = &v
}

// GetNotification returns the Notification field value
func (o *SearchScheduleSyncDefinition) GetNotification() ScheduleNotificationSyncDefinition {
	if o == nil {
		var ret ScheduleNotificationSyncDefinition
		return ret
	}

	return o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetNotificationOk() (*ScheduleNotificationSyncDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Notification, true
}

// SetNotification sets field value
func (o *SearchScheduleSyncDefinition) SetNotification(v ScheduleNotificationSyncDefinition) {
	o.Notification = v
}

// GetScheduleType returns the ScheduleType field value
func (o *SearchScheduleSyncDefinition) GetScheduleType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScheduleType
}

// GetScheduleTypeOk returns a tuple with the ScheduleType field value
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetScheduleTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ScheduleType, true
}

// SetScheduleType sets field value
func (o *SearchScheduleSyncDefinition) SetScheduleType(v string) {
	o.ScheduleType = v
}

// GetMuteErrorEmails returns the MuteErrorEmails field value if set, zero value otherwise.
func (o *SearchScheduleSyncDefinition) GetMuteErrorEmails() bool {
	if o == nil || o.MuteErrorEmails == nil {
		var ret bool
		return ret
	}
	return *o.MuteErrorEmails
}

// GetMuteErrorEmailsOk returns a tuple with the MuteErrorEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetMuteErrorEmailsOk() (*bool, bool) {
	if o == nil || o.MuteErrorEmails == nil {
		return nil, false
	}
	return o.MuteErrorEmails, true
}

// HasMuteErrorEmails returns a boolean if a field has been set.
func (o *SearchScheduleSyncDefinition) HasMuteErrorEmails() bool {
	if o != nil && o.MuteErrorEmails != nil {
		return true
	}

	return false
}

// SetMuteErrorEmails gets a reference to the given bool and assigns it to the MuteErrorEmails field.
func (o *SearchScheduleSyncDefinition) SetMuteErrorEmails(v bool) {
	o.MuteErrorEmails = &v
}

// GetParameters returns the Parameters field value
func (o *SearchScheduleSyncDefinition) GetParameters() []ScheduleSearchParameterSyncDefinition {
	if o == nil {
		var ret []ScheduleSearchParameterSyncDefinition
		return ret
	}

	return o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value
// and a boolean to check if the value has been set.
func (o *SearchScheduleSyncDefinition) GetParametersOk() (*[]ScheduleSearchParameterSyncDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Parameters, true
}

// SetParameters sets field value
func (o *SearchScheduleSyncDefinition) SetParameters(v []ScheduleSearchParameterSyncDefinition) {
	o.Parameters = v
}

func (o SearchScheduleSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CronExpression != nil {
		toSerialize["cronExpression"] = o.CronExpression
	}
	if o.DisplayableTimeRange != nil {
		toSerialize["displayableTimeRange"] = o.DisplayableTimeRange
	}
	if true {
		toSerialize["parseableTimeRange"] = o.ParseableTimeRange
	}
	if true {
		toSerialize["timeZone"] = o.TimeZone
	}
	if o.Threshold != nil {
		toSerialize["threshold"] = o.Threshold
	}
	if true {
		toSerialize["notification"] = o.Notification
	}
	if true {
		toSerialize["scheduleType"] = o.ScheduleType
	}
	if o.MuteErrorEmails != nil {
		toSerialize["muteErrorEmails"] = o.MuteErrorEmails
	}
	if true {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableSearchScheduleSyncDefinition struct {
	value *SearchScheduleSyncDefinition
	isSet bool
}

func (v NullableSearchScheduleSyncDefinition) Get() *SearchScheduleSyncDefinition {
	return v.value
}

func (v *NullableSearchScheduleSyncDefinition) Set(val *SearchScheduleSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchScheduleSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchScheduleSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchScheduleSyncDefinition(val *SearchScheduleSyncDefinition) *NullableSearchScheduleSyncDefinition {
	return &NullableSearchScheduleSyncDefinition{value: val, isSet: true}
}

func (v NullableSearchScheduleSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchScheduleSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



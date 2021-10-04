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

// ReportPanelSyncDefinition struct for ReportPanelSyncDefinition
type ReportPanelSyncDefinition struct {
	// The title of the panel.
	Name string `json:"name"`
	// Type of [area chart](https://help.sumologic.com/Dashboards-and-Alerts/Dashboards/Chart-Panel-Types). Supported values are:   1. `table` for Table   2. `bar` for Bar Chart   3. `column` for Column Chart   4. `line` for Line Chart   5. `area` for Area Chart   6. `pie` for Pie Chart   7. `svv` for Single Value Viewer   8. `title` for Title Panel   9. `text` for Text Panel  Values 1-7 are used for Data Panels.
	ViewerType string `json:"viewerType"`
	// Supported values are:   - `1` for small   - `2` for medium   - `3` for large
	DetailLevel int32 `json:"detailLevel"`
	// The query to run, for panels associated to log searches.
	QueryString string `json:"queryString"`
	// The query or queries to run, for panels associated to metrics searches.
	MetricsQueries []MetricsQuerySyncDefinition `json:"metricsQueries"`
	TimeRange ResolvableTimeRange `json:"timeRange"`
	// The horizontal position of the panel. A sumo screen is divided into 24 columns. The value for x can be any integer from 0 to 24.
	X int32 `json:"x"`
	// The vertical position of the panel. A sumo screen is divided into 24 rows. The value for y can be any integer from 0 to 24.
	Y int32 `json:"y"`
	// The width of the panel.
	Width int32 `json:"width"`
	// The height of the panel.
	Height int32 `json:"height"`
	// Visual settings for the panel.
	Properties string `json:"properties"`
	// A string identifier that you can use to refer to the panel in filters.panelIds.
	Id string `json:"id"`
	// The quantization interval aligns your time series data to common intervals on the time axis (for example every one minute) to optimize the visualization and performance.
	DesiredQuantizationInSecs *int32 `json:"desiredQuantizationInSecs,omitempty"`
	// The parameters for parameterized searches.
	QueryParameters []QueryParameterSyncDefinition `json:"queryParameters"`
	AutoParsingInfo *ReportAutoParsingInfo `json:"autoParsingInfo,omitempty"`
}

// NewReportPanelSyncDefinition instantiates a new ReportPanelSyncDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReportPanelSyncDefinition(name string, viewerType string, detailLevel int32, queryString string, metricsQueries []MetricsQuerySyncDefinition, timeRange ResolvableTimeRange, x int32, y int32, width int32, height int32, properties string, id string, queryParameters []QueryParameterSyncDefinition) *ReportPanelSyncDefinition {
	this := ReportPanelSyncDefinition{}
	this.Name = name
	this.ViewerType = viewerType
	this.DetailLevel = detailLevel
	this.QueryString = queryString
	this.MetricsQueries = metricsQueries
	this.TimeRange = timeRange
	this.X = x
	this.Y = y
	this.Width = width
	this.Height = height
	this.Properties = properties
	this.Id = id
	this.QueryParameters = queryParameters
	return &this
}

// NewReportPanelSyncDefinitionWithDefaults instantiates a new ReportPanelSyncDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReportPanelSyncDefinitionWithDefaults() *ReportPanelSyncDefinition {
	this := ReportPanelSyncDefinition{}
	return &this
}

// GetName returns the Name field value
func (o *ReportPanelSyncDefinition) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReportPanelSyncDefinition) SetName(v string) {
	o.Name = v
}

// GetViewerType returns the ViewerType field value
func (o *ReportPanelSyncDefinition) GetViewerType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ViewerType
}

// GetViewerTypeOk returns a tuple with the ViewerType field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetViewerTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ViewerType, true
}

// SetViewerType sets field value
func (o *ReportPanelSyncDefinition) SetViewerType(v string) {
	o.ViewerType = v
}

// GetDetailLevel returns the DetailLevel field value
func (o *ReportPanelSyncDefinition) GetDetailLevel() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DetailLevel
}

// GetDetailLevelOk returns a tuple with the DetailLevel field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetDetailLevelOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DetailLevel, true
}

// SetDetailLevel sets field value
func (o *ReportPanelSyncDefinition) SetDetailLevel(v int32) {
	o.DetailLevel = v
}

// GetQueryString returns the QueryString field value
func (o *ReportPanelSyncDefinition) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetQueryStringOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *ReportPanelSyncDefinition) SetQueryString(v string) {
	o.QueryString = v
}

// GetMetricsQueries returns the MetricsQueries field value
func (o *ReportPanelSyncDefinition) GetMetricsQueries() []MetricsQuerySyncDefinition {
	if o == nil {
		var ret []MetricsQuerySyncDefinition
		return ret
	}

	return o.MetricsQueries
}

// GetMetricsQueriesOk returns a tuple with the MetricsQueries field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetMetricsQueriesOk() (*[]MetricsQuerySyncDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MetricsQueries, true
}

// SetMetricsQueries sets field value
func (o *ReportPanelSyncDefinition) SetMetricsQueries(v []MetricsQuerySyncDefinition) {
	o.MetricsQueries = v
}

// GetTimeRange returns the TimeRange field value
func (o *ReportPanelSyncDefinition) GetTimeRange() ResolvableTimeRange {
	if o == nil {
		var ret ResolvableTimeRange
		return ret
	}

	return o.TimeRange
}

// GetTimeRangeOk returns a tuple with the TimeRange field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetTimeRangeOk() (*ResolvableTimeRange, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeRange, true
}

// SetTimeRange sets field value
func (o *ReportPanelSyncDefinition) SetTimeRange(v ResolvableTimeRange) {
	o.TimeRange = v
}

// GetX returns the X field value
func (o *ReportPanelSyncDefinition) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetXOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ReportPanelSyncDefinition) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *ReportPanelSyncDefinition) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetYOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ReportPanelSyncDefinition) SetY(v int32) {
	o.Y = v
}

// GetWidth returns the Width field value
func (o *ReportPanelSyncDefinition) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetWidthOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *ReportPanelSyncDefinition) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *ReportPanelSyncDefinition) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetHeightOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *ReportPanelSyncDefinition) SetHeight(v int32) {
	o.Height = v
}

// GetProperties returns the Properties field value
func (o *ReportPanelSyncDefinition) GetProperties() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetPropertiesOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *ReportPanelSyncDefinition) SetProperties(v string) {
	o.Properties = v
}

// GetId returns the Id field value
func (o *ReportPanelSyncDefinition) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReportPanelSyncDefinition) SetId(v string) {
	o.Id = v
}

// GetDesiredQuantizationInSecs returns the DesiredQuantizationInSecs field value if set, zero value otherwise.
func (o *ReportPanelSyncDefinition) GetDesiredQuantizationInSecs() int32 {
	if o == nil || o.DesiredQuantizationInSecs == nil {
		var ret int32
		return ret
	}
	return *o.DesiredQuantizationInSecs
}

// GetDesiredQuantizationInSecsOk returns a tuple with the DesiredQuantizationInSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetDesiredQuantizationInSecsOk() (*int32, bool) {
	if o == nil || o.DesiredQuantizationInSecs == nil {
		return nil, false
	}
	return o.DesiredQuantizationInSecs, true
}

// HasDesiredQuantizationInSecs returns a boolean if a field has been set.
func (o *ReportPanelSyncDefinition) HasDesiredQuantizationInSecs() bool {
	if o != nil && o.DesiredQuantizationInSecs != nil {
		return true
	}

	return false
}

// SetDesiredQuantizationInSecs gets a reference to the given int32 and assigns it to the DesiredQuantizationInSecs field.
func (o *ReportPanelSyncDefinition) SetDesiredQuantizationInSecs(v int32) {
	o.DesiredQuantizationInSecs = &v
}

// GetQueryParameters returns the QueryParameters field value
func (o *ReportPanelSyncDefinition) GetQueryParameters() []QueryParameterSyncDefinition {
	if o == nil {
		var ret []QueryParameterSyncDefinition
		return ret
	}

	return o.QueryParameters
}

// GetQueryParametersOk returns a tuple with the QueryParameters field value
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetQueryParametersOk() (*[]QueryParameterSyncDefinition, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.QueryParameters, true
}

// SetQueryParameters sets field value
func (o *ReportPanelSyncDefinition) SetQueryParameters(v []QueryParameterSyncDefinition) {
	o.QueryParameters = v
}

// GetAutoParsingInfo returns the AutoParsingInfo field value if set, zero value otherwise.
func (o *ReportPanelSyncDefinition) GetAutoParsingInfo() ReportAutoParsingInfo {
	if o == nil || o.AutoParsingInfo == nil {
		var ret ReportAutoParsingInfo
		return ret
	}
	return *o.AutoParsingInfo
}

// GetAutoParsingInfoOk returns a tuple with the AutoParsingInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReportPanelSyncDefinition) GetAutoParsingInfoOk() (*ReportAutoParsingInfo, bool) {
	if o == nil || o.AutoParsingInfo == nil {
		return nil, false
	}
	return o.AutoParsingInfo, true
}

// HasAutoParsingInfo returns a boolean if a field has been set.
func (o *ReportPanelSyncDefinition) HasAutoParsingInfo() bool {
	if o != nil && o.AutoParsingInfo != nil {
		return true
	}

	return false
}

// SetAutoParsingInfo gets a reference to the given ReportAutoParsingInfo and assigns it to the AutoParsingInfo field.
func (o *ReportPanelSyncDefinition) SetAutoParsingInfo(v ReportAutoParsingInfo) {
	o.AutoParsingInfo = &v
}

func (o ReportPanelSyncDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["viewerType"] = o.ViewerType
	}
	if true {
		toSerialize["detailLevel"] = o.DetailLevel
	}
	if true {
		toSerialize["queryString"] = o.QueryString
	}
	if true {
		toSerialize["metricsQueries"] = o.MetricsQueries
	}
	if true {
		toSerialize["timeRange"] = o.TimeRange
	}
	if true {
		toSerialize["x"] = o.X
	}
	if true {
		toSerialize["y"] = o.Y
	}
	if true {
		toSerialize["width"] = o.Width
	}
	if true {
		toSerialize["height"] = o.Height
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.DesiredQuantizationInSecs != nil {
		toSerialize["desiredQuantizationInSecs"] = o.DesiredQuantizationInSecs
	}
	if true {
		toSerialize["queryParameters"] = o.QueryParameters
	}
	if o.AutoParsingInfo != nil {
		toSerialize["autoParsingInfo"] = o.AutoParsingInfo
	}
	return json.Marshal(toSerialize)
}

type NullableReportPanelSyncDefinition struct {
	value *ReportPanelSyncDefinition
	isSet bool
}

func (v NullableReportPanelSyncDefinition) Get() *ReportPanelSyncDefinition {
	return v.value
}

func (v *NullableReportPanelSyncDefinition) Set(val *ReportPanelSyncDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableReportPanelSyncDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableReportPanelSyncDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportPanelSyncDefinition(val *ReportPanelSyncDefinition) *NullableReportPanelSyncDefinition {
	return &NullableReportPanelSyncDefinition{value: val, isSet: true}
}

func (v NullableReportPanelSyncDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportPanelSyncDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



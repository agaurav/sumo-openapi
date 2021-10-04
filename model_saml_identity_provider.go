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

// SamlIdentityProvider struct for SamlIdentityProvider
type SamlIdentityProvider struct {
	// The identifier used to generate a unique URL for user login.
	SpInitiatedLoginPath *string `json:"spInitiatedLoginPath,omitempty"`
	// Name of the SSO policy or another name used to describe the policy internally.
	ConfigurationName string `json:"configurationName"`
	// The unique URL assigned to the organization by the SAML Identity Provider.
	Issuer string `json:"issuer"`
	// True if Sumo Logic redirects users to your identity provider with a SAML AuthnRequest when signing in.
	SpInitiatedLoginEnabled *bool `json:"spInitiatedLoginEnabled,omitempty"`
	// The URL that the identity provider has assigned for Sumo Logic to submit SAML authentication requests to the identity provider.
	AuthnRequestUrl *string `json:"authnRequestUrl,omitempty"`
	// The certificate is used to verify the signature in SAML assertions.
	X509cert1 string `json:"x509cert1"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires.
	X509cert2 *string `json:"x509cert2,omitempty"`
	// The backup certificate used to verify the signature in SAML assertions when x509cert1 expires and x509cert2 is empty.
	X509cert3 *string `json:"x509cert3,omitempty"`
	OnDemandProvisioningEnabled *OnDemandProvisioningInfo `json:"onDemandProvisioningEnabled,omitempty"`
	// The role that Sumo Logic will assign to users when they sign in.
	RolesAttribute *string `json:"rolesAttribute,omitempty"`
	// True if users are redirected to a URL after signing out of Sumo Logic.
	LogoutEnabled *bool `json:"logoutEnabled,omitempty"`
	// The URL that users will be redirected to after signing out of Sumo Logic.
	LogoutUrl *string `json:"logoutUrl,omitempty"`
	// The email address of the new user account.
	EmailAttribute *string `json:"emailAttribute,omitempty"`
	// True if additional details are included when a user fails to sign in.
	DebugMode *bool `json:"debugMode,omitempty"`
	// True if Sumo Logic will send signed Authn requests to the identity provider.
	SignAuthnRequest *bool `json:"signAuthnRequest,omitempty"`
	// True if Sumo Logic will include the RequestedAuthnContext element of the SAML AuthnRequests it sends to the identity provider.
	DisableRequestedAuthnContext *bool `json:"disableRequestedAuthnContext,omitempty"`
	// True if the SAML binding is of HTTP Redirect type.
	IsRedirectBinding *bool `json:"isRedirectBinding,omitempty"`
	// Authentication Request Signing Certificate for the user.
	Certificate string `json:"certificate"`
	// Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format.
	CreatedAt time.Time `json:"createdAt"`
	// Identifier of the user who created the resource.
	CreatedBy string `json:"createdBy"`
	// Last modification timestamp in UTC.
	ModifiedAt time.Time `json:"modifiedAt"`
	// Identifier of the user who last modified the resource.
	ModifiedBy string `json:"modifiedBy"`
	// Unique identifier of the SAML Identity Provider.
	Id string `json:"id"`
	// The URL on Sumo Logic where the IdP will redirect to with its authentication response.
	AssertionConsumerUrl *string `json:"assertionConsumerUrl,omitempty"`
	// A unique identifier that is the intended audience of the SAML assertion.
	EntityId *string `json:"entityId,omitempty"`
}

// NewSamlIdentityProvider instantiates a new SamlIdentityProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSamlIdentityProvider(configurationName string, issuer string, x509cert1 string, certificate string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string) *SamlIdentityProvider {
	this := SamlIdentityProvider{}
	var spInitiatedLoginPath string = ""
	this.SpInitiatedLoginPath = &spInitiatedLoginPath
	this.ConfigurationName = configurationName
	this.Issuer = issuer
	var spInitiatedLoginEnabled bool = false
	this.SpInitiatedLoginEnabled = &spInitiatedLoginEnabled
	var authnRequestUrl string = ""
	this.AuthnRequestUrl = &authnRequestUrl
	this.X509cert1 = x509cert1
	var x509cert2 string = ""
	this.X509cert2 = &x509cert2
	var x509cert3 string = ""
	this.X509cert3 = &x509cert3
	var rolesAttribute string = ""
	this.RolesAttribute = &rolesAttribute
	var logoutEnabled bool = false
	this.LogoutEnabled = &logoutEnabled
	var logoutUrl string = ""
	this.LogoutUrl = &logoutUrl
	var emailAttribute string = ""
	this.EmailAttribute = &emailAttribute
	var debugMode bool = false
	this.DebugMode = &debugMode
	var signAuthnRequest bool = false
	this.SignAuthnRequest = &signAuthnRequest
	var disableRequestedAuthnContext bool = false
	this.DisableRequestedAuthnContext = &disableRequestedAuthnContext
	var isRedirectBinding bool = false
	this.IsRedirectBinding = &isRedirectBinding
	this.Certificate = certificate
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.ModifiedAt = modifiedAt
	this.ModifiedBy = modifiedBy
	this.Id = id
	var assertionConsumerUrl string = ""
	this.AssertionConsumerUrl = &assertionConsumerUrl
	var entityId string = ""
	this.EntityId = &entityId
	return &this
}

// NewSamlIdentityProviderWithDefaults instantiates a new SamlIdentityProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSamlIdentityProviderWithDefaults() *SamlIdentityProvider {
	this := SamlIdentityProvider{}
	var spInitiatedLoginPath string = ""
	this.SpInitiatedLoginPath = &spInitiatedLoginPath
	var spInitiatedLoginEnabled bool = false
	this.SpInitiatedLoginEnabled = &spInitiatedLoginEnabled
	var authnRequestUrl string = ""
	this.AuthnRequestUrl = &authnRequestUrl
	var x509cert2 string = ""
	this.X509cert2 = &x509cert2
	var x509cert3 string = ""
	this.X509cert3 = &x509cert3
	var rolesAttribute string = ""
	this.RolesAttribute = &rolesAttribute
	var logoutEnabled bool = false
	this.LogoutEnabled = &logoutEnabled
	var logoutUrl string = ""
	this.LogoutUrl = &logoutUrl
	var emailAttribute string = ""
	this.EmailAttribute = &emailAttribute
	var debugMode bool = false
	this.DebugMode = &debugMode
	var signAuthnRequest bool = false
	this.SignAuthnRequest = &signAuthnRequest
	var disableRequestedAuthnContext bool = false
	this.DisableRequestedAuthnContext = &disableRequestedAuthnContext
	var isRedirectBinding bool = false
	this.IsRedirectBinding = &isRedirectBinding
	var assertionConsumerUrl string = ""
	this.AssertionConsumerUrl = &assertionConsumerUrl
	var entityId string = ""
	this.EntityId = &entityId
	return &this
}

// GetSpInitiatedLoginPath returns the SpInitiatedLoginPath field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetSpInitiatedLoginPath() string {
	if o == nil || o.SpInitiatedLoginPath == nil {
		var ret string
		return ret
	}
	return *o.SpInitiatedLoginPath
}

// GetSpInitiatedLoginPathOk returns a tuple with the SpInitiatedLoginPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetSpInitiatedLoginPathOk() (*string, bool) {
	if o == nil || o.SpInitiatedLoginPath == nil {
		return nil, false
	}
	return o.SpInitiatedLoginPath, true
}

// HasSpInitiatedLoginPath returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasSpInitiatedLoginPath() bool {
	if o != nil && o.SpInitiatedLoginPath != nil {
		return true
	}

	return false
}

// SetSpInitiatedLoginPath gets a reference to the given string and assigns it to the SpInitiatedLoginPath field.
func (o *SamlIdentityProvider) SetSpInitiatedLoginPath(v string) {
	o.SpInitiatedLoginPath = &v
}

// GetConfigurationName returns the ConfigurationName field value
func (o *SamlIdentityProvider) GetConfigurationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConfigurationName
}

// GetConfigurationNameOk returns a tuple with the ConfigurationName field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetConfigurationNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigurationName, true
}

// SetConfigurationName sets field value
func (o *SamlIdentityProvider) SetConfigurationName(v string) {
	o.ConfigurationName = v
}

// GetIssuer returns the Issuer field value
func (o *SamlIdentityProvider) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetIssuerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *SamlIdentityProvider) SetIssuer(v string) {
	o.Issuer = v
}

// GetSpInitiatedLoginEnabled returns the SpInitiatedLoginEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetSpInitiatedLoginEnabled() bool {
	if o == nil || o.SpInitiatedLoginEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SpInitiatedLoginEnabled
}

// GetSpInitiatedLoginEnabledOk returns a tuple with the SpInitiatedLoginEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetSpInitiatedLoginEnabledOk() (*bool, bool) {
	if o == nil || o.SpInitiatedLoginEnabled == nil {
		return nil, false
	}
	return o.SpInitiatedLoginEnabled, true
}

// HasSpInitiatedLoginEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasSpInitiatedLoginEnabled() bool {
	if o != nil && o.SpInitiatedLoginEnabled != nil {
		return true
	}

	return false
}

// SetSpInitiatedLoginEnabled gets a reference to the given bool and assigns it to the SpInitiatedLoginEnabled field.
func (o *SamlIdentityProvider) SetSpInitiatedLoginEnabled(v bool) {
	o.SpInitiatedLoginEnabled = &v
}

// GetAuthnRequestUrl returns the AuthnRequestUrl field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetAuthnRequestUrl() string {
	if o == nil || o.AuthnRequestUrl == nil {
		var ret string
		return ret
	}
	return *o.AuthnRequestUrl
}

// GetAuthnRequestUrlOk returns a tuple with the AuthnRequestUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetAuthnRequestUrlOk() (*string, bool) {
	if o == nil || o.AuthnRequestUrl == nil {
		return nil, false
	}
	return o.AuthnRequestUrl, true
}

// HasAuthnRequestUrl returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasAuthnRequestUrl() bool {
	if o != nil && o.AuthnRequestUrl != nil {
		return true
	}

	return false
}

// SetAuthnRequestUrl gets a reference to the given string and assigns it to the AuthnRequestUrl field.
func (o *SamlIdentityProvider) SetAuthnRequestUrl(v string) {
	o.AuthnRequestUrl = &v
}

// GetX509cert1 returns the X509cert1 field value
func (o *SamlIdentityProvider) GetX509cert1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.X509cert1
}

// GetX509cert1Ok returns a tuple with the X509cert1 field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetX509cert1Ok() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.X509cert1, true
}

// SetX509cert1 sets field value
func (o *SamlIdentityProvider) SetX509cert1(v string) {
	o.X509cert1 = v
}

// GetX509cert2 returns the X509cert2 field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetX509cert2() string {
	if o == nil || o.X509cert2 == nil {
		var ret string
		return ret
	}
	return *o.X509cert2
}

// GetX509cert2Ok returns a tuple with the X509cert2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetX509cert2Ok() (*string, bool) {
	if o == nil || o.X509cert2 == nil {
		return nil, false
	}
	return o.X509cert2, true
}

// HasX509cert2 returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasX509cert2() bool {
	if o != nil && o.X509cert2 != nil {
		return true
	}

	return false
}

// SetX509cert2 gets a reference to the given string and assigns it to the X509cert2 field.
func (o *SamlIdentityProvider) SetX509cert2(v string) {
	o.X509cert2 = &v
}

// GetX509cert3 returns the X509cert3 field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetX509cert3() string {
	if o == nil || o.X509cert3 == nil {
		var ret string
		return ret
	}
	return *o.X509cert3
}

// GetX509cert3Ok returns a tuple with the X509cert3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetX509cert3Ok() (*string, bool) {
	if o == nil || o.X509cert3 == nil {
		return nil, false
	}
	return o.X509cert3, true
}

// HasX509cert3 returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasX509cert3() bool {
	if o != nil && o.X509cert3 != nil {
		return true
	}

	return false
}

// SetX509cert3 gets a reference to the given string and assigns it to the X509cert3 field.
func (o *SamlIdentityProvider) SetX509cert3(v string) {
	o.X509cert3 = &v
}

// GetOnDemandProvisioningEnabled returns the OnDemandProvisioningEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetOnDemandProvisioningEnabled() OnDemandProvisioningInfo {
	if o == nil || o.OnDemandProvisioningEnabled == nil {
		var ret OnDemandProvisioningInfo
		return ret
	}
	return *o.OnDemandProvisioningEnabled
}

// GetOnDemandProvisioningEnabledOk returns a tuple with the OnDemandProvisioningEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetOnDemandProvisioningEnabledOk() (*OnDemandProvisioningInfo, bool) {
	if o == nil || o.OnDemandProvisioningEnabled == nil {
		return nil, false
	}
	return o.OnDemandProvisioningEnabled, true
}

// HasOnDemandProvisioningEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasOnDemandProvisioningEnabled() bool {
	if o != nil && o.OnDemandProvisioningEnabled != nil {
		return true
	}

	return false
}

// SetOnDemandProvisioningEnabled gets a reference to the given OnDemandProvisioningInfo and assigns it to the OnDemandProvisioningEnabled field.
func (o *SamlIdentityProvider) SetOnDemandProvisioningEnabled(v OnDemandProvisioningInfo) {
	o.OnDemandProvisioningEnabled = &v
}

// GetRolesAttribute returns the RolesAttribute field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetRolesAttribute() string {
	if o == nil || o.RolesAttribute == nil {
		var ret string
		return ret
	}
	return *o.RolesAttribute
}

// GetRolesAttributeOk returns a tuple with the RolesAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetRolesAttributeOk() (*string, bool) {
	if o == nil || o.RolesAttribute == nil {
		return nil, false
	}
	return o.RolesAttribute, true
}

// HasRolesAttribute returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasRolesAttribute() bool {
	if o != nil && o.RolesAttribute != nil {
		return true
	}

	return false
}

// SetRolesAttribute gets a reference to the given string and assigns it to the RolesAttribute field.
func (o *SamlIdentityProvider) SetRolesAttribute(v string) {
	o.RolesAttribute = &v
}

// GetLogoutEnabled returns the LogoutEnabled field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetLogoutEnabled() bool {
	if o == nil || o.LogoutEnabled == nil {
		var ret bool
		return ret
	}
	return *o.LogoutEnabled
}

// GetLogoutEnabledOk returns a tuple with the LogoutEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetLogoutEnabledOk() (*bool, bool) {
	if o == nil || o.LogoutEnabled == nil {
		return nil, false
	}
	return o.LogoutEnabled, true
}

// HasLogoutEnabled returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasLogoutEnabled() bool {
	if o != nil && o.LogoutEnabled != nil {
		return true
	}

	return false
}

// SetLogoutEnabled gets a reference to the given bool and assigns it to the LogoutEnabled field.
func (o *SamlIdentityProvider) SetLogoutEnabled(v bool) {
	o.LogoutEnabled = &v
}

// GetLogoutUrl returns the LogoutUrl field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetLogoutUrl() string {
	if o == nil || o.LogoutUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoutUrl
}

// GetLogoutUrlOk returns a tuple with the LogoutUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetLogoutUrlOk() (*string, bool) {
	if o == nil || o.LogoutUrl == nil {
		return nil, false
	}
	return o.LogoutUrl, true
}

// HasLogoutUrl returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasLogoutUrl() bool {
	if o != nil && o.LogoutUrl != nil {
		return true
	}

	return false
}

// SetLogoutUrl gets a reference to the given string and assigns it to the LogoutUrl field.
func (o *SamlIdentityProvider) SetLogoutUrl(v string) {
	o.LogoutUrl = &v
}

// GetEmailAttribute returns the EmailAttribute field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetEmailAttribute() string {
	if o == nil || o.EmailAttribute == nil {
		var ret string
		return ret
	}
	return *o.EmailAttribute
}

// GetEmailAttributeOk returns a tuple with the EmailAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetEmailAttributeOk() (*string, bool) {
	if o == nil || o.EmailAttribute == nil {
		return nil, false
	}
	return o.EmailAttribute, true
}

// HasEmailAttribute returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasEmailAttribute() bool {
	if o != nil && o.EmailAttribute != nil {
		return true
	}

	return false
}

// SetEmailAttribute gets a reference to the given string and assigns it to the EmailAttribute field.
func (o *SamlIdentityProvider) SetEmailAttribute(v string) {
	o.EmailAttribute = &v
}

// GetDebugMode returns the DebugMode field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetDebugMode() bool {
	if o == nil || o.DebugMode == nil {
		var ret bool
		return ret
	}
	return *o.DebugMode
}

// GetDebugModeOk returns a tuple with the DebugMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetDebugModeOk() (*bool, bool) {
	if o == nil || o.DebugMode == nil {
		return nil, false
	}
	return o.DebugMode, true
}

// HasDebugMode returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasDebugMode() bool {
	if o != nil && o.DebugMode != nil {
		return true
	}

	return false
}

// SetDebugMode gets a reference to the given bool and assigns it to the DebugMode field.
func (o *SamlIdentityProvider) SetDebugMode(v bool) {
	o.DebugMode = &v
}

// GetSignAuthnRequest returns the SignAuthnRequest field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetSignAuthnRequest() bool {
	if o == nil || o.SignAuthnRequest == nil {
		var ret bool
		return ret
	}
	return *o.SignAuthnRequest
}

// GetSignAuthnRequestOk returns a tuple with the SignAuthnRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetSignAuthnRequestOk() (*bool, bool) {
	if o == nil || o.SignAuthnRequest == nil {
		return nil, false
	}
	return o.SignAuthnRequest, true
}

// HasSignAuthnRequest returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasSignAuthnRequest() bool {
	if o != nil && o.SignAuthnRequest != nil {
		return true
	}

	return false
}

// SetSignAuthnRequest gets a reference to the given bool and assigns it to the SignAuthnRequest field.
func (o *SamlIdentityProvider) SetSignAuthnRequest(v bool) {
	o.SignAuthnRequest = &v
}

// GetDisableRequestedAuthnContext returns the DisableRequestedAuthnContext field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetDisableRequestedAuthnContext() bool {
	if o == nil || o.DisableRequestedAuthnContext == nil {
		var ret bool
		return ret
	}
	return *o.DisableRequestedAuthnContext
}

// GetDisableRequestedAuthnContextOk returns a tuple with the DisableRequestedAuthnContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetDisableRequestedAuthnContextOk() (*bool, bool) {
	if o == nil || o.DisableRequestedAuthnContext == nil {
		return nil, false
	}
	return o.DisableRequestedAuthnContext, true
}

// HasDisableRequestedAuthnContext returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasDisableRequestedAuthnContext() bool {
	if o != nil && o.DisableRequestedAuthnContext != nil {
		return true
	}

	return false
}

// SetDisableRequestedAuthnContext gets a reference to the given bool and assigns it to the DisableRequestedAuthnContext field.
func (o *SamlIdentityProvider) SetDisableRequestedAuthnContext(v bool) {
	o.DisableRequestedAuthnContext = &v
}

// GetIsRedirectBinding returns the IsRedirectBinding field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetIsRedirectBinding() bool {
	if o == nil || o.IsRedirectBinding == nil {
		var ret bool
		return ret
	}
	return *o.IsRedirectBinding
}

// GetIsRedirectBindingOk returns a tuple with the IsRedirectBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetIsRedirectBindingOk() (*bool, bool) {
	if o == nil || o.IsRedirectBinding == nil {
		return nil, false
	}
	return o.IsRedirectBinding, true
}

// HasIsRedirectBinding returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasIsRedirectBinding() bool {
	if o != nil && o.IsRedirectBinding != nil {
		return true
	}

	return false
}

// SetIsRedirectBinding gets a reference to the given bool and assigns it to the IsRedirectBinding field.
func (o *SamlIdentityProvider) SetIsRedirectBinding(v bool) {
	o.IsRedirectBinding = &v
}

// GetCertificate returns the Certificate field value
func (o *SamlIdentityProvider) GetCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Certificate, true
}

// SetCertificate sets field value
func (o *SamlIdentityProvider) SetCertificate(v string) {
	o.Certificate = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SamlIdentityProvider) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SamlIdentityProvider) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SamlIdentityProvider) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetCreatedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SamlIdentityProvider) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetModifiedAt returns the ModifiedAt field value
func (o *SamlIdentityProvider) GetModifiedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ModifiedAt
}

// GetModifiedAtOk returns a tuple with the ModifiedAt field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetModifiedAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedAt, true
}

// SetModifiedAt sets field value
func (o *SamlIdentityProvider) SetModifiedAt(v time.Time) {
	o.ModifiedAt = v
}

// GetModifiedBy returns the ModifiedBy field value
func (o *SamlIdentityProvider) GetModifiedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ModifiedBy
}

// GetModifiedByOk returns a tuple with the ModifiedBy field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetModifiedByOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ModifiedBy, true
}

// SetModifiedBy sets field value
func (o *SamlIdentityProvider) SetModifiedBy(v string) {
	o.ModifiedBy = v
}

// GetId returns the Id field value
func (o *SamlIdentityProvider) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SamlIdentityProvider) SetId(v string) {
	o.Id = v
}

// GetAssertionConsumerUrl returns the AssertionConsumerUrl field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetAssertionConsumerUrl() string {
	if o == nil || o.AssertionConsumerUrl == nil {
		var ret string
		return ret
	}
	return *o.AssertionConsumerUrl
}

// GetAssertionConsumerUrlOk returns a tuple with the AssertionConsumerUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetAssertionConsumerUrlOk() (*string, bool) {
	if o == nil || o.AssertionConsumerUrl == nil {
		return nil, false
	}
	return o.AssertionConsumerUrl, true
}

// HasAssertionConsumerUrl returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasAssertionConsumerUrl() bool {
	if o != nil && o.AssertionConsumerUrl != nil {
		return true
	}

	return false
}

// SetAssertionConsumerUrl gets a reference to the given string and assigns it to the AssertionConsumerUrl field.
func (o *SamlIdentityProvider) SetAssertionConsumerUrl(v string) {
	o.AssertionConsumerUrl = &v
}

// GetEntityId returns the EntityId field value if set, zero value otherwise.
func (o *SamlIdentityProvider) GetEntityId() string {
	if o == nil || o.EntityId == nil {
		var ret string
		return ret
	}
	return *o.EntityId
}

// GetEntityIdOk returns a tuple with the EntityId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SamlIdentityProvider) GetEntityIdOk() (*string, bool) {
	if o == nil || o.EntityId == nil {
		return nil, false
	}
	return o.EntityId, true
}

// HasEntityId returns a boolean if a field has been set.
func (o *SamlIdentityProvider) HasEntityId() bool {
	if o != nil && o.EntityId != nil {
		return true
	}

	return false
}

// SetEntityId gets a reference to the given string and assigns it to the EntityId field.
func (o *SamlIdentityProvider) SetEntityId(v string) {
	o.EntityId = &v
}

func (o SamlIdentityProvider) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpInitiatedLoginPath != nil {
		toSerialize["spInitiatedLoginPath"] = o.SpInitiatedLoginPath
	}
	if true {
		toSerialize["configurationName"] = o.ConfigurationName
	}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if o.SpInitiatedLoginEnabled != nil {
		toSerialize["spInitiatedLoginEnabled"] = o.SpInitiatedLoginEnabled
	}
	if o.AuthnRequestUrl != nil {
		toSerialize["authnRequestUrl"] = o.AuthnRequestUrl
	}
	if true {
		toSerialize["x509cert1"] = o.X509cert1
	}
	if o.X509cert2 != nil {
		toSerialize["x509cert2"] = o.X509cert2
	}
	if o.X509cert3 != nil {
		toSerialize["x509cert3"] = o.X509cert3
	}
	if o.OnDemandProvisioningEnabled != nil {
		toSerialize["onDemandProvisioningEnabled"] = o.OnDemandProvisioningEnabled
	}
	if o.RolesAttribute != nil {
		toSerialize["rolesAttribute"] = o.RolesAttribute
	}
	if o.LogoutEnabled != nil {
		toSerialize["logoutEnabled"] = o.LogoutEnabled
	}
	if o.LogoutUrl != nil {
		toSerialize["logoutUrl"] = o.LogoutUrl
	}
	if o.EmailAttribute != nil {
		toSerialize["emailAttribute"] = o.EmailAttribute
	}
	if o.DebugMode != nil {
		toSerialize["debugMode"] = o.DebugMode
	}
	if o.SignAuthnRequest != nil {
		toSerialize["signAuthnRequest"] = o.SignAuthnRequest
	}
	if o.DisableRequestedAuthnContext != nil {
		toSerialize["disableRequestedAuthnContext"] = o.DisableRequestedAuthnContext
	}
	if o.IsRedirectBinding != nil {
		toSerialize["isRedirectBinding"] = o.IsRedirectBinding
	}
	if true {
		toSerialize["certificate"] = o.Certificate
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
	if o.AssertionConsumerUrl != nil {
		toSerialize["assertionConsumerUrl"] = o.AssertionConsumerUrl
	}
	if o.EntityId != nil {
		toSerialize["entityId"] = o.EntityId
	}
	return json.Marshal(toSerialize)
}

type NullableSamlIdentityProvider struct {
	value *SamlIdentityProvider
	isSet bool
}

func (v NullableSamlIdentityProvider) Get() *SamlIdentityProvider {
	return v.value
}

func (v *NullableSamlIdentityProvider) Set(val *SamlIdentityProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableSamlIdentityProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableSamlIdentityProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSamlIdentityProvider(val *SamlIdentityProvider) *NullableSamlIdentityProvider {
	return &NullableSamlIdentityProvider{value: val, isSet: true}
}

func (v NullableSamlIdentityProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSamlIdentityProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



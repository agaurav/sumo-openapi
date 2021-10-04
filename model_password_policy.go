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

// PasswordPolicy Password Policy
type PasswordPolicy struct {
	// The minimum length of the password.
	MinLength *int32 `json:"minLength,omitempty"`
	// The maximum length of the password. (Setting this to any value other than 128 is no longer supported; this field may be deprecated in the future.)
	MaxLength *int32 `json:"maxLength,omitempty"`
	// If the password must contain lower case characters.
	MustContainLowercase *bool `json:"mustContainLowercase,omitempty"`
	// If the password must contain upper case characters.
	MustContainUppercase *bool `json:"mustContainUppercase,omitempty"`
	// If the password must contain digits.
	MustContainDigits *bool `json:"mustContainDigits,omitempty"`
	// If the password must contain special characters.
	MustContainSpecialChars *bool `json:"mustContainSpecialChars,omitempty"`
	// Maximum number of days that a password can be used before user is required to change it. Put -1 if the user should not have to change their password.
	MaxPasswordAgeInDays *int32 `json:"maxPasswordAgeInDays,omitempty"`
	// The minimum number of unique new passwords that a user must use before an old password can be reused.
	MinUniquePasswords *int32 `json:"minUniquePasswords,omitempty"`
	// Number of failed login attempts allowed before account is locked-out.
	AccountLockoutThreshold *int32 `json:"accountLockoutThreshold,omitempty"`
	// The duration of time in minutes that must elapse from the first failed login attempt after which failed login count is reset to 0.
	FailedLoginResetDurationInMins *int32 `json:"failedLoginResetDurationInMins,omitempty"`
	// The duration of time in minutes that a locked-out account remained locked before getting unlocked automatically.
	AccountLockoutDurationInMins *int32 `json:"accountLockoutDurationInMins,omitempty"`
	// If MFA should be required to log in. By default, this field is set to `false`.
	RequireMfa *bool `json:"requireMfa,omitempty"`
	// If MFA should be remembered on the browser.
	RememberMfa *bool `json:"rememberMfa,omitempty"`
}

// NewPasswordPolicy instantiates a new PasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicy() *PasswordPolicy {
	this := PasswordPolicy{}
	var minLength int32 = 8
	this.MinLength = &minLength
	var maxLength int32 = 128
	this.MaxLength = &maxLength
	var mustContainLowercase bool = true
	this.MustContainLowercase = &mustContainLowercase
	var mustContainUppercase bool = true
	this.MustContainUppercase = &mustContainUppercase
	var mustContainDigits bool = true
	this.MustContainDigits = &mustContainDigits
	var mustContainSpecialChars bool = true
	this.MustContainSpecialChars = &mustContainSpecialChars
	var maxPasswordAgeInDays int32 = 365
	this.MaxPasswordAgeInDays = &maxPasswordAgeInDays
	var minUniquePasswords int32 = 10
	this.MinUniquePasswords = &minUniquePasswords
	var accountLockoutThreshold int32 = 6
	this.AccountLockoutThreshold = &accountLockoutThreshold
	var failedLoginResetDurationInMins int32 = 10
	this.FailedLoginResetDurationInMins = &failedLoginResetDurationInMins
	var accountLockoutDurationInMins int32 = 30
	this.AccountLockoutDurationInMins = &accountLockoutDurationInMins
	var requireMfa bool = false
	this.RequireMfa = &requireMfa
	var rememberMfa bool = true
	this.RememberMfa = &rememberMfa
	return &this
}

// NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyWithDefaults() *PasswordPolicy {
	this := PasswordPolicy{}
	var minLength int32 = 8
	this.MinLength = &minLength
	var maxLength int32 = 128
	this.MaxLength = &maxLength
	var mustContainLowercase bool = true
	this.MustContainLowercase = &mustContainLowercase
	var mustContainUppercase bool = true
	this.MustContainUppercase = &mustContainUppercase
	var mustContainDigits bool = true
	this.MustContainDigits = &mustContainDigits
	var mustContainSpecialChars bool = true
	this.MustContainSpecialChars = &mustContainSpecialChars
	var maxPasswordAgeInDays int32 = 365
	this.MaxPasswordAgeInDays = &maxPasswordAgeInDays
	var minUniquePasswords int32 = 10
	this.MinUniquePasswords = &minUniquePasswords
	var accountLockoutThreshold int32 = 6
	this.AccountLockoutThreshold = &accountLockoutThreshold
	var failedLoginResetDurationInMins int32 = 10
	this.FailedLoginResetDurationInMins = &failedLoginResetDurationInMins
	var accountLockoutDurationInMins int32 = 30
	this.AccountLockoutDurationInMins = &accountLockoutDurationInMins
	var requireMfa bool = false
	this.RequireMfa = &requireMfa
	var rememberMfa bool = true
	this.RememberMfa = &rememberMfa
	return &this
}

// GetMinLength returns the MinLength field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMinLength() int32 {
	if o == nil || o.MinLength == nil {
		var ret int32
		return ret
	}
	return *o.MinLength
}

// GetMinLengthOk returns a tuple with the MinLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMinLengthOk() (*int32, bool) {
	if o == nil || o.MinLength == nil {
		return nil, false
	}
	return o.MinLength, true
}

// HasMinLength returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMinLength() bool {
	if o != nil && o.MinLength != nil {
		return true
	}

	return false
}

// SetMinLength gets a reference to the given int32 and assigns it to the MinLength field.
func (o *PasswordPolicy) SetMinLength(v int32) {
	o.MinLength = &v
}

// GetMaxLength returns the MaxLength field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMaxLength() int32 {
	if o == nil || o.MaxLength == nil {
		var ret int32
		return ret
	}
	return *o.MaxLength
}

// GetMaxLengthOk returns a tuple with the MaxLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMaxLengthOk() (*int32, bool) {
	if o == nil || o.MaxLength == nil {
		return nil, false
	}
	return o.MaxLength, true
}

// HasMaxLength returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMaxLength() bool {
	if o != nil && o.MaxLength != nil {
		return true
	}

	return false
}

// SetMaxLength gets a reference to the given int32 and assigns it to the MaxLength field.
func (o *PasswordPolicy) SetMaxLength(v int32) {
	o.MaxLength = &v
}

// GetMustContainLowercase returns the MustContainLowercase field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMustContainLowercase() bool {
	if o == nil || o.MustContainLowercase == nil {
		var ret bool
		return ret
	}
	return *o.MustContainLowercase
}

// GetMustContainLowercaseOk returns a tuple with the MustContainLowercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMustContainLowercaseOk() (*bool, bool) {
	if o == nil || o.MustContainLowercase == nil {
		return nil, false
	}
	return o.MustContainLowercase, true
}

// HasMustContainLowercase returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMustContainLowercase() bool {
	if o != nil && o.MustContainLowercase != nil {
		return true
	}

	return false
}

// SetMustContainLowercase gets a reference to the given bool and assigns it to the MustContainLowercase field.
func (o *PasswordPolicy) SetMustContainLowercase(v bool) {
	o.MustContainLowercase = &v
}

// GetMustContainUppercase returns the MustContainUppercase field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMustContainUppercase() bool {
	if o == nil || o.MustContainUppercase == nil {
		var ret bool
		return ret
	}
	return *o.MustContainUppercase
}

// GetMustContainUppercaseOk returns a tuple with the MustContainUppercase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMustContainUppercaseOk() (*bool, bool) {
	if o == nil || o.MustContainUppercase == nil {
		return nil, false
	}
	return o.MustContainUppercase, true
}

// HasMustContainUppercase returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMustContainUppercase() bool {
	if o != nil && o.MustContainUppercase != nil {
		return true
	}

	return false
}

// SetMustContainUppercase gets a reference to the given bool and assigns it to the MustContainUppercase field.
func (o *PasswordPolicy) SetMustContainUppercase(v bool) {
	o.MustContainUppercase = &v
}

// GetMustContainDigits returns the MustContainDigits field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMustContainDigits() bool {
	if o == nil || o.MustContainDigits == nil {
		var ret bool
		return ret
	}
	return *o.MustContainDigits
}

// GetMustContainDigitsOk returns a tuple with the MustContainDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMustContainDigitsOk() (*bool, bool) {
	if o == nil || o.MustContainDigits == nil {
		return nil, false
	}
	return o.MustContainDigits, true
}

// HasMustContainDigits returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMustContainDigits() bool {
	if o != nil && o.MustContainDigits != nil {
		return true
	}

	return false
}

// SetMustContainDigits gets a reference to the given bool and assigns it to the MustContainDigits field.
func (o *PasswordPolicy) SetMustContainDigits(v bool) {
	o.MustContainDigits = &v
}

// GetMustContainSpecialChars returns the MustContainSpecialChars field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMustContainSpecialChars() bool {
	if o == nil || o.MustContainSpecialChars == nil {
		var ret bool
		return ret
	}
	return *o.MustContainSpecialChars
}

// GetMustContainSpecialCharsOk returns a tuple with the MustContainSpecialChars field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMustContainSpecialCharsOk() (*bool, bool) {
	if o == nil || o.MustContainSpecialChars == nil {
		return nil, false
	}
	return o.MustContainSpecialChars, true
}

// HasMustContainSpecialChars returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMustContainSpecialChars() bool {
	if o != nil && o.MustContainSpecialChars != nil {
		return true
	}

	return false
}

// SetMustContainSpecialChars gets a reference to the given bool and assigns it to the MustContainSpecialChars field.
func (o *PasswordPolicy) SetMustContainSpecialChars(v bool) {
	o.MustContainSpecialChars = &v
}

// GetMaxPasswordAgeInDays returns the MaxPasswordAgeInDays field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMaxPasswordAgeInDays() int32 {
	if o == nil || o.MaxPasswordAgeInDays == nil {
		var ret int32
		return ret
	}
	return *o.MaxPasswordAgeInDays
}

// GetMaxPasswordAgeInDaysOk returns a tuple with the MaxPasswordAgeInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMaxPasswordAgeInDaysOk() (*int32, bool) {
	if o == nil || o.MaxPasswordAgeInDays == nil {
		return nil, false
	}
	return o.MaxPasswordAgeInDays, true
}

// HasMaxPasswordAgeInDays returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMaxPasswordAgeInDays() bool {
	if o != nil && o.MaxPasswordAgeInDays != nil {
		return true
	}

	return false
}

// SetMaxPasswordAgeInDays gets a reference to the given int32 and assigns it to the MaxPasswordAgeInDays field.
func (o *PasswordPolicy) SetMaxPasswordAgeInDays(v int32) {
	o.MaxPasswordAgeInDays = &v
}

// GetMinUniquePasswords returns the MinUniquePasswords field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMinUniquePasswords() int32 {
	if o == nil || o.MinUniquePasswords == nil {
		var ret int32
		return ret
	}
	return *o.MinUniquePasswords
}

// GetMinUniquePasswordsOk returns a tuple with the MinUniquePasswords field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMinUniquePasswordsOk() (*int32, bool) {
	if o == nil || o.MinUniquePasswords == nil {
		return nil, false
	}
	return o.MinUniquePasswords, true
}

// HasMinUniquePasswords returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMinUniquePasswords() bool {
	if o != nil && o.MinUniquePasswords != nil {
		return true
	}

	return false
}

// SetMinUniquePasswords gets a reference to the given int32 and assigns it to the MinUniquePasswords field.
func (o *PasswordPolicy) SetMinUniquePasswords(v int32) {
	o.MinUniquePasswords = &v
}

// GetAccountLockoutThreshold returns the AccountLockoutThreshold field value if set, zero value otherwise.
func (o *PasswordPolicy) GetAccountLockoutThreshold() int32 {
	if o == nil || o.AccountLockoutThreshold == nil {
		var ret int32
		return ret
	}
	return *o.AccountLockoutThreshold
}

// GetAccountLockoutThresholdOk returns a tuple with the AccountLockoutThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetAccountLockoutThresholdOk() (*int32, bool) {
	if o == nil || o.AccountLockoutThreshold == nil {
		return nil, false
	}
	return o.AccountLockoutThreshold, true
}

// HasAccountLockoutThreshold returns a boolean if a field has been set.
func (o *PasswordPolicy) HasAccountLockoutThreshold() bool {
	if o != nil && o.AccountLockoutThreshold != nil {
		return true
	}

	return false
}

// SetAccountLockoutThreshold gets a reference to the given int32 and assigns it to the AccountLockoutThreshold field.
func (o *PasswordPolicy) SetAccountLockoutThreshold(v int32) {
	o.AccountLockoutThreshold = &v
}

// GetFailedLoginResetDurationInMins returns the FailedLoginResetDurationInMins field value if set, zero value otherwise.
func (o *PasswordPolicy) GetFailedLoginResetDurationInMins() int32 {
	if o == nil || o.FailedLoginResetDurationInMins == nil {
		var ret int32
		return ret
	}
	return *o.FailedLoginResetDurationInMins
}

// GetFailedLoginResetDurationInMinsOk returns a tuple with the FailedLoginResetDurationInMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetFailedLoginResetDurationInMinsOk() (*int32, bool) {
	if o == nil || o.FailedLoginResetDurationInMins == nil {
		return nil, false
	}
	return o.FailedLoginResetDurationInMins, true
}

// HasFailedLoginResetDurationInMins returns a boolean if a field has been set.
func (o *PasswordPolicy) HasFailedLoginResetDurationInMins() bool {
	if o != nil && o.FailedLoginResetDurationInMins != nil {
		return true
	}

	return false
}

// SetFailedLoginResetDurationInMins gets a reference to the given int32 and assigns it to the FailedLoginResetDurationInMins field.
func (o *PasswordPolicy) SetFailedLoginResetDurationInMins(v int32) {
	o.FailedLoginResetDurationInMins = &v
}

// GetAccountLockoutDurationInMins returns the AccountLockoutDurationInMins field value if set, zero value otherwise.
func (o *PasswordPolicy) GetAccountLockoutDurationInMins() int32 {
	if o == nil || o.AccountLockoutDurationInMins == nil {
		var ret int32
		return ret
	}
	return *o.AccountLockoutDurationInMins
}

// GetAccountLockoutDurationInMinsOk returns a tuple with the AccountLockoutDurationInMins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetAccountLockoutDurationInMinsOk() (*int32, bool) {
	if o == nil || o.AccountLockoutDurationInMins == nil {
		return nil, false
	}
	return o.AccountLockoutDurationInMins, true
}

// HasAccountLockoutDurationInMins returns a boolean if a field has been set.
func (o *PasswordPolicy) HasAccountLockoutDurationInMins() bool {
	if o != nil && o.AccountLockoutDurationInMins != nil {
		return true
	}

	return false
}

// SetAccountLockoutDurationInMins gets a reference to the given int32 and assigns it to the AccountLockoutDurationInMins field.
func (o *PasswordPolicy) SetAccountLockoutDurationInMins(v int32) {
	o.AccountLockoutDurationInMins = &v
}

// GetRequireMfa returns the RequireMfa field value if set, zero value otherwise.
func (o *PasswordPolicy) GetRequireMfa() bool {
	if o == nil || o.RequireMfa == nil {
		var ret bool
		return ret
	}
	return *o.RequireMfa
}

// GetRequireMfaOk returns a tuple with the RequireMfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetRequireMfaOk() (*bool, bool) {
	if o == nil || o.RequireMfa == nil {
		return nil, false
	}
	return o.RequireMfa, true
}

// HasRequireMfa returns a boolean if a field has been set.
func (o *PasswordPolicy) HasRequireMfa() bool {
	if o != nil && o.RequireMfa != nil {
		return true
	}

	return false
}

// SetRequireMfa gets a reference to the given bool and assigns it to the RequireMfa field.
func (o *PasswordPolicy) SetRequireMfa(v bool) {
	o.RequireMfa = &v
}

// GetRememberMfa returns the RememberMfa field value if set, zero value otherwise.
func (o *PasswordPolicy) GetRememberMfa() bool {
	if o == nil || o.RememberMfa == nil {
		var ret bool
		return ret
	}
	return *o.RememberMfa
}

// GetRememberMfaOk returns a tuple with the RememberMfa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetRememberMfaOk() (*bool, bool) {
	if o == nil || o.RememberMfa == nil {
		return nil, false
	}
	return o.RememberMfa, true
}

// HasRememberMfa returns a boolean if a field has been set.
func (o *PasswordPolicy) HasRememberMfa() bool {
	if o != nil && o.RememberMfa != nil {
		return true
	}

	return false
}

// SetRememberMfa gets a reference to the given bool and assigns it to the RememberMfa field.
func (o *PasswordPolicy) SetRememberMfa(v bool) {
	o.RememberMfa = &v
}

func (o PasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MinLength != nil {
		toSerialize["minLength"] = o.MinLength
	}
	if o.MaxLength != nil {
		toSerialize["maxLength"] = o.MaxLength
	}
	if o.MustContainLowercase != nil {
		toSerialize["mustContainLowercase"] = o.MustContainLowercase
	}
	if o.MustContainUppercase != nil {
		toSerialize["mustContainUppercase"] = o.MustContainUppercase
	}
	if o.MustContainDigits != nil {
		toSerialize["mustContainDigits"] = o.MustContainDigits
	}
	if o.MustContainSpecialChars != nil {
		toSerialize["mustContainSpecialChars"] = o.MustContainSpecialChars
	}
	if o.MaxPasswordAgeInDays != nil {
		toSerialize["maxPasswordAgeInDays"] = o.MaxPasswordAgeInDays
	}
	if o.MinUniquePasswords != nil {
		toSerialize["minUniquePasswords"] = o.MinUniquePasswords
	}
	if o.AccountLockoutThreshold != nil {
		toSerialize["accountLockoutThreshold"] = o.AccountLockoutThreshold
	}
	if o.FailedLoginResetDurationInMins != nil {
		toSerialize["failedLoginResetDurationInMins"] = o.FailedLoginResetDurationInMins
	}
	if o.AccountLockoutDurationInMins != nil {
		toSerialize["accountLockoutDurationInMins"] = o.AccountLockoutDurationInMins
	}
	if o.RequireMfa != nil {
		toSerialize["requireMfa"] = o.RequireMfa
	}
	if o.RememberMfa != nil {
		toSerialize["rememberMfa"] = o.RememberMfa
	}
	return json.Marshal(toSerialize)
}

type NullablePasswordPolicy struct {
	value *PasswordPolicy
	isSet bool
}

func (v NullablePasswordPolicy) Get() *PasswordPolicy {
	return v.value
}

func (v *NullablePasswordPolicy) Set(val *PasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicy(val *PasswordPolicy) *NullablePasswordPolicy {
	return &NullablePasswordPolicy{value: val, isSet: true}
}

func (v NullablePasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



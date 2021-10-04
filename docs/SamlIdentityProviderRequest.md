# SamlIdentityProviderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpInitiatedLoginPath** | Pointer to **string** | The identifier used to generate a unique URL for user login. | [optional] [default to ""]
**ConfigurationName** | **string** | Name of the SSO policy or another name used to describe the policy internally. | 
**Issuer** | **string** | The unique URL assigned to the organization by the SAML Identity Provider. | 
**SpInitiatedLoginEnabled** | Pointer to **bool** | True if Sumo Logic redirects users to your identity provider with a SAML AuthnRequest when signing in. | [optional] [default to false]
**AuthnRequestUrl** | Pointer to **string** | The URL that the identity provider has assigned for Sumo Logic to submit SAML authentication requests to the identity provider. | [optional] [default to ""]
**X509cert1** | **string** | The certificate is used to verify the signature in SAML assertions. | 
**X509cert2** | Pointer to **string** | The backup certificate used to verify the signature in SAML assertions when x509cert1 expires. | [optional] [default to ""]
**X509cert3** | Pointer to **string** | The backup certificate used to verify the signature in SAML assertions when x509cert1 expires and x509cert2 is empty. | [optional] [default to ""]
**OnDemandProvisioningEnabled** | Pointer to [**OnDemandProvisioningInfo**](OnDemandProvisioningInfo.md) |  | [optional] 
**RolesAttribute** | Pointer to **string** | The role that Sumo Logic will assign to users when they sign in. | [optional] [default to ""]
**LogoutEnabled** | Pointer to **bool** | True if users are redirected to a URL after signing out of Sumo Logic. | [optional] [default to false]
**LogoutUrl** | Pointer to **string** | The URL that users will be redirected to after signing out of Sumo Logic. | [optional] [default to ""]
**EmailAttribute** | Pointer to **string** | The email address of the new user account. | [optional] [default to ""]
**DebugMode** | Pointer to **bool** | True if additional details are included when a user fails to sign in. | [optional] [default to false]
**SignAuthnRequest** | Pointer to **bool** | True if Sumo Logic will send signed Authn requests to the identity provider. | [optional] [default to false]
**DisableRequestedAuthnContext** | Pointer to **bool** | True if Sumo Logic will include the RequestedAuthnContext element of the SAML AuthnRequests it sends to the identity provider. | [optional] [default to false]
**IsRedirectBinding** | Pointer to **bool** | True if the SAML binding is of HTTP Redirect type. | [optional] [default to false]

## Methods

### NewSamlIdentityProviderRequest

`func NewSamlIdentityProviderRequest(configurationName string, issuer string, x509cert1 string, ) *SamlIdentityProviderRequest`

NewSamlIdentityProviderRequest instantiates a new SamlIdentityProviderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlIdentityProviderRequestWithDefaults

`func NewSamlIdentityProviderRequestWithDefaults() *SamlIdentityProviderRequest`

NewSamlIdentityProviderRequestWithDefaults instantiates a new SamlIdentityProviderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpInitiatedLoginPath

`func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginPath() string`

GetSpInitiatedLoginPath returns the SpInitiatedLoginPath field if non-nil, zero value otherwise.

### GetSpInitiatedLoginPathOk

`func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginPathOk() (*string, bool)`

GetSpInitiatedLoginPathOk returns a tuple with the SpInitiatedLoginPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiatedLoginPath

`func (o *SamlIdentityProviderRequest) SetSpInitiatedLoginPath(v string)`

SetSpInitiatedLoginPath sets SpInitiatedLoginPath field to given value.

### HasSpInitiatedLoginPath

`func (o *SamlIdentityProviderRequest) HasSpInitiatedLoginPath() bool`

HasSpInitiatedLoginPath returns a boolean if a field has been set.

### GetConfigurationName

`func (o *SamlIdentityProviderRequest) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *SamlIdentityProviderRequest) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *SamlIdentityProviderRequest) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.


### GetIssuer

`func (o *SamlIdentityProviderRequest) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SamlIdentityProviderRequest) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SamlIdentityProviderRequest) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSpInitiatedLoginEnabled

`func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginEnabled() bool`

GetSpInitiatedLoginEnabled returns the SpInitiatedLoginEnabled field if non-nil, zero value otherwise.

### GetSpInitiatedLoginEnabledOk

`func (o *SamlIdentityProviderRequest) GetSpInitiatedLoginEnabledOk() (*bool, bool)`

GetSpInitiatedLoginEnabledOk returns a tuple with the SpInitiatedLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiatedLoginEnabled

`func (o *SamlIdentityProviderRequest) SetSpInitiatedLoginEnabled(v bool)`

SetSpInitiatedLoginEnabled sets SpInitiatedLoginEnabled field to given value.

### HasSpInitiatedLoginEnabled

`func (o *SamlIdentityProviderRequest) HasSpInitiatedLoginEnabled() bool`

HasSpInitiatedLoginEnabled returns a boolean if a field has been set.

### GetAuthnRequestUrl

`func (o *SamlIdentityProviderRequest) GetAuthnRequestUrl() string`

GetAuthnRequestUrl returns the AuthnRequestUrl field if non-nil, zero value otherwise.

### GetAuthnRequestUrlOk

`func (o *SamlIdentityProviderRequest) GetAuthnRequestUrlOk() (*string, bool)`

GetAuthnRequestUrlOk returns a tuple with the AuthnRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnRequestUrl

`func (o *SamlIdentityProviderRequest) SetAuthnRequestUrl(v string)`

SetAuthnRequestUrl sets AuthnRequestUrl field to given value.

### HasAuthnRequestUrl

`func (o *SamlIdentityProviderRequest) HasAuthnRequestUrl() bool`

HasAuthnRequestUrl returns a boolean if a field has been set.

### GetX509cert1

`func (o *SamlIdentityProviderRequest) GetX509cert1() string`

GetX509cert1 returns the X509cert1 field if non-nil, zero value otherwise.

### GetX509cert1Ok

`func (o *SamlIdentityProviderRequest) GetX509cert1Ok() (*string, bool)`

GetX509cert1Ok returns a tuple with the X509cert1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509cert1

`func (o *SamlIdentityProviderRequest) SetX509cert1(v string)`

SetX509cert1 sets X509cert1 field to given value.


### GetX509cert2

`func (o *SamlIdentityProviderRequest) GetX509cert2() string`

GetX509cert2 returns the X509cert2 field if non-nil, zero value otherwise.

### GetX509cert2Ok

`func (o *SamlIdentityProviderRequest) GetX509cert2Ok() (*string, bool)`

GetX509cert2Ok returns a tuple with the X509cert2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509cert2

`func (o *SamlIdentityProviderRequest) SetX509cert2(v string)`

SetX509cert2 sets X509cert2 field to given value.

### HasX509cert2

`func (o *SamlIdentityProviderRequest) HasX509cert2() bool`

HasX509cert2 returns a boolean if a field has been set.

### GetX509cert3

`func (o *SamlIdentityProviderRequest) GetX509cert3() string`

GetX509cert3 returns the X509cert3 field if non-nil, zero value otherwise.

### GetX509cert3Ok

`func (o *SamlIdentityProviderRequest) GetX509cert3Ok() (*string, bool)`

GetX509cert3Ok returns a tuple with the X509cert3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509cert3

`func (o *SamlIdentityProviderRequest) SetX509cert3(v string)`

SetX509cert3 sets X509cert3 field to given value.

### HasX509cert3

`func (o *SamlIdentityProviderRequest) HasX509cert3() bool`

HasX509cert3 returns a boolean if a field has been set.

### GetOnDemandProvisioningEnabled

`func (o *SamlIdentityProviderRequest) GetOnDemandProvisioningEnabled() OnDemandProvisioningInfo`

GetOnDemandProvisioningEnabled returns the OnDemandProvisioningEnabled field if non-nil, zero value otherwise.

### GetOnDemandProvisioningEnabledOk

`func (o *SamlIdentityProviderRequest) GetOnDemandProvisioningEnabledOk() (*OnDemandProvisioningInfo, bool)`

GetOnDemandProvisioningEnabledOk returns a tuple with the OnDemandProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandProvisioningEnabled

`func (o *SamlIdentityProviderRequest) SetOnDemandProvisioningEnabled(v OnDemandProvisioningInfo)`

SetOnDemandProvisioningEnabled sets OnDemandProvisioningEnabled field to given value.

### HasOnDemandProvisioningEnabled

`func (o *SamlIdentityProviderRequest) HasOnDemandProvisioningEnabled() bool`

HasOnDemandProvisioningEnabled returns a boolean if a field has been set.

### GetRolesAttribute

`func (o *SamlIdentityProviderRequest) GetRolesAttribute() string`

GetRolesAttribute returns the RolesAttribute field if non-nil, zero value otherwise.

### GetRolesAttributeOk

`func (o *SamlIdentityProviderRequest) GetRolesAttributeOk() (*string, bool)`

GetRolesAttributeOk returns a tuple with the RolesAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesAttribute

`func (o *SamlIdentityProviderRequest) SetRolesAttribute(v string)`

SetRolesAttribute sets RolesAttribute field to given value.

### HasRolesAttribute

`func (o *SamlIdentityProviderRequest) HasRolesAttribute() bool`

HasRolesAttribute returns a boolean if a field has been set.

### GetLogoutEnabled

`func (o *SamlIdentityProviderRequest) GetLogoutEnabled() bool`

GetLogoutEnabled returns the LogoutEnabled field if non-nil, zero value otherwise.

### GetLogoutEnabledOk

`func (o *SamlIdentityProviderRequest) GetLogoutEnabledOk() (*bool, bool)`

GetLogoutEnabledOk returns a tuple with the LogoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutEnabled

`func (o *SamlIdentityProviderRequest) SetLogoutEnabled(v bool)`

SetLogoutEnabled sets LogoutEnabled field to given value.

### HasLogoutEnabled

`func (o *SamlIdentityProviderRequest) HasLogoutEnabled() bool`

HasLogoutEnabled returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *SamlIdentityProviderRequest) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SamlIdentityProviderRequest) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SamlIdentityProviderRequest) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *SamlIdentityProviderRequest) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *SamlIdentityProviderRequest) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *SamlIdentityProviderRequest) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *SamlIdentityProviderRequest) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *SamlIdentityProviderRequest) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetDebugMode

`func (o *SamlIdentityProviderRequest) GetDebugMode() bool`

GetDebugMode returns the DebugMode field if non-nil, zero value otherwise.

### GetDebugModeOk

`func (o *SamlIdentityProviderRequest) GetDebugModeOk() (*bool, bool)`

GetDebugModeOk returns a tuple with the DebugMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugMode

`func (o *SamlIdentityProviderRequest) SetDebugMode(v bool)`

SetDebugMode sets DebugMode field to given value.

### HasDebugMode

`func (o *SamlIdentityProviderRequest) HasDebugMode() bool`

HasDebugMode returns a boolean if a field has been set.

### GetSignAuthnRequest

`func (o *SamlIdentityProviderRequest) GetSignAuthnRequest() bool`

GetSignAuthnRequest returns the SignAuthnRequest field if non-nil, zero value otherwise.

### GetSignAuthnRequestOk

`func (o *SamlIdentityProviderRequest) GetSignAuthnRequestOk() (*bool, bool)`

GetSignAuthnRequestOk returns a tuple with the SignAuthnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthnRequest

`func (o *SamlIdentityProviderRequest) SetSignAuthnRequest(v bool)`

SetSignAuthnRequest sets SignAuthnRequest field to given value.

### HasSignAuthnRequest

`func (o *SamlIdentityProviderRequest) HasSignAuthnRequest() bool`

HasSignAuthnRequest returns a boolean if a field has been set.

### GetDisableRequestedAuthnContext

`func (o *SamlIdentityProviderRequest) GetDisableRequestedAuthnContext() bool`

GetDisableRequestedAuthnContext returns the DisableRequestedAuthnContext field if non-nil, zero value otherwise.

### GetDisableRequestedAuthnContextOk

`func (o *SamlIdentityProviderRequest) GetDisableRequestedAuthnContextOk() (*bool, bool)`

GetDisableRequestedAuthnContextOk returns a tuple with the DisableRequestedAuthnContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRequestedAuthnContext

`func (o *SamlIdentityProviderRequest) SetDisableRequestedAuthnContext(v bool)`

SetDisableRequestedAuthnContext sets DisableRequestedAuthnContext field to given value.

### HasDisableRequestedAuthnContext

`func (o *SamlIdentityProviderRequest) HasDisableRequestedAuthnContext() bool`

HasDisableRequestedAuthnContext returns a boolean if a field has been set.

### GetIsRedirectBinding

`func (o *SamlIdentityProviderRequest) GetIsRedirectBinding() bool`

GetIsRedirectBinding returns the IsRedirectBinding field if non-nil, zero value otherwise.

### GetIsRedirectBindingOk

`func (o *SamlIdentityProviderRequest) GetIsRedirectBindingOk() (*bool, bool)`

GetIsRedirectBindingOk returns a tuple with the IsRedirectBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRedirectBinding

`func (o *SamlIdentityProviderRequest) SetIsRedirectBinding(v bool)`

SetIsRedirectBinding sets IsRedirectBinding field to given value.

### HasIsRedirectBinding

`func (o *SamlIdentityProviderRequest) HasIsRedirectBinding() bool`

HasIsRedirectBinding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



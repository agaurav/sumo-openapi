# SamlIdentityProvider

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
**Certificate** | **string** | Authentication Request Signing Certificate for the user. | 
**CreatedAt** | **time.Time** | Creation timestamp in UTC in [RFC3339](https://tools.ietf.org/html/rfc3339) format. | 
**CreatedBy** | **string** | Identifier of the user who created the resource. | 
**ModifiedAt** | **time.Time** | Last modification timestamp in UTC. | 
**ModifiedBy** | **string** | Identifier of the user who last modified the resource. | 
**Id** | **string** | Unique identifier of the SAML Identity Provider. | 
**AssertionConsumerUrl** | Pointer to **string** | The URL on Sumo Logic where the IdP will redirect to with its authentication response. | [optional] [default to ""]
**EntityId** | Pointer to **string** | A unique identifier that is the intended audience of the SAML assertion. | [optional] [default to ""]

## Methods

### NewSamlIdentityProvider

`func NewSamlIdentityProvider(configurationName string, issuer string, x509cert1 string, certificate string, createdAt time.Time, createdBy string, modifiedAt time.Time, modifiedBy string, id string, ) *SamlIdentityProvider`

NewSamlIdentityProvider instantiates a new SamlIdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSamlIdentityProviderWithDefaults

`func NewSamlIdentityProviderWithDefaults() *SamlIdentityProvider`

NewSamlIdentityProviderWithDefaults instantiates a new SamlIdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpInitiatedLoginPath

`func (o *SamlIdentityProvider) GetSpInitiatedLoginPath() string`

GetSpInitiatedLoginPath returns the SpInitiatedLoginPath field if non-nil, zero value otherwise.

### GetSpInitiatedLoginPathOk

`func (o *SamlIdentityProvider) GetSpInitiatedLoginPathOk() (*string, bool)`

GetSpInitiatedLoginPathOk returns a tuple with the SpInitiatedLoginPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiatedLoginPath

`func (o *SamlIdentityProvider) SetSpInitiatedLoginPath(v string)`

SetSpInitiatedLoginPath sets SpInitiatedLoginPath field to given value.

### HasSpInitiatedLoginPath

`func (o *SamlIdentityProvider) HasSpInitiatedLoginPath() bool`

HasSpInitiatedLoginPath returns a boolean if a field has been set.

### GetConfigurationName

`func (o *SamlIdentityProvider) GetConfigurationName() string`

GetConfigurationName returns the ConfigurationName field if non-nil, zero value otherwise.

### GetConfigurationNameOk

`func (o *SamlIdentityProvider) GetConfigurationNameOk() (*string, bool)`

GetConfigurationNameOk returns a tuple with the ConfigurationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationName

`func (o *SamlIdentityProvider) SetConfigurationName(v string)`

SetConfigurationName sets ConfigurationName field to given value.


### GetIssuer

`func (o *SamlIdentityProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *SamlIdentityProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *SamlIdentityProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSpInitiatedLoginEnabled

`func (o *SamlIdentityProvider) GetSpInitiatedLoginEnabled() bool`

GetSpInitiatedLoginEnabled returns the SpInitiatedLoginEnabled field if non-nil, zero value otherwise.

### GetSpInitiatedLoginEnabledOk

`func (o *SamlIdentityProvider) GetSpInitiatedLoginEnabledOk() (*bool, bool)`

GetSpInitiatedLoginEnabledOk returns a tuple with the SpInitiatedLoginEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpInitiatedLoginEnabled

`func (o *SamlIdentityProvider) SetSpInitiatedLoginEnabled(v bool)`

SetSpInitiatedLoginEnabled sets SpInitiatedLoginEnabled field to given value.

### HasSpInitiatedLoginEnabled

`func (o *SamlIdentityProvider) HasSpInitiatedLoginEnabled() bool`

HasSpInitiatedLoginEnabled returns a boolean if a field has been set.

### GetAuthnRequestUrl

`func (o *SamlIdentityProvider) GetAuthnRequestUrl() string`

GetAuthnRequestUrl returns the AuthnRequestUrl field if non-nil, zero value otherwise.

### GetAuthnRequestUrlOk

`func (o *SamlIdentityProvider) GetAuthnRequestUrlOk() (*string, bool)`

GetAuthnRequestUrlOk returns a tuple with the AuthnRequestUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnRequestUrl

`func (o *SamlIdentityProvider) SetAuthnRequestUrl(v string)`

SetAuthnRequestUrl sets AuthnRequestUrl field to given value.

### HasAuthnRequestUrl

`func (o *SamlIdentityProvider) HasAuthnRequestUrl() bool`

HasAuthnRequestUrl returns a boolean if a field has been set.

### GetX509cert1

`func (o *SamlIdentityProvider) GetX509cert1() string`

GetX509cert1 returns the X509cert1 field if non-nil, zero value otherwise.

### GetX509cert1Ok

`func (o *SamlIdentityProvider) GetX509cert1Ok() (*string, bool)`

GetX509cert1Ok returns a tuple with the X509cert1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509cert1

`func (o *SamlIdentityProvider) SetX509cert1(v string)`

SetX509cert1 sets X509cert1 field to given value.


### GetX509cert2

`func (o *SamlIdentityProvider) GetX509cert2() string`

GetX509cert2 returns the X509cert2 field if non-nil, zero value otherwise.

### GetX509cert2Ok

`func (o *SamlIdentityProvider) GetX509cert2Ok() (*string, bool)`

GetX509cert2Ok returns a tuple with the X509cert2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509cert2

`func (o *SamlIdentityProvider) SetX509cert2(v string)`

SetX509cert2 sets X509cert2 field to given value.

### HasX509cert2

`func (o *SamlIdentityProvider) HasX509cert2() bool`

HasX509cert2 returns a boolean if a field has been set.

### GetX509cert3

`func (o *SamlIdentityProvider) GetX509cert3() string`

GetX509cert3 returns the X509cert3 field if non-nil, zero value otherwise.

### GetX509cert3Ok

`func (o *SamlIdentityProvider) GetX509cert3Ok() (*string, bool)`

GetX509cert3Ok returns a tuple with the X509cert3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX509cert3

`func (o *SamlIdentityProvider) SetX509cert3(v string)`

SetX509cert3 sets X509cert3 field to given value.

### HasX509cert3

`func (o *SamlIdentityProvider) HasX509cert3() bool`

HasX509cert3 returns a boolean if a field has been set.

### GetOnDemandProvisioningEnabled

`func (o *SamlIdentityProvider) GetOnDemandProvisioningEnabled() OnDemandProvisioningInfo`

GetOnDemandProvisioningEnabled returns the OnDemandProvisioningEnabled field if non-nil, zero value otherwise.

### GetOnDemandProvisioningEnabledOk

`func (o *SamlIdentityProvider) GetOnDemandProvisioningEnabledOk() (*OnDemandProvisioningInfo, bool)`

GetOnDemandProvisioningEnabledOk returns a tuple with the OnDemandProvisioningEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDemandProvisioningEnabled

`func (o *SamlIdentityProvider) SetOnDemandProvisioningEnabled(v OnDemandProvisioningInfo)`

SetOnDemandProvisioningEnabled sets OnDemandProvisioningEnabled field to given value.

### HasOnDemandProvisioningEnabled

`func (o *SamlIdentityProvider) HasOnDemandProvisioningEnabled() bool`

HasOnDemandProvisioningEnabled returns a boolean if a field has been set.

### GetRolesAttribute

`func (o *SamlIdentityProvider) GetRolesAttribute() string`

GetRolesAttribute returns the RolesAttribute field if non-nil, zero value otherwise.

### GetRolesAttributeOk

`func (o *SamlIdentityProvider) GetRolesAttributeOk() (*string, bool)`

GetRolesAttributeOk returns a tuple with the RolesAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRolesAttribute

`func (o *SamlIdentityProvider) SetRolesAttribute(v string)`

SetRolesAttribute sets RolesAttribute field to given value.

### HasRolesAttribute

`func (o *SamlIdentityProvider) HasRolesAttribute() bool`

HasRolesAttribute returns a boolean if a field has been set.

### GetLogoutEnabled

`func (o *SamlIdentityProvider) GetLogoutEnabled() bool`

GetLogoutEnabled returns the LogoutEnabled field if non-nil, zero value otherwise.

### GetLogoutEnabledOk

`func (o *SamlIdentityProvider) GetLogoutEnabledOk() (*bool, bool)`

GetLogoutEnabledOk returns a tuple with the LogoutEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutEnabled

`func (o *SamlIdentityProvider) SetLogoutEnabled(v bool)`

SetLogoutEnabled sets LogoutEnabled field to given value.

### HasLogoutEnabled

`func (o *SamlIdentityProvider) HasLogoutEnabled() bool`

HasLogoutEnabled returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *SamlIdentityProvider) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *SamlIdentityProvider) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *SamlIdentityProvider) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *SamlIdentityProvider) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *SamlIdentityProvider) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *SamlIdentityProvider) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *SamlIdentityProvider) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *SamlIdentityProvider) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetDebugMode

`func (o *SamlIdentityProvider) GetDebugMode() bool`

GetDebugMode returns the DebugMode field if non-nil, zero value otherwise.

### GetDebugModeOk

`func (o *SamlIdentityProvider) GetDebugModeOk() (*bool, bool)`

GetDebugModeOk returns a tuple with the DebugMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugMode

`func (o *SamlIdentityProvider) SetDebugMode(v bool)`

SetDebugMode sets DebugMode field to given value.

### HasDebugMode

`func (o *SamlIdentityProvider) HasDebugMode() bool`

HasDebugMode returns a boolean if a field has been set.

### GetSignAuthnRequest

`func (o *SamlIdentityProvider) GetSignAuthnRequest() bool`

GetSignAuthnRequest returns the SignAuthnRequest field if non-nil, zero value otherwise.

### GetSignAuthnRequestOk

`func (o *SamlIdentityProvider) GetSignAuthnRequestOk() (*bool, bool)`

GetSignAuthnRequestOk returns a tuple with the SignAuthnRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignAuthnRequest

`func (o *SamlIdentityProvider) SetSignAuthnRequest(v bool)`

SetSignAuthnRequest sets SignAuthnRequest field to given value.

### HasSignAuthnRequest

`func (o *SamlIdentityProvider) HasSignAuthnRequest() bool`

HasSignAuthnRequest returns a boolean if a field has been set.

### GetDisableRequestedAuthnContext

`func (o *SamlIdentityProvider) GetDisableRequestedAuthnContext() bool`

GetDisableRequestedAuthnContext returns the DisableRequestedAuthnContext field if non-nil, zero value otherwise.

### GetDisableRequestedAuthnContextOk

`func (o *SamlIdentityProvider) GetDisableRequestedAuthnContextOk() (*bool, bool)`

GetDisableRequestedAuthnContextOk returns a tuple with the DisableRequestedAuthnContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableRequestedAuthnContext

`func (o *SamlIdentityProvider) SetDisableRequestedAuthnContext(v bool)`

SetDisableRequestedAuthnContext sets DisableRequestedAuthnContext field to given value.

### HasDisableRequestedAuthnContext

`func (o *SamlIdentityProvider) HasDisableRequestedAuthnContext() bool`

HasDisableRequestedAuthnContext returns a boolean if a field has been set.

### GetIsRedirectBinding

`func (o *SamlIdentityProvider) GetIsRedirectBinding() bool`

GetIsRedirectBinding returns the IsRedirectBinding field if non-nil, zero value otherwise.

### GetIsRedirectBindingOk

`func (o *SamlIdentityProvider) GetIsRedirectBindingOk() (*bool, bool)`

GetIsRedirectBindingOk returns a tuple with the IsRedirectBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRedirectBinding

`func (o *SamlIdentityProvider) SetIsRedirectBinding(v bool)`

SetIsRedirectBinding sets IsRedirectBinding field to given value.

### HasIsRedirectBinding

`func (o *SamlIdentityProvider) HasIsRedirectBinding() bool`

HasIsRedirectBinding returns a boolean if a field has been set.

### GetCertificate

`func (o *SamlIdentityProvider) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *SamlIdentityProvider) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *SamlIdentityProvider) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetCreatedAt

`func (o *SamlIdentityProvider) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SamlIdentityProvider) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SamlIdentityProvider) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SamlIdentityProvider) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SamlIdentityProvider) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SamlIdentityProvider) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetModifiedAt

`func (o *SamlIdentityProvider) GetModifiedAt() time.Time`

GetModifiedAt returns the ModifiedAt field if non-nil, zero value otherwise.

### GetModifiedAtOk

`func (o *SamlIdentityProvider) GetModifiedAtOk() (*time.Time, bool)`

GetModifiedAtOk returns a tuple with the ModifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedAt

`func (o *SamlIdentityProvider) SetModifiedAt(v time.Time)`

SetModifiedAt sets ModifiedAt field to given value.


### GetModifiedBy

`func (o *SamlIdentityProvider) GetModifiedBy() string`

GetModifiedBy returns the ModifiedBy field if non-nil, zero value otherwise.

### GetModifiedByOk

`func (o *SamlIdentityProvider) GetModifiedByOk() (*string, bool)`

GetModifiedByOk returns a tuple with the ModifiedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedBy

`func (o *SamlIdentityProvider) SetModifiedBy(v string)`

SetModifiedBy sets ModifiedBy field to given value.


### GetId

`func (o *SamlIdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SamlIdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SamlIdentityProvider) SetId(v string)`

SetId sets Id field to given value.


### GetAssertionConsumerUrl

`func (o *SamlIdentityProvider) GetAssertionConsumerUrl() string`

GetAssertionConsumerUrl returns the AssertionConsumerUrl field if non-nil, zero value otherwise.

### GetAssertionConsumerUrlOk

`func (o *SamlIdentityProvider) GetAssertionConsumerUrlOk() (*string, bool)`

GetAssertionConsumerUrlOk returns a tuple with the AssertionConsumerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertionConsumerUrl

`func (o *SamlIdentityProvider) SetAssertionConsumerUrl(v string)`

SetAssertionConsumerUrl sets AssertionConsumerUrl field to given value.

### HasAssertionConsumerUrl

`func (o *SamlIdentityProvider) HasAssertionConsumerUrl() bool`

HasAssertionConsumerUrl returns a boolean if a field has been set.

### GetEntityId

`func (o *SamlIdentityProvider) GetEntityId() string`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *SamlIdentityProvider) GetEntityIdOk() (*string, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *SamlIdentityProvider) SetEntityId(v string)`

SetEntityId sets EntityId field to given value.

### HasEntityId

`func (o *SamlIdentityProvider) HasEntityId() bool`

HasEntityId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



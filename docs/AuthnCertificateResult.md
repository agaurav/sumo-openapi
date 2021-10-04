# AuthnCertificateResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Authentication Request Signing Certificate for the user. | 

## Methods

### NewAuthnCertificateResult

`func NewAuthnCertificateResult(certificate string, ) *AuthnCertificateResult`

NewAuthnCertificateResult instantiates a new AuthnCertificateResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthnCertificateResultWithDefaults

`func NewAuthnCertificateResultWithDefaults() *AuthnCertificateResult`

NewAuthnCertificateResultWithDefaults instantiates a new AuthnCertificateResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *AuthnCertificateResult) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AuthnCertificateResult) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AuthnCertificateResult) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



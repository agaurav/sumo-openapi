# PasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLength** | Pointer to **int32** | The minimum length of the password. | [optional] [default to 8]
**MaxLength** | Pointer to **int32** | The maximum length of the password. (Setting this to any value other than 128 is no longer supported; this field may be deprecated in the future.) | [optional] [default to 128]
**MustContainLowercase** | Pointer to **bool** | If the password must contain lower case characters. | [optional] [default to true]
**MustContainUppercase** | Pointer to **bool** | If the password must contain upper case characters. | [optional] [default to true]
**MustContainDigits** | Pointer to **bool** | If the password must contain digits. | [optional] [default to true]
**MustContainSpecialChars** | Pointer to **bool** | If the password must contain special characters. | [optional] [default to true]
**MaxPasswordAgeInDays** | Pointer to **int32** | Maximum number of days that a password can be used before user is required to change it. Put -1 if the user should not have to change their password. | [optional] [default to 365]
**MinUniquePasswords** | Pointer to **int32** | The minimum number of unique new passwords that a user must use before an old password can be reused. | [optional] [default to 10]
**AccountLockoutThreshold** | Pointer to **int32** | Number of failed login attempts allowed before account is locked-out. | [optional] [default to 6]
**FailedLoginResetDurationInMins** | Pointer to **int32** | The duration of time in minutes that must elapse from the first failed login attempt after which failed login count is reset to 0. | [optional] [default to 10]
**AccountLockoutDurationInMins** | Pointer to **int32** | The duration of time in minutes that a locked-out account remained locked before getting unlocked automatically. | [optional] [default to 30]
**RequireMfa** | Pointer to **bool** | If MFA should be required to log in. By default, this field is set to &#x60;false&#x60;. | [optional] [default to false]
**RememberMfa** | Pointer to **bool** | If MFA should be remembered on the browser. | [optional] [default to true]

## Methods

### NewPasswordPolicy

`func NewPasswordPolicy() *PasswordPolicy`

NewPasswordPolicy instantiates a new PasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyWithDefaults

`func NewPasswordPolicyWithDefaults() *PasswordPolicy`

NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinLength

`func (o *PasswordPolicy) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *PasswordPolicy) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *PasswordPolicy) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *PasswordPolicy) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetMaxLength

`func (o *PasswordPolicy) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PasswordPolicy) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PasswordPolicy) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PasswordPolicy) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMustContainLowercase

`func (o *PasswordPolicy) GetMustContainLowercase() bool`

GetMustContainLowercase returns the MustContainLowercase field if non-nil, zero value otherwise.

### GetMustContainLowercaseOk

`func (o *PasswordPolicy) GetMustContainLowercaseOk() (*bool, bool)`

GetMustContainLowercaseOk returns a tuple with the MustContainLowercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustContainLowercase

`func (o *PasswordPolicy) SetMustContainLowercase(v bool)`

SetMustContainLowercase sets MustContainLowercase field to given value.

### HasMustContainLowercase

`func (o *PasswordPolicy) HasMustContainLowercase() bool`

HasMustContainLowercase returns a boolean if a field has been set.

### GetMustContainUppercase

`func (o *PasswordPolicy) GetMustContainUppercase() bool`

GetMustContainUppercase returns the MustContainUppercase field if non-nil, zero value otherwise.

### GetMustContainUppercaseOk

`func (o *PasswordPolicy) GetMustContainUppercaseOk() (*bool, bool)`

GetMustContainUppercaseOk returns a tuple with the MustContainUppercase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustContainUppercase

`func (o *PasswordPolicy) SetMustContainUppercase(v bool)`

SetMustContainUppercase sets MustContainUppercase field to given value.

### HasMustContainUppercase

`func (o *PasswordPolicy) HasMustContainUppercase() bool`

HasMustContainUppercase returns a boolean if a field has been set.

### GetMustContainDigits

`func (o *PasswordPolicy) GetMustContainDigits() bool`

GetMustContainDigits returns the MustContainDigits field if non-nil, zero value otherwise.

### GetMustContainDigitsOk

`func (o *PasswordPolicy) GetMustContainDigitsOk() (*bool, bool)`

GetMustContainDigitsOk returns a tuple with the MustContainDigits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustContainDigits

`func (o *PasswordPolicy) SetMustContainDigits(v bool)`

SetMustContainDigits sets MustContainDigits field to given value.

### HasMustContainDigits

`func (o *PasswordPolicy) HasMustContainDigits() bool`

HasMustContainDigits returns a boolean if a field has been set.

### GetMustContainSpecialChars

`func (o *PasswordPolicy) GetMustContainSpecialChars() bool`

GetMustContainSpecialChars returns the MustContainSpecialChars field if non-nil, zero value otherwise.

### GetMustContainSpecialCharsOk

`func (o *PasswordPolicy) GetMustContainSpecialCharsOk() (*bool, bool)`

GetMustContainSpecialCharsOk returns a tuple with the MustContainSpecialChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMustContainSpecialChars

`func (o *PasswordPolicy) SetMustContainSpecialChars(v bool)`

SetMustContainSpecialChars sets MustContainSpecialChars field to given value.

### HasMustContainSpecialChars

`func (o *PasswordPolicy) HasMustContainSpecialChars() bool`

HasMustContainSpecialChars returns a boolean if a field has been set.

### GetMaxPasswordAgeInDays

`func (o *PasswordPolicy) GetMaxPasswordAgeInDays() int32`

GetMaxPasswordAgeInDays returns the MaxPasswordAgeInDays field if non-nil, zero value otherwise.

### GetMaxPasswordAgeInDaysOk

`func (o *PasswordPolicy) GetMaxPasswordAgeInDaysOk() (*int32, bool)`

GetMaxPasswordAgeInDaysOk returns a tuple with the MaxPasswordAgeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPasswordAgeInDays

`func (o *PasswordPolicy) SetMaxPasswordAgeInDays(v int32)`

SetMaxPasswordAgeInDays sets MaxPasswordAgeInDays field to given value.

### HasMaxPasswordAgeInDays

`func (o *PasswordPolicy) HasMaxPasswordAgeInDays() bool`

HasMaxPasswordAgeInDays returns a boolean if a field has been set.

### GetMinUniquePasswords

`func (o *PasswordPolicy) GetMinUniquePasswords() int32`

GetMinUniquePasswords returns the MinUniquePasswords field if non-nil, zero value otherwise.

### GetMinUniquePasswordsOk

`func (o *PasswordPolicy) GetMinUniquePasswordsOk() (*int32, bool)`

GetMinUniquePasswordsOk returns a tuple with the MinUniquePasswords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUniquePasswords

`func (o *PasswordPolicy) SetMinUniquePasswords(v int32)`

SetMinUniquePasswords sets MinUniquePasswords field to given value.

### HasMinUniquePasswords

`func (o *PasswordPolicy) HasMinUniquePasswords() bool`

HasMinUniquePasswords returns a boolean if a field has been set.

### GetAccountLockoutThreshold

`func (o *PasswordPolicy) GetAccountLockoutThreshold() int32`

GetAccountLockoutThreshold returns the AccountLockoutThreshold field if non-nil, zero value otherwise.

### GetAccountLockoutThresholdOk

`func (o *PasswordPolicy) GetAccountLockoutThresholdOk() (*int32, bool)`

GetAccountLockoutThresholdOk returns a tuple with the AccountLockoutThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockoutThreshold

`func (o *PasswordPolicy) SetAccountLockoutThreshold(v int32)`

SetAccountLockoutThreshold sets AccountLockoutThreshold field to given value.

### HasAccountLockoutThreshold

`func (o *PasswordPolicy) HasAccountLockoutThreshold() bool`

HasAccountLockoutThreshold returns a boolean if a field has been set.

### GetFailedLoginResetDurationInMins

`func (o *PasswordPolicy) GetFailedLoginResetDurationInMins() int32`

GetFailedLoginResetDurationInMins returns the FailedLoginResetDurationInMins field if non-nil, zero value otherwise.

### GetFailedLoginResetDurationInMinsOk

`func (o *PasswordPolicy) GetFailedLoginResetDurationInMinsOk() (*int32, bool)`

GetFailedLoginResetDurationInMinsOk returns a tuple with the FailedLoginResetDurationInMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedLoginResetDurationInMins

`func (o *PasswordPolicy) SetFailedLoginResetDurationInMins(v int32)`

SetFailedLoginResetDurationInMins sets FailedLoginResetDurationInMins field to given value.

### HasFailedLoginResetDurationInMins

`func (o *PasswordPolicy) HasFailedLoginResetDurationInMins() bool`

HasFailedLoginResetDurationInMins returns a boolean if a field has been set.

### GetAccountLockoutDurationInMins

`func (o *PasswordPolicy) GetAccountLockoutDurationInMins() int32`

GetAccountLockoutDurationInMins returns the AccountLockoutDurationInMins field if non-nil, zero value otherwise.

### GetAccountLockoutDurationInMinsOk

`func (o *PasswordPolicy) GetAccountLockoutDurationInMinsOk() (*int32, bool)`

GetAccountLockoutDurationInMinsOk returns a tuple with the AccountLockoutDurationInMins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountLockoutDurationInMins

`func (o *PasswordPolicy) SetAccountLockoutDurationInMins(v int32)`

SetAccountLockoutDurationInMins sets AccountLockoutDurationInMins field to given value.

### HasAccountLockoutDurationInMins

`func (o *PasswordPolicy) HasAccountLockoutDurationInMins() bool`

HasAccountLockoutDurationInMins returns a boolean if a field has been set.

### GetRequireMfa

`func (o *PasswordPolicy) GetRequireMfa() bool`

GetRequireMfa returns the RequireMfa field if non-nil, zero value otherwise.

### GetRequireMfaOk

`func (o *PasswordPolicy) GetRequireMfaOk() (*bool, bool)`

GetRequireMfaOk returns a tuple with the RequireMfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireMfa

`func (o *PasswordPolicy) SetRequireMfa(v bool)`

SetRequireMfa sets RequireMfa field to given value.

### HasRequireMfa

`func (o *PasswordPolicy) HasRequireMfa() bool`

HasRequireMfa returns a boolean if a field has been set.

### GetRememberMfa

`func (o *PasswordPolicy) GetRememberMfa() bool`

GetRememberMfa returns the RememberMfa field if non-nil, zero value otherwise.

### GetRememberMfaOk

`func (o *PasswordPolicy) GetRememberMfaOk() (*bool, bool)`

GetRememberMfaOk returns a tuple with the RememberMfa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberMfa

`func (o *PasswordPolicy) SetRememberMfa(v bool)`

SetRememberMfa sets RememberMfa field to given value.

### HasRememberMfa

`func (o *PasswordPolicy) HasRememberMfa() bool`

HasRememberMfa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateAccountSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address | [optional] 
**City** | Pointer to **string** | City | [optional] 
**CompanyName** | Pointer to **string** | Company name | [optional] 
**Country** | Pointer to **string** | Country | [optional] 
**DefaultKeyName** | Pointer to **string** | Set the account default key based on the DFC key item name. Use \&quot;set-original-akeyless-default-key\&quot; to revert to using the original default key of the account. Empty string will change nothing. | [optional] 
**DefaultShareLinkTtlMinutes** | Pointer to **string** | Set the default ttl in minutes for sharing item number between 60 and 43200 | [optional] 
**DefaultVersioning** | Pointer to **string** | If set to true, new item version will be created on each update [true/false] | [optional] 
**DpEnableClassicKeyProtection** | Pointer to **string** | Set to update protection with classic keys state [true/false] | [optional] 
**InvalidCharacters** | Pointer to **string** | Characters that cannot be used for items/targets/roles/auths/event_forwarder names. Empty string will enforce nothing. | [optional] [default to "notReceivedInvalidCharacter"]
**ItemType** | Pointer to **string** | VersionSettingsObjectType defines object types for account version settings | [optional] 
**ItemsDeletionProtection** | Pointer to **string** | Set or unset the default behaviour of items deletion protection [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**JwtTtlDefault** | Pointer to **int64** | Default ttl | [optional] 
**JwtTtlMax** | Pointer to **int64** | Maximum ttl | [optional] 
**JwtTtlMin** | Pointer to **int64** | Minimum ttl | [optional] 
**MaxVersions** | Pointer to **string** | Max versions | [optional] 
**PasswordLength** | Pointer to **int64** | Password length between 5 - to 50 characters | [optional] 
**Phone** | Pointer to **string** | Phone number | [optional] 
**PostalCode** | Pointer to **string** | Postal code | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseLowerLetters** | Pointer to **string** | Password must contain lower case letters [true/false] | [optional] 
**UseNumbers** | Pointer to **string** | Password must contain numbers [true/false] | [optional] 
**UseSpecialCharacters** | Pointer to **string** | Password must contain special characters [true/false] | [optional] 
**UseCapitalLetters** | Pointer to **string** | Password must contain capital letters [true/false] | [optional] 

## Methods

### NewUpdateAccountSettings

`func NewUpdateAccountSettings() *UpdateAccountSettings`

NewUpdateAccountSettings instantiates a new UpdateAccountSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountSettingsWithDefaults

`func NewUpdateAccountSettingsWithDefaults() *UpdateAccountSettings`

NewUpdateAccountSettingsWithDefaults instantiates a new UpdateAccountSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UpdateAccountSettings) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateAccountSettings) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateAccountSettings) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateAccountSettings) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetCity

`func (o *UpdateAccountSettings) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateAccountSettings) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateAccountSettings) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateAccountSettings) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCompanyName

`func (o *UpdateAccountSettings) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateAccountSettings) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateAccountSettings) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UpdateAccountSettings) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetCountry

`func (o *UpdateAccountSettings) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateAccountSettings) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateAccountSettings) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UpdateAccountSettings) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDefaultKeyName

`func (o *UpdateAccountSettings) GetDefaultKeyName() string`

GetDefaultKeyName returns the DefaultKeyName field if non-nil, zero value otherwise.

### GetDefaultKeyNameOk

`func (o *UpdateAccountSettings) GetDefaultKeyNameOk() (*string, bool)`

GetDefaultKeyNameOk returns a tuple with the DefaultKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKeyName

`func (o *UpdateAccountSettings) SetDefaultKeyName(v string)`

SetDefaultKeyName sets DefaultKeyName field to given value.

### HasDefaultKeyName

`func (o *UpdateAccountSettings) HasDefaultKeyName() bool`

HasDefaultKeyName returns a boolean if a field has been set.

### GetDefaultShareLinkTtlMinutes

`func (o *UpdateAccountSettings) GetDefaultShareLinkTtlMinutes() string`

GetDefaultShareLinkTtlMinutes returns the DefaultShareLinkTtlMinutes field if non-nil, zero value otherwise.

### GetDefaultShareLinkTtlMinutesOk

`func (o *UpdateAccountSettings) GetDefaultShareLinkTtlMinutesOk() (*string, bool)`

GetDefaultShareLinkTtlMinutesOk returns a tuple with the DefaultShareLinkTtlMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultShareLinkTtlMinutes

`func (o *UpdateAccountSettings) SetDefaultShareLinkTtlMinutes(v string)`

SetDefaultShareLinkTtlMinutes sets DefaultShareLinkTtlMinutes field to given value.

### HasDefaultShareLinkTtlMinutes

`func (o *UpdateAccountSettings) HasDefaultShareLinkTtlMinutes() bool`

HasDefaultShareLinkTtlMinutes returns a boolean if a field has been set.

### GetDefaultVersioning

`func (o *UpdateAccountSettings) GetDefaultVersioning() string`

GetDefaultVersioning returns the DefaultVersioning field if non-nil, zero value otherwise.

### GetDefaultVersioningOk

`func (o *UpdateAccountSettings) GetDefaultVersioningOk() (*string, bool)`

GetDefaultVersioningOk returns a tuple with the DefaultVersioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVersioning

`func (o *UpdateAccountSettings) SetDefaultVersioning(v string)`

SetDefaultVersioning sets DefaultVersioning field to given value.

### HasDefaultVersioning

`func (o *UpdateAccountSettings) HasDefaultVersioning() bool`

HasDefaultVersioning returns a boolean if a field has been set.

### GetDpEnableClassicKeyProtection

`func (o *UpdateAccountSettings) GetDpEnableClassicKeyProtection() string`

GetDpEnableClassicKeyProtection returns the DpEnableClassicKeyProtection field if non-nil, zero value otherwise.

### GetDpEnableClassicKeyProtectionOk

`func (o *UpdateAccountSettings) GetDpEnableClassicKeyProtectionOk() (*string, bool)`

GetDpEnableClassicKeyProtectionOk returns a tuple with the DpEnableClassicKeyProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpEnableClassicKeyProtection

`func (o *UpdateAccountSettings) SetDpEnableClassicKeyProtection(v string)`

SetDpEnableClassicKeyProtection sets DpEnableClassicKeyProtection field to given value.

### HasDpEnableClassicKeyProtection

`func (o *UpdateAccountSettings) HasDpEnableClassicKeyProtection() bool`

HasDpEnableClassicKeyProtection returns a boolean if a field has been set.

### GetInvalidCharacters

`func (o *UpdateAccountSettings) GetInvalidCharacters() string`

GetInvalidCharacters returns the InvalidCharacters field if non-nil, zero value otherwise.

### GetInvalidCharactersOk

`func (o *UpdateAccountSettings) GetInvalidCharactersOk() (*string, bool)`

GetInvalidCharactersOk returns a tuple with the InvalidCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidCharacters

`func (o *UpdateAccountSettings) SetInvalidCharacters(v string)`

SetInvalidCharacters sets InvalidCharacters field to given value.

### HasInvalidCharacters

`func (o *UpdateAccountSettings) HasInvalidCharacters() bool`

HasInvalidCharacters returns a boolean if a field has been set.

### GetItemType

`func (o *UpdateAccountSettings) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *UpdateAccountSettings) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *UpdateAccountSettings) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *UpdateAccountSettings) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### GetItemsDeletionProtection

`func (o *UpdateAccountSettings) GetItemsDeletionProtection() string`

GetItemsDeletionProtection returns the ItemsDeletionProtection field if non-nil, zero value otherwise.

### GetItemsDeletionProtectionOk

`func (o *UpdateAccountSettings) GetItemsDeletionProtectionOk() (*string, bool)`

GetItemsDeletionProtectionOk returns a tuple with the ItemsDeletionProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsDeletionProtection

`func (o *UpdateAccountSettings) SetItemsDeletionProtection(v string)`

SetItemsDeletionProtection sets ItemsDeletionProtection field to given value.

### HasItemsDeletionProtection

`func (o *UpdateAccountSettings) HasItemsDeletionProtection() bool`

HasItemsDeletionProtection returns a boolean if a field has been set.

### GetJson

`func (o *UpdateAccountSettings) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateAccountSettings) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateAccountSettings) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateAccountSettings) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwtTtlDefault

`func (o *UpdateAccountSettings) GetJwtTtlDefault() int64`

GetJwtTtlDefault returns the JwtTtlDefault field if non-nil, zero value otherwise.

### GetJwtTtlDefaultOk

`func (o *UpdateAccountSettings) GetJwtTtlDefaultOk() (*int64, bool)`

GetJwtTtlDefaultOk returns a tuple with the JwtTtlDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtlDefault

`func (o *UpdateAccountSettings) SetJwtTtlDefault(v int64)`

SetJwtTtlDefault sets JwtTtlDefault field to given value.

### HasJwtTtlDefault

`func (o *UpdateAccountSettings) HasJwtTtlDefault() bool`

HasJwtTtlDefault returns a boolean if a field has been set.

### GetJwtTtlMax

`func (o *UpdateAccountSettings) GetJwtTtlMax() int64`

GetJwtTtlMax returns the JwtTtlMax field if non-nil, zero value otherwise.

### GetJwtTtlMaxOk

`func (o *UpdateAccountSettings) GetJwtTtlMaxOk() (*int64, bool)`

GetJwtTtlMaxOk returns a tuple with the JwtTtlMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtlMax

`func (o *UpdateAccountSettings) SetJwtTtlMax(v int64)`

SetJwtTtlMax sets JwtTtlMax field to given value.

### HasJwtTtlMax

`func (o *UpdateAccountSettings) HasJwtTtlMax() bool`

HasJwtTtlMax returns a boolean if a field has been set.

### GetJwtTtlMin

`func (o *UpdateAccountSettings) GetJwtTtlMin() int64`

GetJwtTtlMin returns the JwtTtlMin field if non-nil, zero value otherwise.

### GetJwtTtlMinOk

`func (o *UpdateAccountSettings) GetJwtTtlMinOk() (*int64, bool)`

GetJwtTtlMinOk returns a tuple with the JwtTtlMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtlMin

`func (o *UpdateAccountSettings) SetJwtTtlMin(v int64)`

SetJwtTtlMin sets JwtTtlMin field to given value.

### HasJwtTtlMin

`func (o *UpdateAccountSettings) HasJwtTtlMin() bool`

HasJwtTtlMin returns a boolean if a field has been set.

### GetMaxVersions

`func (o *UpdateAccountSettings) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *UpdateAccountSettings) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *UpdateAccountSettings) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *UpdateAccountSettings) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetPasswordLength

`func (o *UpdateAccountSettings) GetPasswordLength() int64`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *UpdateAccountSettings) GetPasswordLengthOk() (*int64, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *UpdateAccountSettings) SetPasswordLength(v int64)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *UpdateAccountSettings) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateAccountSettings) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateAccountSettings) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateAccountSettings) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateAccountSettings) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetPostalCode

`func (o *UpdateAccountSettings) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UpdateAccountSettings) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UpdateAccountSettings) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UpdateAccountSettings) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetToken

`func (o *UpdateAccountSettings) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateAccountSettings) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateAccountSettings) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateAccountSettings) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateAccountSettings) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateAccountSettings) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateAccountSettings) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateAccountSettings) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseLowerLetters

`func (o *UpdateAccountSettings) GetUseLowerLetters() string`

GetUseLowerLetters returns the UseLowerLetters field if non-nil, zero value otherwise.

### GetUseLowerLettersOk

`func (o *UpdateAccountSettings) GetUseLowerLettersOk() (*string, bool)`

GetUseLowerLettersOk returns a tuple with the UseLowerLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLowerLetters

`func (o *UpdateAccountSettings) SetUseLowerLetters(v string)`

SetUseLowerLetters sets UseLowerLetters field to given value.

### HasUseLowerLetters

`func (o *UpdateAccountSettings) HasUseLowerLetters() bool`

HasUseLowerLetters returns a boolean if a field has been set.

### GetUseNumbers

`func (o *UpdateAccountSettings) GetUseNumbers() string`

GetUseNumbers returns the UseNumbers field if non-nil, zero value otherwise.

### GetUseNumbersOk

`func (o *UpdateAccountSettings) GetUseNumbersOk() (*string, bool)`

GetUseNumbersOk returns a tuple with the UseNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseNumbers

`func (o *UpdateAccountSettings) SetUseNumbers(v string)`

SetUseNumbers sets UseNumbers field to given value.

### HasUseNumbers

`func (o *UpdateAccountSettings) HasUseNumbers() bool`

HasUseNumbers returns a boolean if a field has been set.

### GetUseSpecialCharacters

`func (o *UpdateAccountSettings) GetUseSpecialCharacters() string`

GetUseSpecialCharacters returns the UseSpecialCharacters field if non-nil, zero value otherwise.

### GetUseSpecialCharactersOk

`func (o *UpdateAccountSettings) GetUseSpecialCharactersOk() (*string, bool)`

GetUseSpecialCharactersOk returns a tuple with the UseSpecialCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSpecialCharacters

`func (o *UpdateAccountSettings) SetUseSpecialCharacters(v string)`

SetUseSpecialCharacters sets UseSpecialCharacters field to given value.

### HasUseSpecialCharacters

`func (o *UpdateAccountSettings) HasUseSpecialCharacters() bool`

HasUseSpecialCharacters returns a boolean if a field has been set.

### GetUseCapitalLetters

`func (o *UpdateAccountSettings) GetUseCapitalLetters() string`

GetUseCapitalLetters returns the UseCapitalLetters field if non-nil, zero value otherwise.

### GetUseCapitalLettersOk

`func (o *UpdateAccountSettings) GetUseCapitalLettersOk() (*string, bool)`

GetUseCapitalLettersOk returns a tuple with the UseCapitalLetters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCapitalLetters

`func (o *UpdateAccountSettings) SetUseCapitalLetters(v string)`

SetUseCapitalLetters sets UseCapitalLetters field to given value.

### HasUseCapitalLetters

`func (o *UpdateAccountSettings) HasUseCapitalLetters() bool`

HasUseCapitalLetters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateAccountSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address | [optional] 
**City** | Pointer to **string** | City | [optional] 
**CompanyName** | Pointer to **string** | Company name | [optional] 
**Country** | Pointer to **string** | Country | [optional] 
**DefaultVersioning** | Pointer to **string** | Should create version by default | [optional] 
**ItemType** | Pointer to **string** | VersionSettingsObjectType defines object types for account version settings | [optional] 
**JwtTtlDefault** | Pointer to **int64** | Default ttl | [optional] 
**JwtTtlMax** | Pointer to **int64** | Maximum ttl | [optional] 
**JwtTtlMin** | Pointer to **int64** | Minimum ttl | [optional] 
**MaxVersions** | Pointer to **string** | Max versions | [optional] 
**Phone** | Pointer to **string** | Phone number | [optional] 
**PostalCode** | Pointer to **string** | Postal code | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



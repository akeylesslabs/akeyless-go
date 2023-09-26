# UpdateWindowsTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | SSL CA certificate in base64 encoding generated from a trusted Certificate Authority (CA) | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Domain** | Pointer to **string** | User domain name | [optional] 
**Hostname** | **string** | Server hostname | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Password** | **string** | Privileged user password | 
**Port** | Pointer to **string** | Server WinRM port | [optional] [default to "5986"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 
**UseTls** | Pointer to **string** | Enable/Disable TLS for WinRM over HTTPS [true/false] | [optional] [default to "true"]
**Username** | **string** | Privileged username | 

## Methods

### NewUpdateWindowsTarget

`func NewUpdateWindowsTarget(hostname string, name string, password string, username string, ) *UpdateWindowsTarget`

NewUpdateWindowsTarget instantiates a new UpdateWindowsTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWindowsTargetWithDefaults

`func NewUpdateWindowsTargetWithDefaults() *UpdateWindowsTarget`

NewUpdateWindowsTargetWithDefaults instantiates a new UpdateWindowsTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *UpdateWindowsTarget) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateWindowsTarget) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateWindowsTarget) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *UpdateWindowsTarget) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateWindowsTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWindowsTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWindowsTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWindowsTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *UpdateWindowsTarget) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateWindowsTarget) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateWindowsTarget) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateWindowsTarget) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHostname

`func (o *UpdateWindowsTarget) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *UpdateWindowsTarget) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *UpdateWindowsTarget) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJson

`func (o *UpdateWindowsTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateWindowsTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateWindowsTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateWindowsTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateWindowsTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateWindowsTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateWindowsTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateWindowsTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateWindowsTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateWindowsTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateWindowsTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateWindowsTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateWindowsTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWindowsTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWindowsTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateWindowsTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateWindowsTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateWindowsTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateWindowsTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateWindowsTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateWindowsTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateWindowsTarget) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPort

`func (o *UpdateWindowsTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *UpdateWindowsTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *UpdateWindowsTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *UpdateWindowsTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *UpdateWindowsTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateWindowsTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateWindowsTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateWindowsTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateWindowsTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateWindowsTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateWindowsTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateWindowsTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateWindowsTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateWindowsTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateWindowsTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateWindowsTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.

### GetUseTls

`func (o *UpdateWindowsTarget) GetUseTls() string`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *UpdateWindowsTarget) GetUseTlsOk() (*string, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *UpdateWindowsTarget) SetUseTls(v string)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *UpdateWindowsTarget) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateWindowsTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateWindowsTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateWindowsTarget) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



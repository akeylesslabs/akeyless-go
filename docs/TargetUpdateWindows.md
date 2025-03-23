# TargetUpdateWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | SSL CA certificate in base64 encoding generated from a trusted Certificate Authority (CA) | [optional] 
**ConnectionType** | Pointer to **string** | Type of connection to Windows Server [credentials/parent-target] | [optional] [default to "credentials"]
**Description** | Pointer to **string** | Description of the object | [optional] 
**Domain** | Pointer to **string** | User domain name | [optional] 
**Hostname** | **string** | Server hostname | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**ParentTargetName** | Pointer to **string** | Name of the parent target, relevant only when connection-type is parent-target | [optional] 
**Password** | **string** | Privileged user password | [default to "dummy_value"]
**Port** | Pointer to **string** | Server WinRM port | [optional] [default to "5986"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseTls** | Pointer to **string** | Enable/Disable TLS for WinRM over HTTPS [true/false] | [optional] [default to "true"]
**Username** | **string** | Privileged username | [default to "dummy_value"]

## Methods

### NewTargetUpdateWindows

`func NewTargetUpdateWindows(hostname string, name string, password string, username string, ) *TargetUpdateWindows`

NewTargetUpdateWindows instantiates a new TargetUpdateWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateWindowsWithDefaults

`func NewTargetUpdateWindowsWithDefaults() *TargetUpdateWindows`

NewTargetUpdateWindowsWithDefaults instantiates a new TargetUpdateWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TargetUpdateWindows) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TargetUpdateWindows) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TargetUpdateWindows) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TargetUpdateWindows) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetConnectionType

`func (o *TargetUpdateWindows) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *TargetUpdateWindows) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *TargetUpdateWindows) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *TargetUpdateWindows) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateWindows) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateWindows) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateWindows) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateWindows) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *TargetUpdateWindows) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TargetUpdateWindows) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TargetUpdateWindows) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TargetUpdateWindows) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHostname

`func (o *TargetUpdateWindows) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TargetUpdateWindows) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TargetUpdateWindows) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJson

`func (o *TargetUpdateWindows) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateWindows) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateWindows) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateWindows) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateWindows) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateWindows) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateWindows) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateWindows) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateWindows) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateWindows) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateWindows) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateWindows) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateWindows) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateWindows) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateWindows) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateWindows) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateWindows) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateWindows) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateWindows) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateWindows) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateWindows) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateWindows) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateWindows) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetParentTargetName

`func (o *TargetUpdateWindows) GetParentTargetName() string`

GetParentTargetName returns the ParentTargetName field if non-nil, zero value otherwise.

### GetParentTargetNameOk

`func (o *TargetUpdateWindows) GetParentTargetNameOk() (*string, bool)`

GetParentTargetNameOk returns a tuple with the ParentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTargetName

`func (o *TargetUpdateWindows) SetParentTargetName(v string)`

SetParentTargetName sets ParentTargetName field to given value.

### HasParentTargetName

`func (o *TargetUpdateWindows) HasParentTargetName() bool`

HasParentTargetName returns a boolean if a field has been set.

### GetPassword

`func (o *TargetUpdateWindows) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetUpdateWindows) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetUpdateWindows) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPort

`func (o *TargetUpdateWindows) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetUpdateWindows) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetUpdateWindows) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetUpdateWindows) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateWindows) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateWindows) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateWindows) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateWindows) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateWindows) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateWindows) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateWindows) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateWindows) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseTls

`func (o *TargetUpdateWindows) GetUseTls() string`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *TargetUpdateWindows) GetUseTlsOk() (*string, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *TargetUpdateWindows) SetUseTls(v string)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *TargetUpdateWindows) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUsername

`func (o *TargetUpdateWindows) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetUpdateWindows) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetUpdateWindows) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



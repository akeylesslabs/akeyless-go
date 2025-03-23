# TargetCreateWindows

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | SSL CA certificate in base64 encoding generated from a trusted Certificate Authority (CA) | [optional] 
**ConnectionType** | Pointer to **string** | Type of connection to Windows Server [credentials/parent-target] | [optional] [default to "credentials"]
**Description** | Pointer to **string** | Description of the object | [optional] 
**Domain** | Pointer to **string** | User domain name | [optional] 
**Hostname** | **string** | Server hostname | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**ParentTargetName** | Pointer to **string** | Name of the parent target, relevant only when connection-type is parent-target | [optional] 
**Password** | **string** | Privileged user password | [default to "dummy_value"]
**Port** | Pointer to **string** | Server WinRM port | [optional] [default to "5986"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseTls** | Pointer to **string** | Enable/Disable TLS for WinRM over HTTPS [true/false] | [optional] [default to "true"]
**Username** | **string** | Privileged username | [default to "dummy_value"]

## Methods

### NewTargetCreateWindows

`func NewTargetCreateWindows(hostname string, name string, password string, username string, ) *TargetCreateWindows`

NewTargetCreateWindows instantiates a new TargetCreateWindows object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateWindowsWithDefaults

`func NewTargetCreateWindowsWithDefaults() *TargetCreateWindows`

NewTargetCreateWindowsWithDefaults instantiates a new TargetCreateWindows object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *TargetCreateWindows) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *TargetCreateWindows) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *TargetCreateWindows) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *TargetCreateWindows) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetConnectionType

`func (o *TargetCreateWindows) GetConnectionType() string`

GetConnectionType returns the ConnectionType field if non-nil, zero value otherwise.

### GetConnectionTypeOk

`func (o *TargetCreateWindows) GetConnectionTypeOk() (*string, bool)`

GetConnectionTypeOk returns a tuple with the ConnectionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionType

`func (o *TargetCreateWindows) SetConnectionType(v string)`

SetConnectionType sets ConnectionType field to given value.

### HasConnectionType

`func (o *TargetCreateWindows) HasConnectionType() bool`

HasConnectionType returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateWindows) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateWindows) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateWindows) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateWindows) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDomain

`func (o *TargetCreateWindows) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *TargetCreateWindows) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *TargetCreateWindows) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *TargetCreateWindows) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetHostname

`func (o *TargetCreateWindows) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *TargetCreateWindows) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *TargetCreateWindows) SetHostname(v string)`

SetHostname sets Hostname field to given value.


### GetJson

`func (o *TargetCreateWindows) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateWindows) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateWindows) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateWindows) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateWindows) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateWindows) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateWindows) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateWindows) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateWindows) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateWindows) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateWindows) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateWindows) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateWindows) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateWindows) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateWindows) SetName(v string)`

SetName sets Name field to given value.


### GetParentTargetName

`func (o *TargetCreateWindows) GetParentTargetName() string`

GetParentTargetName returns the ParentTargetName field if non-nil, zero value otherwise.

### GetParentTargetNameOk

`func (o *TargetCreateWindows) GetParentTargetNameOk() (*string, bool)`

GetParentTargetNameOk returns a tuple with the ParentTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentTargetName

`func (o *TargetCreateWindows) SetParentTargetName(v string)`

SetParentTargetName sets ParentTargetName field to given value.

### HasParentTargetName

`func (o *TargetCreateWindows) HasParentTargetName() bool`

HasParentTargetName returns a boolean if a field has been set.

### GetPassword

`func (o *TargetCreateWindows) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetCreateWindows) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetCreateWindows) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPort

`func (o *TargetCreateWindows) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TargetCreateWindows) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TargetCreateWindows) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *TargetCreateWindows) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateWindows) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateWindows) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateWindows) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateWindows) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateWindows) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateWindows) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateWindows) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateWindows) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseTls

`func (o *TargetCreateWindows) GetUseTls() string`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *TargetCreateWindows) GetUseTlsOk() (*string, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *TargetCreateWindows) SetUseTls(v string)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *TargetCreateWindows) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUsername

`func (o *TargetCreateWindows) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetCreateWindows) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetCreateWindows) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# SSHCertificateIssueDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedDomains** | Pointer to **[]string** | Relevant for host certificate | [optional] 
**AllowedUserKeyLengths** | Pointer to **map[string]int64** |  | [optional] 
**AllowedUsers** | Pointer to **[]string** | Relevant for user certificate | [optional] 
**CertType** | Pointer to **int32** |  | [optional] 
**CriticalOptions** | Pointer to **map[string]string** |  | [optional] 
**Extensions** | Pointer to **map[string]string** |  | [optional] 
**Principals** | Pointer to **[]string** |  | [optional] 
**StaticKeyId** | Pointer to **string** | In case it is empty, the key ID will be combination of user identifiers and a random string | [optional] 

## Methods

### NewSSHCertificateIssueDetails

`func NewSSHCertificateIssueDetails() *SSHCertificateIssueDetails`

NewSSHCertificateIssueDetails instantiates a new SSHCertificateIssueDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHCertificateIssueDetailsWithDefaults

`func NewSSHCertificateIssueDetailsWithDefaults() *SSHCertificateIssueDetails`

NewSSHCertificateIssueDetailsWithDefaults instantiates a new SSHCertificateIssueDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedDomains

`func (o *SSHCertificateIssueDetails) GetAllowedDomains() []string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *SSHCertificateIssueDetails) GetAllowedDomainsOk() (*[]string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *SSHCertificateIssueDetails) SetAllowedDomains(v []string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *SSHCertificateIssueDetails) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetAllowedUserKeyLengths

`func (o *SSHCertificateIssueDetails) GetAllowedUserKeyLengths() map[string]int64`

GetAllowedUserKeyLengths returns the AllowedUserKeyLengths field if non-nil, zero value otherwise.

### GetAllowedUserKeyLengthsOk

`func (o *SSHCertificateIssueDetails) GetAllowedUserKeyLengthsOk() (*map[string]int64, bool)`

GetAllowedUserKeyLengthsOk returns a tuple with the AllowedUserKeyLengths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUserKeyLengths

`func (o *SSHCertificateIssueDetails) SetAllowedUserKeyLengths(v map[string]int64)`

SetAllowedUserKeyLengths sets AllowedUserKeyLengths field to given value.

### HasAllowedUserKeyLengths

`func (o *SSHCertificateIssueDetails) HasAllowedUserKeyLengths() bool`

HasAllowedUserKeyLengths returns a boolean if a field has been set.

### GetAllowedUsers

`func (o *SSHCertificateIssueDetails) GetAllowedUsers() []string`

GetAllowedUsers returns the AllowedUsers field if non-nil, zero value otherwise.

### GetAllowedUsersOk

`func (o *SSHCertificateIssueDetails) GetAllowedUsersOk() (*[]string, bool)`

GetAllowedUsersOk returns a tuple with the AllowedUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUsers

`func (o *SSHCertificateIssueDetails) SetAllowedUsers(v []string)`

SetAllowedUsers sets AllowedUsers field to given value.

### HasAllowedUsers

`func (o *SSHCertificateIssueDetails) HasAllowedUsers() bool`

HasAllowedUsers returns a boolean if a field has been set.

### GetCertType

`func (o *SSHCertificateIssueDetails) GetCertType() int32`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *SSHCertificateIssueDetails) GetCertTypeOk() (*int32, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *SSHCertificateIssueDetails) SetCertType(v int32)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *SSHCertificateIssueDetails) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### GetCriticalOptions

`func (o *SSHCertificateIssueDetails) GetCriticalOptions() map[string]string`

GetCriticalOptions returns the CriticalOptions field if non-nil, zero value otherwise.

### GetCriticalOptionsOk

`func (o *SSHCertificateIssueDetails) GetCriticalOptionsOk() (*map[string]string, bool)`

GetCriticalOptionsOk returns a tuple with the CriticalOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalOptions

`func (o *SSHCertificateIssueDetails) SetCriticalOptions(v map[string]string)`

SetCriticalOptions sets CriticalOptions field to given value.

### HasCriticalOptions

`func (o *SSHCertificateIssueDetails) HasCriticalOptions() bool`

HasCriticalOptions returns a boolean if a field has been set.

### GetExtensions

`func (o *SSHCertificateIssueDetails) GetExtensions() map[string]string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *SSHCertificateIssueDetails) GetExtensionsOk() (*map[string]string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *SSHCertificateIssueDetails) SetExtensions(v map[string]string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *SSHCertificateIssueDetails) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetPrincipals

`func (o *SSHCertificateIssueDetails) GetPrincipals() []string`

GetPrincipals returns the Principals field if non-nil, zero value otherwise.

### GetPrincipalsOk

`func (o *SSHCertificateIssueDetails) GetPrincipalsOk() (*[]string, bool)`

GetPrincipalsOk returns a tuple with the Principals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipals

`func (o *SSHCertificateIssueDetails) SetPrincipals(v []string)`

SetPrincipals sets Principals field to given value.

### HasPrincipals

`func (o *SSHCertificateIssueDetails) HasPrincipals() bool`

HasPrincipals returns a boolean if a field has been set.

### GetStaticKeyId

`func (o *SSHCertificateIssueDetails) GetStaticKeyId() string`

GetStaticKeyId returns the StaticKeyId field if non-nil, zero value otherwise.

### GetStaticKeyIdOk

`func (o *SSHCertificateIssueDetails) GetStaticKeyIdOk() (*string, bool)`

GetStaticKeyIdOk returns a tuple with the StaticKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticKeyId

`func (o *SSHCertificateIssueDetails) SetStaticKeyId(v string)`

SetStaticKeyId sets StaticKeyId field to given value.

### HasStaticKeyId

`func (o *SSHCertificateIssueDetails) HasStaticKeyId() bool`

HasStaticKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UpdateLdapTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindDn** | Pointer to **string** |  | [optional] 
**BindDnPassword** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**LdapCaCert** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**Name** | **string** | Target name | 
**NewVersion** | Pointer to **bool** | Deprecated | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**TokenExpiration** | Pointer to **string** |  | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateLdapTargetDetails

`func NewUpdateLdapTargetDetails(name string, ) *UpdateLdapTargetDetails`

NewUpdateLdapTargetDetails instantiates a new UpdateLdapTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLdapTargetDetailsWithDefaults

`func NewUpdateLdapTargetDetailsWithDefaults() *UpdateLdapTargetDetails`

NewUpdateLdapTargetDetailsWithDefaults instantiates a new UpdateLdapTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindDn

`func (o *UpdateLdapTargetDetails) GetBindDn() string`

GetBindDn returns the BindDn field if non-nil, zero value otherwise.

### GetBindDnOk

`func (o *UpdateLdapTargetDetails) GetBindDnOk() (*string, bool)`

GetBindDnOk returns a tuple with the BindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDn

`func (o *UpdateLdapTargetDetails) SetBindDn(v string)`

SetBindDn sets BindDn field to given value.

### HasBindDn

`func (o *UpdateLdapTargetDetails) HasBindDn() bool`

HasBindDn returns a boolean if a field has been set.

### GetBindDnPassword

`func (o *UpdateLdapTargetDetails) GetBindDnPassword() string`

GetBindDnPassword returns the BindDnPassword field if non-nil, zero value otherwise.

### GetBindDnPasswordOk

`func (o *UpdateLdapTargetDetails) GetBindDnPasswordOk() (*string, bool)`

GetBindDnPasswordOk returns a tuple with the BindDnPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDnPassword

`func (o *UpdateLdapTargetDetails) SetBindDnPassword(v string)`

SetBindDnPassword sets BindDnPassword field to given value.

### HasBindDnPassword

`func (o *UpdateLdapTargetDetails) HasBindDnPassword() bool`

HasBindDnPassword returns a boolean if a field has been set.

### GetJson

`func (o *UpdateLdapTargetDetails) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateLdapTargetDetails) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateLdapTargetDetails) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateLdapTargetDetails) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateLdapTargetDetails) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateLdapTargetDetails) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateLdapTargetDetails) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateLdapTargetDetails) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateLdapTargetDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateLdapTargetDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateLdapTargetDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateLdapTargetDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLdapCaCert

`func (o *UpdateLdapTargetDetails) GetLdapCaCert() string`

GetLdapCaCert returns the LdapCaCert field if non-nil, zero value otherwise.

### GetLdapCaCertOk

`func (o *UpdateLdapTargetDetails) GetLdapCaCertOk() (*string, bool)`

GetLdapCaCertOk returns a tuple with the LdapCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCaCert

`func (o *UpdateLdapTargetDetails) SetLdapCaCert(v string)`

SetLdapCaCert sets LdapCaCert field to given value.

### HasLdapCaCert

`func (o *UpdateLdapTargetDetails) HasLdapCaCert() bool`

HasLdapCaCert returns a boolean if a field has been set.

### GetLdapUrl

`func (o *UpdateLdapTargetDetails) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *UpdateLdapTargetDetails) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *UpdateLdapTargetDetails) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *UpdateLdapTargetDetails) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetName

`func (o *UpdateLdapTargetDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLdapTargetDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLdapTargetDetails) SetName(v string)`

SetName sets Name field to given value.


### GetNewVersion

`func (o *UpdateLdapTargetDetails) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateLdapTargetDetails) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateLdapTargetDetails) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateLdapTargetDetails) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetProtectionKey

`func (o *UpdateLdapTargetDetails) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *UpdateLdapTargetDetails) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *UpdateLdapTargetDetails) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *UpdateLdapTargetDetails) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetToken

`func (o *UpdateLdapTargetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateLdapTargetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateLdapTargetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateLdapTargetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenExpiration

`func (o *UpdateLdapTargetDetails) GetTokenExpiration() string`

GetTokenExpiration returns the TokenExpiration field if non-nil, zero value otherwise.

### GetTokenExpirationOk

`func (o *UpdateLdapTargetDetails) GetTokenExpirationOk() (*string, bool)`

GetTokenExpirationOk returns a tuple with the TokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiration

`func (o *UpdateLdapTargetDetails) SetTokenExpiration(v string)`

SetTokenExpiration sets TokenExpiration field to given value.

### HasTokenExpiration

`func (o *UpdateLdapTargetDetails) HasTokenExpiration() bool`

HasTokenExpiration returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateLdapTargetDetails) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateLdapTargetDetails) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateLdapTargetDetails) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateLdapTargetDetails) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



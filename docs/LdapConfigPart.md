# LdapConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LdapAccessId** | Pointer to **string** |  | [optional] 
**LdapAnonymousSearch** | Pointer to **bool** |  | [optional] 
**LdapBindDn** | Pointer to **string** |  | [optional] 
**LdapBindPassword** | Pointer to **string** |  | [optional] 
**LdapCert** | Pointer to **string** |  | [optional] 
**LdapEnable** | Pointer to **bool** |  | [optional] 
**LdapGroupAttr** | Pointer to **string** |  | [optional] 
**LdapGroupDn** | Pointer to **string** |  | [optional] 
**LdapGroupFilter** | Pointer to **string** |  | [optional] 
**LdapPrivateKey** | Pointer to **string** |  | [optional] 
**LdapTokenExpiration** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**LdapUserAttr** | Pointer to **string** |  | [optional] 
**LdapUserDn** | Pointer to **string** |  | [optional] 

## Methods

### NewLdapConfigPart

`func NewLdapConfigPart() *LdapConfigPart`

NewLdapConfigPart instantiates a new LdapConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLdapConfigPartWithDefaults

`func NewLdapConfigPartWithDefaults() *LdapConfigPart`

NewLdapConfigPartWithDefaults instantiates a new LdapConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLdapAccessId

`func (o *LdapConfigPart) GetLdapAccessId() string`

GetLdapAccessId returns the LdapAccessId field if non-nil, zero value otherwise.

### GetLdapAccessIdOk

`func (o *LdapConfigPart) GetLdapAccessIdOk() (*string, bool)`

GetLdapAccessIdOk returns a tuple with the LdapAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAccessId

`func (o *LdapConfigPart) SetLdapAccessId(v string)`

SetLdapAccessId sets LdapAccessId field to given value.

### HasLdapAccessId

`func (o *LdapConfigPart) HasLdapAccessId() bool`

HasLdapAccessId returns a boolean if a field has been set.

### GetLdapAnonymousSearch

`func (o *LdapConfigPart) GetLdapAnonymousSearch() bool`

GetLdapAnonymousSearch returns the LdapAnonymousSearch field if non-nil, zero value otherwise.

### GetLdapAnonymousSearchOk

`func (o *LdapConfigPart) GetLdapAnonymousSearchOk() (*bool, bool)`

GetLdapAnonymousSearchOk returns a tuple with the LdapAnonymousSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAnonymousSearch

`func (o *LdapConfigPart) SetLdapAnonymousSearch(v bool)`

SetLdapAnonymousSearch sets LdapAnonymousSearch field to given value.

### HasLdapAnonymousSearch

`func (o *LdapConfigPart) HasLdapAnonymousSearch() bool`

HasLdapAnonymousSearch returns a boolean if a field has been set.

### GetLdapBindDn

`func (o *LdapConfigPart) GetLdapBindDn() string`

GetLdapBindDn returns the LdapBindDn field if non-nil, zero value otherwise.

### GetLdapBindDnOk

`func (o *LdapConfigPart) GetLdapBindDnOk() (*string, bool)`

GetLdapBindDnOk returns a tuple with the LdapBindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindDn

`func (o *LdapConfigPart) SetLdapBindDn(v string)`

SetLdapBindDn sets LdapBindDn field to given value.

### HasLdapBindDn

`func (o *LdapConfigPart) HasLdapBindDn() bool`

HasLdapBindDn returns a boolean if a field has been set.

### GetLdapBindPassword

`func (o *LdapConfigPart) GetLdapBindPassword() string`

GetLdapBindPassword returns the LdapBindPassword field if non-nil, zero value otherwise.

### GetLdapBindPasswordOk

`func (o *LdapConfigPart) GetLdapBindPasswordOk() (*string, bool)`

GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindPassword

`func (o *LdapConfigPart) SetLdapBindPassword(v string)`

SetLdapBindPassword sets LdapBindPassword field to given value.

### HasLdapBindPassword

`func (o *LdapConfigPart) HasLdapBindPassword() bool`

HasLdapBindPassword returns a boolean if a field has been set.

### GetLdapCert

`func (o *LdapConfigPart) GetLdapCert() string`

GetLdapCert returns the LdapCert field if non-nil, zero value otherwise.

### GetLdapCertOk

`func (o *LdapConfigPart) GetLdapCertOk() (*string, bool)`

GetLdapCertOk returns a tuple with the LdapCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCert

`func (o *LdapConfigPart) SetLdapCert(v string)`

SetLdapCert sets LdapCert field to given value.

### HasLdapCert

`func (o *LdapConfigPart) HasLdapCert() bool`

HasLdapCert returns a boolean if a field has been set.

### GetLdapEnable

`func (o *LdapConfigPart) GetLdapEnable() bool`

GetLdapEnable returns the LdapEnable field if non-nil, zero value otherwise.

### GetLdapEnableOk

`func (o *LdapConfigPart) GetLdapEnableOk() (*bool, bool)`

GetLdapEnableOk returns a tuple with the LdapEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapEnable

`func (o *LdapConfigPart) SetLdapEnable(v bool)`

SetLdapEnable sets LdapEnable field to given value.

### HasLdapEnable

`func (o *LdapConfigPart) HasLdapEnable() bool`

HasLdapEnable returns a boolean if a field has been set.

### GetLdapGroupAttr

`func (o *LdapConfigPart) GetLdapGroupAttr() string`

GetLdapGroupAttr returns the LdapGroupAttr field if non-nil, zero value otherwise.

### GetLdapGroupAttrOk

`func (o *LdapConfigPart) GetLdapGroupAttrOk() (*string, bool)`

GetLdapGroupAttrOk returns a tuple with the LdapGroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAttr

`func (o *LdapConfigPart) SetLdapGroupAttr(v string)`

SetLdapGroupAttr sets LdapGroupAttr field to given value.

### HasLdapGroupAttr

`func (o *LdapConfigPart) HasLdapGroupAttr() bool`

HasLdapGroupAttr returns a boolean if a field has been set.

### GetLdapGroupDn

`func (o *LdapConfigPart) GetLdapGroupDn() string`

GetLdapGroupDn returns the LdapGroupDn field if non-nil, zero value otherwise.

### GetLdapGroupDnOk

`func (o *LdapConfigPart) GetLdapGroupDnOk() (*string, bool)`

GetLdapGroupDnOk returns a tuple with the LdapGroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupDn

`func (o *LdapConfigPart) SetLdapGroupDn(v string)`

SetLdapGroupDn sets LdapGroupDn field to given value.

### HasLdapGroupDn

`func (o *LdapConfigPart) HasLdapGroupDn() bool`

HasLdapGroupDn returns a boolean if a field has been set.

### GetLdapGroupFilter

`func (o *LdapConfigPart) GetLdapGroupFilter() string`

GetLdapGroupFilter returns the LdapGroupFilter field if non-nil, zero value otherwise.

### GetLdapGroupFilterOk

`func (o *LdapConfigPart) GetLdapGroupFilterOk() (*string, bool)`

GetLdapGroupFilterOk returns a tuple with the LdapGroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupFilter

`func (o *LdapConfigPart) SetLdapGroupFilter(v string)`

SetLdapGroupFilter sets LdapGroupFilter field to given value.

### HasLdapGroupFilter

`func (o *LdapConfigPart) HasLdapGroupFilter() bool`

HasLdapGroupFilter returns a boolean if a field has been set.

### GetLdapPrivateKey

`func (o *LdapConfigPart) GetLdapPrivateKey() string`

GetLdapPrivateKey returns the LdapPrivateKey field if non-nil, zero value otherwise.

### GetLdapPrivateKeyOk

`func (o *LdapConfigPart) GetLdapPrivateKeyOk() (*string, bool)`

GetLdapPrivateKeyOk returns a tuple with the LdapPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPrivateKey

`func (o *LdapConfigPart) SetLdapPrivateKey(v string)`

SetLdapPrivateKey sets LdapPrivateKey field to given value.

### HasLdapPrivateKey

`func (o *LdapConfigPart) HasLdapPrivateKey() bool`

HasLdapPrivateKey returns a boolean if a field has been set.

### GetLdapTokenExpiration

`func (o *LdapConfigPart) GetLdapTokenExpiration() string`

GetLdapTokenExpiration returns the LdapTokenExpiration field if non-nil, zero value otherwise.

### GetLdapTokenExpirationOk

`func (o *LdapConfigPart) GetLdapTokenExpirationOk() (*string, bool)`

GetLdapTokenExpirationOk returns a tuple with the LdapTokenExpiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapTokenExpiration

`func (o *LdapConfigPart) SetLdapTokenExpiration(v string)`

SetLdapTokenExpiration sets LdapTokenExpiration field to given value.

### HasLdapTokenExpiration

`func (o *LdapConfigPart) HasLdapTokenExpiration() bool`

HasLdapTokenExpiration returns a boolean if a field has been set.

### GetLdapUrl

`func (o *LdapConfigPart) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *LdapConfigPart) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *LdapConfigPart) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *LdapConfigPart) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetLdapUserAttr

`func (o *LdapConfigPart) GetLdapUserAttr() string`

GetLdapUserAttr returns the LdapUserAttr field if non-nil, zero value otherwise.

### GetLdapUserAttrOk

`func (o *LdapConfigPart) GetLdapUserAttrOk() (*string, bool)`

GetLdapUserAttrOk returns a tuple with the LdapUserAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserAttr

`func (o *LdapConfigPart) SetLdapUserAttr(v string)`

SetLdapUserAttr sets LdapUserAttr field to given value.

### HasLdapUserAttr

`func (o *LdapConfigPart) HasLdapUserAttr() bool`

HasLdapUserAttr returns a boolean if a field has been set.

### GetLdapUserDn

`func (o *LdapConfigPart) GetLdapUserDn() string`

GetLdapUserDn returns the LdapUserDn field if non-nil, zero value otherwise.

### GetLdapUserDnOk

`func (o *LdapConfigPart) GetLdapUserDnOk() (*string, bool)`

GetLdapUserDnOk returns a tuple with the LdapUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDn

`func (o *LdapConfigPart) SetLdapUserDn(v string)`

SetLdapUserDn sets LdapUserDn field to given value.

### HasLdapUserDn

`func (o *LdapConfigPart) HasLdapUserDn() bool`

HasLdapUserDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



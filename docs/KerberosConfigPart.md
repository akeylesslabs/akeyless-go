# KerberosConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KerberosAccessId** | Pointer to **string** |  | [optional] 
**KerberosKeytab** | Pointer to **string** |  | [optional] 
**KerberosKrb5Conf** | Pointer to **string** |  | [optional] 
**KerberosPrivateKey** | Pointer to **string** |  | [optional] 
**LdapAnonymousSearch** | Pointer to **bool** |  | [optional] 
**LdapBindDn** | Pointer to **string** |  | [optional] 
**LdapBindPassword** | Pointer to **string** |  | [optional] 
**LdapCert** | Pointer to **string** |  | [optional] 
**LdapGroupAttr** | Pointer to **string** |  | [optional] 
**LdapGroupDn** | Pointer to **string** |  | [optional] 
**LdapGroupFilter** | Pointer to **string** |  | [optional] 
**LdapUrl** | Pointer to **string** |  | [optional] 
**LdapUserAttr** | Pointer to **string** |  | [optional] 
**LdapUserDn** | Pointer to **string** |  | [optional] 

## Methods

### NewKerberosConfigPart

`func NewKerberosConfigPart() *KerberosConfigPart`

NewKerberosConfigPart instantiates a new KerberosConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosConfigPartWithDefaults

`func NewKerberosConfigPartWithDefaults() *KerberosConfigPart`

NewKerberosConfigPartWithDefaults instantiates a new KerberosConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKerberosAccessId

`func (o *KerberosConfigPart) GetKerberosAccessId() string`

GetKerberosAccessId returns the KerberosAccessId field if non-nil, zero value otherwise.

### GetKerberosAccessIdOk

`func (o *KerberosConfigPart) GetKerberosAccessIdOk() (*string, bool)`

GetKerberosAccessIdOk returns a tuple with the KerberosAccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosAccessId

`func (o *KerberosConfigPart) SetKerberosAccessId(v string)`

SetKerberosAccessId sets KerberosAccessId field to given value.

### HasKerberosAccessId

`func (o *KerberosConfigPart) HasKerberosAccessId() bool`

HasKerberosAccessId returns a boolean if a field has been set.

### GetKerberosKeytab

`func (o *KerberosConfigPart) GetKerberosKeytab() string`

GetKerberosKeytab returns the KerberosKeytab field if non-nil, zero value otherwise.

### GetKerberosKeytabOk

`func (o *KerberosConfigPart) GetKerberosKeytabOk() (*string, bool)`

GetKerberosKeytabOk returns a tuple with the KerberosKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeytab

`func (o *KerberosConfigPart) SetKerberosKeytab(v string)`

SetKerberosKeytab sets KerberosKeytab field to given value.

### HasKerberosKeytab

`func (o *KerberosConfigPart) HasKerberosKeytab() bool`

HasKerberosKeytab returns a boolean if a field has been set.

### GetKerberosKrb5Conf

`func (o *KerberosConfigPart) GetKerberosKrb5Conf() string`

GetKerberosKrb5Conf returns the KerberosKrb5Conf field if non-nil, zero value otherwise.

### GetKerberosKrb5ConfOk

`func (o *KerberosConfigPart) GetKerberosKrb5ConfOk() (*string, bool)`

GetKerberosKrb5ConfOk returns a tuple with the KerberosKrb5Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKrb5Conf

`func (o *KerberosConfigPart) SetKerberosKrb5Conf(v string)`

SetKerberosKrb5Conf sets KerberosKrb5Conf field to given value.

### HasKerberosKrb5Conf

`func (o *KerberosConfigPart) HasKerberosKrb5Conf() bool`

HasKerberosKrb5Conf returns a boolean if a field has been set.

### GetKerberosPrivateKey

`func (o *KerberosConfigPart) GetKerberosPrivateKey() string`

GetKerberosPrivateKey returns the KerberosPrivateKey field if non-nil, zero value otherwise.

### GetKerberosPrivateKeyOk

`func (o *KerberosConfigPart) GetKerberosPrivateKeyOk() (*string, bool)`

GetKerberosPrivateKeyOk returns a tuple with the KerberosPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosPrivateKey

`func (o *KerberosConfigPart) SetKerberosPrivateKey(v string)`

SetKerberosPrivateKey sets KerberosPrivateKey field to given value.

### HasKerberosPrivateKey

`func (o *KerberosConfigPart) HasKerberosPrivateKey() bool`

HasKerberosPrivateKey returns a boolean if a field has been set.

### GetLdapAnonymousSearch

`func (o *KerberosConfigPart) GetLdapAnonymousSearch() bool`

GetLdapAnonymousSearch returns the LdapAnonymousSearch field if non-nil, zero value otherwise.

### GetLdapAnonymousSearchOk

`func (o *KerberosConfigPart) GetLdapAnonymousSearchOk() (*bool, bool)`

GetLdapAnonymousSearchOk returns a tuple with the LdapAnonymousSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAnonymousSearch

`func (o *KerberosConfigPart) SetLdapAnonymousSearch(v bool)`

SetLdapAnonymousSearch sets LdapAnonymousSearch field to given value.

### HasLdapAnonymousSearch

`func (o *KerberosConfigPart) HasLdapAnonymousSearch() bool`

HasLdapAnonymousSearch returns a boolean if a field has been set.

### GetLdapBindDn

`func (o *KerberosConfigPart) GetLdapBindDn() string`

GetLdapBindDn returns the LdapBindDn field if non-nil, zero value otherwise.

### GetLdapBindDnOk

`func (o *KerberosConfigPart) GetLdapBindDnOk() (*string, bool)`

GetLdapBindDnOk returns a tuple with the LdapBindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindDn

`func (o *KerberosConfigPart) SetLdapBindDn(v string)`

SetLdapBindDn sets LdapBindDn field to given value.

### HasLdapBindDn

`func (o *KerberosConfigPart) HasLdapBindDn() bool`

HasLdapBindDn returns a boolean if a field has been set.

### GetLdapBindPassword

`func (o *KerberosConfigPart) GetLdapBindPassword() string`

GetLdapBindPassword returns the LdapBindPassword field if non-nil, zero value otherwise.

### GetLdapBindPasswordOk

`func (o *KerberosConfigPart) GetLdapBindPasswordOk() (*string, bool)`

GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindPassword

`func (o *KerberosConfigPart) SetLdapBindPassword(v string)`

SetLdapBindPassword sets LdapBindPassword field to given value.

### HasLdapBindPassword

`func (o *KerberosConfigPart) HasLdapBindPassword() bool`

HasLdapBindPassword returns a boolean if a field has been set.

### GetLdapCert

`func (o *KerberosConfigPart) GetLdapCert() string`

GetLdapCert returns the LdapCert field if non-nil, zero value otherwise.

### GetLdapCertOk

`func (o *KerberosConfigPart) GetLdapCertOk() (*string, bool)`

GetLdapCertOk returns a tuple with the LdapCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCert

`func (o *KerberosConfigPart) SetLdapCert(v string)`

SetLdapCert sets LdapCert field to given value.

### HasLdapCert

`func (o *KerberosConfigPart) HasLdapCert() bool`

HasLdapCert returns a boolean if a field has been set.

### GetLdapGroupAttr

`func (o *KerberosConfigPart) GetLdapGroupAttr() string`

GetLdapGroupAttr returns the LdapGroupAttr field if non-nil, zero value otherwise.

### GetLdapGroupAttrOk

`func (o *KerberosConfigPart) GetLdapGroupAttrOk() (*string, bool)`

GetLdapGroupAttrOk returns a tuple with the LdapGroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAttr

`func (o *KerberosConfigPart) SetLdapGroupAttr(v string)`

SetLdapGroupAttr sets LdapGroupAttr field to given value.

### HasLdapGroupAttr

`func (o *KerberosConfigPart) HasLdapGroupAttr() bool`

HasLdapGroupAttr returns a boolean if a field has been set.

### GetLdapGroupDn

`func (o *KerberosConfigPart) GetLdapGroupDn() string`

GetLdapGroupDn returns the LdapGroupDn field if non-nil, zero value otherwise.

### GetLdapGroupDnOk

`func (o *KerberosConfigPart) GetLdapGroupDnOk() (*string, bool)`

GetLdapGroupDnOk returns a tuple with the LdapGroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupDn

`func (o *KerberosConfigPart) SetLdapGroupDn(v string)`

SetLdapGroupDn sets LdapGroupDn field to given value.

### HasLdapGroupDn

`func (o *KerberosConfigPart) HasLdapGroupDn() bool`

HasLdapGroupDn returns a boolean if a field has been set.

### GetLdapGroupFilter

`func (o *KerberosConfigPart) GetLdapGroupFilter() string`

GetLdapGroupFilter returns the LdapGroupFilter field if non-nil, zero value otherwise.

### GetLdapGroupFilterOk

`func (o *KerberosConfigPart) GetLdapGroupFilterOk() (*string, bool)`

GetLdapGroupFilterOk returns a tuple with the LdapGroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupFilter

`func (o *KerberosConfigPart) SetLdapGroupFilter(v string)`

SetLdapGroupFilter sets LdapGroupFilter field to given value.

### HasLdapGroupFilter

`func (o *KerberosConfigPart) HasLdapGroupFilter() bool`

HasLdapGroupFilter returns a boolean if a field has been set.

### GetLdapUrl

`func (o *KerberosConfigPart) GetLdapUrl() string`

GetLdapUrl returns the LdapUrl field if non-nil, zero value otherwise.

### GetLdapUrlOk

`func (o *KerberosConfigPart) GetLdapUrlOk() (*string, bool)`

GetLdapUrlOk returns a tuple with the LdapUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrl

`func (o *KerberosConfigPart) SetLdapUrl(v string)`

SetLdapUrl sets LdapUrl field to given value.

### HasLdapUrl

`func (o *KerberosConfigPart) HasLdapUrl() bool`

HasLdapUrl returns a boolean if a field has been set.

### GetLdapUserAttr

`func (o *KerberosConfigPart) GetLdapUserAttr() string`

GetLdapUserAttr returns the LdapUserAttr field if non-nil, zero value otherwise.

### GetLdapUserAttrOk

`func (o *KerberosConfigPart) GetLdapUserAttrOk() (*string, bool)`

GetLdapUserAttrOk returns a tuple with the LdapUserAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserAttr

`func (o *KerberosConfigPart) SetLdapUserAttr(v string)`

SetLdapUserAttr sets LdapUserAttr field to given value.

### HasLdapUserAttr

`func (o *KerberosConfigPart) HasLdapUserAttr() bool`

HasLdapUserAttr returns a boolean if a field has been set.

### GetLdapUserDn

`func (o *KerberosConfigPart) GetLdapUserDn() string`

GetLdapUserDn returns the LdapUserDn field if non-nil, zero value otherwise.

### GetLdapUserDnOk

`func (o *KerberosConfigPart) GetLdapUserDnOk() (*string, bool)`

GetLdapUserDnOk returns a tuple with the LdapUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDn

`func (o *KerberosConfigPart) SetLdapUserDn(v string)`

SetLdapUserDn sets LdapUserDn field to given value.

### HasLdapUserDn

`func (o *KerberosConfigPart) HasLdapUserDn() bool`

HasLdapUserDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



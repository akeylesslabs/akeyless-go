# KerberosAuthMethodInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KerberosKeytab** | Pointer to **string** |  | [optional] 
**KerberosKrb5Conf** | Pointer to **string** |  | [optional] 
**LdapAnonymousSearch** | Pointer to **bool** |  | [optional] 
**LdapBindDn** | Pointer to **string** |  | [optional] 
**LdapBindPassword** | Pointer to **string** |  | [optional] 
**LdapCertificate** | Pointer to **string** |  | [optional] 
**LdapGroupAttr** | Pointer to **string** |  | [optional] 
**LdapGroupDn** | Pointer to **string** |  | [optional] 
**LdapGroupFilter** | Pointer to **string** |  | [optional] 
**LdapUrlAddress** | Pointer to **string** |  | [optional] 
**LdapUserAttr** | Pointer to **string** |  | [optional] 
**LdapUserDn** | Pointer to **string** |  | [optional] 

## Methods

### NewKerberosAuthMethodInfo

`func NewKerberosAuthMethodInfo() *KerberosAuthMethodInfo`

NewKerberosAuthMethodInfo instantiates a new KerberosAuthMethodInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKerberosAuthMethodInfoWithDefaults

`func NewKerberosAuthMethodInfoWithDefaults() *KerberosAuthMethodInfo`

NewKerberosAuthMethodInfoWithDefaults instantiates a new KerberosAuthMethodInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKerberosKeytab

`func (o *KerberosAuthMethodInfo) GetKerberosKeytab() string`

GetKerberosKeytab returns the KerberosKeytab field if non-nil, zero value otherwise.

### GetKerberosKeytabOk

`func (o *KerberosAuthMethodInfo) GetKerberosKeytabOk() (*string, bool)`

GetKerberosKeytabOk returns a tuple with the KerberosKeytab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKeytab

`func (o *KerberosAuthMethodInfo) SetKerberosKeytab(v string)`

SetKerberosKeytab sets KerberosKeytab field to given value.

### HasKerberosKeytab

`func (o *KerberosAuthMethodInfo) HasKerberosKeytab() bool`

HasKerberosKeytab returns a boolean if a field has been set.

### GetKerberosKrb5Conf

`func (o *KerberosAuthMethodInfo) GetKerberosKrb5Conf() string`

GetKerberosKrb5Conf returns the KerberosKrb5Conf field if non-nil, zero value otherwise.

### GetKerberosKrb5ConfOk

`func (o *KerberosAuthMethodInfo) GetKerberosKrb5ConfOk() (*string, bool)`

GetKerberosKrb5ConfOk returns a tuple with the KerberosKrb5Conf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosKrb5Conf

`func (o *KerberosAuthMethodInfo) SetKerberosKrb5Conf(v string)`

SetKerberosKrb5Conf sets KerberosKrb5Conf field to given value.

### HasKerberosKrb5Conf

`func (o *KerberosAuthMethodInfo) HasKerberosKrb5Conf() bool`

HasKerberosKrb5Conf returns a boolean if a field has been set.

### GetLdapAnonymousSearch

`func (o *KerberosAuthMethodInfo) GetLdapAnonymousSearch() bool`

GetLdapAnonymousSearch returns the LdapAnonymousSearch field if non-nil, zero value otherwise.

### GetLdapAnonymousSearchOk

`func (o *KerberosAuthMethodInfo) GetLdapAnonymousSearchOk() (*bool, bool)`

GetLdapAnonymousSearchOk returns a tuple with the LdapAnonymousSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAnonymousSearch

`func (o *KerberosAuthMethodInfo) SetLdapAnonymousSearch(v bool)`

SetLdapAnonymousSearch sets LdapAnonymousSearch field to given value.

### HasLdapAnonymousSearch

`func (o *KerberosAuthMethodInfo) HasLdapAnonymousSearch() bool`

HasLdapAnonymousSearch returns a boolean if a field has been set.

### GetLdapBindDn

`func (o *KerberosAuthMethodInfo) GetLdapBindDn() string`

GetLdapBindDn returns the LdapBindDn field if non-nil, zero value otherwise.

### GetLdapBindDnOk

`func (o *KerberosAuthMethodInfo) GetLdapBindDnOk() (*string, bool)`

GetLdapBindDnOk returns a tuple with the LdapBindDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindDn

`func (o *KerberosAuthMethodInfo) SetLdapBindDn(v string)`

SetLdapBindDn sets LdapBindDn field to given value.

### HasLdapBindDn

`func (o *KerberosAuthMethodInfo) HasLdapBindDn() bool`

HasLdapBindDn returns a boolean if a field has been set.

### GetLdapBindPassword

`func (o *KerberosAuthMethodInfo) GetLdapBindPassword() string`

GetLdapBindPassword returns the LdapBindPassword field if non-nil, zero value otherwise.

### GetLdapBindPasswordOk

`func (o *KerberosAuthMethodInfo) GetLdapBindPasswordOk() (*string, bool)`

GetLdapBindPasswordOk returns a tuple with the LdapBindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapBindPassword

`func (o *KerberosAuthMethodInfo) SetLdapBindPassword(v string)`

SetLdapBindPassword sets LdapBindPassword field to given value.

### HasLdapBindPassword

`func (o *KerberosAuthMethodInfo) HasLdapBindPassword() bool`

HasLdapBindPassword returns a boolean if a field has been set.

### GetLdapCertificate

`func (o *KerberosAuthMethodInfo) GetLdapCertificate() string`

GetLdapCertificate returns the LdapCertificate field if non-nil, zero value otherwise.

### GetLdapCertificateOk

`func (o *KerberosAuthMethodInfo) GetLdapCertificateOk() (*string, bool)`

GetLdapCertificateOk returns a tuple with the LdapCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapCertificate

`func (o *KerberosAuthMethodInfo) SetLdapCertificate(v string)`

SetLdapCertificate sets LdapCertificate field to given value.

### HasLdapCertificate

`func (o *KerberosAuthMethodInfo) HasLdapCertificate() bool`

HasLdapCertificate returns a boolean if a field has been set.

### GetLdapGroupAttr

`func (o *KerberosAuthMethodInfo) GetLdapGroupAttr() string`

GetLdapGroupAttr returns the LdapGroupAttr field if non-nil, zero value otherwise.

### GetLdapGroupAttrOk

`func (o *KerberosAuthMethodInfo) GetLdapGroupAttrOk() (*string, bool)`

GetLdapGroupAttrOk returns a tuple with the LdapGroupAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupAttr

`func (o *KerberosAuthMethodInfo) SetLdapGroupAttr(v string)`

SetLdapGroupAttr sets LdapGroupAttr field to given value.

### HasLdapGroupAttr

`func (o *KerberosAuthMethodInfo) HasLdapGroupAttr() bool`

HasLdapGroupAttr returns a boolean if a field has been set.

### GetLdapGroupDn

`func (o *KerberosAuthMethodInfo) GetLdapGroupDn() string`

GetLdapGroupDn returns the LdapGroupDn field if non-nil, zero value otherwise.

### GetLdapGroupDnOk

`func (o *KerberosAuthMethodInfo) GetLdapGroupDnOk() (*string, bool)`

GetLdapGroupDnOk returns a tuple with the LdapGroupDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupDn

`func (o *KerberosAuthMethodInfo) SetLdapGroupDn(v string)`

SetLdapGroupDn sets LdapGroupDn field to given value.

### HasLdapGroupDn

`func (o *KerberosAuthMethodInfo) HasLdapGroupDn() bool`

HasLdapGroupDn returns a boolean if a field has been set.

### GetLdapGroupFilter

`func (o *KerberosAuthMethodInfo) GetLdapGroupFilter() string`

GetLdapGroupFilter returns the LdapGroupFilter field if non-nil, zero value otherwise.

### GetLdapGroupFilterOk

`func (o *KerberosAuthMethodInfo) GetLdapGroupFilterOk() (*string, bool)`

GetLdapGroupFilterOk returns a tuple with the LdapGroupFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapGroupFilter

`func (o *KerberosAuthMethodInfo) SetLdapGroupFilter(v string)`

SetLdapGroupFilter sets LdapGroupFilter field to given value.

### HasLdapGroupFilter

`func (o *KerberosAuthMethodInfo) HasLdapGroupFilter() bool`

HasLdapGroupFilter returns a boolean if a field has been set.

### GetLdapUrlAddress

`func (o *KerberosAuthMethodInfo) GetLdapUrlAddress() string`

GetLdapUrlAddress returns the LdapUrlAddress field if non-nil, zero value otherwise.

### GetLdapUrlAddressOk

`func (o *KerberosAuthMethodInfo) GetLdapUrlAddressOk() (*string, bool)`

GetLdapUrlAddressOk returns a tuple with the LdapUrlAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUrlAddress

`func (o *KerberosAuthMethodInfo) SetLdapUrlAddress(v string)`

SetLdapUrlAddress sets LdapUrlAddress field to given value.

### HasLdapUrlAddress

`func (o *KerberosAuthMethodInfo) HasLdapUrlAddress() bool`

HasLdapUrlAddress returns a boolean if a field has been set.

### GetLdapUserAttr

`func (o *KerberosAuthMethodInfo) GetLdapUserAttr() string`

GetLdapUserAttr returns the LdapUserAttr field if non-nil, zero value otherwise.

### GetLdapUserAttrOk

`func (o *KerberosAuthMethodInfo) GetLdapUserAttrOk() (*string, bool)`

GetLdapUserAttrOk returns a tuple with the LdapUserAttr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserAttr

`func (o *KerberosAuthMethodInfo) SetLdapUserAttr(v string)`

SetLdapUserAttr sets LdapUserAttr field to given value.

### HasLdapUserAttr

`func (o *KerberosAuthMethodInfo) HasLdapUserAttr() bool`

HasLdapUserAttr returns a boolean if a field has been set.

### GetLdapUserDn

`func (o *KerberosAuthMethodInfo) GetLdapUserDn() string`

GetLdapUserDn returns the LdapUserDn field if non-nil, zero value otherwise.

### GetLdapUserDnOk

`func (o *KerberosAuthMethodInfo) GetLdapUserDnOk() (*string, bool)`

GetLdapUserDnOk returns a tuple with the LdapUserDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUserDn

`func (o *KerberosAuthMethodInfo) SetLdapUserDn(v string)`

SetLdapUserDn sets LdapUserDn field to given value.

### HasLdapUserDn

`func (o *KerberosAuthMethodInfo) HasLdapUserDn() bool`

HasLdapUserDn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



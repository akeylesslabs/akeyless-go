# ItemGeneralInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIssueDetails** | Pointer to [**CertificateIssueInfo**](CertificateIssueInfo.md) |  | [optional] 
**CertificateChainInfo** | Pointer to [**CertificateChainInfo**](CertificateChainInfo.md) |  | [optional] 
**CertificateFormat** | Pointer to **string** |  | [optional] 
**CertificatesTemplateInfo** | Pointer to [**CertificateTemplateInfo**](CertificateTemplateInfo.md) |  | [optional] 
**ClassicKeyDetails** | Pointer to [**ClassicKeyDetailsInfo**](ClassicKeyDetailsInfo.md) |  | [optional] 
**ClusterGwUrl** | Pointer to **string** |  | [optional] 
**DisplayMetadata** | Pointer to **string** |  | [optional] 
**DynamicSecretProducerDetails** | Pointer to [**DynamicSecretProducerInfo**](DynamicSecretProducerInfo.md) |  | [optional] 
**ExpirationEvents** | Pointer to [**[]CertificateExpirationEvent**](CertificateExpirationEvent.md) |  | [optional] 
**ImporterInfo** | Pointer to [**ImporterInfo**](ImporterInfo.md) |  | [optional] 
**IssuerOverviewInfo** | Pointer to [**IssuerOverviewInfo**](IssuerOverviewInfo.md) |  | [optional] 
**NextRotationEvents** | Pointer to [**[]NextAutoRotationEvent**](NextAutoRotationEvent.md) |  | [optional] 
**OidcClientInfo** | Pointer to [**OidcClientInfo**](OidcClientInfo.md) |  | [optional] 
**PasswordPolicy** | Pointer to [**PasswordPolicyInfo**](PasswordPolicyInfo.md) |  | [optional] 
**RotatedSecretDetails** | Pointer to [**RotatedSecretDetailsInfo**](RotatedSecretDetailsInfo.md) |  | [optional] 
**SecureRemoteAccessDetails** | Pointer to [**SecureRemoteAccess**](SecureRemoteAccess.md) |  | [optional] 
**StaticSecretInfo** | Pointer to [**StaticSecretDetailsInfo**](StaticSecretDetailsInfo.md) |  | [optional] 
**TokenizerInfo** | Pointer to [**TokenizerInfo**](TokenizerInfo.md) |  | [optional] 

## Methods

### NewItemGeneralInfo

`func NewItemGeneralInfo() *ItemGeneralInfo`

NewItemGeneralInfo instantiates a new ItemGeneralInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemGeneralInfoWithDefaults

`func NewItemGeneralInfoWithDefaults() *ItemGeneralInfo`

NewItemGeneralInfoWithDefaults instantiates a new ItemGeneralInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIssueDetails

`func (o *ItemGeneralInfo) GetCertIssueDetails() CertificateIssueInfo`

GetCertIssueDetails returns the CertIssueDetails field if non-nil, zero value otherwise.

### GetCertIssueDetailsOk

`func (o *ItemGeneralInfo) GetCertIssueDetailsOk() (*CertificateIssueInfo, bool)`

GetCertIssueDetailsOk returns a tuple with the CertIssueDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssueDetails

`func (o *ItemGeneralInfo) SetCertIssueDetails(v CertificateIssueInfo)`

SetCertIssueDetails sets CertIssueDetails field to given value.

### HasCertIssueDetails

`func (o *ItemGeneralInfo) HasCertIssueDetails() bool`

HasCertIssueDetails returns a boolean if a field has been set.

### GetCertificateChainInfo

`func (o *ItemGeneralInfo) GetCertificateChainInfo() CertificateChainInfo`

GetCertificateChainInfo returns the CertificateChainInfo field if non-nil, zero value otherwise.

### GetCertificateChainInfoOk

`func (o *ItemGeneralInfo) GetCertificateChainInfoOk() (*CertificateChainInfo, bool)`

GetCertificateChainInfoOk returns a tuple with the CertificateChainInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateChainInfo

`func (o *ItemGeneralInfo) SetCertificateChainInfo(v CertificateChainInfo)`

SetCertificateChainInfo sets CertificateChainInfo field to given value.

### HasCertificateChainInfo

`func (o *ItemGeneralInfo) HasCertificateChainInfo() bool`

HasCertificateChainInfo returns a boolean if a field has been set.

### GetCertificateFormat

`func (o *ItemGeneralInfo) GetCertificateFormat() string`

GetCertificateFormat returns the CertificateFormat field if non-nil, zero value otherwise.

### GetCertificateFormatOk

`func (o *ItemGeneralInfo) GetCertificateFormatOk() (*string, bool)`

GetCertificateFormatOk returns a tuple with the CertificateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFormat

`func (o *ItemGeneralInfo) SetCertificateFormat(v string)`

SetCertificateFormat sets CertificateFormat field to given value.

### HasCertificateFormat

`func (o *ItemGeneralInfo) HasCertificateFormat() bool`

HasCertificateFormat returns a boolean if a field has been set.

### GetCertificatesTemplateInfo

`func (o *ItemGeneralInfo) GetCertificatesTemplateInfo() CertificateTemplateInfo`

GetCertificatesTemplateInfo returns the CertificatesTemplateInfo field if non-nil, zero value otherwise.

### GetCertificatesTemplateInfoOk

`func (o *ItemGeneralInfo) GetCertificatesTemplateInfoOk() (*CertificateTemplateInfo, bool)`

GetCertificatesTemplateInfoOk returns a tuple with the CertificatesTemplateInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificatesTemplateInfo

`func (o *ItemGeneralInfo) SetCertificatesTemplateInfo(v CertificateTemplateInfo)`

SetCertificatesTemplateInfo sets CertificatesTemplateInfo field to given value.

### HasCertificatesTemplateInfo

`func (o *ItemGeneralInfo) HasCertificatesTemplateInfo() bool`

HasCertificatesTemplateInfo returns a boolean if a field has been set.

### GetClassicKeyDetails

`func (o *ItemGeneralInfo) GetClassicKeyDetails() ClassicKeyDetailsInfo`

GetClassicKeyDetails returns the ClassicKeyDetails field if non-nil, zero value otherwise.

### GetClassicKeyDetailsOk

`func (o *ItemGeneralInfo) GetClassicKeyDetailsOk() (*ClassicKeyDetailsInfo, bool)`

GetClassicKeyDetailsOk returns a tuple with the ClassicKeyDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassicKeyDetails

`func (o *ItemGeneralInfo) SetClassicKeyDetails(v ClassicKeyDetailsInfo)`

SetClassicKeyDetails sets ClassicKeyDetails field to given value.

### HasClassicKeyDetails

`func (o *ItemGeneralInfo) HasClassicKeyDetails() bool`

HasClassicKeyDetails returns a boolean if a field has been set.

### GetClusterGwUrl

`func (o *ItemGeneralInfo) GetClusterGwUrl() string`

GetClusterGwUrl returns the ClusterGwUrl field if non-nil, zero value otherwise.

### GetClusterGwUrlOk

`func (o *ItemGeneralInfo) GetClusterGwUrlOk() (*string, bool)`

GetClusterGwUrlOk returns a tuple with the ClusterGwUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterGwUrl

`func (o *ItemGeneralInfo) SetClusterGwUrl(v string)`

SetClusterGwUrl sets ClusterGwUrl field to given value.

### HasClusterGwUrl

`func (o *ItemGeneralInfo) HasClusterGwUrl() bool`

HasClusterGwUrl returns a boolean if a field has been set.

### GetDisplayMetadata

`func (o *ItemGeneralInfo) GetDisplayMetadata() string`

GetDisplayMetadata returns the DisplayMetadata field if non-nil, zero value otherwise.

### GetDisplayMetadataOk

`func (o *ItemGeneralInfo) GetDisplayMetadataOk() (*string, bool)`

GetDisplayMetadataOk returns a tuple with the DisplayMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMetadata

`func (o *ItemGeneralInfo) SetDisplayMetadata(v string)`

SetDisplayMetadata sets DisplayMetadata field to given value.

### HasDisplayMetadata

`func (o *ItemGeneralInfo) HasDisplayMetadata() bool`

HasDisplayMetadata returns a boolean if a field has been set.

### GetDynamicSecretProducerDetails

`func (o *ItemGeneralInfo) GetDynamicSecretProducerDetails() DynamicSecretProducerInfo`

GetDynamicSecretProducerDetails returns the DynamicSecretProducerDetails field if non-nil, zero value otherwise.

### GetDynamicSecretProducerDetailsOk

`func (o *ItemGeneralInfo) GetDynamicSecretProducerDetailsOk() (*DynamicSecretProducerInfo, bool)`

GetDynamicSecretProducerDetailsOk returns a tuple with the DynamicSecretProducerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicSecretProducerDetails

`func (o *ItemGeneralInfo) SetDynamicSecretProducerDetails(v DynamicSecretProducerInfo)`

SetDynamicSecretProducerDetails sets DynamicSecretProducerDetails field to given value.

### HasDynamicSecretProducerDetails

`func (o *ItemGeneralInfo) HasDynamicSecretProducerDetails() bool`

HasDynamicSecretProducerDetails returns a boolean if a field has been set.

### GetExpirationEvents

`func (o *ItemGeneralInfo) GetExpirationEvents() []CertificateExpirationEvent`

GetExpirationEvents returns the ExpirationEvents field if non-nil, zero value otherwise.

### GetExpirationEventsOk

`func (o *ItemGeneralInfo) GetExpirationEventsOk() (*[]CertificateExpirationEvent, bool)`

GetExpirationEventsOk returns a tuple with the ExpirationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationEvents

`func (o *ItemGeneralInfo) SetExpirationEvents(v []CertificateExpirationEvent)`

SetExpirationEvents sets ExpirationEvents field to given value.

### HasExpirationEvents

`func (o *ItemGeneralInfo) HasExpirationEvents() bool`

HasExpirationEvents returns a boolean if a field has been set.

### GetImporterInfo

`func (o *ItemGeneralInfo) GetImporterInfo() ImporterInfo`

GetImporterInfo returns the ImporterInfo field if non-nil, zero value otherwise.

### GetImporterInfoOk

`func (o *ItemGeneralInfo) GetImporterInfoOk() (*ImporterInfo, bool)`

GetImporterInfoOk returns a tuple with the ImporterInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterInfo

`func (o *ItemGeneralInfo) SetImporterInfo(v ImporterInfo)`

SetImporterInfo sets ImporterInfo field to given value.

### HasImporterInfo

`func (o *ItemGeneralInfo) HasImporterInfo() bool`

HasImporterInfo returns a boolean if a field has been set.

### GetIssuerOverviewInfo

`func (o *ItemGeneralInfo) GetIssuerOverviewInfo() IssuerOverviewInfo`

GetIssuerOverviewInfo returns the IssuerOverviewInfo field if non-nil, zero value otherwise.

### GetIssuerOverviewInfoOk

`func (o *ItemGeneralInfo) GetIssuerOverviewInfoOk() (*IssuerOverviewInfo, bool)`

GetIssuerOverviewInfoOk returns a tuple with the IssuerOverviewInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerOverviewInfo

`func (o *ItemGeneralInfo) SetIssuerOverviewInfo(v IssuerOverviewInfo)`

SetIssuerOverviewInfo sets IssuerOverviewInfo field to given value.

### HasIssuerOverviewInfo

`func (o *ItemGeneralInfo) HasIssuerOverviewInfo() bool`

HasIssuerOverviewInfo returns a boolean if a field has been set.

### GetNextRotationEvents

`func (o *ItemGeneralInfo) GetNextRotationEvents() []NextAutoRotationEvent`

GetNextRotationEvents returns the NextRotationEvents field if non-nil, zero value otherwise.

### GetNextRotationEventsOk

`func (o *ItemGeneralInfo) GetNextRotationEventsOk() (*[]NextAutoRotationEvent, bool)`

GetNextRotationEventsOk returns a tuple with the NextRotationEvents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRotationEvents

`func (o *ItemGeneralInfo) SetNextRotationEvents(v []NextAutoRotationEvent)`

SetNextRotationEvents sets NextRotationEvents field to given value.

### HasNextRotationEvents

`func (o *ItemGeneralInfo) HasNextRotationEvents() bool`

HasNextRotationEvents returns a boolean if a field has been set.

### GetOidcClientInfo

`func (o *ItemGeneralInfo) GetOidcClientInfo() OidcClientInfo`

GetOidcClientInfo returns the OidcClientInfo field if non-nil, zero value otherwise.

### GetOidcClientInfoOk

`func (o *ItemGeneralInfo) GetOidcClientInfoOk() (*OidcClientInfo, bool)`

GetOidcClientInfoOk returns a tuple with the OidcClientInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcClientInfo

`func (o *ItemGeneralInfo) SetOidcClientInfo(v OidcClientInfo)`

SetOidcClientInfo sets OidcClientInfo field to given value.

### HasOidcClientInfo

`func (o *ItemGeneralInfo) HasOidcClientInfo() bool`

HasOidcClientInfo returns a boolean if a field has been set.

### GetPasswordPolicy

`func (o *ItemGeneralInfo) GetPasswordPolicy() PasswordPolicyInfo`

GetPasswordPolicy returns the PasswordPolicy field if non-nil, zero value otherwise.

### GetPasswordPolicyOk

`func (o *ItemGeneralInfo) GetPasswordPolicyOk() (*PasswordPolicyInfo, bool)`

GetPasswordPolicyOk returns a tuple with the PasswordPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordPolicy

`func (o *ItemGeneralInfo) SetPasswordPolicy(v PasswordPolicyInfo)`

SetPasswordPolicy sets PasswordPolicy field to given value.

### HasPasswordPolicy

`func (o *ItemGeneralInfo) HasPasswordPolicy() bool`

HasPasswordPolicy returns a boolean if a field has been set.

### GetRotatedSecretDetails

`func (o *ItemGeneralInfo) GetRotatedSecretDetails() RotatedSecretDetailsInfo`

GetRotatedSecretDetails returns the RotatedSecretDetails field if non-nil, zero value otherwise.

### GetRotatedSecretDetailsOk

`func (o *ItemGeneralInfo) GetRotatedSecretDetailsOk() (*RotatedSecretDetailsInfo, bool)`

GetRotatedSecretDetailsOk returns a tuple with the RotatedSecretDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedSecretDetails

`func (o *ItemGeneralInfo) SetRotatedSecretDetails(v RotatedSecretDetailsInfo)`

SetRotatedSecretDetails sets RotatedSecretDetails field to given value.

### HasRotatedSecretDetails

`func (o *ItemGeneralInfo) HasRotatedSecretDetails() bool`

HasRotatedSecretDetails returns a boolean if a field has been set.

### GetSecureRemoteAccessDetails

`func (o *ItemGeneralInfo) GetSecureRemoteAccessDetails() SecureRemoteAccess`

GetSecureRemoteAccessDetails returns the SecureRemoteAccessDetails field if non-nil, zero value otherwise.

### GetSecureRemoteAccessDetailsOk

`func (o *ItemGeneralInfo) GetSecureRemoteAccessDetailsOk() (*SecureRemoteAccess, bool)`

GetSecureRemoteAccessDetailsOk returns a tuple with the SecureRemoteAccessDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureRemoteAccessDetails

`func (o *ItemGeneralInfo) SetSecureRemoteAccessDetails(v SecureRemoteAccess)`

SetSecureRemoteAccessDetails sets SecureRemoteAccessDetails field to given value.

### HasSecureRemoteAccessDetails

`func (o *ItemGeneralInfo) HasSecureRemoteAccessDetails() bool`

HasSecureRemoteAccessDetails returns a boolean if a field has been set.

### GetStaticSecretInfo

`func (o *ItemGeneralInfo) GetStaticSecretInfo() StaticSecretDetailsInfo`

GetStaticSecretInfo returns the StaticSecretInfo field if non-nil, zero value otherwise.

### GetStaticSecretInfoOk

`func (o *ItemGeneralInfo) GetStaticSecretInfoOk() (*StaticSecretDetailsInfo, bool)`

GetStaticSecretInfoOk returns a tuple with the StaticSecretInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticSecretInfo

`func (o *ItemGeneralInfo) SetStaticSecretInfo(v StaticSecretDetailsInfo)`

SetStaticSecretInfo sets StaticSecretInfo field to given value.

### HasStaticSecretInfo

`func (o *ItemGeneralInfo) HasStaticSecretInfo() bool`

HasStaticSecretInfo returns a boolean if a field has been set.

### GetTokenizerInfo

`func (o *ItemGeneralInfo) GetTokenizerInfo() TokenizerInfo`

GetTokenizerInfo returns the TokenizerInfo field if non-nil, zero value otherwise.

### GetTokenizerInfoOk

`func (o *ItemGeneralInfo) GetTokenizerInfoOk() (*TokenizerInfo, bool)`

GetTokenizerInfoOk returns a tuple with the TokenizerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenizerInfo

`func (o *ItemGeneralInfo) SetTokenizerInfo(v TokenizerInfo)`

SetTokenizerInfo sets TokenizerInfo field to given value.

### HasTokenizerInfo

`func (o *ItemGeneralInfo) HasTokenizerInfo() bool`

HasTokenizerInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AuthMethodAccessInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessExpires** | Pointer to **int64** |  | [optional] 
**AccessIdAlias** | Pointer to **string** | for accounts where AccessId holds encrypted email this field will hold generated AccessId, for accounts based on regular AccessId it will be equal to accessId itself | [optional] 
**ApiKeyAccessRules** | Pointer to [**APIKeyAccessRules**](APIKeyAccessRules.md) |  | [optional] 
**AwsIamAccessRules** | Pointer to [**AWSIAMAccessRules**](AWSIAMAccessRules.md) |  | [optional] 
**AzureAdAccessRules** | Pointer to [**AzureADAccessRules**](AzureADAccessRules.md) |  | [optional] 
**CertAccessRules** | Pointer to [**CertAccessRules**](CertAccessRules.md) |  | [optional] 
**CidrWhitelist** | Pointer to **string** |  | [optional] 
**EmailPassAccessRules** | Pointer to [**EmailPassAccessRules**](EmailPassAccessRules.md) |  | [optional] 
**ForceSubClaims** | Pointer to **bool** | if true the role associated with this auth method must include sub claims | [optional] 
**GcpAccessRules** | Pointer to [**GCPAccessRules**](GCPAccessRules.md) |  | [optional] 
**GwCidrWhitelist** | Pointer to **string** |  | [optional] 
**HuaweiAccessRules** | Pointer to [**HuaweiAccessRules**](HuaweiAccessRules.md) |  | [optional] 
**JwtTtl** | Pointer to **int64** |  | [optional] 
**K8sAccessRules** | Pointer to [**KubernetesAccessRules**](KubernetesAccessRules.md) |  | [optional] 
**LdapAccessRules** | Pointer to [**LDAPAccessRules**](LDAPAccessRules.md) |  | [optional] 
**Oauth2AccessRules** | Pointer to [**OAuth2AccessRules**](OAuth2AccessRules.md) |  | [optional] 
**OidcAccessRules** | Pointer to [**OIDCAccessRules**](OIDCAccessRules.md) |  | [optional] 
**RulesType** | Pointer to **string** |  | [optional] 
**SamlAccessRules** | Pointer to [**SAMLAccessRules**](SAMLAccessRules.md) |  | [optional] 
**UniversalIdentityAccessRules** | Pointer to [**UniversalIdentityAccessRules**](UniversalIdentityAccessRules.md) |  | [optional] 

## Methods

### NewAuthMethodAccessInfo

`func NewAuthMethodAccessInfo() *AuthMethodAccessInfo`

NewAuthMethodAccessInfo instantiates a new AuthMethodAccessInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMethodAccessInfoWithDefaults

`func NewAuthMethodAccessInfoWithDefaults() *AuthMethodAccessInfo`

NewAuthMethodAccessInfoWithDefaults instantiates a new AuthMethodAccessInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessExpires

`func (o *AuthMethodAccessInfo) GetAccessExpires() int64`

GetAccessExpires returns the AccessExpires field if non-nil, zero value otherwise.

### GetAccessExpiresOk

`func (o *AuthMethodAccessInfo) GetAccessExpiresOk() (*int64, bool)`

GetAccessExpiresOk returns a tuple with the AccessExpires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessExpires

`func (o *AuthMethodAccessInfo) SetAccessExpires(v int64)`

SetAccessExpires sets AccessExpires field to given value.

### HasAccessExpires

`func (o *AuthMethodAccessInfo) HasAccessExpires() bool`

HasAccessExpires returns a boolean if a field has been set.

### GetAccessIdAlias

`func (o *AuthMethodAccessInfo) GetAccessIdAlias() string`

GetAccessIdAlias returns the AccessIdAlias field if non-nil, zero value otherwise.

### GetAccessIdAliasOk

`func (o *AuthMethodAccessInfo) GetAccessIdAliasOk() (*string, bool)`

GetAccessIdAliasOk returns a tuple with the AccessIdAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessIdAlias

`func (o *AuthMethodAccessInfo) SetAccessIdAlias(v string)`

SetAccessIdAlias sets AccessIdAlias field to given value.

### HasAccessIdAlias

`func (o *AuthMethodAccessInfo) HasAccessIdAlias() bool`

HasAccessIdAlias returns a boolean if a field has been set.

### GetApiKeyAccessRules

`func (o *AuthMethodAccessInfo) GetApiKeyAccessRules() APIKeyAccessRules`

GetApiKeyAccessRules returns the ApiKeyAccessRules field if non-nil, zero value otherwise.

### GetApiKeyAccessRulesOk

`func (o *AuthMethodAccessInfo) GetApiKeyAccessRulesOk() (*APIKeyAccessRules, bool)`

GetApiKeyAccessRulesOk returns a tuple with the ApiKeyAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyAccessRules

`func (o *AuthMethodAccessInfo) SetApiKeyAccessRules(v APIKeyAccessRules)`

SetApiKeyAccessRules sets ApiKeyAccessRules field to given value.

### HasApiKeyAccessRules

`func (o *AuthMethodAccessInfo) HasApiKeyAccessRules() bool`

HasApiKeyAccessRules returns a boolean if a field has been set.

### GetAwsIamAccessRules

`func (o *AuthMethodAccessInfo) GetAwsIamAccessRules() AWSIAMAccessRules`

GetAwsIamAccessRules returns the AwsIamAccessRules field if non-nil, zero value otherwise.

### GetAwsIamAccessRulesOk

`func (o *AuthMethodAccessInfo) GetAwsIamAccessRulesOk() (*AWSIAMAccessRules, bool)`

GetAwsIamAccessRulesOk returns a tuple with the AwsIamAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsIamAccessRules

`func (o *AuthMethodAccessInfo) SetAwsIamAccessRules(v AWSIAMAccessRules)`

SetAwsIamAccessRules sets AwsIamAccessRules field to given value.

### HasAwsIamAccessRules

`func (o *AuthMethodAccessInfo) HasAwsIamAccessRules() bool`

HasAwsIamAccessRules returns a boolean if a field has been set.

### GetAzureAdAccessRules

`func (o *AuthMethodAccessInfo) GetAzureAdAccessRules() AzureADAccessRules`

GetAzureAdAccessRules returns the AzureAdAccessRules field if non-nil, zero value otherwise.

### GetAzureAdAccessRulesOk

`func (o *AuthMethodAccessInfo) GetAzureAdAccessRulesOk() (*AzureADAccessRules, bool)`

GetAzureAdAccessRulesOk returns a tuple with the AzureAdAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdAccessRules

`func (o *AuthMethodAccessInfo) SetAzureAdAccessRules(v AzureADAccessRules)`

SetAzureAdAccessRules sets AzureAdAccessRules field to given value.

### HasAzureAdAccessRules

`func (o *AuthMethodAccessInfo) HasAzureAdAccessRules() bool`

HasAzureAdAccessRules returns a boolean if a field has been set.

### GetCertAccessRules

`func (o *AuthMethodAccessInfo) GetCertAccessRules() CertAccessRules`

GetCertAccessRules returns the CertAccessRules field if non-nil, zero value otherwise.

### GetCertAccessRulesOk

`func (o *AuthMethodAccessInfo) GetCertAccessRulesOk() (*CertAccessRules, bool)`

GetCertAccessRulesOk returns a tuple with the CertAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertAccessRules

`func (o *AuthMethodAccessInfo) SetCertAccessRules(v CertAccessRules)`

SetCertAccessRules sets CertAccessRules field to given value.

### HasCertAccessRules

`func (o *AuthMethodAccessInfo) HasCertAccessRules() bool`

HasCertAccessRules returns a boolean if a field has been set.

### GetCidrWhitelist

`func (o *AuthMethodAccessInfo) GetCidrWhitelist() string`

GetCidrWhitelist returns the CidrWhitelist field if non-nil, zero value otherwise.

### GetCidrWhitelistOk

`func (o *AuthMethodAccessInfo) GetCidrWhitelistOk() (*string, bool)`

GetCidrWhitelistOk returns a tuple with the CidrWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrWhitelist

`func (o *AuthMethodAccessInfo) SetCidrWhitelist(v string)`

SetCidrWhitelist sets CidrWhitelist field to given value.

### HasCidrWhitelist

`func (o *AuthMethodAccessInfo) HasCidrWhitelist() bool`

HasCidrWhitelist returns a boolean if a field has been set.

### GetEmailPassAccessRules

`func (o *AuthMethodAccessInfo) GetEmailPassAccessRules() EmailPassAccessRules`

GetEmailPassAccessRules returns the EmailPassAccessRules field if non-nil, zero value otherwise.

### GetEmailPassAccessRulesOk

`func (o *AuthMethodAccessInfo) GetEmailPassAccessRulesOk() (*EmailPassAccessRules, bool)`

GetEmailPassAccessRulesOk returns a tuple with the EmailPassAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailPassAccessRules

`func (o *AuthMethodAccessInfo) SetEmailPassAccessRules(v EmailPassAccessRules)`

SetEmailPassAccessRules sets EmailPassAccessRules field to given value.

### HasEmailPassAccessRules

`func (o *AuthMethodAccessInfo) HasEmailPassAccessRules() bool`

HasEmailPassAccessRules returns a boolean if a field has been set.

### GetForceSubClaims

`func (o *AuthMethodAccessInfo) GetForceSubClaims() bool`

GetForceSubClaims returns the ForceSubClaims field if non-nil, zero value otherwise.

### GetForceSubClaimsOk

`func (o *AuthMethodAccessInfo) GetForceSubClaimsOk() (*bool, bool)`

GetForceSubClaimsOk returns a tuple with the ForceSubClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSubClaims

`func (o *AuthMethodAccessInfo) SetForceSubClaims(v bool)`

SetForceSubClaims sets ForceSubClaims field to given value.

### HasForceSubClaims

`func (o *AuthMethodAccessInfo) HasForceSubClaims() bool`

HasForceSubClaims returns a boolean if a field has been set.

### GetGcpAccessRules

`func (o *AuthMethodAccessInfo) GetGcpAccessRules() GCPAccessRules`

GetGcpAccessRules returns the GcpAccessRules field if non-nil, zero value otherwise.

### GetGcpAccessRulesOk

`func (o *AuthMethodAccessInfo) GetGcpAccessRulesOk() (*GCPAccessRules, bool)`

GetGcpAccessRulesOk returns a tuple with the GcpAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpAccessRules

`func (o *AuthMethodAccessInfo) SetGcpAccessRules(v GCPAccessRules)`

SetGcpAccessRules sets GcpAccessRules field to given value.

### HasGcpAccessRules

`func (o *AuthMethodAccessInfo) HasGcpAccessRules() bool`

HasGcpAccessRules returns a boolean if a field has been set.

### GetGwCidrWhitelist

`func (o *AuthMethodAccessInfo) GetGwCidrWhitelist() string`

GetGwCidrWhitelist returns the GwCidrWhitelist field if non-nil, zero value otherwise.

### GetGwCidrWhitelistOk

`func (o *AuthMethodAccessInfo) GetGwCidrWhitelistOk() (*string, bool)`

GetGwCidrWhitelistOk returns a tuple with the GwCidrWhitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGwCidrWhitelist

`func (o *AuthMethodAccessInfo) SetGwCidrWhitelist(v string)`

SetGwCidrWhitelist sets GwCidrWhitelist field to given value.

### HasGwCidrWhitelist

`func (o *AuthMethodAccessInfo) HasGwCidrWhitelist() bool`

HasGwCidrWhitelist returns a boolean if a field has been set.

### GetHuaweiAccessRules

`func (o *AuthMethodAccessInfo) GetHuaweiAccessRules() HuaweiAccessRules`

GetHuaweiAccessRules returns the HuaweiAccessRules field if non-nil, zero value otherwise.

### GetHuaweiAccessRulesOk

`func (o *AuthMethodAccessInfo) GetHuaweiAccessRulesOk() (*HuaweiAccessRules, bool)`

GetHuaweiAccessRulesOk returns a tuple with the HuaweiAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHuaweiAccessRules

`func (o *AuthMethodAccessInfo) SetHuaweiAccessRules(v HuaweiAccessRules)`

SetHuaweiAccessRules sets HuaweiAccessRules field to given value.

### HasHuaweiAccessRules

`func (o *AuthMethodAccessInfo) HasHuaweiAccessRules() bool`

HasHuaweiAccessRules returns a boolean if a field has been set.

### GetJwtTtl

`func (o *AuthMethodAccessInfo) GetJwtTtl() int64`

GetJwtTtl returns the JwtTtl field if non-nil, zero value otherwise.

### GetJwtTtlOk

`func (o *AuthMethodAccessInfo) GetJwtTtlOk() (*int64, bool)`

GetJwtTtlOk returns a tuple with the JwtTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwtTtl

`func (o *AuthMethodAccessInfo) SetJwtTtl(v int64)`

SetJwtTtl sets JwtTtl field to given value.

### HasJwtTtl

`func (o *AuthMethodAccessInfo) HasJwtTtl() bool`

HasJwtTtl returns a boolean if a field has been set.

### GetK8sAccessRules

`func (o *AuthMethodAccessInfo) GetK8sAccessRules() KubernetesAccessRules`

GetK8sAccessRules returns the K8sAccessRules field if non-nil, zero value otherwise.

### GetK8sAccessRulesOk

`func (o *AuthMethodAccessInfo) GetK8sAccessRulesOk() (*KubernetesAccessRules, bool)`

GetK8sAccessRulesOk returns a tuple with the K8sAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAccessRules

`func (o *AuthMethodAccessInfo) SetK8sAccessRules(v KubernetesAccessRules)`

SetK8sAccessRules sets K8sAccessRules field to given value.

### HasK8sAccessRules

`func (o *AuthMethodAccessInfo) HasK8sAccessRules() bool`

HasK8sAccessRules returns a boolean if a field has been set.

### GetLdapAccessRules

`func (o *AuthMethodAccessInfo) GetLdapAccessRules() LDAPAccessRules`

GetLdapAccessRules returns the LdapAccessRules field if non-nil, zero value otherwise.

### GetLdapAccessRulesOk

`func (o *AuthMethodAccessInfo) GetLdapAccessRulesOk() (*LDAPAccessRules, bool)`

GetLdapAccessRulesOk returns a tuple with the LdapAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAccessRules

`func (o *AuthMethodAccessInfo) SetLdapAccessRules(v LDAPAccessRules)`

SetLdapAccessRules sets LdapAccessRules field to given value.

### HasLdapAccessRules

`func (o *AuthMethodAccessInfo) HasLdapAccessRules() bool`

HasLdapAccessRules returns a boolean if a field has been set.

### GetOauth2AccessRules

`func (o *AuthMethodAccessInfo) GetOauth2AccessRules() OAuth2AccessRules`

GetOauth2AccessRules returns the Oauth2AccessRules field if non-nil, zero value otherwise.

### GetOauth2AccessRulesOk

`func (o *AuthMethodAccessInfo) GetOauth2AccessRulesOk() (*OAuth2AccessRules, bool)`

GetOauth2AccessRulesOk returns a tuple with the Oauth2AccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth2AccessRules

`func (o *AuthMethodAccessInfo) SetOauth2AccessRules(v OAuth2AccessRules)`

SetOauth2AccessRules sets Oauth2AccessRules field to given value.

### HasOauth2AccessRules

`func (o *AuthMethodAccessInfo) HasOauth2AccessRules() bool`

HasOauth2AccessRules returns a boolean if a field has been set.

### GetOidcAccessRules

`func (o *AuthMethodAccessInfo) GetOidcAccessRules() OIDCAccessRules`

GetOidcAccessRules returns the OidcAccessRules field if non-nil, zero value otherwise.

### GetOidcAccessRulesOk

`func (o *AuthMethodAccessInfo) GetOidcAccessRulesOk() (*OIDCAccessRules, bool)`

GetOidcAccessRulesOk returns a tuple with the OidcAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOidcAccessRules

`func (o *AuthMethodAccessInfo) SetOidcAccessRules(v OIDCAccessRules)`

SetOidcAccessRules sets OidcAccessRules field to given value.

### HasOidcAccessRules

`func (o *AuthMethodAccessInfo) HasOidcAccessRules() bool`

HasOidcAccessRules returns a boolean if a field has been set.

### GetRulesType

`func (o *AuthMethodAccessInfo) GetRulesType() string`

GetRulesType returns the RulesType field if non-nil, zero value otherwise.

### GetRulesTypeOk

`func (o *AuthMethodAccessInfo) GetRulesTypeOk() (*string, bool)`

GetRulesTypeOk returns a tuple with the RulesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRulesType

`func (o *AuthMethodAccessInfo) SetRulesType(v string)`

SetRulesType sets RulesType field to given value.

### HasRulesType

`func (o *AuthMethodAccessInfo) HasRulesType() bool`

HasRulesType returns a boolean if a field has been set.

### GetSamlAccessRules

`func (o *AuthMethodAccessInfo) GetSamlAccessRules() SAMLAccessRules`

GetSamlAccessRules returns the SamlAccessRules field if non-nil, zero value otherwise.

### GetSamlAccessRulesOk

`func (o *AuthMethodAccessInfo) GetSamlAccessRulesOk() (*SAMLAccessRules, bool)`

GetSamlAccessRulesOk returns a tuple with the SamlAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamlAccessRules

`func (o *AuthMethodAccessInfo) SetSamlAccessRules(v SAMLAccessRules)`

SetSamlAccessRules sets SamlAccessRules field to given value.

### HasSamlAccessRules

`func (o *AuthMethodAccessInfo) HasSamlAccessRules() bool`

HasSamlAccessRules returns a boolean if a field has been set.

### GetUniversalIdentityAccessRules

`func (o *AuthMethodAccessInfo) GetUniversalIdentityAccessRules() UniversalIdentityAccessRules`

GetUniversalIdentityAccessRules returns the UniversalIdentityAccessRules field if non-nil, zero value otherwise.

### GetUniversalIdentityAccessRulesOk

`func (o *AuthMethodAccessInfo) GetUniversalIdentityAccessRulesOk() (*UniversalIdentityAccessRules, bool)`

GetUniversalIdentityAccessRulesOk returns a tuple with the UniversalIdentityAccessRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalIdentityAccessRules

`func (o *AuthMethodAccessInfo) SetUniversalIdentityAccessRules(v UniversalIdentityAccessRules)`

SetUniversalIdentityAccessRules sets UniversalIdentityAccessRules field to given value.

### HasUniversalIdentityAccessRules

`func (o *AuthMethodAccessInfo) HasUniversalIdentityAccessRules() bool`

HasUniversalIdentityAccessRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



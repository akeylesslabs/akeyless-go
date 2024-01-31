# Auth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | Access ID | [optional] 
**AccessKey** | Pointer to **string** | Access key (relevant only for access-type&#x3D;access_key) | [optional] 
**AccessType** | Pointer to **string** | Access Type (access_key/password/saml/ldap/k8s/azure_ad/oidc/aws_iam/universal_identity/jwt/gcp/cert) | [optional] [default to "access_key"]
**AccountId** | Pointer to **string** | Account id (relevant only for access-type&#x3D;password where the email address is associated with more than one account) | [optional] 
**AdminEmail** | Pointer to **string** | Email (relevant only for access-type&#x3D;password) | [optional] 
**AdminPassword** | Pointer to **string** | Password (relevant only for access-type&#x3D;password) | [optional] 
**CertData** | Pointer to **string** | Certificate data encoded in base64. Used if file was not provided. (relevant only for access-type&#x3D;cert) | [optional] 
**CloudId** | Pointer to **string** | The cloud identity (relevant only for access-type&#x3D;azure_ad,aws_iam,gcp) | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**GatewayUrl** | Pointer to **string** | Gateway URL for the K8S/OAUTH2 authenticated (relevant only for access-type&#x3D;k8s/oauth2) | [optional] 
**GcpAudience** | Pointer to **string** | GCP JWT audience | [optional] [default to "akeyless.io"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Jwt** | Pointer to **string** | The Json Web Token (relevant only for access-type&#x3D;jwt/oidc) | [optional] 
**K8sAuthConfigName** | Pointer to **string** | The K8S Auth config name (relevant only for access-type&#x3D;k8s) | [optional] 
**K8sServiceAccountToken** | Pointer to **string** | The K8S service account token. (relevant only for access-type&#x3D;k8s) | [optional] 
**KeyData** | Pointer to **string** | Private key data encoded in base64. Used if file was not provided.(relevant only for access-type&#x3D;cert) | [optional] 
**LdapPassword** | Pointer to **string** | LDAP password (relevant only for access-type&#x3D;ldap) | [optional] 
**LdapUsername** | Pointer to **string** | LDAP username (relevant only for access-type&#x3D;ldap) | [optional] 
**OciAuthType** | Pointer to **string** | The type of the OCI configuration to use [instance/apikey/resource] (relevant only for access-type&#x3D;oci) | [optional] [default to "apikey"]
**OciGroupOcid** | Pointer to **[]string** | A list of Oracle Cloud IDs groups (relevant only for access-type&#x3D;oci) | [optional] 
**UidToken** | Pointer to **string** | The universal_identity token (relevant only for access-type&#x3D;universal_identity) | [optional] 

## Methods

### NewAuth

`func NewAuth() *Auth`

NewAuth instantiates a new Auth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthWithDefaults

`func NewAuthWithDefaults() *Auth`

NewAuthWithDefaults instantiates a new Auth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *Auth) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *Auth) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *Auth) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *Auth) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAccessKey

`func (o *Auth) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Auth) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Auth) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *Auth) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessType

`func (o *Auth) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Auth) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Auth) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *Auth) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAccountId

`func (o *Auth) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Auth) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Auth) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Auth) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAdminEmail

`func (o *Auth) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *Auth) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *Auth) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *Auth) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetAdminPassword

`func (o *Auth) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *Auth) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *Auth) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *Auth) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetCertData

`func (o *Auth) GetCertData() string`

GetCertData returns the CertData field if non-nil, zero value otherwise.

### GetCertDataOk

`func (o *Auth) GetCertDataOk() (*string, bool)`

GetCertDataOk returns a tuple with the CertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertData

`func (o *Auth) SetCertData(v string)`

SetCertData sets CertData field to given value.

### HasCertData

`func (o *Auth) HasCertData() bool`

HasCertData returns a boolean if a field has been set.

### GetCloudId

`func (o *Auth) GetCloudId() string`

GetCloudId returns the CloudId field if non-nil, zero value otherwise.

### GetCloudIdOk

`func (o *Auth) GetCloudIdOk() (*string, bool)`

GetCloudIdOk returns a tuple with the CloudId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudId

`func (o *Auth) SetCloudId(v string)`

SetCloudId sets CloudId field to given value.

### HasCloudId

`func (o *Auth) HasCloudId() bool`

HasCloudId returns a boolean if a field has been set.

### GetDebug

`func (o *Auth) GetDebug() bool`

GetDebug returns the Debug field if non-nil, zero value otherwise.

### GetDebugOk

`func (o *Auth) GetDebugOk() (*bool, bool)`

GetDebugOk returns a tuple with the Debug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebug

`func (o *Auth) SetDebug(v bool)`

SetDebug sets Debug field to given value.

### HasDebug

`func (o *Auth) HasDebug() bool`

HasDebug returns a boolean if a field has been set.

### GetGatewayUrl

`func (o *Auth) GetGatewayUrl() string`

GetGatewayUrl returns the GatewayUrl field if non-nil, zero value otherwise.

### GetGatewayUrlOk

`func (o *Auth) GetGatewayUrlOk() (*string, bool)`

GetGatewayUrlOk returns a tuple with the GatewayUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayUrl

`func (o *Auth) SetGatewayUrl(v string)`

SetGatewayUrl sets GatewayUrl field to given value.

### HasGatewayUrl

`func (o *Auth) HasGatewayUrl() bool`

HasGatewayUrl returns a boolean if a field has been set.

### GetGcpAudience

`func (o *Auth) GetGcpAudience() string`

GetGcpAudience returns the GcpAudience field if non-nil, zero value otherwise.

### GetGcpAudienceOk

`func (o *Auth) GetGcpAudienceOk() (*string, bool)`

GetGcpAudienceOk returns a tuple with the GcpAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpAudience

`func (o *Auth) SetGcpAudience(v string)`

SetGcpAudience sets GcpAudience field to given value.

### HasGcpAudience

`func (o *Auth) HasGcpAudience() bool`

HasGcpAudience returns a boolean if a field has been set.

### GetJson

`func (o *Auth) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Auth) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Auth) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *Auth) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetJwt

`func (o *Auth) GetJwt() string`

GetJwt returns the Jwt field if non-nil, zero value otherwise.

### GetJwtOk

`func (o *Auth) GetJwtOk() (*string, bool)`

GetJwtOk returns a tuple with the Jwt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwt

`func (o *Auth) SetJwt(v string)`

SetJwt sets Jwt field to given value.

### HasJwt

`func (o *Auth) HasJwt() bool`

HasJwt returns a boolean if a field has been set.

### GetK8sAuthConfigName

`func (o *Auth) GetK8sAuthConfigName() string`

GetK8sAuthConfigName returns the K8sAuthConfigName field if non-nil, zero value otherwise.

### GetK8sAuthConfigNameOk

`func (o *Auth) GetK8sAuthConfigNameOk() (*string, bool)`

GetK8sAuthConfigNameOk returns a tuple with the K8sAuthConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuthConfigName

`func (o *Auth) SetK8sAuthConfigName(v string)`

SetK8sAuthConfigName sets K8sAuthConfigName field to given value.

### HasK8sAuthConfigName

`func (o *Auth) HasK8sAuthConfigName() bool`

HasK8sAuthConfigName returns a boolean if a field has been set.

### GetK8sServiceAccountToken

`func (o *Auth) GetK8sServiceAccountToken() string`

GetK8sServiceAccountToken returns the K8sServiceAccountToken field if non-nil, zero value otherwise.

### GetK8sServiceAccountTokenOk

`func (o *Auth) GetK8sServiceAccountTokenOk() (*string, bool)`

GetK8sServiceAccountTokenOk returns a tuple with the K8sServiceAccountToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceAccountToken

`func (o *Auth) SetK8sServiceAccountToken(v string)`

SetK8sServiceAccountToken sets K8sServiceAccountToken field to given value.

### HasK8sServiceAccountToken

`func (o *Auth) HasK8sServiceAccountToken() bool`

HasK8sServiceAccountToken returns a boolean if a field has been set.

### GetKeyData

`func (o *Auth) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *Auth) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *Auth) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *Auth) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetLdapPassword

`func (o *Auth) GetLdapPassword() string`

GetLdapPassword returns the LdapPassword field if non-nil, zero value otherwise.

### GetLdapPasswordOk

`func (o *Auth) GetLdapPasswordOk() (*string, bool)`

GetLdapPasswordOk returns a tuple with the LdapPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapPassword

`func (o *Auth) SetLdapPassword(v string)`

SetLdapPassword sets LdapPassword field to given value.

### HasLdapPassword

`func (o *Auth) HasLdapPassword() bool`

HasLdapPassword returns a boolean if a field has been set.

### GetLdapUsername

`func (o *Auth) GetLdapUsername() string`

GetLdapUsername returns the LdapUsername field if non-nil, zero value otherwise.

### GetLdapUsernameOk

`func (o *Auth) GetLdapUsernameOk() (*string, bool)`

GetLdapUsernameOk returns a tuple with the LdapUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapUsername

`func (o *Auth) SetLdapUsername(v string)`

SetLdapUsername sets LdapUsername field to given value.

### HasLdapUsername

`func (o *Auth) HasLdapUsername() bool`

HasLdapUsername returns a boolean if a field has been set.

### GetOciAuthType

`func (o *Auth) GetOciAuthType() string`

GetOciAuthType returns the OciAuthType field if non-nil, zero value otherwise.

### GetOciAuthTypeOk

`func (o *Auth) GetOciAuthTypeOk() (*string, bool)`

GetOciAuthTypeOk returns a tuple with the OciAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciAuthType

`func (o *Auth) SetOciAuthType(v string)`

SetOciAuthType sets OciAuthType field to given value.

### HasOciAuthType

`func (o *Auth) HasOciAuthType() bool`

HasOciAuthType returns a boolean if a field has been set.

### GetOciGroupOcid

`func (o *Auth) GetOciGroupOcid() []string`

GetOciGroupOcid returns the OciGroupOcid field if non-nil, zero value otherwise.

### GetOciGroupOcidOk

`func (o *Auth) GetOciGroupOcidOk() (*[]string, bool)`

GetOciGroupOcidOk returns a tuple with the OciGroupOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciGroupOcid

`func (o *Auth) SetOciGroupOcid(v []string)`

SetOciGroupOcid sets OciGroupOcid field to given value.

### HasOciGroupOcid

`func (o *Auth) HasOciGroupOcid() bool`

HasOciGroupOcid returns a boolean if a field has been set.

### GetUidToken

`func (o *Auth) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *Auth) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *Auth) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *Auth) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



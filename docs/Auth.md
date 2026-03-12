# Auth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | Access ID | [optional] 
**AccessKey** | Pointer to **string** | Access key (relevant only for access-type&#x3D;access_key) | [optional] 
**AccessType** | Pointer to **string** | Access Type (access_key/password/saml/ldap/k8s/azure_ad/oidc/aws_iam/universal_identity/jwt/gcp/cert/oci/kerberos) | [optional] [default to "access_key"]
**AccountId** | Pointer to **string** | Account id (relevant only for access-type&#x3D;password where the email address is associated with more than one account) | [optional] 
**AdminEmail** | Pointer to **string** | Email (relevant only for access-type&#x3D;password) | [optional] 
**AdminPassword** | Pointer to **string** | Password (relevant only for access-type&#x3D;password) | [optional] 
**AzureCloud** | Pointer to **string** | Azure cloud environment to use. Values: AzureCloud (default), AzureUSGovernment, AzureChinaCloud. | [optional] [default to "AzureCloud"]
**CertChallenge** | Pointer to **string** | Certificate challenge encoded in base64. (relevant only for access-type&#x3D;cert) | [optional] 
**CertData** | Pointer to **string** | Certificate data encoded in base64. Used if file was not provided. (relevant only for access-type&#x3D;cert) | [optional] 
**CloudId** | Pointer to **string** | The cloud identity (relevant only for access-type&#x3D;azure_ad,aws_iam,gcp) | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**DisablePafxfast** | Pointer to **string** | Disable the FAST negotiation in the Kerberos authentication method | [optional] 
**GatewaySpn** | Pointer to **string** | The service principal name of the gateway as registered in LDAP (i.e., HTTP/gateway) | [optional] 
**GatewayUrl** | Pointer to **string** | Gateway URL relevant only for access-type&#x3D;k8s/oauth2/saml/oidc | [optional] 
**GcpAudience** | Pointer to **string** | GCP JWT audience | [optional] [default to "akeyless.io"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Jwt** | Pointer to **string** | The Json Web Token (relevant only for access-type&#x3D;jwt/oidc) | [optional] 
**K8sAuthConfigName** | Pointer to **string** | The K8S Auth config name (relevant only for access-type&#x3D;k8s) | [optional] 
**K8sServiceAccountToken** | Pointer to **string** | The K8S service account token. (relevant only for access-type&#x3D;k8s) | [optional] 
**KerberosToken** | Pointer to **string** | KerberosToken represents a Kerberos token generated for the gateway SPN (Service Principal Name). | [optional] 
**KerberosUsername** | Pointer to **string** | TThe username for the entry within the keytab to authenticate via Kerberos | [optional] 
**KeyData** | Pointer to **string** | Private key data encoded in base64. Used if file was not provided.(relevant only for access-type&#x3D;cert) | [optional] 
**KeytabData** | Pointer to **string** | Base64-encoded content of a valid keytab file, containing the service account&#39;s entry. | [optional] 
**Krb5ConfData** | Pointer to **string** | Base64-encoded content of a valid krb5.conf file, specifying the settings and parameters required for Kerberos authentication. | [optional] 
**LdapPassword** | Pointer to **string** | LDAP password (relevant only for access-type&#x3D;ldap) | [optional] 
**OciAuthType** | Pointer to **string** | The type of the OCI configuration to use [instance/apikey/resource] (relevant only for access-type&#x3D;oci) | [optional] [default to "apikey"]
**OciGroupOcid** | Pointer to **[]string** | A list of Oracle Cloud IDs groups (relevant only for access-type&#x3D;oci) | [optional] 
**Otp** | Pointer to **string** |  | [optional] 
**SignedCertChallenge** | Pointer to **string** | Signed certificate challenge encoded in base64. (relevant only for access-type&#x3D;cert) | [optional] 
**UidToken** | Pointer to **string** | The universal_identity token (relevant only for access-type&#x3D;universal_identity) | [optional] 
**UseRemoteBrowser** | Pointer to **bool** | Returns a link to complete the authentication remotely (relevant only for access-type&#x3D;saml/oidc) | [optional] 
**Username** | Pointer to **string** | LDAP username (relevant only for access-type&#x3D;ldap) | [optional] 

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

### GetAzureCloud

`func (o *Auth) GetAzureCloud() string`

GetAzureCloud returns the AzureCloud field if non-nil, zero value otherwise.

### GetAzureCloudOk

`func (o *Auth) GetAzureCloudOk() (*string, bool)`

GetAzureCloudOk returns a tuple with the AzureCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCloud

`func (o *Auth) SetAzureCloud(v string)`

SetAzureCloud sets AzureCloud field to given value.

### HasAzureCloud

`func (o *Auth) HasAzureCloud() bool`

HasAzureCloud returns a boolean if a field has been set.

### GetCertChallenge

`func (o *Auth) GetCertChallenge() string`

GetCertChallenge returns the CertChallenge field if non-nil, zero value otherwise.

### GetCertChallengeOk

`func (o *Auth) GetCertChallengeOk() (*string, bool)`

GetCertChallengeOk returns a tuple with the CertChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertChallenge

`func (o *Auth) SetCertChallenge(v string)`

SetCertChallenge sets CertChallenge field to given value.

### HasCertChallenge

`func (o *Auth) HasCertChallenge() bool`

HasCertChallenge returns a boolean if a field has been set.

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

### GetDisablePafxfast

`func (o *Auth) GetDisablePafxfast() string`

GetDisablePafxfast returns the DisablePafxfast field if non-nil, zero value otherwise.

### GetDisablePafxfastOk

`func (o *Auth) GetDisablePafxfastOk() (*string, bool)`

GetDisablePafxfastOk returns a tuple with the DisablePafxfast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePafxfast

`func (o *Auth) SetDisablePafxfast(v string)`

SetDisablePafxfast sets DisablePafxfast field to given value.

### HasDisablePafxfast

`func (o *Auth) HasDisablePafxfast() bool`

HasDisablePafxfast returns a boolean if a field has been set.

### GetGatewaySpn

`func (o *Auth) GetGatewaySpn() string`

GetGatewaySpn returns the GatewaySpn field if non-nil, zero value otherwise.

### GetGatewaySpnOk

`func (o *Auth) GetGatewaySpnOk() (*string, bool)`

GetGatewaySpnOk returns a tuple with the GatewaySpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySpn

`func (o *Auth) SetGatewaySpn(v string)`

SetGatewaySpn sets GatewaySpn field to given value.

### HasGatewaySpn

`func (o *Auth) HasGatewaySpn() bool`

HasGatewaySpn returns a boolean if a field has been set.

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

### GetKerberosToken

`func (o *Auth) GetKerberosToken() string`

GetKerberosToken returns the KerberosToken field if non-nil, zero value otherwise.

### GetKerberosTokenOk

`func (o *Auth) GetKerberosTokenOk() (*string, bool)`

GetKerberosTokenOk returns a tuple with the KerberosToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosToken

`func (o *Auth) SetKerberosToken(v string)`

SetKerberosToken sets KerberosToken field to given value.

### HasKerberosToken

`func (o *Auth) HasKerberosToken() bool`

HasKerberosToken returns a boolean if a field has been set.

### GetKerberosUsername

`func (o *Auth) GetKerberosUsername() string`

GetKerberosUsername returns the KerberosUsername field if non-nil, zero value otherwise.

### GetKerberosUsernameOk

`func (o *Auth) GetKerberosUsernameOk() (*string, bool)`

GetKerberosUsernameOk returns a tuple with the KerberosUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosUsername

`func (o *Auth) SetKerberosUsername(v string)`

SetKerberosUsername sets KerberosUsername field to given value.

### HasKerberosUsername

`func (o *Auth) HasKerberosUsername() bool`

HasKerberosUsername returns a boolean if a field has been set.

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

### GetKeytabData

`func (o *Auth) GetKeytabData() string`

GetKeytabData returns the KeytabData field if non-nil, zero value otherwise.

### GetKeytabDataOk

`func (o *Auth) GetKeytabDataOk() (*string, bool)`

GetKeytabDataOk returns a tuple with the KeytabData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabData

`func (o *Auth) SetKeytabData(v string)`

SetKeytabData sets KeytabData field to given value.

### HasKeytabData

`func (o *Auth) HasKeytabData() bool`

HasKeytabData returns a boolean if a field has been set.

### GetKrb5ConfData

`func (o *Auth) GetKrb5ConfData() string`

GetKrb5ConfData returns the Krb5ConfData field if non-nil, zero value otherwise.

### GetKrb5ConfDataOk

`func (o *Auth) GetKrb5ConfDataOk() (*string, bool)`

GetKrb5ConfDataOk returns a tuple with the Krb5ConfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5ConfData

`func (o *Auth) SetKrb5ConfData(v string)`

SetKrb5ConfData sets Krb5ConfData field to given value.

### HasKrb5ConfData

`func (o *Auth) HasKrb5ConfData() bool`

HasKrb5ConfData returns a boolean if a field has been set.

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

### GetOtp

`func (o *Auth) GetOtp() string`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *Auth) GetOtpOk() (*string, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *Auth) SetOtp(v string)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *Auth) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetSignedCertChallenge

`func (o *Auth) GetSignedCertChallenge() string`

GetSignedCertChallenge returns the SignedCertChallenge field if non-nil, zero value otherwise.

### GetSignedCertChallengeOk

`func (o *Auth) GetSignedCertChallengeOk() (*string, bool)`

GetSignedCertChallengeOk returns a tuple with the SignedCertChallenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignedCertChallenge

`func (o *Auth) SetSignedCertChallenge(v string)`

SetSignedCertChallenge sets SignedCertChallenge field to given value.

### HasSignedCertChallenge

`func (o *Auth) HasSignedCertChallenge() bool`

HasSignedCertChallenge returns a boolean if a field has been set.

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

### GetUseRemoteBrowser

`func (o *Auth) GetUseRemoteBrowser() bool`

GetUseRemoteBrowser returns the UseRemoteBrowser field if non-nil, zero value otherwise.

### GetUseRemoteBrowserOk

`func (o *Auth) GetUseRemoteBrowserOk() (*bool, bool)`

GetUseRemoteBrowserOk returns a tuple with the UseRemoteBrowser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseRemoteBrowser

`func (o *Auth) SetUseRemoteBrowser(v bool)`

SetUseRemoteBrowser sets UseRemoteBrowser field to given value.

### HasUseRemoteBrowser

`func (o *Auth) HasUseRemoteBrowser() bool`

HasUseRemoteBrowser returns a boolean if a field has been set.

### GetUsername

`func (o *Auth) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Auth) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Auth) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Auth) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



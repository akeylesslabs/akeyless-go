# Configure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | Access ID | [optional] 
**AccessKey** | Pointer to **string** | Access Key | [optional] 
**AccessType** | Pointer to **string** | Access Type (access_key/password/azure_ad/saml/oidc/aws_iam/gcp/k8s/cert) | [optional] [default to "access_key"]
**AccountId** | Pointer to **string** | Account id (relevant only for access-type&#x3D;password where the email address is associated with more than one account) | [optional] 
**AdminEmail** | Pointer to **string** | Email (relevant only for access-type&#x3D;password) | [optional] 
**AdminPassword** | Pointer to **string** | Password (relevant only for access-type&#x3D;password) | [optional] 
**AzureAdObjectId** | Pointer to **string** | Azure Active Directory ObjectId (relevant only for access-type&#x3D;azure_ad) | [optional] 
**AzureCloud** | Pointer to **string** | Azure cloud environment to use. Values: AzureCloud (default), AzureUSGovernment, AzureChinaCloud. | [optional] [default to "AzureCloud"]
**CertData** | Pointer to **string** | Certificate data encoded in base64. Used if file was not provided. (relevant only for access-type&#x3D;cert in Curl Context) | [optional] 
**CertIssuerName** | Pointer to **string** | Certificate Issuer Name | [optional] 
**CertUsername** | Pointer to **string** | The username to sign in the SSH certificate (use a comma-separated list for more than one username) | [optional] 
**DefaultLocationPrefix** | Pointer to **string** | Default path prefix for name of items, targets and auth methods | [optional] 
**DisablePafxfast** | Pointer to **string** | Disable the FAST negotiation in the Kerberos authentication method | [optional] 
**GatewaySpn** | Pointer to **string** | The service principal name of the gateway as registered in LDAP (i.e., HTTP/gateway) | [optional] 
**GcpAudience** | Pointer to **string** | GCP JWT audience | [optional] [default to "akeyless.io"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sAuthConfigName** | Pointer to **string** | The K8S Auth config name (relevant only for access-type&#x3D;k8s) | [optional] 
**KerberosToken** | Pointer to **string** | KerberosToken represents a Kerberos token generated for the gateway SPN (Service Principal Name). | [optional] 
**KerberosUsername** | Pointer to **string** | TThe username for the entry within the keytab to authenticate via Kerberos | [optional] 
**KeyData** | Pointer to **string** | Private key data encoded in base64. Used if file was not provided.(relevant only for access-type&#x3D;cert in Curl Context) | [optional] 
**KeytabData** | Pointer to **string** | Base64-encoded content of a valid keytab file, containing the service account&#39;s entry. | [optional] 
**Krb5ConfData** | Pointer to **string** | Base64-encoded content of a valid krb5.conf file, specifying the settings and parameters required for Kerberos authentication. | [optional] 
**LegacySigningAlgName** | Pointer to **bool** | Set this option to output legacy (&#39;ssh-rsa-cert-v01@openssh.com&#39;) signing algorithm name in the certificate. | [optional] 
**OciAuthType** | Pointer to **string** | The type of the OCI configuration to use [instance/apikey/resource] (relevant only for access-type&#x3D;oci) | [optional] [default to "apikey"]
**OciGroupOcid** | Pointer to **[]string** | A list of Oracle Cloud IDs groups (relevant only for access-type&#x3D;oci) | [optional] 

## Methods

### NewConfigure

`func NewConfigure() *Configure`

NewConfigure instantiates a new Configure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigureWithDefaults

`func NewConfigureWithDefaults() *Configure`

NewConfigureWithDefaults instantiates a new Configure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessId

`func (o *Configure) GetAccessId() string`

GetAccessId returns the AccessId field if non-nil, zero value otherwise.

### GetAccessIdOk

`func (o *Configure) GetAccessIdOk() (*string, bool)`

GetAccessIdOk returns a tuple with the AccessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessId

`func (o *Configure) SetAccessId(v string)`

SetAccessId sets AccessId field to given value.

### HasAccessId

`func (o *Configure) HasAccessId() bool`

HasAccessId returns a boolean if a field has been set.

### GetAccessKey

`func (o *Configure) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *Configure) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *Configure) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *Configure) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessType

`func (o *Configure) GetAccessType() string`

GetAccessType returns the AccessType field if non-nil, zero value otherwise.

### GetAccessTypeOk

`func (o *Configure) GetAccessTypeOk() (*string, bool)`

GetAccessTypeOk returns a tuple with the AccessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessType

`func (o *Configure) SetAccessType(v string)`

SetAccessType sets AccessType field to given value.

### HasAccessType

`func (o *Configure) HasAccessType() bool`

HasAccessType returns a boolean if a field has been set.

### GetAccountId

`func (o *Configure) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Configure) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Configure) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Configure) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAdminEmail

`func (o *Configure) GetAdminEmail() string`

GetAdminEmail returns the AdminEmail field if non-nil, zero value otherwise.

### GetAdminEmailOk

`func (o *Configure) GetAdminEmailOk() (*string, bool)`

GetAdminEmailOk returns a tuple with the AdminEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminEmail

`func (o *Configure) SetAdminEmail(v string)`

SetAdminEmail sets AdminEmail field to given value.

### HasAdminEmail

`func (o *Configure) HasAdminEmail() bool`

HasAdminEmail returns a boolean if a field has been set.

### GetAdminPassword

`func (o *Configure) GetAdminPassword() string`

GetAdminPassword returns the AdminPassword field if non-nil, zero value otherwise.

### GetAdminPasswordOk

`func (o *Configure) GetAdminPasswordOk() (*string, bool)`

GetAdminPasswordOk returns a tuple with the AdminPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPassword

`func (o *Configure) SetAdminPassword(v string)`

SetAdminPassword sets AdminPassword field to given value.

### HasAdminPassword

`func (o *Configure) HasAdminPassword() bool`

HasAdminPassword returns a boolean if a field has been set.

### GetAzureAdObjectId

`func (o *Configure) GetAzureAdObjectId() string`

GetAzureAdObjectId returns the AzureAdObjectId field if non-nil, zero value otherwise.

### GetAzureAdObjectIdOk

`func (o *Configure) GetAzureAdObjectIdOk() (*string, bool)`

GetAzureAdObjectIdOk returns a tuple with the AzureAdObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAdObjectId

`func (o *Configure) SetAzureAdObjectId(v string)`

SetAzureAdObjectId sets AzureAdObjectId field to given value.

### HasAzureAdObjectId

`func (o *Configure) HasAzureAdObjectId() bool`

HasAzureAdObjectId returns a boolean if a field has been set.

### GetAzureCloud

`func (o *Configure) GetAzureCloud() string`

GetAzureCloud returns the AzureCloud field if non-nil, zero value otherwise.

### GetAzureCloudOk

`func (o *Configure) GetAzureCloudOk() (*string, bool)`

GetAzureCloudOk returns a tuple with the AzureCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureCloud

`func (o *Configure) SetAzureCloud(v string)`

SetAzureCloud sets AzureCloud field to given value.

### HasAzureCloud

`func (o *Configure) HasAzureCloud() bool`

HasAzureCloud returns a boolean if a field has been set.

### GetCertData

`func (o *Configure) GetCertData() string`

GetCertData returns the CertData field if non-nil, zero value otherwise.

### GetCertDataOk

`func (o *Configure) GetCertDataOk() (*string, bool)`

GetCertDataOk returns a tuple with the CertData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertData

`func (o *Configure) SetCertData(v string)`

SetCertData sets CertData field to given value.

### HasCertData

`func (o *Configure) HasCertData() bool`

HasCertData returns a boolean if a field has been set.

### GetCertIssuerName

`func (o *Configure) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *Configure) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *Configure) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.

### HasCertIssuerName

`func (o *Configure) HasCertIssuerName() bool`

HasCertIssuerName returns a boolean if a field has been set.

### GetCertUsername

`func (o *Configure) GetCertUsername() string`

GetCertUsername returns the CertUsername field if non-nil, zero value otherwise.

### GetCertUsernameOk

`func (o *Configure) GetCertUsernameOk() (*string, bool)`

GetCertUsernameOk returns a tuple with the CertUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertUsername

`func (o *Configure) SetCertUsername(v string)`

SetCertUsername sets CertUsername field to given value.

### HasCertUsername

`func (o *Configure) HasCertUsername() bool`

HasCertUsername returns a boolean if a field has been set.

### GetDefaultLocationPrefix

`func (o *Configure) GetDefaultLocationPrefix() string`

GetDefaultLocationPrefix returns the DefaultLocationPrefix field if non-nil, zero value otherwise.

### GetDefaultLocationPrefixOk

`func (o *Configure) GetDefaultLocationPrefixOk() (*string, bool)`

GetDefaultLocationPrefixOk returns a tuple with the DefaultLocationPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLocationPrefix

`func (o *Configure) SetDefaultLocationPrefix(v string)`

SetDefaultLocationPrefix sets DefaultLocationPrefix field to given value.

### HasDefaultLocationPrefix

`func (o *Configure) HasDefaultLocationPrefix() bool`

HasDefaultLocationPrefix returns a boolean if a field has been set.

### GetDisablePafxfast

`func (o *Configure) GetDisablePafxfast() string`

GetDisablePafxfast returns the DisablePafxfast field if non-nil, zero value otherwise.

### GetDisablePafxfastOk

`func (o *Configure) GetDisablePafxfastOk() (*string, bool)`

GetDisablePafxfastOk returns a tuple with the DisablePafxfast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisablePafxfast

`func (o *Configure) SetDisablePafxfast(v string)`

SetDisablePafxfast sets DisablePafxfast field to given value.

### HasDisablePafxfast

`func (o *Configure) HasDisablePafxfast() bool`

HasDisablePafxfast returns a boolean if a field has been set.

### GetGatewaySpn

`func (o *Configure) GetGatewaySpn() string`

GetGatewaySpn returns the GatewaySpn field if non-nil, zero value otherwise.

### GetGatewaySpnOk

`func (o *Configure) GetGatewaySpnOk() (*string, bool)`

GetGatewaySpnOk returns a tuple with the GatewaySpn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySpn

`func (o *Configure) SetGatewaySpn(v string)`

SetGatewaySpn sets GatewaySpn field to given value.

### HasGatewaySpn

`func (o *Configure) HasGatewaySpn() bool`

HasGatewaySpn returns a boolean if a field has been set.

### GetGcpAudience

`func (o *Configure) GetGcpAudience() string`

GetGcpAudience returns the GcpAudience field if non-nil, zero value otherwise.

### GetGcpAudienceOk

`func (o *Configure) GetGcpAudienceOk() (*string, bool)`

GetGcpAudienceOk returns a tuple with the GcpAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpAudience

`func (o *Configure) SetGcpAudience(v string)`

SetGcpAudience sets GcpAudience field to given value.

### HasGcpAudience

`func (o *Configure) HasGcpAudience() bool`

HasGcpAudience returns a boolean if a field has been set.

### GetJson

`func (o *Configure) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *Configure) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *Configure) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *Configure) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sAuthConfigName

`func (o *Configure) GetK8sAuthConfigName() string`

GetK8sAuthConfigName returns the K8sAuthConfigName field if non-nil, zero value otherwise.

### GetK8sAuthConfigNameOk

`func (o *Configure) GetK8sAuthConfigNameOk() (*string, bool)`

GetK8sAuthConfigNameOk returns a tuple with the K8sAuthConfigName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAuthConfigName

`func (o *Configure) SetK8sAuthConfigName(v string)`

SetK8sAuthConfigName sets K8sAuthConfigName field to given value.

### HasK8sAuthConfigName

`func (o *Configure) HasK8sAuthConfigName() bool`

HasK8sAuthConfigName returns a boolean if a field has been set.

### GetKerberosToken

`func (o *Configure) GetKerberosToken() string`

GetKerberosToken returns the KerberosToken field if non-nil, zero value otherwise.

### GetKerberosTokenOk

`func (o *Configure) GetKerberosTokenOk() (*string, bool)`

GetKerberosTokenOk returns a tuple with the KerberosToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosToken

`func (o *Configure) SetKerberosToken(v string)`

SetKerberosToken sets KerberosToken field to given value.

### HasKerberosToken

`func (o *Configure) HasKerberosToken() bool`

HasKerberosToken returns a boolean if a field has been set.

### GetKerberosUsername

`func (o *Configure) GetKerberosUsername() string`

GetKerberosUsername returns the KerberosUsername field if non-nil, zero value otherwise.

### GetKerberosUsernameOk

`func (o *Configure) GetKerberosUsernameOk() (*string, bool)`

GetKerberosUsernameOk returns a tuple with the KerberosUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberosUsername

`func (o *Configure) SetKerberosUsername(v string)`

SetKerberosUsername sets KerberosUsername field to given value.

### HasKerberosUsername

`func (o *Configure) HasKerberosUsername() bool`

HasKerberosUsername returns a boolean if a field has been set.

### GetKeyData

`func (o *Configure) GetKeyData() string`

GetKeyData returns the KeyData field if non-nil, zero value otherwise.

### GetKeyDataOk

`func (o *Configure) GetKeyDataOk() (*string, bool)`

GetKeyDataOk returns a tuple with the KeyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyData

`func (o *Configure) SetKeyData(v string)`

SetKeyData sets KeyData field to given value.

### HasKeyData

`func (o *Configure) HasKeyData() bool`

HasKeyData returns a boolean if a field has been set.

### GetKeytabData

`func (o *Configure) GetKeytabData() string`

GetKeytabData returns the KeytabData field if non-nil, zero value otherwise.

### GetKeytabDataOk

`func (o *Configure) GetKeytabDataOk() (*string, bool)`

GetKeytabDataOk returns a tuple with the KeytabData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeytabData

`func (o *Configure) SetKeytabData(v string)`

SetKeytabData sets KeytabData field to given value.

### HasKeytabData

`func (o *Configure) HasKeytabData() bool`

HasKeytabData returns a boolean if a field has been set.

### GetKrb5ConfData

`func (o *Configure) GetKrb5ConfData() string`

GetKrb5ConfData returns the Krb5ConfData field if non-nil, zero value otherwise.

### GetKrb5ConfDataOk

`func (o *Configure) GetKrb5ConfDataOk() (*string, bool)`

GetKrb5ConfDataOk returns a tuple with the Krb5ConfData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKrb5ConfData

`func (o *Configure) SetKrb5ConfData(v string)`

SetKrb5ConfData sets Krb5ConfData field to given value.

### HasKrb5ConfData

`func (o *Configure) HasKrb5ConfData() bool`

HasKrb5ConfData returns a boolean if a field has been set.

### GetLegacySigningAlgName

`func (o *Configure) GetLegacySigningAlgName() bool`

GetLegacySigningAlgName returns the LegacySigningAlgName field if non-nil, zero value otherwise.

### GetLegacySigningAlgNameOk

`func (o *Configure) GetLegacySigningAlgNameOk() (*bool, bool)`

GetLegacySigningAlgNameOk returns a tuple with the LegacySigningAlgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacySigningAlgName

`func (o *Configure) SetLegacySigningAlgName(v bool)`

SetLegacySigningAlgName sets LegacySigningAlgName field to given value.

### HasLegacySigningAlgName

`func (o *Configure) HasLegacySigningAlgName() bool`

HasLegacySigningAlgName returns a boolean if a field has been set.

### GetOciAuthType

`func (o *Configure) GetOciAuthType() string`

GetOciAuthType returns the OciAuthType field if non-nil, zero value otherwise.

### GetOciAuthTypeOk

`func (o *Configure) GetOciAuthTypeOk() (*string, bool)`

GetOciAuthTypeOk returns a tuple with the OciAuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciAuthType

`func (o *Configure) SetOciAuthType(v string)`

SetOciAuthType sets OciAuthType field to given value.

### HasOciAuthType

`func (o *Configure) HasOciAuthType() bool`

HasOciAuthType returns a boolean if a field has been set.

### GetOciGroupOcid

`func (o *Configure) GetOciGroupOcid() []string`

GetOciGroupOcid returns the OciGroupOcid field if non-nil, zero value otherwise.

### GetOciGroupOcidOk

`func (o *Configure) GetOciGroupOcidOk() (*[]string, bool)`

GetOciGroupOcidOk returns a tuple with the OciGroupOcid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOciGroupOcid

`func (o *Configure) SetOciGroupOcid(v []string)`

SetOciGroupOcid sets OciGroupOcid field to given value.

### HasOciGroupOcid

`func (o *Configure) HasOciGroupOcid() bool`

HasOciGroupOcid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



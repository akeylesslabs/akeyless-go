# Auth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | Access ID | [optional] 
**AccessKey** | Pointer to **string** | Access key (relevant only for access-type&#x3D;access_key) | [optional] 
**AccessType** | Pointer to **string** | Access Type (access_key/password/saml/ldap/k8s/azure_ad/aws_iam/universal_identity/jwt/gcp/k8s) | [optional] [default to "access_key"]
**AdminEmail** | Pointer to **string** | Email (relevant only for access-type&#x3D;password) | [optional] 
**AdminPassword** | Pointer to **string** | Password (relevant only for access-type&#x3D;password) | [optional] 
**CloudId** | Pointer to **string** | The cloud identity (relevant only for access-type&#x3D;azure_ad,aws_iam,gcp) | [optional] 
**Debug** | Pointer to **bool** |  | [optional] 
**GcpAudience** | Pointer to **string** | GCP JWT audience | [optional] 
**Jwt** | Pointer to **string** | The Json Web Token (relevant only for access-type&#x3D;jwt/oidc) | [optional] 
**K8sAuthConfigName** | Pointer to **string** | The K8S Auth config name (relevant only for access-type&#x3D;k8s) | [optional] 
**K8sServiceAccountToken** | Pointer to **string** | The K8S service account token. (relevant only for access-type&#x3D;k8s) | [optional] 
**LdapPassword** | Pointer to **string** | LDAP password (relevant only for access-type&#x3D;ldap) | [optional] 
**LdapUsername** | Pointer to **string** | LDAP username (relevant only for access-type&#x3D;ldap) | [optional] 
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



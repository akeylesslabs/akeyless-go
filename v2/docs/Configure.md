# Configure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessId** | Pointer to **string** | Access ID | [optional] 
**AccessKey** | Pointer to **string** | Access Key | [optional] 
**AccessType** | Pointer to **string** | Access Type (access_key/password/azure_ad/saml/oidc/aws_iam/gcp/k8s) | [optional] [default to "access_key"]
**AdminEmail** | Pointer to **string** | Email (relevant only for access-type&#x3D;password) | [optional] 
**AdminPassword** | Pointer to **string** | Password (relevant only for access-type&#x3D;password) | [optional] 
**AzureAdObjectId** | Pointer to **string** | Azure Active Directory ObjectId (relevant only for access-type&#x3D;azure_ad) | [optional] 
**CertData** | Pointer to **string** | Certificate data encoded in base64. Used if file was not provided. (relevant only for access-type&#x3D;cert in Curl Context) | [optional] 
**GcpAudience** | Pointer to **string** | GCP JWT audience | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**K8sAuthConfigName** | Pointer to **string** | The K8S Auth config name (relevant only for access-type&#x3D;k8s) | [optional] 
**KeyData** | Pointer to **string** | Private key data encoded in base64. Used if file was not provided.(relevant only for access-type&#x3D;cert in Curl Context) | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DsCreateK8s

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sAllowedNamespaces** | Pointer to **string** | Comma-separated list of allowed K8S namespaces for the generated ServiceAccount (relevant only for k8s-service-account-type&#x3D;dynamic) | [optional] 
**K8sClusterCaCert** | Pointer to **string** | K8S cluster CA certificate | [optional] 
**K8sClusterEndpoint** | Pointer to **string** | K8S cluster URL endpoint | [optional] 
**K8sClusterToken** | Pointer to **string** | K8S cluster Bearer token | [optional] 
**K8sNamespace** | Pointer to **string** | K8S Namespace where the ServiceAccount exists. | [optional] 
**K8sPredefinedRoleName** | Pointer to **string** | The pre-existing Role or ClusterRole name to bind the generated ServiceAccount to (relevant only for k8s-service-account-type&#x3D;dynamic) | [optional] 
**K8sPredefinedRoleType** | Pointer to **string** | Specifies the type of the pre-existing K8S role [Role, ClusterRole] (relevant only for k8s-service-account-type&#x3D;dynamic) | [optional] 
**K8sRolebindingYamlDef** | Pointer to **string** | Path to yaml file that contains definitions of K8S role and role binding (relevant only for k8s-service-account-type&#x3D;dynamic) | [optional] 
**K8sServiceAccount** | Pointer to **string** | K8S ServiceAccount to extract token from. | [optional] 
**K8sServiceAccountType** | Pointer to **string** | K8S ServiceAccount type [fixed, dynamic]. | [optional] 
**Name** | **string** | Dynamic secret name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SecureAccessAllowPortForwading** | Pointer to **bool** | Enable Port forwarding while using CLI access | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessClusterEndpoint** | Pointer to **string** | The K8s cluster endpoint URL | [optional] 
**SecureAccessDashboardUrl** | Pointer to **string** | The K8s dashboard url | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion | [optional] [default to false]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwServiceAccount** | Pointer to **bool** | Use the GW&#39;s service account | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDsCreateK8s

`func NewDsCreateK8s(name string, ) *DsCreateK8s`

NewDsCreateK8s instantiates a new DsCreateK8s object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsCreateK8sWithDefaults

`func NewDsCreateK8sWithDefaults() *DsCreateK8s`

NewDsCreateK8sWithDefaults instantiates a new DsCreateK8s object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DsCreateK8s) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsCreateK8s) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsCreateK8s) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsCreateK8s) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DsCreateK8s) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsCreateK8s) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsCreateK8s) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsCreateK8s) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sAllowedNamespaces

`func (o *DsCreateK8s) GetK8sAllowedNamespaces() string`

GetK8sAllowedNamespaces returns the K8sAllowedNamespaces field if non-nil, zero value otherwise.

### GetK8sAllowedNamespacesOk

`func (o *DsCreateK8s) GetK8sAllowedNamespacesOk() (*string, bool)`

GetK8sAllowedNamespacesOk returns a tuple with the K8sAllowedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAllowedNamespaces

`func (o *DsCreateK8s) SetK8sAllowedNamespaces(v string)`

SetK8sAllowedNamespaces sets K8sAllowedNamespaces field to given value.

### HasK8sAllowedNamespaces

`func (o *DsCreateK8s) HasK8sAllowedNamespaces() bool`

HasK8sAllowedNamespaces returns a boolean if a field has been set.

### GetK8sClusterCaCert

`func (o *DsCreateK8s) GetK8sClusterCaCert() string`

GetK8sClusterCaCert returns the K8sClusterCaCert field if non-nil, zero value otherwise.

### GetK8sClusterCaCertOk

`func (o *DsCreateK8s) GetK8sClusterCaCertOk() (*string, bool)`

GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCert

`func (o *DsCreateK8s) SetK8sClusterCaCert(v string)`

SetK8sClusterCaCert sets K8sClusterCaCert field to given value.

### HasK8sClusterCaCert

`func (o *DsCreateK8s) HasK8sClusterCaCert() bool`

HasK8sClusterCaCert returns a boolean if a field has been set.

### GetK8sClusterEndpoint

`func (o *DsCreateK8s) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *DsCreateK8s) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *DsCreateK8s) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.

### HasK8sClusterEndpoint

`func (o *DsCreateK8s) HasK8sClusterEndpoint() bool`

HasK8sClusterEndpoint returns a boolean if a field has been set.

### GetK8sClusterToken

`func (o *DsCreateK8s) GetK8sClusterToken() string`

GetK8sClusterToken returns the K8sClusterToken field if non-nil, zero value otherwise.

### GetK8sClusterTokenOk

`func (o *DsCreateK8s) GetK8sClusterTokenOk() (*string, bool)`

GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterToken

`func (o *DsCreateK8s) SetK8sClusterToken(v string)`

SetK8sClusterToken sets K8sClusterToken field to given value.

### HasK8sClusterToken

`func (o *DsCreateK8s) HasK8sClusterToken() bool`

HasK8sClusterToken returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *DsCreateK8s) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *DsCreateK8s) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *DsCreateK8s) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *DsCreateK8s) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetK8sPredefinedRoleName

`func (o *DsCreateK8s) GetK8sPredefinedRoleName() string`

GetK8sPredefinedRoleName returns the K8sPredefinedRoleName field if non-nil, zero value otherwise.

### GetK8sPredefinedRoleNameOk

`func (o *DsCreateK8s) GetK8sPredefinedRoleNameOk() (*string, bool)`

GetK8sPredefinedRoleNameOk returns a tuple with the K8sPredefinedRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPredefinedRoleName

`func (o *DsCreateK8s) SetK8sPredefinedRoleName(v string)`

SetK8sPredefinedRoleName sets K8sPredefinedRoleName field to given value.

### HasK8sPredefinedRoleName

`func (o *DsCreateK8s) HasK8sPredefinedRoleName() bool`

HasK8sPredefinedRoleName returns a boolean if a field has been set.

### GetK8sPredefinedRoleType

`func (o *DsCreateK8s) GetK8sPredefinedRoleType() string`

GetK8sPredefinedRoleType returns the K8sPredefinedRoleType field if non-nil, zero value otherwise.

### GetK8sPredefinedRoleTypeOk

`func (o *DsCreateK8s) GetK8sPredefinedRoleTypeOk() (*string, bool)`

GetK8sPredefinedRoleTypeOk returns a tuple with the K8sPredefinedRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPredefinedRoleType

`func (o *DsCreateK8s) SetK8sPredefinedRoleType(v string)`

SetK8sPredefinedRoleType sets K8sPredefinedRoleType field to given value.

### HasK8sPredefinedRoleType

`func (o *DsCreateK8s) HasK8sPredefinedRoleType() bool`

HasK8sPredefinedRoleType returns a boolean if a field has been set.

### GetK8sRolebindingYamlDef

`func (o *DsCreateK8s) GetK8sRolebindingYamlDef() string`

GetK8sRolebindingYamlDef returns the K8sRolebindingYamlDef field if non-nil, zero value otherwise.

### GetK8sRolebindingYamlDefOk

`func (o *DsCreateK8s) GetK8sRolebindingYamlDefOk() (*string, bool)`

GetK8sRolebindingYamlDefOk returns a tuple with the K8sRolebindingYamlDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sRolebindingYamlDef

`func (o *DsCreateK8s) SetK8sRolebindingYamlDef(v string)`

SetK8sRolebindingYamlDef sets K8sRolebindingYamlDef field to given value.

### HasK8sRolebindingYamlDef

`func (o *DsCreateK8s) HasK8sRolebindingYamlDef() bool`

HasK8sRolebindingYamlDef returns a boolean if a field has been set.

### GetK8sServiceAccount

`func (o *DsCreateK8s) GetK8sServiceAccount() string`

GetK8sServiceAccount returns the K8sServiceAccount field if non-nil, zero value otherwise.

### GetK8sServiceAccountOk

`func (o *DsCreateK8s) GetK8sServiceAccountOk() (*string, bool)`

GetK8sServiceAccountOk returns a tuple with the K8sServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceAccount

`func (o *DsCreateK8s) SetK8sServiceAccount(v string)`

SetK8sServiceAccount sets K8sServiceAccount field to given value.

### HasK8sServiceAccount

`func (o *DsCreateK8s) HasK8sServiceAccount() bool`

HasK8sServiceAccount returns a boolean if a field has been set.

### GetK8sServiceAccountType

`func (o *DsCreateK8s) GetK8sServiceAccountType() string`

GetK8sServiceAccountType returns the K8sServiceAccountType field if non-nil, zero value otherwise.

### GetK8sServiceAccountTypeOk

`func (o *DsCreateK8s) GetK8sServiceAccountTypeOk() (*string, bool)`

GetK8sServiceAccountTypeOk returns a tuple with the K8sServiceAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceAccountType

`func (o *DsCreateK8s) SetK8sServiceAccountType(v string)`

SetK8sServiceAccountType sets K8sServiceAccountType field to given value.

### HasK8sServiceAccountType

`func (o *DsCreateK8s) HasK8sServiceAccountType() bool`

HasK8sServiceAccountType returns a boolean if a field has been set.

### GetName

`func (o *DsCreateK8s) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsCreateK8s) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsCreateK8s) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DsCreateK8s) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DsCreateK8s) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DsCreateK8s) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DsCreateK8s) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessAllowPortForwading

`func (o *DsCreateK8s) GetSecureAccessAllowPortForwading() bool`

GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field if non-nil, zero value otherwise.

### GetSecureAccessAllowPortForwadingOk

`func (o *DsCreateK8s) GetSecureAccessAllowPortForwadingOk() (*bool, bool)`

GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowPortForwading

`func (o *DsCreateK8s) SetSecureAccessAllowPortForwading(v bool)`

SetSecureAccessAllowPortForwading sets SecureAccessAllowPortForwading field to given value.

### HasSecureAccessAllowPortForwading

`func (o *DsCreateK8s) HasSecureAccessAllowPortForwading() bool`

HasSecureAccessAllowPortForwading returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *DsCreateK8s) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *DsCreateK8s) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *DsCreateK8s) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *DsCreateK8s) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessClusterEndpoint

`func (o *DsCreateK8s) GetSecureAccessClusterEndpoint() string`

GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field if non-nil, zero value otherwise.

### GetSecureAccessClusterEndpointOk

`func (o *DsCreateK8s) GetSecureAccessClusterEndpointOk() (*string, bool)`

GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessClusterEndpoint

`func (o *DsCreateK8s) SetSecureAccessClusterEndpoint(v string)`

SetSecureAccessClusterEndpoint sets SecureAccessClusterEndpoint field to given value.

### HasSecureAccessClusterEndpoint

`func (o *DsCreateK8s) HasSecureAccessClusterEndpoint() bool`

HasSecureAccessClusterEndpoint returns a boolean if a field has been set.

### GetSecureAccessDashboardUrl

`func (o *DsCreateK8s) GetSecureAccessDashboardUrl() string`

GetSecureAccessDashboardUrl returns the SecureAccessDashboardUrl field if non-nil, zero value otherwise.

### GetSecureAccessDashboardUrlOk

`func (o *DsCreateK8s) GetSecureAccessDashboardUrlOk() (*string, bool)`

GetSecureAccessDashboardUrlOk returns a tuple with the SecureAccessDashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDashboardUrl

`func (o *DsCreateK8s) SetSecureAccessDashboardUrl(v string)`

SetSecureAccessDashboardUrl sets SecureAccessDashboardUrl field to given value.

### HasSecureAccessDashboardUrl

`func (o *DsCreateK8s) HasSecureAccessDashboardUrl() bool`

HasSecureAccessDashboardUrl returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *DsCreateK8s) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *DsCreateK8s) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *DsCreateK8s) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *DsCreateK8s) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *DsCreateK8s) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *DsCreateK8s) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *DsCreateK8s) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *DsCreateK8s) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *DsCreateK8s) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *DsCreateK8s) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *DsCreateK8s) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *DsCreateK8s) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *DsCreateK8s) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *DsCreateK8s) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *DsCreateK8s) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *DsCreateK8s) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *DsCreateK8s) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsCreateK8s) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsCreateK8s) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsCreateK8s) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsCreateK8s) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsCreateK8s) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsCreateK8s) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsCreateK8s) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsCreateK8s) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsCreateK8s) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsCreateK8s) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsCreateK8s) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DsCreateK8s) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsCreateK8s) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsCreateK8s) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsCreateK8s) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwServiceAccount

`func (o *DsCreateK8s) GetUseGwServiceAccount() bool`

GetUseGwServiceAccount returns the UseGwServiceAccount field if non-nil, zero value otherwise.

### GetUseGwServiceAccountOk

`func (o *DsCreateK8s) GetUseGwServiceAccountOk() (*bool, bool)`

GetUseGwServiceAccountOk returns a tuple with the UseGwServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwServiceAccount

`func (o *DsCreateK8s) SetUseGwServiceAccount(v bool)`

SetUseGwServiceAccount sets UseGwServiceAccount field to given value.

### HasUseGwServiceAccount

`func (o *DsCreateK8s) HasUseGwServiceAccount() bool`

HasUseGwServiceAccount returns a boolean if a field has been set.

### GetUserTtl

`func (o *DsCreateK8s) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DsCreateK8s) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DsCreateK8s) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DsCreateK8s) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



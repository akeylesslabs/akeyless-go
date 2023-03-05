# GatewayUpdateProducerNativeK8S

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
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
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
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayUpdateProducerNativeK8S

`func NewGatewayUpdateProducerNativeK8S(name string, ) *GatewayUpdateProducerNativeK8S`

NewGatewayUpdateProducerNativeK8S instantiates a new GatewayUpdateProducerNativeK8S object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerNativeK8SWithDefaults

`func NewGatewayUpdateProducerNativeK8SWithDefaults() *GatewayUpdateProducerNativeK8S`

NewGatewayUpdateProducerNativeK8SWithDefaults instantiates a new GatewayUpdateProducerNativeK8S object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *GatewayUpdateProducerNativeK8S) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerNativeK8S) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerNativeK8S) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerNativeK8S) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerNativeK8S) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerNativeK8S) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerNativeK8S) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerNativeK8S) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sAllowedNamespaces

`func (o *GatewayUpdateProducerNativeK8S) GetK8sAllowedNamespaces() string`

GetK8sAllowedNamespaces returns the K8sAllowedNamespaces field if non-nil, zero value otherwise.

### GetK8sAllowedNamespacesOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sAllowedNamespacesOk() (*string, bool)`

GetK8sAllowedNamespacesOk returns a tuple with the K8sAllowedNamespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sAllowedNamespaces

`func (o *GatewayUpdateProducerNativeK8S) SetK8sAllowedNamespaces(v string)`

SetK8sAllowedNamespaces sets K8sAllowedNamespaces field to given value.

### HasK8sAllowedNamespaces

`func (o *GatewayUpdateProducerNativeK8S) HasK8sAllowedNamespaces() bool`

HasK8sAllowedNamespaces returns a boolean if a field has been set.

### GetK8sClusterCaCert

`func (o *GatewayUpdateProducerNativeK8S) GetK8sClusterCaCert() string`

GetK8sClusterCaCert returns the K8sClusterCaCert field if non-nil, zero value otherwise.

### GetK8sClusterCaCertOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sClusterCaCertOk() (*string, bool)`

GetK8sClusterCaCertOk returns a tuple with the K8sClusterCaCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterCaCert

`func (o *GatewayUpdateProducerNativeK8S) SetK8sClusterCaCert(v string)`

SetK8sClusterCaCert sets K8sClusterCaCert field to given value.

### HasK8sClusterCaCert

`func (o *GatewayUpdateProducerNativeK8S) HasK8sClusterCaCert() bool`

HasK8sClusterCaCert returns a boolean if a field has been set.

### GetK8sClusterEndpoint

`func (o *GatewayUpdateProducerNativeK8S) GetK8sClusterEndpoint() string`

GetK8sClusterEndpoint returns the K8sClusterEndpoint field if non-nil, zero value otherwise.

### GetK8sClusterEndpointOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sClusterEndpointOk() (*string, bool)`

GetK8sClusterEndpointOk returns a tuple with the K8sClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterEndpoint

`func (o *GatewayUpdateProducerNativeK8S) SetK8sClusterEndpoint(v string)`

SetK8sClusterEndpoint sets K8sClusterEndpoint field to given value.

### HasK8sClusterEndpoint

`func (o *GatewayUpdateProducerNativeK8S) HasK8sClusterEndpoint() bool`

HasK8sClusterEndpoint returns a boolean if a field has been set.

### GetK8sClusterToken

`func (o *GatewayUpdateProducerNativeK8S) GetK8sClusterToken() string`

GetK8sClusterToken returns the K8sClusterToken field if non-nil, zero value otherwise.

### GetK8sClusterTokenOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sClusterTokenOk() (*string, bool)`

GetK8sClusterTokenOk returns a tuple with the K8sClusterToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClusterToken

`func (o *GatewayUpdateProducerNativeK8S) SetK8sClusterToken(v string)`

SetK8sClusterToken sets K8sClusterToken field to given value.

### HasK8sClusterToken

`func (o *GatewayUpdateProducerNativeK8S) HasK8sClusterToken() bool`

HasK8sClusterToken returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *GatewayUpdateProducerNativeK8S) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *GatewayUpdateProducerNativeK8S) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *GatewayUpdateProducerNativeK8S) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetK8sPredefinedRoleName

`func (o *GatewayUpdateProducerNativeK8S) GetK8sPredefinedRoleName() string`

GetK8sPredefinedRoleName returns the K8sPredefinedRoleName field if non-nil, zero value otherwise.

### GetK8sPredefinedRoleNameOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sPredefinedRoleNameOk() (*string, bool)`

GetK8sPredefinedRoleNameOk returns a tuple with the K8sPredefinedRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPredefinedRoleName

`func (o *GatewayUpdateProducerNativeK8S) SetK8sPredefinedRoleName(v string)`

SetK8sPredefinedRoleName sets K8sPredefinedRoleName field to given value.

### HasK8sPredefinedRoleName

`func (o *GatewayUpdateProducerNativeK8S) HasK8sPredefinedRoleName() bool`

HasK8sPredefinedRoleName returns a boolean if a field has been set.

### GetK8sPredefinedRoleType

`func (o *GatewayUpdateProducerNativeK8S) GetK8sPredefinedRoleType() string`

GetK8sPredefinedRoleType returns the K8sPredefinedRoleType field if non-nil, zero value otherwise.

### GetK8sPredefinedRoleTypeOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sPredefinedRoleTypeOk() (*string, bool)`

GetK8sPredefinedRoleTypeOk returns a tuple with the K8sPredefinedRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPredefinedRoleType

`func (o *GatewayUpdateProducerNativeK8S) SetK8sPredefinedRoleType(v string)`

SetK8sPredefinedRoleType sets K8sPredefinedRoleType field to given value.

### HasK8sPredefinedRoleType

`func (o *GatewayUpdateProducerNativeK8S) HasK8sPredefinedRoleType() bool`

HasK8sPredefinedRoleType returns a boolean if a field has been set.

### GetK8sRolebindingYamlDef

`func (o *GatewayUpdateProducerNativeK8S) GetK8sRolebindingYamlDef() string`

GetK8sRolebindingYamlDef returns the K8sRolebindingYamlDef field if non-nil, zero value otherwise.

### GetK8sRolebindingYamlDefOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sRolebindingYamlDefOk() (*string, bool)`

GetK8sRolebindingYamlDefOk returns a tuple with the K8sRolebindingYamlDef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sRolebindingYamlDef

`func (o *GatewayUpdateProducerNativeK8S) SetK8sRolebindingYamlDef(v string)`

SetK8sRolebindingYamlDef sets K8sRolebindingYamlDef field to given value.

### HasK8sRolebindingYamlDef

`func (o *GatewayUpdateProducerNativeK8S) HasK8sRolebindingYamlDef() bool`

HasK8sRolebindingYamlDef returns a boolean if a field has been set.

### GetK8sServiceAccount

`func (o *GatewayUpdateProducerNativeK8S) GetK8sServiceAccount() string`

GetK8sServiceAccount returns the K8sServiceAccount field if non-nil, zero value otherwise.

### GetK8sServiceAccountOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sServiceAccountOk() (*string, bool)`

GetK8sServiceAccountOk returns a tuple with the K8sServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceAccount

`func (o *GatewayUpdateProducerNativeK8S) SetK8sServiceAccount(v string)`

SetK8sServiceAccount sets K8sServiceAccount field to given value.

### HasK8sServiceAccount

`func (o *GatewayUpdateProducerNativeK8S) HasK8sServiceAccount() bool`

HasK8sServiceAccount returns a boolean if a field has been set.

### GetK8sServiceAccountType

`func (o *GatewayUpdateProducerNativeK8S) GetK8sServiceAccountType() string`

GetK8sServiceAccountType returns the K8sServiceAccountType field if non-nil, zero value otherwise.

### GetK8sServiceAccountTypeOk

`func (o *GatewayUpdateProducerNativeK8S) GetK8sServiceAccountTypeOk() (*string, bool)`

GetK8sServiceAccountTypeOk returns a tuple with the K8sServiceAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sServiceAccountType

`func (o *GatewayUpdateProducerNativeK8S) SetK8sServiceAccountType(v string)`

SetK8sServiceAccountType sets K8sServiceAccountType field to given value.

### HasK8sServiceAccountType

`func (o *GatewayUpdateProducerNativeK8S) HasK8sServiceAccountType() bool`

HasK8sServiceAccountType returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerNativeK8S) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerNativeK8S) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerNativeK8S) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerNativeK8S) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerNativeK8S) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerNativeK8S) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerNativeK8S) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerNativeK8S) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerNativeK8S) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerNativeK8S) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerNativeK8S) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSecureAccessAllowPortForwading

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessAllowPortForwading() bool`

GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field if non-nil, zero value otherwise.

### GetSecureAccessAllowPortForwadingOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessAllowPortForwadingOk() (*bool, bool)`

GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowPortForwading

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessAllowPortForwading(v bool)`

SetSecureAccessAllowPortForwading sets SecureAccessAllowPortForwading field to given value.

### HasSecureAccessAllowPortForwading

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessAllowPortForwading() bool`

HasSecureAccessAllowPortForwading returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessClusterEndpoint

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessClusterEndpoint() string`

GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field if non-nil, zero value otherwise.

### GetSecureAccessClusterEndpointOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessClusterEndpointOk() (*string, bool)`

GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessClusterEndpoint

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessClusterEndpoint(v string)`

SetSecureAccessClusterEndpoint sets SecureAccessClusterEndpoint field to given value.

### HasSecureAccessClusterEndpoint

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessClusterEndpoint() bool`

HasSecureAccessClusterEndpoint returns a boolean if a field has been set.

### GetSecureAccessDashboardUrl

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessDashboardUrl() string`

GetSecureAccessDashboardUrl returns the SecureAccessDashboardUrl field if non-nil, zero value otherwise.

### GetSecureAccessDashboardUrlOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessDashboardUrlOk() (*string, bool)`

GetSecureAccessDashboardUrlOk returns a tuple with the SecureAccessDashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDashboardUrl

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessDashboardUrl(v string)`

SetSecureAccessDashboardUrl sets SecureAccessDashboardUrl field to given value.

### HasSecureAccessDashboardUrl

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessDashboardUrl() bool`

HasSecureAccessDashboardUrl returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *GatewayUpdateProducerNativeK8S) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *GatewayUpdateProducerNativeK8S) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *GatewayUpdateProducerNativeK8S) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerNativeK8S) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerNativeK8S) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerNativeK8S) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerNativeK8S) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerNativeK8S) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerNativeK8S) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerNativeK8S) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerNativeK8S) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerNativeK8S) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerNativeK8S) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerNativeK8S) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerNativeK8S) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerNativeK8S) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerNativeK8S) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerNativeK8S) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerNativeK8S) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerNativeK8S) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerNativeK8S) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerNativeK8S) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerNativeK8S) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



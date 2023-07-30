# UpdateItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**CertFileData** | Pointer to **string** | PEM Certificate in a Base64 format. Used for updating RSA keys&#39; certificates. | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_metadata"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Current item name | 
**NewMetadata** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_metadata"]
**NewName** | Pointer to **string** | New item name | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotateAfterDisconnect** | Pointer to **string** | Rotate the value of the secret after SRA session ends [true/false] (relevant only for Rotated-secret) | [optional] [default to "false"]
**SecureAccessAddHost** | Pointer to **[]string** | List of the new hosts that will be attached to SRA servers host | [optional] 
**SecureAccessAllowExternalUser** | Pointer to **string** | Allow providing external user for a domain users [true/false] | [optional] 
**SecureAccessAllowPortForwading** | Pointer to **bool** | Enable Port forwarding while using CLI access (relevant only for EKS/GKE/K8s Dynamic-Secret) | [optional] 
**SecureAccessAwsAccountId** | Pointer to **string** | The AWS account id (relevant only for aws) | [optional] 
**SecureAccessAwsNativeCli** | Pointer to **bool** | The AWS native cli (relevant only for aws) | [optional] 
**SecureAccessAwsRegion** | Pointer to **string** | The AWS region (relevant only for aws) | [optional] 
**SecureAccessBastionApi** | Pointer to **string** | Bastion&#39;s SSH control API endpoint. E.g. https://my.bastion:9900 (relevant only for ssh cert issuer) | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessBastionSsh** | Pointer to **string** | Bastion&#39;s SSH server. E.g. my.bastion:22 (relevant only for ssh cert issuer) | [optional] 
**SecureAccessClusterEndpoint** | Pointer to **string** | The K8s cluster endpoint URL (relevant only for EKS/GKE/K8s Dynamic-Secret) | [optional] 
**SecureAccessDashboardUrl** | Pointer to **string** | The K8s dashboard url (relevant only for k8s) | [optional] 
**SecureAccessDbName** | Pointer to **string** | The DB name (relevant only for DB Dynamic-Secret) | [optional] 
**SecureAccessDbSchema** | Pointer to **string** | The DB schema (relevant only for DB Dynamic-Secret) | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user (relevant only for RDP Dynamic-Secret) | [optional] 
**SecureAccessRdpUser** | Pointer to **string** | Override the RDP Domain username | [optional] 
**SecureAccessRmHost** | Pointer to **[]string** | List of the existent hosts that will be removed from SRA servers host | [optional] 
**SecureAccessSshCreds** | Pointer to **string** | Secret values contains SSH Credentials, either Private Key or Password [password/private-key] (relevant only for Static-Secret or Rotated-secret) | [optional] 
**SecureAccessSshCredsUser** | Pointer to **string** | SSH username to connect to target server, must be in &#39;Allowed Users&#39; list (relevant only for ssh cert issuer) | [optional] 
**SecureAccessUrl** | Pointer to **string** | Destination URL to inject secrets | [optional] 
**SecureAccessUseInternalBastion** | Pointer to **bool** | Use internal SSH Bastion | [optional] 
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion | [optional] [default to false]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateItem

`func NewUpdateItem(name string, ) *UpdateItem`

NewUpdateItem instantiates a new UpdateItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateItemWithDefaults

`func NewUpdateItemWithDefaults() *UpdateItem`

NewUpdateItemWithDefaults instantiates a new UpdateItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *UpdateItem) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *UpdateItem) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *UpdateItem) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *UpdateItem) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAddTag

`func (o *UpdateItem) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateItem) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateItem) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateItem) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetCertFileData

`func (o *UpdateItem) GetCertFileData() string`

GetCertFileData returns the CertFileData field if non-nil, zero value otherwise.

### GetCertFileDataOk

`func (o *UpdateItem) GetCertFileDataOk() (*string, bool)`

GetCertFileDataOk returns a tuple with the CertFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFileData

`func (o *UpdateItem) SetCertFileData(v string)`

SetCertFileData sets CertFileData field to given value.

### HasCertFileData

`func (o *UpdateItem) HasCertFileData() bool`

HasCertFileData returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *UpdateItem) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *UpdateItem) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *UpdateItem) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *UpdateItem) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *UpdateItem) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateItem) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateItem) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateItem) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *UpdateItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateItem) SetName(v string)`

SetName sets Name field to given value.


### GetNewMetadata

`func (o *UpdateItem) GetNewMetadata() string`

GetNewMetadata returns the NewMetadata field if non-nil, zero value otherwise.

### GetNewMetadataOk

`func (o *UpdateItem) GetNewMetadataOk() (*string, bool)`

GetNewMetadataOk returns a tuple with the NewMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMetadata

`func (o *UpdateItem) SetNewMetadata(v string)`

SetNewMetadata sets NewMetadata field to given value.

### HasNewMetadata

`func (o *UpdateItem) HasNewMetadata() bool`

HasNewMetadata returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateItem) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateItem) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateItem) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateItem) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdateItem) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateItem) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateItem) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateItem) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotateAfterDisconnect

`func (o *UpdateItem) GetRotateAfterDisconnect() string`

GetRotateAfterDisconnect returns the RotateAfterDisconnect field if non-nil, zero value otherwise.

### GetRotateAfterDisconnectOk

`func (o *UpdateItem) GetRotateAfterDisconnectOk() (*string, bool)`

GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAfterDisconnect

`func (o *UpdateItem) SetRotateAfterDisconnect(v string)`

SetRotateAfterDisconnect sets RotateAfterDisconnect field to given value.

### HasRotateAfterDisconnect

`func (o *UpdateItem) HasRotateAfterDisconnect() bool`

HasRotateAfterDisconnect returns a boolean if a field has been set.

### GetSecureAccessAddHost

`func (o *UpdateItem) GetSecureAccessAddHost() []string`

GetSecureAccessAddHost returns the SecureAccessAddHost field if non-nil, zero value otherwise.

### GetSecureAccessAddHostOk

`func (o *UpdateItem) GetSecureAccessAddHostOk() (*[]string, bool)`

GetSecureAccessAddHostOk returns a tuple with the SecureAccessAddHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAddHost

`func (o *UpdateItem) SetSecureAccessAddHost(v []string)`

SetSecureAccessAddHost sets SecureAccessAddHost field to given value.

### HasSecureAccessAddHost

`func (o *UpdateItem) HasSecureAccessAddHost() bool`

HasSecureAccessAddHost returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *UpdateItem) GetSecureAccessAllowExternalUser() string`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *UpdateItem) GetSecureAccessAllowExternalUserOk() (*string, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *UpdateItem) SetSecureAccessAllowExternalUser(v string)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *UpdateItem) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessAllowPortForwading

`func (o *UpdateItem) GetSecureAccessAllowPortForwading() bool`

GetSecureAccessAllowPortForwading returns the SecureAccessAllowPortForwading field if non-nil, zero value otherwise.

### GetSecureAccessAllowPortForwadingOk

`func (o *UpdateItem) GetSecureAccessAllowPortForwadingOk() (*bool, bool)`

GetSecureAccessAllowPortForwadingOk returns a tuple with the SecureAccessAllowPortForwading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowPortForwading

`func (o *UpdateItem) SetSecureAccessAllowPortForwading(v bool)`

SetSecureAccessAllowPortForwading sets SecureAccessAllowPortForwading field to given value.

### HasSecureAccessAllowPortForwading

`func (o *UpdateItem) HasSecureAccessAllowPortForwading() bool`

HasSecureAccessAllowPortForwading returns a boolean if a field has been set.

### GetSecureAccessAwsAccountId

`func (o *UpdateItem) GetSecureAccessAwsAccountId() string`

GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field if non-nil, zero value otherwise.

### GetSecureAccessAwsAccountIdOk

`func (o *UpdateItem) GetSecureAccessAwsAccountIdOk() (*string, bool)`

GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsAccountId

`func (o *UpdateItem) SetSecureAccessAwsAccountId(v string)`

SetSecureAccessAwsAccountId sets SecureAccessAwsAccountId field to given value.

### HasSecureAccessAwsAccountId

`func (o *UpdateItem) HasSecureAccessAwsAccountId() bool`

HasSecureAccessAwsAccountId returns a boolean if a field has been set.

### GetSecureAccessAwsNativeCli

`func (o *UpdateItem) GetSecureAccessAwsNativeCli() bool`

GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field if non-nil, zero value otherwise.

### GetSecureAccessAwsNativeCliOk

`func (o *UpdateItem) GetSecureAccessAwsNativeCliOk() (*bool, bool)`

GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsNativeCli

`func (o *UpdateItem) SetSecureAccessAwsNativeCli(v bool)`

SetSecureAccessAwsNativeCli sets SecureAccessAwsNativeCli field to given value.

### HasSecureAccessAwsNativeCli

`func (o *UpdateItem) HasSecureAccessAwsNativeCli() bool`

HasSecureAccessAwsNativeCli returns a boolean if a field has been set.

### GetSecureAccessAwsRegion

`func (o *UpdateItem) GetSecureAccessAwsRegion() string`

GetSecureAccessAwsRegion returns the SecureAccessAwsRegion field if non-nil, zero value otherwise.

### GetSecureAccessAwsRegionOk

`func (o *UpdateItem) GetSecureAccessAwsRegionOk() (*string, bool)`

GetSecureAccessAwsRegionOk returns a tuple with the SecureAccessAwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsRegion

`func (o *UpdateItem) SetSecureAccessAwsRegion(v string)`

SetSecureAccessAwsRegion sets SecureAccessAwsRegion field to given value.

### HasSecureAccessAwsRegion

`func (o *UpdateItem) HasSecureAccessAwsRegion() bool`

HasSecureAccessAwsRegion returns a boolean if a field has been set.

### GetSecureAccessBastionApi

`func (o *UpdateItem) GetSecureAccessBastionApi() string`

GetSecureAccessBastionApi returns the SecureAccessBastionApi field if non-nil, zero value otherwise.

### GetSecureAccessBastionApiOk

`func (o *UpdateItem) GetSecureAccessBastionApiOk() (*string, bool)`

GetSecureAccessBastionApiOk returns a tuple with the SecureAccessBastionApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionApi

`func (o *UpdateItem) SetSecureAccessBastionApi(v string)`

SetSecureAccessBastionApi sets SecureAccessBastionApi field to given value.

### HasSecureAccessBastionApi

`func (o *UpdateItem) HasSecureAccessBastionApi() bool`

HasSecureAccessBastionApi returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *UpdateItem) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *UpdateItem) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *UpdateItem) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *UpdateItem) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessBastionSsh

`func (o *UpdateItem) GetSecureAccessBastionSsh() string`

GetSecureAccessBastionSsh returns the SecureAccessBastionSsh field if non-nil, zero value otherwise.

### GetSecureAccessBastionSshOk

`func (o *UpdateItem) GetSecureAccessBastionSshOk() (*string, bool)`

GetSecureAccessBastionSshOk returns a tuple with the SecureAccessBastionSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionSsh

`func (o *UpdateItem) SetSecureAccessBastionSsh(v string)`

SetSecureAccessBastionSsh sets SecureAccessBastionSsh field to given value.

### HasSecureAccessBastionSsh

`func (o *UpdateItem) HasSecureAccessBastionSsh() bool`

HasSecureAccessBastionSsh returns a boolean if a field has been set.

### GetSecureAccessClusterEndpoint

`func (o *UpdateItem) GetSecureAccessClusterEndpoint() string`

GetSecureAccessClusterEndpoint returns the SecureAccessClusterEndpoint field if non-nil, zero value otherwise.

### GetSecureAccessClusterEndpointOk

`func (o *UpdateItem) GetSecureAccessClusterEndpointOk() (*string, bool)`

GetSecureAccessClusterEndpointOk returns a tuple with the SecureAccessClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessClusterEndpoint

`func (o *UpdateItem) SetSecureAccessClusterEndpoint(v string)`

SetSecureAccessClusterEndpoint sets SecureAccessClusterEndpoint field to given value.

### HasSecureAccessClusterEndpoint

`func (o *UpdateItem) HasSecureAccessClusterEndpoint() bool`

HasSecureAccessClusterEndpoint returns a boolean if a field has been set.

### GetSecureAccessDashboardUrl

`func (o *UpdateItem) GetSecureAccessDashboardUrl() string`

GetSecureAccessDashboardUrl returns the SecureAccessDashboardUrl field if non-nil, zero value otherwise.

### GetSecureAccessDashboardUrlOk

`func (o *UpdateItem) GetSecureAccessDashboardUrlOk() (*string, bool)`

GetSecureAccessDashboardUrlOk returns a tuple with the SecureAccessDashboardUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDashboardUrl

`func (o *UpdateItem) SetSecureAccessDashboardUrl(v string)`

SetSecureAccessDashboardUrl sets SecureAccessDashboardUrl field to given value.

### HasSecureAccessDashboardUrl

`func (o *UpdateItem) HasSecureAccessDashboardUrl() bool`

HasSecureAccessDashboardUrl returns a boolean if a field has been set.

### GetSecureAccessDbName

`func (o *UpdateItem) GetSecureAccessDbName() string`

GetSecureAccessDbName returns the SecureAccessDbName field if non-nil, zero value otherwise.

### GetSecureAccessDbNameOk

`func (o *UpdateItem) GetSecureAccessDbNameOk() (*string, bool)`

GetSecureAccessDbNameOk returns a tuple with the SecureAccessDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbName

`func (o *UpdateItem) SetSecureAccessDbName(v string)`

SetSecureAccessDbName sets SecureAccessDbName field to given value.

### HasSecureAccessDbName

`func (o *UpdateItem) HasSecureAccessDbName() bool`

HasSecureAccessDbName returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *UpdateItem) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *UpdateItem) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *UpdateItem) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *UpdateItem) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *UpdateItem) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *UpdateItem) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *UpdateItem) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *UpdateItem) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *UpdateItem) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *UpdateItem) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *UpdateItem) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *UpdateItem) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *UpdateItem) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *UpdateItem) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *UpdateItem) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *UpdateItem) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *UpdateItem) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *UpdateItem) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *UpdateItem) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *UpdateItem) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetSecureAccessRmHost

`func (o *UpdateItem) GetSecureAccessRmHost() []string`

GetSecureAccessRmHost returns the SecureAccessRmHost field if non-nil, zero value otherwise.

### GetSecureAccessRmHostOk

`func (o *UpdateItem) GetSecureAccessRmHostOk() (*[]string, bool)`

GetSecureAccessRmHostOk returns a tuple with the SecureAccessRmHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRmHost

`func (o *UpdateItem) SetSecureAccessRmHost(v []string)`

SetSecureAccessRmHost sets SecureAccessRmHost field to given value.

### HasSecureAccessRmHost

`func (o *UpdateItem) HasSecureAccessRmHost() bool`

HasSecureAccessRmHost returns a boolean if a field has been set.

### GetSecureAccessSshCreds

`func (o *UpdateItem) GetSecureAccessSshCreds() string`

GetSecureAccessSshCreds returns the SecureAccessSshCreds field if non-nil, zero value otherwise.

### GetSecureAccessSshCredsOk

`func (o *UpdateItem) GetSecureAccessSshCredsOk() (*string, bool)`

GetSecureAccessSshCredsOk returns a tuple with the SecureAccessSshCreds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSshCreds

`func (o *UpdateItem) SetSecureAccessSshCreds(v string)`

SetSecureAccessSshCreds sets SecureAccessSshCreds field to given value.

### HasSecureAccessSshCreds

`func (o *UpdateItem) HasSecureAccessSshCreds() bool`

HasSecureAccessSshCreds returns a boolean if a field has been set.

### GetSecureAccessSshCredsUser

`func (o *UpdateItem) GetSecureAccessSshCredsUser() string`

GetSecureAccessSshCredsUser returns the SecureAccessSshCredsUser field if non-nil, zero value otherwise.

### GetSecureAccessSshCredsUserOk

`func (o *UpdateItem) GetSecureAccessSshCredsUserOk() (*string, bool)`

GetSecureAccessSshCredsUserOk returns a tuple with the SecureAccessSshCredsUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessSshCredsUser

`func (o *UpdateItem) SetSecureAccessSshCredsUser(v string)`

SetSecureAccessSshCredsUser sets SecureAccessSshCredsUser field to given value.

### HasSecureAccessSshCredsUser

`func (o *UpdateItem) HasSecureAccessSshCredsUser() bool`

HasSecureAccessSshCredsUser returns a boolean if a field has been set.

### GetSecureAccessUrl

`func (o *UpdateItem) GetSecureAccessUrl() string`

GetSecureAccessUrl returns the SecureAccessUrl field if non-nil, zero value otherwise.

### GetSecureAccessUrlOk

`func (o *UpdateItem) GetSecureAccessUrlOk() (*string, bool)`

GetSecureAccessUrlOk returns a tuple with the SecureAccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUrl

`func (o *UpdateItem) SetSecureAccessUrl(v string)`

SetSecureAccessUrl sets SecureAccessUrl field to given value.

### HasSecureAccessUrl

`func (o *UpdateItem) HasSecureAccessUrl() bool`

HasSecureAccessUrl returns a boolean if a field has been set.

### GetSecureAccessUseInternalBastion

`func (o *UpdateItem) GetSecureAccessUseInternalBastion() bool`

GetSecureAccessUseInternalBastion returns the SecureAccessUseInternalBastion field if non-nil, zero value otherwise.

### GetSecureAccessUseInternalBastionOk

`func (o *UpdateItem) GetSecureAccessUseInternalBastionOk() (*bool, bool)`

GetSecureAccessUseInternalBastionOk returns a tuple with the SecureAccessUseInternalBastion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessUseInternalBastion

`func (o *UpdateItem) SetSecureAccessUseInternalBastion(v bool)`

SetSecureAccessUseInternalBastion sets SecureAccessUseInternalBastion field to given value.

### HasSecureAccessUseInternalBastion

`func (o *UpdateItem) HasSecureAccessUseInternalBastion() bool`

HasSecureAccessUseInternalBastion returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *UpdateItem) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *UpdateItem) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *UpdateItem) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *UpdateItem) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *UpdateItem) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *UpdateItem) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *UpdateItem) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *UpdateItem) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetToken

`func (o *UpdateItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateItem) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateItem) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateItem) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateItem) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateItem) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateItem) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



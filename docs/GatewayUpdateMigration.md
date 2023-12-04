# GatewayUpdateMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var1passwordEmail** | Pointer to **string** | 1Password user email to connect to the API | [optional] 
**Var1passwordPassword** | Pointer to **string** | 1Password user password to connect to the API | [optional] 
**Var1passwordSecretKey** | Pointer to **string** | 1Password user secret key to connect to the API | [optional] 
**Var1passwordUrl** | Pointer to **string** | 1Password api container url | [optional] 
**Var1passwordVaults** | Pointer to **[]string** | 1Password list of vault to get the items from | [optional] 
**AdDiscoverServices** | Pointer to **string** | Enable/Disable discovery of Windows services from each domain server as part of the SSH/Windows Rotated Secrets. Default is false. (Relevant only for Active Directory migration) | [optional] [default to "false"]
**AdSshPort** | Pointer to **string** | Set the SSH Port for further connection to the domain servers. Default is port 22 (Relevant only for Active Directory migration) | [optional] [default to "22"]
**AdTargetsType** | Pointer to **string** | Set the target type of the domain servers [ssh/windows](Relevant only for Active Directory migration) | [optional] [default to "windows"]
**AdWinrmOverHttp** | Pointer to **string** | Use WinRM over HTTP, by default runs over HTTPS | [optional] [default to "false"]
**AdWinrmPort** | Pointer to **string** | Set the WinRM Port for further connection to the domain servers. Default is 5986 (Relevant only for Active Directory migration) | [optional] [default to "5986"]
**AdAutoRotate** | Pointer to **string** | Enable/Disable automatic/recurrent rotation for migrated secrets. Default is false: only manual rotation is allowed for migrated secrets. If set to true, this command should be combined with --ad-rotation-interval and --ad-rotation-hour parameters (Relevant only for Active Directory migration) | [optional] 
**AdComputerBaseDn** | Pointer to **string** | Distinguished Name of Computer objects (servers) to search in Active Directory e.g.: CN&#x3D;Computers,DC&#x3D;example,DC&#x3D;com (Relevant only for Active Directory migration) | [optional] 
**AdDiscoverLocalUsers** | Pointer to **string** | Enable/Disable discovery of local users from each domain server and migrate them as SSH/Windows Rotated Secrets. Default is false: only domain users will be migrated. Discovery of local users might require further installation of SSH on the servers, based on the supplied computer base DN. This will be implemented automatically as part of the migration process (Relevant only for Active Directory migration) | [optional] 
**AdDomainName** | Pointer to **string** | Active Directory Domain Name (Relevant only for Active Directory migration) | [optional] 
**AdDomainUsersPathTemplate** | Pointer to **string** | Path location template for migrating domain users as Rotated Secrets e.g.: .../DomainUsers/{{USERNAME}} (Relevant only for Active Directory migration) | [optional] 
**AdLocalUsersIgnore** | Pointer to **string** | Comma-separated list of Local Users which should not be migrated (Relevant only for Active Directory migration) | [optional] 
**AdLocalUsersPathTemplate** | Pointer to **string** | Path location template for migrating domain users as Rotated Secrets e.g.: .../LocalUsers/{{COMPUTER_NAME}}/{{USERNAME}} (Relevant only for Active Directory migration) | [optional] 
**AdRotationHour** | Pointer to **int32** | The hour of the scheduled rotation in UTC (Relevant only for Active Directory migration) | [optional] 
**AdRotationInterval** | Pointer to **int32** | The number of days to wait between every automatic rotation [1-365] (Relevant only for Active Directory migration) | [optional] 
**AdSraEnableRdp** | Pointer to **string** | Enable/Disable RDP Secure Remote Access for the migrated local users rotated secrets. Default is false: rotated secrets will not be created with SRA (Relevant only for Active Directory migration) | [optional] 
**AdTargetName** | Pointer to **string** | Active Directory LDAP Target Name. Server type should be Active Directory (Relevant only for Active Directory migration) | [optional] 
**AdTargetsPathTemplate** | Pointer to **string** | Path location template for migrating domain servers as SSH/Windows Targets e.g.: .../Servers/{{COMPUTER_NAME}} (Relevant only for Active Directory migration) | [optional] 
**AdUserBaseDn** | Pointer to **string** | Distinguished Name of User objects to search in Active Directory, e.g.: CN&#x3D;Users,DC&#x3D;example,DC&#x3D;com (Relevant only for Active Directory migration) | [optional] 
**AdUserGroups** | Pointer to **string** | Comma-separated list of domain groups from which privileged domain users will be migrated. If empty, migrate all users based on the --ad-user-base-dn (Relevant only for Active Directory migration) | [optional] 
**AwsKey** | Pointer to **string** | AWS Secret Access Key (relevant only for AWS migration) | [optional] 
**AwsKeyId** | Pointer to **string** | AWS Access Key ID with sufficient permissions to get all secrets, e.g. &#39;arn:aws:secretsmanager:[Region]:[AccountId]:secret:[/path/to/secrets/_*]&#39; (relevant only for AWS migration) | [optional] 
**AwsRegion** | Pointer to **string** | AWS region of the required Secrets Manager (relevant only for AWS migration) | [optional] [default to "us-east-2"]
**AzureClientId** | Pointer to **string** | Azure Key Vault Access client ID, should be Azure AD App with a service principal (relevant only for Azure Key Vault migration) | [optional] 
**AzureKvName** | Pointer to **string** | Azure Key Vault Name (relevant only for Azure Key Vault migration) | [optional] 
**AzureSecret** | Pointer to **string** | Azure Key Vault secret (relevant only for Azure Key Vault migration) | [optional] 
**AzureTenantId** | Pointer to **string** | Azure Key Vault Access tenant ID (relevant only for Azure Key Vault migration) | [optional] 
**GcpKey** | Pointer to **string** | Base64-encoded GCP Service Account private key text with sufficient permissions to Secrets Manager, Minimum required permission is Secret Manager Secret Accessor, e.g. &#39;roles/secretmanager.secretAccessor&#39; (relevant only for GCP migration) | [optional] 
**HashiJson** | Pointer to **string** | Import secret key as json value or independent secrets (relevant only for HasiCorp Vault migration) [true/false] | [optional] [default to "true"]
**HashiNs** | Pointer to **[]string** | HashiCorp Vault Namespaces is a comma-separated list of namespaces which need to be imported into Akeyless Vault. For every provided namespace, all its child namespaces are imported as well, e.g. nmsp/subnmsp1/subnmsp2,nmsp/anothernmsp. By default, import all namespaces (relevant only for HasiCorp Vault migration) | [optional] 
**HashiToken** | Pointer to **string** | HashiCorp Vault access token with sufficient permissions to preform list &amp; read operations on secrets objects (relevant only for HasiCorp Vault migration) | [optional] 
**HashiUrl** | Pointer to **string** | HashiCorp Vault API URL, e.g. https://vault-mgr01:8200 (relevant only for HasiCorp Vault migration) | [optional] 
**Id** | Pointer to **string** | Migration ID (Can be retrieved with gateway-list-migration command) | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**K8sCaCertificate** | Pointer to **[]int32** | For Certificate Authentication method K8s Cluster CA certificate (relevant only for K8s migration with Certificate Authentication method) | [optional] 
**K8sClientCertificate** | Pointer to **[]int32** | K8s Client certificate with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Certificate Authentication method) | [optional] 
**K8sClientKey** | Pointer to **[]int32** | K8s Client key (relevant only for K8s migration with Certificate Authentication method) | [optional] 
**K8sNamespace** | Pointer to **string** | K8s Namespace, Use this field to import secrets from a particular namespace only. By default, the secrets are imported from all namespaces (relevant only for K8s migration) | [optional] 
**K8sPassword** | Pointer to **string** | K8s Client password (relevant only for K8s migration with Password Authentication method) | [optional] 
**K8sSkipSystem** | Pointer to **bool** | K8s Skip Control Plane Secrets, This option allows to avoid importing secrets from system namespaces (relevant only for K8s migration) | [optional] 
**K8sToken** | Pointer to **string** | For Token Authentication method K8s Bearer Token with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Token Authentication method) | [optional] 
**K8sUrl** | Pointer to **string** | K8s API Server URL, e.g. https://k8s-api.mycompany.com:6443 (relevant only for K8s migration) | [optional] 
**K8sUsername** | Pointer to **string** | For Password Authentication method K8s Client username with sufficient permission to list and get secrets in the namespace(s) you selected (relevant only for K8s migration with Password Authentication method) | [optional] 
**Name** | Pointer to **string** | Migration name | [optional] 
**NewName** | Pointer to **string** | New migration name | [optional] 
**ProtectionKey** | Pointer to **string** | The name of the key that protects the classic key value (if empty, the account default key will be used) | [optional] 
**SiAutoRotate** | Pointer to **string** | Enable/Disable automatic/recurrent rotation for migrated secrets. Default is false: only manual rotation is allowed for migrated secrets. If set to true, this command should be combined with --si-rotation-interval and --si-rotation-hour parameters (Relevant only for Server Inventory migration) | [optional] 
**SiRotationHour** | Pointer to **int32** | The hour of the scheduled rotation in UTC (Relevant only for Server Inventory migration) | [optional] 
**SiRotationInterval** | Pointer to **int32** | The number of days to wait between every automatic rotation [1-365] (Relevant only for Server Inventory migration) | [optional] 
**SiSraEnableRdp** | Pointer to **string** | Enable/Disable RDP Secure Remote Access for the migrated local users rotated secrets. Default is false: rotated secrets will not be created with SRA (Relevant only for Server Inventory migration) | [optional] [default to "false"]
**SiTargetName** | **string** | SSH, Windows or Linked Target Name. (Relevant only for Server Inventory migration) | 
**SiUsersIgnore** | Pointer to **string** | Comma-separated list of Local Users which should not be migrated (Relevant only for Server Inventory migration) | [optional] 
**SiUsersPathTemplate** | **string** | Path location template for migrating users as Rotated Secrets e.g.: .../Users/{{COMPUTER_NAME}}/{{USERNAME}} (Relevant only for Server Inventory migration) | 
**TargetLocation** | **string** | Target location in Akeyless for imported secrets | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateMigration

`func NewGatewayUpdateMigration(siTargetName string, siUsersPathTemplate string, targetLocation string, ) *GatewayUpdateMigration`

NewGatewayUpdateMigration instantiates a new GatewayUpdateMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateMigrationWithDefaults

`func NewGatewayUpdateMigrationWithDefaults() *GatewayUpdateMigration`

NewGatewayUpdateMigrationWithDefaults instantiates a new GatewayUpdateMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar1passwordEmail

`func (o *GatewayUpdateMigration) GetVar1passwordEmail() string`

GetVar1passwordEmail returns the Var1passwordEmail field if non-nil, zero value otherwise.

### GetVar1passwordEmailOk

`func (o *GatewayUpdateMigration) GetVar1passwordEmailOk() (*string, bool)`

GetVar1passwordEmailOk returns a tuple with the Var1passwordEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordEmail

`func (o *GatewayUpdateMigration) SetVar1passwordEmail(v string)`

SetVar1passwordEmail sets Var1passwordEmail field to given value.

### HasVar1passwordEmail

`func (o *GatewayUpdateMigration) HasVar1passwordEmail() bool`

HasVar1passwordEmail returns a boolean if a field has been set.

### GetVar1passwordPassword

`func (o *GatewayUpdateMigration) GetVar1passwordPassword() string`

GetVar1passwordPassword returns the Var1passwordPassword field if non-nil, zero value otherwise.

### GetVar1passwordPasswordOk

`func (o *GatewayUpdateMigration) GetVar1passwordPasswordOk() (*string, bool)`

GetVar1passwordPasswordOk returns a tuple with the Var1passwordPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordPassword

`func (o *GatewayUpdateMigration) SetVar1passwordPassword(v string)`

SetVar1passwordPassword sets Var1passwordPassword field to given value.

### HasVar1passwordPassword

`func (o *GatewayUpdateMigration) HasVar1passwordPassword() bool`

HasVar1passwordPassword returns a boolean if a field has been set.

### GetVar1passwordSecretKey

`func (o *GatewayUpdateMigration) GetVar1passwordSecretKey() string`

GetVar1passwordSecretKey returns the Var1passwordSecretKey field if non-nil, zero value otherwise.

### GetVar1passwordSecretKeyOk

`func (o *GatewayUpdateMigration) GetVar1passwordSecretKeyOk() (*string, bool)`

GetVar1passwordSecretKeyOk returns a tuple with the Var1passwordSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordSecretKey

`func (o *GatewayUpdateMigration) SetVar1passwordSecretKey(v string)`

SetVar1passwordSecretKey sets Var1passwordSecretKey field to given value.

### HasVar1passwordSecretKey

`func (o *GatewayUpdateMigration) HasVar1passwordSecretKey() bool`

HasVar1passwordSecretKey returns a boolean if a field has been set.

### GetVar1passwordUrl

`func (o *GatewayUpdateMigration) GetVar1passwordUrl() string`

GetVar1passwordUrl returns the Var1passwordUrl field if non-nil, zero value otherwise.

### GetVar1passwordUrlOk

`func (o *GatewayUpdateMigration) GetVar1passwordUrlOk() (*string, bool)`

GetVar1passwordUrlOk returns a tuple with the Var1passwordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordUrl

`func (o *GatewayUpdateMigration) SetVar1passwordUrl(v string)`

SetVar1passwordUrl sets Var1passwordUrl field to given value.

### HasVar1passwordUrl

`func (o *GatewayUpdateMigration) HasVar1passwordUrl() bool`

HasVar1passwordUrl returns a boolean if a field has been set.

### GetVar1passwordVaults

`func (o *GatewayUpdateMigration) GetVar1passwordVaults() []string`

GetVar1passwordVaults returns the Var1passwordVaults field if non-nil, zero value otherwise.

### GetVar1passwordVaultsOk

`func (o *GatewayUpdateMigration) GetVar1passwordVaultsOk() (*[]string, bool)`

GetVar1passwordVaultsOk returns a tuple with the Var1passwordVaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordVaults

`func (o *GatewayUpdateMigration) SetVar1passwordVaults(v []string)`

SetVar1passwordVaults sets Var1passwordVaults field to given value.

### HasVar1passwordVaults

`func (o *GatewayUpdateMigration) HasVar1passwordVaults() bool`

HasVar1passwordVaults returns a boolean if a field has been set.

### GetAdDiscoverServices

`func (o *GatewayUpdateMigration) GetAdDiscoverServices() string`

GetAdDiscoverServices returns the AdDiscoverServices field if non-nil, zero value otherwise.

### GetAdDiscoverServicesOk

`func (o *GatewayUpdateMigration) GetAdDiscoverServicesOk() (*string, bool)`

GetAdDiscoverServicesOk returns a tuple with the AdDiscoverServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDiscoverServices

`func (o *GatewayUpdateMigration) SetAdDiscoverServices(v string)`

SetAdDiscoverServices sets AdDiscoverServices field to given value.

### HasAdDiscoverServices

`func (o *GatewayUpdateMigration) HasAdDiscoverServices() bool`

HasAdDiscoverServices returns a boolean if a field has been set.

### GetAdSshPort

`func (o *GatewayUpdateMigration) GetAdSshPort() string`

GetAdSshPort returns the AdSshPort field if non-nil, zero value otherwise.

### GetAdSshPortOk

`func (o *GatewayUpdateMigration) GetAdSshPortOk() (*string, bool)`

GetAdSshPortOk returns a tuple with the AdSshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSshPort

`func (o *GatewayUpdateMigration) SetAdSshPort(v string)`

SetAdSshPort sets AdSshPort field to given value.

### HasAdSshPort

`func (o *GatewayUpdateMigration) HasAdSshPort() bool`

HasAdSshPort returns a boolean if a field has been set.

### GetAdTargetsType

`func (o *GatewayUpdateMigration) GetAdTargetsType() string`

GetAdTargetsType returns the AdTargetsType field if non-nil, zero value otherwise.

### GetAdTargetsTypeOk

`func (o *GatewayUpdateMigration) GetAdTargetsTypeOk() (*string, bool)`

GetAdTargetsTypeOk returns a tuple with the AdTargetsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdTargetsType

`func (o *GatewayUpdateMigration) SetAdTargetsType(v string)`

SetAdTargetsType sets AdTargetsType field to given value.

### HasAdTargetsType

`func (o *GatewayUpdateMigration) HasAdTargetsType() bool`

HasAdTargetsType returns a boolean if a field has been set.

### GetAdWinrmOverHttp

`func (o *GatewayUpdateMigration) GetAdWinrmOverHttp() string`

GetAdWinrmOverHttp returns the AdWinrmOverHttp field if non-nil, zero value otherwise.

### GetAdWinrmOverHttpOk

`func (o *GatewayUpdateMigration) GetAdWinrmOverHttpOk() (*string, bool)`

GetAdWinrmOverHttpOk returns a tuple with the AdWinrmOverHttp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdWinrmOverHttp

`func (o *GatewayUpdateMigration) SetAdWinrmOverHttp(v string)`

SetAdWinrmOverHttp sets AdWinrmOverHttp field to given value.

### HasAdWinrmOverHttp

`func (o *GatewayUpdateMigration) HasAdWinrmOverHttp() bool`

HasAdWinrmOverHttp returns a boolean if a field has been set.

### GetAdWinrmPort

`func (o *GatewayUpdateMigration) GetAdWinrmPort() string`

GetAdWinrmPort returns the AdWinrmPort field if non-nil, zero value otherwise.

### GetAdWinrmPortOk

`func (o *GatewayUpdateMigration) GetAdWinrmPortOk() (*string, bool)`

GetAdWinrmPortOk returns a tuple with the AdWinrmPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdWinrmPort

`func (o *GatewayUpdateMigration) SetAdWinrmPort(v string)`

SetAdWinrmPort sets AdWinrmPort field to given value.

### HasAdWinrmPort

`func (o *GatewayUpdateMigration) HasAdWinrmPort() bool`

HasAdWinrmPort returns a boolean if a field has been set.

### GetAdAutoRotate

`func (o *GatewayUpdateMigration) GetAdAutoRotate() string`

GetAdAutoRotate returns the AdAutoRotate field if non-nil, zero value otherwise.

### GetAdAutoRotateOk

`func (o *GatewayUpdateMigration) GetAdAutoRotateOk() (*string, bool)`

GetAdAutoRotateOk returns a tuple with the AdAutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAutoRotate

`func (o *GatewayUpdateMigration) SetAdAutoRotate(v string)`

SetAdAutoRotate sets AdAutoRotate field to given value.

### HasAdAutoRotate

`func (o *GatewayUpdateMigration) HasAdAutoRotate() bool`

HasAdAutoRotate returns a boolean if a field has been set.

### GetAdComputerBaseDn

`func (o *GatewayUpdateMigration) GetAdComputerBaseDn() string`

GetAdComputerBaseDn returns the AdComputerBaseDn field if non-nil, zero value otherwise.

### GetAdComputerBaseDnOk

`func (o *GatewayUpdateMigration) GetAdComputerBaseDnOk() (*string, bool)`

GetAdComputerBaseDnOk returns a tuple with the AdComputerBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdComputerBaseDn

`func (o *GatewayUpdateMigration) SetAdComputerBaseDn(v string)`

SetAdComputerBaseDn sets AdComputerBaseDn field to given value.

### HasAdComputerBaseDn

`func (o *GatewayUpdateMigration) HasAdComputerBaseDn() bool`

HasAdComputerBaseDn returns a boolean if a field has been set.

### GetAdDiscoverLocalUsers

`func (o *GatewayUpdateMigration) GetAdDiscoverLocalUsers() string`

GetAdDiscoverLocalUsers returns the AdDiscoverLocalUsers field if non-nil, zero value otherwise.

### GetAdDiscoverLocalUsersOk

`func (o *GatewayUpdateMigration) GetAdDiscoverLocalUsersOk() (*string, bool)`

GetAdDiscoverLocalUsersOk returns a tuple with the AdDiscoverLocalUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDiscoverLocalUsers

`func (o *GatewayUpdateMigration) SetAdDiscoverLocalUsers(v string)`

SetAdDiscoverLocalUsers sets AdDiscoverLocalUsers field to given value.

### HasAdDiscoverLocalUsers

`func (o *GatewayUpdateMigration) HasAdDiscoverLocalUsers() bool`

HasAdDiscoverLocalUsers returns a boolean if a field has been set.

### GetAdDomainName

`func (o *GatewayUpdateMigration) GetAdDomainName() string`

GetAdDomainName returns the AdDomainName field if non-nil, zero value otherwise.

### GetAdDomainNameOk

`func (o *GatewayUpdateMigration) GetAdDomainNameOk() (*string, bool)`

GetAdDomainNameOk returns a tuple with the AdDomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainName

`func (o *GatewayUpdateMigration) SetAdDomainName(v string)`

SetAdDomainName sets AdDomainName field to given value.

### HasAdDomainName

`func (o *GatewayUpdateMigration) HasAdDomainName() bool`

HasAdDomainName returns a boolean if a field has been set.

### GetAdDomainUsersPathTemplate

`func (o *GatewayUpdateMigration) GetAdDomainUsersPathTemplate() string`

GetAdDomainUsersPathTemplate returns the AdDomainUsersPathTemplate field if non-nil, zero value otherwise.

### GetAdDomainUsersPathTemplateOk

`func (o *GatewayUpdateMigration) GetAdDomainUsersPathTemplateOk() (*string, bool)`

GetAdDomainUsersPathTemplateOk returns a tuple with the AdDomainUsersPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainUsersPathTemplate

`func (o *GatewayUpdateMigration) SetAdDomainUsersPathTemplate(v string)`

SetAdDomainUsersPathTemplate sets AdDomainUsersPathTemplate field to given value.

### HasAdDomainUsersPathTemplate

`func (o *GatewayUpdateMigration) HasAdDomainUsersPathTemplate() bool`

HasAdDomainUsersPathTemplate returns a boolean if a field has been set.

### GetAdLocalUsersIgnore

`func (o *GatewayUpdateMigration) GetAdLocalUsersIgnore() string`

GetAdLocalUsersIgnore returns the AdLocalUsersIgnore field if non-nil, zero value otherwise.

### GetAdLocalUsersIgnoreOk

`func (o *GatewayUpdateMigration) GetAdLocalUsersIgnoreOk() (*string, bool)`

GetAdLocalUsersIgnoreOk returns a tuple with the AdLocalUsersIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdLocalUsersIgnore

`func (o *GatewayUpdateMigration) SetAdLocalUsersIgnore(v string)`

SetAdLocalUsersIgnore sets AdLocalUsersIgnore field to given value.

### HasAdLocalUsersIgnore

`func (o *GatewayUpdateMigration) HasAdLocalUsersIgnore() bool`

HasAdLocalUsersIgnore returns a boolean if a field has been set.

### GetAdLocalUsersPathTemplate

`func (o *GatewayUpdateMigration) GetAdLocalUsersPathTemplate() string`

GetAdLocalUsersPathTemplate returns the AdLocalUsersPathTemplate field if non-nil, zero value otherwise.

### GetAdLocalUsersPathTemplateOk

`func (o *GatewayUpdateMigration) GetAdLocalUsersPathTemplateOk() (*string, bool)`

GetAdLocalUsersPathTemplateOk returns a tuple with the AdLocalUsersPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdLocalUsersPathTemplate

`func (o *GatewayUpdateMigration) SetAdLocalUsersPathTemplate(v string)`

SetAdLocalUsersPathTemplate sets AdLocalUsersPathTemplate field to given value.

### HasAdLocalUsersPathTemplate

`func (o *GatewayUpdateMigration) HasAdLocalUsersPathTemplate() bool`

HasAdLocalUsersPathTemplate returns a boolean if a field has been set.

### GetAdRotationHour

`func (o *GatewayUpdateMigration) GetAdRotationHour() int32`

GetAdRotationHour returns the AdRotationHour field if non-nil, zero value otherwise.

### GetAdRotationHourOk

`func (o *GatewayUpdateMigration) GetAdRotationHourOk() (*int32, bool)`

GetAdRotationHourOk returns a tuple with the AdRotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRotationHour

`func (o *GatewayUpdateMigration) SetAdRotationHour(v int32)`

SetAdRotationHour sets AdRotationHour field to given value.

### HasAdRotationHour

`func (o *GatewayUpdateMigration) HasAdRotationHour() bool`

HasAdRotationHour returns a boolean if a field has been set.

### GetAdRotationInterval

`func (o *GatewayUpdateMigration) GetAdRotationInterval() int32`

GetAdRotationInterval returns the AdRotationInterval field if non-nil, zero value otherwise.

### GetAdRotationIntervalOk

`func (o *GatewayUpdateMigration) GetAdRotationIntervalOk() (*int32, bool)`

GetAdRotationIntervalOk returns a tuple with the AdRotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdRotationInterval

`func (o *GatewayUpdateMigration) SetAdRotationInterval(v int32)`

SetAdRotationInterval sets AdRotationInterval field to given value.

### HasAdRotationInterval

`func (o *GatewayUpdateMigration) HasAdRotationInterval() bool`

HasAdRotationInterval returns a boolean if a field has been set.

### GetAdSraEnableRdp

`func (o *GatewayUpdateMigration) GetAdSraEnableRdp() string`

GetAdSraEnableRdp returns the AdSraEnableRdp field if non-nil, zero value otherwise.

### GetAdSraEnableRdpOk

`func (o *GatewayUpdateMigration) GetAdSraEnableRdpOk() (*string, bool)`

GetAdSraEnableRdpOk returns a tuple with the AdSraEnableRdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdSraEnableRdp

`func (o *GatewayUpdateMigration) SetAdSraEnableRdp(v string)`

SetAdSraEnableRdp sets AdSraEnableRdp field to given value.

### HasAdSraEnableRdp

`func (o *GatewayUpdateMigration) HasAdSraEnableRdp() bool`

HasAdSraEnableRdp returns a boolean if a field has been set.

### GetAdTargetName

`func (o *GatewayUpdateMigration) GetAdTargetName() string`

GetAdTargetName returns the AdTargetName field if non-nil, zero value otherwise.

### GetAdTargetNameOk

`func (o *GatewayUpdateMigration) GetAdTargetNameOk() (*string, bool)`

GetAdTargetNameOk returns a tuple with the AdTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdTargetName

`func (o *GatewayUpdateMigration) SetAdTargetName(v string)`

SetAdTargetName sets AdTargetName field to given value.

### HasAdTargetName

`func (o *GatewayUpdateMigration) HasAdTargetName() bool`

HasAdTargetName returns a boolean if a field has been set.

### GetAdTargetsPathTemplate

`func (o *GatewayUpdateMigration) GetAdTargetsPathTemplate() string`

GetAdTargetsPathTemplate returns the AdTargetsPathTemplate field if non-nil, zero value otherwise.

### GetAdTargetsPathTemplateOk

`func (o *GatewayUpdateMigration) GetAdTargetsPathTemplateOk() (*string, bool)`

GetAdTargetsPathTemplateOk returns a tuple with the AdTargetsPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdTargetsPathTemplate

`func (o *GatewayUpdateMigration) SetAdTargetsPathTemplate(v string)`

SetAdTargetsPathTemplate sets AdTargetsPathTemplate field to given value.

### HasAdTargetsPathTemplate

`func (o *GatewayUpdateMigration) HasAdTargetsPathTemplate() bool`

HasAdTargetsPathTemplate returns a boolean if a field has been set.

### GetAdUserBaseDn

`func (o *GatewayUpdateMigration) GetAdUserBaseDn() string`

GetAdUserBaseDn returns the AdUserBaseDn field if non-nil, zero value otherwise.

### GetAdUserBaseDnOk

`func (o *GatewayUpdateMigration) GetAdUserBaseDnOk() (*string, bool)`

GetAdUserBaseDnOk returns a tuple with the AdUserBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserBaseDn

`func (o *GatewayUpdateMigration) SetAdUserBaseDn(v string)`

SetAdUserBaseDn sets AdUserBaseDn field to given value.

### HasAdUserBaseDn

`func (o *GatewayUpdateMigration) HasAdUserBaseDn() bool`

HasAdUserBaseDn returns a boolean if a field has been set.

### GetAdUserGroups

`func (o *GatewayUpdateMigration) GetAdUserGroups() string`

GetAdUserGroups returns the AdUserGroups field if non-nil, zero value otherwise.

### GetAdUserGroupsOk

`func (o *GatewayUpdateMigration) GetAdUserGroupsOk() (*string, bool)`

GetAdUserGroupsOk returns a tuple with the AdUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserGroups

`func (o *GatewayUpdateMigration) SetAdUserGroups(v string)`

SetAdUserGroups sets AdUserGroups field to given value.

### HasAdUserGroups

`func (o *GatewayUpdateMigration) HasAdUserGroups() bool`

HasAdUserGroups returns a boolean if a field has been set.

### GetAwsKey

`func (o *GatewayUpdateMigration) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *GatewayUpdateMigration) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *GatewayUpdateMigration) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.

### HasAwsKey

`func (o *GatewayUpdateMigration) HasAwsKey() bool`

HasAwsKey returns a boolean if a field has been set.

### GetAwsKeyId

`func (o *GatewayUpdateMigration) GetAwsKeyId() string`

GetAwsKeyId returns the AwsKeyId field if non-nil, zero value otherwise.

### GetAwsKeyIdOk

`func (o *GatewayUpdateMigration) GetAwsKeyIdOk() (*string, bool)`

GetAwsKeyIdOk returns a tuple with the AwsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKeyId

`func (o *GatewayUpdateMigration) SetAwsKeyId(v string)`

SetAwsKeyId sets AwsKeyId field to given value.

### HasAwsKeyId

`func (o *GatewayUpdateMigration) HasAwsKeyId() bool`

HasAwsKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *GatewayUpdateMigration) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *GatewayUpdateMigration) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *GatewayUpdateMigration) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *GatewayUpdateMigration) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureClientId

`func (o *GatewayUpdateMigration) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *GatewayUpdateMigration) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *GatewayUpdateMigration) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *GatewayUpdateMigration) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureKvName

`func (o *GatewayUpdateMigration) GetAzureKvName() string`

GetAzureKvName returns the AzureKvName field if non-nil, zero value otherwise.

### GetAzureKvNameOk

`func (o *GatewayUpdateMigration) GetAzureKvNameOk() (*string, bool)`

GetAzureKvNameOk returns a tuple with the AzureKvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvName

`func (o *GatewayUpdateMigration) SetAzureKvName(v string)`

SetAzureKvName sets AzureKvName field to given value.

### HasAzureKvName

`func (o *GatewayUpdateMigration) HasAzureKvName() bool`

HasAzureKvName returns a boolean if a field has been set.

### GetAzureSecret

`func (o *GatewayUpdateMigration) GetAzureSecret() string`

GetAzureSecret returns the AzureSecret field if non-nil, zero value otherwise.

### GetAzureSecretOk

`func (o *GatewayUpdateMigration) GetAzureSecretOk() (*string, bool)`

GetAzureSecretOk returns a tuple with the AzureSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSecret

`func (o *GatewayUpdateMigration) SetAzureSecret(v string)`

SetAzureSecret sets AzureSecret field to given value.

### HasAzureSecret

`func (o *GatewayUpdateMigration) HasAzureSecret() bool`

HasAzureSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *GatewayUpdateMigration) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *GatewayUpdateMigration) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *GatewayUpdateMigration) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *GatewayUpdateMigration) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetGcpKey

`func (o *GatewayUpdateMigration) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *GatewayUpdateMigration) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *GatewayUpdateMigration) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *GatewayUpdateMigration) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetHashiJson

`func (o *GatewayUpdateMigration) GetHashiJson() string`

GetHashiJson returns the HashiJson field if non-nil, zero value otherwise.

### GetHashiJsonOk

`func (o *GatewayUpdateMigration) GetHashiJsonOk() (*string, bool)`

GetHashiJsonOk returns a tuple with the HashiJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiJson

`func (o *GatewayUpdateMigration) SetHashiJson(v string)`

SetHashiJson sets HashiJson field to given value.

### HasHashiJson

`func (o *GatewayUpdateMigration) HasHashiJson() bool`

HasHashiJson returns a boolean if a field has been set.

### GetHashiNs

`func (o *GatewayUpdateMigration) GetHashiNs() []string`

GetHashiNs returns the HashiNs field if non-nil, zero value otherwise.

### GetHashiNsOk

`func (o *GatewayUpdateMigration) GetHashiNsOk() (*[]string, bool)`

GetHashiNsOk returns a tuple with the HashiNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiNs

`func (o *GatewayUpdateMigration) SetHashiNs(v []string)`

SetHashiNs sets HashiNs field to given value.

### HasHashiNs

`func (o *GatewayUpdateMigration) HasHashiNs() bool`

HasHashiNs returns a boolean if a field has been set.

### GetHashiToken

`func (o *GatewayUpdateMigration) GetHashiToken() string`

GetHashiToken returns the HashiToken field if non-nil, zero value otherwise.

### GetHashiTokenOk

`func (o *GatewayUpdateMigration) GetHashiTokenOk() (*string, bool)`

GetHashiTokenOk returns a tuple with the HashiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiToken

`func (o *GatewayUpdateMigration) SetHashiToken(v string)`

SetHashiToken sets HashiToken field to given value.

### HasHashiToken

`func (o *GatewayUpdateMigration) HasHashiToken() bool`

HasHashiToken returns a boolean if a field has been set.

### GetHashiUrl

`func (o *GatewayUpdateMigration) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *GatewayUpdateMigration) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *GatewayUpdateMigration) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *GatewayUpdateMigration) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetId

`func (o *GatewayUpdateMigration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayUpdateMigration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayUpdateMigration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayUpdateMigration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateMigration) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateMigration) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateMigration) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateMigration) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetK8sCaCertificate

`func (o *GatewayUpdateMigration) GetK8sCaCertificate() []int32`

GetK8sCaCertificate returns the K8sCaCertificate field if non-nil, zero value otherwise.

### GetK8sCaCertificateOk

`func (o *GatewayUpdateMigration) GetK8sCaCertificateOk() (*[]int32, bool)`

GetK8sCaCertificateOk returns a tuple with the K8sCaCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sCaCertificate

`func (o *GatewayUpdateMigration) SetK8sCaCertificate(v []int32)`

SetK8sCaCertificate sets K8sCaCertificate field to given value.

### HasK8sCaCertificate

`func (o *GatewayUpdateMigration) HasK8sCaCertificate() bool`

HasK8sCaCertificate returns a boolean if a field has been set.

### GetK8sClientCertificate

`func (o *GatewayUpdateMigration) GetK8sClientCertificate() []int32`

GetK8sClientCertificate returns the K8sClientCertificate field if non-nil, zero value otherwise.

### GetK8sClientCertificateOk

`func (o *GatewayUpdateMigration) GetK8sClientCertificateOk() (*[]int32, bool)`

GetK8sClientCertificateOk returns a tuple with the K8sClientCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientCertificate

`func (o *GatewayUpdateMigration) SetK8sClientCertificate(v []int32)`

SetK8sClientCertificate sets K8sClientCertificate field to given value.

### HasK8sClientCertificate

`func (o *GatewayUpdateMigration) HasK8sClientCertificate() bool`

HasK8sClientCertificate returns a boolean if a field has been set.

### GetK8sClientKey

`func (o *GatewayUpdateMigration) GetK8sClientKey() []int32`

GetK8sClientKey returns the K8sClientKey field if non-nil, zero value otherwise.

### GetK8sClientKeyOk

`func (o *GatewayUpdateMigration) GetK8sClientKeyOk() (*[]int32, bool)`

GetK8sClientKeyOk returns a tuple with the K8sClientKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sClientKey

`func (o *GatewayUpdateMigration) SetK8sClientKey(v []int32)`

SetK8sClientKey sets K8sClientKey field to given value.

### HasK8sClientKey

`func (o *GatewayUpdateMigration) HasK8sClientKey() bool`

HasK8sClientKey returns a boolean if a field has been set.

### GetK8sNamespace

`func (o *GatewayUpdateMigration) GetK8sNamespace() string`

GetK8sNamespace returns the K8sNamespace field if non-nil, zero value otherwise.

### GetK8sNamespaceOk

`func (o *GatewayUpdateMigration) GetK8sNamespaceOk() (*string, bool)`

GetK8sNamespaceOk returns a tuple with the K8sNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sNamespace

`func (o *GatewayUpdateMigration) SetK8sNamespace(v string)`

SetK8sNamespace sets K8sNamespace field to given value.

### HasK8sNamespace

`func (o *GatewayUpdateMigration) HasK8sNamespace() bool`

HasK8sNamespace returns a boolean if a field has been set.

### GetK8sPassword

`func (o *GatewayUpdateMigration) GetK8sPassword() string`

GetK8sPassword returns the K8sPassword field if non-nil, zero value otherwise.

### GetK8sPasswordOk

`func (o *GatewayUpdateMigration) GetK8sPasswordOk() (*string, bool)`

GetK8sPasswordOk returns a tuple with the K8sPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sPassword

`func (o *GatewayUpdateMigration) SetK8sPassword(v string)`

SetK8sPassword sets K8sPassword field to given value.

### HasK8sPassword

`func (o *GatewayUpdateMigration) HasK8sPassword() bool`

HasK8sPassword returns a boolean if a field has been set.

### GetK8sSkipSystem

`func (o *GatewayUpdateMigration) GetK8sSkipSystem() bool`

GetK8sSkipSystem returns the K8sSkipSystem field if non-nil, zero value otherwise.

### GetK8sSkipSystemOk

`func (o *GatewayUpdateMigration) GetK8sSkipSystemOk() (*bool, bool)`

GetK8sSkipSystemOk returns a tuple with the K8sSkipSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sSkipSystem

`func (o *GatewayUpdateMigration) SetK8sSkipSystem(v bool)`

SetK8sSkipSystem sets K8sSkipSystem field to given value.

### HasK8sSkipSystem

`func (o *GatewayUpdateMigration) HasK8sSkipSystem() bool`

HasK8sSkipSystem returns a boolean if a field has been set.

### GetK8sToken

`func (o *GatewayUpdateMigration) GetK8sToken() string`

GetK8sToken returns the K8sToken field if non-nil, zero value otherwise.

### GetK8sTokenOk

`func (o *GatewayUpdateMigration) GetK8sTokenOk() (*string, bool)`

GetK8sTokenOk returns a tuple with the K8sToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sToken

`func (o *GatewayUpdateMigration) SetK8sToken(v string)`

SetK8sToken sets K8sToken field to given value.

### HasK8sToken

`func (o *GatewayUpdateMigration) HasK8sToken() bool`

HasK8sToken returns a boolean if a field has been set.

### GetK8sUrl

`func (o *GatewayUpdateMigration) GetK8sUrl() string`

GetK8sUrl returns the K8sUrl field if non-nil, zero value otherwise.

### GetK8sUrlOk

`func (o *GatewayUpdateMigration) GetK8sUrlOk() (*string, bool)`

GetK8sUrlOk returns a tuple with the K8sUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sUrl

`func (o *GatewayUpdateMigration) SetK8sUrl(v string)`

SetK8sUrl sets K8sUrl field to given value.

### HasK8sUrl

`func (o *GatewayUpdateMigration) HasK8sUrl() bool`

HasK8sUrl returns a boolean if a field has been set.

### GetK8sUsername

`func (o *GatewayUpdateMigration) GetK8sUsername() string`

GetK8sUsername returns the K8sUsername field if non-nil, zero value otherwise.

### GetK8sUsernameOk

`func (o *GatewayUpdateMigration) GetK8sUsernameOk() (*string, bool)`

GetK8sUsernameOk returns a tuple with the K8sUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sUsername

`func (o *GatewayUpdateMigration) SetK8sUsername(v string)`

SetK8sUsername sets K8sUsername field to given value.

### HasK8sUsername

`func (o *GatewayUpdateMigration) HasK8sUsername() bool`

HasK8sUsername returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateMigration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayUpdateMigration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNewName

`func (o *GatewayUpdateMigration) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateMigration) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateMigration) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateMigration) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProtectionKey

`func (o *GatewayUpdateMigration) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *GatewayUpdateMigration) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *GatewayUpdateMigration) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *GatewayUpdateMigration) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetSiAutoRotate

`func (o *GatewayUpdateMigration) GetSiAutoRotate() string`

GetSiAutoRotate returns the SiAutoRotate field if non-nil, zero value otherwise.

### GetSiAutoRotateOk

`func (o *GatewayUpdateMigration) GetSiAutoRotateOk() (*string, bool)`

GetSiAutoRotateOk returns a tuple with the SiAutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiAutoRotate

`func (o *GatewayUpdateMigration) SetSiAutoRotate(v string)`

SetSiAutoRotate sets SiAutoRotate field to given value.

### HasSiAutoRotate

`func (o *GatewayUpdateMigration) HasSiAutoRotate() bool`

HasSiAutoRotate returns a boolean if a field has been set.

### GetSiRotationHour

`func (o *GatewayUpdateMigration) GetSiRotationHour() int32`

GetSiRotationHour returns the SiRotationHour field if non-nil, zero value otherwise.

### GetSiRotationHourOk

`func (o *GatewayUpdateMigration) GetSiRotationHourOk() (*int32, bool)`

GetSiRotationHourOk returns a tuple with the SiRotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiRotationHour

`func (o *GatewayUpdateMigration) SetSiRotationHour(v int32)`

SetSiRotationHour sets SiRotationHour field to given value.

### HasSiRotationHour

`func (o *GatewayUpdateMigration) HasSiRotationHour() bool`

HasSiRotationHour returns a boolean if a field has been set.

### GetSiRotationInterval

`func (o *GatewayUpdateMigration) GetSiRotationInterval() int32`

GetSiRotationInterval returns the SiRotationInterval field if non-nil, zero value otherwise.

### GetSiRotationIntervalOk

`func (o *GatewayUpdateMigration) GetSiRotationIntervalOk() (*int32, bool)`

GetSiRotationIntervalOk returns a tuple with the SiRotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiRotationInterval

`func (o *GatewayUpdateMigration) SetSiRotationInterval(v int32)`

SetSiRotationInterval sets SiRotationInterval field to given value.

### HasSiRotationInterval

`func (o *GatewayUpdateMigration) HasSiRotationInterval() bool`

HasSiRotationInterval returns a boolean if a field has been set.

### GetSiSraEnableRdp

`func (o *GatewayUpdateMigration) GetSiSraEnableRdp() string`

GetSiSraEnableRdp returns the SiSraEnableRdp field if non-nil, zero value otherwise.

### GetSiSraEnableRdpOk

`func (o *GatewayUpdateMigration) GetSiSraEnableRdpOk() (*string, bool)`

GetSiSraEnableRdpOk returns a tuple with the SiSraEnableRdp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiSraEnableRdp

`func (o *GatewayUpdateMigration) SetSiSraEnableRdp(v string)`

SetSiSraEnableRdp sets SiSraEnableRdp field to given value.

### HasSiSraEnableRdp

`func (o *GatewayUpdateMigration) HasSiSraEnableRdp() bool`

HasSiSraEnableRdp returns a boolean if a field has been set.

### GetSiTargetName

`func (o *GatewayUpdateMigration) GetSiTargetName() string`

GetSiTargetName returns the SiTargetName field if non-nil, zero value otherwise.

### GetSiTargetNameOk

`func (o *GatewayUpdateMigration) GetSiTargetNameOk() (*string, bool)`

GetSiTargetNameOk returns a tuple with the SiTargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiTargetName

`func (o *GatewayUpdateMigration) SetSiTargetName(v string)`

SetSiTargetName sets SiTargetName field to given value.


### GetSiUsersIgnore

`func (o *GatewayUpdateMigration) GetSiUsersIgnore() string`

GetSiUsersIgnore returns the SiUsersIgnore field if non-nil, zero value otherwise.

### GetSiUsersIgnoreOk

`func (o *GatewayUpdateMigration) GetSiUsersIgnoreOk() (*string, bool)`

GetSiUsersIgnoreOk returns a tuple with the SiUsersIgnore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiUsersIgnore

`func (o *GatewayUpdateMigration) SetSiUsersIgnore(v string)`

SetSiUsersIgnore sets SiUsersIgnore field to given value.

### HasSiUsersIgnore

`func (o *GatewayUpdateMigration) HasSiUsersIgnore() bool`

HasSiUsersIgnore returns a boolean if a field has been set.

### GetSiUsersPathTemplate

`func (o *GatewayUpdateMigration) GetSiUsersPathTemplate() string`

GetSiUsersPathTemplate returns the SiUsersPathTemplate field if non-nil, zero value otherwise.

### GetSiUsersPathTemplateOk

`func (o *GatewayUpdateMigration) GetSiUsersPathTemplateOk() (*string, bool)`

GetSiUsersPathTemplateOk returns a tuple with the SiUsersPathTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiUsersPathTemplate

`func (o *GatewayUpdateMigration) SetSiUsersPathTemplate(v string)`

SetSiUsersPathTemplate sets SiUsersPathTemplate field to given value.


### GetTargetLocation

`func (o *GatewayUpdateMigration) GetTargetLocation() string`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *GatewayUpdateMigration) GetTargetLocationOk() (*string, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *GatewayUpdateMigration) SetTargetLocation(v string)`

SetTargetLocation sets TargetLocation field to given value.


### GetToken

`func (o *GatewayUpdateMigration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateMigration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateMigration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateMigration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateMigration) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateMigration) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateMigration) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateMigration) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



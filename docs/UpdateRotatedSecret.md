# UpdateRotatedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddTag** | Pointer to **[]string** | List of the new tags that will be attached to this item | [optional] 
**ApiId** | Pointer to **string** | API ID to rotate | [optional] 
**ApiKey** | Pointer to **string** | API key to rotate | [optional] 
**AutoRotate** | Pointer to **string** | Whether to automatically rotate every --rotation-interval days, or disable existing automatic rotation [true/false] | [optional] 
**AwsRegion** | Pointer to **string** | Region (used in aws) | [optional] [default to "us-east-2"]
**CustomPayload** | Pointer to **string** | Secret payload to be sent with rotation request (relevant only for rotator-type&#x3D;custom) | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] [default to "default_metadata"]
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Secret name | 
**NewMetadata** | Pointer to **string** | Deprecated - use description | [optional] [default to "default_metadata"]
**NewName** | Pointer to **string** | New item name | [optional] 
**NewVersion** | Pointer to **bool** | Deprecated | [optional] 
**RmTag** | Pointer to **[]string** | List of the existent tags that will be removed from this item | [optional] 
**RotateAfterDisconnect** | Pointer to **string** | Rotate the value of the secret after SRA session ends [true/false] | [optional] [default to "false"]
**RotatedPassword** | Pointer to **string** | rotated-username password | [optional] 
**RotatedUsername** | Pointer to **string** | username to be rotated, if selected use-self-creds at rotator-creds-type, this username will try to rotate it&#39;s own password, if use-target-creds is selected, target credentials will be use to rotate the rotated-password | [optional] 
**RotationHour** | Pointer to **int32** | The Hour of the rotation in UTC | [optional] 
**RotationInterval** | Pointer to **string** | The number of days to wait between every automatic key rotation (7-365) | [optional] 
**RotatorCredsType** | Pointer to **string** | The credentials to connect with use-self-creds/use-target-creds | [optional] [default to "use-self-creds"]
**RotatorCustomCmd** | Pointer to **string** | \&quot;Custom rotation command (relevant only for ssh target) | [optional] 
**SamePassword** | Pointer to **string** | Rotate same password for each host from the Linked Target (relevant only for Linked Target) | [optional] 
**SecureAccessAllowExternalUser** | Pointer to **bool** | Allow providing external user for a domain users (relevant only for rdp) | [optional] [default to false]
**SecureAccessAwsAccountId** | Pointer to **string** | The AWS account id (relevant only for aws) | [optional] 
**SecureAccessAwsNativeCli** | Pointer to **bool** | The AWS native cli | [optional] 
**SecureAccessBastionIssuer** | Pointer to **string** | Path to the SSH Certificate Issuer for your Akeyless Bastion | [optional] 
**SecureAccessDbName** | Pointer to **string** | The DB name (relevant only for DB Dynamic-Secret) | [optional] 
**SecureAccessDbSchema** | Pointer to **string** | The db schema (relevant only for mssql or postgresql) | [optional] 
**SecureAccessEnable** | Pointer to **string** | Enable/Disable secure remote access [true/false] | [optional] 
**SecureAccessHost** | Pointer to **[]string** | Target servers for connections (In case of Linked Target association, host(s) will inherit Linked Target hosts - Relevant only for Dynamic Secrets/producers) | [optional] 
**SecureAccessRdpDomain** | Pointer to **string** | Required when the Dynamic Secret is used for a domain user (relevant only for RDP Dynamic-Secret) | [optional] 
**SecureAccessRdpUser** | Pointer to **string** | Override the RDP Domain username (relevant only for rdp) | [optional] 
**SecureAccessWeb** | Pointer to **bool** | Enable Web Secure Remote Access | [optional] [default to false]
**SecureAccessWebBrowsing** | Pointer to **bool** | Secure browser via Akeyless Web Access Bastion (relevant only for aws or azure) | [optional] [default to false]
**SecureAccessWebProxy** | Pointer to **bool** | Web-Proxy via Akeyless Web Access Bastion (relevant only for aws or azure) | [optional] [default to false]
**SshPassword** | Pointer to **string** | Deprecated: use RotatedPassword | [optional] 
**SshUsername** | Pointer to **string** | Deprecated: use RotatedUser | [optional] 
**StorageAccountKeyName** | Pointer to **string** | The name of the storage account key to rotate [key1/key2/kerb1/kerb2] | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateRotatedSecret

`func NewUpdateRotatedSecret(name string, ) *UpdateRotatedSecret`

NewUpdateRotatedSecret instantiates a new UpdateRotatedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRotatedSecretWithDefaults

`func NewUpdateRotatedSecretWithDefaults() *UpdateRotatedSecret`

NewUpdateRotatedSecretWithDefaults instantiates a new UpdateRotatedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddTag

`func (o *UpdateRotatedSecret) GetAddTag() []string`

GetAddTag returns the AddTag field if non-nil, zero value otherwise.

### GetAddTagOk

`func (o *UpdateRotatedSecret) GetAddTagOk() (*[]string, bool)`

GetAddTagOk returns a tuple with the AddTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTag

`func (o *UpdateRotatedSecret) SetAddTag(v []string)`

SetAddTag sets AddTag field to given value.

### HasAddTag

`func (o *UpdateRotatedSecret) HasAddTag() bool`

HasAddTag returns a boolean if a field has been set.

### GetApiId

`func (o *UpdateRotatedSecret) GetApiId() string`

GetApiId returns the ApiId field if non-nil, zero value otherwise.

### GetApiIdOk

`func (o *UpdateRotatedSecret) GetApiIdOk() (*string, bool)`

GetApiIdOk returns a tuple with the ApiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiId

`func (o *UpdateRotatedSecret) SetApiId(v string)`

SetApiId sets ApiId field to given value.

### HasApiId

`func (o *UpdateRotatedSecret) HasApiId() bool`

HasApiId returns a boolean if a field has been set.

### GetApiKey

`func (o *UpdateRotatedSecret) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateRotatedSecret) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateRotatedSecret) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *UpdateRotatedSecret) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetAutoRotate

`func (o *UpdateRotatedSecret) GetAutoRotate() string`

GetAutoRotate returns the AutoRotate field if non-nil, zero value otherwise.

### GetAutoRotateOk

`func (o *UpdateRotatedSecret) GetAutoRotateOk() (*string, bool)`

GetAutoRotateOk returns a tuple with the AutoRotate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRotate

`func (o *UpdateRotatedSecret) SetAutoRotate(v string)`

SetAutoRotate sets AutoRotate field to given value.

### HasAutoRotate

`func (o *UpdateRotatedSecret) HasAutoRotate() bool`

HasAutoRotate returns a boolean if a field has been set.

### GetAwsRegion

`func (o *UpdateRotatedSecret) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *UpdateRotatedSecret) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *UpdateRotatedSecret) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *UpdateRotatedSecret) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetCustomPayload

`func (o *UpdateRotatedSecret) GetCustomPayload() string`

GetCustomPayload returns the CustomPayload field if non-nil, zero value otherwise.

### GetCustomPayloadOk

`func (o *UpdateRotatedSecret) GetCustomPayloadOk() (*string, bool)`

GetCustomPayloadOk returns a tuple with the CustomPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPayload

`func (o *UpdateRotatedSecret) SetCustomPayload(v string)`

SetCustomPayload sets CustomPayload field to given value.

### HasCustomPayload

`func (o *UpdateRotatedSecret) HasCustomPayload() bool`

HasCustomPayload returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateRotatedSecret) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRotatedSecret) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRotatedSecret) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRotatedSecret) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcpKey

`func (o *UpdateRotatedSecret) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *UpdateRotatedSecret) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *UpdateRotatedSecret) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *UpdateRotatedSecret) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetJson

`func (o *UpdateRotatedSecret) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateRotatedSecret) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateRotatedSecret) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateRotatedSecret) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateRotatedSecret) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateRotatedSecret) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateRotatedSecret) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateRotatedSecret) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateRotatedSecret) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateRotatedSecret) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateRotatedSecret) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateRotatedSecret) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateRotatedSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRotatedSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRotatedSecret) SetName(v string)`

SetName sets Name field to given value.


### GetNewMetadata

`func (o *UpdateRotatedSecret) GetNewMetadata() string`

GetNewMetadata returns the NewMetadata field if non-nil, zero value otherwise.

### GetNewMetadataOk

`func (o *UpdateRotatedSecret) GetNewMetadataOk() (*string, bool)`

GetNewMetadataOk returns a tuple with the NewMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewMetadata

`func (o *UpdateRotatedSecret) SetNewMetadata(v string)`

SetNewMetadata sets NewMetadata field to given value.

### HasNewMetadata

`func (o *UpdateRotatedSecret) HasNewMetadata() bool`

HasNewMetadata returns a boolean if a field has been set.

### GetNewName

`func (o *UpdateRotatedSecret) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateRotatedSecret) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateRotatedSecret) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateRotatedSecret) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetNewVersion

`func (o *UpdateRotatedSecret) GetNewVersion() bool`

GetNewVersion returns the NewVersion field if non-nil, zero value otherwise.

### GetNewVersionOk

`func (o *UpdateRotatedSecret) GetNewVersionOk() (*bool, bool)`

GetNewVersionOk returns a tuple with the NewVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewVersion

`func (o *UpdateRotatedSecret) SetNewVersion(v bool)`

SetNewVersion sets NewVersion field to given value.

### HasNewVersion

`func (o *UpdateRotatedSecret) HasNewVersion() bool`

HasNewVersion returns a boolean if a field has been set.

### GetRmTag

`func (o *UpdateRotatedSecret) GetRmTag() []string`

GetRmTag returns the RmTag field if non-nil, zero value otherwise.

### GetRmTagOk

`func (o *UpdateRotatedSecret) GetRmTagOk() (*[]string, bool)`

GetRmTagOk returns a tuple with the RmTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRmTag

`func (o *UpdateRotatedSecret) SetRmTag(v []string)`

SetRmTag sets RmTag field to given value.

### HasRmTag

`func (o *UpdateRotatedSecret) HasRmTag() bool`

HasRmTag returns a boolean if a field has been set.

### GetRotateAfterDisconnect

`func (o *UpdateRotatedSecret) GetRotateAfterDisconnect() string`

GetRotateAfterDisconnect returns the RotateAfterDisconnect field if non-nil, zero value otherwise.

### GetRotateAfterDisconnectOk

`func (o *UpdateRotatedSecret) GetRotateAfterDisconnectOk() (*string, bool)`

GetRotateAfterDisconnectOk returns a tuple with the RotateAfterDisconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateAfterDisconnect

`func (o *UpdateRotatedSecret) SetRotateAfterDisconnect(v string)`

SetRotateAfterDisconnect sets RotateAfterDisconnect field to given value.

### HasRotateAfterDisconnect

`func (o *UpdateRotatedSecret) HasRotateAfterDisconnect() bool`

HasRotateAfterDisconnect returns a boolean if a field has been set.

### GetRotatedPassword

`func (o *UpdateRotatedSecret) GetRotatedPassword() string`

GetRotatedPassword returns the RotatedPassword field if non-nil, zero value otherwise.

### GetRotatedPasswordOk

`func (o *UpdateRotatedSecret) GetRotatedPasswordOk() (*string, bool)`

GetRotatedPasswordOk returns a tuple with the RotatedPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedPassword

`func (o *UpdateRotatedSecret) SetRotatedPassword(v string)`

SetRotatedPassword sets RotatedPassword field to given value.

### HasRotatedPassword

`func (o *UpdateRotatedSecret) HasRotatedPassword() bool`

HasRotatedPassword returns a boolean if a field has been set.

### GetRotatedUsername

`func (o *UpdateRotatedSecret) GetRotatedUsername() string`

GetRotatedUsername returns the RotatedUsername field if non-nil, zero value otherwise.

### GetRotatedUsernameOk

`func (o *UpdateRotatedSecret) GetRotatedUsernameOk() (*string, bool)`

GetRotatedUsernameOk returns a tuple with the RotatedUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatedUsername

`func (o *UpdateRotatedSecret) SetRotatedUsername(v string)`

SetRotatedUsername sets RotatedUsername field to given value.

### HasRotatedUsername

`func (o *UpdateRotatedSecret) HasRotatedUsername() bool`

HasRotatedUsername returns a boolean if a field has been set.

### GetRotationHour

`func (o *UpdateRotatedSecret) GetRotationHour() int32`

GetRotationHour returns the RotationHour field if non-nil, zero value otherwise.

### GetRotationHourOk

`func (o *UpdateRotatedSecret) GetRotationHourOk() (*int32, bool)`

GetRotationHourOk returns a tuple with the RotationHour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationHour

`func (o *UpdateRotatedSecret) SetRotationHour(v int32)`

SetRotationHour sets RotationHour field to given value.

### HasRotationHour

`func (o *UpdateRotatedSecret) HasRotationHour() bool`

HasRotationHour returns a boolean if a field has been set.

### GetRotationInterval

`func (o *UpdateRotatedSecret) GetRotationInterval() string`

GetRotationInterval returns the RotationInterval field if non-nil, zero value otherwise.

### GetRotationIntervalOk

`func (o *UpdateRotatedSecret) GetRotationIntervalOk() (*string, bool)`

GetRotationIntervalOk returns a tuple with the RotationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationInterval

`func (o *UpdateRotatedSecret) SetRotationInterval(v string)`

SetRotationInterval sets RotationInterval field to given value.

### HasRotationInterval

`func (o *UpdateRotatedSecret) HasRotationInterval() bool`

HasRotationInterval returns a boolean if a field has been set.

### GetRotatorCredsType

`func (o *UpdateRotatedSecret) GetRotatorCredsType() string`

GetRotatorCredsType returns the RotatorCredsType field if non-nil, zero value otherwise.

### GetRotatorCredsTypeOk

`func (o *UpdateRotatedSecret) GetRotatorCredsTypeOk() (*string, bool)`

GetRotatorCredsTypeOk returns a tuple with the RotatorCredsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCredsType

`func (o *UpdateRotatedSecret) SetRotatorCredsType(v string)`

SetRotatorCredsType sets RotatorCredsType field to given value.

### HasRotatorCredsType

`func (o *UpdateRotatedSecret) HasRotatorCredsType() bool`

HasRotatorCredsType returns a boolean if a field has been set.

### GetRotatorCustomCmd

`func (o *UpdateRotatedSecret) GetRotatorCustomCmd() string`

GetRotatorCustomCmd returns the RotatorCustomCmd field if non-nil, zero value otherwise.

### GetRotatorCustomCmdOk

`func (o *UpdateRotatedSecret) GetRotatorCustomCmdOk() (*string, bool)`

GetRotatorCustomCmdOk returns a tuple with the RotatorCustomCmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotatorCustomCmd

`func (o *UpdateRotatedSecret) SetRotatorCustomCmd(v string)`

SetRotatorCustomCmd sets RotatorCustomCmd field to given value.

### HasRotatorCustomCmd

`func (o *UpdateRotatedSecret) HasRotatorCustomCmd() bool`

HasRotatorCustomCmd returns a boolean if a field has been set.

### GetSamePassword

`func (o *UpdateRotatedSecret) GetSamePassword() string`

GetSamePassword returns the SamePassword field if non-nil, zero value otherwise.

### GetSamePasswordOk

`func (o *UpdateRotatedSecret) GetSamePasswordOk() (*string, bool)`

GetSamePasswordOk returns a tuple with the SamePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamePassword

`func (o *UpdateRotatedSecret) SetSamePassword(v string)`

SetSamePassword sets SamePassword field to given value.

### HasSamePassword

`func (o *UpdateRotatedSecret) HasSamePassword() bool`

HasSamePassword returns a boolean if a field has been set.

### GetSecureAccessAllowExternalUser

`func (o *UpdateRotatedSecret) GetSecureAccessAllowExternalUser() bool`

GetSecureAccessAllowExternalUser returns the SecureAccessAllowExternalUser field if non-nil, zero value otherwise.

### GetSecureAccessAllowExternalUserOk

`func (o *UpdateRotatedSecret) GetSecureAccessAllowExternalUserOk() (*bool, bool)`

GetSecureAccessAllowExternalUserOk returns a tuple with the SecureAccessAllowExternalUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAllowExternalUser

`func (o *UpdateRotatedSecret) SetSecureAccessAllowExternalUser(v bool)`

SetSecureAccessAllowExternalUser sets SecureAccessAllowExternalUser field to given value.

### HasSecureAccessAllowExternalUser

`func (o *UpdateRotatedSecret) HasSecureAccessAllowExternalUser() bool`

HasSecureAccessAllowExternalUser returns a boolean if a field has been set.

### GetSecureAccessAwsAccountId

`func (o *UpdateRotatedSecret) GetSecureAccessAwsAccountId() string`

GetSecureAccessAwsAccountId returns the SecureAccessAwsAccountId field if non-nil, zero value otherwise.

### GetSecureAccessAwsAccountIdOk

`func (o *UpdateRotatedSecret) GetSecureAccessAwsAccountIdOk() (*string, bool)`

GetSecureAccessAwsAccountIdOk returns a tuple with the SecureAccessAwsAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsAccountId

`func (o *UpdateRotatedSecret) SetSecureAccessAwsAccountId(v string)`

SetSecureAccessAwsAccountId sets SecureAccessAwsAccountId field to given value.

### HasSecureAccessAwsAccountId

`func (o *UpdateRotatedSecret) HasSecureAccessAwsAccountId() bool`

HasSecureAccessAwsAccountId returns a boolean if a field has been set.

### GetSecureAccessAwsNativeCli

`func (o *UpdateRotatedSecret) GetSecureAccessAwsNativeCli() bool`

GetSecureAccessAwsNativeCli returns the SecureAccessAwsNativeCli field if non-nil, zero value otherwise.

### GetSecureAccessAwsNativeCliOk

`func (o *UpdateRotatedSecret) GetSecureAccessAwsNativeCliOk() (*bool, bool)`

GetSecureAccessAwsNativeCliOk returns a tuple with the SecureAccessAwsNativeCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessAwsNativeCli

`func (o *UpdateRotatedSecret) SetSecureAccessAwsNativeCli(v bool)`

SetSecureAccessAwsNativeCli sets SecureAccessAwsNativeCli field to given value.

### HasSecureAccessAwsNativeCli

`func (o *UpdateRotatedSecret) HasSecureAccessAwsNativeCli() bool`

HasSecureAccessAwsNativeCli returns a boolean if a field has been set.

### GetSecureAccessBastionIssuer

`func (o *UpdateRotatedSecret) GetSecureAccessBastionIssuer() string`

GetSecureAccessBastionIssuer returns the SecureAccessBastionIssuer field if non-nil, zero value otherwise.

### GetSecureAccessBastionIssuerOk

`func (o *UpdateRotatedSecret) GetSecureAccessBastionIssuerOk() (*string, bool)`

GetSecureAccessBastionIssuerOk returns a tuple with the SecureAccessBastionIssuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessBastionIssuer

`func (o *UpdateRotatedSecret) SetSecureAccessBastionIssuer(v string)`

SetSecureAccessBastionIssuer sets SecureAccessBastionIssuer field to given value.

### HasSecureAccessBastionIssuer

`func (o *UpdateRotatedSecret) HasSecureAccessBastionIssuer() bool`

HasSecureAccessBastionIssuer returns a boolean if a field has been set.

### GetSecureAccessDbName

`func (o *UpdateRotatedSecret) GetSecureAccessDbName() string`

GetSecureAccessDbName returns the SecureAccessDbName field if non-nil, zero value otherwise.

### GetSecureAccessDbNameOk

`func (o *UpdateRotatedSecret) GetSecureAccessDbNameOk() (*string, bool)`

GetSecureAccessDbNameOk returns a tuple with the SecureAccessDbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbName

`func (o *UpdateRotatedSecret) SetSecureAccessDbName(v string)`

SetSecureAccessDbName sets SecureAccessDbName field to given value.

### HasSecureAccessDbName

`func (o *UpdateRotatedSecret) HasSecureAccessDbName() bool`

HasSecureAccessDbName returns a boolean if a field has been set.

### GetSecureAccessDbSchema

`func (o *UpdateRotatedSecret) GetSecureAccessDbSchema() string`

GetSecureAccessDbSchema returns the SecureAccessDbSchema field if non-nil, zero value otherwise.

### GetSecureAccessDbSchemaOk

`func (o *UpdateRotatedSecret) GetSecureAccessDbSchemaOk() (*string, bool)`

GetSecureAccessDbSchemaOk returns a tuple with the SecureAccessDbSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessDbSchema

`func (o *UpdateRotatedSecret) SetSecureAccessDbSchema(v string)`

SetSecureAccessDbSchema sets SecureAccessDbSchema field to given value.

### HasSecureAccessDbSchema

`func (o *UpdateRotatedSecret) HasSecureAccessDbSchema() bool`

HasSecureAccessDbSchema returns a boolean if a field has been set.

### GetSecureAccessEnable

`func (o *UpdateRotatedSecret) GetSecureAccessEnable() string`

GetSecureAccessEnable returns the SecureAccessEnable field if non-nil, zero value otherwise.

### GetSecureAccessEnableOk

`func (o *UpdateRotatedSecret) GetSecureAccessEnableOk() (*string, bool)`

GetSecureAccessEnableOk returns a tuple with the SecureAccessEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessEnable

`func (o *UpdateRotatedSecret) SetSecureAccessEnable(v string)`

SetSecureAccessEnable sets SecureAccessEnable field to given value.

### HasSecureAccessEnable

`func (o *UpdateRotatedSecret) HasSecureAccessEnable() bool`

HasSecureAccessEnable returns a boolean if a field has been set.

### GetSecureAccessHost

`func (o *UpdateRotatedSecret) GetSecureAccessHost() []string`

GetSecureAccessHost returns the SecureAccessHost field if non-nil, zero value otherwise.

### GetSecureAccessHostOk

`func (o *UpdateRotatedSecret) GetSecureAccessHostOk() (*[]string, bool)`

GetSecureAccessHostOk returns a tuple with the SecureAccessHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessHost

`func (o *UpdateRotatedSecret) SetSecureAccessHost(v []string)`

SetSecureAccessHost sets SecureAccessHost field to given value.

### HasSecureAccessHost

`func (o *UpdateRotatedSecret) HasSecureAccessHost() bool`

HasSecureAccessHost returns a boolean if a field has been set.

### GetSecureAccessRdpDomain

`func (o *UpdateRotatedSecret) GetSecureAccessRdpDomain() string`

GetSecureAccessRdpDomain returns the SecureAccessRdpDomain field if non-nil, zero value otherwise.

### GetSecureAccessRdpDomainOk

`func (o *UpdateRotatedSecret) GetSecureAccessRdpDomainOk() (*string, bool)`

GetSecureAccessRdpDomainOk returns a tuple with the SecureAccessRdpDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpDomain

`func (o *UpdateRotatedSecret) SetSecureAccessRdpDomain(v string)`

SetSecureAccessRdpDomain sets SecureAccessRdpDomain field to given value.

### HasSecureAccessRdpDomain

`func (o *UpdateRotatedSecret) HasSecureAccessRdpDomain() bool`

HasSecureAccessRdpDomain returns a boolean if a field has been set.

### GetSecureAccessRdpUser

`func (o *UpdateRotatedSecret) GetSecureAccessRdpUser() string`

GetSecureAccessRdpUser returns the SecureAccessRdpUser field if non-nil, zero value otherwise.

### GetSecureAccessRdpUserOk

`func (o *UpdateRotatedSecret) GetSecureAccessRdpUserOk() (*string, bool)`

GetSecureAccessRdpUserOk returns a tuple with the SecureAccessRdpUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessRdpUser

`func (o *UpdateRotatedSecret) SetSecureAccessRdpUser(v string)`

SetSecureAccessRdpUser sets SecureAccessRdpUser field to given value.

### HasSecureAccessRdpUser

`func (o *UpdateRotatedSecret) HasSecureAccessRdpUser() bool`

HasSecureAccessRdpUser returns a boolean if a field has been set.

### GetSecureAccessWeb

`func (o *UpdateRotatedSecret) GetSecureAccessWeb() bool`

GetSecureAccessWeb returns the SecureAccessWeb field if non-nil, zero value otherwise.

### GetSecureAccessWebOk

`func (o *UpdateRotatedSecret) GetSecureAccessWebOk() (*bool, bool)`

GetSecureAccessWebOk returns a tuple with the SecureAccessWeb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWeb

`func (o *UpdateRotatedSecret) SetSecureAccessWeb(v bool)`

SetSecureAccessWeb sets SecureAccessWeb field to given value.

### HasSecureAccessWeb

`func (o *UpdateRotatedSecret) HasSecureAccessWeb() bool`

HasSecureAccessWeb returns a boolean if a field has been set.

### GetSecureAccessWebBrowsing

`func (o *UpdateRotatedSecret) GetSecureAccessWebBrowsing() bool`

GetSecureAccessWebBrowsing returns the SecureAccessWebBrowsing field if non-nil, zero value otherwise.

### GetSecureAccessWebBrowsingOk

`func (o *UpdateRotatedSecret) GetSecureAccessWebBrowsingOk() (*bool, bool)`

GetSecureAccessWebBrowsingOk returns a tuple with the SecureAccessWebBrowsing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebBrowsing

`func (o *UpdateRotatedSecret) SetSecureAccessWebBrowsing(v bool)`

SetSecureAccessWebBrowsing sets SecureAccessWebBrowsing field to given value.

### HasSecureAccessWebBrowsing

`func (o *UpdateRotatedSecret) HasSecureAccessWebBrowsing() bool`

HasSecureAccessWebBrowsing returns a boolean if a field has been set.

### GetSecureAccessWebProxy

`func (o *UpdateRotatedSecret) GetSecureAccessWebProxy() bool`

GetSecureAccessWebProxy returns the SecureAccessWebProxy field if non-nil, zero value otherwise.

### GetSecureAccessWebProxyOk

`func (o *UpdateRotatedSecret) GetSecureAccessWebProxyOk() (*bool, bool)`

GetSecureAccessWebProxyOk returns a tuple with the SecureAccessWebProxy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecureAccessWebProxy

`func (o *UpdateRotatedSecret) SetSecureAccessWebProxy(v bool)`

SetSecureAccessWebProxy sets SecureAccessWebProxy field to given value.

### HasSecureAccessWebProxy

`func (o *UpdateRotatedSecret) HasSecureAccessWebProxy() bool`

HasSecureAccessWebProxy returns a boolean if a field has been set.

### GetSshPassword

`func (o *UpdateRotatedSecret) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *UpdateRotatedSecret) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *UpdateRotatedSecret) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *UpdateRotatedSecret) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshUsername

`func (o *UpdateRotatedSecret) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *UpdateRotatedSecret) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *UpdateRotatedSecret) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *UpdateRotatedSecret) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetStorageAccountKeyName

`func (o *UpdateRotatedSecret) GetStorageAccountKeyName() string`

GetStorageAccountKeyName returns the StorageAccountKeyName field if non-nil, zero value otherwise.

### GetStorageAccountKeyNameOk

`func (o *UpdateRotatedSecret) GetStorageAccountKeyNameOk() (*string, bool)`

GetStorageAccountKeyNameOk returns a tuple with the StorageAccountKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccountKeyName

`func (o *UpdateRotatedSecret) SetStorageAccountKeyName(v string)`

SetStorageAccountKeyName sets StorageAccountKeyName field to given value.

### HasStorageAccountKeyName

`func (o *UpdateRotatedSecret) HasStorageAccountKeyName() bool`

HasStorageAccountKeyName returns a boolean if a field has been set.

### GetToken

`func (o *UpdateRotatedSecret) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateRotatedSecret) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateRotatedSecret) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateRotatedSecret) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateRotatedSecret) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateRotatedSecret) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateRotatedSecret) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateRotatedSecret) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



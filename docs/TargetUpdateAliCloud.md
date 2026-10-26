# TargetUpdateAliCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AliCloud RAM access key secret | [optional] 
**AccessKeyId** | Pointer to **string** | AliCloud RAM access key ID | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GenerateExternalId** | Pointer to **bool** | A unique auto-generated value used in your AliCloud account when configuring the RAM role trust policy to securely delegate access to Akeyless. Relevant only when using GW cloud ID | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Region** | Pointer to **string** | AliCloud region. Selects the regional STS endpoint. If empty, the global STS endpoint is used | [optional] 
**RoleArn** | Pointer to **string** | AliCloud RAM role ARN that the Gateway will assume in your AliCloud account (acs:ram::&lt;account-id&gt;:role/&lt;role-name&gt;), relevant only when using external ID | [optional] 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**SessionToken** | Pointer to **string** | Required only for temporary security credentials retrieved using STS | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** | Use the Gateway cloud identity | [optional] 

## Methods

### NewTargetUpdateAliCloud

`func NewTargetUpdateAliCloud(name string, ) *TargetUpdateAliCloud`

NewTargetUpdateAliCloud instantiates a new TargetUpdateAliCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateAliCloudWithDefaults

`func NewTargetUpdateAliCloudWithDefaults() *TargetUpdateAliCloud`

NewTargetUpdateAliCloudWithDefaults instantiates a new TargetUpdateAliCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *TargetUpdateAliCloud) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *TargetUpdateAliCloud) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *TargetUpdateAliCloud) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *TargetUpdateAliCloud) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *TargetUpdateAliCloud) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *TargetUpdateAliCloud) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *TargetUpdateAliCloud) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *TargetUpdateAliCloud) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *TargetUpdateAliCloud) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *TargetUpdateAliCloud) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *TargetUpdateAliCloud) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *TargetUpdateAliCloud) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateAliCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateAliCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateAliCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateAliCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenerateExternalId

`func (o *TargetUpdateAliCloud) GetGenerateExternalId() bool`

GetGenerateExternalId returns the GenerateExternalId field if non-nil, zero value otherwise.

### GetGenerateExternalIdOk

`func (o *TargetUpdateAliCloud) GetGenerateExternalIdOk() (*bool, bool)`

GetGenerateExternalIdOk returns a tuple with the GenerateExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateExternalId

`func (o *TargetUpdateAliCloud) SetGenerateExternalId(v bool)`

SetGenerateExternalId sets GenerateExternalId field to given value.

### HasGenerateExternalId

`func (o *TargetUpdateAliCloud) HasGenerateExternalId() bool`

HasGenerateExternalId returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateAliCloud) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateAliCloud) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateAliCloud) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateAliCloud) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateAliCloud) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateAliCloud) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateAliCloud) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateAliCloud) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateAliCloud) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateAliCloud) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateAliCloud) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateAliCloud) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *TargetUpdateAliCloud) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *TargetUpdateAliCloud) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *TargetUpdateAliCloud) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *TargetUpdateAliCloud) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *TargetUpdateAliCloud) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *TargetUpdateAliCloud) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *TargetUpdateAliCloud) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *TargetUpdateAliCloud) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateAliCloud) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateAliCloud) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateAliCloud) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateAliCloud) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateAliCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateAliCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateAliCloud) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateAliCloud) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateAliCloud) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateAliCloud) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateAliCloud) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRegion

`func (o *TargetUpdateAliCloud) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TargetUpdateAliCloud) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TargetUpdateAliCloud) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TargetUpdateAliCloud) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleArn

`func (o *TargetUpdateAliCloud) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *TargetUpdateAliCloud) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *TargetUpdateAliCloud) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *TargetUpdateAliCloud) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *TargetUpdateAliCloud) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *TargetUpdateAliCloud) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *TargetUpdateAliCloud) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *TargetUpdateAliCloud) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetSessionToken

`func (o *TargetUpdateAliCloud) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *TargetUpdateAliCloud) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *TargetUpdateAliCloud) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *TargetUpdateAliCloud) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateAliCloud) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateAliCloud) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateAliCloud) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateAliCloud) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateAliCloud) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateAliCloud) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateAliCloud) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateAliCloud) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *TargetUpdateAliCloud) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *TargetUpdateAliCloud) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *TargetUpdateAliCloud) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *TargetUpdateAliCloud) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



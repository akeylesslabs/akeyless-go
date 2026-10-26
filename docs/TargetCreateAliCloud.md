# TargetCreateAliCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | AliCloud RAM access key secret | [optional] 
**AccessKeyId** | Pointer to **string** | AliCloud RAM access key ID | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GenerateExternalId** | Pointer to **bool** | A unique auto-generated value used in your AliCloud account when configuring the RAM role trust policy to securely delegate access to Akeyless. Relevant only when using GW cloud ID | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**LockOnRead** | Pointer to **string** | Lock this secret after each successful value read | [optional] 
**LockTtl** | Pointer to **string** | Lock TTL in minutes | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Region** | Pointer to **string** | AliCloud region. Selects the regional STS endpoint. If empty, the global STS endpoint is used | [optional] 
**RoleArn** | Pointer to **string** | AliCloud RAM role ARN that the Gateway will assume in your AliCloud account (acs:ram::&lt;account-id&gt;:role/&lt;role-name&gt;), relevant only when using external ID | [optional] 
**RotateOnUnlock** | Pointer to **string** | Rotate this secret after it is unlocked | [optional] 
**SessionToken** | Pointer to **string** | Required only for temporary security credentials retrieved using STS | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UseGwCloudIdentity** | Pointer to **bool** | Use the Gateway cloud identity | [optional] 

## Methods

### NewTargetCreateAliCloud

`func NewTargetCreateAliCloud(name string, ) *TargetCreateAliCloud`

NewTargetCreateAliCloud instantiates a new TargetCreateAliCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateAliCloudWithDefaults

`func NewTargetCreateAliCloudWithDefaults() *TargetCreateAliCloud`

NewTargetCreateAliCloudWithDefaults instantiates a new TargetCreateAliCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *TargetCreateAliCloud) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *TargetCreateAliCloud) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *TargetCreateAliCloud) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *TargetCreateAliCloud) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetAccessKeyId

`func (o *TargetCreateAliCloud) GetAccessKeyId() string`

GetAccessKeyId returns the AccessKeyId field if non-nil, zero value otherwise.

### GetAccessKeyIdOk

`func (o *TargetCreateAliCloud) GetAccessKeyIdOk() (*string, bool)`

GetAccessKeyIdOk returns a tuple with the AccessKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKeyId

`func (o *TargetCreateAliCloud) SetAccessKeyId(v string)`

SetAccessKeyId sets AccessKeyId field to given value.

### HasAccessKeyId

`func (o *TargetCreateAliCloud) HasAccessKeyId() bool`

HasAccessKeyId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *TargetCreateAliCloud) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *TargetCreateAliCloud) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *TargetCreateAliCloud) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *TargetCreateAliCloud) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateAliCloud) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateAliCloud) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateAliCloud) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateAliCloud) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGenerateExternalId

`func (o *TargetCreateAliCloud) GetGenerateExternalId() bool`

GetGenerateExternalId returns the GenerateExternalId field if non-nil, zero value otherwise.

### GetGenerateExternalIdOk

`func (o *TargetCreateAliCloud) GetGenerateExternalIdOk() (*bool, bool)`

GetGenerateExternalIdOk returns a tuple with the GenerateExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateExternalId

`func (o *TargetCreateAliCloud) SetGenerateExternalId(v bool)`

SetGenerateExternalId sets GenerateExternalId field to given value.

### HasGenerateExternalId

`func (o *TargetCreateAliCloud) HasGenerateExternalId() bool`

HasGenerateExternalId returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateAliCloud) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateAliCloud) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateAliCloud) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateAliCloud) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateAliCloud) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateAliCloud) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateAliCloud) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateAliCloud) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLockOnRead

`func (o *TargetCreateAliCloud) GetLockOnRead() string`

GetLockOnRead returns the LockOnRead field if non-nil, zero value otherwise.

### GetLockOnReadOk

`func (o *TargetCreateAliCloud) GetLockOnReadOk() (*string, bool)`

GetLockOnReadOk returns a tuple with the LockOnRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockOnRead

`func (o *TargetCreateAliCloud) SetLockOnRead(v string)`

SetLockOnRead sets LockOnRead field to given value.

### HasLockOnRead

`func (o *TargetCreateAliCloud) HasLockOnRead() bool`

HasLockOnRead returns a boolean if a field has been set.

### GetLockTtl

`func (o *TargetCreateAliCloud) GetLockTtl() string`

GetLockTtl returns the LockTtl field if non-nil, zero value otherwise.

### GetLockTtlOk

`func (o *TargetCreateAliCloud) GetLockTtlOk() (*string, bool)`

GetLockTtlOk returns a tuple with the LockTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTtl

`func (o *TargetCreateAliCloud) SetLockTtl(v string)`

SetLockTtl sets LockTtl field to given value.

### HasLockTtl

`func (o *TargetCreateAliCloud) HasLockTtl() bool`

HasLockTtl returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateAliCloud) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateAliCloud) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateAliCloud) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateAliCloud) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateAliCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateAliCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateAliCloud) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *TargetCreateAliCloud) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TargetCreateAliCloud) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TargetCreateAliCloud) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TargetCreateAliCloud) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetRoleArn

`func (o *TargetCreateAliCloud) GetRoleArn() string`

GetRoleArn returns the RoleArn field if non-nil, zero value otherwise.

### GetRoleArnOk

`func (o *TargetCreateAliCloud) GetRoleArnOk() (*string, bool)`

GetRoleArnOk returns a tuple with the RoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleArn

`func (o *TargetCreateAliCloud) SetRoleArn(v string)`

SetRoleArn sets RoleArn field to given value.

### HasRoleArn

`func (o *TargetCreateAliCloud) HasRoleArn() bool`

HasRoleArn returns a boolean if a field has been set.

### GetRotateOnUnlock

`func (o *TargetCreateAliCloud) GetRotateOnUnlock() string`

GetRotateOnUnlock returns the RotateOnUnlock field if non-nil, zero value otherwise.

### GetRotateOnUnlockOk

`func (o *TargetCreateAliCloud) GetRotateOnUnlockOk() (*string, bool)`

GetRotateOnUnlockOk returns a tuple with the RotateOnUnlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateOnUnlock

`func (o *TargetCreateAliCloud) SetRotateOnUnlock(v string)`

SetRotateOnUnlock sets RotateOnUnlock field to given value.

### HasRotateOnUnlock

`func (o *TargetCreateAliCloud) HasRotateOnUnlock() bool`

HasRotateOnUnlock returns a boolean if a field has been set.

### GetSessionToken

`func (o *TargetCreateAliCloud) GetSessionToken() string`

GetSessionToken returns the SessionToken field if non-nil, zero value otherwise.

### GetSessionTokenOk

`func (o *TargetCreateAliCloud) GetSessionTokenOk() (*string, bool)`

GetSessionTokenOk returns a tuple with the SessionToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionToken

`func (o *TargetCreateAliCloud) SetSessionToken(v string)`

SetSessionToken sets SessionToken field to given value.

### HasSessionToken

`func (o *TargetCreateAliCloud) HasSessionToken() bool`

HasSessionToken returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateAliCloud) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateAliCloud) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateAliCloud) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateAliCloud) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateAliCloud) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateAliCloud) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateAliCloud) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateAliCloud) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUseGwCloudIdentity

`func (o *TargetCreateAliCloud) GetUseGwCloudIdentity() bool`

GetUseGwCloudIdentity returns the UseGwCloudIdentity field if non-nil, zero value otherwise.

### GetUseGwCloudIdentityOk

`func (o *TargetCreateAliCloud) GetUseGwCloudIdentityOk() (*bool, bool)`

GetUseGwCloudIdentityOk returns a tuple with the UseGwCloudIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseGwCloudIdentity

`func (o *TargetCreateAliCloud) SetUseGwCloudIdentity(v bool)`

SetUseGwCloudIdentity sets UseGwCloudIdentity field to given value.

### HasUseGwCloudIdentity

`func (o *TargetCreateAliCloud) HasUseGwCloudIdentity() bool`

HasUseGwCloudIdentity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



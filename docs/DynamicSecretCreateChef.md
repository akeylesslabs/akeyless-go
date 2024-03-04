# DynamicSecretCreateChef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChefOrgs** | Pointer to **string** | Organizations | [optional] 
**ChefServerKey** | Pointer to **string** | Server key | [optional] 
**ChefServerUrl** | Pointer to **string** | Server URL | [optional] 
**ChefServerUsername** | Pointer to **string** | Server username | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**PasswordLength** | Pointer to **string** | The length of the password to be generated | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SkipSsl** | Pointer to **bool** | Skip SSL | [optional] [default to true]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateChef

`func NewDynamicSecretCreateChef(name string, ) *DynamicSecretCreateChef`

NewDynamicSecretCreateChef instantiates a new DynamicSecretCreateChef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateChefWithDefaults

`func NewDynamicSecretCreateChefWithDefaults() *DynamicSecretCreateChef`

NewDynamicSecretCreateChefWithDefaults instantiates a new DynamicSecretCreateChef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChefOrgs

`func (o *DynamicSecretCreateChef) GetChefOrgs() string`

GetChefOrgs returns the ChefOrgs field if non-nil, zero value otherwise.

### GetChefOrgsOk

`func (o *DynamicSecretCreateChef) GetChefOrgsOk() (*string, bool)`

GetChefOrgsOk returns a tuple with the ChefOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefOrgs

`func (o *DynamicSecretCreateChef) SetChefOrgs(v string)`

SetChefOrgs sets ChefOrgs field to given value.

### HasChefOrgs

`func (o *DynamicSecretCreateChef) HasChefOrgs() bool`

HasChefOrgs returns a boolean if a field has been set.

### GetChefServerKey

`func (o *DynamicSecretCreateChef) GetChefServerKey() string`

GetChefServerKey returns the ChefServerKey field if non-nil, zero value otherwise.

### GetChefServerKeyOk

`func (o *DynamicSecretCreateChef) GetChefServerKeyOk() (*string, bool)`

GetChefServerKeyOk returns a tuple with the ChefServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerKey

`func (o *DynamicSecretCreateChef) SetChefServerKey(v string)`

SetChefServerKey sets ChefServerKey field to given value.

### HasChefServerKey

`func (o *DynamicSecretCreateChef) HasChefServerKey() bool`

HasChefServerKey returns a boolean if a field has been set.

### GetChefServerUrl

`func (o *DynamicSecretCreateChef) GetChefServerUrl() string`

GetChefServerUrl returns the ChefServerUrl field if non-nil, zero value otherwise.

### GetChefServerUrlOk

`func (o *DynamicSecretCreateChef) GetChefServerUrlOk() (*string, bool)`

GetChefServerUrlOk returns a tuple with the ChefServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUrl

`func (o *DynamicSecretCreateChef) SetChefServerUrl(v string)`

SetChefServerUrl sets ChefServerUrl field to given value.

### HasChefServerUrl

`func (o *DynamicSecretCreateChef) HasChefServerUrl() bool`

HasChefServerUrl returns a boolean if a field has been set.

### GetChefServerUsername

`func (o *DynamicSecretCreateChef) GetChefServerUsername() string`

GetChefServerUsername returns the ChefServerUsername field if non-nil, zero value otherwise.

### GetChefServerUsernameOk

`func (o *DynamicSecretCreateChef) GetChefServerUsernameOk() (*string, bool)`

GetChefServerUsernameOk returns a tuple with the ChefServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUsername

`func (o *DynamicSecretCreateChef) SetChefServerUsername(v string)`

SetChefServerUsername sets ChefServerUsername field to given value.

### HasChefServerUsername

`func (o *DynamicSecretCreateChef) HasChefServerUsername() bool`

HasChefServerUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateChef) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateChef) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateChef) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateChef) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateChef) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateChef) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateChef) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateChef) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateChef) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateChef) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateChef) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateChef) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateChef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateChef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateChef) SetName(v string)`

SetName sets Name field to given value.


### GetPasswordLength

`func (o *DynamicSecretCreateChef) GetPasswordLength() string`

GetPasswordLength returns the PasswordLength field if non-nil, zero value otherwise.

### GetPasswordLengthOk

`func (o *DynamicSecretCreateChef) GetPasswordLengthOk() (*string, bool)`

GetPasswordLengthOk returns a tuple with the PasswordLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordLength

`func (o *DynamicSecretCreateChef) SetPasswordLength(v string)`

SetPasswordLength sets PasswordLength field to given value.

### HasPasswordLength

`func (o *DynamicSecretCreateChef) HasPasswordLength() bool`

HasPasswordLength returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateChef) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateChef) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateChef) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateChef) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSkipSsl

`func (o *DynamicSecretCreateChef) GetSkipSsl() bool`

GetSkipSsl returns the SkipSsl field if non-nil, zero value otherwise.

### GetSkipSslOk

`func (o *DynamicSecretCreateChef) GetSkipSslOk() (*bool, bool)`

GetSkipSslOk returns a tuple with the SkipSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSsl

`func (o *DynamicSecretCreateChef) SetSkipSsl(v bool)`

SetSkipSsl sets SkipSsl field to given value.

### HasSkipSsl

`func (o *DynamicSecretCreateChef) HasSkipSsl() bool`

HasSkipSsl returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateChef) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateChef) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateChef) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateChef) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateChef) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateChef) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateChef) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateChef) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateChef) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateChef) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateChef) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateChef) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateChef) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateChef) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateChef) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateChef) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateChef) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateChef) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateChef) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateChef) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



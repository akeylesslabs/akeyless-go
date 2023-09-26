# DynamicSecretUpdateChef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChefOrgs** | Pointer to **string** | Organizations | [optional] 
**ChefServerKey** | Pointer to **string** | Server key | [optional] 
**ChefServerUrl** | Pointer to **string** | Server URL | [optional] 
**ChefServerUsername** | Pointer to **string** | Server username | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SkipSsl** | Pointer to **bool** | Skip SSL | [optional] [default to true]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateChef

`func NewDynamicSecretUpdateChef(name string, ) *DynamicSecretUpdateChef`

NewDynamicSecretUpdateChef instantiates a new DynamicSecretUpdateChef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateChefWithDefaults

`func NewDynamicSecretUpdateChefWithDefaults() *DynamicSecretUpdateChef`

NewDynamicSecretUpdateChefWithDefaults instantiates a new DynamicSecretUpdateChef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChefOrgs

`func (o *DynamicSecretUpdateChef) GetChefOrgs() string`

GetChefOrgs returns the ChefOrgs field if non-nil, zero value otherwise.

### GetChefOrgsOk

`func (o *DynamicSecretUpdateChef) GetChefOrgsOk() (*string, bool)`

GetChefOrgsOk returns a tuple with the ChefOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefOrgs

`func (o *DynamicSecretUpdateChef) SetChefOrgs(v string)`

SetChefOrgs sets ChefOrgs field to given value.

### HasChefOrgs

`func (o *DynamicSecretUpdateChef) HasChefOrgs() bool`

HasChefOrgs returns a boolean if a field has been set.

### GetChefServerKey

`func (o *DynamicSecretUpdateChef) GetChefServerKey() string`

GetChefServerKey returns the ChefServerKey field if non-nil, zero value otherwise.

### GetChefServerKeyOk

`func (o *DynamicSecretUpdateChef) GetChefServerKeyOk() (*string, bool)`

GetChefServerKeyOk returns a tuple with the ChefServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerKey

`func (o *DynamicSecretUpdateChef) SetChefServerKey(v string)`

SetChefServerKey sets ChefServerKey field to given value.

### HasChefServerKey

`func (o *DynamicSecretUpdateChef) HasChefServerKey() bool`

HasChefServerKey returns a boolean if a field has been set.

### GetChefServerUrl

`func (o *DynamicSecretUpdateChef) GetChefServerUrl() string`

GetChefServerUrl returns the ChefServerUrl field if non-nil, zero value otherwise.

### GetChefServerUrlOk

`func (o *DynamicSecretUpdateChef) GetChefServerUrlOk() (*string, bool)`

GetChefServerUrlOk returns a tuple with the ChefServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUrl

`func (o *DynamicSecretUpdateChef) SetChefServerUrl(v string)`

SetChefServerUrl sets ChefServerUrl field to given value.

### HasChefServerUrl

`func (o *DynamicSecretUpdateChef) HasChefServerUrl() bool`

HasChefServerUrl returns a boolean if a field has been set.

### GetChefServerUsername

`func (o *DynamicSecretUpdateChef) GetChefServerUsername() string`

GetChefServerUsername returns the ChefServerUsername field if non-nil, zero value otherwise.

### GetChefServerUsernameOk

`func (o *DynamicSecretUpdateChef) GetChefServerUsernameOk() (*string, bool)`

GetChefServerUsernameOk returns a tuple with the ChefServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUsername

`func (o *DynamicSecretUpdateChef) SetChefServerUsername(v string)`

SetChefServerUsername sets ChefServerUsername field to given value.

### HasChefServerUsername

`func (o *DynamicSecretUpdateChef) HasChefServerUsername() bool`

HasChefServerUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateChef) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateChef) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateChef) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateChef) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateChef) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateChef) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateChef) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateChef) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateChef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateChef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateChef) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateChef) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateChef) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateChef) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateChef) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateChef) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateChef) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateChef) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateChef) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSkipSsl

`func (o *DynamicSecretUpdateChef) GetSkipSsl() bool`

GetSkipSsl returns the SkipSsl field if non-nil, zero value otherwise.

### GetSkipSslOk

`func (o *DynamicSecretUpdateChef) GetSkipSslOk() (*bool, bool)`

GetSkipSslOk returns a tuple with the SkipSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSsl

`func (o *DynamicSecretUpdateChef) SetSkipSsl(v bool)`

SetSkipSsl sets SkipSsl field to given value.

### HasSkipSsl

`func (o *DynamicSecretUpdateChef) HasSkipSsl() bool`

HasSkipSsl returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateChef) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateChef) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateChef) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateChef) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateChef) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateChef) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateChef) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateChef) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateChef) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateChef) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateChef) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateChef) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateChef) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateChef) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateChef) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateChef) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateChef) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateChef) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateChef) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateChef) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



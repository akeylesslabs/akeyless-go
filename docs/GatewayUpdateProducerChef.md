# GatewayUpdateProducerChef

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
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**SkipSsl** | Pointer to **bool** | Skip SSL | [optional] [default to true]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayUpdateProducerChef

`func NewGatewayUpdateProducerChef(name string, ) *GatewayUpdateProducerChef`

NewGatewayUpdateProducerChef instantiates a new GatewayUpdateProducerChef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerChefWithDefaults

`func NewGatewayUpdateProducerChefWithDefaults() *GatewayUpdateProducerChef`

NewGatewayUpdateProducerChefWithDefaults instantiates a new GatewayUpdateProducerChef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChefOrgs

`func (o *GatewayUpdateProducerChef) GetChefOrgs() string`

GetChefOrgs returns the ChefOrgs field if non-nil, zero value otherwise.

### GetChefOrgsOk

`func (o *GatewayUpdateProducerChef) GetChefOrgsOk() (*string, bool)`

GetChefOrgsOk returns a tuple with the ChefOrgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefOrgs

`func (o *GatewayUpdateProducerChef) SetChefOrgs(v string)`

SetChefOrgs sets ChefOrgs field to given value.

### HasChefOrgs

`func (o *GatewayUpdateProducerChef) HasChefOrgs() bool`

HasChefOrgs returns a boolean if a field has been set.

### GetChefServerKey

`func (o *GatewayUpdateProducerChef) GetChefServerKey() string`

GetChefServerKey returns the ChefServerKey field if non-nil, zero value otherwise.

### GetChefServerKeyOk

`func (o *GatewayUpdateProducerChef) GetChefServerKeyOk() (*string, bool)`

GetChefServerKeyOk returns a tuple with the ChefServerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerKey

`func (o *GatewayUpdateProducerChef) SetChefServerKey(v string)`

SetChefServerKey sets ChefServerKey field to given value.

### HasChefServerKey

`func (o *GatewayUpdateProducerChef) HasChefServerKey() bool`

HasChefServerKey returns a boolean if a field has been set.

### GetChefServerUrl

`func (o *GatewayUpdateProducerChef) GetChefServerUrl() string`

GetChefServerUrl returns the ChefServerUrl field if non-nil, zero value otherwise.

### GetChefServerUrlOk

`func (o *GatewayUpdateProducerChef) GetChefServerUrlOk() (*string, bool)`

GetChefServerUrlOk returns a tuple with the ChefServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUrl

`func (o *GatewayUpdateProducerChef) SetChefServerUrl(v string)`

SetChefServerUrl sets ChefServerUrl field to given value.

### HasChefServerUrl

`func (o *GatewayUpdateProducerChef) HasChefServerUrl() bool`

HasChefServerUrl returns a boolean if a field has been set.

### GetChefServerUsername

`func (o *GatewayUpdateProducerChef) GetChefServerUsername() string`

GetChefServerUsername returns the ChefServerUsername field if non-nil, zero value otherwise.

### GetChefServerUsernameOk

`func (o *GatewayUpdateProducerChef) GetChefServerUsernameOk() (*string, bool)`

GetChefServerUsernameOk returns a tuple with the ChefServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChefServerUsername

`func (o *GatewayUpdateProducerChef) SetChefServerUsername(v string)`

SetChefServerUsername sets ChefServerUsername field to given value.

### HasChefServerUsername

`func (o *GatewayUpdateProducerChef) HasChefServerUsername() bool`

HasChefServerUsername returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GatewayUpdateProducerChef) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerChef) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerChef) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerChef) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateProducerChef) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateProducerChef) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateProducerChef) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateProducerChef) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerChef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerChef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerChef) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerChef) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerChef) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerChef) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerChef) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerChef) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerChef) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerChef) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerChef) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetSkipSsl

`func (o *GatewayUpdateProducerChef) GetSkipSsl() bool`

GetSkipSsl returns the SkipSsl field if non-nil, zero value otherwise.

### GetSkipSslOk

`func (o *GatewayUpdateProducerChef) GetSkipSslOk() (*bool, bool)`

GetSkipSslOk returns a tuple with the SkipSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipSsl

`func (o *GatewayUpdateProducerChef) SetSkipSsl(v bool)`

SetSkipSsl sets SkipSsl field to given value.

### HasSkipSsl

`func (o *GatewayUpdateProducerChef) HasSkipSsl() bool`

HasSkipSsl returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerChef) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerChef) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerChef) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerChef) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerChef) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerChef) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerChef) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerChef) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerChef) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerChef) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerChef) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerChef) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerChef) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerChef) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerChef) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerChef) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerChef) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerChef) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerChef) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerChef) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



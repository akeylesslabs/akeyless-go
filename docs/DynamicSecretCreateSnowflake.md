# DynamicSecretCreateSnowflake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account name | [optional] 
**AccountPassword** | Pointer to **string** | Database Password | [optional] 
**AccountUsername** | Pointer to **string** | Database Username | [optional] 
**DbName** | Pointer to **string** | Database name | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**PrivateKey** | Pointer to **string** | RSA Private key (base64 encoded) | [optional] 
**PrivateKeyPassphrase** | Pointer to **string** | The Private key passphrase | [optional] 
**Role** | Pointer to **string** | User role | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "24h"]
**Warehouse** | Pointer to **string** | Warehouse name | [optional] 

## Methods

### NewDynamicSecretCreateSnowflake

`func NewDynamicSecretCreateSnowflake(name string, ) *DynamicSecretCreateSnowflake`

NewDynamicSecretCreateSnowflake instantiates a new DynamicSecretCreateSnowflake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateSnowflakeWithDefaults

`func NewDynamicSecretCreateSnowflakeWithDefaults() *DynamicSecretCreateSnowflake`

NewDynamicSecretCreateSnowflakeWithDefaults instantiates a new DynamicSecretCreateSnowflake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DynamicSecretCreateSnowflake) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DynamicSecretCreateSnowflake) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DynamicSecretCreateSnowflake) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *DynamicSecretCreateSnowflake) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountPassword

`func (o *DynamicSecretCreateSnowflake) GetAccountPassword() string`

GetAccountPassword returns the AccountPassword field if non-nil, zero value otherwise.

### GetAccountPasswordOk

`func (o *DynamicSecretCreateSnowflake) GetAccountPasswordOk() (*string, bool)`

GetAccountPasswordOk returns a tuple with the AccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPassword

`func (o *DynamicSecretCreateSnowflake) SetAccountPassword(v string)`

SetAccountPassword sets AccountPassword field to given value.

### HasAccountPassword

`func (o *DynamicSecretCreateSnowflake) HasAccountPassword() bool`

HasAccountPassword returns a boolean if a field has been set.

### GetAccountUsername

`func (o *DynamicSecretCreateSnowflake) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *DynamicSecretCreateSnowflake) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *DynamicSecretCreateSnowflake) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *DynamicSecretCreateSnowflake) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetDbName

`func (o *DynamicSecretCreateSnowflake) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DynamicSecretCreateSnowflake) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DynamicSecretCreateSnowflake) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DynamicSecretCreateSnowflake) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretCreateSnowflake) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateSnowflake) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateSnowflake) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateSnowflake) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateSnowflake) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateSnowflake) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateSnowflake) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateSnowflake) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateSnowflake) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateSnowflake) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateSnowflake) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateKey

`func (o *DynamicSecretCreateSnowflake) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *DynamicSecretCreateSnowflake) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *DynamicSecretCreateSnowflake) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *DynamicSecretCreateSnowflake) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyPassphrase

`func (o *DynamicSecretCreateSnowflake) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *DynamicSecretCreateSnowflake) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *DynamicSecretCreateSnowflake) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *DynamicSecretCreateSnowflake) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### GetRole

`func (o *DynamicSecretCreateSnowflake) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DynamicSecretCreateSnowflake) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DynamicSecretCreateSnowflake) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DynamicSecretCreateSnowflake) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateSnowflake) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateSnowflake) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateSnowflake) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateSnowflake) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateSnowflake) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateSnowflake) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateSnowflake) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateSnowflake) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateSnowflake) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateSnowflake) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateSnowflake) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateSnowflake) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateSnowflake) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateSnowflake) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateSnowflake) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateSnowflake) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateSnowflake) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateSnowflake) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateSnowflake) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateSnowflake) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarehouse

`func (o *DynamicSecretCreateSnowflake) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *DynamicSecretCreateSnowflake) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *DynamicSecretCreateSnowflake) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *DynamicSecretCreateSnowflake) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



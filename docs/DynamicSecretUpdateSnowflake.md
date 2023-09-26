# DynamicSecretUpdateSnowflake

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
**NewName** | Pointer to **string** | Dynamic secret name | [optional] 
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

### NewDynamicSecretUpdateSnowflake

`func NewDynamicSecretUpdateSnowflake(name string, ) *DynamicSecretUpdateSnowflake`

NewDynamicSecretUpdateSnowflake instantiates a new DynamicSecretUpdateSnowflake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateSnowflakeWithDefaults

`func NewDynamicSecretUpdateSnowflakeWithDefaults() *DynamicSecretUpdateSnowflake`

NewDynamicSecretUpdateSnowflakeWithDefaults instantiates a new DynamicSecretUpdateSnowflake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *DynamicSecretUpdateSnowflake) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *DynamicSecretUpdateSnowflake) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *DynamicSecretUpdateSnowflake) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *DynamicSecretUpdateSnowflake) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountPassword

`func (o *DynamicSecretUpdateSnowflake) GetAccountPassword() string`

GetAccountPassword returns the AccountPassword field if non-nil, zero value otherwise.

### GetAccountPasswordOk

`func (o *DynamicSecretUpdateSnowflake) GetAccountPasswordOk() (*string, bool)`

GetAccountPasswordOk returns a tuple with the AccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPassword

`func (o *DynamicSecretUpdateSnowflake) SetAccountPassword(v string)`

SetAccountPassword sets AccountPassword field to given value.

### HasAccountPassword

`func (o *DynamicSecretUpdateSnowflake) HasAccountPassword() bool`

HasAccountPassword returns a boolean if a field has been set.

### GetAccountUsername

`func (o *DynamicSecretUpdateSnowflake) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *DynamicSecretUpdateSnowflake) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *DynamicSecretUpdateSnowflake) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *DynamicSecretUpdateSnowflake) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetDbName

`func (o *DynamicSecretUpdateSnowflake) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *DynamicSecretUpdateSnowflake) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *DynamicSecretUpdateSnowflake) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *DynamicSecretUpdateSnowflake) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *DynamicSecretUpdateSnowflake) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateSnowflake) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateSnowflake) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateSnowflake) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateSnowflake) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateSnowflake) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateSnowflake) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateSnowflake) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateSnowflake) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateSnowflake) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateSnowflake) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateSnowflake) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateSnowflake) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateSnowflake) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateSnowflake) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPrivateKey

`func (o *DynamicSecretUpdateSnowflake) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *DynamicSecretUpdateSnowflake) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *DynamicSecretUpdateSnowflake) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *DynamicSecretUpdateSnowflake) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyPassphrase

`func (o *DynamicSecretUpdateSnowflake) GetPrivateKeyPassphrase() string`

GetPrivateKeyPassphrase returns the PrivateKeyPassphrase field if non-nil, zero value otherwise.

### GetPrivateKeyPassphraseOk

`func (o *DynamicSecretUpdateSnowflake) GetPrivateKeyPassphraseOk() (*string, bool)`

GetPrivateKeyPassphraseOk returns a tuple with the PrivateKeyPassphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyPassphrase

`func (o *DynamicSecretUpdateSnowflake) SetPrivateKeyPassphrase(v string)`

SetPrivateKeyPassphrase sets PrivateKeyPassphrase field to given value.

### HasPrivateKeyPassphrase

`func (o *DynamicSecretUpdateSnowflake) HasPrivateKeyPassphrase() bool`

HasPrivateKeyPassphrase returns a boolean if a field has been set.

### GetRole

`func (o *DynamicSecretUpdateSnowflake) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *DynamicSecretUpdateSnowflake) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *DynamicSecretUpdateSnowflake) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *DynamicSecretUpdateSnowflake) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateSnowflake) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateSnowflake) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateSnowflake) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateSnowflake) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateSnowflake) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateSnowflake) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateSnowflake) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateSnowflake) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateSnowflake) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateSnowflake) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateSnowflake) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateSnowflake) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateSnowflake) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateSnowflake) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateSnowflake) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateSnowflake) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateSnowflake) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateSnowflake) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateSnowflake) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateSnowflake) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarehouse

`func (o *DynamicSecretUpdateSnowflake) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *DynamicSecretUpdateSnowflake) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *DynamicSecretUpdateSnowflake) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *DynamicSecretUpdateSnowflake) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



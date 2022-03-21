# GatewayUpdateProducerSnowflake

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account name | [optional] 
**AccountPassword** | Pointer to **string** | Database Password | [optional] 
**AccountUsername** | Pointer to **string** | Database Username | [optional] 
**DbName** | Pointer to **string** | Database name | [optional] 
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**Role** | Pointer to **string** | User role | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "24h"]
**Warehouse** | Pointer to **string** | Warehouse name | [optional] 

## Methods

### NewGatewayUpdateProducerSnowflake

`func NewGatewayUpdateProducerSnowflake(name string, ) *GatewayUpdateProducerSnowflake`

NewGatewayUpdateProducerSnowflake instantiates a new GatewayUpdateProducerSnowflake object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerSnowflakeWithDefaults

`func NewGatewayUpdateProducerSnowflakeWithDefaults() *GatewayUpdateProducerSnowflake`

NewGatewayUpdateProducerSnowflakeWithDefaults instantiates a new GatewayUpdateProducerSnowflake object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GatewayUpdateProducerSnowflake) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GatewayUpdateProducerSnowflake) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GatewayUpdateProducerSnowflake) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GatewayUpdateProducerSnowflake) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAccountPassword

`func (o *GatewayUpdateProducerSnowflake) GetAccountPassword() string`

GetAccountPassword returns the AccountPassword field if non-nil, zero value otherwise.

### GetAccountPasswordOk

`func (o *GatewayUpdateProducerSnowflake) GetAccountPasswordOk() (*string, bool)`

GetAccountPasswordOk returns a tuple with the AccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountPassword

`func (o *GatewayUpdateProducerSnowflake) SetAccountPassword(v string)`

SetAccountPassword sets AccountPassword field to given value.

### HasAccountPassword

`func (o *GatewayUpdateProducerSnowflake) HasAccountPassword() bool`

HasAccountPassword returns a boolean if a field has been set.

### GetAccountUsername

`func (o *GatewayUpdateProducerSnowflake) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *GatewayUpdateProducerSnowflake) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *GatewayUpdateProducerSnowflake) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *GatewayUpdateProducerSnowflake) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetDbName

`func (o *GatewayUpdateProducerSnowflake) GetDbName() string`

GetDbName returns the DbName field if non-nil, zero value otherwise.

### GetDbNameOk

`func (o *GatewayUpdateProducerSnowflake) GetDbNameOk() (*string, bool)`

GetDbNameOk returns a tuple with the DbName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbName

`func (o *GatewayUpdateProducerSnowflake) SetDbName(v string)`

SetDbName sets DbName field to given value.

### HasDbName

`func (o *GatewayUpdateProducerSnowflake) HasDbName() bool`

HasDbName returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerSnowflake) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerSnowflake) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerSnowflake) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerSnowflake) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerSnowflake) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerSnowflake) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerSnowflake) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetRole

`func (o *GatewayUpdateProducerSnowflake) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GatewayUpdateProducerSnowflake) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GatewayUpdateProducerSnowflake) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *GatewayUpdateProducerSnowflake) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerSnowflake) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerSnowflake) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerSnowflake) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerSnowflake) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerSnowflake) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerSnowflake) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerSnowflake) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerSnowflake) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerSnowflake) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerSnowflake) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerSnowflake) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerSnowflake) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerSnowflake) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerSnowflake) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerSnowflake) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerSnowflake) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerSnowflake) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerSnowflake) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerSnowflake) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerSnowflake) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetWarehouse

`func (o *GatewayUpdateProducerSnowflake) GetWarehouse() string`

GetWarehouse returns the Warehouse field if non-nil, zero value otherwise.

### GetWarehouseOk

`func (o *GatewayUpdateProducerSnowflake) GetWarehouseOk() (*string, bool)`

GetWarehouseOk returns a tuple with the Warehouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouse

`func (o *GatewayUpdateProducerSnowflake) SetWarehouse(v string)`

SetWarehouse sets Warehouse field to given value.

### HasWarehouse

`func (o *GatewayUpdateProducerSnowflake) HasWarehouse() bool`

HasWarehouse returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



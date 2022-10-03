# GatewayMigratePersonalItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Var1passwordEmail** | Pointer to **string** | 1Password user email to connect to the API | [optional] 
**Var1passwordPassword** | Pointer to **string** | 1Password user password to connect to the API | [optional] 
**Var1passwordSecretKey** | Pointer to **string** | 1Password user secret key to connect to the API | [optional] 
**Var1passwordUrl** | Pointer to **string** | 1Password api container url | [optional] 
**Var1passwordVaults** | Pointer to **[]string** | 1Password list of vault to get the items from | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**ProtectionKey** | Pointer to **string** | The name of a key that used to encrypt the secret value | [optional] 
**TargetLocation** | Pointer to **string** | Target location in Akeyless for migrated secrets | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** | Migration type for now only 1password. | [optional] [default to "1password"]
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayMigratePersonalItems

`func NewGatewayMigratePersonalItems() *GatewayMigratePersonalItems`

NewGatewayMigratePersonalItems instantiates a new GatewayMigratePersonalItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMigratePersonalItemsWithDefaults

`func NewGatewayMigratePersonalItemsWithDefaults() *GatewayMigratePersonalItems`

NewGatewayMigratePersonalItemsWithDefaults instantiates a new GatewayMigratePersonalItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVar1passwordEmail

`func (o *GatewayMigratePersonalItems) GetVar1passwordEmail() string`

GetVar1passwordEmail returns the Var1passwordEmail field if non-nil, zero value otherwise.

### GetVar1passwordEmailOk

`func (o *GatewayMigratePersonalItems) GetVar1passwordEmailOk() (*string, bool)`

GetVar1passwordEmailOk returns a tuple with the Var1passwordEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordEmail

`func (o *GatewayMigratePersonalItems) SetVar1passwordEmail(v string)`

SetVar1passwordEmail sets Var1passwordEmail field to given value.

### HasVar1passwordEmail

`func (o *GatewayMigratePersonalItems) HasVar1passwordEmail() bool`

HasVar1passwordEmail returns a boolean if a field has been set.

### GetVar1passwordPassword

`func (o *GatewayMigratePersonalItems) GetVar1passwordPassword() string`

GetVar1passwordPassword returns the Var1passwordPassword field if non-nil, zero value otherwise.

### GetVar1passwordPasswordOk

`func (o *GatewayMigratePersonalItems) GetVar1passwordPasswordOk() (*string, bool)`

GetVar1passwordPasswordOk returns a tuple with the Var1passwordPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordPassword

`func (o *GatewayMigratePersonalItems) SetVar1passwordPassword(v string)`

SetVar1passwordPassword sets Var1passwordPassword field to given value.

### HasVar1passwordPassword

`func (o *GatewayMigratePersonalItems) HasVar1passwordPassword() bool`

HasVar1passwordPassword returns a boolean if a field has been set.

### GetVar1passwordSecretKey

`func (o *GatewayMigratePersonalItems) GetVar1passwordSecretKey() string`

GetVar1passwordSecretKey returns the Var1passwordSecretKey field if non-nil, zero value otherwise.

### GetVar1passwordSecretKeyOk

`func (o *GatewayMigratePersonalItems) GetVar1passwordSecretKeyOk() (*string, bool)`

GetVar1passwordSecretKeyOk returns a tuple with the Var1passwordSecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordSecretKey

`func (o *GatewayMigratePersonalItems) SetVar1passwordSecretKey(v string)`

SetVar1passwordSecretKey sets Var1passwordSecretKey field to given value.

### HasVar1passwordSecretKey

`func (o *GatewayMigratePersonalItems) HasVar1passwordSecretKey() bool`

HasVar1passwordSecretKey returns a boolean if a field has been set.

### GetVar1passwordUrl

`func (o *GatewayMigratePersonalItems) GetVar1passwordUrl() string`

GetVar1passwordUrl returns the Var1passwordUrl field if non-nil, zero value otherwise.

### GetVar1passwordUrlOk

`func (o *GatewayMigratePersonalItems) GetVar1passwordUrlOk() (*string, bool)`

GetVar1passwordUrlOk returns a tuple with the Var1passwordUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordUrl

`func (o *GatewayMigratePersonalItems) SetVar1passwordUrl(v string)`

SetVar1passwordUrl sets Var1passwordUrl field to given value.

### HasVar1passwordUrl

`func (o *GatewayMigratePersonalItems) HasVar1passwordUrl() bool`

HasVar1passwordUrl returns a boolean if a field has been set.

### GetVar1passwordVaults

`func (o *GatewayMigratePersonalItems) GetVar1passwordVaults() []string`

GetVar1passwordVaults returns the Var1passwordVaults field if non-nil, zero value otherwise.

### GetVar1passwordVaultsOk

`func (o *GatewayMigratePersonalItems) GetVar1passwordVaultsOk() (*[]string, bool)`

GetVar1passwordVaultsOk returns a tuple with the Var1passwordVaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVar1passwordVaults

`func (o *GatewayMigratePersonalItems) SetVar1passwordVaults(v []string)`

SetVar1passwordVaults sets Var1passwordVaults field to given value.

### HasVar1passwordVaults

`func (o *GatewayMigratePersonalItems) HasVar1passwordVaults() bool`

HasVar1passwordVaults returns a boolean if a field has been set.

### GetJson

`func (o *GatewayMigratePersonalItems) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayMigratePersonalItems) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayMigratePersonalItems) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayMigratePersonalItems) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetProtectionKey

`func (o *GatewayMigratePersonalItems) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *GatewayMigratePersonalItems) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *GatewayMigratePersonalItems) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *GatewayMigratePersonalItems) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTargetLocation

`func (o *GatewayMigratePersonalItems) GetTargetLocation() string`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *GatewayMigratePersonalItems) GetTargetLocationOk() (*string, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *GatewayMigratePersonalItems) SetTargetLocation(v string)`

SetTargetLocation sets TargetLocation field to given value.

### HasTargetLocation

`func (o *GatewayMigratePersonalItems) HasTargetLocation() bool`

HasTargetLocation returns a boolean if a field has been set.

### GetToken

`func (o *GatewayMigratePersonalItems) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayMigratePersonalItems) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayMigratePersonalItems) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayMigratePersonalItems) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *GatewayMigratePersonalItems) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayMigratePersonalItems) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayMigratePersonalItems) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayMigratePersonalItems) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayMigratePersonalItems) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayMigratePersonalItems) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayMigratePersonalItems) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayMigratePersonalItems) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



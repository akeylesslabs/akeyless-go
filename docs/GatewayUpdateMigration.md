# GatewayUpdateMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AwsKey** | Pointer to **string** |  | [optional] 
**AwsKeyId** | Pointer to **string** |  | [optional] 
**AwsRegion** | Pointer to **string** |  | [optional] 
**AzureClientId** | Pointer to **string** |  | [optional] 
**AzureKvName** | Pointer to **string** |  | [optional] 
**AzureSecret** | Pointer to **string** |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**HashiJson** | Pointer to **string** |  | [optional] 
**HashiNs** | Pointer to **[]string** |  | [optional] 
**HashiToken** | Pointer to **string** |  | [optional] 
**HashiUrl** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | **string** | Migration name | 
**ProtectionKey** | Pointer to **string** |  | [optional] 
**TargetLocation** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateMigration

`func NewGatewayUpdateMigration(name string, ) *GatewayUpdateMigration`

NewGatewayUpdateMigration instantiates a new GatewayUpdateMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateMigrationWithDefaults

`func NewGatewayUpdateMigrationWithDefaults() *GatewayUpdateMigration`

NewGatewayUpdateMigrationWithDefaults instantiates a new GatewayUpdateMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *GatewayUpdateMigration) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *GatewayUpdateMigration) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *GatewayUpdateMigration) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.

### HasAwsKey

`func (o *GatewayUpdateMigration) HasAwsKey() bool`

HasAwsKey returns a boolean if a field has been set.

### GetAwsKeyId

`func (o *GatewayUpdateMigration) GetAwsKeyId() string`

GetAwsKeyId returns the AwsKeyId field if non-nil, zero value otherwise.

### GetAwsKeyIdOk

`func (o *GatewayUpdateMigration) GetAwsKeyIdOk() (*string, bool)`

GetAwsKeyIdOk returns a tuple with the AwsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKeyId

`func (o *GatewayUpdateMigration) SetAwsKeyId(v string)`

SetAwsKeyId sets AwsKeyId field to given value.

### HasAwsKeyId

`func (o *GatewayUpdateMigration) HasAwsKeyId() bool`

HasAwsKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *GatewayUpdateMigration) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *GatewayUpdateMigration) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *GatewayUpdateMigration) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *GatewayUpdateMigration) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureClientId

`func (o *GatewayUpdateMigration) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *GatewayUpdateMigration) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *GatewayUpdateMigration) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *GatewayUpdateMigration) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureKvName

`func (o *GatewayUpdateMigration) GetAzureKvName() string`

GetAzureKvName returns the AzureKvName field if non-nil, zero value otherwise.

### GetAzureKvNameOk

`func (o *GatewayUpdateMigration) GetAzureKvNameOk() (*string, bool)`

GetAzureKvNameOk returns a tuple with the AzureKvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvName

`func (o *GatewayUpdateMigration) SetAzureKvName(v string)`

SetAzureKvName sets AzureKvName field to given value.

### HasAzureKvName

`func (o *GatewayUpdateMigration) HasAzureKvName() bool`

HasAzureKvName returns a boolean if a field has been set.

### GetAzureSecret

`func (o *GatewayUpdateMigration) GetAzureSecret() string`

GetAzureSecret returns the AzureSecret field if non-nil, zero value otherwise.

### GetAzureSecretOk

`func (o *GatewayUpdateMigration) GetAzureSecretOk() (*string, bool)`

GetAzureSecretOk returns a tuple with the AzureSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSecret

`func (o *GatewayUpdateMigration) SetAzureSecret(v string)`

SetAzureSecret sets AzureSecret field to given value.

### HasAzureSecret

`func (o *GatewayUpdateMigration) HasAzureSecret() bool`

HasAzureSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *GatewayUpdateMigration) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *GatewayUpdateMigration) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *GatewayUpdateMigration) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *GatewayUpdateMigration) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetHashiJson

`func (o *GatewayUpdateMigration) GetHashiJson() string`

GetHashiJson returns the HashiJson field if non-nil, zero value otherwise.

### GetHashiJsonOk

`func (o *GatewayUpdateMigration) GetHashiJsonOk() (*string, bool)`

GetHashiJsonOk returns a tuple with the HashiJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiJson

`func (o *GatewayUpdateMigration) SetHashiJson(v string)`

SetHashiJson sets HashiJson field to given value.

### HasHashiJson

`func (o *GatewayUpdateMigration) HasHashiJson() bool`

HasHashiJson returns a boolean if a field has been set.

### GetHashiNs

`func (o *GatewayUpdateMigration) GetHashiNs() []string`

GetHashiNs returns the HashiNs field if non-nil, zero value otherwise.

### GetHashiNsOk

`func (o *GatewayUpdateMigration) GetHashiNsOk() (*[]string, bool)`

GetHashiNsOk returns a tuple with the HashiNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiNs

`func (o *GatewayUpdateMigration) SetHashiNs(v []string)`

SetHashiNs sets HashiNs field to given value.

### HasHashiNs

`func (o *GatewayUpdateMigration) HasHashiNs() bool`

HasHashiNs returns a boolean if a field has been set.

### GetHashiToken

`func (o *GatewayUpdateMigration) GetHashiToken() string`

GetHashiToken returns the HashiToken field if non-nil, zero value otherwise.

### GetHashiTokenOk

`func (o *GatewayUpdateMigration) GetHashiTokenOk() (*string, bool)`

GetHashiTokenOk returns a tuple with the HashiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiToken

`func (o *GatewayUpdateMigration) SetHashiToken(v string)`

SetHashiToken sets HashiToken field to given value.

### HasHashiToken

`func (o *GatewayUpdateMigration) HasHashiToken() bool`

HasHashiToken returns a boolean if a field has been set.

### GetHashiUrl

`func (o *GatewayUpdateMigration) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *GatewayUpdateMigration) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *GatewayUpdateMigration) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *GatewayUpdateMigration) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetId

`func (o *GatewayUpdateMigration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayUpdateMigration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayUpdateMigration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayUpdateMigration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateMigration) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *GatewayUpdateMigration) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *GatewayUpdateMigration) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *GatewayUpdateMigration) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *GatewayUpdateMigration) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTargetLocation

`func (o *GatewayUpdateMigration) GetTargetLocation() string`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *GatewayUpdateMigration) GetTargetLocationOk() (*string, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *GatewayUpdateMigration) SetTargetLocation(v string)`

SetTargetLocation sets TargetLocation field to given value.

### HasTargetLocation

`func (o *GatewayUpdateMigration) HasTargetLocation() bool`

HasTargetLocation returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateMigration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateMigration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateMigration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateMigration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *GatewayUpdateMigration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayUpdateMigration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayUpdateMigration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayUpdateMigration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateMigration) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateMigration) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateMigration) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateMigration) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



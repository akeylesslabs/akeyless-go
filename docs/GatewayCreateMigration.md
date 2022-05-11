# GatewayCreateMigration

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
**Name** | **string** | Migration name | 
**ProtectionKey** | Pointer to **string** |  | [optional] 
**TargetLocation** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayCreateMigration

`func NewGatewayCreateMigration(name string, ) *GatewayCreateMigration`

NewGatewayCreateMigration instantiates a new GatewayCreateMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateMigrationWithDefaults

`func NewGatewayCreateMigrationWithDefaults() *GatewayCreateMigration`

NewGatewayCreateMigrationWithDefaults instantiates a new GatewayCreateMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAwsKey

`func (o *GatewayCreateMigration) GetAwsKey() string`

GetAwsKey returns the AwsKey field if non-nil, zero value otherwise.

### GetAwsKeyOk

`func (o *GatewayCreateMigration) GetAwsKeyOk() (*string, bool)`

GetAwsKeyOk returns a tuple with the AwsKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKey

`func (o *GatewayCreateMigration) SetAwsKey(v string)`

SetAwsKey sets AwsKey field to given value.

### HasAwsKey

`func (o *GatewayCreateMigration) HasAwsKey() bool`

HasAwsKey returns a boolean if a field has been set.

### GetAwsKeyId

`func (o *GatewayCreateMigration) GetAwsKeyId() string`

GetAwsKeyId returns the AwsKeyId field if non-nil, zero value otherwise.

### GetAwsKeyIdOk

`func (o *GatewayCreateMigration) GetAwsKeyIdOk() (*string, bool)`

GetAwsKeyIdOk returns a tuple with the AwsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsKeyId

`func (o *GatewayCreateMigration) SetAwsKeyId(v string)`

SetAwsKeyId sets AwsKeyId field to given value.

### HasAwsKeyId

`func (o *GatewayCreateMigration) HasAwsKeyId() bool`

HasAwsKeyId returns a boolean if a field has been set.

### GetAwsRegion

`func (o *GatewayCreateMigration) GetAwsRegion() string`

GetAwsRegion returns the AwsRegion field if non-nil, zero value otherwise.

### GetAwsRegionOk

`func (o *GatewayCreateMigration) GetAwsRegionOk() (*string, bool)`

GetAwsRegionOk returns a tuple with the AwsRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsRegion

`func (o *GatewayCreateMigration) SetAwsRegion(v string)`

SetAwsRegion sets AwsRegion field to given value.

### HasAwsRegion

`func (o *GatewayCreateMigration) HasAwsRegion() bool`

HasAwsRegion returns a boolean if a field has been set.

### GetAzureClientId

`func (o *GatewayCreateMigration) GetAzureClientId() string`

GetAzureClientId returns the AzureClientId field if non-nil, zero value otherwise.

### GetAzureClientIdOk

`func (o *GatewayCreateMigration) GetAzureClientIdOk() (*string, bool)`

GetAzureClientIdOk returns a tuple with the AzureClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureClientId

`func (o *GatewayCreateMigration) SetAzureClientId(v string)`

SetAzureClientId sets AzureClientId field to given value.

### HasAzureClientId

`func (o *GatewayCreateMigration) HasAzureClientId() bool`

HasAzureClientId returns a boolean if a field has been set.

### GetAzureKvName

`func (o *GatewayCreateMigration) GetAzureKvName() string`

GetAzureKvName returns the AzureKvName field if non-nil, zero value otherwise.

### GetAzureKvNameOk

`func (o *GatewayCreateMigration) GetAzureKvNameOk() (*string, bool)`

GetAzureKvNameOk returns a tuple with the AzureKvName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvName

`func (o *GatewayCreateMigration) SetAzureKvName(v string)`

SetAzureKvName sets AzureKvName field to given value.

### HasAzureKvName

`func (o *GatewayCreateMigration) HasAzureKvName() bool`

HasAzureKvName returns a boolean if a field has been set.

### GetAzureSecret

`func (o *GatewayCreateMigration) GetAzureSecret() string`

GetAzureSecret returns the AzureSecret field if non-nil, zero value otherwise.

### GetAzureSecretOk

`func (o *GatewayCreateMigration) GetAzureSecretOk() (*string, bool)`

GetAzureSecretOk returns a tuple with the AzureSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureSecret

`func (o *GatewayCreateMigration) SetAzureSecret(v string)`

SetAzureSecret sets AzureSecret field to given value.

### HasAzureSecret

`func (o *GatewayCreateMigration) HasAzureSecret() bool`

HasAzureSecret returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *GatewayCreateMigration) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *GatewayCreateMigration) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *GatewayCreateMigration) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *GatewayCreateMigration) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetHashiJson

`func (o *GatewayCreateMigration) GetHashiJson() string`

GetHashiJson returns the HashiJson field if non-nil, zero value otherwise.

### GetHashiJsonOk

`func (o *GatewayCreateMigration) GetHashiJsonOk() (*string, bool)`

GetHashiJsonOk returns a tuple with the HashiJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiJson

`func (o *GatewayCreateMigration) SetHashiJson(v string)`

SetHashiJson sets HashiJson field to given value.

### HasHashiJson

`func (o *GatewayCreateMigration) HasHashiJson() bool`

HasHashiJson returns a boolean if a field has been set.

### GetHashiNs

`func (o *GatewayCreateMigration) GetHashiNs() []string`

GetHashiNs returns the HashiNs field if non-nil, zero value otherwise.

### GetHashiNsOk

`func (o *GatewayCreateMigration) GetHashiNsOk() (*[]string, bool)`

GetHashiNsOk returns a tuple with the HashiNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiNs

`func (o *GatewayCreateMigration) SetHashiNs(v []string)`

SetHashiNs sets HashiNs field to given value.

### HasHashiNs

`func (o *GatewayCreateMigration) HasHashiNs() bool`

HasHashiNs returns a boolean if a field has been set.

### GetHashiToken

`func (o *GatewayCreateMigration) GetHashiToken() string`

GetHashiToken returns the HashiToken field if non-nil, zero value otherwise.

### GetHashiTokenOk

`func (o *GatewayCreateMigration) GetHashiTokenOk() (*string, bool)`

GetHashiTokenOk returns a tuple with the HashiToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiToken

`func (o *GatewayCreateMigration) SetHashiToken(v string)`

SetHashiToken sets HashiToken field to given value.

### HasHashiToken

`func (o *GatewayCreateMigration) HasHashiToken() bool`

HasHashiToken returns a boolean if a field has been set.

### GetHashiUrl

`func (o *GatewayCreateMigration) GetHashiUrl() string`

GetHashiUrl returns the HashiUrl field if non-nil, zero value otherwise.

### GetHashiUrlOk

`func (o *GatewayCreateMigration) GetHashiUrlOk() (*string, bool)`

GetHashiUrlOk returns a tuple with the HashiUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiUrl

`func (o *GatewayCreateMigration) SetHashiUrl(v string)`

SetHashiUrl sets HashiUrl field to given value.

### HasHashiUrl

`func (o *GatewayCreateMigration) HasHashiUrl() bool`

HasHashiUrl returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateMigration) SetName(v string)`

SetName sets Name field to given value.


### GetProtectionKey

`func (o *GatewayCreateMigration) GetProtectionKey() string`

GetProtectionKey returns the ProtectionKey field if non-nil, zero value otherwise.

### GetProtectionKeyOk

`func (o *GatewayCreateMigration) GetProtectionKeyOk() (*string, bool)`

GetProtectionKeyOk returns a tuple with the ProtectionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKey

`func (o *GatewayCreateMigration) SetProtectionKey(v string)`

SetProtectionKey sets ProtectionKey field to given value.

### HasProtectionKey

`func (o *GatewayCreateMigration) HasProtectionKey() bool`

HasProtectionKey returns a boolean if a field has been set.

### GetTargetLocation

`func (o *GatewayCreateMigration) GetTargetLocation() string`

GetTargetLocation returns the TargetLocation field if non-nil, zero value otherwise.

### GetTargetLocationOk

`func (o *GatewayCreateMigration) GetTargetLocationOk() (*string, bool)`

GetTargetLocationOk returns a tuple with the TargetLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetLocation

`func (o *GatewayCreateMigration) SetTargetLocation(v string)`

SetTargetLocation sets TargetLocation field to given value.

### HasTargetLocation

`func (o *GatewayCreateMigration) HasTargetLocation() bool`

HasTargetLocation returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateMigration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateMigration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateMigration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateMigration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetType

`func (o *GatewayCreateMigration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayCreateMigration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayCreateMigration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayCreateMigration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateMigration) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateMigration) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateMigration) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateMigration) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



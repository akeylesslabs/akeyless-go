# MigrationsConfigPart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveDirectoryMigrations** | Pointer to [**[]ActiveDirectoryMigration**](ActiveDirectoryMigration.md) |  | [optional] 
**AwsSecretsMigrations** | Pointer to [**[]AWSSecretsMigration**](AWSSecretsMigration.md) |  | [optional] 
**AzureKvMigrations** | Pointer to [**[]AzureKeyVaultMigration**](AzureKeyVaultMigration.md) |  | [optional] 
**GcpSecretsMigrations** | Pointer to [**[]GCPSecretsMigration**](GCPSecretsMigration.md) |  | [optional] 
**HashiMigrations** | Pointer to [**[]HashiMigration**](HashiMigration.md) |  | [optional] 
**K8sMigrations** | Pointer to [**[]K8SMigration**](K8SMigration.md) |  | [optional] 
**MockMigrations** | Pointer to [**[]MockMigration**](MockMigration.md) |  | [optional] 
**OnePasswordMigrations** | Pointer to [**[]OnePasswordMigration**](OnePasswordMigration.md) |  | [optional] 
**ServerInventoryMigrations** | Pointer to [**[]ServerInventoryMigration**](ServerInventoryMigration.md) |  | [optional] 

## Methods

### NewMigrationsConfigPart

`func NewMigrationsConfigPart() *MigrationsConfigPart`

NewMigrationsConfigPart instantiates a new MigrationsConfigPart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMigrationsConfigPartWithDefaults

`func NewMigrationsConfigPartWithDefaults() *MigrationsConfigPart`

NewMigrationsConfigPartWithDefaults instantiates a new MigrationsConfigPart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveDirectoryMigrations

`func (o *MigrationsConfigPart) GetActiveDirectoryMigrations() []ActiveDirectoryMigration`

GetActiveDirectoryMigrations returns the ActiveDirectoryMigrations field if non-nil, zero value otherwise.

### GetActiveDirectoryMigrationsOk

`func (o *MigrationsConfigPart) GetActiveDirectoryMigrationsOk() (*[]ActiveDirectoryMigration, bool)`

GetActiveDirectoryMigrationsOk returns a tuple with the ActiveDirectoryMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveDirectoryMigrations

`func (o *MigrationsConfigPart) SetActiveDirectoryMigrations(v []ActiveDirectoryMigration)`

SetActiveDirectoryMigrations sets ActiveDirectoryMigrations field to given value.

### HasActiveDirectoryMigrations

`func (o *MigrationsConfigPart) HasActiveDirectoryMigrations() bool`

HasActiveDirectoryMigrations returns a boolean if a field has been set.

### GetAwsSecretsMigrations

`func (o *MigrationsConfigPart) GetAwsSecretsMigrations() []AWSSecretsMigration`

GetAwsSecretsMigrations returns the AwsSecretsMigrations field if non-nil, zero value otherwise.

### GetAwsSecretsMigrationsOk

`func (o *MigrationsConfigPart) GetAwsSecretsMigrationsOk() (*[]AWSSecretsMigration, bool)`

GetAwsSecretsMigrationsOk returns a tuple with the AwsSecretsMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAwsSecretsMigrations

`func (o *MigrationsConfigPart) SetAwsSecretsMigrations(v []AWSSecretsMigration)`

SetAwsSecretsMigrations sets AwsSecretsMigrations field to given value.

### HasAwsSecretsMigrations

`func (o *MigrationsConfigPart) HasAwsSecretsMigrations() bool`

HasAwsSecretsMigrations returns a boolean if a field has been set.

### GetAzureKvMigrations

`func (o *MigrationsConfigPart) GetAzureKvMigrations() []AzureKeyVaultMigration`

GetAzureKvMigrations returns the AzureKvMigrations field if non-nil, zero value otherwise.

### GetAzureKvMigrationsOk

`func (o *MigrationsConfigPart) GetAzureKvMigrationsOk() (*[]AzureKeyVaultMigration, bool)`

GetAzureKvMigrationsOk returns a tuple with the AzureKvMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureKvMigrations

`func (o *MigrationsConfigPart) SetAzureKvMigrations(v []AzureKeyVaultMigration)`

SetAzureKvMigrations sets AzureKvMigrations field to given value.

### HasAzureKvMigrations

`func (o *MigrationsConfigPart) HasAzureKvMigrations() bool`

HasAzureKvMigrations returns a boolean if a field has been set.

### GetGcpSecretsMigrations

`func (o *MigrationsConfigPart) GetGcpSecretsMigrations() []GCPSecretsMigration`

GetGcpSecretsMigrations returns the GcpSecretsMigrations field if non-nil, zero value otherwise.

### GetGcpSecretsMigrationsOk

`func (o *MigrationsConfigPart) GetGcpSecretsMigrationsOk() (*[]GCPSecretsMigration, bool)`

GetGcpSecretsMigrationsOk returns a tuple with the GcpSecretsMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSecretsMigrations

`func (o *MigrationsConfigPart) SetGcpSecretsMigrations(v []GCPSecretsMigration)`

SetGcpSecretsMigrations sets GcpSecretsMigrations field to given value.

### HasGcpSecretsMigrations

`func (o *MigrationsConfigPart) HasGcpSecretsMigrations() bool`

HasGcpSecretsMigrations returns a boolean if a field has been set.

### GetHashiMigrations

`func (o *MigrationsConfigPart) GetHashiMigrations() []HashiMigration`

GetHashiMigrations returns the HashiMigrations field if non-nil, zero value otherwise.

### GetHashiMigrationsOk

`func (o *MigrationsConfigPart) GetHashiMigrationsOk() (*[]HashiMigration, bool)`

GetHashiMigrationsOk returns a tuple with the HashiMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashiMigrations

`func (o *MigrationsConfigPart) SetHashiMigrations(v []HashiMigration)`

SetHashiMigrations sets HashiMigrations field to given value.

### HasHashiMigrations

`func (o *MigrationsConfigPart) HasHashiMigrations() bool`

HasHashiMigrations returns a boolean if a field has been set.

### GetK8sMigrations

`func (o *MigrationsConfigPart) GetK8sMigrations() []K8SMigration`

GetK8sMigrations returns the K8sMigrations field if non-nil, zero value otherwise.

### GetK8sMigrationsOk

`func (o *MigrationsConfigPart) GetK8sMigrationsOk() (*[]K8SMigration, bool)`

GetK8sMigrationsOk returns a tuple with the K8sMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK8sMigrations

`func (o *MigrationsConfigPart) SetK8sMigrations(v []K8SMigration)`

SetK8sMigrations sets K8sMigrations field to given value.

### HasK8sMigrations

`func (o *MigrationsConfigPart) HasK8sMigrations() bool`

HasK8sMigrations returns a boolean if a field has been set.

### GetMockMigrations

`func (o *MigrationsConfigPart) GetMockMigrations() []MockMigration`

GetMockMigrations returns the MockMigrations field if non-nil, zero value otherwise.

### GetMockMigrationsOk

`func (o *MigrationsConfigPart) GetMockMigrationsOk() (*[]MockMigration, bool)`

GetMockMigrationsOk returns a tuple with the MockMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMockMigrations

`func (o *MigrationsConfigPart) SetMockMigrations(v []MockMigration)`

SetMockMigrations sets MockMigrations field to given value.

### HasMockMigrations

`func (o *MigrationsConfigPart) HasMockMigrations() bool`

HasMockMigrations returns a boolean if a field has been set.

### GetOnePasswordMigrations

`func (o *MigrationsConfigPart) GetOnePasswordMigrations() []OnePasswordMigration`

GetOnePasswordMigrations returns the OnePasswordMigrations field if non-nil, zero value otherwise.

### GetOnePasswordMigrationsOk

`func (o *MigrationsConfigPart) GetOnePasswordMigrationsOk() (*[]OnePasswordMigration, bool)`

GetOnePasswordMigrationsOk returns a tuple with the OnePasswordMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnePasswordMigrations

`func (o *MigrationsConfigPart) SetOnePasswordMigrations(v []OnePasswordMigration)`

SetOnePasswordMigrations sets OnePasswordMigrations field to given value.

### HasOnePasswordMigrations

`func (o *MigrationsConfigPart) HasOnePasswordMigrations() bool`

HasOnePasswordMigrations returns a boolean if a field has been set.

### GetServerInventoryMigrations

`func (o *MigrationsConfigPart) GetServerInventoryMigrations() []ServerInventoryMigration`

GetServerInventoryMigrations returns the ServerInventoryMigrations field if non-nil, zero value otherwise.

### GetServerInventoryMigrationsOk

`func (o *MigrationsConfigPart) GetServerInventoryMigrationsOk() (*[]ServerInventoryMigration, bool)`

GetServerInventoryMigrationsOk returns a tuple with the ServerInventoryMigrations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerInventoryMigrations

`func (o *MigrationsConfigPart) SetServerInventoryMigrations(v []ServerInventoryMigration)`

SetServerInventoryMigrations sets ServerInventoryMigrations field to given value.

### HasServerInventoryMigrations

`func (o *MigrationsConfigPart) HasServerInventoryMigrations() bool`

HasServerInventoryMigrations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# AzureKeyVaultMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**MigrationGeneral**](MigrationGeneral.md) |  | [optional] 
**Payload** | Pointer to [**AzurePayload**](AzurePayload.md) |  | [optional] 

## Methods

### NewAzureKeyVaultMigration

`func NewAzureKeyVaultMigration() *AzureKeyVaultMigration`

NewAzureKeyVaultMigration instantiates a new AzureKeyVaultMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureKeyVaultMigrationWithDefaults

`func NewAzureKeyVaultMigrationWithDefaults() *AzureKeyVaultMigration`

NewAzureKeyVaultMigrationWithDefaults instantiates a new AzureKeyVaultMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *AzureKeyVaultMigration) GetGeneral() MigrationGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *AzureKeyVaultMigration) GetGeneralOk() (*MigrationGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *AzureKeyVaultMigration) SetGeneral(v MigrationGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *AzureKeyVaultMigration) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetPayload

`func (o *AzureKeyVaultMigration) GetPayload() AzurePayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AzureKeyVaultMigration) GetPayloadOk() (*AzurePayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AzureKeyVaultMigration) SetPayload(v AzurePayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AzureKeyVaultMigration) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



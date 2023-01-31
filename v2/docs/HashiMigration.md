# HashiMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**MigrationGeneral**](MigrationGeneral.md) |  | [optional] 
**Payload** | Pointer to [**HashiPayload**](HashiPayload.md) |  | [optional] 

## Methods

### NewHashiMigration

`func NewHashiMigration() *HashiMigration`

NewHashiMigration instantiates a new HashiMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHashiMigrationWithDefaults

`func NewHashiMigrationWithDefaults() *HashiMigration`

NewHashiMigrationWithDefaults instantiates a new HashiMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *HashiMigration) GetGeneral() MigrationGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *HashiMigration) GetGeneralOk() (*MigrationGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *HashiMigration) SetGeneral(v MigrationGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *HashiMigration) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetPayload

`func (o *HashiMigration) GetPayload() HashiPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *HashiMigration) GetPayloadOk() (*HashiPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *HashiMigration) SetPayload(v HashiPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *HashiMigration) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



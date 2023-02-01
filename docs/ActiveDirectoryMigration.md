# ActiveDirectoryMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**MigrationGeneral**](MigrationGeneral.md) |  | [optional] 
**Payload** | Pointer to [**ActiveDirectoryPayload**](ActiveDirectoryPayload.md) |  | [optional] 

## Methods

### NewActiveDirectoryMigration

`func NewActiveDirectoryMigration() *ActiveDirectoryMigration`

NewActiveDirectoryMigration instantiates a new ActiveDirectoryMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActiveDirectoryMigrationWithDefaults

`func NewActiveDirectoryMigrationWithDefaults() *ActiveDirectoryMigration`

NewActiveDirectoryMigrationWithDefaults instantiates a new ActiveDirectoryMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *ActiveDirectoryMigration) GetGeneral() MigrationGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *ActiveDirectoryMigration) GetGeneralOk() (*MigrationGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *ActiveDirectoryMigration) SetGeneral(v MigrationGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *ActiveDirectoryMigration) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetPayload

`func (o *ActiveDirectoryMigration) GetPayload() ActiveDirectoryPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ActiveDirectoryMigration) GetPayloadOk() (*ActiveDirectoryPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ActiveDirectoryMigration) SetPayload(v ActiveDirectoryPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ActiveDirectoryMigration) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



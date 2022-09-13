# OnePasswordMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**MigrationGeneral**](MigrationGeneral.md) |  | [optional] 
**Payload** | Pointer to [**OnePasswordPayload**](OnePasswordPayload.md) |  | [optional] 

## Methods

### NewOnePasswordMigration

`func NewOnePasswordMigration() *OnePasswordMigration`

NewOnePasswordMigration instantiates a new OnePasswordMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnePasswordMigrationWithDefaults

`func NewOnePasswordMigrationWithDefaults() *OnePasswordMigration`

NewOnePasswordMigrationWithDefaults instantiates a new OnePasswordMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *OnePasswordMigration) GetGeneral() MigrationGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *OnePasswordMigration) GetGeneralOk() (*MigrationGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *OnePasswordMigration) SetGeneral(v MigrationGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *OnePasswordMigration) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetPayload

`func (o *OnePasswordMigration) GetPayload() OnePasswordPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *OnePasswordMigration) GetPayloadOk() (*OnePasswordPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *OnePasswordMigration) SetPayload(v OnePasswordPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *OnePasswordMigration) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



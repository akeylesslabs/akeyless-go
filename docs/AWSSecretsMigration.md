# AWSSecretsMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**General** | Pointer to [**MigrationGeneral**](MigrationGeneral.md) |  | [optional] 
**Payload** | Pointer to [**AWSPayload**](AWSPayload.md) |  | [optional] 

## Methods

### NewAWSSecretsMigration

`func NewAWSSecretsMigration() *AWSSecretsMigration`

NewAWSSecretsMigration instantiates a new AWSSecretsMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAWSSecretsMigrationWithDefaults

`func NewAWSSecretsMigrationWithDefaults() *AWSSecretsMigration`

NewAWSSecretsMigrationWithDefaults instantiates a new AWSSecretsMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneral

`func (o *AWSSecretsMigration) GetGeneral() MigrationGeneral`

GetGeneral returns the General field if non-nil, zero value otherwise.

### GetGeneralOk

`func (o *AWSSecretsMigration) GetGeneralOk() (*MigrationGeneral, bool)`

GetGeneralOk returns a tuple with the General field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneral

`func (o *AWSSecretsMigration) SetGeneral(v MigrationGeneral)`

SetGeneral sets General field to given value.

### HasGeneral

`func (o *AWSSecretsMigration) HasGeneral() bool`

HasGeneral returns a boolean if a field has been set.

### GetPayload

`func (o *AWSSecretsMigration) GetPayload() AWSPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AWSSecretsMigration) GetPayloadOk() (*AWSPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AWSSecretsMigration) SetPayload(v AWSPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AWSSecretsMigration) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



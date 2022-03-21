# GatewaySyncMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Migration name | 
**StartSync** | Pointer to **bool** |  | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewaySyncMigration

`func NewGatewaySyncMigration(name string, ) *GatewaySyncMigration`

NewGatewaySyncMigration instantiates a new GatewaySyncMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewaySyncMigrationWithDefaults

`func NewGatewaySyncMigrationWithDefaults() *GatewaySyncMigration`

NewGatewaySyncMigrationWithDefaults instantiates a new GatewaySyncMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GatewaySyncMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewaySyncMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewaySyncMigration) SetName(v string)`

SetName sets Name field to given value.


### GetStartSync

`func (o *GatewaySyncMigration) GetStartSync() bool`

GetStartSync returns the StartSync field if non-nil, zero value otherwise.

### GetStartSyncOk

`func (o *GatewaySyncMigration) GetStartSyncOk() (*bool, bool)`

GetStartSyncOk returns a tuple with the StartSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartSync

`func (o *GatewaySyncMigration) SetStartSync(v bool)`

SetStartSync sets StartSync field to given value.

### HasStartSync

`func (o *GatewaySyncMigration) HasStartSync() bool`

HasStartSync returns a boolean if a field has been set.

### GetToken

`func (o *GatewaySyncMigration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewaySyncMigration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewaySyncMigration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewaySyncMigration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewaySyncMigration) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewaySyncMigration) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewaySyncMigration) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewaySyncMigration) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



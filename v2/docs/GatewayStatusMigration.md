# GatewayStatusMigration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Migration ID | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Name** | Pointer to **string** | Migration name to display | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayStatusMigration

`func NewGatewayStatusMigration() *GatewayStatusMigration`

NewGatewayStatusMigration instantiates a new GatewayStatusMigration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayStatusMigrationWithDefaults

`func NewGatewayStatusMigrationWithDefaults() *GatewayStatusMigration`

NewGatewayStatusMigrationWithDefaults instantiates a new GatewayStatusMigration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayStatusMigration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayStatusMigration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayStatusMigration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayStatusMigration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJson

`func (o *GatewayStatusMigration) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayStatusMigration) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayStatusMigration) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayStatusMigration) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayStatusMigration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayStatusMigration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayStatusMigration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayStatusMigration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayStatusMigration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayStatusMigration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayStatusMigration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayStatusMigration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayStatusMigration) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayStatusMigration) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayStatusMigration) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayStatusMigration) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



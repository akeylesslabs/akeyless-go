# GatewayUpdateProducerCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateSyncUrl** | **string** | URL of an endpoint that implements /sync/create method, for example https://webhook.example.com/sync/create | 
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Payload** | Pointer to **string** | Secret payload to be sent with each create/revoke webhook request | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RevokeSyncUrl** | **string** | URL of an endpoint that implements /sync/revoke method, for example https://webhook.example.com/sync/revoke | 
**RotateSyncUrl** | Pointer to **string** | URL of an endpoint that implements /sync/rotate method, for example https://webhook.example.com/sync/rotate | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TimeoutSec** | Pointer to **int64** | Maximum allowed time in seconds for the webhook to return the results | [optional] [default to 60]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayUpdateProducerCustom

`func NewGatewayUpdateProducerCustom(createSyncUrl string, name string, revokeSyncUrl string, ) *GatewayUpdateProducerCustom`

NewGatewayUpdateProducerCustom instantiates a new GatewayUpdateProducerCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerCustomWithDefaults

`func NewGatewayUpdateProducerCustomWithDefaults() *GatewayUpdateProducerCustom`

NewGatewayUpdateProducerCustomWithDefaults instantiates a new GatewayUpdateProducerCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateSyncUrl

`func (o *GatewayUpdateProducerCustom) GetCreateSyncUrl() string`

GetCreateSyncUrl returns the CreateSyncUrl field if non-nil, zero value otherwise.

### GetCreateSyncUrlOk

`func (o *GatewayUpdateProducerCustom) GetCreateSyncUrlOk() (*string, bool)`

GetCreateSyncUrlOk returns a tuple with the CreateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSyncUrl

`func (o *GatewayUpdateProducerCustom) SetCreateSyncUrl(v string)`

SetCreateSyncUrl sets CreateSyncUrl field to given value.


### GetName

`func (o *GatewayUpdateProducerCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerCustom) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerCustom) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerCustom) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerCustom) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerCustom) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPassword

`func (o *GatewayUpdateProducerCustom) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayUpdateProducerCustom) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayUpdateProducerCustom) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayUpdateProducerCustom) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPayload

`func (o *GatewayUpdateProducerCustom) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GatewayUpdateProducerCustom) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GatewayUpdateProducerCustom) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *GatewayUpdateProducerCustom) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerCustom) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerCustom) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerCustom) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerCustom) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRevokeSyncUrl

`func (o *GatewayUpdateProducerCustom) GetRevokeSyncUrl() string`

GetRevokeSyncUrl returns the RevokeSyncUrl field if non-nil, zero value otherwise.

### GetRevokeSyncUrlOk

`func (o *GatewayUpdateProducerCustom) GetRevokeSyncUrlOk() (*string, bool)`

GetRevokeSyncUrlOk returns a tuple with the RevokeSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSyncUrl

`func (o *GatewayUpdateProducerCustom) SetRevokeSyncUrl(v string)`

SetRevokeSyncUrl sets RevokeSyncUrl field to given value.


### GetRotateSyncUrl

`func (o *GatewayUpdateProducerCustom) GetRotateSyncUrl() string`

GetRotateSyncUrl returns the RotateSyncUrl field if non-nil, zero value otherwise.

### GetRotateSyncUrlOk

`func (o *GatewayUpdateProducerCustom) GetRotateSyncUrlOk() (*string, bool)`

GetRotateSyncUrlOk returns a tuple with the RotateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateSyncUrl

`func (o *GatewayUpdateProducerCustom) SetRotateSyncUrl(v string)`

SetRotateSyncUrl sets RotateSyncUrl field to given value.

### HasRotateSyncUrl

`func (o *GatewayUpdateProducerCustom) HasRotateSyncUrl() bool`

HasRotateSyncUrl returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerCustom) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerCustom) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerCustom) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerCustom) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *GatewayUpdateProducerCustom) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *GatewayUpdateProducerCustom) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *GatewayUpdateProducerCustom) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *GatewayUpdateProducerCustom) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerCustom) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerCustom) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerCustom) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerCustom) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerCustom) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerCustom) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerCustom) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerCustom) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerCustom) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerCustom) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerCustom) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerCustom) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayUpdateProducerCustom) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayUpdateProducerCustom) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayUpdateProducerCustom) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayUpdateProducerCustom) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


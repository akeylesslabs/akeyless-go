# GatewayCreateProducerCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateSyncUrl** | **string** | URL of an endpoint that implements /sync/create method, for example https://webhook.example.com/sync/create | 
**Name** | **string** | Producer name | 
**Payload** | Pointer to **string** | Secret payload to be sent with each create/revoke webhook request | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RevokeSyncUrl** | **string** | URL of an endpoint that implements /sync/revoke method, for example https://webhook.example.com/sync/revoke | 
**RotateSyncUrl** | Pointer to **string** | URL of an endpoint that implements /sync/rotate method, for example https://webhook.example.com/sync/rotate | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TimeoutSec** | Pointer to **int64** | Maximum allowed time in seconds for the webhook to return the results | [optional] [default to 60]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerCustom

`func NewGatewayCreateProducerCustom(createSyncUrl string, name string, revokeSyncUrl string, ) *GatewayCreateProducerCustom`

NewGatewayCreateProducerCustom instantiates a new GatewayCreateProducerCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerCustomWithDefaults

`func NewGatewayCreateProducerCustomWithDefaults() *GatewayCreateProducerCustom`

NewGatewayCreateProducerCustomWithDefaults instantiates a new GatewayCreateProducerCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateSyncUrl

`func (o *GatewayCreateProducerCustom) GetCreateSyncUrl() string`

GetCreateSyncUrl returns the CreateSyncUrl field if non-nil, zero value otherwise.

### GetCreateSyncUrlOk

`func (o *GatewayCreateProducerCustom) GetCreateSyncUrlOk() (*string, bool)`

GetCreateSyncUrlOk returns a tuple with the CreateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSyncUrl

`func (o *GatewayCreateProducerCustom) SetCreateSyncUrl(v string)`

SetCreateSyncUrl sets CreateSyncUrl field to given value.


### GetName

`func (o *GatewayCreateProducerCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerCustom) SetName(v string)`

SetName sets Name field to given value.


### GetPayload

`func (o *GatewayCreateProducerCustom) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *GatewayCreateProducerCustom) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *GatewayCreateProducerCustom) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *GatewayCreateProducerCustom) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerCustom) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerCustom) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerCustom) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerCustom) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRevokeSyncUrl

`func (o *GatewayCreateProducerCustom) GetRevokeSyncUrl() string`

GetRevokeSyncUrl returns the RevokeSyncUrl field if non-nil, zero value otherwise.

### GetRevokeSyncUrlOk

`func (o *GatewayCreateProducerCustom) GetRevokeSyncUrlOk() (*string, bool)`

GetRevokeSyncUrlOk returns a tuple with the RevokeSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSyncUrl

`func (o *GatewayCreateProducerCustom) SetRevokeSyncUrl(v string)`

SetRevokeSyncUrl sets RevokeSyncUrl field to given value.


### GetRotateSyncUrl

`func (o *GatewayCreateProducerCustom) GetRotateSyncUrl() string`

GetRotateSyncUrl returns the RotateSyncUrl field if non-nil, zero value otherwise.

### GetRotateSyncUrlOk

`func (o *GatewayCreateProducerCustom) GetRotateSyncUrlOk() (*string, bool)`

GetRotateSyncUrlOk returns a tuple with the RotateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateSyncUrl

`func (o *GatewayCreateProducerCustom) SetRotateSyncUrl(v string)`

SetRotateSyncUrl sets RotateSyncUrl field to given value.

### HasRotateSyncUrl

`func (o *GatewayCreateProducerCustom) HasRotateSyncUrl() bool`

HasRotateSyncUrl returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerCustom) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerCustom) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerCustom) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerCustom) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *GatewayCreateProducerCustom) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *GatewayCreateProducerCustom) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *GatewayCreateProducerCustom) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *GatewayCreateProducerCustom) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerCustom) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerCustom) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerCustom) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerCustom) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerCustom) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerCustom) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerCustom) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerCustom) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerCustom) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerCustom) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerCustom) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerCustom) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



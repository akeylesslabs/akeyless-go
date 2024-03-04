# DynamicSecretCreateCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminRotationIntervalDays** | Pointer to **int64** | Define rotation interval in days | [optional] 
**CreateSyncUrl** | **string** | URL of an endpoint that implements /sync/create method, for example https://webhook.example.com/sync/create | 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**EnableAdminRotation** | Pointer to **bool** | Should admin credentials be rotated | [optional] [default to false]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**Payload** | Pointer to **string** | Secret payload to be sent with each create/revoke webhook request | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RevokeSyncUrl** | **string** | URL of an endpoint that implements /sync/revoke method, for example https://webhook.example.com/sync/revoke | 
**RotateSyncUrl** | Pointer to **string** | URL of an endpoint that implements /sync/rotate method, for example https://webhook.example.com/sync/rotate | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TimeoutSec** | Pointer to **int64** | Maximum allowed time in seconds for the webhook to return the results | [optional] [default to 60]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateCustom

`func NewDynamicSecretCreateCustom(createSyncUrl string, name string, revokeSyncUrl string, ) *DynamicSecretCreateCustom`

NewDynamicSecretCreateCustom instantiates a new DynamicSecretCreateCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateCustomWithDefaults

`func NewDynamicSecretCreateCustomWithDefaults() *DynamicSecretCreateCustom`

NewDynamicSecretCreateCustomWithDefaults instantiates a new DynamicSecretCreateCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminRotationIntervalDays

`func (o *DynamicSecretCreateCustom) GetAdminRotationIntervalDays() int64`

GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field if non-nil, zero value otherwise.

### GetAdminRotationIntervalDaysOk

`func (o *DynamicSecretCreateCustom) GetAdminRotationIntervalDaysOk() (*int64, bool)`

GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminRotationIntervalDays

`func (o *DynamicSecretCreateCustom) SetAdminRotationIntervalDays(v int64)`

SetAdminRotationIntervalDays sets AdminRotationIntervalDays field to given value.

### HasAdminRotationIntervalDays

`func (o *DynamicSecretCreateCustom) HasAdminRotationIntervalDays() bool`

HasAdminRotationIntervalDays returns a boolean if a field has been set.

### GetCreateSyncUrl

`func (o *DynamicSecretCreateCustom) GetCreateSyncUrl() string`

GetCreateSyncUrl returns the CreateSyncUrl field if non-nil, zero value otherwise.

### GetCreateSyncUrlOk

`func (o *DynamicSecretCreateCustom) GetCreateSyncUrlOk() (*string, bool)`

GetCreateSyncUrlOk returns a tuple with the CreateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSyncUrl

`func (o *DynamicSecretCreateCustom) SetCreateSyncUrl(v string)`

SetCreateSyncUrl sets CreateSyncUrl field to given value.


### GetDeleteProtection

`func (o *DynamicSecretCreateCustom) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateCustom) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateCustom) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateCustom) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateCustom) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateCustom) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateCustom) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateCustom) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAdminRotation

`func (o *DynamicSecretCreateCustom) GetEnableAdminRotation() bool`

GetEnableAdminRotation returns the EnableAdminRotation field if non-nil, zero value otherwise.

### GetEnableAdminRotationOk

`func (o *DynamicSecretCreateCustom) GetEnableAdminRotationOk() (*bool, bool)`

GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdminRotation

`func (o *DynamicSecretCreateCustom) SetEnableAdminRotation(v bool)`

SetEnableAdminRotation sets EnableAdminRotation field to given value.

### HasEnableAdminRotation

`func (o *DynamicSecretCreateCustom) HasEnableAdminRotation() bool`

HasEnableAdminRotation returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateCustom) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateCustom) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateCustom) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateCustom) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateCustom) SetName(v string)`

SetName sets Name field to given value.


### GetPayload

`func (o *DynamicSecretCreateCustom) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DynamicSecretCreateCustom) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DynamicSecretCreateCustom) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DynamicSecretCreateCustom) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateCustom) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateCustom) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateCustom) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateCustom) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRevokeSyncUrl

`func (o *DynamicSecretCreateCustom) GetRevokeSyncUrl() string`

GetRevokeSyncUrl returns the RevokeSyncUrl field if non-nil, zero value otherwise.

### GetRevokeSyncUrlOk

`func (o *DynamicSecretCreateCustom) GetRevokeSyncUrlOk() (*string, bool)`

GetRevokeSyncUrlOk returns a tuple with the RevokeSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSyncUrl

`func (o *DynamicSecretCreateCustom) SetRevokeSyncUrl(v string)`

SetRevokeSyncUrl sets RevokeSyncUrl field to given value.


### GetRotateSyncUrl

`func (o *DynamicSecretCreateCustom) GetRotateSyncUrl() string`

GetRotateSyncUrl returns the RotateSyncUrl field if non-nil, zero value otherwise.

### GetRotateSyncUrlOk

`func (o *DynamicSecretCreateCustom) GetRotateSyncUrlOk() (*string, bool)`

GetRotateSyncUrlOk returns a tuple with the RotateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateSyncUrl

`func (o *DynamicSecretCreateCustom) SetRotateSyncUrl(v string)`

SetRotateSyncUrl sets RotateSyncUrl field to given value.

### HasRotateSyncUrl

`func (o *DynamicSecretCreateCustom) HasRotateSyncUrl() bool`

HasRotateSyncUrl returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretCreateCustom) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateCustom) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateCustom) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateCustom) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *DynamicSecretCreateCustom) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *DynamicSecretCreateCustom) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *DynamicSecretCreateCustom) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *DynamicSecretCreateCustom) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateCustom) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateCustom) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateCustom) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateCustom) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateCustom) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateCustom) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateCustom) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateCustom) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateCustom) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateCustom) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateCustom) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateCustom) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DynamicSecretUpdateCustom

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
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
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

### NewDynamicSecretUpdateCustom

`func NewDynamicSecretUpdateCustom(createSyncUrl string, name string, revokeSyncUrl string, ) *DynamicSecretUpdateCustom`

NewDynamicSecretUpdateCustom instantiates a new DynamicSecretUpdateCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateCustomWithDefaults

`func NewDynamicSecretUpdateCustomWithDefaults() *DynamicSecretUpdateCustom`

NewDynamicSecretUpdateCustomWithDefaults instantiates a new DynamicSecretUpdateCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminRotationIntervalDays

`func (o *DynamicSecretUpdateCustom) GetAdminRotationIntervalDays() int64`

GetAdminRotationIntervalDays returns the AdminRotationIntervalDays field if non-nil, zero value otherwise.

### GetAdminRotationIntervalDaysOk

`func (o *DynamicSecretUpdateCustom) GetAdminRotationIntervalDaysOk() (*int64, bool)`

GetAdminRotationIntervalDaysOk returns a tuple with the AdminRotationIntervalDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminRotationIntervalDays

`func (o *DynamicSecretUpdateCustom) SetAdminRotationIntervalDays(v int64)`

SetAdminRotationIntervalDays sets AdminRotationIntervalDays field to given value.

### HasAdminRotationIntervalDays

`func (o *DynamicSecretUpdateCustom) HasAdminRotationIntervalDays() bool`

HasAdminRotationIntervalDays returns a boolean if a field has been set.

### GetCreateSyncUrl

`func (o *DynamicSecretUpdateCustom) GetCreateSyncUrl() string`

GetCreateSyncUrl returns the CreateSyncUrl field if non-nil, zero value otherwise.

### GetCreateSyncUrlOk

`func (o *DynamicSecretUpdateCustom) GetCreateSyncUrlOk() (*string, bool)`

GetCreateSyncUrlOk returns a tuple with the CreateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSyncUrl

`func (o *DynamicSecretUpdateCustom) SetCreateSyncUrl(v string)`

SetCreateSyncUrl sets CreateSyncUrl field to given value.


### GetDeleteProtection

`func (o *DynamicSecretUpdateCustom) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateCustom) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateCustom) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateCustom) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretUpdateCustom) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretUpdateCustom) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretUpdateCustom) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretUpdateCustom) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnableAdminRotation

`func (o *DynamicSecretUpdateCustom) GetEnableAdminRotation() bool`

GetEnableAdminRotation returns the EnableAdminRotation field if non-nil, zero value otherwise.

### GetEnableAdminRotationOk

`func (o *DynamicSecretUpdateCustom) GetEnableAdminRotationOk() (*bool, bool)`

GetEnableAdminRotationOk returns a tuple with the EnableAdminRotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAdminRotation

`func (o *DynamicSecretUpdateCustom) SetEnableAdminRotation(v bool)`

SetEnableAdminRotation sets EnableAdminRotation field to given value.

### HasEnableAdminRotation

`func (o *DynamicSecretUpdateCustom) HasEnableAdminRotation() bool`

HasEnableAdminRotation returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateCustom) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateCustom) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateCustom) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateCustom) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateCustom) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateCustom) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateCustom) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateCustom) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateCustom) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetPayload

`func (o *DynamicSecretUpdateCustom) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *DynamicSecretUpdateCustom) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *DynamicSecretUpdateCustom) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *DynamicSecretUpdateCustom) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateCustom) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateCustom) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateCustom) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateCustom) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRevokeSyncUrl

`func (o *DynamicSecretUpdateCustom) GetRevokeSyncUrl() string`

GetRevokeSyncUrl returns the RevokeSyncUrl field if non-nil, zero value otherwise.

### GetRevokeSyncUrlOk

`func (o *DynamicSecretUpdateCustom) GetRevokeSyncUrlOk() (*string, bool)`

GetRevokeSyncUrlOk returns a tuple with the RevokeSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeSyncUrl

`func (o *DynamicSecretUpdateCustom) SetRevokeSyncUrl(v string)`

SetRevokeSyncUrl sets RevokeSyncUrl field to given value.


### GetRotateSyncUrl

`func (o *DynamicSecretUpdateCustom) GetRotateSyncUrl() string`

GetRotateSyncUrl returns the RotateSyncUrl field if non-nil, zero value otherwise.

### GetRotateSyncUrlOk

`func (o *DynamicSecretUpdateCustom) GetRotateSyncUrlOk() (*string, bool)`

GetRotateSyncUrlOk returns a tuple with the RotateSyncUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotateSyncUrl

`func (o *DynamicSecretUpdateCustom) SetRotateSyncUrl(v string)`

SetRotateSyncUrl sets RotateSyncUrl field to given value.

### HasRotateSyncUrl

`func (o *DynamicSecretUpdateCustom) HasRotateSyncUrl() bool`

HasRotateSyncUrl returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateCustom) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateCustom) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateCustom) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateCustom) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimeoutSec

`func (o *DynamicSecretUpdateCustom) GetTimeoutSec() int64`

GetTimeoutSec returns the TimeoutSec field if non-nil, zero value otherwise.

### GetTimeoutSecOk

`func (o *DynamicSecretUpdateCustom) GetTimeoutSecOk() (*int64, bool)`

GetTimeoutSecOk returns a tuple with the TimeoutSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSec

`func (o *DynamicSecretUpdateCustom) SetTimeoutSec(v int64)`

SetTimeoutSec sets TimeoutSec field to given value.

### HasTimeoutSec

`func (o *DynamicSecretUpdateCustom) HasTimeoutSec() bool`

HasTimeoutSec returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateCustom) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateCustom) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateCustom) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateCustom) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateCustom) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateCustom) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateCustom) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateCustom) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateCustom) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateCustom) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateCustom) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateCustom) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



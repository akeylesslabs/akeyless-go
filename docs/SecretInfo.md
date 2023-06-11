# SecretInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Expiration** | Pointer to **time.Time** |  | [optional] 
**LastRetrieved** | Pointer to **time.Time** |  | [optional] 
**Location** | Pointer to **map[string]interface{}** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SecretId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Tags** | Pointer to **map[string]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewSecretInfo

`func NewSecretInfo() *SecretInfo`

NewSecretInfo instantiates a new SecretInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecretInfoWithDefaults

`func NewSecretInfoWithDefaults() *SecretInfo`

NewSecretInfoWithDefaults instantiates a new SecretInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SecretInfo) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SecretInfo) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SecretInfo) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SecretInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *SecretInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecretInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecretInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecretInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpiration

`func (o *SecretInfo) GetExpiration() time.Time`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *SecretInfo) GetExpirationOk() (*time.Time, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *SecretInfo) SetExpiration(v time.Time)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *SecretInfo) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetLastRetrieved

`func (o *SecretInfo) GetLastRetrieved() time.Time`

GetLastRetrieved returns the LastRetrieved field if non-nil, zero value otherwise.

### GetLastRetrievedOk

`func (o *SecretInfo) GetLastRetrievedOk() (*time.Time, bool)`

GetLastRetrievedOk returns a tuple with the LastRetrieved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRetrieved

`func (o *SecretInfo) SetLastRetrieved(v time.Time)`

SetLastRetrieved sets LastRetrieved field to given value.

### HasLastRetrieved

`func (o *SecretInfo) HasLastRetrieved() bool`

HasLastRetrieved returns a boolean if a field has been set.

### GetLocation

`func (o *SecretInfo) GetLocation() map[string]interface{}`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SecretInfo) GetLocationOk() (*map[string]interface{}, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SecretInfo) SetLocation(v map[string]interface{})`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SecretInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *SecretInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecretInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecretInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecretInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecretId

`func (o *SecretInfo) GetSecretId() string`

GetSecretId returns the SecretId field if non-nil, zero value otherwise.

### GetSecretIdOk

`func (o *SecretInfo) GetSecretIdOk() (*string, bool)`

GetSecretIdOk returns a tuple with the SecretId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretId

`func (o *SecretInfo) SetSecretId(v string)`

SetSecretId sets SecretId field to given value.

### HasSecretId

`func (o *SecretInfo) HasSecretId() bool`

HasSecretId returns a boolean if a field has been set.

### GetStatus

`func (o *SecretInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecretInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecretInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SecretInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *SecretInfo) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SecretInfo) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SecretInfo) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SecretInfo) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *SecretInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecretInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecretInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecretInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



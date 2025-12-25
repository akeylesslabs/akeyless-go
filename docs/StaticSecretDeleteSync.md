# StaticSecretDeleteSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteFromUsc** | Pointer to **bool** | Delete the secret from the remote target USC as well | [optional] [default to false]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Static secret name | 
**RemoteSecretName** | Pointer to **string** | Remote Secret Name to disambiguate when multiple syncs exist under the same USC | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Universal Secret Connector name | 

## Methods

### NewStaticSecretDeleteSync

`func NewStaticSecretDeleteSync(name string, uscName string, ) *StaticSecretDeleteSync`

NewStaticSecretDeleteSync instantiates a new StaticSecretDeleteSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticSecretDeleteSyncWithDefaults

`func NewStaticSecretDeleteSyncWithDefaults() *StaticSecretDeleteSync`

NewStaticSecretDeleteSyncWithDefaults instantiates a new StaticSecretDeleteSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteFromUsc

`func (o *StaticSecretDeleteSync) GetDeleteFromUsc() bool`

GetDeleteFromUsc returns the DeleteFromUsc field if non-nil, zero value otherwise.

### GetDeleteFromUscOk

`func (o *StaticSecretDeleteSync) GetDeleteFromUscOk() (*bool, bool)`

GetDeleteFromUscOk returns a tuple with the DeleteFromUsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFromUsc

`func (o *StaticSecretDeleteSync) SetDeleteFromUsc(v bool)`

SetDeleteFromUsc sets DeleteFromUsc field to given value.

### HasDeleteFromUsc

`func (o *StaticSecretDeleteSync) HasDeleteFromUsc() bool`

HasDeleteFromUsc returns a boolean if a field has been set.

### GetJson

`func (o *StaticSecretDeleteSync) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *StaticSecretDeleteSync) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *StaticSecretDeleteSync) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *StaticSecretDeleteSync) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *StaticSecretDeleteSync) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticSecretDeleteSync) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticSecretDeleteSync) SetName(v string)`

SetName sets Name field to given value.


### GetRemoteSecretName

`func (o *StaticSecretDeleteSync) GetRemoteSecretName() string`

GetRemoteSecretName returns the RemoteSecretName field if non-nil, zero value otherwise.

### GetRemoteSecretNameOk

`func (o *StaticSecretDeleteSync) GetRemoteSecretNameOk() (*string, bool)`

GetRemoteSecretNameOk returns a tuple with the RemoteSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSecretName

`func (o *StaticSecretDeleteSync) SetRemoteSecretName(v string)`

SetRemoteSecretName sets RemoteSecretName field to given value.

### HasRemoteSecretName

`func (o *StaticSecretDeleteSync) HasRemoteSecretName() bool`

HasRemoteSecretName returns a boolean if a field has been set.

### GetToken

`func (o *StaticSecretDeleteSync) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *StaticSecretDeleteSync) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *StaticSecretDeleteSync) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *StaticSecretDeleteSync) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *StaticSecretDeleteSync) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *StaticSecretDeleteSync) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *StaticSecretDeleteSync) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *StaticSecretDeleteSync) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *StaticSecretDeleteSync) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *StaticSecretDeleteSync) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *StaticSecretDeleteSync) SetUscName(v string)`

SetUscName sets UscName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



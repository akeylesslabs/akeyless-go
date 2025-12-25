# RotatedSecretSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteRemote** | Pointer to **bool** | Delete the secret from remote secret manager (for association create/update) | [optional] 
**FilterSecretValue** | Pointer to **string** | JQ expression to filter or transform the secret value | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Rotated secret name | 
**Namespace** | Pointer to **string** | Vault namespace, releavnt only for Hashicorp Vault Target | [optional] 
**RemoteSecretName** | Pointer to **string** | Remote Secret Name that will be synced on the remote endpoint | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | Pointer to **string** | Universal Secret Connector name, If not provided all attached USC&#39;s will be synced | [optional] 

## Methods

### NewRotatedSecretSync

`func NewRotatedSecretSync(name string, ) *RotatedSecretSync`

NewRotatedSecretSync instantiates a new RotatedSecretSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotatedSecretSyncWithDefaults

`func NewRotatedSecretSyncWithDefaults() *RotatedSecretSync`

NewRotatedSecretSyncWithDefaults instantiates a new RotatedSecretSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteRemote

`func (o *RotatedSecretSync) GetDeleteRemote() bool`

GetDeleteRemote returns the DeleteRemote field if non-nil, zero value otherwise.

### GetDeleteRemoteOk

`func (o *RotatedSecretSync) GetDeleteRemoteOk() (*bool, bool)`

GetDeleteRemoteOk returns a tuple with the DeleteRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRemote

`func (o *RotatedSecretSync) SetDeleteRemote(v bool)`

SetDeleteRemote sets DeleteRemote field to given value.

### HasDeleteRemote

`func (o *RotatedSecretSync) HasDeleteRemote() bool`

HasDeleteRemote returns a boolean if a field has been set.

### GetFilterSecretValue

`func (o *RotatedSecretSync) GetFilterSecretValue() string`

GetFilterSecretValue returns the FilterSecretValue field if non-nil, zero value otherwise.

### GetFilterSecretValueOk

`func (o *RotatedSecretSync) GetFilterSecretValueOk() (*string, bool)`

GetFilterSecretValueOk returns a tuple with the FilterSecretValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterSecretValue

`func (o *RotatedSecretSync) SetFilterSecretValue(v string)`

SetFilterSecretValue sets FilterSecretValue field to given value.

### HasFilterSecretValue

`func (o *RotatedSecretSync) HasFilterSecretValue() bool`

HasFilterSecretValue returns a boolean if a field has been set.

### GetJson

`func (o *RotatedSecretSync) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *RotatedSecretSync) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *RotatedSecretSync) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *RotatedSecretSync) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *RotatedSecretSync) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RotatedSecretSync) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RotatedSecretSync) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *RotatedSecretSync) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RotatedSecretSync) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RotatedSecretSync) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RotatedSecretSync) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRemoteSecretName

`func (o *RotatedSecretSync) GetRemoteSecretName() string`

GetRemoteSecretName returns the RemoteSecretName field if non-nil, zero value otherwise.

### GetRemoteSecretNameOk

`func (o *RotatedSecretSync) GetRemoteSecretNameOk() (*string, bool)`

GetRemoteSecretNameOk returns a tuple with the RemoteSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSecretName

`func (o *RotatedSecretSync) SetRemoteSecretName(v string)`

SetRemoteSecretName sets RemoteSecretName field to given value.

### HasRemoteSecretName

`func (o *RotatedSecretSync) HasRemoteSecretName() bool`

HasRemoteSecretName returns a boolean if a field has been set.

### GetToken

`func (o *RotatedSecretSync) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *RotatedSecretSync) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *RotatedSecretSync) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *RotatedSecretSync) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *RotatedSecretSync) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *RotatedSecretSync) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *RotatedSecretSync) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *RotatedSecretSync) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *RotatedSecretSync) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *RotatedSecretSync) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *RotatedSecretSync) SetUscName(v string)`

SetUscName sets UscName field to given value.

### HasUscName

`func (o *RotatedSecretSync) HasUscName() bool`

HasUscName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



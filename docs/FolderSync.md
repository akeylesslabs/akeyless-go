# FolderSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**DeleteRemote** | Pointer to **bool** | Delete the secret from the remote target as well | [optional] 
**EngineName** | Pointer to **string** | Hashi Vault engine name prefix, must end with &#39;/&#39; | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Folder name | 
**Namespace** | Pointer to **string** | Vault namespace, relevant only for Hashicorp Vault Target | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | Pointer to **string** | Universal Secret Connector name, If not provided all attached USC&#39;s will be synced | [optional] 

## Methods

### NewFolderSync

`func NewFolderSync(name string, ) *FolderSync`

NewFolderSync instantiates a new FolderSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderSyncWithDefaults

`func NewFolderSyncWithDefaults() *FolderSync`

NewFolderSyncWithDefaults instantiates a new FolderSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *FolderSync) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *FolderSync) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *FolderSync) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *FolderSync) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetDeleteRemote

`func (o *FolderSync) GetDeleteRemote() bool`

GetDeleteRemote returns the DeleteRemote field if non-nil, zero value otherwise.

### GetDeleteRemoteOk

`func (o *FolderSync) GetDeleteRemoteOk() (*bool, bool)`

GetDeleteRemoteOk returns a tuple with the DeleteRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteRemote

`func (o *FolderSync) SetDeleteRemote(v bool)`

SetDeleteRemote sets DeleteRemote field to given value.

### HasDeleteRemote

`func (o *FolderSync) HasDeleteRemote() bool`

HasDeleteRemote returns a boolean if a field has been set.

### GetEngineName

`func (o *FolderSync) GetEngineName() string`

GetEngineName returns the EngineName field if non-nil, zero value otherwise.

### GetEngineNameOk

`func (o *FolderSync) GetEngineNameOk() (*string, bool)`

GetEngineNameOk returns a tuple with the EngineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineName

`func (o *FolderSync) SetEngineName(v string)`

SetEngineName sets EngineName field to given value.

### HasEngineName

`func (o *FolderSync) HasEngineName() bool`

HasEngineName returns a boolean if a field has been set.

### GetJson

`func (o *FolderSync) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *FolderSync) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *FolderSync) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *FolderSync) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *FolderSync) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderSync) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderSync) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *FolderSync) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *FolderSync) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *FolderSync) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *FolderSync) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetToken

`func (o *FolderSync) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FolderSync) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FolderSync) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FolderSync) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *FolderSync) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *FolderSync) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *FolderSync) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *FolderSync) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *FolderSync) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *FolderSync) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *FolderSync) SetUscName(v string)`

SetUscName sets UscName field to given value.

### HasUscName

`func (o *FolderSync) HasUscName() bool`

HasUscName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



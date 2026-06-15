# FolderDeleteSync

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**DeleteFromUsc** | Pointer to **bool** | Delete the secrets from the remote target usc as well | [optional] [default to false]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Folder name | 
**RemoteSecretName** | Pointer to **string** | Remote Secret Name to delete when multiple syncs exist under the same USC | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscName** | **string** | Universal Secret Connector name | 

## Methods

### NewFolderDeleteSync

`func NewFolderDeleteSync(name string, uscName string, ) *FolderDeleteSync`

NewFolderDeleteSync instantiates a new FolderDeleteSync object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderDeleteSyncWithDefaults

`func NewFolderDeleteSyncWithDefaults() *FolderDeleteSync`

NewFolderDeleteSyncWithDefaults instantiates a new FolderDeleteSync object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *FolderDeleteSync) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *FolderDeleteSync) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *FolderDeleteSync) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *FolderDeleteSync) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetDeleteFromUsc

`func (o *FolderDeleteSync) GetDeleteFromUsc() bool`

GetDeleteFromUsc returns the DeleteFromUsc field if non-nil, zero value otherwise.

### GetDeleteFromUscOk

`func (o *FolderDeleteSync) GetDeleteFromUscOk() (*bool, bool)`

GetDeleteFromUscOk returns a tuple with the DeleteFromUsc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteFromUsc

`func (o *FolderDeleteSync) SetDeleteFromUsc(v bool)`

SetDeleteFromUsc sets DeleteFromUsc field to given value.

### HasDeleteFromUsc

`func (o *FolderDeleteSync) HasDeleteFromUsc() bool`

HasDeleteFromUsc returns a boolean if a field has been set.

### GetJson

`func (o *FolderDeleteSync) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *FolderDeleteSync) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *FolderDeleteSync) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *FolderDeleteSync) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *FolderDeleteSync) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderDeleteSync) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderDeleteSync) SetName(v string)`

SetName sets Name field to given value.


### GetRemoteSecretName

`func (o *FolderDeleteSync) GetRemoteSecretName() string`

GetRemoteSecretName returns the RemoteSecretName field if non-nil, zero value otherwise.

### GetRemoteSecretNameOk

`func (o *FolderDeleteSync) GetRemoteSecretNameOk() (*string, bool)`

GetRemoteSecretNameOk returns a tuple with the RemoteSecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteSecretName

`func (o *FolderDeleteSync) SetRemoteSecretName(v string)`

SetRemoteSecretName sets RemoteSecretName field to given value.

### HasRemoteSecretName

`func (o *FolderDeleteSync) HasRemoteSecretName() bool`

HasRemoteSecretName returns a boolean if a field has been set.

### GetToken

`func (o *FolderDeleteSync) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FolderDeleteSync) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FolderDeleteSync) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FolderDeleteSync) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *FolderDeleteSync) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *FolderDeleteSync) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *FolderDeleteSync) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *FolderDeleteSync) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscName

`func (o *FolderDeleteSync) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *FolderDeleteSync) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *FolderDeleteSync) SetUscName(v string)`

SetUscName sets UscName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



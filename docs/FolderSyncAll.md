# FolderSyncAll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Folder name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewFolderSyncAll

`func NewFolderSyncAll(name string, ) *FolderSyncAll`

NewFolderSyncAll instantiates a new FolderSyncAll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFolderSyncAllWithDefaults

`func NewFolderSyncAllWithDefaults() *FolderSyncAll`

NewFolderSyncAllWithDefaults instantiates a new FolderSyncAll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *FolderSyncAll) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *FolderSyncAll) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *FolderSyncAll) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *FolderSyncAll) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetJson

`func (o *FolderSyncAll) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *FolderSyncAll) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *FolderSyncAll) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *FolderSyncAll) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *FolderSyncAll) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FolderSyncAll) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FolderSyncAll) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *FolderSyncAll) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *FolderSyncAll) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *FolderSyncAll) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *FolderSyncAll) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *FolderSyncAll) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *FolderSyncAll) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *FolderSyncAll) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *FolderSyncAll) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



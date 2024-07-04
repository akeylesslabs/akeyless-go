# TargetUpdateDockerhub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**DockerhubPassword** | Pointer to **string** | Password for docker repository | [optional] 
**DockerhubUsername** | Pointer to **string** | Username for docker repository | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateDockerhub

`func NewTargetUpdateDockerhub(name string, ) *TargetUpdateDockerhub`

NewTargetUpdateDockerhub instantiates a new TargetUpdateDockerhub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateDockerhubWithDefaults

`func NewTargetUpdateDockerhubWithDefaults() *TargetUpdateDockerhub`

NewTargetUpdateDockerhubWithDefaults instantiates a new TargetUpdateDockerhub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TargetUpdateDockerhub) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateDockerhub) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateDockerhub) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateDockerhub) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDockerhubPassword

`func (o *TargetUpdateDockerhub) GetDockerhubPassword() string`

GetDockerhubPassword returns the DockerhubPassword field if non-nil, zero value otherwise.

### GetDockerhubPasswordOk

`func (o *TargetUpdateDockerhub) GetDockerhubPasswordOk() (*string, bool)`

GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubPassword

`func (o *TargetUpdateDockerhub) SetDockerhubPassword(v string)`

SetDockerhubPassword sets DockerhubPassword field to given value.

### HasDockerhubPassword

`func (o *TargetUpdateDockerhub) HasDockerhubPassword() bool`

HasDockerhubPassword returns a boolean if a field has been set.

### GetDockerhubUsername

`func (o *TargetUpdateDockerhub) GetDockerhubUsername() string`

GetDockerhubUsername returns the DockerhubUsername field if non-nil, zero value otherwise.

### GetDockerhubUsernameOk

`func (o *TargetUpdateDockerhub) GetDockerhubUsernameOk() (*string, bool)`

GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubUsername

`func (o *TargetUpdateDockerhub) SetDockerhubUsername(v string)`

SetDockerhubUsername sets DockerhubUsername field to given value.

### HasDockerhubUsername

`func (o *TargetUpdateDockerhub) HasDockerhubUsername() bool`

HasDockerhubUsername returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateDockerhub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateDockerhub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateDockerhub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateDockerhub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateDockerhub) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateDockerhub) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateDockerhub) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateDockerhub) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateDockerhub) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateDockerhub) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateDockerhub) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateDockerhub) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateDockerhub) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateDockerhub) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateDockerhub) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateDockerhub) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateDockerhub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateDockerhub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateDockerhub) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateDockerhub) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateDockerhub) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateDockerhub) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateDockerhub) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateDockerhub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateDockerhub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateDockerhub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateDockerhub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateDockerhub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateDockerhub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateDockerhub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateDockerhub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



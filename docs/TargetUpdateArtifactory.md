# TargetUpdateArtifactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryAdminName** | **string** | Artifactory Admin Name | 
**ArtifactoryAdminPwd** | **string** | Artifactory Admin password | 
**BaseUrl** | **string** | Base URL | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateArtifactory

`func NewTargetUpdateArtifactory(artifactoryAdminName string, artifactoryAdminPwd string, baseUrl string, name string, ) *TargetUpdateArtifactory`

NewTargetUpdateArtifactory instantiates a new TargetUpdateArtifactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateArtifactoryWithDefaults

`func NewTargetUpdateArtifactoryWithDefaults() *TargetUpdateArtifactory`

NewTargetUpdateArtifactoryWithDefaults instantiates a new TargetUpdateArtifactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryAdminName

`func (o *TargetUpdateArtifactory) GetArtifactoryAdminName() string`

GetArtifactoryAdminName returns the ArtifactoryAdminName field if non-nil, zero value otherwise.

### GetArtifactoryAdminNameOk

`func (o *TargetUpdateArtifactory) GetArtifactoryAdminNameOk() (*string, bool)`

GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminName

`func (o *TargetUpdateArtifactory) SetArtifactoryAdminName(v string)`

SetArtifactoryAdminName sets ArtifactoryAdminName field to given value.


### GetArtifactoryAdminPwd

`func (o *TargetUpdateArtifactory) GetArtifactoryAdminPwd() string`

GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field if non-nil, zero value otherwise.

### GetArtifactoryAdminPwdOk

`func (o *TargetUpdateArtifactory) GetArtifactoryAdminPwdOk() (*string, bool)`

GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminPwd

`func (o *TargetUpdateArtifactory) SetArtifactoryAdminPwd(v string)`

SetArtifactoryAdminPwd sets ArtifactoryAdminPwd field to given value.


### GetBaseUrl

`func (o *TargetUpdateArtifactory) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *TargetUpdateArtifactory) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *TargetUpdateArtifactory) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetDescription

`func (o *TargetUpdateArtifactory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateArtifactory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateArtifactory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateArtifactory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateArtifactory) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateArtifactory) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateArtifactory) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateArtifactory) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateArtifactory) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateArtifactory) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateArtifactory) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateArtifactory) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateArtifactory) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateArtifactory) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateArtifactory) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateArtifactory) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateArtifactory) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateArtifactory) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateArtifactory) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateArtifactory) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateArtifactory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateArtifactory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateArtifactory) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateArtifactory) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateArtifactory) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateArtifactory) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateArtifactory) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateArtifactory) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateArtifactory) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateArtifactory) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateArtifactory) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateArtifactory) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateArtifactory) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateArtifactory) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateArtifactory) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



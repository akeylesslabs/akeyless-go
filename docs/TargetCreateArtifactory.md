# TargetCreateArtifactory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArtifactoryAdminName** | **string** | Artifactory Admin Name | 
**ArtifactoryAdminPwd** | **string** | Artifactory Admin password | 
**BaseUrl** | **string** | Base URL | 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateArtifactory

`func NewTargetCreateArtifactory(artifactoryAdminName string, artifactoryAdminPwd string, baseUrl string, name string, ) *TargetCreateArtifactory`

NewTargetCreateArtifactory instantiates a new TargetCreateArtifactory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateArtifactoryWithDefaults

`func NewTargetCreateArtifactoryWithDefaults() *TargetCreateArtifactory`

NewTargetCreateArtifactoryWithDefaults instantiates a new TargetCreateArtifactory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArtifactoryAdminName

`func (o *TargetCreateArtifactory) GetArtifactoryAdminName() string`

GetArtifactoryAdminName returns the ArtifactoryAdminName field if non-nil, zero value otherwise.

### GetArtifactoryAdminNameOk

`func (o *TargetCreateArtifactory) GetArtifactoryAdminNameOk() (*string, bool)`

GetArtifactoryAdminNameOk returns a tuple with the ArtifactoryAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminName

`func (o *TargetCreateArtifactory) SetArtifactoryAdminName(v string)`

SetArtifactoryAdminName sets ArtifactoryAdminName field to given value.


### GetArtifactoryAdminPwd

`func (o *TargetCreateArtifactory) GetArtifactoryAdminPwd() string`

GetArtifactoryAdminPwd returns the ArtifactoryAdminPwd field if non-nil, zero value otherwise.

### GetArtifactoryAdminPwdOk

`func (o *TargetCreateArtifactory) GetArtifactoryAdminPwdOk() (*string, bool)`

GetArtifactoryAdminPwdOk returns a tuple with the ArtifactoryAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArtifactoryAdminPwd

`func (o *TargetCreateArtifactory) SetArtifactoryAdminPwd(v string)`

SetArtifactoryAdminPwd sets ArtifactoryAdminPwd field to given value.


### GetBaseUrl

`func (o *TargetCreateArtifactory) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *TargetCreateArtifactory) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *TargetCreateArtifactory) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetDescription

`func (o *TargetCreateArtifactory) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateArtifactory) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateArtifactory) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateArtifactory) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateArtifactory) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateArtifactory) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateArtifactory) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateArtifactory) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateArtifactory) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateArtifactory) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateArtifactory) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateArtifactory) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateArtifactory) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateArtifactory) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateArtifactory) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateArtifactory) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateArtifactory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateArtifactory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateArtifactory) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *TargetCreateArtifactory) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateArtifactory) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateArtifactory) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateArtifactory) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateArtifactory) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateArtifactory) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateArtifactory) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateArtifactory) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



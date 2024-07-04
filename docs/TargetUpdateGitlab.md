# TargetUpdateGitlab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**GitlabAccessToken** | Pointer to **string** | Gitlab access token | [optional] 
**GitlabCertificate** | Pointer to **string** | Gitlab tls certificate (base64 encoded) | [optional] 
**GitlabUrl** | Pointer to **string** | Gitlab base url | [optional] [default to "https://gitlab.com/"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateGitlab

`func NewTargetUpdateGitlab(name string, ) *TargetUpdateGitlab`

NewTargetUpdateGitlab instantiates a new TargetUpdateGitlab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateGitlabWithDefaults

`func NewTargetUpdateGitlabWithDefaults() *TargetUpdateGitlab`

NewTargetUpdateGitlabWithDefaults instantiates a new TargetUpdateGitlab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TargetUpdateGitlab) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateGitlab) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateGitlab) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateGitlab) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGitlabAccessToken

`func (o *TargetUpdateGitlab) GetGitlabAccessToken() string`

GetGitlabAccessToken returns the GitlabAccessToken field if non-nil, zero value otherwise.

### GetGitlabAccessTokenOk

`func (o *TargetUpdateGitlab) GetGitlabAccessTokenOk() (*string, bool)`

GetGitlabAccessTokenOk returns a tuple with the GitlabAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabAccessToken

`func (o *TargetUpdateGitlab) SetGitlabAccessToken(v string)`

SetGitlabAccessToken sets GitlabAccessToken field to given value.

### HasGitlabAccessToken

`func (o *TargetUpdateGitlab) HasGitlabAccessToken() bool`

HasGitlabAccessToken returns a boolean if a field has been set.

### GetGitlabCertificate

`func (o *TargetUpdateGitlab) GetGitlabCertificate() string`

GetGitlabCertificate returns the GitlabCertificate field if non-nil, zero value otherwise.

### GetGitlabCertificateOk

`func (o *TargetUpdateGitlab) GetGitlabCertificateOk() (*string, bool)`

GetGitlabCertificateOk returns a tuple with the GitlabCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabCertificate

`func (o *TargetUpdateGitlab) SetGitlabCertificate(v string)`

SetGitlabCertificate sets GitlabCertificate field to given value.

### HasGitlabCertificate

`func (o *TargetUpdateGitlab) HasGitlabCertificate() bool`

HasGitlabCertificate returns a boolean if a field has been set.

### GetGitlabUrl

`func (o *TargetUpdateGitlab) GetGitlabUrl() string`

GetGitlabUrl returns the GitlabUrl field if non-nil, zero value otherwise.

### GetGitlabUrlOk

`func (o *TargetUpdateGitlab) GetGitlabUrlOk() (*string, bool)`

GetGitlabUrlOk returns a tuple with the GitlabUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabUrl

`func (o *TargetUpdateGitlab) SetGitlabUrl(v string)`

SetGitlabUrl sets GitlabUrl field to given value.

### HasGitlabUrl

`func (o *TargetUpdateGitlab) HasGitlabUrl() bool`

HasGitlabUrl returns a boolean if a field has been set.

### GetJson

`func (o *TargetUpdateGitlab) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateGitlab) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateGitlab) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateGitlab) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateGitlab) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateGitlab) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateGitlab) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateGitlab) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateGitlab) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateGitlab) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateGitlab) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateGitlab) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateGitlab) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateGitlab) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateGitlab) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateGitlab) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateGitlab) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateGitlab) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateGitlab) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateGitlab) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateGitlab) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateGitlab) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateGitlab) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateGitlab) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateGitlab) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateGitlab) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateGitlab) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateGitlab) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateGitlab) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateGitlab) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateGitlab) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TargetCreateGitlab

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**GitlabAccessToken** | Pointer to **string** | Gitlab access token | [optional] 
**GitlabCertificate** | Pointer to **string** | Gitlab tls certificate (base64 encoded) | [optional] 
**GitlabUrl** | Pointer to **string** | Gitlab base url | [optional] [default to "https://gitlab.com/"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetCreateGitlab

`func NewTargetCreateGitlab(name string, ) *TargetCreateGitlab`

NewTargetCreateGitlab instantiates a new TargetCreateGitlab object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateGitlabWithDefaults

`func NewTargetCreateGitlabWithDefaults() *TargetCreateGitlab`

NewTargetCreateGitlabWithDefaults instantiates a new TargetCreateGitlab object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *TargetCreateGitlab) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateGitlab) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateGitlab) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateGitlab) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGitlabAccessToken

`func (o *TargetCreateGitlab) GetGitlabAccessToken() string`

GetGitlabAccessToken returns the GitlabAccessToken field if non-nil, zero value otherwise.

### GetGitlabAccessTokenOk

`func (o *TargetCreateGitlab) GetGitlabAccessTokenOk() (*string, bool)`

GetGitlabAccessTokenOk returns a tuple with the GitlabAccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabAccessToken

`func (o *TargetCreateGitlab) SetGitlabAccessToken(v string)`

SetGitlabAccessToken sets GitlabAccessToken field to given value.

### HasGitlabAccessToken

`func (o *TargetCreateGitlab) HasGitlabAccessToken() bool`

HasGitlabAccessToken returns a boolean if a field has been set.

### GetGitlabCertificate

`func (o *TargetCreateGitlab) GetGitlabCertificate() string`

GetGitlabCertificate returns the GitlabCertificate field if non-nil, zero value otherwise.

### GetGitlabCertificateOk

`func (o *TargetCreateGitlab) GetGitlabCertificateOk() (*string, bool)`

GetGitlabCertificateOk returns a tuple with the GitlabCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabCertificate

`func (o *TargetCreateGitlab) SetGitlabCertificate(v string)`

SetGitlabCertificate sets GitlabCertificate field to given value.

### HasGitlabCertificate

`func (o *TargetCreateGitlab) HasGitlabCertificate() bool`

HasGitlabCertificate returns a boolean if a field has been set.

### GetGitlabUrl

`func (o *TargetCreateGitlab) GetGitlabUrl() string`

GetGitlabUrl returns the GitlabUrl field if non-nil, zero value otherwise.

### GetGitlabUrlOk

`func (o *TargetCreateGitlab) GetGitlabUrlOk() (*string, bool)`

GetGitlabUrlOk returns a tuple with the GitlabUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitlabUrl

`func (o *TargetCreateGitlab) SetGitlabUrl(v string)`

SetGitlabUrl sets GitlabUrl field to given value.

### HasGitlabUrl

`func (o *TargetCreateGitlab) HasGitlabUrl() bool`

HasGitlabUrl returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateGitlab) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateGitlab) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateGitlab) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateGitlab) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateGitlab) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateGitlab) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateGitlab) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateGitlab) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateGitlab) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateGitlab) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateGitlab) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateGitlab) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateGitlab) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateGitlab) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateGitlab) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *TargetCreateGitlab) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateGitlab) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateGitlab) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateGitlab) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateGitlab) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateGitlab) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateGitlab) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateGitlab) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TargetCreateSplunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | Splunk token audience (required when using token authentication for rotation) | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | Splunk Password (used when authenticating with username/password) | [optional] 
**Token** | Pointer to **string** | Splunk Token (used when authenticating with token) | [optional] 
**TokenOwner** | Pointer to **string** | Splunk Token Owner (required when using token authentication for rotation) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Url** | **string** | Splunk server URL | 
**UseTls** | Pointer to **bool** | Use TLS certificate verification when connecting to the Splunk management API | [optional] [default to true]
**Username** | Pointer to **string** | Splunk Username (used when authenticating with username/password) | [optional] 

## Methods

### NewTargetCreateSplunk

`func NewTargetCreateSplunk(name string, url string, ) *TargetCreateSplunk`

NewTargetCreateSplunk instantiates a new TargetCreateSplunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetCreateSplunkWithDefaults

`func NewTargetCreateSplunkWithDefaults() *TargetCreateSplunk`

NewTargetCreateSplunkWithDefaults instantiates a new TargetCreateSplunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *TargetCreateSplunk) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *TargetCreateSplunk) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *TargetCreateSplunk) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *TargetCreateSplunk) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetDescription

`func (o *TargetCreateSplunk) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetCreateSplunk) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetCreateSplunk) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetCreateSplunk) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *TargetCreateSplunk) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetCreateSplunk) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetCreateSplunk) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetCreateSplunk) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *TargetCreateSplunk) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetCreateSplunk) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetCreateSplunk) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetCreateSplunk) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetCreateSplunk) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetCreateSplunk) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetCreateSplunk) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetCreateSplunk) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetCreateSplunk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetCreateSplunk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetCreateSplunk) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *TargetCreateSplunk) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *TargetCreateSplunk) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *TargetCreateSplunk) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *TargetCreateSplunk) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *TargetCreateSplunk) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetCreateSplunk) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetCreateSplunk) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetCreateSplunk) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenOwner

`func (o *TargetCreateSplunk) GetTokenOwner() string`

GetTokenOwner returns the TokenOwner field if non-nil, zero value otherwise.

### GetTokenOwnerOk

`func (o *TargetCreateSplunk) GetTokenOwnerOk() (*string, bool)`

GetTokenOwnerOk returns a tuple with the TokenOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOwner

`func (o *TargetCreateSplunk) SetTokenOwner(v string)`

SetTokenOwner sets TokenOwner field to given value.

### HasTokenOwner

`func (o *TargetCreateSplunk) HasTokenOwner() bool`

HasTokenOwner returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetCreateSplunk) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetCreateSplunk) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetCreateSplunk) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetCreateSplunk) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUrl

`func (o *TargetCreateSplunk) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *TargetCreateSplunk) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *TargetCreateSplunk) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUseTls

`func (o *TargetCreateSplunk) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *TargetCreateSplunk) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *TargetCreateSplunk) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *TargetCreateSplunk) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUsername

`func (o *TargetCreateSplunk) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TargetCreateSplunk) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TargetCreateSplunk) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *TargetCreateSplunk) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



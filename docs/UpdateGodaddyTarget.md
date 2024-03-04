# UpdateGodaddyTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Key of the api credentials to the Godaddy account | 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ImapFqdn** | **string** | ImapFQDN of the IMAP service, FQDN or IPv4 address. Must be FQDN if the IMAP is using TLS | 
**ImapPassword** | **string** | ImapPassword to access the IMAP service | 
**ImapPort** | Pointer to **string** | ImapPort of the IMAP service | [optional] [default to "993"]
**ImapUsername** | **string** | ImapUsername to access the IMAP service | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Secret** | **string** | Secret of the api credentials to the Godaddy account | 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdateGodaddyTarget

`func NewUpdateGodaddyTarget(apiKey string, imapFqdn string, imapPassword string, imapUsername string, name string, secret string, ) *UpdateGodaddyTarget`

NewUpdateGodaddyTarget instantiates a new UpdateGodaddyTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGodaddyTargetWithDefaults

`func NewUpdateGodaddyTargetWithDefaults() *UpdateGodaddyTarget`

NewUpdateGodaddyTargetWithDefaults instantiates a new UpdateGodaddyTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *UpdateGodaddyTarget) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateGodaddyTarget) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateGodaddyTarget) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetComment

`func (o *UpdateGodaddyTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateGodaddyTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateGodaddyTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateGodaddyTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateGodaddyTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateGodaddyTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateGodaddyTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateGodaddyTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImapFqdn

`func (o *UpdateGodaddyTarget) GetImapFqdn() string`

GetImapFqdn returns the ImapFqdn field if non-nil, zero value otherwise.

### GetImapFqdnOk

`func (o *UpdateGodaddyTarget) GetImapFqdnOk() (*string, bool)`

GetImapFqdnOk returns a tuple with the ImapFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapFqdn

`func (o *UpdateGodaddyTarget) SetImapFqdn(v string)`

SetImapFqdn sets ImapFqdn field to given value.


### GetImapPassword

`func (o *UpdateGodaddyTarget) GetImapPassword() string`

GetImapPassword returns the ImapPassword field if non-nil, zero value otherwise.

### GetImapPasswordOk

`func (o *UpdateGodaddyTarget) GetImapPasswordOk() (*string, bool)`

GetImapPasswordOk returns a tuple with the ImapPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapPassword

`func (o *UpdateGodaddyTarget) SetImapPassword(v string)`

SetImapPassword sets ImapPassword field to given value.


### GetImapPort

`func (o *UpdateGodaddyTarget) GetImapPort() string`

GetImapPort returns the ImapPort field if non-nil, zero value otherwise.

### GetImapPortOk

`func (o *UpdateGodaddyTarget) GetImapPortOk() (*string, bool)`

GetImapPortOk returns a tuple with the ImapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapPort

`func (o *UpdateGodaddyTarget) SetImapPort(v string)`

SetImapPort sets ImapPort field to given value.

### HasImapPort

`func (o *UpdateGodaddyTarget) HasImapPort() bool`

HasImapPort returns a boolean if a field has been set.

### GetImapUsername

`func (o *UpdateGodaddyTarget) GetImapUsername() string`

GetImapUsername returns the ImapUsername field if non-nil, zero value otherwise.

### GetImapUsernameOk

`func (o *UpdateGodaddyTarget) GetImapUsernameOk() (*string, bool)`

GetImapUsernameOk returns a tuple with the ImapUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapUsername

`func (o *UpdateGodaddyTarget) SetImapUsername(v string)`

SetImapUsername sets ImapUsername field to given value.


### GetJson

`func (o *UpdateGodaddyTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateGodaddyTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateGodaddyTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateGodaddyTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateGodaddyTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateGodaddyTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateGodaddyTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateGodaddyTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateGodaddyTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateGodaddyTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateGodaddyTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateGodaddyTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateGodaddyTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateGodaddyTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateGodaddyTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateGodaddyTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateGodaddyTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateGodaddyTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateGodaddyTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateGodaddyTarget) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateGodaddyTarget) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateGodaddyTarget) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetTimeout

`func (o *UpdateGodaddyTarget) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateGodaddyTarget) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateGodaddyTarget) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateGodaddyTarget) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *UpdateGodaddyTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateGodaddyTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateGodaddyTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateGodaddyTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateGodaddyTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateGodaddyTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateGodaddyTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateGodaddyTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateGodaddyTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateGodaddyTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateGodaddyTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateGodaddyTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



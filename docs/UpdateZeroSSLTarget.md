# UpdateZeroSSLTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | API Key of the ZeroSSLTarget account | 
**Comment** | Pointer to **string** | Deprecated - use description | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ImapFqdn** | **string** | ImapFQDN of the IMAP service, FQDN or IPv4 address. Must be FQDN if the IMAP is using TLS | 
**ImapPassword** | **string** | ImapPassword to access the IMAP service | 
**ImapPort** | Pointer to **string** | ImapPort of the IMAP service | [optional] [default to "993"]
**ImapTargetEmail** | Pointer to **string** | ImapValidationEmail to use when asking ZeroSSL to send a validation email, if empty will user imap-username | [optional] 
**ImapUsername** | **string** | ImapUsername to access the IMAP service | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UpdateVersion** | Pointer to **bool** | Deprecated | [optional] 

## Methods

### NewUpdateZeroSSLTarget

`func NewUpdateZeroSSLTarget(apiKey string, imapFqdn string, imapPassword string, imapUsername string, name string, ) *UpdateZeroSSLTarget`

NewUpdateZeroSSLTarget instantiates a new UpdateZeroSSLTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateZeroSSLTargetWithDefaults

`func NewUpdateZeroSSLTargetWithDefaults() *UpdateZeroSSLTarget`

NewUpdateZeroSSLTargetWithDefaults instantiates a new UpdateZeroSSLTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *UpdateZeroSSLTarget) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *UpdateZeroSSLTarget) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *UpdateZeroSSLTarget) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetComment

`func (o *UpdateZeroSSLTarget) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *UpdateZeroSSLTarget) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *UpdateZeroSSLTarget) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *UpdateZeroSSLTarget) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateZeroSSLTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateZeroSSLTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateZeroSSLTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateZeroSSLTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImapFqdn

`func (o *UpdateZeroSSLTarget) GetImapFqdn() string`

GetImapFqdn returns the ImapFqdn field if non-nil, zero value otherwise.

### GetImapFqdnOk

`func (o *UpdateZeroSSLTarget) GetImapFqdnOk() (*string, bool)`

GetImapFqdnOk returns a tuple with the ImapFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapFqdn

`func (o *UpdateZeroSSLTarget) SetImapFqdn(v string)`

SetImapFqdn sets ImapFqdn field to given value.


### GetImapPassword

`func (o *UpdateZeroSSLTarget) GetImapPassword() string`

GetImapPassword returns the ImapPassword field if non-nil, zero value otherwise.

### GetImapPasswordOk

`func (o *UpdateZeroSSLTarget) GetImapPasswordOk() (*string, bool)`

GetImapPasswordOk returns a tuple with the ImapPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapPassword

`func (o *UpdateZeroSSLTarget) SetImapPassword(v string)`

SetImapPassword sets ImapPassword field to given value.


### GetImapPort

`func (o *UpdateZeroSSLTarget) GetImapPort() string`

GetImapPort returns the ImapPort field if non-nil, zero value otherwise.

### GetImapPortOk

`func (o *UpdateZeroSSLTarget) GetImapPortOk() (*string, bool)`

GetImapPortOk returns a tuple with the ImapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapPort

`func (o *UpdateZeroSSLTarget) SetImapPort(v string)`

SetImapPort sets ImapPort field to given value.

### HasImapPort

`func (o *UpdateZeroSSLTarget) HasImapPort() bool`

HasImapPort returns a boolean if a field has been set.

### GetImapTargetEmail

`func (o *UpdateZeroSSLTarget) GetImapTargetEmail() string`

GetImapTargetEmail returns the ImapTargetEmail field if non-nil, zero value otherwise.

### GetImapTargetEmailOk

`func (o *UpdateZeroSSLTarget) GetImapTargetEmailOk() (*string, bool)`

GetImapTargetEmailOk returns a tuple with the ImapTargetEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapTargetEmail

`func (o *UpdateZeroSSLTarget) SetImapTargetEmail(v string)`

SetImapTargetEmail sets ImapTargetEmail field to given value.

### HasImapTargetEmail

`func (o *UpdateZeroSSLTarget) HasImapTargetEmail() bool`

HasImapTargetEmail returns a boolean if a field has been set.

### GetImapUsername

`func (o *UpdateZeroSSLTarget) GetImapUsername() string`

GetImapUsername returns the ImapUsername field if non-nil, zero value otherwise.

### GetImapUsernameOk

`func (o *UpdateZeroSSLTarget) GetImapUsernameOk() (*string, bool)`

GetImapUsernameOk returns a tuple with the ImapUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapUsername

`func (o *UpdateZeroSSLTarget) SetImapUsername(v string)`

SetImapUsername sets ImapUsername field to given value.


### GetJson

`func (o *UpdateZeroSSLTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateZeroSSLTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateZeroSSLTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateZeroSSLTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *UpdateZeroSSLTarget) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *UpdateZeroSSLTarget) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *UpdateZeroSSLTarget) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *UpdateZeroSSLTarget) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *UpdateZeroSSLTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UpdateZeroSSLTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UpdateZeroSSLTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UpdateZeroSSLTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UpdateZeroSSLTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateZeroSSLTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateZeroSSLTarget) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *UpdateZeroSSLTarget) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *UpdateZeroSSLTarget) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *UpdateZeroSSLTarget) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *UpdateZeroSSLTarget) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetTimeout

`func (o *UpdateZeroSSLTarget) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UpdateZeroSSLTarget) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UpdateZeroSSLTarget) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UpdateZeroSSLTarget) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *UpdateZeroSSLTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateZeroSSLTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateZeroSSLTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateZeroSSLTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateZeroSSLTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateZeroSSLTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateZeroSSLTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateZeroSSLTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUpdateVersion

`func (o *UpdateZeroSSLTarget) GetUpdateVersion() bool`

GetUpdateVersion returns the UpdateVersion field if non-nil, zero value otherwise.

### GetUpdateVersionOk

`func (o *UpdateZeroSSLTarget) GetUpdateVersionOk() (*bool, bool)`

GetUpdateVersionOk returns a tuple with the UpdateVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateVersion

`func (o *UpdateZeroSSLTarget) SetUpdateVersion(v bool)`

SetUpdateVersion sets UpdateVersion field to given value.

### HasUpdateVersion

`func (o *UpdateZeroSSLTarget) HasUpdateVersion() bool`

HasUpdateVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# TargetUpdateGodaddy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** | Key of the api credentials to the Godaddy account | 
**CustomerId** | Pointer to **string** | Customer ID (ShopperId) required for renewal of imported certificates | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**ImapFqdn** | **string** | ImapFQDN of the IMAP service, FQDN or IPv4 address. Must be FQDN if the IMAP is using TLS | 
**ImapPassword** | **string** | ImapPassword to access the IMAP service | 
**ImapPort** | Pointer to **string** | ImapPort of the IMAP service | [optional] [default to "993"]
**ImapUsername** | **string** | ImapUsername to access the IMAP service | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeepPrevVersion** | Pointer to **string** | Whether to keep previous version [true/false]. If not set, use default according to account settings | [optional] 
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**MaxVersions** | Pointer to **string** | Set the maximum number of versions, limited by the account settings defaults. | [optional] 
**Name** | **string** | Target name | 
**NewName** | Pointer to **string** | New target name | [optional] 
**Secret** | **string** | Secret of the api credentials to the Godaddy account | 
**Timeout** | Pointer to **string** | Timeout waiting for certificate validation in Duration format (1h - 1 Hour, 20m - 20 Minutes, 33m3s - 33 Minutes and 3 Seconds), maximum 1h. | [optional] [default to "5m"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewTargetUpdateGodaddy

`func NewTargetUpdateGodaddy(apiKey string, imapFqdn string, imapPassword string, imapUsername string, name string, secret string, ) *TargetUpdateGodaddy`

NewTargetUpdateGodaddy instantiates a new TargetUpdateGodaddy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetUpdateGodaddyWithDefaults

`func NewTargetUpdateGodaddyWithDefaults() *TargetUpdateGodaddy`

NewTargetUpdateGodaddyWithDefaults instantiates a new TargetUpdateGodaddy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *TargetUpdateGodaddy) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *TargetUpdateGodaddy) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *TargetUpdateGodaddy) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetCustomerId

`func (o *TargetUpdateGodaddy) GetCustomerId() string`

GetCustomerId returns the CustomerId field if non-nil, zero value otherwise.

### GetCustomerIdOk

`func (o *TargetUpdateGodaddy) GetCustomerIdOk() (*string, bool)`

GetCustomerIdOk returns a tuple with the CustomerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerId

`func (o *TargetUpdateGodaddy) SetCustomerId(v string)`

SetCustomerId sets CustomerId field to given value.

### HasCustomerId

`func (o *TargetUpdateGodaddy) HasCustomerId() bool`

HasCustomerId returns a boolean if a field has been set.

### GetDescription

`func (o *TargetUpdateGodaddy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TargetUpdateGodaddy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TargetUpdateGodaddy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TargetUpdateGodaddy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImapFqdn

`func (o *TargetUpdateGodaddy) GetImapFqdn() string`

GetImapFqdn returns the ImapFqdn field if non-nil, zero value otherwise.

### GetImapFqdnOk

`func (o *TargetUpdateGodaddy) GetImapFqdnOk() (*string, bool)`

GetImapFqdnOk returns a tuple with the ImapFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapFqdn

`func (o *TargetUpdateGodaddy) SetImapFqdn(v string)`

SetImapFqdn sets ImapFqdn field to given value.


### GetImapPassword

`func (o *TargetUpdateGodaddy) GetImapPassword() string`

GetImapPassword returns the ImapPassword field if non-nil, zero value otherwise.

### GetImapPasswordOk

`func (o *TargetUpdateGodaddy) GetImapPasswordOk() (*string, bool)`

GetImapPasswordOk returns a tuple with the ImapPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapPassword

`func (o *TargetUpdateGodaddy) SetImapPassword(v string)`

SetImapPassword sets ImapPassword field to given value.


### GetImapPort

`func (o *TargetUpdateGodaddy) GetImapPort() string`

GetImapPort returns the ImapPort field if non-nil, zero value otherwise.

### GetImapPortOk

`func (o *TargetUpdateGodaddy) GetImapPortOk() (*string, bool)`

GetImapPortOk returns a tuple with the ImapPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapPort

`func (o *TargetUpdateGodaddy) SetImapPort(v string)`

SetImapPort sets ImapPort field to given value.

### HasImapPort

`func (o *TargetUpdateGodaddy) HasImapPort() bool`

HasImapPort returns a boolean if a field has been set.

### GetImapUsername

`func (o *TargetUpdateGodaddy) GetImapUsername() string`

GetImapUsername returns the ImapUsername field if non-nil, zero value otherwise.

### GetImapUsernameOk

`func (o *TargetUpdateGodaddy) GetImapUsernameOk() (*string, bool)`

GetImapUsernameOk returns a tuple with the ImapUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImapUsername

`func (o *TargetUpdateGodaddy) SetImapUsername(v string)`

SetImapUsername sets ImapUsername field to given value.


### GetJson

`func (o *TargetUpdateGodaddy) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *TargetUpdateGodaddy) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *TargetUpdateGodaddy) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *TargetUpdateGodaddy) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeepPrevVersion

`func (o *TargetUpdateGodaddy) GetKeepPrevVersion() string`

GetKeepPrevVersion returns the KeepPrevVersion field if non-nil, zero value otherwise.

### GetKeepPrevVersionOk

`func (o *TargetUpdateGodaddy) GetKeepPrevVersionOk() (*string, bool)`

GetKeepPrevVersionOk returns a tuple with the KeepPrevVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepPrevVersion

`func (o *TargetUpdateGodaddy) SetKeepPrevVersion(v string)`

SetKeepPrevVersion sets KeepPrevVersion field to given value.

### HasKeepPrevVersion

`func (o *TargetUpdateGodaddy) HasKeepPrevVersion() bool`

HasKeepPrevVersion returns a boolean if a field has been set.

### GetKey

`func (o *TargetUpdateGodaddy) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *TargetUpdateGodaddy) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *TargetUpdateGodaddy) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *TargetUpdateGodaddy) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetMaxVersions

`func (o *TargetUpdateGodaddy) GetMaxVersions() string`

GetMaxVersions returns the MaxVersions field if non-nil, zero value otherwise.

### GetMaxVersionsOk

`func (o *TargetUpdateGodaddy) GetMaxVersionsOk() (*string, bool)`

GetMaxVersionsOk returns a tuple with the MaxVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxVersions

`func (o *TargetUpdateGodaddy) SetMaxVersions(v string)`

SetMaxVersions sets MaxVersions field to given value.

### HasMaxVersions

`func (o *TargetUpdateGodaddy) HasMaxVersions() bool`

HasMaxVersions returns a boolean if a field has been set.

### GetName

`func (o *TargetUpdateGodaddy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TargetUpdateGodaddy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TargetUpdateGodaddy) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *TargetUpdateGodaddy) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *TargetUpdateGodaddy) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *TargetUpdateGodaddy) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *TargetUpdateGodaddy) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetSecret

`func (o *TargetUpdateGodaddy) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TargetUpdateGodaddy) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TargetUpdateGodaddy) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetTimeout

`func (o *TargetUpdateGodaddy) GetTimeout() string`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TargetUpdateGodaddy) GetTimeoutOk() (*string, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TargetUpdateGodaddy) SetTimeout(v string)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TargetUpdateGodaddy) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetToken

`func (o *TargetUpdateGodaddy) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *TargetUpdateGodaddy) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *TargetUpdateGodaddy) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *TargetUpdateGodaddy) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *TargetUpdateGodaddy) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *TargetUpdateGodaddy) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *TargetUpdateGodaddy) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *TargetUpdateGodaddy) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# GatewayUpdateRemoteAccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedSshUrl** | Pointer to **string** | Specify a valid SSH-URL to tunnel to SSH session | [optional] [default to "use-existing"]
**AllowedUrls** | Pointer to **string** | List of valid URLs to redirect from the Portal back to the remote access server (in a comma-delimited list) | [optional] [default to "use-existing"]
**DefaultSessionTtlMinutes** | Pointer to **string** | Default session TTL in minutes | [optional] [default to "use-existing"]
**HideSessionRecording** | Pointer to **string** | Specifies whether to show/hide if the session is currently recorded [true/false] | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Kexalgs** | Pointer to **string** | Decide which algorithm will be used as part of the SSH initial hand-shake process | [optional] [default to "use-existing"]
**KeyboardLayout** | Pointer to **string** | Enable support for additional keyboard layouts | [optional] [default to "use-existing"]
**LegacySshAlgorithm** | Pointer to **string** | Signs SSH certificates using legacy ssh-rsa-cert-01@openssh.com signing algorithm [true/false] | [optional] 
**RdpTargetConfiguration** | Pointer to **string** | Specify the usernameSubClaim that exists inside the IDP JWT, e.g. email | [optional] [default to "use-existing"]
**SshTargetConfiguration** | Pointer to **string** | Specify the usernameSubClaim that exists inside the IDP JWT, e.g. email | [optional] [default to "use-existing"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGatewayUpdateRemoteAccess

`func NewGatewayUpdateRemoteAccess() *GatewayUpdateRemoteAccess`

NewGatewayUpdateRemoteAccess instantiates a new GatewayUpdateRemoteAccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateRemoteAccessWithDefaults

`func NewGatewayUpdateRemoteAccessWithDefaults() *GatewayUpdateRemoteAccess`

NewGatewayUpdateRemoteAccessWithDefaults instantiates a new GatewayUpdateRemoteAccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedSshUrl

`func (o *GatewayUpdateRemoteAccess) GetAllowedSshUrl() string`

GetAllowedSshUrl returns the AllowedSshUrl field if non-nil, zero value otherwise.

### GetAllowedSshUrlOk

`func (o *GatewayUpdateRemoteAccess) GetAllowedSshUrlOk() (*string, bool)`

GetAllowedSshUrlOk returns a tuple with the AllowedSshUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedSshUrl

`func (o *GatewayUpdateRemoteAccess) SetAllowedSshUrl(v string)`

SetAllowedSshUrl sets AllowedSshUrl field to given value.

### HasAllowedSshUrl

`func (o *GatewayUpdateRemoteAccess) HasAllowedSshUrl() bool`

HasAllowedSshUrl returns a boolean if a field has been set.

### GetAllowedUrls

`func (o *GatewayUpdateRemoteAccess) GetAllowedUrls() string`

GetAllowedUrls returns the AllowedUrls field if non-nil, zero value otherwise.

### GetAllowedUrlsOk

`func (o *GatewayUpdateRemoteAccess) GetAllowedUrlsOk() (*string, bool)`

GetAllowedUrlsOk returns a tuple with the AllowedUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedUrls

`func (o *GatewayUpdateRemoteAccess) SetAllowedUrls(v string)`

SetAllowedUrls sets AllowedUrls field to given value.

### HasAllowedUrls

`func (o *GatewayUpdateRemoteAccess) HasAllowedUrls() bool`

HasAllowedUrls returns a boolean if a field has been set.

### GetDefaultSessionTtlMinutes

`func (o *GatewayUpdateRemoteAccess) GetDefaultSessionTtlMinutes() string`

GetDefaultSessionTtlMinutes returns the DefaultSessionTtlMinutes field if non-nil, zero value otherwise.

### GetDefaultSessionTtlMinutesOk

`func (o *GatewayUpdateRemoteAccess) GetDefaultSessionTtlMinutesOk() (*string, bool)`

GetDefaultSessionTtlMinutesOk returns a tuple with the DefaultSessionTtlMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSessionTtlMinutes

`func (o *GatewayUpdateRemoteAccess) SetDefaultSessionTtlMinutes(v string)`

SetDefaultSessionTtlMinutes sets DefaultSessionTtlMinutes field to given value.

### HasDefaultSessionTtlMinutes

`func (o *GatewayUpdateRemoteAccess) HasDefaultSessionTtlMinutes() bool`

HasDefaultSessionTtlMinutes returns a boolean if a field has been set.

### GetHideSessionRecording

`func (o *GatewayUpdateRemoteAccess) GetHideSessionRecording() string`

GetHideSessionRecording returns the HideSessionRecording field if non-nil, zero value otherwise.

### GetHideSessionRecordingOk

`func (o *GatewayUpdateRemoteAccess) GetHideSessionRecordingOk() (*string, bool)`

GetHideSessionRecordingOk returns a tuple with the HideSessionRecording field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideSessionRecording

`func (o *GatewayUpdateRemoteAccess) SetHideSessionRecording(v string)`

SetHideSessionRecording sets HideSessionRecording field to given value.

### HasHideSessionRecording

`func (o *GatewayUpdateRemoteAccess) HasHideSessionRecording() bool`

HasHideSessionRecording returns a boolean if a field has been set.

### GetJson

`func (o *GatewayUpdateRemoteAccess) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayUpdateRemoteAccess) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayUpdateRemoteAccess) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayUpdateRemoteAccess) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKexalgs

`func (o *GatewayUpdateRemoteAccess) GetKexalgs() string`

GetKexalgs returns the Kexalgs field if non-nil, zero value otherwise.

### GetKexalgsOk

`func (o *GatewayUpdateRemoteAccess) GetKexalgsOk() (*string, bool)`

GetKexalgsOk returns a tuple with the Kexalgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKexalgs

`func (o *GatewayUpdateRemoteAccess) SetKexalgs(v string)`

SetKexalgs sets Kexalgs field to given value.

### HasKexalgs

`func (o *GatewayUpdateRemoteAccess) HasKexalgs() bool`

HasKexalgs returns a boolean if a field has been set.

### GetKeyboardLayout

`func (o *GatewayUpdateRemoteAccess) GetKeyboardLayout() string`

GetKeyboardLayout returns the KeyboardLayout field if non-nil, zero value otherwise.

### GetKeyboardLayoutOk

`func (o *GatewayUpdateRemoteAccess) GetKeyboardLayoutOk() (*string, bool)`

GetKeyboardLayoutOk returns a tuple with the KeyboardLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyboardLayout

`func (o *GatewayUpdateRemoteAccess) SetKeyboardLayout(v string)`

SetKeyboardLayout sets KeyboardLayout field to given value.

### HasKeyboardLayout

`func (o *GatewayUpdateRemoteAccess) HasKeyboardLayout() bool`

HasKeyboardLayout returns a boolean if a field has been set.

### GetLegacySshAlgorithm

`func (o *GatewayUpdateRemoteAccess) GetLegacySshAlgorithm() string`

GetLegacySshAlgorithm returns the LegacySshAlgorithm field if non-nil, zero value otherwise.

### GetLegacySshAlgorithmOk

`func (o *GatewayUpdateRemoteAccess) GetLegacySshAlgorithmOk() (*string, bool)`

GetLegacySshAlgorithmOk returns a tuple with the LegacySshAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacySshAlgorithm

`func (o *GatewayUpdateRemoteAccess) SetLegacySshAlgorithm(v string)`

SetLegacySshAlgorithm sets LegacySshAlgorithm field to given value.

### HasLegacySshAlgorithm

`func (o *GatewayUpdateRemoteAccess) HasLegacySshAlgorithm() bool`

HasLegacySshAlgorithm returns a boolean if a field has been set.

### GetRdpTargetConfiguration

`func (o *GatewayUpdateRemoteAccess) GetRdpTargetConfiguration() string`

GetRdpTargetConfiguration returns the RdpTargetConfiguration field if non-nil, zero value otherwise.

### GetRdpTargetConfigurationOk

`func (o *GatewayUpdateRemoteAccess) GetRdpTargetConfigurationOk() (*string, bool)`

GetRdpTargetConfigurationOk returns a tuple with the RdpTargetConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpTargetConfiguration

`func (o *GatewayUpdateRemoteAccess) SetRdpTargetConfiguration(v string)`

SetRdpTargetConfiguration sets RdpTargetConfiguration field to given value.

### HasRdpTargetConfiguration

`func (o *GatewayUpdateRemoteAccess) HasRdpTargetConfiguration() bool`

HasRdpTargetConfiguration returns a boolean if a field has been set.

### GetSshTargetConfiguration

`func (o *GatewayUpdateRemoteAccess) GetSshTargetConfiguration() string`

GetSshTargetConfiguration returns the SshTargetConfiguration field if non-nil, zero value otherwise.

### GetSshTargetConfigurationOk

`func (o *GatewayUpdateRemoteAccess) GetSshTargetConfigurationOk() (*string, bool)`

GetSshTargetConfigurationOk returns a tuple with the SshTargetConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshTargetConfiguration

`func (o *GatewayUpdateRemoteAccess) SetSshTargetConfiguration(v string)`

SetSshTargetConfiguration sets SshTargetConfiguration field to given value.

### HasSshTargetConfiguration

`func (o *GatewayUpdateRemoteAccess) HasSshTargetConfiguration() bool`

HasSshTargetConfiguration returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateRemoteAccess) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateRemoteAccess) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateRemoteAccess) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateRemoteAccess) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateRemoteAccess) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateRemoteAccess) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateRemoteAccess) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateRemoteAccess) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



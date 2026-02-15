# SplunkTargetDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **string** | Token audience | [optional] 
**AuthMode** | Pointer to **string** | Authentication mode: \&quot;username\&quot; or \&quot;token\&quot; | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SplunkPayload** | Pointer to [**SplunkPayload**](SplunkPayload.md) |  | [optional] 
**SplunkUrl** | Pointer to **string** | Splunk server URL | [optional] 
**Token** | Pointer to **string** | Token is used when AuthMode &#x3D;&#x3D; \&quot;token\&quot; | [optional] 
**TokenOwner** | Pointer to **string** | Token owner (the Splunk user who owns the token, required for token rotation) | [optional] 
**UseTls** | Pointer to **bool** | Use TLS certificate verification when connecting to the Splunk management API. | [optional] 
**Username** | Pointer to **string** | Username &amp; Password are used when AuthMode &#x3D;&#x3D; \&quot;username\&quot; | [optional] 

## Methods

### NewSplunkTargetDetails

`func NewSplunkTargetDetails() *SplunkTargetDetails`

NewSplunkTargetDetails instantiates a new SplunkTargetDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSplunkTargetDetailsWithDefaults

`func NewSplunkTargetDetailsWithDefaults() *SplunkTargetDetails`

NewSplunkTargetDetailsWithDefaults instantiates a new SplunkTargetDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *SplunkTargetDetails) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SplunkTargetDetails) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SplunkTargetDetails) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SplunkTargetDetails) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### GetAuthMode

`func (o *SplunkTargetDetails) GetAuthMode() string`

GetAuthMode returns the AuthMode field if non-nil, zero value otherwise.

### GetAuthModeOk

`func (o *SplunkTargetDetails) GetAuthModeOk() (*string, bool)`

GetAuthModeOk returns a tuple with the AuthMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthMode

`func (o *SplunkTargetDetails) SetAuthMode(v string)`

SetAuthMode sets AuthMode field to given value.

### HasAuthMode

`func (o *SplunkTargetDetails) HasAuthMode() bool`

HasAuthMode returns a boolean if a field has been set.

### GetPassword

`func (o *SplunkTargetDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *SplunkTargetDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *SplunkTargetDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *SplunkTargetDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSplunkPayload

`func (o *SplunkTargetDetails) GetSplunkPayload() SplunkPayload`

GetSplunkPayload returns the SplunkPayload field if non-nil, zero value otherwise.

### GetSplunkPayloadOk

`func (o *SplunkTargetDetails) GetSplunkPayloadOk() (*SplunkPayload, bool)`

GetSplunkPayloadOk returns a tuple with the SplunkPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkPayload

`func (o *SplunkTargetDetails) SetSplunkPayload(v SplunkPayload)`

SetSplunkPayload sets SplunkPayload field to given value.

### HasSplunkPayload

`func (o *SplunkTargetDetails) HasSplunkPayload() bool`

HasSplunkPayload returns a boolean if a field has been set.

### GetSplunkUrl

`func (o *SplunkTargetDetails) GetSplunkUrl() string`

GetSplunkUrl returns the SplunkUrl field if non-nil, zero value otherwise.

### GetSplunkUrlOk

`func (o *SplunkTargetDetails) GetSplunkUrlOk() (*string, bool)`

GetSplunkUrlOk returns a tuple with the SplunkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplunkUrl

`func (o *SplunkTargetDetails) SetSplunkUrl(v string)`

SetSplunkUrl sets SplunkUrl field to given value.

### HasSplunkUrl

`func (o *SplunkTargetDetails) HasSplunkUrl() bool`

HasSplunkUrl returns a boolean if a field has been set.

### GetToken

`func (o *SplunkTargetDetails) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *SplunkTargetDetails) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *SplunkTargetDetails) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *SplunkTargetDetails) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenOwner

`func (o *SplunkTargetDetails) GetTokenOwner() string`

GetTokenOwner returns the TokenOwner field if non-nil, zero value otherwise.

### GetTokenOwnerOk

`func (o *SplunkTargetDetails) GetTokenOwnerOk() (*string, bool)`

GetTokenOwnerOk returns a tuple with the TokenOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenOwner

`func (o *SplunkTargetDetails) SetTokenOwner(v string)`

SetTokenOwner sets TokenOwner field to given value.

### HasTokenOwner

`func (o *SplunkTargetDetails) HasTokenOwner() bool`

HasTokenOwner returns a boolean if a field has been set.

### GetUseTls

`func (o *SplunkTargetDetails) GetUseTls() bool`

GetUseTls returns the UseTls field if non-nil, zero value otherwise.

### GetUseTlsOk

`func (o *SplunkTargetDetails) GetUseTlsOk() (*bool, bool)`

GetUseTlsOk returns a tuple with the UseTls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseTls

`func (o *SplunkTargetDetails) SetUseTls(v bool)`

SetUseTls sets UseTls field to given value.

### HasUseTls

`func (o *SplunkTargetDetails) HasUseTls() bool`

HasUseTls returns a boolean if a field has been set.

### GetUsername

`func (o *SplunkTargetDetails) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SplunkTargetDetails) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SplunkTargetDetails) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SplunkTargetDetails) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



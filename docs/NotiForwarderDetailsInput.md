# NotiForwarderDetailsInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppPrivateKeyPemBase64** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** | Auth - JWT | [optional] 
**Password** | Pointer to **string** | Auth - User Password | [optional] 
**SlackNotiForwarderDetails** | Pointer to [**SlackNotiForwarderDetails**](SlackNotiForwarderDetails.md) |  | [optional] 
**TeamsNotiForwarderDetails** | Pointer to [**TeamsNotiForwarderDetails**](TeamsNotiForwarderDetails.md) |  | [optional] 
**WebhookNotiForwarderDetails** | Pointer to [**WebhookNotiForwarderDetails**](WebhookNotiForwarderDetails.md) |  | [optional] 

## Methods

### NewNotiForwarderDetailsInput

`func NewNotiForwarderDetailsInput() *NotiForwarderDetailsInput`

NewNotiForwarderDetailsInput instantiates a new NotiForwarderDetailsInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotiForwarderDetailsInputWithDefaults

`func NewNotiForwarderDetailsInputWithDefaults() *NotiForwarderDetailsInput`

NewNotiForwarderDetailsInputWithDefaults instantiates a new NotiForwarderDetailsInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppPrivateKeyPemBase64

`func (o *NotiForwarderDetailsInput) GetAppPrivateKeyPemBase64() string`

GetAppPrivateKeyPemBase64 returns the AppPrivateKeyPemBase64 field if non-nil, zero value otherwise.

### GetAppPrivateKeyPemBase64Ok

`func (o *NotiForwarderDetailsInput) GetAppPrivateKeyPemBase64Ok() (*string, bool)`

GetAppPrivateKeyPemBase64Ok returns a tuple with the AppPrivateKeyPemBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPrivateKeyPemBase64

`func (o *NotiForwarderDetailsInput) SetAppPrivateKeyPemBase64(v string)`

SetAppPrivateKeyPemBase64 sets AppPrivateKeyPemBase64 field to given value.

### HasAppPrivateKeyPemBase64

`func (o *NotiForwarderDetailsInput) HasAppPrivateKeyPemBase64() bool`

HasAppPrivateKeyPemBase64 returns a boolean if a field has been set.

### GetClientSecret

`func (o *NotiForwarderDetailsInput) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *NotiForwarderDetailsInput) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *NotiForwarderDetailsInput) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *NotiForwarderDetailsInput) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetPassword

`func (o *NotiForwarderDetailsInput) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotiForwarderDetailsInput) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotiForwarderDetailsInput) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotiForwarderDetailsInput) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSlackNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) GetSlackNotiForwarderDetails() SlackNotiForwarderDetails`

GetSlackNotiForwarderDetails returns the SlackNotiForwarderDetails field if non-nil, zero value otherwise.

### GetSlackNotiForwarderDetailsOk

`func (o *NotiForwarderDetailsInput) GetSlackNotiForwarderDetailsOk() (*SlackNotiForwarderDetails, bool)`

GetSlackNotiForwarderDetailsOk returns a tuple with the SlackNotiForwarderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlackNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) SetSlackNotiForwarderDetails(v SlackNotiForwarderDetails)`

SetSlackNotiForwarderDetails sets SlackNotiForwarderDetails field to given value.

### HasSlackNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) HasSlackNotiForwarderDetails() bool`

HasSlackNotiForwarderDetails returns a boolean if a field has been set.

### GetTeamsNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) GetTeamsNotiForwarderDetails() TeamsNotiForwarderDetails`

GetTeamsNotiForwarderDetails returns the TeamsNotiForwarderDetails field if non-nil, zero value otherwise.

### GetTeamsNotiForwarderDetailsOk

`func (o *NotiForwarderDetailsInput) GetTeamsNotiForwarderDetailsOk() (*TeamsNotiForwarderDetails, bool)`

GetTeamsNotiForwarderDetailsOk returns a tuple with the TeamsNotiForwarderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamsNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) SetTeamsNotiForwarderDetails(v TeamsNotiForwarderDetails)`

SetTeamsNotiForwarderDetails sets TeamsNotiForwarderDetails field to given value.

### HasTeamsNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) HasTeamsNotiForwarderDetails() bool`

HasTeamsNotiForwarderDetails returns a boolean if a field has been set.

### GetWebhookNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) GetWebhookNotiForwarderDetails() WebhookNotiForwarderDetails`

GetWebhookNotiForwarderDetails returns the WebhookNotiForwarderDetails field if non-nil, zero value otherwise.

### GetWebhookNotiForwarderDetailsOk

`func (o *NotiForwarderDetailsInput) GetWebhookNotiForwarderDetailsOk() (*WebhookNotiForwarderDetails, bool)`

GetWebhookNotiForwarderDetailsOk returns a tuple with the WebhookNotiForwarderDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) SetWebhookNotiForwarderDetails(v WebhookNotiForwarderDetails)`

SetWebhookNotiForwarderDetails sets WebhookNotiForwarderDetails field to given value.

### HasWebhookNotiForwarderDetails

`func (o *NotiForwarderDetailsInput) HasWebhookNotiForwarderDetails() bool`

HasWebhookNotiForwarderDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



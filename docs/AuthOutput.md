# AuthOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Creds** | Pointer to [**SystemAccessCredentialsReplyObj**](SystemAccessCredentialsReplyObj.md) |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 

## Methods

### NewAuthOutput

`func NewAuthOutput() *AuthOutput`

NewAuthOutput instantiates a new AuthOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthOutputWithDefaults

`func NewAuthOutputWithDefaults() *AuthOutput`

NewAuthOutputWithDefaults instantiates a new AuthOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreds

`func (o *AuthOutput) GetCreds() SystemAccessCredentialsReplyObj`

GetCreds returns the Creds field if non-nil, zero value otherwise.

### GetCredsOk

`func (o *AuthOutput) GetCredsOk() (*SystemAccessCredentialsReplyObj, bool)`

GetCredsOk returns a tuple with the Creds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreds

`func (o *AuthOutput) SetCreds(v SystemAccessCredentialsReplyObj)`

SetCreds sets Creds field to given value.

### HasCreds

`func (o *AuthOutput) HasCreds() bool`

HasCreds returns a boolean if a field has been set.

### GetToken

`func (o *AuthOutput) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AuthOutput) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AuthOutput) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AuthOutput) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



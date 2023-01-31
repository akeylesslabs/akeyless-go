# OnePasswordPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**SecretKey** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Vaults** | Pointer to **[]string** |  | [optional] 

## Methods

### NewOnePasswordPayload

`func NewOnePasswordPayload() *OnePasswordPayload`

NewOnePasswordPayload instantiates a new OnePasswordPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOnePasswordPayloadWithDefaults

`func NewOnePasswordPayloadWithDefaults() *OnePasswordPayload`

NewOnePasswordPayloadWithDefaults instantiates a new OnePasswordPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OnePasswordPayload) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OnePasswordPayload) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OnePasswordPayload) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *OnePasswordPayload) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *OnePasswordPayload) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OnePasswordPayload) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OnePasswordPayload) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *OnePasswordPayload) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSecretKey

`func (o *OnePasswordPayload) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *OnePasswordPayload) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *OnePasswordPayload) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *OnePasswordPayload) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUrl

`func (o *OnePasswordPayload) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *OnePasswordPayload) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *OnePasswordPayload) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *OnePasswordPayload) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetVaults

`func (o *OnePasswordPayload) GetVaults() []string`

GetVaults returns the Vaults field if non-nil, zero value otherwise.

### GetVaultsOk

`func (o *OnePasswordPayload) GetVaultsOk() (*[]string, bool)`

GetVaultsOk returns a tuple with the Vaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVaults

`func (o *OnePasswordPayload) SetVaults(v []string)`

SetVaults sets Vaults field to given value.

### HasVaults

`func (o *OnePasswordPayload) HasVaults() bool`

HasVaults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# DecryptFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewDecryptFile

`func NewDecryptFile() *DecryptFile`

NewDecryptFile instantiates a new DecryptFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecryptFileWithDefaults

`func NewDecryptFileWithDefaults() *DecryptFile`

NewDecryptFileWithDefaults instantiates a new DecryptFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *DecryptFile) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *DecryptFile) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *DecryptFile) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *DecryptFile) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *DecryptFile) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DecryptFile) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DecryptFile) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DecryptFile) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DecryptFile) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DecryptFile) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DecryptFile) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DecryptFile) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *DecryptFile) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *DecryptFile) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *DecryptFile) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *DecryptFile) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



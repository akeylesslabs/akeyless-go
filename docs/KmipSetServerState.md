# KmipSetServerState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**State** | **string** |  | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewKmipSetServerState

`func NewKmipSetServerState(state string, ) *KmipSetServerState`

NewKmipSetServerState instantiates a new KmipSetServerState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKmipSetServerStateWithDefaults

`func NewKmipSetServerStateWithDefaults() *KmipSetServerState`

NewKmipSetServerStateWithDefaults instantiates a new KmipSetServerState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *KmipSetServerState) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KmipSetServerState) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KmipSetServerState) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KmipSetServerState) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetState

`func (o *KmipSetServerState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KmipSetServerState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KmipSetServerState) SetState(v string)`

SetState sets State field to given value.


### GetToken

`func (o *KmipSetServerState) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *KmipSetServerState) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *KmipSetServerState) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *KmipSetServerState) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *KmipSetServerState) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *KmipSetServerState) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *KmipSetServerState) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *KmipSetServerState) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *KmipSetServerState) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *KmipSetServerState) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *KmipSetServerState) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *KmipSetServerState) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



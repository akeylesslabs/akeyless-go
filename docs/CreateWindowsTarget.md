# CreateWindowsTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the object | [optional] 
**Hostname** | Pointer to **string** | Server hostname | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Key** | Pointer to **string** | The name of a key that used to encrypt the target secret value (if empty, the account default protectionKey key will be used) | [optional] 
**Name** | **string** | Target name | 
**Password** | Pointer to **string** | The privileged user password | [optional] 
**Port** | Pointer to **string** | Server WinRM HTTPS port | [optional] [default to "5986"]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Privileged username | [optional] 

## Methods

### NewCreateWindowsTarget

`func NewCreateWindowsTarget(name string, ) *CreateWindowsTarget`

NewCreateWindowsTarget instantiates a new CreateWindowsTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWindowsTargetWithDefaults

`func NewCreateWindowsTargetWithDefaults() *CreateWindowsTarget`

NewCreateWindowsTargetWithDefaults instantiates a new CreateWindowsTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateWindowsTarget) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWindowsTarget) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWindowsTarget) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWindowsTarget) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHostname

`func (o *CreateWindowsTarget) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *CreateWindowsTarget) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *CreateWindowsTarget) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *CreateWindowsTarget) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetJson

`func (o *CreateWindowsTarget) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateWindowsTarget) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateWindowsTarget) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateWindowsTarget) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKey

`func (o *CreateWindowsTarget) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateWindowsTarget) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateWindowsTarget) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *CreateWindowsTarget) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *CreateWindowsTarget) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWindowsTarget) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWindowsTarget) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *CreateWindowsTarget) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateWindowsTarget) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateWindowsTarget) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateWindowsTarget) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPort

`func (o *CreateWindowsTarget) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CreateWindowsTarget) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CreateWindowsTarget) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *CreateWindowsTarget) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetToken

`func (o *CreateWindowsTarget) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateWindowsTarget) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateWindowsTarget) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateWindowsTarget) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateWindowsTarget) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateWindowsTarget) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateWindowsTarget) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateWindowsTarget) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *CreateWindowsTarget) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CreateWindowsTarget) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CreateWindowsTarget) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CreateWindowsTarget) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



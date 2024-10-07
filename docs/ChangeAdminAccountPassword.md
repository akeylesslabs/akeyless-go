# ChangeAdminAccountPassword

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPassword** | **string** | Current password | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**NewPassword** | **string** | New password | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewChangeAdminAccountPassword

`func NewChangeAdminAccountPassword(currentPassword string, newPassword string, ) *ChangeAdminAccountPassword`

NewChangeAdminAccountPassword instantiates a new ChangeAdminAccountPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChangeAdminAccountPasswordWithDefaults

`func NewChangeAdminAccountPasswordWithDefaults() *ChangeAdminAccountPassword`

NewChangeAdminAccountPasswordWithDefaults instantiates a new ChangeAdminAccountPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPassword

`func (o *ChangeAdminAccountPassword) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *ChangeAdminAccountPassword) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *ChangeAdminAccountPassword) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.


### GetJson

`func (o *ChangeAdminAccountPassword) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *ChangeAdminAccountPassword) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *ChangeAdminAccountPassword) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *ChangeAdminAccountPassword) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNewPassword

`func (o *ChangeAdminAccountPassword) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ChangeAdminAccountPassword) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ChangeAdminAccountPassword) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetToken

`func (o *ChangeAdminAccountPassword) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ChangeAdminAccountPassword) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ChangeAdminAccountPassword) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ChangeAdminAccountPassword) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *ChangeAdminAccountPassword) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *ChangeAdminAccountPassword) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *ChangeAdminAccountPassword) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *ChangeAdminAccountPassword) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



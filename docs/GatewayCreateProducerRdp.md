# GatewayCreateProducerRdp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FixedUserOnly** | Pointer to **string** | Fixed user | [optional] [default to "false"]
**Name** | **string** | Producer name | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RdpAdminName** | **string** | RDP Admin Name | 
**RdpAdminPwd** | **string** | RDP Admin password | 
**RdpHostName** | **string** | Hostname | 
**RdpHostPort** | Pointer to **string** | Port | [optional] [default to "22"]
**RdpUserGroups** | **string** | Groups | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewGatewayCreateProducerRdp

`func NewGatewayCreateProducerRdp(name string, rdpAdminName string, rdpAdminPwd string, rdpHostName string, rdpUserGroups string, ) *GatewayCreateProducerRdp`

NewGatewayCreateProducerRdp instantiates a new GatewayCreateProducerRdp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerRdpWithDefaults

`func NewGatewayCreateProducerRdpWithDefaults() *GatewayCreateProducerRdp`

NewGatewayCreateProducerRdpWithDefaults instantiates a new GatewayCreateProducerRdp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFixedUserOnly

`func (o *GatewayCreateProducerRdp) GetFixedUserOnly() string`

GetFixedUserOnly returns the FixedUserOnly field if non-nil, zero value otherwise.

### GetFixedUserOnlyOk

`func (o *GatewayCreateProducerRdp) GetFixedUserOnlyOk() (*string, bool)`

GetFixedUserOnlyOk returns a tuple with the FixedUserOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedUserOnly

`func (o *GatewayCreateProducerRdp) SetFixedUserOnly(v string)`

SetFixedUserOnly sets FixedUserOnly field to given value.

### HasFixedUserOnly

`func (o *GatewayCreateProducerRdp) HasFixedUserOnly() bool`

HasFixedUserOnly returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerRdp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerRdp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerRdp) SetName(v string)`

SetName sets Name field to given value.


### GetPassword

`func (o *GatewayCreateProducerRdp) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GatewayCreateProducerRdp) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GatewayCreateProducerRdp) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GatewayCreateProducerRdp) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerRdp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerRdp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerRdp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerRdp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRdpAdminName

`func (o *GatewayCreateProducerRdp) GetRdpAdminName() string`

GetRdpAdminName returns the RdpAdminName field if non-nil, zero value otherwise.

### GetRdpAdminNameOk

`func (o *GatewayCreateProducerRdp) GetRdpAdminNameOk() (*string, bool)`

GetRdpAdminNameOk returns a tuple with the RdpAdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminName

`func (o *GatewayCreateProducerRdp) SetRdpAdminName(v string)`

SetRdpAdminName sets RdpAdminName field to given value.


### GetRdpAdminPwd

`func (o *GatewayCreateProducerRdp) GetRdpAdminPwd() string`

GetRdpAdminPwd returns the RdpAdminPwd field if non-nil, zero value otherwise.

### GetRdpAdminPwdOk

`func (o *GatewayCreateProducerRdp) GetRdpAdminPwdOk() (*string, bool)`

GetRdpAdminPwdOk returns a tuple with the RdpAdminPwd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpAdminPwd

`func (o *GatewayCreateProducerRdp) SetRdpAdminPwd(v string)`

SetRdpAdminPwd sets RdpAdminPwd field to given value.


### GetRdpHostName

`func (o *GatewayCreateProducerRdp) GetRdpHostName() string`

GetRdpHostName returns the RdpHostName field if non-nil, zero value otherwise.

### GetRdpHostNameOk

`func (o *GatewayCreateProducerRdp) GetRdpHostNameOk() (*string, bool)`

GetRdpHostNameOk returns a tuple with the RdpHostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostName

`func (o *GatewayCreateProducerRdp) SetRdpHostName(v string)`

SetRdpHostName sets RdpHostName field to given value.


### GetRdpHostPort

`func (o *GatewayCreateProducerRdp) GetRdpHostPort() string`

GetRdpHostPort returns the RdpHostPort field if non-nil, zero value otherwise.

### GetRdpHostPortOk

`func (o *GatewayCreateProducerRdp) GetRdpHostPortOk() (*string, bool)`

GetRdpHostPortOk returns a tuple with the RdpHostPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpHostPort

`func (o *GatewayCreateProducerRdp) SetRdpHostPort(v string)`

SetRdpHostPort sets RdpHostPort field to given value.

### HasRdpHostPort

`func (o *GatewayCreateProducerRdp) HasRdpHostPort() bool`

HasRdpHostPort returns a boolean if a field has been set.

### GetRdpUserGroups

`func (o *GatewayCreateProducerRdp) GetRdpUserGroups() string`

GetRdpUserGroups returns the RdpUserGroups field if non-nil, zero value otherwise.

### GetRdpUserGroupsOk

`func (o *GatewayCreateProducerRdp) GetRdpUserGroupsOk() (*string, bool)`

GetRdpUserGroupsOk returns a tuple with the RdpUserGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdpUserGroups

`func (o *GatewayCreateProducerRdp) SetRdpUserGroups(v string)`

SetRdpUserGroups sets RdpUserGroups field to given value.


### GetToken

`func (o *GatewayCreateProducerRdp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerRdp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerRdp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerRdp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerRdp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerRdp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerRdp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerRdp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerRdp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerRdp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerRdp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerRdp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.

### GetUsername

`func (o *GatewayCreateProducerRdp) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GatewayCreateProducerRdp) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GatewayCreateProducerRdp) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GatewayCreateProducerRdp) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



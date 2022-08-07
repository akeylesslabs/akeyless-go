# GatewayUpdateProducerDockerhub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item | [optional] 
**DockerhubPassword** | Pointer to **string** | DockerhubPassword is either the user&#39;s password access token to manage the repository | [optional] 
**DockerhubTokenScopes** | Pointer to **string** | Access token scopes list (comma-separated) to give the dynamic secret valid options are in \&quot;repo:admin\&quot;, \&quot;repo:write\&quot;, \&quot;repo:read\&quot;, \&quot;repo:public_read\&quot; | [optional] 
**DockerhubUsername** | Pointer to **string** | DockerhubUsername is the name of the user in dockerhub | [optional] 
**Name** | **string** | Producer name | 
**NewName** | Pointer to **string** | Producer name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | List of the tags attached to this secret | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayUpdateProducerDockerhub

`func NewGatewayUpdateProducerDockerhub(name string, ) *GatewayUpdateProducerDockerhub`

NewGatewayUpdateProducerDockerhub instantiates a new GatewayUpdateProducerDockerhub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayUpdateProducerDockerhubWithDefaults

`func NewGatewayUpdateProducerDockerhubWithDefaults() *GatewayUpdateProducerDockerhub`

NewGatewayUpdateProducerDockerhubWithDefaults instantiates a new GatewayUpdateProducerDockerhub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *GatewayUpdateProducerDockerhub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayUpdateProducerDockerhub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayUpdateProducerDockerhub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayUpdateProducerDockerhub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDockerhubPassword

`func (o *GatewayUpdateProducerDockerhub) GetDockerhubPassword() string`

GetDockerhubPassword returns the DockerhubPassword field if non-nil, zero value otherwise.

### GetDockerhubPasswordOk

`func (o *GatewayUpdateProducerDockerhub) GetDockerhubPasswordOk() (*string, bool)`

GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubPassword

`func (o *GatewayUpdateProducerDockerhub) SetDockerhubPassword(v string)`

SetDockerhubPassword sets DockerhubPassword field to given value.

### HasDockerhubPassword

`func (o *GatewayUpdateProducerDockerhub) HasDockerhubPassword() bool`

HasDockerhubPassword returns a boolean if a field has been set.

### GetDockerhubTokenScopes

`func (o *GatewayUpdateProducerDockerhub) GetDockerhubTokenScopes() string`

GetDockerhubTokenScopes returns the DockerhubTokenScopes field if non-nil, zero value otherwise.

### GetDockerhubTokenScopesOk

`func (o *GatewayUpdateProducerDockerhub) GetDockerhubTokenScopesOk() (*string, bool)`

GetDockerhubTokenScopesOk returns a tuple with the DockerhubTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubTokenScopes

`func (o *GatewayUpdateProducerDockerhub) SetDockerhubTokenScopes(v string)`

SetDockerhubTokenScopes sets DockerhubTokenScopes field to given value.

### HasDockerhubTokenScopes

`func (o *GatewayUpdateProducerDockerhub) HasDockerhubTokenScopes() bool`

HasDockerhubTokenScopes returns a boolean if a field has been set.

### GetDockerhubUsername

`func (o *GatewayUpdateProducerDockerhub) GetDockerhubUsername() string`

GetDockerhubUsername returns the DockerhubUsername field if non-nil, zero value otherwise.

### GetDockerhubUsernameOk

`func (o *GatewayUpdateProducerDockerhub) GetDockerhubUsernameOk() (*string, bool)`

GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubUsername

`func (o *GatewayUpdateProducerDockerhub) SetDockerhubUsername(v string)`

SetDockerhubUsername sets DockerhubUsername field to given value.

### HasDockerhubUsername

`func (o *GatewayUpdateProducerDockerhub) HasDockerhubUsername() bool`

HasDockerhubUsername returns a boolean if a field has been set.

### GetName

`func (o *GatewayUpdateProducerDockerhub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayUpdateProducerDockerhub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayUpdateProducerDockerhub) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *GatewayUpdateProducerDockerhub) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *GatewayUpdateProducerDockerhub) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *GatewayUpdateProducerDockerhub) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *GatewayUpdateProducerDockerhub) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerDockerhub) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayUpdateProducerDockerhub) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayUpdateProducerDockerhub) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayUpdateProducerDockerhub) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *GatewayUpdateProducerDockerhub) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayUpdateProducerDockerhub) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayUpdateProducerDockerhub) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayUpdateProducerDockerhub) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayUpdateProducerDockerhub) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayUpdateProducerDockerhub) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayUpdateProducerDockerhub) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayUpdateProducerDockerhub) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayUpdateProducerDockerhub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayUpdateProducerDockerhub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayUpdateProducerDockerhub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayUpdateProducerDockerhub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayUpdateProducerDockerhub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayUpdateProducerDockerhub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayUpdateProducerDockerhub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayUpdateProducerDockerhub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayUpdateProducerDockerhub) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayUpdateProducerDockerhub) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayUpdateProducerDockerhub) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayUpdateProducerDockerhub) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



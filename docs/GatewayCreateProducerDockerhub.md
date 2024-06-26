# GatewayCreateProducerDockerhub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**DockerhubPassword** | Pointer to **string** | DockerhubPassword is either the user&#39;s password access token to manage the repository | [optional] 
**DockerhubTokenScopes** | Pointer to **string** | Access token scopes list (comma-separated) to give the dynamic secret valid options are in \&quot;repo:admin\&quot;, \&quot;repo:write\&quot;, \&quot;repo:read\&quot;, \&quot;repo:public_read\&quot; | [optional] 
**DockerhubUsername** | Pointer to **string** | DockerhubUsername is the name of the user in dockerhub | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewGatewayCreateProducerDockerhub

`func NewGatewayCreateProducerDockerhub(name string, ) *GatewayCreateProducerDockerhub`

NewGatewayCreateProducerDockerhub instantiates a new GatewayCreateProducerDockerhub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateProducerDockerhubWithDefaults

`func NewGatewayCreateProducerDockerhubWithDefaults() *GatewayCreateProducerDockerhub`

NewGatewayCreateProducerDockerhubWithDefaults instantiates a new GatewayCreateProducerDockerhub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *GatewayCreateProducerDockerhub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GatewayCreateProducerDockerhub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GatewayCreateProducerDockerhub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GatewayCreateProducerDockerhub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDockerhubPassword

`func (o *GatewayCreateProducerDockerhub) GetDockerhubPassword() string`

GetDockerhubPassword returns the DockerhubPassword field if non-nil, zero value otherwise.

### GetDockerhubPasswordOk

`func (o *GatewayCreateProducerDockerhub) GetDockerhubPasswordOk() (*string, bool)`

GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubPassword

`func (o *GatewayCreateProducerDockerhub) SetDockerhubPassword(v string)`

SetDockerhubPassword sets DockerhubPassword field to given value.

### HasDockerhubPassword

`func (o *GatewayCreateProducerDockerhub) HasDockerhubPassword() bool`

HasDockerhubPassword returns a boolean if a field has been set.

### GetDockerhubTokenScopes

`func (o *GatewayCreateProducerDockerhub) GetDockerhubTokenScopes() string`

GetDockerhubTokenScopes returns the DockerhubTokenScopes field if non-nil, zero value otherwise.

### GetDockerhubTokenScopesOk

`func (o *GatewayCreateProducerDockerhub) GetDockerhubTokenScopesOk() (*string, bool)`

GetDockerhubTokenScopesOk returns a tuple with the DockerhubTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubTokenScopes

`func (o *GatewayCreateProducerDockerhub) SetDockerhubTokenScopes(v string)`

SetDockerhubTokenScopes sets DockerhubTokenScopes field to given value.

### HasDockerhubTokenScopes

`func (o *GatewayCreateProducerDockerhub) HasDockerhubTokenScopes() bool`

HasDockerhubTokenScopes returns a boolean if a field has been set.

### GetDockerhubUsername

`func (o *GatewayCreateProducerDockerhub) GetDockerhubUsername() string`

GetDockerhubUsername returns the DockerhubUsername field if non-nil, zero value otherwise.

### GetDockerhubUsernameOk

`func (o *GatewayCreateProducerDockerhub) GetDockerhubUsernameOk() (*string, bool)`

GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubUsername

`func (o *GatewayCreateProducerDockerhub) SetDockerhubUsername(v string)`

SetDockerhubUsername sets DockerhubUsername field to given value.

### HasDockerhubUsername

`func (o *GatewayCreateProducerDockerhub) HasDockerhubUsername() bool`

HasDockerhubUsername returns a boolean if a field has been set.

### GetJson

`func (o *GatewayCreateProducerDockerhub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GatewayCreateProducerDockerhub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GatewayCreateProducerDockerhub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GatewayCreateProducerDockerhub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GatewayCreateProducerDockerhub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateProducerDockerhub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateProducerDockerhub) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *GatewayCreateProducerDockerhub) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *GatewayCreateProducerDockerhub) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *GatewayCreateProducerDockerhub) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *GatewayCreateProducerDockerhub) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *GatewayCreateProducerDockerhub) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GatewayCreateProducerDockerhub) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GatewayCreateProducerDockerhub) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GatewayCreateProducerDockerhub) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *GatewayCreateProducerDockerhub) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *GatewayCreateProducerDockerhub) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *GatewayCreateProducerDockerhub) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *GatewayCreateProducerDockerhub) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *GatewayCreateProducerDockerhub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GatewayCreateProducerDockerhub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GatewayCreateProducerDockerhub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GatewayCreateProducerDockerhub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GatewayCreateProducerDockerhub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GatewayCreateProducerDockerhub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GatewayCreateProducerDockerhub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GatewayCreateProducerDockerhub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *GatewayCreateProducerDockerhub) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *GatewayCreateProducerDockerhub) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *GatewayCreateProducerDockerhub) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *GatewayCreateProducerDockerhub) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



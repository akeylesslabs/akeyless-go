# DynamicSecretUpdateDockerhub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**DockerhubPassword** | Pointer to **string** | DockerhubPassword is either the user&#39;s password access token to manage the repository | [optional] 
**DockerhubTokenScopes** | Pointer to **string** | Access token scopes list (comma-separated) to give the dynamic secret valid options are in \&quot;repo:admin\&quot;, \&quot;repo:write\&quot;, \&quot;repo:read\&quot;, \&quot;repo:public_read\&quot; | [optional] 
**DockerhubUsername** | Pointer to **string** | DockerhubUsername is the name of the user in dockerhub | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**NewName** | Pointer to **string** | Dynamic secret new name | [optional] 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretUpdateDockerhub

`func NewDynamicSecretUpdateDockerhub(name string, ) *DynamicSecretUpdateDockerhub`

NewDynamicSecretUpdateDockerhub instantiates a new DynamicSecretUpdateDockerhub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretUpdateDockerhubWithDefaults

`func NewDynamicSecretUpdateDockerhubWithDefaults() *DynamicSecretUpdateDockerhub`

NewDynamicSecretUpdateDockerhubWithDefaults instantiates a new DynamicSecretUpdateDockerhub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretUpdateDockerhub) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretUpdateDockerhub) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretUpdateDockerhub) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretUpdateDockerhub) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDockerhubPassword

`func (o *DynamicSecretUpdateDockerhub) GetDockerhubPassword() string`

GetDockerhubPassword returns the DockerhubPassword field if non-nil, zero value otherwise.

### GetDockerhubPasswordOk

`func (o *DynamicSecretUpdateDockerhub) GetDockerhubPasswordOk() (*string, bool)`

GetDockerhubPasswordOk returns a tuple with the DockerhubPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubPassword

`func (o *DynamicSecretUpdateDockerhub) SetDockerhubPassword(v string)`

SetDockerhubPassword sets DockerhubPassword field to given value.

### HasDockerhubPassword

`func (o *DynamicSecretUpdateDockerhub) HasDockerhubPassword() bool`

HasDockerhubPassword returns a boolean if a field has been set.

### GetDockerhubTokenScopes

`func (o *DynamicSecretUpdateDockerhub) GetDockerhubTokenScopes() string`

GetDockerhubTokenScopes returns the DockerhubTokenScopes field if non-nil, zero value otherwise.

### GetDockerhubTokenScopesOk

`func (o *DynamicSecretUpdateDockerhub) GetDockerhubTokenScopesOk() (*string, bool)`

GetDockerhubTokenScopesOk returns a tuple with the DockerhubTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubTokenScopes

`func (o *DynamicSecretUpdateDockerhub) SetDockerhubTokenScopes(v string)`

SetDockerhubTokenScopes sets DockerhubTokenScopes field to given value.

### HasDockerhubTokenScopes

`func (o *DynamicSecretUpdateDockerhub) HasDockerhubTokenScopes() bool`

HasDockerhubTokenScopes returns a boolean if a field has been set.

### GetDockerhubUsername

`func (o *DynamicSecretUpdateDockerhub) GetDockerhubUsername() string`

GetDockerhubUsername returns the DockerhubUsername field if non-nil, zero value otherwise.

### GetDockerhubUsernameOk

`func (o *DynamicSecretUpdateDockerhub) GetDockerhubUsernameOk() (*string, bool)`

GetDockerhubUsernameOk returns a tuple with the DockerhubUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerhubUsername

`func (o *DynamicSecretUpdateDockerhub) SetDockerhubUsername(v string)`

SetDockerhubUsername sets DockerhubUsername field to given value.

### HasDockerhubUsername

`func (o *DynamicSecretUpdateDockerhub) HasDockerhubUsername() bool`

HasDockerhubUsername returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretUpdateDockerhub) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretUpdateDockerhub) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretUpdateDockerhub) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretUpdateDockerhub) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretUpdateDockerhub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretUpdateDockerhub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretUpdateDockerhub) SetName(v string)`

SetName sets Name field to given value.


### GetNewName

`func (o *DynamicSecretUpdateDockerhub) GetNewName() string`

GetNewName returns the NewName field if non-nil, zero value otherwise.

### GetNewNameOk

`func (o *DynamicSecretUpdateDockerhub) GetNewNameOk() (*string, bool)`

GetNewNameOk returns a tuple with the NewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewName

`func (o *DynamicSecretUpdateDockerhub) SetNewName(v string)`

SetNewName sets NewName field to given value.

### HasNewName

`func (o *DynamicSecretUpdateDockerhub) HasNewName() bool`

HasNewName returns a boolean if a field has been set.

### GetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateDockerhub) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretUpdateDockerhub) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretUpdateDockerhub) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretUpdateDockerhub) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetTags

`func (o *DynamicSecretUpdateDockerhub) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretUpdateDockerhub) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretUpdateDockerhub) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretUpdateDockerhub) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretUpdateDockerhub) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretUpdateDockerhub) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretUpdateDockerhub) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretUpdateDockerhub) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretUpdateDockerhub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretUpdateDockerhub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretUpdateDockerhub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretUpdateDockerhub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretUpdateDockerhub) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretUpdateDockerhub) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretUpdateDockerhub) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretUpdateDockerhub) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretUpdateDockerhub) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretUpdateDockerhub) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretUpdateDockerhub) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretUpdateDockerhub) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



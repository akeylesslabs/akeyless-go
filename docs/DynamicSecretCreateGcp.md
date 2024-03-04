# DynamicSecretCreateGcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**GcpCredType** | Pointer to **string** |  | [optional] 
**GcpKey** | Pointer to **string** | Base64-encoded service account private key text | [optional] 
**GcpKeyAlgo** | Pointer to **string** | Service account key algorithm, e.g. KEY_ALG_RSA_1024 | [optional] 
**GcpSaEmail** | Pointer to **string** | The email of the fixed service acocunt to generate keys or tokens for. (revelant for service-account-type&#x3D;fixed) | [optional] 
**GcpTokenScopes** | Pointer to **string** | Access token scopes list, e.g. scope1,scope2 | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Dynamic secret name | 
**ProducerEncryptionKeyName** | Pointer to **string** | Dynamic producer encryption key | [optional] 
**RoleBinding** | Pointer to **string** | Role binding definitions in json format | [optional] 
**ServiceAccountType** | **string** | The type of the gcp dynamic secret. Options[fixed, dynamic] | [default to "fixed"]
**Tags** | Pointer to **[]string** | Add tags attached to this object | [optional] 
**TargetName** | Pointer to **string** | Target name | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UserTtl** | Pointer to **string** | User TTL | [optional] [default to "60m"]

## Methods

### NewDynamicSecretCreateGcp

`func NewDynamicSecretCreateGcp(name string, serviceAccountType string, ) *DynamicSecretCreateGcp`

NewDynamicSecretCreateGcp instantiates a new DynamicSecretCreateGcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDynamicSecretCreateGcpWithDefaults

`func NewDynamicSecretCreateGcpWithDefaults() *DynamicSecretCreateGcp`

NewDynamicSecretCreateGcpWithDefaults instantiates a new DynamicSecretCreateGcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DynamicSecretCreateGcp) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DynamicSecretCreateGcp) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DynamicSecretCreateGcp) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DynamicSecretCreateGcp) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *DynamicSecretCreateGcp) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DynamicSecretCreateGcp) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DynamicSecretCreateGcp) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DynamicSecretCreateGcp) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGcpCredType

`func (o *DynamicSecretCreateGcp) GetGcpCredType() string`

GetGcpCredType returns the GcpCredType field if non-nil, zero value otherwise.

### GetGcpCredTypeOk

`func (o *DynamicSecretCreateGcp) GetGcpCredTypeOk() (*string, bool)`

GetGcpCredTypeOk returns a tuple with the GcpCredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredType

`func (o *DynamicSecretCreateGcp) SetGcpCredType(v string)`

SetGcpCredType sets GcpCredType field to given value.

### HasGcpCredType

`func (o *DynamicSecretCreateGcp) HasGcpCredType() bool`

HasGcpCredType returns a boolean if a field has been set.

### GetGcpKey

`func (o *DynamicSecretCreateGcp) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *DynamicSecretCreateGcp) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *DynamicSecretCreateGcp) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *DynamicSecretCreateGcp) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetGcpKeyAlgo

`func (o *DynamicSecretCreateGcp) GetGcpKeyAlgo() string`

GetGcpKeyAlgo returns the GcpKeyAlgo field if non-nil, zero value otherwise.

### GetGcpKeyAlgoOk

`func (o *DynamicSecretCreateGcp) GetGcpKeyAlgoOk() (*string, bool)`

GetGcpKeyAlgoOk returns a tuple with the GcpKeyAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKeyAlgo

`func (o *DynamicSecretCreateGcp) SetGcpKeyAlgo(v string)`

SetGcpKeyAlgo sets GcpKeyAlgo field to given value.

### HasGcpKeyAlgo

`func (o *DynamicSecretCreateGcp) HasGcpKeyAlgo() bool`

HasGcpKeyAlgo returns a boolean if a field has been set.

### GetGcpSaEmail

`func (o *DynamicSecretCreateGcp) GetGcpSaEmail() string`

GetGcpSaEmail returns the GcpSaEmail field if non-nil, zero value otherwise.

### GetGcpSaEmailOk

`func (o *DynamicSecretCreateGcp) GetGcpSaEmailOk() (*string, bool)`

GetGcpSaEmailOk returns a tuple with the GcpSaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSaEmail

`func (o *DynamicSecretCreateGcp) SetGcpSaEmail(v string)`

SetGcpSaEmail sets GcpSaEmail field to given value.

### HasGcpSaEmail

`func (o *DynamicSecretCreateGcp) HasGcpSaEmail() bool`

HasGcpSaEmail returns a boolean if a field has been set.

### GetGcpTokenScopes

`func (o *DynamicSecretCreateGcp) GetGcpTokenScopes() string`

GetGcpTokenScopes returns the GcpTokenScopes field if non-nil, zero value otherwise.

### GetGcpTokenScopesOk

`func (o *DynamicSecretCreateGcp) GetGcpTokenScopesOk() (*string, bool)`

GetGcpTokenScopesOk returns a tuple with the GcpTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTokenScopes

`func (o *DynamicSecretCreateGcp) SetGcpTokenScopes(v string)`

SetGcpTokenScopes sets GcpTokenScopes field to given value.

### HasGcpTokenScopes

`func (o *DynamicSecretCreateGcp) HasGcpTokenScopes() bool`

HasGcpTokenScopes returns a boolean if a field has been set.

### GetJson

`func (o *DynamicSecretCreateGcp) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DynamicSecretCreateGcp) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DynamicSecretCreateGcp) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DynamicSecretCreateGcp) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DynamicSecretCreateGcp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DynamicSecretCreateGcp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DynamicSecretCreateGcp) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DynamicSecretCreateGcp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DynamicSecretCreateGcp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DynamicSecretCreateGcp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DynamicSecretCreateGcp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRoleBinding

`func (o *DynamicSecretCreateGcp) GetRoleBinding() string`

GetRoleBinding returns the RoleBinding field if non-nil, zero value otherwise.

### GetRoleBindingOk

`func (o *DynamicSecretCreateGcp) GetRoleBindingOk() (*string, bool)`

GetRoleBindingOk returns a tuple with the RoleBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleBinding

`func (o *DynamicSecretCreateGcp) SetRoleBinding(v string)`

SetRoleBinding sets RoleBinding field to given value.

### HasRoleBinding

`func (o *DynamicSecretCreateGcp) HasRoleBinding() bool`

HasRoleBinding returns a boolean if a field has been set.

### GetServiceAccountType

`func (o *DynamicSecretCreateGcp) GetServiceAccountType() string`

GetServiceAccountType returns the ServiceAccountType field if non-nil, zero value otherwise.

### GetServiceAccountTypeOk

`func (o *DynamicSecretCreateGcp) GetServiceAccountTypeOk() (*string, bool)`

GetServiceAccountTypeOk returns a tuple with the ServiceAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountType

`func (o *DynamicSecretCreateGcp) SetServiceAccountType(v string)`

SetServiceAccountType sets ServiceAccountType field to given value.


### GetTags

`func (o *DynamicSecretCreateGcp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DynamicSecretCreateGcp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DynamicSecretCreateGcp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DynamicSecretCreateGcp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DynamicSecretCreateGcp) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DynamicSecretCreateGcp) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DynamicSecretCreateGcp) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DynamicSecretCreateGcp) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DynamicSecretCreateGcp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DynamicSecretCreateGcp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DynamicSecretCreateGcp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DynamicSecretCreateGcp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DynamicSecretCreateGcp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DynamicSecretCreateGcp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DynamicSecretCreateGcp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DynamicSecretCreateGcp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DynamicSecretCreateGcp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DynamicSecretCreateGcp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DynamicSecretCreateGcp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DynamicSecretCreateGcp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



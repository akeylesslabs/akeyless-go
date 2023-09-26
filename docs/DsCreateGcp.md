# DsCreateGcp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item [true/false] | [optional] 
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

### NewDsCreateGcp

`func NewDsCreateGcp(name string, serviceAccountType string, ) *DsCreateGcp`

NewDsCreateGcp instantiates a new DsCreateGcp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDsCreateGcpWithDefaults

`func NewDsCreateGcpWithDefaults() *DsCreateGcp`

NewDsCreateGcpWithDefaults instantiates a new DsCreateGcp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleteProtection

`func (o *DsCreateGcp) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *DsCreateGcp) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *DsCreateGcp) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *DsCreateGcp) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetGcpCredType

`func (o *DsCreateGcp) GetGcpCredType() string`

GetGcpCredType returns the GcpCredType field if non-nil, zero value otherwise.

### GetGcpCredTypeOk

`func (o *DsCreateGcp) GetGcpCredTypeOk() (*string, bool)`

GetGcpCredTypeOk returns a tuple with the GcpCredType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpCredType

`func (o *DsCreateGcp) SetGcpCredType(v string)`

SetGcpCredType sets GcpCredType field to given value.

### HasGcpCredType

`func (o *DsCreateGcp) HasGcpCredType() bool`

HasGcpCredType returns a boolean if a field has been set.

### GetGcpKey

`func (o *DsCreateGcp) GetGcpKey() string`

GetGcpKey returns the GcpKey field if non-nil, zero value otherwise.

### GetGcpKeyOk

`func (o *DsCreateGcp) GetGcpKeyOk() (*string, bool)`

GetGcpKeyOk returns a tuple with the GcpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKey

`func (o *DsCreateGcp) SetGcpKey(v string)`

SetGcpKey sets GcpKey field to given value.

### HasGcpKey

`func (o *DsCreateGcp) HasGcpKey() bool`

HasGcpKey returns a boolean if a field has been set.

### GetGcpKeyAlgo

`func (o *DsCreateGcp) GetGcpKeyAlgo() string`

GetGcpKeyAlgo returns the GcpKeyAlgo field if non-nil, zero value otherwise.

### GetGcpKeyAlgoOk

`func (o *DsCreateGcp) GetGcpKeyAlgoOk() (*string, bool)`

GetGcpKeyAlgoOk returns a tuple with the GcpKeyAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpKeyAlgo

`func (o *DsCreateGcp) SetGcpKeyAlgo(v string)`

SetGcpKeyAlgo sets GcpKeyAlgo field to given value.

### HasGcpKeyAlgo

`func (o *DsCreateGcp) HasGcpKeyAlgo() bool`

HasGcpKeyAlgo returns a boolean if a field has been set.

### GetGcpSaEmail

`func (o *DsCreateGcp) GetGcpSaEmail() string`

GetGcpSaEmail returns the GcpSaEmail field if non-nil, zero value otherwise.

### GetGcpSaEmailOk

`func (o *DsCreateGcp) GetGcpSaEmailOk() (*string, bool)`

GetGcpSaEmailOk returns a tuple with the GcpSaEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpSaEmail

`func (o *DsCreateGcp) SetGcpSaEmail(v string)`

SetGcpSaEmail sets GcpSaEmail field to given value.

### HasGcpSaEmail

`func (o *DsCreateGcp) HasGcpSaEmail() bool`

HasGcpSaEmail returns a boolean if a field has been set.

### GetGcpTokenScopes

`func (o *DsCreateGcp) GetGcpTokenScopes() string`

GetGcpTokenScopes returns the GcpTokenScopes field if non-nil, zero value otherwise.

### GetGcpTokenScopesOk

`func (o *DsCreateGcp) GetGcpTokenScopesOk() (*string, bool)`

GetGcpTokenScopesOk returns a tuple with the GcpTokenScopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcpTokenScopes

`func (o *DsCreateGcp) SetGcpTokenScopes(v string)`

SetGcpTokenScopes sets GcpTokenScopes field to given value.

### HasGcpTokenScopes

`func (o *DsCreateGcp) HasGcpTokenScopes() bool`

HasGcpTokenScopes returns a boolean if a field has been set.

### GetJson

`func (o *DsCreateGcp) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DsCreateGcp) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DsCreateGcp) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DsCreateGcp) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *DsCreateGcp) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DsCreateGcp) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DsCreateGcp) SetName(v string)`

SetName sets Name field to given value.


### GetProducerEncryptionKeyName

`func (o *DsCreateGcp) GetProducerEncryptionKeyName() string`

GetProducerEncryptionKeyName returns the ProducerEncryptionKeyName field if non-nil, zero value otherwise.

### GetProducerEncryptionKeyNameOk

`func (o *DsCreateGcp) GetProducerEncryptionKeyNameOk() (*string, bool)`

GetProducerEncryptionKeyNameOk returns a tuple with the ProducerEncryptionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducerEncryptionKeyName

`func (o *DsCreateGcp) SetProducerEncryptionKeyName(v string)`

SetProducerEncryptionKeyName sets ProducerEncryptionKeyName field to given value.

### HasProducerEncryptionKeyName

`func (o *DsCreateGcp) HasProducerEncryptionKeyName() bool`

HasProducerEncryptionKeyName returns a boolean if a field has been set.

### GetRoleBinding

`func (o *DsCreateGcp) GetRoleBinding() string`

GetRoleBinding returns the RoleBinding field if non-nil, zero value otherwise.

### GetRoleBindingOk

`func (o *DsCreateGcp) GetRoleBindingOk() (*string, bool)`

GetRoleBindingOk returns a tuple with the RoleBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleBinding

`func (o *DsCreateGcp) SetRoleBinding(v string)`

SetRoleBinding sets RoleBinding field to given value.

### HasRoleBinding

`func (o *DsCreateGcp) HasRoleBinding() bool`

HasRoleBinding returns a boolean if a field has been set.

### GetServiceAccountType

`func (o *DsCreateGcp) GetServiceAccountType() string`

GetServiceAccountType returns the ServiceAccountType field if non-nil, zero value otherwise.

### GetServiceAccountTypeOk

`func (o *DsCreateGcp) GetServiceAccountTypeOk() (*string, bool)`

GetServiceAccountTypeOk returns a tuple with the ServiceAccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountType

`func (o *DsCreateGcp) SetServiceAccountType(v string)`

SetServiceAccountType sets ServiceAccountType field to given value.


### GetTags

`func (o *DsCreateGcp) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DsCreateGcp) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DsCreateGcp) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DsCreateGcp) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargetName

`func (o *DsCreateGcp) GetTargetName() string`

GetTargetName returns the TargetName field if non-nil, zero value otherwise.

### GetTargetNameOk

`func (o *DsCreateGcp) GetTargetNameOk() (*string, bool)`

GetTargetNameOk returns a tuple with the TargetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetName

`func (o *DsCreateGcp) SetTargetName(v string)`

SetTargetName sets TargetName field to given value.

### HasTargetName

`func (o *DsCreateGcp) HasTargetName() bool`

HasTargetName returns a boolean if a field has been set.

### GetToken

`func (o *DsCreateGcp) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DsCreateGcp) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DsCreateGcp) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DsCreateGcp) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DsCreateGcp) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DsCreateGcp) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DsCreateGcp) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DsCreateGcp) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUserTtl

`func (o *DsCreateGcp) GetUserTtl() string`

GetUserTtl returns the UserTtl field if non-nil, zero value otherwise.

### GetUserTtlOk

`func (o *DsCreateGcp) GetUserTtlOk() (*string, bool)`

GetUserTtlOk returns a tuple with the UserTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTtl

`func (o *DsCreateGcp) SetUserTtl(v string)`

SetUserTtl sets UserTtl field to given value.

### HasUserTtl

`func (o *DsCreateGcp) HasUserTtl() bool`

HasUserTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



# UscCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BinaryValue** | Pointer to **bool** | Use this option if the universal secrets value is a base64 encoded binary | [optional] 
**Description** | Pointer to **string** | Description of the universal secrets | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Namespace** | Pointer to **string** | The namespace (relevant for Hashi vault target) | [optional] 
**ObjectType** | Pointer to **string** |  | [optional] 
**PfxPassword** | Pointer to **string** | Optional, the passphrase that protects the private key within the pfx certificate (Relevant only for Azure KV certificates) | [optional] 
**Region** | Pointer to **string** | Optional, create secret in a specific region (GCP only). If empty, a global secret will be created (provider default). | [optional] 
**SecretName** | **string** | Name for the new universal secrets | 
**Tags** | Pointer to **map[string]string** | Tags for the universal secrets | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**UscEncryptionKey** | Pointer to **string** | Optional, The name of the remote key that used to encrypt the secret value (if empty, the default key will be used) | [optional] 
**UscName** | **string** | Name of the Universal Secrets Connector item | 
**Value** | **string** | Value of the universal secrets item, either text or base64 encoded binary | 

## Methods

### NewUscCreate

`func NewUscCreate(secretName string, uscName string, value string, ) *UscCreate`

NewUscCreate instantiates a new UscCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUscCreateWithDefaults

`func NewUscCreateWithDefaults() *UscCreate`

NewUscCreateWithDefaults instantiates a new UscCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBinaryValue

`func (o *UscCreate) GetBinaryValue() bool`

GetBinaryValue returns the BinaryValue field if non-nil, zero value otherwise.

### GetBinaryValueOk

`func (o *UscCreate) GetBinaryValueOk() (*bool, bool)`

GetBinaryValueOk returns a tuple with the BinaryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinaryValue

`func (o *UscCreate) SetBinaryValue(v bool)`

SetBinaryValue sets BinaryValue field to given value.

### HasBinaryValue

`func (o *UscCreate) HasBinaryValue() bool`

HasBinaryValue returns a boolean if a field has been set.

### GetDescription

`func (o *UscCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UscCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UscCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UscCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *UscCreate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UscCreate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UscCreate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UscCreate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetNamespace

`func (o *UscCreate) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *UscCreate) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *UscCreate) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *UscCreate) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetObjectType

`func (o *UscCreate) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *UscCreate) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *UscCreate) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *UscCreate) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetPfxPassword

`func (o *UscCreate) GetPfxPassword() string`

GetPfxPassword returns the PfxPassword field if non-nil, zero value otherwise.

### GetPfxPasswordOk

`func (o *UscCreate) GetPfxPasswordOk() (*string, bool)`

GetPfxPasswordOk returns a tuple with the PfxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfxPassword

`func (o *UscCreate) SetPfxPassword(v string)`

SetPfxPassword sets PfxPassword field to given value.

### HasPfxPassword

`func (o *UscCreate) HasPfxPassword() bool`

HasPfxPassword returns a boolean if a field has been set.

### GetRegion

`func (o *UscCreate) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UscCreate) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UscCreate) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UscCreate) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSecretName

`func (o *UscCreate) GetSecretName() string`

GetSecretName returns the SecretName field if non-nil, zero value otherwise.

### GetSecretNameOk

`func (o *UscCreate) GetSecretNameOk() (*string, bool)`

GetSecretNameOk returns a tuple with the SecretName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretName

`func (o *UscCreate) SetSecretName(v string)`

SetSecretName sets SecretName field to given value.


### GetTags

`func (o *UscCreate) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UscCreate) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UscCreate) SetTags(v map[string]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UscCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetToken

`func (o *UscCreate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UscCreate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UscCreate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UscCreate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UscCreate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UscCreate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UscCreate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UscCreate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUscEncryptionKey

`func (o *UscCreate) GetUscEncryptionKey() string`

GetUscEncryptionKey returns the UscEncryptionKey field if non-nil, zero value otherwise.

### GetUscEncryptionKeyOk

`func (o *UscCreate) GetUscEncryptionKeyOk() (*string, bool)`

GetUscEncryptionKeyOk returns a tuple with the UscEncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscEncryptionKey

`func (o *UscCreate) SetUscEncryptionKey(v string)`

SetUscEncryptionKey sets UscEncryptionKey field to given value.

### HasUscEncryptionKey

`func (o *UscCreate) HasUscEncryptionKey() bool`

HasUscEncryptionKey returns a boolean if a field has been set.

### GetUscName

`func (o *UscCreate) GetUscName() string`

GetUscName returns the UscName field if non-nil, zero value otherwise.

### GetUscNameOk

`func (o *UscCreate) GetUscNameOk() (*string, bool)`

GetUscNameOk returns a tuple with the UscName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscName

`func (o *UscCreate) SetUscName(v string)`

SetUscName sets UscName field to given value.


### GetValue

`func (o *UscCreate) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UscCreate) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UscCreate) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



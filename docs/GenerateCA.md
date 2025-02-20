# GenerateCA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**AllowedDomains** | **string** | A list of the allowed domains that clients can request to be included in the certificate (in a comma-delimited list) | 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**ExtendedKeyUsage** | Pointer to **string** | A comma sepereted list of extended key usage for the intermediate (serverauth / clientauth / codesigning) | [optional] [default to "serverauth,clientauth"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyType** | Pointer to **string** |  | [optional] 
**MaxPathLen** | Pointer to **int64** | The maximum number of intermediate certificates that can appear in a certification path | [optional] [default to 1]
**PkiChainName** | **string** | PKI chain name | 
**ProtectionKeyName** | Pointer to **string** | The name of a key that used to encrypt the secrets values (if empty, the account default protectionKey key will be used) | [optional] 
**SplitLevel** | Pointer to **int64** | The number of fragments that the item will be split into | [optional] [default to 3]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | **string** | The maximum requested Time To Live for issued certificate by default in seconds, supported formats are s,m,h,d | 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGenerateCA

`func NewGenerateCA(allowedDomains string, pkiChainName string, ttl string, ) *GenerateCA`

NewGenerateCA instantiates a new GenerateCA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateCAWithDefaults

`func NewGenerateCAWithDefaults() *GenerateCA`

NewGenerateCAWithDefaults instantiates a new GenerateCA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *GenerateCA) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *GenerateCA) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *GenerateCA) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *GenerateCA) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *GenerateCA) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *GenerateCA) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *GenerateCA) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.


### GetDeleteProtection

`func (o *GenerateCA) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GenerateCA) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GenerateCA) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GenerateCA) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *GenerateCA) GetExtendedKeyUsage() string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *GenerateCA) GetExtendedKeyUsageOk() (*string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *GenerateCA) SetExtendedKeyUsage(v string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *GenerateCA) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetJson

`func (o *GenerateCA) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GenerateCA) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GenerateCA) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GenerateCA) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyType

`func (o *GenerateCA) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *GenerateCA) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *GenerateCA) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *GenerateCA) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetMaxPathLen

`func (o *GenerateCA) GetMaxPathLen() int64`

GetMaxPathLen returns the MaxPathLen field if non-nil, zero value otherwise.

### GetMaxPathLenOk

`func (o *GenerateCA) GetMaxPathLenOk() (*int64, bool)`

GetMaxPathLenOk returns a tuple with the MaxPathLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLen

`func (o *GenerateCA) SetMaxPathLen(v int64)`

SetMaxPathLen sets MaxPathLen field to given value.

### HasMaxPathLen

`func (o *GenerateCA) HasMaxPathLen() bool`

HasMaxPathLen returns a boolean if a field has been set.

### GetPkiChainName

`func (o *GenerateCA) GetPkiChainName() string`

GetPkiChainName returns the PkiChainName field if non-nil, zero value otherwise.

### GetPkiChainNameOk

`func (o *GenerateCA) GetPkiChainNameOk() (*string, bool)`

GetPkiChainNameOk returns a tuple with the PkiChainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPkiChainName

`func (o *GenerateCA) SetPkiChainName(v string)`

SetPkiChainName sets PkiChainName field to given value.


### GetProtectionKeyName

`func (o *GenerateCA) GetProtectionKeyName() string`

GetProtectionKeyName returns the ProtectionKeyName field if non-nil, zero value otherwise.

### GetProtectionKeyNameOk

`func (o *GenerateCA) GetProtectionKeyNameOk() (*string, bool)`

GetProtectionKeyNameOk returns a tuple with the ProtectionKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectionKeyName

`func (o *GenerateCA) SetProtectionKeyName(v string)`

SetProtectionKeyName sets ProtectionKeyName field to given value.

### HasProtectionKeyName

`func (o *GenerateCA) HasProtectionKeyName() bool`

HasProtectionKeyName returns a boolean if a field has been set.

### GetSplitLevel

`func (o *GenerateCA) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *GenerateCA) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *GenerateCA) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *GenerateCA) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetToken

`func (o *GenerateCA) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GenerateCA) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GenerateCA) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GenerateCA) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *GenerateCA) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GenerateCA) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GenerateCA) SetTtl(v string)`

SetTtl sets Ttl field to given value.


### GetUidToken

`func (o *GenerateCA) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GenerateCA) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GenerateCA) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GenerateCA) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



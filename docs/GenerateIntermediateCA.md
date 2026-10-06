# GenerateIntermediateCA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | Pointer to **string** |  | [optional] 
**AllowedDomains** | Pointer to **string** | Allowed domains for future leaf issuance, not inherited into the SCEP subordinate CA certificate | [optional] 
**CommonName** | Pointer to **string** | Optional Common Name for the intermediate CA certificate | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this object [true/false] | [optional] 
**DestinationPath** | Pointer to **string** | Destination path for SCEP-issued leaf certificates. Not derived from the CA certificate item path. | [optional] 
**EnableScep** | Pointer to **bool** | Enable the fixed SCEP Stage 1 profile | [optional] 
**ExtendedKeyUsage** | Pointer to **string** | Extended key usage for future leaf issuance (serverauth / clientauth / codesigning) | [optional] [default to "serverauth,clientauth"]
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**MaxPathLen** | Pointer to **int64** | The maximum path length of the generated intermediate CA certificate | [optional] [default to 0]
**Name** | **string** | Base path for derived intermediate CA resources | 
**ParentCaName** | Pointer to **string** | Parent PKI certificate issuer name | [optional] 
**ScepPassword** | Pointer to **string** | SCEP static challenge password. Request-only; never returned | [optional] 
**SplitLevel** | Pointer to **int64** | The number of fragments that the DFC key will be split into | [optional] [default to 3]
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**Ttl** | Pointer to **string** | Maximum TTL for certificates issued by the new intermediate issuer, supported formats are s,m,h,d | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewGenerateIntermediateCA

`func NewGenerateIntermediateCA(name string, ) *GenerateIntermediateCA`

NewGenerateIntermediateCA instantiates a new GenerateIntermediateCA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateIntermediateCAWithDefaults

`func NewGenerateIntermediateCAWithDefaults() *GenerateIntermediateCA`

NewGenerateIntermediateCAWithDefaults instantiates a new GenerateIntermediateCA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *GenerateIntermediateCA) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *GenerateIntermediateCA) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *GenerateIntermediateCA) SetAlg(v string)`

SetAlg sets Alg field to given value.

### HasAlg

`func (o *GenerateIntermediateCA) HasAlg() bool`

HasAlg returns a boolean if a field has been set.

### GetAllowedDomains

`func (o *GenerateIntermediateCA) GetAllowedDomains() string`

GetAllowedDomains returns the AllowedDomains field if non-nil, zero value otherwise.

### GetAllowedDomainsOk

`func (o *GenerateIntermediateCA) GetAllowedDomainsOk() (*string, bool)`

GetAllowedDomainsOk returns a tuple with the AllowedDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedDomains

`func (o *GenerateIntermediateCA) SetAllowedDomains(v string)`

SetAllowedDomains sets AllowedDomains field to given value.

### HasAllowedDomains

`func (o *GenerateIntermediateCA) HasAllowedDomains() bool`

HasAllowedDomains returns a boolean if a field has been set.

### GetCommonName

`func (o *GenerateIntermediateCA) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *GenerateIntermediateCA) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *GenerateIntermediateCA) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *GenerateIntermediateCA) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *GenerateIntermediateCA) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *GenerateIntermediateCA) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *GenerateIntermediateCA) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *GenerateIntermediateCA) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDestinationPath

`func (o *GenerateIntermediateCA) GetDestinationPath() string`

GetDestinationPath returns the DestinationPath field if non-nil, zero value otherwise.

### GetDestinationPathOk

`func (o *GenerateIntermediateCA) GetDestinationPathOk() (*string, bool)`

GetDestinationPathOk returns a tuple with the DestinationPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPath

`func (o *GenerateIntermediateCA) SetDestinationPath(v string)`

SetDestinationPath sets DestinationPath field to given value.

### HasDestinationPath

`func (o *GenerateIntermediateCA) HasDestinationPath() bool`

HasDestinationPath returns a boolean if a field has been set.

### GetEnableScep

`func (o *GenerateIntermediateCA) GetEnableScep() bool`

GetEnableScep returns the EnableScep field if non-nil, zero value otherwise.

### GetEnableScepOk

`func (o *GenerateIntermediateCA) GetEnableScepOk() (*bool, bool)`

GetEnableScepOk returns a tuple with the EnableScep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableScep

`func (o *GenerateIntermediateCA) SetEnableScep(v bool)`

SetEnableScep sets EnableScep field to given value.

### HasEnableScep

`func (o *GenerateIntermediateCA) HasEnableScep() bool`

HasEnableScep returns a boolean if a field has been set.

### GetExtendedKeyUsage

`func (o *GenerateIntermediateCA) GetExtendedKeyUsage() string`

GetExtendedKeyUsage returns the ExtendedKeyUsage field if non-nil, zero value otherwise.

### GetExtendedKeyUsageOk

`func (o *GenerateIntermediateCA) GetExtendedKeyUsageOk() (*string, bool)`

GetExtendedKeyUsageOk returns a tuple with the ExtendedKeyUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtendedKeyUsage

`func (o *GenerateIntermediateCA) SetExtendedKeyUsage(v string)`

SetExtendedKeyUsage sets ExtendedKeyUsage field to given value.

### HasExtendedKeyUsage

`func (o *GenerateIntermediateCA) HasExtendedKeyUsage() bool`

HasExtendedKeyUsage returns a boolean if a field has been set.

### GetJson

`func (o *GenerateIntermediateCA) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GenerateIntermediateCA) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GenerateIntermediateCA) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GenerateIntermediateCA) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMaxPathLen

`func (o *GenerateIntermediateCA) GetMaxPathLen() int64`

GetMaxPathLen returns the MaxPathLen field if non-nil, zero value otherwise.

### GetMaxPathLenOk

`func (o *GenerateIntermediateCA) GetMaxPathLenOk() (*int64, bool)`

GetMaxPathLenOk returns a tuple with the MaxPathLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPathLen

`func (o *GenerateIntermediateCA) SetMaxPathLen(v int64)`

SetMaxPathLen sets MaxPathLen field to given value.

### HasMaxPathLen

`func (o *GenerateIntermediateCA) HasMaxPathLen() bool`

HasMaxPathLen returns a boolean if a field has been set.

### GetName

`func (o *GenerateIntermediateCA) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateIntermediateCA) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateIntermediateCA) SetName(v string)`

SetName sets Name field to given value.


### GetParentCaName

`func (o *GenerateIntermediateCA) GetParentCaName() string`

GetParentCaName returns the ParentCaName field if non-nil, zero value otherwise.

### GetParentCaNameOk

`func (o *GenerateIntermediateCA) GetParentCaNameOk() (*string, bool)`

GetParentCaNameOk returns a tuple with the ParentCaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentCaName

`func (o *GenerateIntermediateCA) SetParentCaName(v string)`

SetParentCaName sets ParentCaName field to given value.

### HasParentCaName

`func (o *GenerateIntermediateCA) HasParentCaName() bool`

HasParentCaName returns a boolean if a field has been set.

### GetScepPassword

`func (o *GenerateIntermediateCA) GetScepPassword() string`

GetScepPassword returns the ScepPassword field if non-nil, zero value otherwise.

### GetScepPasswordOk

`func (o *GenerateIntermediateCA) GetScepPasswordOk() (*string, bool)`

GetScepPasswordOk returns a tuple with the ScepPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScepPassword

`func (o *GenerateIntermediateCA) SetScepPassword(v string)`

SetScepPassword sets ScepPassword field to given value.

### HasScepPassword

`func (o *GenerateIntermediateCA) HasScepPassword() bool`

HasScepPassword returns a boolean if a field has been set.

### GetSplitLevel

`func (o *GenerateIntermediateCA) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *GenerateIntermediateCA) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *GenerateIntermediateCA) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *GenerateIntermediateCA) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetToken

`func (o *GenerateIntermediateCA) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GenerateIntermediateCA) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GenerateIntermediateCA) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GenerateIntermediateCA) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTtl

`func (o *GenerateIntermediateCA) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *GenerateIntermediateCA) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *GenerateIntermediateCA) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *GenerateIntermediateCA) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUidToken

`func (o *GenerateIntermediateCA) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GenerateIntermediateCA) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GenerateIntermediateCA) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GenerateIntermediateCA) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



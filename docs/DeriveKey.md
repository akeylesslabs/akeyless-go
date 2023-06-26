# DeriveKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accessibility** | Pointer to **string** | for personal password manager | [optional] [default to "regular"]
**Alg** | **string** | Kdf algorithm | [default to "pbkdf2"]
**HashFunction** | Pointer to **string** | HashFunction the hash function to use (relevant for pbkdf2) | [optional] [default to "sha256"]
**Iter** | **int64** | IterationCount the number of iterations | 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**KeyLen** | **int64** | KeyLength the byte length of the generated key | 
**Mem** | Pointer to **int64** | MemorySizeInKb the memory paramter in kb (relevant for argon2id) | [optional] [default to 16384]
**Name** | **string** | Static Secret full name | 
**Parallelism** | Pointer to **int64** | Parallelism the number of threads to use (relevant for argon2id) | [optional] [default to 1]
**Salt** | Pointer to **string** | Salt Base64 encoded salt value. If not provided, the salt will be generated as part of the operation. The salt should be securely-generated random bytes, minimum 64 bits, 128 bits is recommended | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewDeriveKey

`func NewDeriveKey(alg string, iter int64, keyLen int64, name string, ) *DeriveKey`

NewDeriveKey instantiates a new DeriveKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeriveKeyWithDefaults

`func NewDeriveKeyWithDefaults() *DeriveKey`

NewDeriveKeyWithDefaults instantiates a new DeriveKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessibility

`func (o *DeriveKey) GetAccessibility() string`

GetAccessibility returns the Accessibility field if non-nil, zero value otherwise.

### GetAccessibilityOk

`func (o *DeriveKey) GetAccessibilityOk() (*string, bool)`

GetAccessibilityOk returns a tuple with the Accessibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessibility

`func (o *DeriveKey) SetAccessibility(v string)`

SetAccessibility sets Accessibility field to given value.

### HasAccessibility

`func (o *DeriveKey) HasAccessibility() bool`

HasAccessibility returns a boolean if a field has been set.

### GetAlg

`func (o *DeriveKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *DeriveKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *DeriveKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetHashFunction

`func (o *DeriveKey) GetHashFunction() string`

GetHashFunction returns the HashFunction field if non-nil, zero value otherwise.

### GetHashFunctionOk

`func (o *DeriveKey) GetHashFunctionOk() (*string, bool)`

GetHashFunctionOk returns a tuple with the HashFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashFunction

`func (o *DeriveKey) SetHashFunction(v string)`

SetHashFunction sets HashFunction field to given value.

### HasHashFunction

`func (o *DeriveKey) HasHashFunction() bool`

HasHashFunction returns a boolean if a field has been set.

### GetIter

`func (o *DeriveKey) GetIter() int64`

GetIter returns the Iter field if non-nil, zero value otherwise.

### GetIterOk

`func (o *DeriveKey) GetIterOk() (*int64, bool)`

GetIterOk returns a tuple with the Iter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIter

`func (o *DeriveKey) SetIter(v int64)`

SetIter sets Iter field to given value.


### GetJson

`func (o *DeriveKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *DeriveKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *DeriveKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *DeriveKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetKeyLen

`func (o *DeriveKey) GetKeyLen() int64`

GetKeyLen returns the KeyLen field if non-nil, zero value otherwise.

### GetKeyLenOk

`func (o *DeriveKey) GetKeyLenOk() (*int64, bool)`

GetKeyLenOk returns a tuple with the KeyLen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLen

`func (o *DeriveKey) SetKeyLen(v int64)`

SetKeyLen sets KeyLen field to given value.


### GetMem

`func (o *DeriveKey) GetMem() int64`

GetMem returns the Mem field if non-nil, zero value otherwise.

### GetMemOk

`func (o *DeriveKey) GetMemOk() (*int64, bool)`

GetMemOk returns a tuple with the Mem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMem

`func (o *DeriveKey) SetMem(v int64)`

SetMem sets Mem field to given value.

### HasMem

`func (o *DeriveKey) HasMem() bool`

HasMem returns a boolean if a field has been set.

### GetName

`func (o *DeriveKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeriveKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeriveKey) SetName(v string)`

SetName sets Name field to given value.


### GetParallelism

`func (o *DeriveKey) GetParallelism() int64`

GetParallelism returns the Parallelism field if non-nil, zero value otherwise.

### GetParallelismOk

`func (o *DeriveKey) GetParallelismOk() (*int64, bool)`

GetParallelismOk returns a tuple with the Parallelism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParallelism

`func (o *DeriveKey) SetParallelism(v int64)`

SetParallelism sets Parallelism field to given value.

### HasParallelism

`func (o *DeriveKey) HasParallelism() bool`

HasParallelism returns a boolean if a field has been set.

### GetSalt

`func (o *DeriveKey) GetSalt() string`

GetSalt returns the Salt field if non-nil, zero value otherwise.

### GetSaltOk

`func (o *DeriveKey) GetSaltOk() (*string, bool)`

GetSaltOk returns a tuple with the Salt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalt

`func (o *DeriveKey) SetSalt(v string)`

SetSalt sets Salt field to given value.

### HasSalt

`func (o *DeriveKey) HasSalt() bool`

HasSalt returns a boolean if a field has been set.

### GetToken

`func (o *DeriveKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeriveKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeriveKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeriveKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *DeriveKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *DeriveKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *DeriveKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *DeriveKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



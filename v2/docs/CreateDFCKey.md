# CreateDFCKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | DFCKey type; options: [AES128GCM, AES256GCM, AES128SIV, AES256SIV, AES128CBC, AES256CBC, RSA1024, RSA2048, RSA3072, RSA4096] | 
**CustomerFrgId** | Pointer to **string** | The customer fragment ID that will be used to create the DFC key (if empty, the key will be created independently of a customer fragment) | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item | [optional] 
**Description** | Pointer to **string** | Description of the object | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] 
**Metadata** | Pointer to **string** | Deprecated - use description | [optional] 
**Name** | **string** | DFCKey name | 
**SplitLevel** | Pointer to **int64** | The number of fragments that the item will be split into (not includes customer fragment) | [optional] [default to 3]
**Tag** | Pointer to **[]string** | List of the tags attached to this DFC key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewCreateDFCKey

`func NewCreateDFCKey(alg string, name string, ) *CreateDFCKey`

NewCreateDFCKey instantiates a new CreateDFCKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDFCKeyWithDefaults

`func NewCreateDFCKeyWithDefaults() *CreateDFCKey`

NewCreateDFCKeyWithDefaults instantiates a new CreateDFCKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *CreateDFCKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *CreateDFCKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *CreateDFCKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetCustomerFrgId

`func (o *CreateDFCKey) GetCustomerFrgId() string`

GetCustomerFrgId returns the CustomerFrgId field if non-nil, zero value otherwise.

### GetCustomerFrgIdOk

`func (o *CreateDFCKey) GetCustomerFrgIdOk() (*string, bool)`

GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFrgId

`func (o *CreateDFCKey) SetCustomerFrgId(v string)`

SetCustomerFrgId sets CustomerFrgId field to given value.

### HasCustomerFrgId

`func (o *CreateDFCKey) HasCustomerFrgId() bool`

HasCustomerFrgId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *CreateDFCKey) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *CreateDFCKey) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *CreateDFCKey) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *CreateDFCKey) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetDescription

`func (o *CreateDFCKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateDFCKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateDFCKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateDFCKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJson

`func (o *CreateDFCKey) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *CreateDFCKey) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *CreateDFCKey) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *CreateDFCKey) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateDFCKey) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateDFCKey) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateDFCKey) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateDFCKey) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *CreateDFCKey) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDFCKey) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDFCKey) SetName(v string)`

SetName sets Name field to given value.


### GetSplitLevel

`func (o *CreateDFCKey) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *CreateDFCKey) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *CreateDFCKey) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *CreateDFCKey) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetTag

`func (o *CreateDFCKey) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *CreateDFCKey) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *CreateDFCKey) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *CreateDFCKey) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetToken

`func (o *CreateDFCKey) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateDFCKey) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateDFCKey) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CreateDFCKey) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *CreateDFCKey) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *CreateDFCKey) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *CreateDFCKey) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *CreateDFCKey) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



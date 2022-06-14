# UploadRSA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | Key type. options: [RSA1024, RSA2048, RSA3072, RSA4096] | 
**CertFileData** | Pointer to **string** | Certificate in a PEM format. | [optional] 
**CustomerFrgId** | Pointer to **string** | The customer fragment ID that will be used to split the key (if empty, the key will be created independently of a customer fragment) | [optional] 
**DeleteProtection** | Pointer to **string** | Protection from accidental deletion of this item | [optional] 
**Metadata** | Pointer to **string** | A metadata about the key | [optional] 
**Name** | **string** | Name of key to be created | 
**RsaFileData** | Pointer to **string** | RSA private key data, base64 encoded | [optional] 
**SplitLevel** | Pointer to **int64** | The number of fragments that the item will be split into | [optional] [default to 2]
**Tag** | Pointer to **[]string** | List of the tags attached to this key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUploadRSA

`func NewUploadRSA(alg string, name string, ) *UploadRSA`

NewUploadRSA instantiates a new UploadRSA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadRSAWithDefaults

`func NewUploadRSAWithDefaults() *UploadRSA`

NewUploadRSAWithDefaults instantiates a new UploadRSA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *UploadRSA) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *UploadRSA) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *UploadRSA) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetCertFileData

`func (o *UploadRSA) GetCertFileData() string`

GetCertFileData returns the CertFileData field if non-nil, zero value otherwise.

### GetCertFileDataOk

`func (o *UploadRSA) GetCertFileDataOk() (*string, bool)`

GetCertFileDataOk returns a tuple with the CertFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFileData

`func (o *UploadRSA) SetCertFileData(v string)`

SetCertFileData sets CertFileData field to given value.

### HasCertFileData

`func (o *UploadRSA) HasCertFileData() bool`

HasCertFileData returns a boolean if a field has been set.

### GetCustomerFrgId

`func (o *UploadRSA) GetCustomerFrgId() string`

GetCustomerFrgId returns the CustomerFrgId field if non-nil, zero value otherwise.

### GetCustomerFrgIdOk

`func (o *UploadRSA) GetCustomerFrgIdOk() (*string, bool)`

GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFrgId

`func (o *UploadRSA) SetCustomerFrgId(v string)`

SetCustomerFrgId sets CustomerFrgId field to given value.

### HasCustomerFrgId

`func (o *UploadRSA) HasCustomerFrgId() bool`

HasCustomerFrgId returns a boolean if a field has been set.

### GetDeleteProtection

`func (o *UploadRSA) GetDeleteProtection() string`

GetDeleteProtection returns the DeleteProtection field if non-nil, zero value otherwise.

### GetDeleteProtectionOk

`func (o *UploadRSA) GetDeleteProtectionOk() (*string, bool)`

GetDeleteProtectionOk returns a tuple with the DeleteProtection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteProtection

`func (o *UploadRSA) SetDeleteProtection(v string)`

SetDeleteProtection sets DeleteProtection field to given value.

### HasDeleteProtection

`func (o *UploadRSA) HasDeleteProtection() bool`

HasDeleteProtection returns a boolean if a field has been set.

### GetMetadata

`func (o *UploadRSA) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UploadRSA) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UploadRSA) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UploadRSA) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UploadRSA) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadRSA) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadRSA) SetName(v string)`

SetName sets Name field to given value.


### GetRsaFileData

`func (o *UploadRSA) GetRsaFileData() string`

GetRsaFileData returns the RsaFileData field if non-nil, zero value otherwise.

### GetRsaFileDataOk

`func (o *UploadRSA) GetRsaFileDataOk() (*string, bool)`

GetRsaFileDataOk returns a tuple with the RsaFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsaFileData

`func (o *UploadRSA) SetRsaFileData(v string)`

SetRsaFileData sets RsaFileData field to given value.

### HasRsaFileData

`func (o *UploadRSA) HasRsaFileData() bool`

HasRsaFileData returns a boolean if a field has been set.

### GetSplitLevel

`func (o *UploadRSA) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *UploadRSA) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *UploadRSA) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *UploadRSA) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetTag

`func (o *UploadRSA) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UploadRSA) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UploadRSA) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *UploadRSA) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetToken

`func (o *UploadRSA) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UploadRSA) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UploadRSA) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UploadRSA) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UploadRSA) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UploadRSA) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UploadRSA) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UploadRSA) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



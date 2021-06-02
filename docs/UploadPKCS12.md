# UploadPKCS12

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cert** | Pointer to **string** | Path to a file that contain the certificate in a PEM format. If this parameter is not empty, the certificate will be taken from here and not from the PKCS#12 input file | [optional] 
**CustomerFrgId** | Pointer to **string** | The customer fragment ID that will be used to split the key (if empty, the key will be created independently of a customer fragment) | [optional] 
**In** | **string** | PKCS#12 input file (private key and certificate only) | 
**Metadata** | Pointer to **string** | A metadata about the key | [optional] 
**Name** | **string** | Name of key to be created | 
**Passphrase** | **string** | Passphrase to unlock the pkcs#12 bundle | 
**Password** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 
**SplitLevel** | Pointer to **int64** | The number of fragments that the item will be split into | [optional] [default to 2]
**Tag** | Pointer to **[]string** | List of the tags attached to this key | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Username** | Pointer to **string** | Required only when the authentication process requires a username and password | [optional] 

## Methods

### NewUploadPKCS12

`func NewUploadPKCS12(in string, name string, passphrase string, ) *UploadPKCS12`

NewUploadPKCS12 instantiates a new UploadPKCS12 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadPKCS12WithDefaults

`func NewUploadPKCS12WithDefaults() *UploadPKCS12`

NewUploadPKCS12WithDefaults instantiates a new UploadPKCS12 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCert

`func (o *UploadPKCS12) GetCert() string`

GetCert returns the Cert field if non-nil, zero value otherwise.

### GetCertOk

`func (o *UploadPKCS12) GetCertOk() (*string, bool)`

GetCertOk returns a tuple with the Cert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCert

`func (o *UploadPKCS12) SetCert(v string)`

SetCert sets Cert field to given value.

### HasCert

`func (o *UploadPKCS12) HasCert() bool`

HasCert returns a boolean if a field has been set.

### GetCustomerFrgId

`func (o *UploadPKCS12) GetCustomerFrgId() string`

GetCustomerFrgId returns the CustomerFrgId field if non-nil, zero value otherwise.

### GetCustomerFrgIdOk

`func (o *UploadPKCS12) GetCustomerFrgIdOk() (*string, bool)`

GetCustomerFrgIdOk returns a tuple with the CustomerFrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFrgId

`func (o *UploadPKCS12) SetCustomerFrgId(v string)`

SetCustomerFrgId sets CustomerFrgId field to given value.

### HasCustomerFrgId

`func (o *UploadPKCS12) HasCustomerFrgId() bool`

HasCustomerFrgId returns a boolean if a field has been set.

### GetIn

`func (o *UploadPKCS12) GetIn() string`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *UploadPKCS12) GetInOk() (*string, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *UploadPKCS12) SetIn(v string)`

SetIn sets In field to given value.


### GetMetadata

`func (o *UploadPKCS12) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UploadPKCS12) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UploadPKCS12) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UploadPKCS12) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *UploadPKCS12) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadPKCS12) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadPKCS12) SetName(v string)`

SetName sets Name field to given value.


### GetPassphrase

`func (o *UploadPKCS12) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *UploadPKCS12) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *UploadPKCS12) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.


### GetPassword

`func (o *UploadPKCS12) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UploadPKCS12) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UploadPKCS12) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UploadPKCS12) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetSplitLevel

`func (o *UploadPKCS12) GetSplitLevel() int64`

GetSplitLevel returns the SplitLevel field if non-nil, zero value otherwise.

### GetSplitLevelOk

`func (o *UploadPKCS12) GetSplitLevelOk() (*int64, bool)`

GetSplitLevelOk returns a tuple with the SplitLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitLevel

`func (o *UploadPKCS12) SetSplitLevel(v int64)`

SetSplitLevel sets SplitLevel field to given value.

### HasSplitLevel

`func (o *UploadPKCS12) HasSplitLevel() bool`

HasSplitLevel returns a boolean if a field has been set.

### GetTag

`func (o *UploadPKCS12) GetTag() []string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *UploadPKCS12) GetTagOk() (*[]string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *UploadPKCS12) SetTag(v []string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *UploadPKCS12) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetToken

`func (o *UploadPKCS12) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UploadPKCS12) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UploadPKCS12) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UploadPKCS12) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UploadPKCS12) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UploadPKCS12) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UploadPKCS12) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UploadPKCS12) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetUsername

`func (o *UploadPKCS12) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UploadPKCS12) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UploadPKCS12) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UploadPKCS12) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



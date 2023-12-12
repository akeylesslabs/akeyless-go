# UpdateClassicKeyCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertFileData** | Pointer to **string** | PEM Certificate in a Base64 format. Used for updating RSA keys&#39; certificates. | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | ClassicKey name | 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 

## Methods

### NewUpdateClassicKeyCertificate

`func NewUpdateClassicKeyCertificate(name string, ) *UpdateClassicKeyCertificate`

NewUpdateClassicKeyCertificate instantiates a new UpdateClassicKeyCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClassicKeyCertificateWithDefaults

`func NewUpdateClassicKeyCertificateWithDefaults() *UpdateClassicKeyCertificate`

NewUpdateClassicKeyCertificateWithDefaults instantiates a new UpdateClassicKeyCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertFileData

`func (o *UpdateClassicKeyCertificate) GetCertFileData() string`

GetCertFileData returns the CertFileData field if non-nil, zero value otherwise.

### GetCertFileDataOk

`func (o *UpdateClassicKeyCertificate) GetCertFileDataOk() (*string, bool)`

GetCertFileDataOk returns a tuple with the CertFileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFileData

`func (o *UpdateClassicKeyCertificate) SetCertFileData(v string)`

SetCertFileData sets CertFileData field to given value.

### HasCertFileData

`func (o *UpdateClassicKeyCertificate) HasCertFileData() bool`

HasCertFileData returns a boolean if a field has been set.

### GetJson

`func (o *UpdateClassicKeyCertificate) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *UpdateClassicKeyCertificate) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *UpdateClassicKeyCertificate) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *UpdateClassicKeyCertificate) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *UpdateClassicKeyCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClassicKeyCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClassicKeyCertificate) SetName(v string)`

SetName sets Name field to given value.


### GetToken

`func (o *UpdateClassicKeyCertificate) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *UpdateClassicKeyCertificate) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *UpdateClassicKeyCertificate) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *UpdateClassicKeyCertificate) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *UpdateClassicKeyCertificate) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *UpdateClassicKeyCertificate) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *UpdateClassicKeyCertificate) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *UpdateClassicKeyCertificate) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)



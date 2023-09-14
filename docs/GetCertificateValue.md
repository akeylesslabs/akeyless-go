# GetCertificateValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertIssuerName** | Pointer to **string** | The parent PKI Certificate Issuer&#39;s name of the certificate, required when used with display-id and token | [optional] 
**CertificateFileOutput** | Pointer to **string** | File to write the certificates to. | [optional] 
**DisplayId** | Pointer to **string** | Certificate display ID | [optional] 
**IgnoreCache** | Pointer to **string** | Retrieve the Secret value without checking the Gateway&#39;s cache [true/false]. This flag is only relevant when using the RestAPI | [optional] [default to "false"]
**IssuanceToken** | Pointer to **string** | Token for getting the issued certificate | [optional] 
**Json** | Pointer to **bool** | Set output format to JSON | [optional] [default to false]
**Name** | **string** | Certificate name | [default to "dummy_certificate_name"]
**PrivateKeyFileOutput** | Pointer to **string** | File to write the private key to. | [optional] 
**Token** | Pointer to **string** | Authentication token (see &#x60;/auth&#x60; and &#x60;/configure&#x60;) | [optional] 
**UidToken** | Pointer to **string** | The universal identity token, Required only for universal_identity authentication | [optional] 
**Version** | Pointer to **int32** | Certificate version | [optional] 

## Methods

### NewGetCertificateValue

`func NewGetCertificateValue(name string, ) *GetCertificateValue`

NewGetCertificateValue instantiates a new GetCertificateValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCertificateValueWithDefaults

`func NewGetCertificateValueWithDefaults() *GetCertificateValue`

NewGetCertificateValueWithDefaults instantiates a new GetCertificateValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertIssuerName

`func (o *GetCertificateValue) GetCertIssuerName() string`

GetCertIssuerName returns the CertIssuerName field if non-nil, zero value otherwise.

### GetCertIssuerNameOk

`func (o *GetCertificateValue) GetCertIssuerNameOk() (*string, bool)`

GetCertIssuerNameOk returns a tuple with the CertIssuerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertIssuerName

`func (o *GetCertificateValue) SetCertIssuerName(v string)`

SetCertIssuerName sets CertIssuerName field to given value.

### HasCertIssuerName

`func (o *GetCertificateValue) HasCertIssuerName() bool`

HasCertIssuerName returns a boolean if a field has been set.

### GetCertificateFileOutput

`func (o *GetCertificateValue) GetCertificateFileOutput() string`

GetCertificateFileOutput returns the CertificateFileOutput field if non-nil, zero value otherwise.

### GetCertificateFileOutputOk

`func (o *GetCertificateValue) GetCertificateFileOutputOk() (*string, bool)`

GetCertificateFileOutputOk returns a tuple with the CertificateFileOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFileOutput

`func (o *GetCertificateValue) SetCertificateFileOutput(v string)`

SetCertificateFileOutput sets CertificateFileOutput field to given value.

### HasCertificateFileOutput

`func (o *GetCertificateValue) HasCertificateFileOutput() bool`

HasCertificateFileOutput returns a boolean if a field has been set.

### GetDisplayId

`func (o *GetCertificateValue) GetDisplayId() string`

GetDisplayId returns the DisplayId field if non-nil, zero value otherwise.

### GetDisplayIdOk

`func (o *GetCertificateValue) GetDisplayIdOk() (*string, bool)`

GetDisplayIdOk returns a tuple with the DisplayId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayId

`func (o *GetCertificateValue) SetDisplayId(v string)`

SetDisplayId sets DisplayId field to given value.

### HasDisplayId

`func (o *GetCertificateValue) HasDisplayId() bool`

HasDisplayId returns a boolean if a field has been set.

### GetIgnoreCache

`func (o *GetCertificateValue) GetIgnoreCache() string`

GetIgnoreCache returns the IgnoreCache field if non-nil, zero value otherwise.

### GetIgnoreCacheOk

`func (o *GetCertificateValue) GetIgnoreCacheOk() (*string, bool)`

GetIgnoreCacheOk returns a tuple with the IgnoreCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreCache

`func (o *GetCertificateValue) SetIgnoreCache(v string)`

SetIgnoreCache sets IgnoreCache field to given value.

### HasIgnoreCache

`func (o *GetCertificateValue) HasIgnoreCache() bool`

HasIgnoreCache returns a boolean if a field has been set.

### GetIssuanceToken

`func (o *GetCertificateValue) GetIssuanceToken() string`

GetIssuanceToken returns the IssuanceToken field if non-nil, zero value otherwise.

### GetIssuanceTokenOk

`func (o *GetCertificateValue) GetIssuanceTokenOk() (*string, bool)`

GetIssuanceTokenOk returns a tuple with the IssuanceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceToken

`func (o *GetCertificateValue) SetIssuanceToken(v string)`

SetIssuanceToken sets IssuanceToken field to given value.

### HasIssuanceToken

`func (o *GetCertificateValue) HasIssuanceToken() bool`

HasIssuanceToken returns a boolean if a field has been set.

### GetJson

`func (o *GetCertificateValue) GetJson() bool`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *GetCertificateValue) GetJsonOk() (*bool, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *GetCertificateValue) SetJson(v bool)`

SetJson sets Json field to given value.

### HasJson

`func (o *GetCertificateValue) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetName

`func (o *GetCertificateValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCertificateValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCertificateValue) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateKeyFileOutput

`func (o *GetCertificateValue) GetPrivateKeyFileOutput() string`

GetPrivateKeyFileOutput returns the PrivateKeyFileOutput field if non-nil, zero value otherwise.

### GetPrivateKeyFileOutputOk

`func (o *GetCertificateValue) GetPrivateKeyFileOutputOk() (*string, bool)`

GetPrivateKeyFileOutputOk returns a tuple with the PrivateKeyFileOutput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFileOutput

`func (o *GetCertificateValue) SetPrivateKeyFileOutput(v string)`

SetPrivateKeyFileOutput sets PrivateKeyFileOutput field to given value.

### HasPrivateKeyFileOutput

`func (o *GetCertificateValue) HasPrivateKeyFileOutput() bool`

HasPrivateKeyFileOutput returns a boolean if a field has been set.

### GetToken

`func (o *GetCertificateValue) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetCertificateValue) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetCertificateValue) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetCertificateValue) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUidToken

`func (o *GetCertificateValue) GetUidToken() string`

GetUidToken returns the UidToken field if non-nil, zero value otherwise.

### GetUidTokenOk

`func (o *GetCertificateValue) GetUidTokenOk() (*string, bool)`

GetUidTokenOk returns a tuple with the UidToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUidToken

`func (o *GetCertificateValue) SetUidToken(v string)`

SetUidToken sets UidToken field to given value.

### HasUidToken

`func (o *GetCertificateValue) HasUidToken() bool`

HasUidToken returns a boolean if a field has been set.

### GetVersion

`func (o *GetCertificateValue) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetCertificateValue) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetCertificateValue) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetCertificateValue) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


